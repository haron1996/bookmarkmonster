<script lang="ts">
	import { page } from '$app/stores';
	import { apiHost, bookmarks, session } from '../stores/stores';
	import { closeCreateBookmark } from '../utils/closeCreateBookmark';
	import { sortBookmarksByTitle } from '../utils/sortBookmarks';

	let bookmark: string = '';
	let notes: string = '';
	let savingBookmark: boolean = false;
	let folderID: string = '';

	function getFolderID() {
		const id: string | null = $page.url.searchParams.get('id');
		if (id) {
			if (id === 'root') {
				folderID = '';
			} else {
				folderID = id;
			}
		}
	}

	async function createBookmark() {
		savingBookmark = true;
		getFolderID();
		const response = await fetch(`${$apiHost}/authenticated/bookmarks/createBookmark`, {
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
			body: JSON.stringify({ bookmark: bookmark, notes: notes, folder_id: folderID })
		});

		const result = await response.json();

		bookmarks.update((values) => [result[0], ...values]);

		savingBookmark = false;

		bookmark = '';

		notes = '';

		// sortBookmarksByTitle($bookmarks);

		closeCreateBookmark();
	}
</script>

<div class="createBookmark" id="createBookmark" on:click={closeCreateBookmark} role="none">
	<form
		on:click|stopPropagation={() => {}}
		role="none"
		class="animate__animated animate__backInDown"
	>
		<p>Create bookmark</p>
		<input
			type="url"
			name="url"
			id="url"
			placeholder="https://example.com"
			pattern="https://.*"
			autocomplete="off"
			bind:value={bookmark}
		/>
		<textarea autocomplete="off" placeholder="Enter bookmark notes" bind:value={notes} />
		<button type="submit" on:click|preventDefault={createBookmark}>
			{#if savingBookmark}
				<div class="loader" />
				<span>processing...</span>
			{:else}
				<span>Save and close</span>
			{/if}
		</button>
	</form>
</div>

<style lang="scss">
	.createBookmark {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 5;
		background-color: rgb(0, 0, 0, 0.4);
		transition: all ease 300ms;
		display: none;

		form {
			width: 60rem;
			min-height: max-content;
			background-color: white;
			display: flex;
			flex-direction: column;
			padding: 1.5em;
			border-radius: 0.5rem;
			gap: 3em;
			box-shadow: rgba(0, 0, 0, 0.15) 0px 5px 15px 0px;

			p {
				font-family: 'Segoe UI', sans-serif;
				font-size: 1.2rem;
				font-weight: 600;
				text-transform: uppercase;
			}

			input {
				padding: 0.7em;
				border-radius: 0.3rem;
				border: 0.2rem solid rgb(232, 240, 254);
				outline: none;
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
				background-color: rgb(232, 240, 254);
			}

			textarea {
				width: 100%;
				height: 20rem;
				background-color: rgb(232, 240, 254);
				resize: none;
				border-radius: 0.3rem;
				padding: 0.7em;
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
				outline: none;
				border: 0.2rem solid rgb(232, 240, 254);
			}

			button {
				padding: 1em 0.7em;
				border: none;
				cursor: pointer;
				background-color: #0079ff;
				border-radius: 0.3rem;
				color: white;
				box-shadow: rgba(0, 0, 0, 0.4) 0px 2px 4px, rgba(0, 0, 0, 0.3) 0px 7px 13px -3px,
					rgba(0, 0, 0, 0.2) 0px -3px 0px inset;
				display: flex;
				align-items: center;
				justify-content: center;
				gap: 0.7em;

				.loader {
					border: 2px solid #f3f3f3; /* Light grey */
					border-top: 2px solid #3498db; /* Blue */
					border-radius: 50%;
					width: 2rem;
					height: 2rem;
					animation: spin 0.5s linear infinite;

					@keyframes spin {
						0% {
							transform: rotate(0deg);
						}
						100% {
							transform: rotate(360deg);
						}
					}
				}

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					font-weight: 600;
				}
			}
		}
	}
</style>
