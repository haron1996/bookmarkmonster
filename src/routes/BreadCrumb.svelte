<script lang="ts">
	import {
		apiHost,
		folderPath,
		folders,
		foldersToMove,
		selectedFolders,
		session,
		breadCrumbIsOverflowing,
		collapsedFolders,
		showCollapsedFolders
	} from '../stores/stores';
	import { removeSelectedClassFromAllDomFolders } from '../utils/removeSelectedClassFromAllDomFolders';
	import { removeSlideInAnimationFromActionBarV2 } from '../utils/removeSlideInAnimationFromActionBarV2';
	import type { Folder } from '../types/folder';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let root: string = 'My collections';

	async function moveCollectionsToRoot(f: Folder[], a: HTMLAnchorElement) {
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
				collections: f
			})
		});

		const result = await response.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			return;
		}

		const fs: Folder[] = result[0];

		if (fs === null) {
			selectedFolders.set([]);

			removeSelectedClassFromAllDomFolders();

			removeSlideInAnimationFromActionBarV2();

			a.style.border = 'none';
			return;
		}

		fs.forEach((f) => {
			$folders = $folders.filter((value) => {
				return value.folder_id !== f.folder_id;
			});

			folders.set($folders);
		});

		selectedFolders.set([]);

		removeSelectedClassFromAllDomFolders();

		removeSlideInAnimationFromActionBarV2();

		a.style.border = 'none';
	}

	function handleBreadcrumbRootDragOver(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.border = '0.2rem dashed #272829';
	}

	function handleBreadcrumbRootDragLeave(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.border = 'none';
	}

	function handleBreadcrumbRootDrop(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		if ($selectedFolders.length >= 1) {
			moveCollectionsToRoot($selectedFolders, a);
		} else {
			moveCollectionsToRoot($foldersToMove, a);
		}
	}

	function handleBreadcrumbFolderDragOver(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.border = '0.2rem dashed #272829';
	}

	function handleBreadcrumbFolderDragLeave(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.border = 'none';
	}

	function handleDropOnBreadcrumbFolder(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;

		const f: Folder = {
			created_at: a.dataset.createdat,
			deleted_at: a.dataset.deletedat,
			folder_description: a.dataset.folderdescription,
			folder_id: a.dataset.id,
			folder_name: a.dataset.foldername,
			label: a.dataset.label,
			path: a.dataset.path,
			starred: a.dataset.starred ? false : true,
			subfolder_of: a.dataset.subfolderof,
			updated_at: a.dataset.updatedat,
			user_id: a.dataset.userid
		};

		if ($selectedFolders.length >= 1) {
			$selectedFolders.forEach((sf) => {
				if (sf.folder_id === f.folder_id) {
					$selectedFolders = $selectedFolders.filter((value) => {
						return value.folder_id !== f.folder_id;
					});

					selectedFolders.set($selectedFolders);
				}
			});
			moveFolders(f, a, $selectedFolders);
		} else {
			if (f.folder_id === $foldersToMove[0].folder_id) {
				foldersToMove.set([]);
				return;
			}
			moveFolders(f, a, $foldersToMove);
		}
	}

	async function moveFolders(destination: Folder, a: HTMLAnchorElement, f: Folder[]) {
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
				destination_folder: destination,
				folders_to_move: f
			})
		});

		const result = await response.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			a.style.border = '0.1rem solid #e7ebed';
			return;
		}

		const fs: Folder[] = result[0];

		fs.forEach((f) => {
			$folders = $folders.filter((value) => {
				return value.folder_id !== f.folder_id;
			});

			folders.set($folders);
		});

		a.style.border = 'none';
	}

	function handleClickOnColapsedFolder(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const a = target.closest('a.collapsedFolder') as HTMLAnchorElement | null;

		if (a === null) return;

		goto(`/dashboard/my_collections?id=${a.dataset.id}`);
	}
</script>

<div class="breadCrumb" id="breadCrumb">
	{#if root === 'My collections'}
		<a
			data-sveltekit-preload-data="tap"
			href="/dashboard/my_collections?id=root"
			class="root"
			class:disabled={$folderPath.length < 1 && $collapsedFolders.length < 1}
			on:dragover|preventDefault={handleBreadcrumbRootDragOver}
			on:dragleave={handleBreadcrumbRootDragLeave}
			on:drop={handleBreadcrumbRootDrop}
		>
			{root}
		</a>
	{:else if root === 'Shared'}
		<a
			data-sveltekit-preload-data="tap"
			href="/dashboard/my_collections?id=root&which=shared"
			class="root"
			class:disabled={$folderPath.length < 1 && $collapsedFolders.length < 1}
			on:dragover|preventDefault={handleBreadcrumbRootDragOver}
		>
			{root}
		</a>
	{:else if (root = 'Recycle bin')}
		<a data-sveltekit-preload-data="tap" href="/dashboard/my_collections?id=root" class="disabled">
			{root}
		</a>
	{/if}

	{#if $collapsedFolders.length}
		<i
			class="las la-ellipsis-h"
			role="none"
			on:click|stopPropagation={() => {
				$showCollapsedFolders = !$showCollapsedFolders;
			}}
		>
			{#if $showCollapsedFolders}
				<div id="collapsedFolders" on:click|stopPropagation={() => {}} role="none">
					{#each $collapsedFolders as { created_at, deleted_at, folder_description, folder_id, folder_name, label, path, starred, subfolder_of, updated_at, user_id }}
						<a
							data-sveltekit-preload-data="tap"
							id="collapsedFolder"
							class="collapsedFolder"
							data-id={folder_id}
							data-label={label}
							data-createdat={created_at}
							data-deletedat={deleted_at}
							data-folderdescription={folder_description}
							data-path={path}
							data-starred={starred}
							data-subfolderof={subfolder_of}
							data-updatedat={updated_at}
							data-userid={user_id}
							data-foldername={folder_name}
							href={`/dashboard/my_collections?id=${folder_id}`}
							on:click|preventDefault={handleClickOnColapsedFolder}
						>
							{folder_name}
						</a>
					{/each}
				</div>
			{/if}
		</i>
	{/if}
	{#each $folderPath as { label, created_at, deleted_at, folder_description, folder_id, folder_name, path, starred, subfolder_of, updated_at, user_id }}
		<i class="las la-long-arrow-alt-right" />
		<a
			class="folderPath"
			data-sveltekit-preload-data="tap"
			href={`/dashboard/my_collections?id=${folder_id}`}
			data-id={folder_id}
			data-label={label}
			data-createdat={created_at}
			data-deletedat={deleted_at}
			data-folderdescription={folder_description}
			data-path={path}
			data-starred={starred}
			data-subfolderof={subfolder_of}
			data-updatedat={updated_at}
			data-userid={user_id}
			data-foldername={folder_name}
			on:dragover|preventDefault={handleBreadcrumbFolderDragOver}
			on:dragleave={handleBreadcrumbFolderDragLeave}
			on:drop={handleDropOnBreadcrumbFolder}
			>{folder_name}
		</a>
	{/each}
</div>

<style lang="scss">
	.breadCrumb {
		height: 5rem;
		width: 70vw;
		background-color: white;
		display: flex;
		align-items: center;
		padding: 0 1em;
		gap: 0.1em;

		i.la-ellipsis-h {
			font-size: 1.3rem;
			cursor: pointer;
			color: #272829;
			padding: 0.1em 0.6em;
			border-radius: 0.3rem;
			position: relative;
			color: red;

			div#collapsedFolders {
				position: absolute;
				top: 100%;
				left: 0;
				width: 20rem;
				height: min-content;
				max-height: 40rem;
				overflow-y: auto;
				overflow-x: hidden;
				z-index: 1;
				background-color: white;
				box-shadow: rgba(0, 0, 0, 0.15) 0px 15px 25px, rgba(0, 0, 0, 0.05) 0px 5px 10px;
				border-radius: 0.1rem;
				display: flex;
				flex-direction: column;

				a#collapsedFolder {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					padding: 0.7em 0.5em;
					border-radius: none;
					width: 100%;
					white-space: nowrap;
					overflow-x: hidden;
					text-overflow: ellipsis;
					transition: all 300ms ease;
					text-decoration: none;
					color: #61677a;
					font-weight: 500;

					&:hover {
						background-color: #eeeeee;
					}
				}
			}
		}

		i {
			font-size: 1.5rem;
		}

		a.folderPath,
		a.root {
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			text-decoration: none;
			transition: all ease 300ms;
			color: #61677a;
			padding: 0.3em 0.6em;
			border-radius: 1em 0.5rem;
			min-width: max-content;

			&:hover {
				background-color: #eeeeee;
			}
		}

		a:last-of-type:not(.root):not(#collapsedFolder) {
			pointer-events: none;
			color: #272829;
		}

		.disabled {
			pointer-events: none;
			text-decoration: none;
			font-family: 'Segoe UI', sans-serif;
			font-size: 1.5rem;
			color: #495464;
			padding: 0 !important;
		}
	}
</style>
