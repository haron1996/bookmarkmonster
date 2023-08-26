<script lang="ts">
	import { apiHost, bookmarks, selectedBookmarks, session } from '../stores/stores';
	import type { Bookmark } from '../types/bookmark';
	import { closeTrashBookmarkConfirmation } from '../utils/closeTrashBookmarkConfirmation';

	async function handleDeleteButtonClick() {
		const response = await fetch(`${$apiHost}/authenticated/bookmarks/trashBookmarks`, {
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
			body: JSON.stringify({ bookmarks: $selectedBookmarks })
		});

		const result = await response.json();

		const trashedBookmarks: Bookmark[] = result[0];

		trashedBookmarks.forEach((b) => {
			$bookmarks = $bookmarks.filter((value) => {
				return value.id !== b.id;
			});
		});

		selectedBookmarks.set([]);

		closeTrashBookmarkConfirmation();

		const bookmarkNodes = document.querySelectorAll(
			'.bookmark'
		) as NodeListOf<HTMLDivElement> | null;

		if (bookmarkNodes === null) return;

		bookmarkNodes.forEach((bn) => {
			bn.classList.remove('bookmarkSelected');
		});
	}
</script>

<div
	class="wrapper"
	id="trashBookmarkConfirmation"
	role="none"
	on:click={closeTrashBookmarkConfirmation}
>
	<form on:click|stopPropagation={() => {}} role="none">
		<p>
			Are you sure you want to delete {$selectedBookmarks.length > 1
				? 'these bookmarks'
				: 'this bookmark'} 🤔
		</p>
		<div class="buttons">
			<button type="submit" on:click|preventDefault|stopPropagation={handleDeleteButtonClick}>
				<span>Delete</span>
			</button>
			<button on:click|preventDefault|stopPropagation={closeTrashBookmarkConfirmation}>
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

			p {
				font-size: 1.5rem;
				font-family: 'Arial CE', sans-serif;
				line-height: 1.6;
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
		}
	}
</style>
