<script lang="ts">
	import { apiHost, bookmarks, selectedBookmarks, session, error } from '../stores/stores';
	import type { Bookmark } from '../types/bookmark';
	import { hideRenameBookmark } from '../utils/hideRenameBookmark';

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

	$: $selectedBookmarks[0] && $selectedBookmarks[0].title === ''
		? error.set('bookmark title cannot be empty')
		: () => {};
</script>

<div class="wrapper" id="renameBookmarkComponent" role="none" on:click={hideRenameBookmark}>
	<form
		on:click|stopPropagation={() => {}}
		role="none"
		on:submit|preventDefault|stopPropagation={hideRenameBookmark}
	>
		<div class="top">
			<p>Rename bookmark</p>
			<i class="las la-times" role="none" on:click|stopPropagation={hideRenameBookmark} />
		</div>
		{#if $selectedBookmarks && $selectedBookmarks.length >= 1}
			<input
				type="text"
				name="title"
				id="renameBookmarkInput"
				autocomplete="off"
				spellcheck="false"
				bind:value={$selectedBookmarks[0].title}
				on:input={renameBookmark}
			/>
		{/if}
		<div class="buttons">
			<button type="submit" on:click|preventDefault|stopPropagation={hideRenameBookmark}>
				<span>Rename</span>
			</button>
			<button on:click|preventDefault|stopPropagation={hideRenameBookmark}>
				<span>Cancel</span>
			</button>
		</div>
	</form>
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

		form {
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

			input {
				padding: 0.5em;
				border-radius: inherit;
				border: 0.1rem solid rgb(0, 0, 0, 0.2);
				font-family: 'Arial CE', sans-serif;

				&:focus {
					outline-color: rgb(96, 1, 255);
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

					span {
						color: rgb(255, 255, 255);
					}
				}
			}

			@media only screen and (width <= 600px) {
				width: 95%;
			}
		}
	}
</style>
