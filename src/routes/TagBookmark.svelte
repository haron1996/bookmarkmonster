<script lang="ts">
	import {
		apiHost,
		bookmarkTags,
		matchedTagsFromDB,
		selectedBookmarks,
		tagName,
		tags,
		error
	} from '../stores/stores';
	import { session } from '../stores/stores';
	import type { Tag } from '../types/tag';
	import { hideTagBookmarkComp } from '../utils/hideTagBookmarkComp';
	import { removeError } from '../utils/removeError';

	let tagid: string = '';
	let savingChanges: boolean = false;

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

	const handleTagInput = async () => {
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
	};

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

	async function handleTagFormSubmit() {
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
		savingChanges = true;

		setTimeout(() => {
			savingChanges = false;
		}, 3000);

		setTimeout(() => {
			hideTagBookmarkComp();
		}, 3500);
	}

	$: $error,
		$error === ''
			? () => {}
			: setTimeout(() => {
					removeError();
			  }, 5000);
</script>

<div class="wrapper" id="tagBookmarkComponent" role="none" on:click={hideTagBookmarkComp}>
	<div class="card" role="none" on:click|stopPropagation={() => {}}>
		<div class="top">
			<p>Tag bookmark</p>
			<i class="las la-times" />
		</div>
		<div class="currentTags">
			{#if $bookmarkTags}
				{#each $bookmarkTags as { added, deleted, id, name, updated, user_id }}
					<div class="tag" data-id={id} data-userid={user_id} data-name={name}>
						<i class="las la-hashtag" />
						<span>{name}</span>
						<i class="las la-times" role="none" on:click={deleteTagFromBookmark} />
					</div>
				{/each}
			{/if}
		</div>
		<form on:submit|preventDefault={handleTagFormSubmit}>
			<input
				type="text"
				name="tag"
				id="tag"
				placeholder="Type and enter to add tag"
				autocomplete="off"
				spellcheck="false"
				bind:value={$tagName}
				on:input|stopPropagation={handleTagInput}
			/>
			<div class="matchingTags">
				{#if $matchedTagsFromDB}
					{#each $matchedTagsFromDB as { added, deleted, id, name, updated, user_id }}
						<div
							class="tag"
							data-name={name}
							data-id={id}
							data-userid={user_id}
							role="none"
							on:click={handleClickOnMatchingTag}
						>
							<i class="las la-hashtag" />
							<span>{name}</span>
						</div>
					{/each}
				{/if}
			</div>
		</form>
		<div class="buttons">
			<button
				type="submit"
				on:click|stopPropagation|preventDefault={saveChanges}
				class:savingChanges={savingChanges || $tagName !== ''}
			>
				<span>{savingChanges ? `Updating...` : 'Update'}</span>
			</button>
			<button on:click|preventDefault|stopPropagation={hideTagBookmarkComp}>
				<span>Cancel</span>
			</button>
		</div>
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
			width: 40rem;
			min-height: max-content;
			padding: 1em;
			background-color: rgb(255, 255, 255);
			box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
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
				}

				i {
					font-size: 1.5rem;
				}
			}

			.currentTags {
				display: flex;
				align-items: center;
				flex-flow: row wrap;
				gap: 1em;

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

			form {
				min-height: max-content;
				width: 100%;
				border-radius: 0.3rem;
				position: relative;

				input {
					padding: 0.5em;
					border-radius: inherit;
					border: 0.1rem solid rgb(0, 0, 0, 0.2);
					width: 100%;
					font-family: 'Arial CE', sans-serif;

					&:focus {
						outline-color: rgb(96, 1, 255);
					}
				}

				.matchingTags {
					position: absolute;
					left: 0;
					right: 0;
					width: inherit;
					top: 100%;
					min-height: max-content;
					max-height: 20rem;
					overflow-y: auto;
					box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
					z-index: 3;
					background-color: rgb(245, 245, 245);

					.tag {
						width: 100%;
						min-height: 4rem;
						display: flex;
						align-items: center;
						gap: 0.5em;
						padding: 0.5em;
						border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
						cursor: pointer;
						transition: all ease 0.3s;

						i {
							font-size: 1.5rem;
						}

						span {
							font-family: 'Arial CE', sans-serif;
							font-size: 1.3rem;
						}

						&:hover {
							background-color: rgb(96, 1, 255, 0.1);
						}
					}
				}
			}

			.buttons {
				display: flex;
				align-items: center;
				justify-content: space-between;

				button {
					border: none;
					padding: 0.5em 1em;
					border-radius: 0.3rem;
					cursor: pointer;

					span {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
					}
				}

				button[type='submit'] {
					background-color: rgb(96, 1, 255);
					box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

					span {
						color: rgb(255, 255, 255);
					}
				}

				.savingChanges {
					pointer-events: none;
					opacity: 0.5;
				}
			}
		}
	}
</style>
