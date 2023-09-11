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
	$: $page.data.folders,
		$page.data.folders.length >= 1 ? folders.set($page.data.folders) : folders.set([]);
	$: $page.data.bookmarks,
		$page.data.bookmarks.length >= 1 ? bookmarks.set($page.data.bookmarks) : bookmarks.set([]);
	$: $page.data.folderPath,
		$page.data.folderPath.length >= 1 ? folderPath.set($page.data.folderPath) : folderPath.set([]);
	$: $folders, sortFoldersByName($folders);
	$: $bookmarks, sortBookmarksByTitle($bookmarks);
</script>

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
	<div class="myBookmarks">
		{#if $loadingItems}
			<div class="loadingItems">
				<div class="loader" />
			</div>
		{:else if !$loadingItems && $allItems.length >= 1}
			<Folder />
			<PlainBookmark />
		{:else if !$loadingItems && $allItems.length < 1}
			<div class="noItems">
				<span>No collections</span>
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
			display: flex;
			height: calc(100vh - calc(8vh + 7vh));
			overflow-y: auto;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			gap: 1em;
			padding: 1em;
		}
	}
</style>
