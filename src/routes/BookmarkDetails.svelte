<script lang="ts">
	import {
		apiHost,
		bookmarkTags,
		matchedTagsFromDB,
		selectedBookmarks,
		session,
		tagName,
		tags,
		error,
		bookmarks
	} from '../stores/stores';
	import { hideBookmarkDetails } from '../utils/hideBookmarkDetails';
	import type { Tag } from '../types/tag';
	import type { Bookmark } from '../types/bookmark';

	let processing: boolean = false;
	let tagid: string = '';

	// function selectElementContents() {
	// 	const el = document.getElementById('addTag') as HTMLDivElement | null;

	// 	if (el === null) return;

	// 	var range: Range = document.createRange();

	// 	range.selectNodeContents(el);

	// 	var sel: Selection | null = window.getSelection();

	// 	if (sel === null) return;

	// 	sel.removeAllRanges();

	// 	sel.addRange(range);
	// }

	function handleContentEditableBlur(e: FocusEvent) {
		const target = e.currentTarget as HTMLElement;

		const div = target.closest('.addTag') as HTMLDivElement | null;

		if (div === null) return;

		console.log(div.textContent);
	}

	async function fetchMatchingTags() {
		if ($tagName === '') {
			matchedTagsFromDB.set([]);
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/${$tagName}`, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		const tgs: Tag[] = result[0];

		matchedTagsFromDB.set(tgs);
	}

	async function handleClickOnMatchingTag(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const div = target.closest('.tag') as HTMLDivElement | null;

		if (div === null) return;

		const tag: Tag = { id: div.dataset.id, name: div.dataset.name, user_id: div.dataset.userid };

		let tgs: Tag[] = [];

		tgs = [tag, ...tgs];

		if (
			$bookmarkTags
				.map((value) => {
					return value.id;
				})
				.includes(tag.id)
		) {
			// ignore
			matchedTagsFromDB.set([]);
			tagName.set('');
		} else {
			// add tag to bookmark tags
			const response = await fetch(`${$apiHost}/authenticated/bookmarks/tagBookmark`, {
				method: 'PATCH',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${$session.AccessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer',
				body: JSON.stringify({ bookmark_id: $selectedBookmarks[0].id, tags: tgs })
			});

			const result = await response.json();

			bookmarkTags.update((values) => [result[0], ...values]);

			matchedTagsFromDB.set([]);

			tagName.set('');
		}
	}

	async function tagBookmark(tgs: Tag[]) {
		const response = await fetch(`${$apiHost}/authenticated/bookmarks/tagBookmark`, {
			method: 'PATCH',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ bookmark_id: $selectedBookmarks[0].id, tags: tgs })
		});
		const result = await response.json();
		bookmarkTags.update((values) => [result[0], ...values]);
		matchedTagsFromDB.set([]);
		tagName.set('');
	}

	async function createTagAndTagBookmarkWithIt() {
		let response = await fetch(`${$apiHost}/authenticated/tags/create-tag`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ name: $tagName })
		});
		let result = await response.json();
		tags.update((values) => [result[0], ...values]);
		let tgs: Tag[] = [];
		tgs = [result[0], ...tgs];
		tagBookmark(tgs);
	}

	async function handleTagSubmit() {
		if (
			$matchedTagsFromDB.filter((value) => {
				return value.name === $tagName;
			})[0] !== undefined
		) {
			const t = $matchedTagsFromDB.filter((value) => {
				return value.name === $tagName;
			});

			if (
				$bookmarkTags.filter((bt) => {
					return bt.name === t[0].name;
				})[0] !== undefined
			) {
				matchedTagsFromDB.set([]);

				tagName.set('');
			} else {
				let tgs: Tag[] = [];

				tgs = [t[0], ...tgs];

				tagBookmark(tgs);
			}
		} else {
			createTagAndTagBookmarkWithIt();
		}
	}

	function saveChanges() {
		processing = true;

		setTimeout(() => {
			processing = false;
		}, 3000);

		setTimeout(() => {
			hideBookmarkDetails();
		}, 3500);
	}

	function getTagId(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('.tag') as HTMLElement | null;

		if (tag === null) return;

		if (tag.dataset.id) {
			tagid = tag.dataset.id;
		}
	}

	async function deleteTagFromBookmark(e: MouseEvent) {
		getTagId(e);

		const response = await fetch(`${$apiHost}/authenticated/bookmarks/deleteTagFromBookmark`, {
			method: 'DELETE',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				bookmarkid: $selectedBookmarks[0].id,
				tagid: tagid
			})
		});

		const result = await response.json();

		if (result.message === 'bookmark must have at least one tag') {
			error.set(result.message);
			return;
		}

		if (result.message === 'something went wrong') {
			error.set(result.message);
			return;
		}

		const tag: Tag = result[0];

		$bookmarkTags = $bookmarkTags.filter((value) => {
			return value.id !== tag.id;
		});
	}
	async function renameBookmark() {
		if ($selectedBookmarks[0].title === '') return;

		const response = await fetch(`${$apiHost}/authenticated/bookmarks/renameBookmark`, {
			method: 'PATCH',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ title: $selectedBookmarks[0].title, id: $selectedBookmarks[0].id })
		});

		const result = await response.json();

		const b: Bookmark = result[0];

		const index: number = $bookmarks.findIndex((value) => {
			return value.id === b.id;
		});

		$bookmarks.splice(index, 1, b);

		bookmarks.set($bookmarks);
	}
</script>

<div
	class="wrapper"
	id="bookmarkDetails"
	role="none"
	on:click|stopPropagation={hideBookmarkDetails}
>
	<div class="card" role="none" on:click|stopPropagation={() => {}}>
		{#if $selectedBookmarks[0]}
			<div class="top">
				<p>Bookmark details</p>
				<i class="las la-times" role="none" on:click|stopPropagation={hideBookmarkDetails} />
			</div>
			<div class="detail">
				<p>Title</p>
				<input
					type="text"
					name="bookmarkTitle"
					id="bookmarkTitle"
					spellcheck="false"
					autocomplete="off"
					bind:value={$selectedBookmarks[0].title}
					on:input={renameBookmark}
				/>
			</div>
			<div class="detail">
				<p>Notes</p>

				{#if $selectedBookmarks[0].notes}
					<span>{$selectedBookmarks[0].notes}</span>
				{/if}
			</div>
			<div class="detail">
				<p>URL</p>
				<span>{$selectedBookmarks[0].bookmark}</span>
			</div>
			<div class="detail">
				<p>Added</p>
				<span
					>{$selectedBookmarks[0].added?.getDate()} / {$selectedBookmarks[0].added?.getMonth()} / {$selectedBookmarks[0].added?.getFullYear()}</span
				>
			</div>
			<div class="detail">
				<p>Updated</p>
				{#if $selectedBookmarks[0].updated}
					<span
						>{$selectedBookmarks[0].updated?.getDate()} / {$selectedBookmarks[0].updated?.getMonth()}
						/
						{$selectedBookmarks[0].updated?.getFullYear()}</span
					>
				{/if}
			</div>
			<div class="detail">
				<p>Tags</p>
				<div class="tags">
					{#if $bookmarkTags}
						{#each $bookmarkTags as { added, deleted, id, name, updated, user_id }}
							<div class="tag" data-id={id} data-userid={user_id} data-name={name}>
								<i class="las la-hashtag" />
								<span>{name}</span>
								<i class="las la-times" role="none" on:click={deleteTagFromBookmark} />
							</div>
						{/each}
					{/if}
					<form class="addTag" on:submit|preventDefault={handleTagSubmit}>
						<i class="las la-hashtag" />
						<input
							type="text"
							name="tagInput"
							id="tagInput"
							placeholder="add tag"
							spellcheck="false"
							autocomplete="off"
							bind:value={$tagName}
							on:input={fetchMatchingTags}
						/>
						<div class="matchingTags">
							{#each $matchedTagsFromDB as { added, deleted, id, name, updated, user_id }}
								<div
									class="tag"
									role="none"
									data-id={id}
									data-name={name}
									data-userid={user_id}
									on:click={handleClickOnMatchingTag}
								>
									<i class="las la-hashtag" />
									<span>{name}</span>
								</div>
							{/each}
						</div>
					</form>
				</div>
			</div>
		{/if}
		<button on:click|stopPropagation|preventDefault={saveChanges}>
			{#if processing}
				<span>Updating....</span>
			{:else}
				<span>Update and Close</span>
			{/if}
		</button>
	</div>
</div>

<style lang="scss">
	.wrapper {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		z-index: 1;
		background-color: rgb(0, 0, 0, 0.6);
		display: flex;
		align-items: center;
		justify-content: center;
		display: none;

		.card {
			width: 60rem;
			min-height: max-content;
			box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
			background-color: rgb(255, 255, 255);
			padding: 1em;
			border-radius: 0.3rem;
			display: flex;
			flex-direction: column;
			gap: 2em;

			@media only screen and (max-width: 425px) {
				max-width: 98%;
			}

			.top {
				display: flex;
				align-items: center;
				justify-content: space-between;

				p {
					font-size: 1.5rem;
					font-family: 'Arial CE', sans-serif;
					font-weight: 600;
				}

				i {
					font-size: 1.5rem;
				}
			}

			.detail {
				width: 100%;
				display: flex;
				flex-direction: column;
				gap: 0.5em;

				input[type='text'] {
					width: 100%;
					border: none;
					outline: none;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
				}

				p {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					font-weight: 600;
				}

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
				}
			}

			div:has(div.tags) {
				display: flex;

				.tags {
					display: flex;
					align-items: center;
					flex-flow: row wrap;
					gap: 0.5em;

					form.addTag {
						min-width: max-content;
						padding: 0.5em;
						position: relative;
						display: flex;
						align-items: center;
						gap: 0.5em;
						background-color: #dde6ed;
						z-index: 3;

						i {
							font-size: 1.5rem;
							background-color: #8cc0de;
							color: rgb(255, 255, 255);
						}

						input {
							outline: none;
							border: none;
							background-color: inherit;
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							text-transform: lowercase;
							width: max-content;
						}

						.matchingTags {
							position: absolute;
							top: 100%;
							left: 0;
							right: 0;
							width: inherit;
							display: flex;
							flex-direction: column;
							background-color: #dde6ed;
							min-height: max-content;
							overflow-y: auto;
							max-width: 40rem;
							box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;

							.tag {
								padding: 1em 0.5em;
								cursor: pointer;
								display: flex;
								align-items: center;
								gap: 0.5em;
								width: 100%;

								span {
									font-size: 1.3rem;
									font-family: 'Arial CE', sans-serif;
									text-transform: lowercase;
									color: rgb(0, 0, 0, 0.8);
								}
							}
						}
					}

					.tag {
						width: max-content;
						padding: 0.5em;
						background-color: #dde6ed;
						display: flex;
						align-items: center;
						gap: 0.5em;

						i.la-hashtag {
							font-size: 1.5rem;
							background-color: #8cc0de;
							color: rgb(255, 255, 255);
						}

						span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
						}

						i.la-times {
							font-size: 1.5rem;
							cursor: pointer;
						}
					}
				}
			}

			button {
				border: none;
				padding: 0.5em 1em;
				width: max-content;
				border-radius: 0.3rem;
				cursor: pointer;
				background-color: #6528f7;
				box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
					color: rgb(255, 255, 255);
				}
			}
		}
	}
</style>
