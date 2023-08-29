<script lang="ts">
	import type { Tag } from '../types/tag';
	import { apiHost, error, matchedTagsFromDB, session, tags, tagName } from '../stores/stores';
	import { hideOverlay } from '../utils/hideOverlay';
	import { showTagCreatedAlert } from '../utils/showTagCreatedAlert';
	import { hideTagCreatedAlert } from '../utils/hideTagCreatedAlert';
	import { showDuplicateTagAlert } from '../utils/showDuplicateTagAlert';
	import { hideDuplicateTagAlert } from '../utils/hideDuplicateTagAlert';
	import { page } from '$app/stores';

	let tag: Tag = {};

	const showSuggestedTags = () => {
		if ($tagName != '') return;

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
		if ($tagName === '') {
			error.set('tag name required');
			return;
		}

		const t = $matchedTagsFromDB.filter((value) => {
			return value.name === $tagName;
		});

		if (t[0]) {
			error.set('tag name already exists');
			tagName.set('');
			matchedTagsFromDB.set([]);
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
			body: JSON.stringify({ name: $tagName })
		});

		if (response.ok) {
			const result = await response.json();

			if (result.message) {
				console.log(result.message);
				return;
			}

			if (result.message === 'duplicate tag') {
				error.set('tag already exists');

				return;
			}

			tag = result[0];

			tags.update((tags) => [tag, ...tags]);
		} else {
			console.log(response.status, response.statusText);
		}

		tagName.set('');

		hideOverlay();

		matchedTagsFromDB.set([]);
	};

	async function fetchUserMatchingTags() {
		if ($tagName === '') {
			matchedTagsFromDB.set([]);
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/${$tagName}`, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		const tgs: Tag[] = result[0];

		matchedTagsFromDB.set(tgs);
	}

	function handleInputKeyDown(e: KeyboardEvent) {
		if (e.code === 'Space') {
			if ($tagName === '') {
				e.preventDefault();
				return;
			}
		}
	}

	function handleClickOnMatchedTagFromDB(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('.tag') as HTMLDivElement | null;

		if (tag === null) return;

		const t = $matchedTagsFromDB.filter((value) => {
			return value.name === tag.innerText;
		});

		if (t[0]) {
			error.set('tag name already exists');
			tagName.set('');
			matchedTagsFromDB.set([]);
			return;
		}
	}
</script>

<div class="container" id="createTag">
	<div class="top">
		<p>New tag</p>
		<i class="las la-times" on:click|stopPropagation={hideOverlay} role="none" />
	</div>
	<div class="input-and-submit-button">
		<form>
			<div class="inputContainer">
				<div class="input">
					<div class="icon">
						<i class="las la-hashtag" />
					</div>
					<input
						type="text"
						name="tag"
						id="tag"
						placeholder="eg AI"
						autocomplete="off"
						spellcheck="false"
						bind:value={$tagName}
						on:input={fetchUserMatchingTags}
						on:focus|stopPropagation={showSuggestedTags}
						on:blur|stopPropagation={hideSuggestedTags}
						on:paste|stopPropagation={hideSuggestedTags}
						on:keydown={handleInputKeyDown}
					/>
				</div>
				<div class="matchingTags" id="matchingTags">
					{#each $matchedTagsFromDB as { added, deleted, id, name, updated, user_id }}
						<div class="tag" role="none" on:click={handleClickOnMatchedTagFromDB}>
							<i class="las la-hashtag" />
							<span>{name}</span>
						</div>
					{/each}
				</div>
			</div>
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
			<div class="buttons">
				<button type="submit" on:click|stopPropagation|preventDefault={createTag}>
					<span>Create tag</span>
				</button>
				<button class="cancel" on:click|stopPropagation|preventDefault={hideOverlay}>
					<span>Cancel</span>
				</button>
			</div>
		</form>
	</div>
</div>

<style lang="scss">
	.container {
		z-index: 2;
		position: fixed;
		top: 0;
		left: 0;
		width: 27.5rem;
		height: 100vh;
		background-color: rgb(245, 245, 245);
		display: flex;
		flex-direction: column;
		transform: translateX(0);
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
			padding: 0.5em;

			p {
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
			padding: 0.1em;

			form {
				width: 100%;
				min-height: 3.5rem;
				border-radius: 0.3rem;
				position: relative;
				display: flex;
				flex-direction: column;
				gap: 2em;

				.inputContainer {
					display: flex;
					flex-direction: column;
					width: 100%;
					position: relative;

					.input {
						display: flex;
						align-items: center;
						border: 0.1rem solid rgb(2, 84, 100, 0.1);
						background-color: rgb(255, 255, 255);
						height: 3rem;
						border-radius: 0.2rem;

						.icon {
							min-width: 10%;
							height: 100%;
							display: flex;
							align-items: center;
							justify-content: center;

							i {
								background-color: rgb(120, 193, 243);
								font-size: 1.5rem;
								color: rgb(255, 255, 255);
							}
						}

						input[type='text'] {
							border: none;
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							outline: none;
							width: 90%;
							height: 100%;
							padding: 0 0.2em;
						}
					}

					.matchingTags {
						position: absolute;
						top: 100%;
						left: 0;
						right: 0;
						width: inherit;
						min-height: max-content;
						max-height: 20rem;
						overflow-y: auto;
						background-color: rgb(255, 255, 255);
						box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em,
							rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em, rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;

						.tag {
							min-height: 4rem;
							display: flex;
							align-items: center;
							gap: 1em;
							padding: 0.3em;
							border-bottom: 0.1rem solid rgb(2, 84, 100, 0.1);
							cursor: pointer;

							i {
								background-color: rgb(120, 193, 243);
								font-size: 1.5rem;
								color: rgb(255, 255, 255);
							}

							span {
								font-size: 1.3rem;
								font-family: 'Arial CE', sans-serif;
							}
						}
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

				.buttons {
					display: flex;
					align-items: center;
					justify-content: space-between;
					gap: 0.5em;

					button {
						border: none;
						min-width: max-content;
						padding: 0.7em;
						cursor: pointer;
						border-radius: 0.2rem;

						span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
						}
					}

					button.cancel {
						display: none;
					}

					button[type='submit'] {
						width: 100%;
						background-color: #025464;

						span {
							color: rgb(255, 255, 255);
						}
					}
				}
			}
		}
	}
</style>
