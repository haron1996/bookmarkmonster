<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/stores';
	import {
		allItems,
		bookmarks,
		folderPath,
		folders,
		loadingItems,
		selectedBookmarks,
		selectedFolders
	} from '../../../stores/stores';
	import { removeSelectedClassFromAllDomFolders } from '../../../utils/removeSelectedClassFromAllDomFolders';
	import { sortBookmarksByTitle } from '../../../utils/sortBookmarks';
	import { sortFoldersByName } from '../../../utils/sortFolders';
	import CreateFolder from '../../CreateFolder.svelte';
	import Folder from '../../Folder.svelte';
	import CreateBookmark from '../../createBookmark.svelte';
	import type { Bookmark } from '../../../types/bookmark';
	import { removeSelectedClassFromAllDomBookmarks } from '../../../utils/removeSelectedClassFromAllSelectedDomBookmarks';
	import TopBar from '../../TopBar.svelte';
	import ActionBarV2 from '../../ActionBarV2.svelte';
	import PlainBookmark from '../../PlainBookmark.svelte';
	import nodata from '$lib/images/no-data.jpg';

	let root: string = 'My collections';

	afterNavigate(() => {
		removeSelectedClassFromAllDomFolders();
		selectedFolders.set([]);
		removeSelectedClassFromAllDomBookmarks();
		selectedBookmarks.set([]);
	});

	function handleChangeInSearchParams() {
		const which: string | null = $page.url.searchParams.get('which');
		if (which) {
			if (which === 'recycle_bin') {
				root = 'Recycle bin';
				return;
			}
			return;
		}

		const id: string | null = $page.url.searchParams.get('id');

		if (id) {
			if (id === 'root') {
				root = 'My collections';
			}
		}
	}

	$: $page.url.searchParams, handleChangeInSearchParams();

	$: $page.data.streamed.folders,
		$page.data.streamed.folders.length >= 1
			? folders.set($page.data.streamed.folders)
			: folders.set([]);

	$: $page.data.streamed.bookmarks,
		$page.data.streamed.bookmarks.length >= 1
			? bookmarks.set($page.data.streamed.bookmarks)
			: bookmarks.set([]);

	$: $page.data.streamed.folderPath,
		$page.data.streamed.folderPath.length >= 1
			? folderPath.set($page.data.streamed.folderPath)
			: folderPath.set([]);

	$: $folders, sortFoldersByName($folders);

	$: $bookmarks, sortBookmarksByTitle($bookmarks);
</script>

<svelte:head>
	<title>My collections | Bookmarkmonster</title>
</svelte:head>

<CreateBookmark />

<CreateFolder />

<div class="app">
	<ActionBarV2 />
	<TopBar />
	<div class="breadcrumb">
		{#if root === 'My collections'}
			<a
				data-sveltekit-preload-data="tap"
				href="/dashboard/my_collections?id=root"
				class:disabled={$folderPath.length < 1}
			>
				{root}
			</a>
		{:else if root === 'Shared'}
			<a
				data-sveltekit-preload-data="tap"
				href="/dashboard/my_collections?id=root&which=shared"
				class:disabled={$folderPath.length < 1}
			>
				{root}
			</a>
		{:else if (root = 'Recycle bin')}
			<a
				data-sveltekit-preload-data="tap"
				href="/dashboard/my_collections?id=root"
				class="disabled"
			>
				{root}
			</a>
		{/if}
		{#each $folderPath as { label, created_at, deleted_at, folder_description, folder_id, folder_name, path, starred, subfolder_of, updated_at, user_id }}
			<i class="las la-long-arrow-alt-right" />
			<a data-sveltekit-preload-data="tap" href={`/dashboard/my_collections?id=${folder_id}`}
				>{folder_name}</a
			>
		{/each}
	</div>
	<div class="myBookmarks" id="myBookmarks">
		{#if $loadingItems}
			<div class="loadingItems">
				<p>Loading...</p>
				<div class="loader" />
			</div>
		{:else if !$loadingItems && $allItems.length >= 1}
			<Folder />
			<PlainBookmark />
		{:else if !$loadingItems && $allItems.length < 1}
			<div class="noItems">
				<img src={nodata} alt="no data" />
				<p>You have no collections yet</p>
			</div>
		{/if}
	</div>
</div>

<style lang="scss">
	.app {
		width: 100%;
		height: 100vh;
		display: flex;
		flex-direction: column;
		background-color: rgb(255, 255, 255);

		.breadcrumb {
			width: 100%;
			padding: 0em 1em;
			display: flex;
			align-items: center;
			gap: 0.5em;
			min-height: 7vh;
			max-width: fit-content;
			white-space: nowrap;
			overflow: hidden;
			text-overflow: ellipsis;
			flex: 1;
			background-color: white;

			i {
				font-size: 1.5rem;
			}

			a {
				font-size: 1.5rem;
				font-family: 'Arial CE', sans-serif;
				text-decoration: none;
				transition: all ease 300ms;
				color: #61677a;
				padding: 0.3em 0.6em;
				border-radius: 0.5rem;

				&:hover {
					background-color: #eeeeee;
				}
			}

			.disabled {
				pointer-events: none;
				text-decoration: none;
				font-family: 'Segoe UI', sans-serif;
				font-size: 1.8rem;
				color: #495464;
				padding: 0 !important;
			}
		}

		.myBookmarks {
			height: calc(100vh - calc(8vh + 7vh));
			overflow-y: auto;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			gap: 1em;
			padding: 1em;

			.loadingItems {
				width: 100%;
				height: 100%;
				display: flex;
				align-items: center;
				justify-content: center;
				flex-direction: column;
				gap: 1em;

				p {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
				}

				.loader {
					border: 2px solid #f3f3f3; /* Light grey */
					border-top: 2px solid #3498db; /* Blue */
					border-radius: 50%;
					width: 3rem;
					height: 3rem;
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
			}

			.noItems {
				width: 100%;
				height: 100%;
				display: flex;
				align-items: center;
				justify-content: center;
				flex-direction: column;

				img {
					// max-inline-size: 100%;
					// object-fit: contain;
					width: 50%;
					height: 50%;
					object-fit: contain;
				}

				p {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
				}
			}
		}
	}
</style>
