<script lang="ts">
	import { updatingFolder, selectedFolders, showUpdateFolder } from '../stores/stores';
	import { updateFolder } from '../utils/updateFolder';
</script>

<div
	class="updateFolder"
	class:showUpdateFolder={$showUpdateFolder}
	role="none"
	on:click={() => {
		showUpdateFolder.set(false);
	}}
>
	{#if $selectedFolders[0]}
		<form
			class="animate__animated animate__backInDown"
			role="none"
			on:click|stopPropagation={() => {}}
		>
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
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: all 300ms ease;
		background-color: rgb(0, 0, 0, 0.7);
		flex-direction: column;
		z-index: 10;
		display: none;

		form {
			width: 60rem;
			display: flex;
			flex-direction: column;
			padding: 1.5em;
			border-radius: 0.5rem;
			gap: 3em;
			transition: all 300ms ease;
			display: none;
			background-color: white;
			box-shadow: rgba(0, 0, 0, 0.07) 0px 1px 2px, rgba(0, 0, 0, 0.07) 0px 2px 4px,
				rgba(0, 0, 0, 0.07) 0px 4px 8px, rgba(0, 0, 0, 0.07) 0px 8px 16px,
				rgba(0, 0, 0, 0.07) 0px 16px 32px, rgba(0, 0, 0, 0.07) 0px 32px 64px;

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
				background-color: rgb(4, 13, 18);
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

			@media only screen and (width <= 600px) {
				width: 95%;
			}
		}
	}

	.showUpdateFolder {
		display: flex;

		form {
			display: flex;
		}
	}
</style>
