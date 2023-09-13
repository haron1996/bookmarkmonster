<script lang="ts">
	import UpdateFolder from '../UpdateFolder.svelte';
	import SideBar from '../SideBar.svelte';
	import UpdateBookmark from '../UpdateBookmark.svelte';
	import { afterNavigate } from '$app/navigation';
	import {
		destinationFolder,
		loadingItems,
		destinationFolderPath,
		showMoveItemsPopup,
		childrenOfDestinationFolder,
		bookmarksOfDestinationFolder,
		selectedFolders,
		selectedBookmarks
	} from '../../stores/stores';
	import NewBookmark from '../NewBookmark.svelte';
	import MoveItems from '../MoveItems.svelte';
	import { page } from '$app/stores';
	import { getUserRootFoldersAndPlainBookmarks } from '../../utils/getRootFoldersAndPlainBookmarks';
	import type { Folder } from '../../types/folder';
	import type { Bookmark } from '../../types/bookmark';
	import { getFolderSubFoldersAndBookmarks } from '../../utils/getFolderSubfoldersAndBookmarks';
	import { getFolder } from '../../utils/getFolder';
	import { getFolderPath } from '../../utils/getFolderPath';
	import { browser } from '$app/environment';
	import { sortFoldersByName } from '../../utils/sortFolders';

	let f: Folder = {};
	let fs: Folder[] = [];
	let bs: Bookmark[] = [];

	afterNavigate(() => {
		loadingItems.set(false);
	});

	async function getDestinationFolder() {
		if (browser) {
			const id: string | null = $page.url.searchParams.get('id');

			if (id) {
				if (id === 'root') {
					destinationFolder.set({});
				} else {
					f = await getFolder(id);

					destinationFolder.set(f);
				}
			} else {
				destinationFolder.set({});
			}
		}
	}

	function resetMoveItems() {
		destinationFolder.set({});
		destinationFolderPath.set([]);
		childrenOfDestinationFolder.set([]);
		bookmarksOfDestinationFolder.set([]);
	}

	async function getDestinationFolderPathSubfoldersAndBookmarks(id: string) {
		fs = await getFolderPath(id, fetch);

		destinationFolderPath.set(fs);

		if (browser) {
			[fs, bs] = await getFolderSubFoldersAndBookmarks(fetch, id);
			fs.length >= 1 ? childrenOfDestinationFolder.set(fs) : childrenOfDestinationFolder.set([]);
			bs.length >= 1 ? bookmarksOfDestinationFolder.set(bs) : bookmarksOfDestinationFolder.set([]);
		}
	}

	async function setDestinationFolderPathEmptyAndGetRootFoldersAndBookmarks() {
		destinationFolderPath.set([]);

		if (browser) {
			[fs, bs] = await getUserRootFoldersAndPlainBookmarks(fetch);
			fs.length >= 1 ? childrenOfDestinationFolder.set(fs) : childrenOfDestinationFolder.set([]);
			bs.length >= 1 ? bookmarksOfDestinationFolder.set(bs) : bookmarksOfDestinationFolder.set([]);
		}
	}

	function handleChangeInChildrenOfDestinationFolder() {
		$childrenOfDestinationFolder.forEach((value) => {
			$selectedFolders.forEach((f) => {
				if (value.folder_id === f.folder_id) {
					// disable moving folder itself
					setTimeout(() => {
						const domFolders = document.querySelectorAll(
							'.folderV2'
						) as NodeListOf<HTMLAnchorElement> | null;

						if (domFolders === null) return;

						domFolders.forEach((df) => {
							const id = df.dataset.id;
							const name = df.dataset.foldername;

							if (id === undefined) return;

							if (name === undefined) return;

							df.classList.remove('blocked');

							if (value.folder_id === id) {
								df.classList.add('blocked');
							}
						});
					}, 10);
				}
			});
		});

		// disable moving folder to its current location
		// ie disable move button when children of destination folder include selected folder
		$selectedFolders.forEach((value) => {
			const f = $childrenOfDestinationFolder.filter((v) => {
				return v.folder_id === value.folder_id;
			});

			const fn = $childrenOfDestinationFolder.filter((v) => {
				return v.folder_name === value.folder_name;
			});

			const btn = document.getElementById('moveButton') as HTMLButtonElement | null;

			if (btn === null) return;

			if (f.length >= 1 || fn.length >= 1) {
				btn.classList.add('moveButtonDisabled');
			} else {
				btn.classList.remove('moveButtonDisabled');
			}
		});

		$selectedBookmarks.forEach((value) => {
			const f = $bookmarksOfDestinationFolder.filter((v) => {
				return v.id === value.id;
			});

			const btn = document.getElementById('moveButton') as HTMLButtonElement | null;

			if (btn === null) return;

			if (f.length >= 1) {
				btn.classList.add('moveButtonDisabled');
			} else {
				btn.classList.remove('moveButtonDisabled');
			}
		});
	}

	// listen to showMoveItems store event
	$: $showMoveItemsPopup, $showMoveItemsPopup ? getDestinationFolder() : resetMoveItems();
	// listen to change in destination folder
	$: $destinationFolder,
		$destinationFolder.folder_id !== undefined
			? getDestinationFolderPathSubfoldersAndBookmarks($destinationFolder.folder_id)
			: setDestinationFolderPathEmptyAndGetRootFoldersAndBookmarks();
	// get path to destination folder(ancestors)
	// get children folders of the destination folder
	// get bookmarks of the destination folder
	$: $childrenOfDestinationFolder, handleChangeInChildrenOfDestinationFolder();
	$: $childrenOfDestinationFolder, sortFoldersByName($childrenOfDestinationFolder);
</script>

<div class="app">
	<NewBookmark />
	<UpdateFolder />
	<UpdateBookmark />
	<MoveItems />
	<div class="bottom">
		<SideBar />
		<div class="content">
			<slot />
		</div>
	</div>
</div>

<style lang="scss">
	@import url('https://fonts.cdnfonts.com/css/arial');
	@import url('https://fonts.cdnfonts.com/css/segoe-ui-4');
	.app {
		position: fixed;
		left: 0;
		top: 0;
		height: 100vh;
		width: 100vw;
		overflow: hidden;
		display: flex;
		flex-direction: column;
		background-color: rgb(250, 250, 250);

		.bottom {
			display: flex;

			.content {
				height: 100%;
				width: calc(100vw - 9rem);
			}
		}
	}
</style>
