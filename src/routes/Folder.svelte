<script lang="ts">
	import { folders, selectedFolders } from '../stores/stores';
	import type { Folder } from '../types/folder';

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
			starred: f.dataset.starred,
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
</script>

{#each $folders as { label, created_at, deleted_at, folder_description, folder_id, folder_name, path, starred, subfolder_of, updated_at, user_id }}
	<a
		data-sveltekit-preload-data="tap"
		href={`/dashboard/my_collections?id=${folder_id}`}
		class="folder"
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
				fill: #0079ff;
				stroke: #0079ff;
			}
		}

		span {
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			white-space: nowrap;
			overflow: hidden;
			text-overflow: ellipsis;
			max-width: 95%;
			height: max-content;
			padding: 0.5em;
			border-radius: 0.3rem;
			color: #080202;
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
		}

		&:hover {
			box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px, rgba(17, 17, 26, 0.1) 0px 0px 8px;

			.selector {
				opacity: 1;
			}
		}
	}

	:global(.selected) {
		box-shadow: #3740ff 0px 0px 0px 3px;

		.selector {
			background-color: #0079ff !important;
			box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
			opacity: 1 !important;
			border-color: white !important;

			i {
				opacity: 1 !important;
				color: white !important;
			}
		}

		&:hover {
			box-shadow: #3740ff 0px 0px 0px 3px !important;
		}
	}
</style>
