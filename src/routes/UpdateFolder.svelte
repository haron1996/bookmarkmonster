<script lang="ts">
	import { updatingFolder, selectedFolders, showUpdateFolder } from '../stores/stores';
	import { updateFolder } from '../utils/updateFolder';
</script>

<div class="updateFolder" class:showUpdateFolder={$showUpdateFolder}>
	{#if $selectedFolders[0]}
		<form class="animate__animated animate__backInDown">
			<div class="top">
				<p>update <span>{$selectedFolders[0].folder_name}</span></p>
				<span
					class="close"
					role="none"
					on:click={() => {
						showUpdateFolder.set(false);
					}}>cancel</span
				>
			</div>
			<input
				type="text"
				name="folderName"
				id="folderName"
				autocomplete="off"
				spellcheck="false"
				bind:value={$selectedFolders[0].folder_name}
			/>
			<textarea
				name="description"
				id="description"
				autocomplete="off"
				spellcheck="false"
				bind:value={$selectedFolders[0].folder_description}
			/>
			<button type="submit" on:click|preventDefault={updateFolder}>
				{#if $updatingFolder}
					<div class="loader" />
					<span>processing...</span>
				{:else}
					<span>Save and close</span>
				{/if}
			</button>
		</form>
	{/if}
</div>

<style lang="scss">
	.updateFolder {
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
				justify-content: space-between;
				align-items: center;
				min-height: 5vh;

				p {
					font-size: 1.2rem;
					font-weight: 600;
					text-transform: uppercase;
					font-family: 'Segoe UI', sans-serif;
					white-space: nowrap;
					overflow: hidden;
					text-overflow: ellipsis;
					max-width: 75%;

					span {
						color: #0079ff;
					}
				}

				span.close {
					background-color: rgb(232, 240, 254);
					font-size: 1.2rem;
					font-family: 'Segoe UI', sans-serif;
					text-transform: uppercase;
					padding: 0.5em 1em;
					border-radius: 0.3rem;
					cursor: pointer;
					box-shadow: rgba(255, 255, 255, 0.2) 0px 0px 0px 1px inset,
						rgba(0, 0, 0, 0.9) 0px 0px 0px 1px;
				}
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
				text-transform: uppercase;

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

	.showUpdateFolder {
		min-height: 100%;
		max-height: 100%;

		form {
			display: flex;
		}
	}
</style>
