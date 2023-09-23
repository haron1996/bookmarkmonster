<script lang="ts">
	import { browser } from '$app/environment';
	import {
		apiHost,
		currentTag,
		matchedTagsFromDB,
		session,
		showCreatBookmarkFromTagPage,
		successMessage,
		tags
	} from '../../../stores/stores';
	import type { Tag } from '../../../types/tag';

	let bookmark_url: string = '';
	let tagName: string = '';
	let selectedTags: Tag[] = [];

	async function createTag() {
		if (tagName === 'Untitled tag' || tagName === '') {
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/create-tag`, {
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
			body: JSON.stringify({ name: tagName })
		});

		const result = await response.json();

		tags.update((values) => [result[0], ...values]);

		//setCurrentTag();

		selectedTags = [result[0], ...selectedTags];

		selectedTags = selectedTags;

		tagName = '';

		matchedTagsFromDB.set([]);

		setTimeout(() => {
			addActiveClassToCurrentTag();
		}, 100);
	}

	function handleTagInputKeyDown(e: KeyboardEvent) {
		if (e.code === 'Space' && tagName === '') {
			e.preventDefault();
			return;
		}

		if (e.code === 'Enter') {
			const t = $matchedTagsFromDB.filter((value) => {
				return value.name === tagName;
			});

			if (t.length === 1) {
				const tag = t[0];

				const tgs = selectedTags.filter((value) => {
					return value.id === tag.id;
				});

				if (tgs.length === 1) {
					// ignore
					tagName = '';
					matchedTagsFromDB.set([]);
				} else {
					// add
					selectedTags = [tag, ...selectedTags];
					selectedTags = selectedTags;
					tagName = '';
					matchedTagsFromDB.set([]);
				}
			} else {
				createTag();
			}
		}
	}

	async function handleTagInput() {
		if (tagName === '') {
			matchedTagsFromDB.set($tags);
		} else {
			searchUserTags();
		}
	}

	async function searchUserTags() {
		const response = await fetch(`${$apiHost}/authenticated/tags/${tagName}`, {
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

		const msg = result.message;

		if (msg) {
			alert(msg);
			return;
		}

		matchedTagsFromDB.set(result[0]);
	}

	function selectTag(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('span') as HTMLSpanElement | null;

		if (tag === null) return;

		const t: Tag = {
			added: tag.dataset.added,
			deleted: tag.dataset.deleted,
			id: tag.dataset.id,
			name: tag.dataset.name,
			updated: tag.dataset.updated,
			user_id: tag.dataset.userid
		};

		const et = selectedTags.filter((value) => {
			return value.id === t.id;
		});

		if (et.length < 1) {
			selectedTags = [t, ...selectedTags];
			selectedTags = selectedTags;
		} else {
			tagName = '';
			matchedTagsFromDB.set([]);
		}

		tagName = '';

		matchedTagsFromDB.set([]);
	}

	function removeTagFromSelectedTags(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('div.tag') as HTMLSpanElement | null;

		if (tag === null) return;

		const t: Tag = {
			added: tag.dataset.added,
			deleted: tag.dataset.deleted,
			id: tag.dataset.id,
			name: tag.dataset.name,
			updated: tag.dataset.updated,
			user_id: tag.dataset.userid
		};

		selectedTags = selectedTags.filter((value) => {
			return value.id !== t.id;
		});

		selectedTags = selectedTags;
	}

	function addActiveClassToCurrentTag() {
		const domTags = document.querySelectorAll('.tag') as NodeListOf<HTMLDivElement>;

		if (domTags.length < 1) return;

		domTags.forEach((dt) => {
			dt.classList.remove('activeTag');
			if (dt.dataset.id === $currentTag.id) {
				dt.classList.add('activeTag');
			}
		});
	}

	async function createAndTagBookmark() {
		if (bookmark_url === '' || selectedTags.length < 1) return;

		if (tagName !== '') {
			const t = $matchedTagsFromDB.filter((value) => {
				return value.name === tagName;
			});

			if (t.length === 1) {
				const tag = t[0];

				const tgs = selectedTags.filter((value) => {
					return value.id === tag.id;
				});

				if (tgs.length === 1) {
					// ignore
					tagName = '';
					matchedTagsFromDB.set([]);
				} else {
					// add
					selectedTags = [tag, ...selectedTags];
					selectedTags = selectedTags;
					tagName = '';
					matchedTagsFromDB.set([]);
				}
			} else {
				createTag();
			}
			return;
		}

		showCreatBookmarkFromTagPage.set(false);

		const response = await fetch(`${$apiHost}/authenticated/bookmarks/createAndTagBookmark`, {
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
			body: JSON.stringify({
				bookmark: bookmark_url,
				tags: selectedTags.map((v) => {
					return v.id;
				})
			})
		});

		const result = await response.json();

		console.log(result[0]);

		successMessage.set('Success! Your link has been saved');

		tagName = '';

		selectedTags = [];

		matchedTagsFromDB.set([]);

		bookmark_url = '';
	}
</script>

{#if $showCreatBookmarkFromTagPage}
	<div
		class="container"
		role="none"
		on:click={() => {
			tagName = '';
			matchedTagsFromDB.set([]);
			showCreatBookmarkFromTagPage.set(false);
			selectedTags = [];
		}}
	>
		<form
			role="none"
			on:click|stopPropagation={() => {
				matchedTagsFromDB.set([]);
			}}
			on:submit|preventDefault
		>
			<div class="top">
				<span>Add bookmark</span>
				<i
					class="las la-times"
					role="none"
					on:click|stopPropagation={() => {
						tagName = '';
						matchedTagsFromDB.set([]);
						showCreatBookmarkFromTagPage.set(false);
						selectedTags = [];
					}}
				/>
			</div>
			<div class="inputs">
				<input
					type="url"
					name="url"
					id="url"
					placeholder="https://example.com"
					pattern="https://.*"
					bind:value={bookmark_url}
				/>
				<div class="chooseTag">
					{#if selectedTags.length >= 1}
						<div class="selectedTags">
							{#each selectedTags as { added, deleted, id, name, updated, user_id }}
								<div
									class="tag"
									data-added={added}
									data-deleted={deleted}
									data-id={id}
									data-updated={updated}
									data-userid={user_id}
									data-name={name}
									role="none"
									on:click|stopPropagation
								>
									<span>{name}</span>
									<i class="las la-times" role="none" on:click={removeTagFromSelectedTags} />
								</div>
							{/each}
						</div>
					{/if}
					<input
						type="text"
						id="tag"
						class="tag"
						placeholder="Select one or multiple tags"
						autocomplete="off"
						spellcheck="false"
						bind:value={tagName}
						on:keydown={handleTagInputKeyDown}
						on:pointerdown={handleTagInput}
						on:input={searchUserTags}
						on:click|stopPropagation
					/>

					{#if $matchedTagsFromDB.length >= 1}
						<div class="tags" role="none" on:click|stopPropagation>
							{#each $matchedTagsFromDB as { added, deleted, id, name, updated, user_id }}
								<span
									data-added={added}
									data-deleted={deleted}
									data-id={id}
									data-updated={updated}
									data-userid={user_id}
									data-name={name}
									role="none"
									on:click|stopPropagation={selectTag}>{name}</span
								>
							{/each}
						</div>
					{/if}
				</div>
			</div>
			<div class="buttons">
				<button type="submit" on:click|preventDefault|stopPropagation={createAndTagBookmark}>
					Add bookmark
				</button>
				<button
					on:click|preventDefault|stopPropagation={() => {
						tagName = '';
						matchedTagsFromDB.set([]);
						showCreatBookmarkFromTagPage.set(false);
						selectedTags = [];
					}}>Cancel</button
				>
			</div>
		</form>
	</div>
{/if}

<style lang="scss">
	.container {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-color: rgb(0, 0, 0, 0.4);
		z-index: 10;
		display: flex;
		align-items: center;
		justify-content: center;

		form {
			width: 50rem;
			min-height: max-content;
			padding: 1em;
			background-color: white;
			box-shadow: rgba(0, 0, 0, 0.07) 0px 1px 2px, rgba(0, 0, 0, 0.07) 0px 2px 4px,
				rgba(0, 0, 0, 0.07) 0px 4px 8px, rgba(0, 0, 0, 0.07) 0px 8px 16px,
				rgba(0, 0, 0, 0.07) 0px 16px 32px, rgba(0, 0, 0, 0.07) 0px 32px 64px;
			border-radius: 0.3rem;
			display: flex;
			flex-direction: column;
			gap: 3em;

			.top {
				display: flex;
				align-items: center;
				justify-content: space-between;

				span {
					font-size: 1.5rem;
					font-family: 'Arial CE', sans-serif;
				}

				i {
					font-size: 1.5rem;
					cursor: pointer;
				}
			}

			.inputs {
				width: 100%;
				display: flex;
				flex-direction: column;
				min-height: max-content;
				gap: 2em;

				input {
					width: 100%;
					height: 3.5rem;
					padding: 0.5em;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					border-radius: 0.3rem;
					outline: none;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
				}

				.chooseTag {
					position: relative;
					width: 100%;
					height: max-content;
					display: flex;
					flex-direction: column;
					gap: 1em;

					.selectedTags {
						display: flex;
						flex-flow: row wrap;
						align-content: start;
						gap: 1em;
						padding: 0.5em;
						max-height: 20rem;
						overflow-y: auto;

						.tag {
							display: flex;
							align-items: center;
							gap: 1em;
							padding: 0.5em 0.8em;
							background-color: #cee6f3;
							border-radius: 0.3rem;

							span {
								font-family: 'Arial CE', sans-serif;
								font-size: 1.3rem;
							}

							i {
								font-size: 1.5rem;
								cursor: pointer;
							}
						}
					}

					.tags {
						position: absolute;
						top: 100%;
						left: 0;
						width: inherit;
						height: inherit;
						height: max-content;
						max-height: 25rem; // change to max content // set 25rem to max height
						overflow-y: auto;
						border-radius: inherit;
						box-shadow: rgba(0, 0, 0, 0.07) 0px 1px 2px, rgba(0, 0, 0, 0.07) 0px 2px 4px,
							rgba(0, 0, 0, 0.07) 0px 4px 8px, rgba(0, 0, 0, 0.07) 0px 8px 16px,
							rgba(0, 0, 0, 0.07) 0px 16px 32px, rgba(0, 0, 0, 0.07) 0px 32px 64px;
						padding: 1em;
						background-color: white;
						display: flex;
						flex-flow: row wrap;
						align-content: start;
						gap: 1em;

						span {
							padding: 0.3em 0.8em;
							font-family: 'Arial CE', sans-serif;
							font-size: 1.3rem;
							background-color: #cee6f3;
							border-radius: 0.3rem;
							cursor: pointer;
						}
					}
				}
			}

			.buttons {
				display: flex;
				align-items: center;
				justify-content: flex-end;
				gap: 0.7em;
				border-radius: 0.2rem;

				button {
					padding: 0.5em 1em;
					border: none;
					cursor: pointer;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					background-color: rgb(4, 13, 18);
					color: white;
				}

				button:last-of-type {
					background-color: #f8f1f1;
					color: rgb(4, 13, 18);
				}
			}

			// only form
			@media only screen and (width <= 560px) {
				width: 95%;
			}
		}
	}
</style>
