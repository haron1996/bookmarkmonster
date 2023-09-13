<script lang="ts">
	import {
		apiHost,
		bookmarks,
		bookmarksOfDestinationFolder,
		childrenOfDestinationFolder,
		destinationFolder,
		destinationFolderPath,
		folders,
		selectedBookmarks,
		selectedFolders,
		selectedItems,
		session,
		showCreateFolderFromMoveFoldersPopup,
		showMoveItemsPopup
	} from '../stores/stores';
	import type { Bookmark } from '../types/bookmark';
	import type { Folder } from '../types/folder';
	import { removeSelectedClassFromAllDomFolders } from '../utils/removeSelectedClassFromAllDomFolders';
	import { removeSelectedClassFromAllDomBookmarks } from '../utils/removeSelectedClassFromAllSelectedDomBookmarks';
	import { removeSlideInAnimationFromActionBarV2 } from '../utils/removeSlideInAnimationFromActionBarV2';
	import FolderV2 from './FolderV2.svelte';
	import NewCollectionFromMoveFolders from './NewCollectionFromMoveFolders.svelte';
	import PlainBookmarkV2 from './PlainBookmarkV2.svelte';

	function handleClickOnMoveFoldersBreadcrumbElement(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const span = t.closest('span') as HTMLSpanElement | null;

		if (span === null) return;

		const f: Folder = {
			created_at: span.dataset.createdat,
			deleted_at: span.dataset.deletedat,
			folder_description: span.dataset.folderdescription,
			folder_id: span.dataset.folderid,
			folder_name: span.dataset.foldername,
			label: span.dataset.label,
			path: span.dataset.path,
			starred: span.dataset.starred ? true : false,
			subfolder_of: span.dataset.subfolderof,
			updated_at: span.dataset.updatedat,
			user_id: span.dataset.userid
		};

		destinationFolder.set(f);

		doShit();
	}

	async function moveCollectionsToAnother() {
		const response = await fetch(`${$apiHost}/authenticated/collections/moveCollectionsToAnother`, {
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
			body: JSON.stringify({
				destination_folder: $destinationFolder,
				folders_to_move: $selectedFolders
			})
		});

		const result = await response.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			return;
		}

		const fs: Folder[] = result[0];

		fs.forEach((f) => {
			$folders = $folders.filter((value) => {
				return value.folder_id !== f.folder_id;
			});

			folders.set($folders);
		});

		selectedFolders.set([]);

		removeSelectedClassFromAllDomFolders();

		removeSlideInAnimationFromActionBarV2();

		showMoveItemsPopup.set(false);
	}

	async function moveCollectionsToRoot() {
		const response = await fetch(`${$apiHost}/authenticated/collections/moveCollectionsToRoot`, {
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
			body: JSON.stringify({
				collections: $selectedFolders
			})
		});

		const result = await response.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			return;
		}

		const fs: Folder[] = result[0];

		fs.forEach((f) => {
			$folders = $folders.filter((value) => {
				return value.folder_id !== f.folder_id;
			});

			folders.set($folders);
		});

		selectedFolders.set([]);

		removeSelectedClassFromAllDomFolders();

		removeSlideInAnimationFromActionBarV2();

		showMoveItemsPopup.set(false);
	}

	async function moveBookmarks() {
		const response = await fetch(`${$apiHost}/authenticated/bookmarks/moveBookmarks`, {
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
			body: JSON.stringify({
				bookmarks: $selectedBookmarks,
				destination: $destinationFolder
			})
		});

		const result = await response.json();

		const bs: Bookmark[] = result[0];

		bs.forEach((b) => {
			$bookmarks = $bookmarks.filter((value) => {
				return value.id !== b.id;
			});

			bookmarks.set($bookmarks);
		});

		selectedBookmarks.set([]);

		removeSelectedClassFromAllDomBookmarks();

		removeSlideInAnimationFromActionBarV2();

		showMoveItemsPopup.set(false);
	}

	async function handleClickOnMoveButton() {
		if ($selectedFolders.length >= 1 && $selectedBookmarks.length >= 1) {
			// alert('move folders and bookmarks');
			if ($destinationFolder.folder_id === undefined) {
				// alert('move folders to root');
				moveCollectionsToRoot();
				// alert('move bookmarks');
				moveBookmarks();
			} else {
				// alert('move folders to another');
				moveCollectionsToAnother();
				// alert('move bookmarks');
				moveBookmarks();
			}
		} else if ($selectedFolders.length >= 1 && $selectedBookmarks.length < 1) {
			// alert('move folders');
			if ($destinationFolder.folder_id === undefined) {
				// alert('move folders to root');
				moveCollectionsToRoot();
			} else {
				// alert('move folders to another');
				moveCollectionsToAnother();
			}
		} else if ($selectedFolders.length < 1 && $selectedBookmarks.length >= 1) {
			// alert('move bookmarks');
			moveBookmarks();
		}
	}

	function doShit() {
		const domFolders = document.querySelectorAll(
			'.folderV2'
		) as NodeListOf<HTMLAnchorElement> | null;

		if (domFolders === null) return;

		domFolders.forEach((df) => {
			const id = df.dataset.id;

			if (id === undefined) return;

			df.classList.remove('blocked');
		});
	}
</script>

<NewCollectionFromMoveFolders />

{#if $showMoveItemsPopup}
	<div class="container animate__animated animate__zoomInDown">
		<div class="breadcrumb">
			<span
				role="none"
				on:click={() => {
					destinationFolder.set({});
					doShit();
				}}>My collections</span
			>
			{#each $destinationFolderPath as { created_at, deleted_at, folder_description, folder_id, folder_name, label, path, starred, subfolder_of, updated_at, user_id }}
				<i class="las la-long-arrow-alt-right" />
				<span
					data-createdat={created_at}
					data-deleted={deleted_at}
					data-folderdescription={folder_description}
					data-folderid={folder_id}
					data-label={label}
					data-path={path}
					data-starred={starred}
					data-subfolderof={subfolder_of}
					data-updatedat={updated_at}
					data-foldername={folder_name}
					data-userid={user_id}
					role="none"
					on:click={handleClickOnMoveFoldersBreadcrumbElement}>{folder_name}</span
				>
			{/each}
		</div>
		<div class="items">
			<FolderV2 />
			<PlainBookmarkV2 />
		</div>
		<div class="bottom">
			<div class="left">
				<span>Move {$selectedItems.length} {$selectedItems.length > 1 ? 'items' : 'item'}</span>
			</div>
			<div class="right">
				<button on:click|preventDefault={handleClickOnMoveButton} id="moveButton">
					<span>Move here</span>
				</button>
				<button
					on:click|preventDefault={() => {
						showCreateFolderFromMoveFoldersPopup.set(true);
					}}
				>
					<span>New collection</span>
				</button>
				<button
					on:click|preventDefault={() => {
						showMoveItemsPopup.set(false);
					}}
				>
					<span>Cancel</span>
				</button>
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	.container {
		min-height: 100%;
		max-height: 100%;
		width: 100%;
		background-color: white;
		display: flex;
		flex-direction: column;

		.breadcrumb {
			min-height: 10vh;
			// background-color: lightblue;
			display: flex;
			align-items: center;
			padding: 1em;
			border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
			gap: 0.5em;
			overflow-y: auto;

			i {
				font-size: 1.3rem;
			}

			span {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.3rem;
				color: #040d12;
				cursor: pointer;
				padding: 0.3em 0.5em;
				border-radius: 0.3rem;

				&:last-of-type {
					pointer-events: none;
				}

				&:hover {
					background-color: #eeeeee;
				}
			}
		}

		.items {
			min-height: 80vh;
			max-height: 80vh;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			gap: 1em;
			padding: 1em;
			overflow-y: auto;
		}

		.bottom {
			min-height: 10vh;
			// background-color: lightblue;
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 1em;
			border-top: 0.1rem solid rgb(0, 0, 0, 0.1);

			.left {
				display: flex;
				align-items: center;
				width: max-content;
				max-width: 50%;

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
					color: #040d12;
				}
			}

			.right {
				display: flex;
				align-items: center;
				width: max-content;
				gap: 0.5em;

				button {
					border: none;
					padding: 0.7em 1em;
					border-radius: 0.3rem;
					cursor: pointer;
					background-color: aliceblue;
					transition: all 300ms ease;

					span {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
						color: #040d12;
						transition: all 300ms ease;
					}

					&:hover {
						background-color: #040d12;

						span {
							color: white;
						}
					}

					&:first-of-type {
						background-color: #040d12;

						span {
							color: white;
						}
					}
				}
			}
		}
	}

	:global(.moveButtonDisabled) {
		opacity: 0.5;
		pointer-events: none;
	}
</style>
