<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import {
		apiHost,
		bookmarks,
		currentTag,
		errorMessage,
		hideSideBar,
		query,
		session,
		showCreatBookmarkFromTagPage,
		successMessage,
		tags,
		tagsBarCollapsed
	} from '../../../stores/stores';
	import { sortBookmarksByTitle } from '../../../utils/sortBookmarks';
	import TagsBar from '../../TagsBar.svelte';
	import NewBookmarkFromTagsPage from './NewBookmarkFromTagsPage.svelte';
	import Mark from 'mark.js';

	afterNavigate(() => {
		getUserTags();
		getAndSetActiveTag();
		setTimeout(() => {
			addActiveClassToCurrentTag();
		}, 100);
	});

	let searchingBookmarks: boolean = false;

	function getAndSetActiveTag() {
		const currentTagString = localStorage.getItem('currentTag');

		if (currentTagString === null) return;

		currentTag.set(JSON.parse(currentTagString));
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
	}

	function handleSearchInputKeydown(e: KeyboardEvent) {
		if (e.code === 'Space' && $query === '') {
			e.preventDefault();
			return;
		}

		if (e.code === 'Backspace') {
			// handle search input clear by backspace
		}
	}

	async function searchUserTagBookmarks() {
		if ($query === '') {
			getBookmarksOfCurrentTag();
			return;
		}

		const response = await fetch(
			`${$apiHost}/authenticated/tags/searchTagBookmarks/${$currentTag.id}/${$query}`,
			{
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
			}
		);

		const result = await response.json();

		const message = result.message;

		if (message) {
			errorMessage.set(message);
			return;
		}

		bookmarks.set(result[0]);

		const context = document.getElementById('actualBookmarks') as HTMLDivElement | null;

		if (context === null) return;

		const markInstance = new Mark(context);

		markInstance.unmark({
			done: () => {
				markInstance.mark($query, {
					accuracy: 'exactly',
					caseSensitive: false,
					done: () => {},
					separateWordSearch: false,
					diacritics: true
				});
			}
		});
	}

	async function getBookmarksOfCurrentTag() {
		console.log('geting bookmarks of tag: ' + $currentTag.name + '...');
		const response = await fetch(`${$apiHost}/authenticated/bookmarks/${$currentTag.id}`, {
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
			return;
		}

		if (result[0]) {
			bookmarks.set(result[0]);
		} else {
			bookmarks.set([]);
		}
	}

	function addActiveClassToCurrentTag() {
		const domTags = document.querySelectorAll('.tag') as NodeListOf<HTMLDivElement>;

		if (domTags.length < 1) return;

		domTags.forEach((dt) => {
			if (dt.dataset.id === $currentTag.id) {
				dt.classList.add('activeTag');
			}
		});
	}

	function clearMarkJs() {
		const context = document.getElementById('actualBookmarks') as HTMLDivElement | null;

		if (context === null) return;

		const markInstance = new Mark(context);

		markInstance.unmark();
	}

	$: $currentTag.id ? getBookmarksOfCurrentTag() : () => {};
	$: $bookmarks, $bookmarks ? sortBookmarksByTitle($bookmarks) : () => {};
</script>

<svelte:head>
	<title>Tags | Bookmarkmonster</title>
</svelte:head>

<NewBookmarkFromTagsPage />

<div class="app">
	<TagsBar />
	<div class="bookmarks">
		<div class="top">
			<div class="left">
				{#if $tagsBarCollapsed}
					<i
						class="las la-angle-double-right"
						role="none"
						on:click={() => {
							setTimeout(() => {
								addActiveClassToCurrentTag();
							}, 100);
							$tagsBarCollapsed = !$tagsBarCollapsed;
						}}
					/>
				{/if}
				<i
					class="las la-bars"
					role="none"
					on:click={() => {
						$hideSideBar = !$hideSideBar;
					}}
				/>
				<i
					class="las la-sync"
					role="none"
					on:click={() => {
						getBookmarksOfCurrentTag;
						successMessage.set('Bookmarks have been refreshed');
					}}
				/>
			</div>
			<form class="search" id="searchForm" on:submit|preventDefault={searchUserTagBookmarks}>
				<div class="searchInputContainer">
					<input
						type="search"
						name="search"
						id="searchInput"
						placeholder={`Search in ${$currentTag.name ? `${$currentTag.name}` : '....'}`}
						autocomplete="off"
						bind:value={$query}
						on:keydown={handleSearchInputKeydown}
						on:input={searchUserTagBookmarks}
					/>
				</div>
				{#if $query !== ''}
					<div class="loadOrClear">
						{#if searchingBookmarks}
							<div class="loader" />
						{:else}
							<i
								class="las la-times"
								role="none"
								on:click={() => {
									query.set('');
									getBookmarksOfCurrentTag();
									clearMarkJs();
								}}
							/>
						{/if}
					</div>
				{/if}
			</form>
			<div class="right">
				<i
					class="las la-plus"
					role="none"
					on:click={() => {
						showCreatBookmarkFromTagPage.set(true);
					}}
				/>
			</div>
		</div>
		<div class="actualBookmarks" id="actualBookmarks">
			{#if $bookmarks}
				{#each $bookmarks as { added, bookmark, deleted, favicon, host, id, notes, thumbnail, title, updated, user_id }}
					<div class="bookmark">
						<a href={bookmark} target="_blank">
							{title}
						</a>
						<div class="actions">
							<i class="las la-pen" />
							<i class="las la-share" />
							<i class="las la-trash-alt" />
							<i class="las la-tag" />
							<i class="las la-ellipsis-h" />
						</div>
					</div>
				{/each}
			{/if}
		</div>
	</div>
</div>

<style lang="scss">
	.app {
		width: 100%;
		height: 100%;
		display: flex;

		.bookmarks {
			width: 100%;
			//border: 0.2rem solid red;
			display: flex;
			flex-direction: column;

			.top {
				height: 5rem;
				width: 100%;
				background-color: white;
				display: flex;
				align-items: center;

				.left {
					background-color: inherit;
					height: 100%;
					flex: 1;
					display: flex;
					align-items: center;
					padding-left: 0.5em;
					gap: 1em;

					i {
						font-size: 1.5rem;
						cursor: pointer;
						color: #272829;
					}

					i.la-bars,
					i.la-sync {
						font-size: 2rem;
					}
				}

				form {
					display: flex;
					height: 3.5rem;
					flex: 1;
					align-items: center;
					background-color: white;
					border-radius: 0.3rem;
					width: 40rem;
					max-width: 60rem;

					.searchInputContainer {
						height: 100%;
						background-color: inherit;
						width: 85%;
						border-radius: inherit;

						input {
							height: 100%;
							width: 100%;
							padding: 0.5em;
							padding-left: 1em;
							font-family: 'Arial CE', sans-serif;
							font-size: 1.3rem;
							border: none;
							outline: none;
							background-color: inherit;
							border-radius: inherit;

							&::placeholder {
								color: #d8d8d8;
							}

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
						border-radius: inherit;

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
				}

				.right {
					background-color: inherit;
					height: 100%;
					flex: 1;
					display: flex;
					align-items: center;
					justify-content: flex-end;
					padding-right: 0.5em;

					i {
						font-size: 2rem;
						cursor: pointer;
					}
				}
			}

			.actualBookmarks {
				width: 100%;
				height: max-content;
				background-color: white;
				display: flex;
				flex-direction: column;
				height: 100%;
				overflow-x: auto;
				padding: 0.1em;

				.bookmark {
					height: max-content;
					display: flex;
					align-items: center;
					background-color: inherit;
					padding: 0.5em;
					gap: 1em;

					a {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						display: -webkit-box;
						-webkit-box-orient: vertical;
						-webkit-line-clamp: 1;
						overflow: hidden;
						text-decoration: underline;
						color: rgb(4, 13, 18);
						transition: all 300ms ease;

						&:hover {
							color: #75c2f6;
						}
					}

					.actions {
						display: flex;
						align-items: center;
						gap: 1em;
						display: none;

						i {
							font-size: 1.6rem;
							cursor: pointer;
							transform: all 300ms ease;

							&:hover {
								color: #75c2f6;
							}
						}
					}

					&:hover {
						.actions {
							display: flex;
						}
					}
				}
			}
		}
	}
</style>
