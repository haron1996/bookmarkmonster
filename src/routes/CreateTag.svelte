<script lang="ts">
	import type { Tag } from '../types/tag';
	import { apiHost, session, tags } from '../stores/stores';
	import { hideOverlay } from '../utils/hideOverlay';
	import { showTagCreatedAlert } from '../utils/showTagCreatedAlert';
	import { hideTagCreatedAlert } from '../utils/hideTagCreatedAlert';
	import { showDuplicateTagAlert } from '../utils/showDuplicateTagAlert';
	import { hideDuplicateTagAlert } from '../utils/hideDuplicateTagAlert';
	import { page } from '$app/stores';

	let tag: Tag = {};

	let tagName: string = '';

	const showSuggestedTags = () => {
		if (tagName != '') return;

		const suggestedTags = document.getElementById('suggestedTags') as HTMLDivElement | null;

		if (suggestedTags === null) return;

		suggestedTags.style.display = 'none';
	};

	const hideSuggestedTags = () => {
		const suggestedTags = document.getElementById('suggestedTags') as HTMLDivElement | null;

		if (suggestedTags === null) return;

		suggestedTags.style.display = 'none';
	};

	const createTag = async () => {
		if (tagName === '') {
			console.log('tag name required');
			return;
		}

		if (
			$tags
				.map((t) => {
					return t.name;
				})
				.includes(tagName)
		) {
			console.log('tag already exists');
			tagName = '';

			hideOverlay();

			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/create-tag`, {
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
			body: JSON.stringify({ name: tagName })
		});

		if (response.status === 200) {
			const result = await response.json();

			if (result.message) {
				console.log(result.message);
				return;
			}

			if (result.message === 'duplicate tag') {
				showDuplicateTagAlert();

				setTimeout(() => {
					hideDuplicateTagAlert();
				}, 3000);

				return;
			}

			tag = result[0];

			tags.update((tags) => [tag, ...tags]);
		} else {
			console.log(response.status, response.statusText);
		}

		tagName = '';

		hideOverlay();
	};
</script>

<div class="container" id="createTag">
	<div class="top">
		<h6>Create Tag</h6>
		<i class="las la-times" on:click|stopPropagation={hideOverlay} role="none" />
	</div>
	<div class="input-and-submit-button">
		<div class="input">
			<input
				type="text"
				name="tag"
				id="tag"
				placeholder="AI"
				autocomplete="off"
				spellcheck="false"
				bind:value={tagName}
				on:focus|stopPropagation={showSuggestedTags}
				on:blur|stopPropagation={hideSuggestedTags}
				on:paste|stopPropagation={hideSuggestedTags}
				on:input|stopPropagation={hideSuggestedTags}
			/>
			<div class="suggestedTags" id="suggestedTags">
				<div class="suggested-tag">
					<i class="las la-hashtag" />
					<span>Education</span>
				</div>
				<div class="suggested-tag">
					<i class="las la-hashtag" />
					<span>Politics</span>
				</div>
				<div class="suggested-tag">
					<i class="las la-hashtag" />
					<span>Sports</span>
				</div>
				<div class="suggested-tag">
					<i class="las la-hashtag" />
					<span>Entertainment</span>
				</div>
			</div>
		</div>
		<div class="buttons">
			<button class="cancel" on:click|stopPropagation={hideOverlay}>
				<span>Cancel</span>
			</button>
			<button type="submit" on:click|stopPropagation={createTag}>
				<span>Create</span>
			</button>
		</div>
	</div>
</div>

<style lang="scss">
	.container {
		z-index: 2;
		position: fixed;
		top: 1%;
		left: 0.5%;
		width: 27.5rem;
		height: 98vh;
		max-height: 98vh;
		background-color: rgb(245, 245, 245);
		display: flex;
		flex-direction: column;
		transform: translateX(0);
		border-radius: 0.5rem;
		border: 0.1rem solid rgb(0, 0, 0, 0.1);
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		gap: 1em;
		transform: translateX(-200%);
		transition: all ease 300ms;
		overflow-y: auto;

		.top {
			display: flex;
			align-items: center;
			justify-content: space-between;
			min-height: 7vh;
			padding: 1em;

			h6 {
				font-size: 1.5rem;
				font-family: 'Arial CE', sans-serif;
			}

			i {
				font-size: 2rem;
				cursor: pointer;
			}
		}

		.input-and-submit-button {
			min-height: max-content;
			display: flex;
			flex-direction: column;
			gap: 2em;
			padding: 1em;

			.input {
				width: 100%;
				min-height: 3.5rem;
				border-radius: 0.3rem;
				position: relative;

				input[type='text'] {
					width: 100%;
					min-height: 3.5rem;
					padding: 0.5em;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					border-radius: inherit;

					&:placeholder-shown {
						font-family: 'Arial CE', sans-serif;
					}

					&::placeholder {
						font-family: 'Arial CE', sans-serif;
					}
				}

				.suggestedTags {
					position: absolute;
					top: 101%;
					left: 0;
					right: 0;
					width: inherit;
					min-height: max-content;
					max-height: 70vh;
					overflow-y: auto;
					background-color: rgb(255, 255, 255);
					box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
					border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
					display: flex;
					flex-direction: column;
					border-radius: 0.3rem;
					display: none;

					.suggested-tag {
						width: 100%;
						padding: 0.5em;
						min-height: 3.5rem;
						display: flex;
						align-items: center;
						gap: 0.5em;
						cursor: pointer;
						background-color: inherit;
						border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
						color: rgb(0, 0, 0, 0.5);

						i {
							font-size: 1.8rem;
						}

						span {
							font-size: 1.2rem;
							font-family: 'Arial CE', sans-serif;
						}

						&:hover {
							background-color: rgb(255, 234, 221);
						}
					}
				}
			}

			.buttons {
				display: flex;
				align-items: center;
				justify-content: space-between;
				gap: 0.5em;

				button {
					border: none;
					min-width: 7rem;
					min-height: 3.5rem;
					padding: 0.5em;
					cursor: pointer;
					border-radius: 0.3rem;

					span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
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
