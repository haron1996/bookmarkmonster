<script lang="ts">
	import { creatingFolder, newFolder } from '../stores/stores';
	import { closeCreateFolder } from '../utils/closeCreateFolder';
	import { createFolder } from '../utils/createFolder';
</script>

<div class="createFolder" id="createFolder" on:click={closeCreateFolder} role="none">
	<form
		on:click|stopPropagation={() => {}}
		role="none"
		class="animate__animated animate__backInDown"
	>
		<p>create collection</p>
		<input
			type="text"
			name="name"
			id="name"
			autocomplete="off"
			autocorrect="false"
			placeholder="Collection name"
			bind:value={$newFolder.folder_name}
		/>
		<textarea
			name="description"
			id="description"
			placeholder="Collection description"
			bind:value={$newFolder.folder_description}
		/>
		<button on:click|preventDefault={createFolder}>
			{#if $creatingFolder}
				<div class="loader" />
				<span>processing...</span>
			{:else}
				<span>Create and close</span>
			{/if}
		</button>
	</form>
</div>

<style lang="scss">
	.createFolder {
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
				font-size: 1.2rem;
				font-weight: 600;
				text-transform: uppercase;
				font-family: 'Segoe UI', sans-serif;
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
</style>
