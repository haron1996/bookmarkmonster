<script lang="ts">
	import { hideAddBookmarkComponent } from '../utils/hideAddBookmarkComponent';
	import { hideOverlay } from '../utils/hideOverlay';

	let bookmark: string = '';

	const showDefaultURLs = () => {
		const defaultURLS = document.getElementById('defaultURLs') as HTMLDivElement | null;

		if (defaultURLS === null) return;

		defaultURLS.style.display = 'flex';
	};

	const hideDefaultURLs = () => {
		const defaultURLS = document.getElementById('defaultURLs') as HTMLDivElement | null;

		if (defaultURLS === null) return;

		defaultURLS.style.display = 'none';
	};
</script>

<div class="container" id="addBookmark">
	<div class="top">
		<h6>Add Bookmark</h6>
		<i class="las la-times" on:click|stopPropagation={hideAddBookmarkComponent} role="none" />
	</div>
	<div class="input-and-buttons">
		<div class="input">
			<input
				id="myURL"
				name="myURL"
				type="url"
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
		</div>
		<div class="buttons">
			<button class="cancel" on:click|stopPropagation={hideAddBookmarkComponent}>
				<span>Cancel</span>
			</button>
			<button type="submit">
				<span>Save</span>
			</button>
		</div>
	</div>
</div>

<style lang="scss">
	.container {
		z-index: 2;
		position: fixed;
		top: 1%;
		right: 0.5%;
		min-width: 40rem;
		width: auto;
		height: 98vh;
		background-color: rgb(245, 245, 245);
		display: flex;
		flex-direction: column;
		padding: 1em;
		gap: 1em;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		border-radius: 0.5rem;
		transform: translateX(200%);
		transition: all ease 300ms;

		.top {
			display: flex;
			align-items: center;
			justify-content: space-between;
			min-height: 7vh;

			h6 {
				font-size: 1.5rem;
				font-family: 'Arial CE', sans-serif;
			}

			i {
				font-size: 2rem;
				cursor: pointer;
			}
		}

		.input-and-buttons {
			display: flex;
			flex-direction: column;
			gap: 2em;

			.input {
				border-radius: 0.3rem;
				position: relative;

				input[type='url'] {
					width: 100%;
					min-height: 3.5rem;
					padding: 0.5em;
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					border-radius: inherit;
					border-radius: 0.3rem;

					&::placeholder {
						font-family: 'Arial CE', sans-serif;
					}

					&:placeholder-shown {
						background-color: rgb(255, 234, 221);
					}
				}

				#defaultURLs {
					box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					height: 40rem;
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
				padding: 0.5em;
				gap: 1em;

				button {
					width: 7rem;
					min-height: 3.5rem;
					min-width: max-content;
					padding: 0.5em;
					cursor: pointer;
					display: flex;
					align-items: center;
					justify-content: center;
					border: none;
					border-radius: 0.3rem;

					span {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
					}
				}

				button.cancel {
					width: 50%;
					&:hover {
						background-color: rgb(252, 174, 174);
					}
				}

				button[type='submit'] {
					width: 50%;
					background-color: rgb(0, 121, 255);

					span {
						color: rgb(255, 255, 255);
					}

					&:hover {
						background-color: rgb(78, 79, 235);
					}
				}
			}
		}
	}
</style>
