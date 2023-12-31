<script lang="ts">
	import {
		apiHost,
		folders,
		foldersToMove,
		selectedFolders,
		session,
		showMoveItemsPopup
	} from '../stores/stores';
	import type { Folder } from '../types/folder';
	import interact from 'interactjs';
	import { removeSelectedClassFromAllDomFolders } from '../utils/removeSelectedClassFromAllDomFolders';

	let folder: Folder = {};

	function selectFolder(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const f = t.closest('.folder') as HTMLAnchorElement | null;

		if (f === null) return;

		const div = document.getElementById('actionBar') as HTMLDivElement | null;

		if (div === null) return;

		folder = {
			created_at: f.dataset.createdat,
			deleted_at: f.dataset.deletedat,
			folder_description: f.dataset.folderdescription,
			folder_id: f.dataset.id,
			folder_name: f.dataset.foldername,
			label: f.dataset.label,
			path: f.dataset.path,
			starred: f.dataset.starred ? true : false,
			subfolder_of: f.dataset.subfolderof,
			updated_at: f.dataset.updatedat,
			user_id: f.dataset.userid
		};

		const exists = $selectedFolders
			.map((value) => {
				return value.folder_id;
			})
			.includes(folder.folder_id);

		if (exists) {
			$selectedFolders = $selectedFolders.filter((value) => {
				return value.folder_id !== folder.folder_id;
			});

			selectedFolders.set($selectedFolders);

			f.classList.remove('selected');

			div.classList.remove('animate__animated');

			div.classList.remove('animate__fadeInDown');
		} else {
			selectedFolders.update((values) => [folder, ...values]);

			f.classList.add('selected');

			div.classList.add('animate__animated');

			div.classList.add('animate__fadeInDown');
		}
	}

	function handleFolderDragStart(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.opacity = '0';

		let f: Folder = {
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

		foldersToMove.set([]);

		foldersToMove.update((values) => [f, ...values]);
	}

	function handleFolderDragEnd(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.opacity = '1';
		a.style.border = '0.1rem solid #e7ebed';
		const domFolders = document.querySelectorAll('a.f') as NodeListOf<HTMLAnchorElement> | null;

		if (domFolders === null) return;

		domFolders.forEach((f) => {
			f.style.opacity = '1';
		});

		selectedFolders.set([]);
		removeSelectedClassFromAllDomFolders();
	}

	function handleFolderDragOver(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.border = 'thick dashed red';
	}

	function handleFolderDragLeave(e: DragEvent) {
		const target = e.currentTarget as HTMLElement;
		const a = target.closest('a') as HTMLAnchorElement | null;
		if (a === null) return;
		a.style.border = '0.1rem solid #e7ebed';
	}

	function handleFolderDrop(e: DragEvent) {
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
				// alert('cannot move collection to itself');
				foldersToMove.set([]);
				return;
			}

			moveFolders(f, a, $foldersToMove);
		}
	}

	function handleFolderDragging(e: DragEvent) {
		removeSelectedClassFromAllDomFolders();
		const domFolders = document.querySelectorAll('a.f') as NodeListOf<HTMLAnchorElement> | null;

		if (domFolders === null) return;

		domFolders.forEach((folder) => {
			$selectedFolders.forEach((f) => {
				if (folder.dataset.id === f.folder_id) {
					folder.style.opacity = '.4';
				}
			});
		});
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

		a.style.border = '0.1rem solid #e7ebed';
	}
</script>

{#each $folders as { label, created_at, deleted_at, folder_description, folder_id, folder_name, path, starred, subfolder_of, updated_at, user_id }}
	<a
		data-sveltekit-preload-data="tap"
		href={`/dashboard/my_collections?id=${folder_id}`}
		class="folder f"
		id="folder"
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
		draggable="true"
		on:dragstart={handleFolderDragStart}
		on:dragend={handleFolderDragEnd}
		on:dragover|preventDefault={handleFolderDragOver}
		on:dragleave={handleFolderDragLeave}
		on:drop={handleFolderDrop}
		on:drag={handleFolderDragging}
	>
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			role="none"
			><path
				d="M2 11V4.6a.6.6 0 01.6-.6h6.178a.6.6 0 01.39.144l3.164 2.712a.6.6 0 00.39.144H21.4a.6.6 0 01.6.6V11M2 11v8.4a.6.6 0 00.6.6h18.8a.6.6 0 00.6-.6V11M2 11h20"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
		</svg>
		<span>{folder_name}</span>
		<div
			class="selector"
			id="folderSelector"
			on:click|preventDefault|stopPropagation={selectFolder}
			role="none"
		>
			<i class="las la-check" />
		</div>
	</a>
{/each}

<style lang="scss">
	a.folder {
		background-color: inherit;
		display: flex;
		flex-direction: column;
		width: max-content;
		align-items: center;
		justify-content: center;
		padding: 1em;
		border-radius: 0.3rem;
		width: 30rem;
		height: 25rem;
		text-decoration: none;
		transition: all ease 300ms;
		position: relative;
		border: 0.1rem solid #e7ebed;

		svg {
			height: 13rem;
			width: 13rem;

			path {
				fill: #454545;
				stroke: #454545;
			}
		}

		span {
			font-size: 1.3rem;
			white-space: nowrap;
			overflow: hidden;
			text-overflow: ellipsis;
			max-width: 95%;
			height: max-content;
			padding: 0.5em;
			border-radius: 0.3rem;
			color: #040d12;
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;

			&::first-letter {
				text-transform: uppercase;
			}
		}

		.selector {
			height: 1.8rem;
			width: 1.8rem;
			position: absolute;
			top: 0.5rem;
			right: 0.5rem;
			background-color: white;
			display: flex;
			align-items: center;
			justify-content: center;
			border: 0.1rem solid #9ba4b5;
			transition: all 300ms ease;
			opacity: 0;
			border-radius: 100vh;
			cursor: default;
			cursor: pointer;

			i {
				font-size: 1.3rem;
				opacity: 0;
			}

			&:hover {
				i {
					opacity: 1;
				}
			}

			@media only screen and (width <= 768px) {
				opacity: 1;
			}
		}

		&:hover {
			box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px, rgba(17, 17, 26, 0.1) 0px 0px 8px;

			.selector {
				opacity: 1;
			}
		}

		@media only screen and (width <= 425px) {
			width: 100%;
		}

		@media only screen and (426px <= width <= 768px) {
			width: 30rem;
		}
	}

	:global(.selected) {
		box-shadow: rgb(4, 13, 18) 0px 0px 0px 3px;

		.selector {
			background-color: rgb(4, 13, 18) !important;
			box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
			opacity: 1 !important;
			border-color: white !important;

			i {
				opacity: 1 !important;
				color: white !important;
			}
		}

		&:hover {
			box-shadow: rgb(4, 13, 18) 0px 0px 0px 3px !important;
		}
	}
</style>
