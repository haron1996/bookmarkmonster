<script lang="ts">
	import { matchedTagsFromDB, apiHost, session, error, selectedTags, tags } from '../stores/stores';
	import type { Tag } from '../types/tag';
	import { closeRenameTagModal } from '../utils/closeRenameTagModal';

	async function fetchUserMatchingTags() {
		if ($selectedTags[0].name === '') {
			matchedTagsFromDB.set([]);
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/${$selectedTags[0].name}`, {
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
			if ($selectedTags[0].name === '') {
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
			// tagName.set('');
			$selectedTags[0].name = '';
			matchedTagsFromDB.set([]);
			return;
		}
	}

	async function handleFornSubmit() {
		if ($selectedTags[0].name === '') {
			error.set('tag name cannot be empty');
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/renameTag`, {
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
			body: JSON.stringify({ tag: $selectedTags[0] })
		});

		const result = await response.json();

		const t: Tag = result[0];

		const index: number = $tags.findIndex((value) => {
			return value.id === t.id;
		});

		$tags.splice(index, 1, t);

		tags.set($tags);

		closeRenameTagModal();
	}

	$: $selectedTags, $selectedTags.length >= 1 ? fetchUserMatchingTags() : () => {};
</script>

<div class="wrapper" id="renameTag" role="none" on:click={closeRenameTagModal}>
	<form
		role="none"
		on:click|stopPropagation={() => {}}
		on:submit|preventDefault|stopPropagation={handleFornSubmit}
	>
		<div class="heading">
			<div class="left">
				<p>Rename Tag</p>
				<div class="tagNameToUpdate">
					<i class="las la-hashtag" />
					<span>copywriting</span>
				</div>
			</div>
			<i class="las la-times" role="none" on:click|stopPropagation={closeRenameTagModal} />
		</div>
		<div class="inputContainer">
			<div class="input">
				<div class="icon">
					<i class="las la-hashtag" />
				</div>
				{#if $selectedTags.length >= 1}
					<input
						type="text"
						name="tag"
						id="renameTagInput"
						placeholder="new tag name"
						autocomplete="off"
						spellcheck="false"
						bind:value={$selectedTags[0].name}
						on:input={fetchUserMatchingTags}
						on:keydown={handleInputKeyDown}
					/>
				{/if}
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
		<button type="submit" on:click|preventDefault|stopPropagation={handleFornSubmit}>
			<span>Rename tag</span>
		</button>
	</form>
</div>

<style lang="scss">
	.wrapper {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 1;
		background-color: rgb(0, 0, 0, 0);
		display: none;
		transition: background-color 0.5s ease;

		form {
			display: flex;
			flex-direction: column;
			gap: 3em;
			border-radius: 0.3rem;
			box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em, rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em,
				rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
			background-color: rgba(255, 255, 255);
			width: 40rem;
			min-height: max-content;
			padding: 1em;

			@media only screen and (max-width: 425px) {
				max-width: 98%;
			}

			.heading {
				display: flex;
				align-content: center;
				justify-content: space-between;

				.left {
					display: flex;
					gap: 1em;
					align-items: center;

					p {
						font-size: 1.3rem;
						font-weight: 600;
						font-family: 'Arial CE', sans-serif;
					}

					.tagNameToUpdate {
						display: flex;
						align-items: center;
						gap: 0.5em;
						background-color: rgb(238, 238, 238);
						padding: 0.3em 0.5em;
						border-radius: 0.3rem;

						i {
							background-color: #78c1f3;
							color: white;
							font-size: 1.3rem;
						}

						span {
							font-family: 'Arial CE', sans-serif;
							font-size: 1.3rem;
						}
					}
				}

				i {
					font-size: 1.5rem;
					cursor: pointer;
				}
			}

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

			button {
				width: 100%;
				padding: 0.7em;
				display: flex;
				align-items: center;
				justify-content: center;
				border: none;
				background-color: #025464;
				cursor: pointer;

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					color: white;
				}
			}
		}
	}
</style>
