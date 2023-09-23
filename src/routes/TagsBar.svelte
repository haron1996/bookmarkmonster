<script lang="ts">
	import { browser } from '$app/environment';
	import {
		apiHost,
		currentTag,
		deletedTag,
		selectedTag,
		session,
		tags,
		tagsBarCollapsed
	} from '../stores/stores';
	import type { Tag } from '../types/tag';
	import RenameTagV2 from './RenameTagV2.svelte';

	let searchingTags: boolean = false;
	let showNewTagForm: boolean = false;
	let tagName: string = 'Untitled tag';
	let tagQuery: string = '';
	let formError: boolean = false;
	let searchFormError: boolean = false;

	function selectTagInput() {
		tagName = 'Untitled tag';

		setTimeout(() => {
			const input = document.getElementById('tagInput') as HTMLInputElement | null;

			if (input === null) return;

			input.select();
		}, 10);
	}

	// function setCurrentTag() {
	// 	if (!browser) return;

	// 	setTimeout(() => {
	// 		const allTags = document.querySelectorAll('div.tag') as NodeListOf<HTMLDivElement> | null;

	// 		if (allTags === null || allTags.length < 1) return;

	// 		const t: Tag = {
	// 			added: allTags[0].dataset.added,
	// 			deleted: allTags[0].dataset.deleted,
	// 			id: allTags[0].dataset.id,
	// 			name: allTags[0].dataset.name,
	// 			updated: allTags[0].dataset.updated,
	// 			user_id: allTags[0].dataset.userid
	// 		};

	// 		currentTag.set(t);

	// 		allTags[0].classList.add('activeTag');
	// 	}, 1000);
	// }

	async function handleNewTagFormSubmit() {
		if (tagName === 'Untitled tag' || tagName === '') {
			formError = true;
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

		const result = await response.json();

		tags.update((values) => [result[0], ...values]);

		selectTagInput();

		setTimeout(() => {
			addActiveClassToCurrentTag();
		}, 100);
	}

	function handleClickOnTag(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('.tag') as HTMLDivElement | null;

		if (tag === null) return;

		const allTags = document.querySelectorAll('div.tag') as NodeListOf<HTMLDivElement> | null;

		if (allTags === null || allTags.length < 1) return;

		allTags.forEach((tag) => {
			tag.classList.remove('activeTag');
		});

		tag.classList.add('activeTag');

		const t: Tag = {
			added: tag.dataset.added,
			deleted: tag.dataset.deleted,
			id: tag.dataset.id,
			name: tag.dataset.name,
			updated: tag.dataset.updated,
			user_id: tag.dataset.userid
		};

		currentTag.set(t);

		window.localStorage.setItem('currentTag', JSON.stringify(t));
	}

	function handleClickOnPenIcon(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('.tag') as HTMLDivElement | null;

		if (tag === null) return;

		const t: Tag = {
			added: tag.dataset.added,
			deleted: tag.dataset.deleted,
			id: tag.dataset.id,
			name: tag.dataset.name,
			updated: tag.dataset.updated,
			user_id: tag.dataset.userid
		};

		selectedTag.set(t);

		const form = document.getElementById('renameTagV2') as HTMLFormElement | null;

		if (form === null) return;

		const tagRect = tag.getBoundingClientRect();

		form.style.left = `${tagRect.left}px`;

		form.style.top = `${tagRect.top}px`;

		form.style.minHeight = `${tag.offsetHeight}px`;

		form.style.width = `${tag.offsetWidth}px`;

		form.style.display = 'flex';

		const input = form.firstElementChild as HTMLInputElement | null;

		if (input === null) return;

		setTimeout(() => {
			input.select();
		}, 10);

		const tags = document.getElementById('tags') as HTMLDivElement | null;

		if (tags === null) return;

		tags.style.overflowY = 'hidden';
	}

	async function getUserTags() {
		const response = await fetch(`${$apiHost}/authenticated/tags`, {
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

		const msg = result.message;

		if (msg) {
			console.log(msg);
			return;
		}

		tags.set(result[0]);

		//setCurrentTag();
	}

	async function handleClickOnTrashIcon(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const tag = target.closest('.tag') as HTMLDivElement | null;

		if (tag === null) return;

		const t: Tag = {
			added: tag.dataset.added,
			deleted: tag.dataset.deleted,
			id: tag.dataset.id,
			name: tag.dataset.name,
			updated: tag.dataset.updated,
			user_id: tag.dataset.userid
		};

		selectedTag.set(t);

		deletedTag.set(t);

		const response = await fetch(`${$apiHost}/authenticated/tags/trashTag`, {
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
			body: JSON.stringify({ tag: $selectedTag })
		});

		const result = await response.json();

		const allTags = document.querySelectorAll('div.tag') as NodeListOf<HTMLDivElement>;

		if (allTags.length < 1) return;

		allTags.forEach((tag) => {
			if (tag.dataset.id === result[0].id) {
				tag.classList.add('animate__animated');

				tag.classList.add('animate__bounceOut');
			}
		});

		setTimeout(() => {
			$tags = $tags.filter((value) => {
				return value.id !== result[0].id;
			});

			tags.set($tags);

			getUserTags();

			const allTags = document.querySelectorAll('div.tag') as NodeListOf<HTMLDivElement>;

			if (allTags.length < 1) {
				selectedTag.set({});
				return;
			}

			allTags.forEach((tag) => {
				tag.classList.remove('animate__animated');
				tag.classList.remove('animate__bounceOut');
				tag.classList.remove('activeTag');
			});

			selectedTag.set({});
		}, 2000);

		setTimeout(() => {
			allTags.forEach((tag) => {
				if (tag.dataset.id === $currentTag.id) {
					tag.classList.add('activeTag');
				}
			});
		}, 2000);
	}

	async function searchUserTags() {
		searchingTags = true;

		if (tagQuery === '') {
			getUserTags();
			searchingTags = false;
			setTimeout(() => {
				addActiveClassToCurrentTag();
			}, 100);
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/${tagQuery}`, {
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

		const msg = result.message;

		if (msg) {
			alert(msg);
			searchingTags = false;
			return;
		}

		tags.set(result[0]);

		searchingTags = false;

		const allTags = document.querySelectorAll('div.tag') as NodeListOf<HTMLDivElement> | null;

		if (allTags === null || allTags.length < 1) return;

		allTags.forEach((tag) => {
			tag.classList.remove('activeTag');
		});
	}

	function handleSearchInputKeydown(e: KeyboardEvent) {
		if (e.code === 'Space' && tagQuery === '') {
			e.preventDefault();
			return;
		}
	}

	function addActiveClassToCurrentTag() {
		const domTags = document.querySelectorAll('.tag') as NodeListOf<HTMLDivElement>;

		if (domTags.length < 1) return;

		domTags.forEach((dt) => {
			dt.classList.remove('activeTag');
			if (dt.dataset.id === $currentTag.id) {
				dt.classList.add('activeTag');
			}
		});
	}

	$: showNewTagForm ? selectTagInput() : () => {};
	$: formError
		? setTimeout(() => {
				formError = false;
		  }, 4000)
		: () => {};
	$: searchFormError
		? setTimeout(() => {
				searchFormError = false;
		  }, 4000)
		: () => {};
</script>

<RenameTagV2 />

{#if !$tagsBarCollapsed}
	<div class="tagsBar" id="tagsBar">
		<div class="top">
			<form
				class="search"
				class:searchFormError
				id="searchForm"
				on:submit|preventDefault={searchUserTags}
			>
				<div class="searchInputContainer">
					<input
						type="search"
						name="search"
						id="searchInput"
						placeholder="Search tags..."
						autocomplete="off"
						bind:value={tagQuery}
						on:keydown={handleSearchInputKeydown}
						on:input={searchUserTags}
					/>
				</div>
				{#if tagQuery !== ''}
					<div class="loadOrClear">
						{#if searchingTags}
							<div class="loader" />
						{:else}
							<i
								class="las la-times"
								role="none"
								on:click={() => {
									tagQuery = '';
									getUserTags();
									setTimeout(() => {
										addActiveClassToCurrentTag();
									}, 100);
								}}
							/>
						{/if}
					</div>
				{:else}
					<div class="collapseTagsBar">
						<i
							class="las la-angle-double-left"
							role="none"
							on:click={() => {
								$tagsBarCollapsed = !$tagsBarCollapsed;
							}}
						/>
					</div>
				{/if}
			</form>
		</div>
		<div class="tags" id="tags">
			{#if $tags}
				{#each $tags as { added, deleted, id, name, updated, user_id }}
					<div
						class="tag"
						role="none"
						on:click={handleClickOnTag}
						data-added={added}
						data-deleted={deleted}
						data-id={id}
						data-updated={updated}
						data-userid={user_id}
						data-name={name}
					>
						<i class="las la-hashtag" />
						<span>{name}</span>
						<div class="options">
							<i class="las la-pen" role="none" on:click|stopPropagation={handleClickOnPenIcon} />
							<i
								class="las la-trash"
								role="none"
								on:click|stopPropagation={handleClickOnTrashIcon}
							/>
						</div>
					</div>
				{/each}
			{/if}
		</div>
		<div class="addTag">
			{#if showNewTagForm}
				<form class:formError on:submit|preventDefault={handleNewTagFormSubmit}>
					<input
						type="text"
						name="tag"
						id="tagInput"
						placeholder="Enter tag name..."
						bind:value={tagName}
						autocomplete="off"
					/>
					<div class="close">
						<i
							class="las la-times"
							role="none"
							on:click={() => {
								showNewTagForm = false;
							}}
						/>
					</div>
				</form>
			{:else}
				<button
					on:click|preventDefault={() => {
						showNewTagForm = true;
					}}
				>
					<i class="las la-plus" />
					<span>New Tag</span>
				</button>
			{/if}
		</div>
	</div>
{/if}

<style lang="scss">
	.tagsBar {
		width: 40rem;
		background-color: #f0f0f0;
		display: flex;
		flex-direction: column;
		border-right: 0.1rem solid rgb(4, 13, 18, 0.1);
		transition: all 300ms ease-in-out;

		.top {
			height: max-content;
			width: 100%;
			border-bottom: 0.1rem solid rgb(4, 13, 18, 0.1);
			transition: all 300ms ease-in-out;

			form {
				display: flex;
				background-color: inherit;
				height: 4rem;

				.searchInputContainer {
					height: 100%;
					background-color: inherit;
					width: 85%;

					input {
						height: 100%;
						width: 100%;
						padding: 0.5em;
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
						border: none;
						outline: none;
						background-color: inherit;

						&::-webkit-search-cancel-button {
							-webkit-appearance: none;
						}
					}
				}

				.loadOrClear {
					width: 15%;
					display: flex;
					align-items: center;
					justify-content: center;
					background-color: inherit;

					.loader {
						border: 0.1rem solid #f3f3f3;
						border-top: 0.1rem solid #3498db;
						border-bottom: 0.1rem solid #3498db;
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

					i {
						font-size: 1.5rem;
						cursor: pointer;
					}
				}

				.collapseTagsBar {
					width: 15%;
					display: flex;
					align-items: center;
					justify-content: center;
					background-color: inherit;

					i {
						font-size: 1.5rem;
						cursor: pointer;
					}
				}

				&:focus-within {
					.collapseTagsBar {
						i {
							display: none;
						}
					}
				}
			}

			.searchFormError {
				border: 0.2rem solid red;
			}

			&:focus-within {
				background-color: white;
			}
		}

		.tags {
			width: 100%;
			height: 100%;
			overflow-y: auto;
			background-color: inherit;
			display: flex;
			flex-direction: column;

			.tag {
				min-height: 4rem;
				border-bottom: 0.1rem solid rgb(4, 13, 18, 0.1);
				display: flex;
				align-items: center;
				gap: 0.5em;
				padding: 0 0.5em;
				cursor: pointer;
				width: 100%;
				position: relative;

				i {
					font-size: 1.5rem;
					transition: all 300ms ease;
				}

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
					white-space: nowrap;
					overflow: hidden;
					text-overflow: ellipsis;

					@media only screen and (width <= 600px) {
						max-width: 50%;
					}
				}

				.options {
					display: flex;
					align-items: center;
					gap: 0.5em;
					opacity: 0;
					transition: all 300ms ease;

					i {
						&:hover {
							color: #0079ff;
						}
					}

					@media only screen and (width <= 768px) {
						opacity: 1;
					}
				}

				&:hover {
					.options {
						opacity: 1;
					}
				}

				@media only screen and (width <= 768px) {
					gap: 1em;
				}
			}
		}

		.addTag {
			height: max-content;
			width: 100%;
			border-top: 0.1rem solid rgb(4, 13, 18, 0.1);

			button {
				width: 100%;
				height: 4rem;
				padding: 0em 0.5em;
				display: flex;
				align-items: center;
				justify-content: start;
				gap: 0.5em;
				border: none;
				cursor: pointer;
				transition: all 300ms ease;

				i {
					font-size: 1.5rem;
				}

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
				}
			}

			form {
				display: flex;
				align-items: center;
				width: 100%;
				background-color: inherit;

				input {
					width: 85%;
					height: 4rem;
					padding: 0em 0.5em;
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
					background-color: white;
					outline: none;
					border: none;
					transition: all 300ms ease-in-out;
					text-transform: lowercase;
				}

				.close {
					height: 4rem;
					width: 15%;
					display: flex;
					align-items: center;
					justify-content: center;

					i {
						font-size: 1.5rem;
						cursor: pointer;
					}
				}
			}

			.formError {
				input {
					border: 0.2rem solid red;
				}
			}
		}

		// upto 600px
		@media only screen and (width <= 600px) {
			min-width: 100%;
			max-width: 100%;
		}

		// 600 - 768px
		@media only screen and (601px <= width <= 768px) {
			min-width: 50%;
			max-width: 50%;
		}
	}

	:global(.activeTag) {
		i.la-hashtag {
			color: #0079ff !important;
		}
		span {
			color: #0079ff !important;
		}
	}
</style>
