<script lang="ts">
	import {
		apiHost,
		childrenOfDestinationFolder,
		destinationFolder,
		folders,
		session,
		showCreateFolderFromMoveFoldersPopup
	} from '../stores/stores';
	import type { Folder } from '../types/folder';
	import { sortFoldersByName } from '../utils/sortFolders';

	let collectionName: string = '';
	let collectionDescription: string = '';
	let parentFolderID: string = '';

	async function createCollection() {
		if (collectionName === '') {
			alert('collection name required');
			return;
		}
		$destinationFolder.folder_id
			? (parentFolderID = $destinationFolder.folder_id)
			: (parentFolderID = '');

		const response = await fetch(`${$apiHost}/authenticated/collections/createCollection`, {
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
				name: collectionName,
				parent_folder_id: parentFolderID,
				description: collectionDescription
			})
		});

		const result = await response.json();

		if (result.message) {
			alert(result.message);
			return;
		}

		const folder: Folder = result[0];

		childrenOfDestinationFolder.update((values) => [folder, ...values]);

		sortFoldersByName($childrenOfDestinationFolder);

		showCreateFolderFromMoveFoldersPopup.set(false);

		collectionName = '';

		collectionDescription = '';

		destinationFolder.set(folder);
	}
</script>

{#if $showCreateFolderFromMoveFoldersPopup}
	<div
		class="container"
		on:click={() => {
			showCreateFolderFromMoveFoldersPopup.set(false);
		}}
		role="none"
	>
		<form role="none" on:click|stopPropagation={() => {}}>
			<div class="top">
				<p>Create collection</p>
			</div>
			<div class="inputs">
				<div class="name">
					<label for="name">Collection name</label>
					<input
						type="text"
						name="name"
						id="name"
						autocomplete="off"
						autocorrect="false"
						placeholder="Collection name"
						bind:value={collectionName}
					/>
				</div>
				<div class="description">
					<label for="description">Collection description</label>
					<textarea
						name="description"
						id="description"
						autocomplete="off"
						autocorrect="false"
						placeholder="Collection description"
						bind:value={collectionDescription}
					/>
				</div>
			</div>
			<button on:click|preventDefault={createCollection}>Create collection</button>
		</form>
	</div>
{/if}

<style lang="scss">
	.container {
		position: fixed;
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
		background-color: rgb(0, 0, 0, 0.4);
		z-index: 11;
		display: flex;
		align-items: center;
		justify-content: center;

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

			.top {
				display: flex;
				align-items: center;

				p {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					text-transform: uppercase;
				}
			}

			.inputs {
				display: flex;
				flex-direction: column;
				gap: 1em;

				div {
					display: flex;
					flex-direction: column;
					gap: 0.5em;

					label {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
					}

					input[type='text'] {
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
				}
			}

			button {
				padding: 1em 0.7em;
				border: none;
				cursor: pointer;
				background-color: #040d12;
				border-radius: 0.3rem;
				color: white;
				box-shadow: rgba(0, 0, 0, 0.4) 0px 2px 4px, rgba(0, 0, 0, 0.3) 0px 7px 13px -3px,
					rgba(0, 0, 0, 0.2) 0px -3px 0px inset;
				display: flex;
				align-items: center;
				justify-content: center;
				gap: 0.7em;
				font-family: 'Segoe UI', sans-serif;
				font-weight: 600;
				font-size: 1.2rem;
			}

			@media only screen and (width <= 600px) {
				width: 95%;
			}
		}
	}
</style>
