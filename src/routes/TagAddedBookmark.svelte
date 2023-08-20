<script lang="ts">
	import type { Tag } from '../types/tag';
	import { apiHost, lastAddedBookmark, session, tags } from '../stores/stores';
	import GlobePNG from '$lib/images/globe.png';
	import { hideTagCreatedBookmarkForm } from '../utils/hideTagCreatedBookmarkForm';

	let tagName: string = '';
	let matchedTagsFromDB: Tag[] = [];
	let selectedTags: Tag[] = [];
	let taggingBookmark: boolean = false;

	const getMatchingTags = async () => {
		if (tagName === '') {
			matchedTagsFromDB = [];
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/${tagName}`, {
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

		matchedTagsFromDB = tgs;
	};

	function handleKeyDownOnTagsInput(e: KeyboardEvent) {
		// check if key is space bar
		if (e.code === 'Space' && tagName === '') {
			e.preventDefault();
			return;
		}

		// check if key is enter
		if ((e.code === 'Enter' || e.code === 'Comma') && tagName != '') {
			selectTag();
			return;
		}
	}

	function handleTagInputEmptied() {}

	async function selectTag() {
		const matchedTagsContainsTagName: boolean = matchedTagsFromDB
			.map((tag) => {
				return tag.name;
			})
			.includes(tagName);

		if (matchedTagsContainsTagName) {
			const matchingTag: Tag = matchedTagsFromDB.filter((tag) => tag.name === tagName)[0];

			const selectedTagsIncludeMatchingTag: boolean = selectedTags
				.map((st) => {
					return st.name;
				})
				.includes(matchingTag.name);

			if (selectedTagsIncludeMatchingTag) {
				tagName = '';
				matchedTagsFromDB = [];
			} else {
				selectedTags = [matchingTag, ...selectedTags];
			}
		} else if (!matchedTagsContainsTagName) {
			// create new tag
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

			if (response.ok) {
				const result = await response.json();

				tags.update((tags) => [result[0], ...tags]);

				selectedTags = [result[0], ...selectedTags];
			} else {
				console.log(response.status);
			}
		}

		tagName = '';

		matchedTagsFromDB = [];
	}

	function handleClickOnMatchedTag(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('div') as HTMLDivElement | null;

		if (tag === null) return;

		if (tag.dataset.name === undefined) return;

		tagName = tag.dataset.name;

		selectTag();
	}

	function removeTagFromSelectedTags(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('div') as HTMLDivElement | null;

		if (tag === null) return;

		if (tag.dataset.name === undefined) return;

		selectedTags = selectedTags.filter((st) => st.name != tag.dataset.name);
	}

	const tagBookmark = async () => {
		if (tagName === '' && selectedTags.length === 0) return;

		//selectTag();

		const response = await fetch(`${$apiHost}/authenticated/bookmarks/tagBookmark`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`,
				'Access-Control-Allow-Origin': '*'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ bookmark_id: $lastAddedBookmark.id, tags: selectedTags })
		});

		if (response.ok) {
			const result = await response.json();

			if (result.message === 'bookmark tagged successfully') {
				tagName = '';
				matchedTagsFromDB = [];
				selectedTags = [];
				hideTagCreatedBookmarkForm();
			}
		} else {
			console.log(response.status);
		}
	};
</script>

<div class="wrapper" id="tagCreatedBookmarkWrapper">
	<div class="card">
		<p>Tag your created bookmark:</p>
		{#if $lastAddedBookmark.id !== undefined}
			<div
				class="bookmark"
				data-id=""
				data-bookmark=""
				data-title=""
				data-thumbnail=""
				data-notes=""
				data-userid=""
				data-favicon=""
				data-updated=""
				data-added=""
				data-deleted=""
			>
				<div class="thumbnail">
					<img
						src={$lastAddedBookmark.thumbnail}
						alt="bookmark thumbnail"
						loading="lazy"
						draggable="false"
					/>
				</div>
				<div class="title-favicon-and-domain">
					<div class="title">
						<span>{$lastAddedBookmark.title}</span>
					</div>
					<div class="favicon-and-domain">
						{#if $lastAddedBookmark.favicon === ''}
							<img src={GlobePNG} alt="bookmark favicon" draggable="false" />
						{:else}
							<img src={$lastAddedBookmark.favicon} alt="bookmark favicon" draggable="false" />
						{/if}
						<span>{$lastAddedBookmark.host}</span>
					</div>
				</div>
			</div>
		{/if}
		<div class="tagsInputAndMatchedTags">
			<div class="tags">
				{#each selectedTags as { id, name, user_id, added, updated }}
					<div class="tag" data-id={id} data-name={name} data-userid={user_id}>
						<i class="las la-hashtag" />
						<span>{name}</span>
						<i class="las la-times" on:click={removeTagFromSelectedTags} role="none" />
					</div>
				{/each}
			</div>
			<div class="input">
				<input
					type="text"
					name="tagsInput"
					id="tagsInput"
					autocomplete="off"
					spellcheck="false"
					placeholder="enter tag name and press comma/enter"
					bind:value={tagName}
					on:input|stopPropagation={getMatchingTags}
					on:keydown|stopPropagation={handleKeyDownOnTagsInput}
					on:emptied={handleTagInputEmptied}
				/>
				<div class="matchedTags">
					{#each matchedTagsFromDB as { name, id, user_id, added, updated }}
						<div
							class="tag"
							data-name={name}
							data-id={id}
							data-userid={user_id}
							on:click={handleClickOnMatchedTag}
							role="none"
						>
							<i class="las la-hashtag" />
							<span>{name}</span>
						</div>
					{/each}
				</div>
			</div>
		</div>
		<button
			class:button-disabled={taggingBookmark || selectedTags.length === 0 || tagName != ''}
			on:click={tagBookmark}
		>
			<span>Tag bookmark</span>
		</button>
	</div>
</div>

<style lang="scss">
	.wrapper {
		position: fixed;
		top: 0;
		left: 0;
		height: 100vh;
		width: 100vw;
		background-color: rgb(0, 0, 0, 0);
		display: flex;
		align-items: center;
		justify-content: flex-end;
		z-index: 2;
		padding: 1em;
		overflow-y: auto;
		transform: translateX(200%);
		transition: transform ease 0.3s, background-color ease 6s;

		.card {
			width: 30rem;
			height: 100%;
			border-radius: 0.5rem;
			background-color: rgb(250, 250, 250);
			display: flex;
			flex-direction: column;
			padding: 1em;
			gap: 2em;
			box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;

			p {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.5rem;
			}

			.bookmark {
				width: 100%;
				border: 0.1rem solid rgb(0, 0, 0, 0.1);
				display: flex;
				flex-direction: column;
				border-radius: 0.6rem;
				gap: 1em;
				box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
				padding: 0.5em;

				.thumbnail {
					height: 70%;
					width: 100%;
					display: flex;
					align-items: center;
					justify-content: center;

					img {
						height: 100%;
						width: 100%;
						max-inline-size: 100%;
						object-fit: fill;
					}
				}

				.title-favicon-and-domain {
					height: 30%;
					display: flex;
					flex-direction: column;
					align-items: center;
					gap: 0.5em;

					.title {
						height: 50%;
						width: 100%;
						display: flex;
						align-items: center;

						span {
							display: -webkit-box;
							-webkit-line-clamp: 2;
							-webkit-box-orient: vertical;
							overflow: hidden;
							text-overflow: ellipsis;
							font-size: 1.3rem;
							line-height: 1.6;
							color: rgb(24, 23, 40);
							font-family: 'Arial CE', sans-serif;
							text-decoration: underline;
						}
					}

					.favicon-and-domain {
						width: 100%;
						height: 50%;
						display: flex;
						align-items: center;
						gap: 1em;
						//justify-content: space-between;

						img {
							height: 2rem;
							width: 2rem;
							object-fit: cover;
							border-radius: 100vh;
						}

						span {
							color: rgb(24, 23, 40);
							font-size: 1.3rem;
							white-space: nowrap;
							overflow: hidden;
							text-overflow: ellipsis;
							max-width: 90%;
							font-family: 'Arial CE', sans-serif;
						}
					}
				}

				&:hover {
					box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
				}
			}

			.tagsInputAndMatchedTags {
				display: flex;
				flex-direction: column;
				gap: 1em;

				.tags {
					display: flex;
					flex-wrap: wrap;
					align-content: flex-start;
					gap: 0.5em;

					.tag {
						background-color: #91c8e4;
						display: flex;
						align-items: center;
						gap: 0.5em;
						padding: 0.5em;
						border-radius: 0.3rem;

						i {
							font-size: 1.3em;
						}

						span {
							font-family: 'Arial CE', sans-serif;
							font-size: 1.3rem;
						}
					}
				}

				.input {
					height: 3.5rem;
					width: 100%;
					background-color: orange;
					border-radius: 0.3rem;
					position: relative;

					input {
						width: 100%;
						height: 100%;
						padding: 0.5em;
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
						border-radius: inherit;
						outline: none;
						border: 0.1rem solid rgb(0, 0, 0, 0.1);

						&:focus {
							border-color: rgb(6, 143, 255);
						}
					}

					.matchedTags {
						position: absolute;
						z-index: 2;
						top: 101%;
						left: 0;
						right: 0;
						width: inherit;
						min-height: max-content;
						max-height: 40vh;
						overflow-y: auto;
						background-color: rgb(255, 255, 255);
						box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;

						.tag {
							width: 100%;
							height: 4rem;
							padding: 0.5em;
							border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
							display: flex;
							align-items: center;
							gap: 0.5em;
							color: rgb(39, 55, 77);

							i {
								font-size: 1.5rem;
							}

							span {
								font-family: 'Arial CE', sans-serif;
								font-size: 1.3rem;
							}
						}
					}
				}
			}

			button {
				border: none;
				height: 4rem;
				width: 100%;
				background-color: rgb(0, 121, 255);
				display: flex;
				align-items: center;
				justify-content: center;
				border-radius: 0.3rem;
				cursor: pointer;

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					color: rgb(240, 240, 240);
				}

				&:hover {
					background-color: rgb(6, 143, 255);
				}
			}

			.button-disabled {
				opacity: 0.5;
				pointer-events: none;
			}
		}
	}
</style>
