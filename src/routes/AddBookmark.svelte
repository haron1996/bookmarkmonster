<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import {
		apiHost,
		bookmarks,
		currentTagID,
		error,
		lastAddedBookmark,
		processingBookmark,
		session,
		tags
	} from '../stores/stores';
	import type { Bookmark } from '../types/bookmark';
	import type { Tag } from '../types/tag';
	import { hideAddBookmarkComponent } from '../utils/hideAddBookmarkComponent';
	import { hideOverlay } from '../utils/hideOverlay';

	let bookmark: string = '';
	let tagName: string = '';
	let tag: Tag = {};
	let matchingTags: Tag[] = [];
	let selectedTags: Tag[] = [];

	const showDefaultURLs = () => {
		const defaultURLS = document.getElementById('defaultURLs') as HTMLDivElement | null;

		if (defaultURLS === null) return;

		defaultURLS.style.display = 'none';
	};

	const hideDefaultURLs = () => {
		const defaultURLS = document.getElementById('defaultURLs') as HTMLDivElement | null;

		if (defaultURLS === null) return;

		defaultURLS.style.display = 'none';
	};

	const getAllUserTagsIfInputEmpty = async () => {
		if (tagName != '') return;

		const sessionString = localStorage.getItem('session') as string | null;

		if (sessionString === null) return;

		const response = await fetch(`${$apiHost}/authenticated/tags`, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${JSON.parse(sessionString).AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		matchingTags = result[0];
	};

	const getUserMatchingTags = async () => {
		if (tagName === '') {
			matchingTags = [];
			return;
		}

		const sessionString = localStorage.getItem('session') as string | null;

		if (sessionString === null) return;

		const response = await fetch(`${$apiHost}/authenticated/tags/${tagName}`, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${JSON.parse(sessionString).AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		matchingTags = result[0];
	};

	const handleTagsInputBlurEvent = () => {
		if (tagName === '') {
			matchingTags = [];
		}
	};

	const handleTagFormSubmit = async () => {
		if (tagName === '') {
			error.set('tag name must not be empty');
			return;
		}

		if (matchingTags.length > 0) {
			matchingTags.forEach((tag) => {
				if (tag.name === tagName) {
					if (
						selectedTags
							.map((t) => {
								return t.name;
							})
							.includes(tagName)
					) {
						tagName = '';

						return;
					} else {
						selectedTags = [...selectedTags, tag];

						tagName = '';

						return;
					}
				}
			});
		} else {
			const sessionString = localStorage.getItem('session') as string | null;

			if (sessionString === null) return;

			const response = await fetch(`${$apiHost}/authenticated/tags/create-tag`, {
				method: 'POST',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${JSON.parse(sessionString).AccessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer',
				body: JSON.stringify({ name: tagName })
			});

			if (response.status === 200) {
				const result = await response.json();

				const tag: Tag = result[0];

				selectedTags = [...selectedTags, tag];

				tags.update((tags) => [...tags, tag]);

				tagName = '';
			} else {
				console.log('could not create tag', response.status);
				tagName = '';
				return;
			}
		}
	};

	const handleClickOnMatchingTag = (e: MouseEvent) => {
		const clickTarget = e.currentTarget as HTMLElement;

		const clickedMatchingTag = clickTarget.closest('.tag') as HTMLDivElement | null;

		if (clickedMatchingTag === null) return;

		matchingTags.forEach((tag) => {
			if (tag.id === clickedMatchingTag.dataset.id) {
				if (
					selectedTags
						.map((t) => {
							return t.name;
						})
						.includes(tag.name)
				) {
					tagName = '';

					return;
				} else {
					selectedTags = [...selectedTags, tag];

					tagName = '';
				}
			}
		});
	};

	const removeTagFromSelectedTags = (e: MouseEvent) => {
		const clickTarget = e.currentTarget as HTMLElement;

		const clickedSelectedTag = clickTarget.closest('.tag') as HTMLDivElement | null;

		if (clickedSelectedTag === null) return;

		selectedTags.forEach((tag) => {
			if (tag.id === clickedSelectedTag.dataset.id) {
				selectedTags = selectedTags.filter((tag) => tag.id != clickedSelectedTag.dataset.id);
			}
		});
	};

	function isValidURL(url: string) {
		try {
			const urlFromString = new URL(url);
			if (urlFromString.protocol === 'https:') {
				return true;
			} else if (urlFromString.protocol === 'http:') {
				return false;
			}
		} catch (error) {
			return false;
		}
	}

	const handleAddBookmarkFormSubmit = async () => {
		// if (!isValidURL(bookmark)) {
		// 	console.log('enter a valid url');
		// 	return;
		// }

		if (bookmark === '') {
			error.set('url must not be empty');
			return;
		}

		processingBookmark.set(true);

		hideAddBookmarkComponent();

		const sessionString = localStorage.getItem('session') as string | null;

		if (sessionString === null) return;

		const response = await fetch(`${$apiHost}/authenticated/bookmarks/add`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${JSON.parse(sessionString).AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ bookmark: bookmark })
		});

		if (response.ok) {
			const result = await response.json();

			const createdBookmark: Bookmark = result[0];

			lastAddedBookmark.set(createdBookmark);

			bookmarks.update((bookmarks) => [...bookmarks, createdBookmark]);

			if ($currentTagID === 'all-tags') {
				const getUserBookmarks = async () => {
					const response = await fetch(`${$apiHost}/authenticated/bookmarks`, {
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

					//const b: Bookmark[] = result[0];

					bookmarks.set(result[0]);
				};

				await getUserBookmarks();
			} else {
				console.log($currentTagID);
				const allTagsDiv = document.getElementById('allTagsDiv') as HTMLDivElement | null;

				if (allTagsDiv === null) {
					processingBookmark.set(false);
					return;
				}

				allTagsDiv.click();
			}

			//todo
			// track created tag
			// go to that tag and get related bookmarks
		} else {
			console.log(response.status, response.statusText);
		}

		bookmark = '';

		selectedTags = [];

		processingBookmark.set(false);
	};

	function handleKeyDownInTagInput(e: KeyboardEvent) {
		if (e.code === 'Comma') {
			e.preventDefault();
			handleTagFormSubmit();
		}
	}
</script>

<div class="container" id="addBookmark">
	<div class="top">
		<p>New bookmark</p>
		<i class="las la-times" on:click|stopPropagation={hideAddBookmarkComponent} role="none" />
	</div>

	<form class="bookmark-input" on:submit|preventDefault={handleAddBookmarkFormSubmit} id="urlForm">
		<input
			id="myURL"
			name="myURL"
			type="text"
			placeholder="https://example.com"
			pattern="https://.*"
			autocomplete="off"
			spellcheck="false"
			bind:value={bookmark}
			on:focus|stopPropagation={showDefaultURLs}
			on:blur|stopPropagation={hideDefaultURLs}
			on:paste|stopPropagation={hideDefaultURLs}
			on:input|stopPropagation={hideDefaultURLs}
		/>

		<div id="defaultURLs">
			<div class="option">
				<span class="url">https://developer.mozilla.org/en-US/</span>
				<span class="label">MDN Web Docs</span>
			</div>
			<div class="option">
				<span class="url">https://stackoverflow.com/</span>
				<span class="label">Stack Overflow</span>
			</div>
		</div>

		<div class="buttons">
			<button type="submit">
				<span>Add bookmark</span>
			</button>
			<button class="cancel" on:click|stopPropagation={hideAddBookmarkComponent}>
				<span>Cancel</span>
			</button>
		</div>
	</form>
	<!-- <form class="tags-input" on:submit|preventDefault={handleTagFormSubmit}>
			<div class="selectedTags">
				{#if selectedTags.length > 0}
					{#each selectedTags as { name, id, user_id, added, updated, deleted }}
						<div class="tag" data-name={name} data-id={id} data-userid={user_id}>
							<i class="las la-hashtag" />
							<span>{name}</span>
							<i
								class="las la-times"
								on:click|stopPropagation={removeTagFromSelectedTags}
								role="none"
							/>
						</div>
					{/each}
				{/if}
			</div>
			<input
				type="text"
				name="tags"
				id="tags"
				placeholder="attach at least one tag"
				autocomplete="off"
				spellcheck="false"
				bind:value={tagName}
				on:input|stopPropagation={getUserMatchingTags}
				on:mousedown={getAllUserTagsIfInputEmpty}
				on:blur={handleTagsInputBlurEvent}
				on:keydown|stopPropagation={handleKeyDownInTagInput}
			/>
			<div class="matchingTags">
				{#if matchingTags.length > 0}
					{#each matchingTags as { name, id, user_id, added, deleted, updated }}
						<div
							class="tag"
							data-id={id}
							data-name={name}
							data-userid={user_id}
							on:mousedown={handleClickOnMatchingTag}
							role="none"
						>
							<i class="las la-hashtag" />
							<span>{name}</span>
						</div>
					{/each}
				{/if}
			</div>
		</form> -->
</div>

<style lang="scss">
	.container {
		z-index: 2;
		position: fixed;
		top: 0;
		right: 0;
		width: 40rem;
		height: 100vh;
		overflow-y: auto;
		background-color: rgb(245, 245, 245);
		display: flex;
		flex-direction: column;
		gap: 1em;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		transform: translateX(200%);
		transition: all ease 300ms;

		@media only screen and (max-width: 425px) {
			width: 99vw;
		}

		.top {
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 1em;

			p {
				font-size: 1.5rem;
				font-family: 'Arial CE', sans-serif;
			}

			i {
				font-size: 2rem;
				cursor: pointer;
			}
		}

		form {
			border-radius: 0.3rem;
			position: relative;
			background-color: inherit;
			display: flex;
			flex-direction: column;
			gap: 2em;
			padding: 1em;

			input[type='text'] {
				width: 100%;
				padding: 0.7em;
				font-family: 'Arial CE', sans-serif;
				font-size: 1.3rem;
				border: 0.1rem solid rgb(0, 0, 0, 0.1);
				border-radius: inherit;
				border-radius: 0.3rem;
				outline: none;
			}

			#defaultURLs {
				box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
				border: 0.1rem solid rgb(0, 0, 0, 0.1);
				min-height: max-content;
				max-height: 70vh;
				overflow-y: auto;
				position: absolute;
				top: 101%;
				right: 0;
				left: 0;
				width: inherit;
				border-radius: 0.3rem;
				background-color: rgb(255, 255, 255);
				display: flex;
				flex-direction: column;
				z-index: 3;
				display: none;

				.option {
					width: 100%;
					border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
					padding: 0.5em;
					min-height: 3.5rem;
					display: flex;
					justify-content: center;
					flex-direction: column;
					gap: 0.5em;
					cursor: pointer;
					background-color: inherit;

					span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
					}

					span.label {
						color: rgb(0, 0, 0, 0.6);
					}

					&:hover {
						background-color: rgb(255, 234, 221);
					}
				}
			}
		}

		.buttons {
			width: 100%;
			min-height: max-content;
			display: flex;
			align-items: center;
			justify-content: flex-end;
			gap: 1em;

			button {
				min-width: max-content;
				padding: 0.7em;
				cursor: pointer;
				display: flex;
				align-items: center;
				justify-content: center;
				border: none;
				border-radius: 0.3rem;
				background-color: rgb(238, 238, 238);

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
				}
			}

			button.cancel {
				display: none;
			}

			button[type='submit'] {
				background-color: #025464;
				width: 100%;

				span {
					color: rgb(255, 255, 255);
				}
			}
		}
	}
</style>
