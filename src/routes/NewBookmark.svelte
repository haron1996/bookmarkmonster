<script lang="ts">
	import { apiHost, bookmarks, session, showAddNewBookmark } from '../stores/stores';
	import type { Bookmark } from '../types/bookmark';

	let savingBookmark: boolean = false;

	let newBookmark: Bookmark = {};

	async function createBookmark() {
		savingBookmark = true;

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
			body: JSON.stringify({
				bookmark: newBookmark.bookmark,
				notes: newBookmark.notes,
				folder_id: ''
			})
		});

		const result = await response.json();

		bookmarks.update((values) => [result[0], ...values]);

		savingBookmark = false;

		newBookmark = {};

		showAddNewBookmark.set(false);
	}
</script>

<div class="newBookmark" class:showAddNewBookmark={$showAddNewBookmark}>
	<form class="animate__animated animate__backInDown">
		<div class="top">
			<p>add bookmark</p>
			<i
				class="las la-times"
				role="none"
				on:click={() => {
					showAddNewBookmark.set(false);
				}}
			/>
		</div>
		<div class="URL">
			<label for="url">URL</label>
			<input
				type="url"
				name="url"
				id="url"
				placeholder="https://example.com"
				pattern="https://.*"
				autocomplete="off"
				spellcheck="false"
				bind:value={newBookmark.bookmark}
			/>
		</div>
		<div class="notes">
			<label for="notes">Notes</label>
			<textarea
				name="notes"
				id="notes"
				autocomplete="off"
				spellcheck="false"
				bind:value={newBookmark.notes}
			/>
		</div>
		<button
			type="submit"
			class:buttonDisabled={newBookmark.bookmark === '' ||
				newBookmark.bookmark === undefined ||
				savingBookmark}
			on:click|preventDefault={createBookmark}
		>
			{#if savingBookmark}
				<div class="loader" />
				<span>processing...</span>
			{:else}
				<span>Save and close</span>
			{/if}
			{#if newBookmark.bookmark === '' || newBookmark.bookmark === undefined || savingBookmark}
				<div class="blockButton" on:click|stopPropagation|preventDefault={() => {}} role="none" />
			{/if}
		</button>
	</form>
</div>

<style lang="scss">
	.newBookmark {
		width: 100%;
		min-height: 0%;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: all 300ms ease;
		background-color: white;
		flex-direction: column;

		form {
			width: 60rem;
			display: flex;
			flex-direction: column;
			padding: 1.5em;
			border-radius: 0.5rem;
			gap: 3em;
			transition: all 300ms ease;
			box-shadow: rgba(3, 102, 214, 0.3) 0px 0px 0px 3px;
			display: none;

			.top {
				display: flex;
				align-items: center;
				justify-content: space-between;

				p {
					font-size: 1.2rem;
					text-transform: uppercase;
					font-family: 'Segoe UI', sans-serif;
					font-weight: 600;
				}

				i {
					font-size: 1.8rem;
					cursor: pointer;
				}
			}

			.URL {
				display: flex;
				flex-direction: column;
				gap: 0.5em;

				label {
					font-family: 'Segoe UI', sans-serif;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
				}

				input {
					padding: 0.7em;
					border-radius: 0.3rem;
					border: 0.2rem solid rgb(232, 240, 254);
					outline: none;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					background-color: rgb(232, 240, 254);
					text-transform: lowercase;
				}
			}

			.notes {
				display: flex;
				flex-direction: column;
				gap: 0.5em;

				label {
					font-family: 'Segoe UI', sans-serif;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
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
			}

			button {
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
				text-transform: uppercase;
				position: relative;
				padding: 0.5em;
				width: 100%;
				height: 4rem;

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

				.blockButton {
					position: absolute;
					left: 0;
					top: 0;
					background-color: transparent;
					height: inherit;
					width: inherit;
					cursor: not-allowed;
				}
			}

			.buttonDisabled {
				opacity: 0.5;
			}
		}
	}

	.showAddNewBookmark {
		min-height: 100%;
		max-height: 100%;

		form {
			display: flex;
		}
	}
</style>
