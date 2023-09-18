<script lang="ts">
	import {
		childrenOfDestinationFolder,
		destinationFolder,
		folders,
		selectedFolders,
		showMoveItemsPopup
	} from '../stores/stores';
	import type { Folder } from '../types/folder';

	let folder: Folder = {};

	function handleClickOnFolder(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const f = t.closest('.folder') as HTMLAnchorElement | null;

		if (f === null) return;

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

		destinationFolder.set(folder);

		doShit();
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

{#each $childrenOfDestinationFolder as { label, created_at, deleted_at, folder_description, folder_id, folder_name, path, starred, subfolder_of, updated_at, user_id }}
	<a
		href="/"
		data-sveltekit-preload-data="tap"
		class="folder folderV2"
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
		on:click|preventDefault={handleClickOnFolder}
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
		width: 40rem;
		height: 25rem;
		text-decoration: none;
		transition: all ease 300ms;
		position: relative;
		border: 0.1rem solid #e7ebed;

		svg {
			height: 13rem;
			width: 13rem;

			path {
				fill: #040d12;
				stroke: #040d12;
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

		&:hover {
			background-color: #eeeeee;
		}
	}

	:global(.blocked) {
		opacity: 0.5;
		pointer-events: none;
	}
</style>
