<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/stores';
	import {
		allItems,
		bookmarks,
		breadCrumbIsOverflowing,
		collapsedFolders,
		folderPath,
		folders,
		loadingItems,
		selectedBookmarks,
		selectedFolders,
		showCollapsedFolders
	} from '../../../stores/stores';
	import { removeSelectedClassFromAllDomFolders } from '../../../utils/removeSelectedClassFromAllDomFolders';
	import { sortBookmarksByTitle } from '../../../utils/sortBookmarks';
	import { sortFoldersByName } from '../../../utils/sortFolders';
	import CreateFolder from '../../CreateFolder.svelte';
	import Folder from '../../Folder.svelte';
	import CreateBookmark from '../../createBookmark.svelte';
	import { removeSelectedClassFromAllDomBookmarks } from '../../../utils/removeSelectedClassFromAllSelectedDomBookmarks';
	import TopBar from '../../TopBar.svelte';
	import ActionBarV2 from '../../ActionBarV2.svelte';
	import PlainBookmark from '../../PlainBookmark.svelte';
	import nodata from '$lib/images/no-data.jpg';
	import BreadCrumb from '../../BreadCrumb.svelte';
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';

	let root: string = 'My collections';
	let interval: number;

	afterNavigate(() => {
		showCollapsedFolders.set(false);

		breadCrumbIsOverflowing.set(false);

		collapsedFolders.set([]);

		checkIfBreadCrumbIsOverflowing();

		removeSelectedClassFromAllDomFolders();

		selectedFolders.set([]);

		removeSelectedClassFromAllDomBookmarks();

		selectedBookmarks.set([]);
	});

	function checkIfBreadCrumbIsOverflowing() {
		const breadcrumb = document.getElementById('breadCrumb') as HTMLDivElement | null;

		if (breadcrumb === null) return;

		const clientWidth = breadcrumb.clientWidth;
		const scrollWidth = breadcrumb.scrollWidth;

		breadCrumbIsOverflowing.set(scrollWidth > clientWidth);
	}

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

	function collapseBreadCrumb() {
		interval = setInterval(() => {
			if ($folderPath.length >= 1) {
				collapsedFolders.update((values) => [...values, $folderPath.slice(0, 1)[0]]);

				$folderPath = $folderPath.filter((value) => {
					return value.folder_id !== $folderPath.slice(0, 1)[0].folder_id;
				});

				folderPath.set($folderPath);

				checkIfBreadCrumbIsOverflowing();
			}
		}, 0.1);
	}

	function removeInterval() {
		clearInterval(interval);
	}

	$: $page.url.searchParams, handleChangeInSearchParams();

	$: $page.data.folders,
		$page.data.folders.length >= 1 ? folders.set($page.data.folders) : folders.set([]);

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

	$: $breadCrumbIsOverflowing ? collapseBreadCrumb() : removeInterval();
</script>

<svelte:head>
	<title>My collections | Bookmarkmonster</title>
</svelte:head>

<CreateBookmark />

<CreateFolder />

<div class="app">
	<ActionBarV2 />
	<TopBar />
	<BreadCrumb />
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

		.myBookmarks {
			height: calc(100vh - calc(10rem));
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
