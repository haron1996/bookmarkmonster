<script lang="ts">
	import { onMount } from 'svelte';
	import logo from '../../static/logo.png';
	import Success from './Success.svelte';
	import { pageSaved } from '../../store/stores';

	let tagName = '';
	let collectionName = '';
	let url = '';
	let title = '';
	//const apiURL = 'http://localhost:5000';
	const apiURL = 'https://api.bookmarkmonster.xyz';
	let accessToken = '';

	onMount(() => {
		getCurrentTabUrlAndTitle();
		getAccessToken();
	});

	async function getAccessToken() {
		const t = await chrome.cookies.getAll({
			domain: 'localhost',
			secure: true,
			name: 'accessTokenCookie'
		});

		accessToken = t[0].value;
	}

	interface Tag {
		added?: string;
		deleted?: string | null;
		id?: string;
		name?: string;
		updated?: string | null;
		user_id?: string;
	}

	interface Folder {
		created_at?: string | null;
		deleted_at?: string | null;
		folder_description?: string;
		folder_id?: string;
		folder_name?: string;
		label?: string;
		path?: string;
		starred?: boolean;
		subfolder_of?: string | null;
		updated_at?: string;
		user_id?: string;
	}

	let chosenTags: Tag[] = [];
	let tagsMatchingQuery: Tag[] = [];
	let chosenFolders: Folder[] = [];
	let foldersMatchingQuery: Folder[] = [];

	async function getCurrentTabUrlAndTitle() {
		chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
			const currentTab = tabs[0];
			if (currentTab.url) {
				url = currentTab.url;
			}

			if (currentTab.title) {
				title = currentTab.title;
			}
		});
	}

	async function searchTagsMatchingQuery() {
		if (tagName === '') {
			tagsMatchingQuery = [];
			return;
		}

		const response = await fetch(`${apiURL}/authenticated/tags/${tagName}`, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${accessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		tagsMatchingQuery = result[0];
	}

	function syncChromeBookmarks() {
		chrome.bookmarks.getTree((tree) => {
			const nodes = tree[0].children;

			if (nodes) {
				for (const node of nodes) {
					const bookmarks = node.children;
					if (bookmarks) {
						for (const bookmark of bookmarks) {
							alert(bookmark);
						}
					}
				}
			}
		});
	}

	function handleTagInputKeydown(e: KeyboardEvent) {
		if (e.code === 'Space' && tagName === '') {
			e.preventDefault();
			return;
		}

		if (e.code === 'Enter') {
			e.preventDefault();

			if (tagName === '') return;

			const exists = tagsMatchingQuery.some((obj) => obj.name === tagName);

			if (exists) {
				const tag = tagsMatchingQuery.filter((obj) => {
					return obj.name === tagName;
				});

				const exists = chosenTags.some((obj) => obj.id === tag[0].id && obj.name === tag[0].name);

				if (!exists) {
					addTagToChosenTags(tag[0]);
				}

				tagName = '';

				tagsMatchingQuery = [];
			} else {
				createTag(tagName);
				tagName = '';
			}

			return;
		}
	}

	function chooseTag(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const span = target.closest('span') as HTMLDivElement | null;

		if (span) {
			const tag: Tag = {
				added: span.dataset.added,
				deleted: span.dataset.deleted,
				id: span.dataset.id,
				name: span.dataset.name,
				updated: span.dataset.updated,
				user_id: span.dataset.userid
			};

			const exists = chosenTags.some((obj) => obj.id === tag.id && obj.name === tag.name);

			if (!exists) {
				addTagToChosenTags(tag);
			}

			tagName = '';

			tagsMatchingQuery = [];
		}
	}

	function addTagToChosenTags(tag: Tag) {
		chosenTags = [tag, ...chosenTags];
	}

	function removeTagFromChosenTags(tag: Tag) {
		chosenTags = chosenTags.filter((obj) => {
			return obj.id !== tag.id;
		});

		chosenTags = chosenTags;
	}

	async function createTag(tagName: string) {
		const response = await fetch(`${apiURL}/authenticated/tags/create-tag`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${accessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ name: tagName })
		});

		const result = await response.json();

		if (result.message) {
			alert(result.message);
			tagName = '';
			tagsMatchingQuery = [];
			return;
		}

		chosenTags = [result[0], ...chosenTags];
	}

	function RemoveTagFromChosenTags(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const div = target.closest('div.tag') as HTMLDivElement | null;

		if (div) {
			const tag: Tag = {
				added: div.dataset.added,
				deleted: div.dataset.deleted,
				id: div.dataset.id,
				name: div.dataset.name,
				updated: div.dataset.updated,
				user_id: div.dataset.userid
			};

			const exists = chosenTags.some((obj) => obj.id === tag.id && obj.name === tag.name);

			if (exists) {
				removeTagFromChosenTags(tag);
			}
		}
	}

	async function searchUserFolders() {
		if (collectionName === '') {
			foldersMatchingQuery = [];
			return;
		}

		const response = await fetch(
			`${apiURL}/authenticated/collections/searchUserFolders/${collectionName}`,
			{
				method: 'GET',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${accessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer'
			}
		);

		const result = await response.json();

		if (result.message) {
			alert(result.message);
			collectionName = '';
			foldersMatchingQuery = [];
			return;
		}

		foldersMatchingQuery = result[0];
	}

	function chooseFolder(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const span = target.closest('span') as HTMLDivElement | null;

		if (span) {
			const folder: Folder = {
				created_at: span.dataset.created,
				deleted_at: span.dataset.deleted,
				folder_description: span.dataset.description,
				folder_id: span.dataset.id,
				folder_name: span.dataset.name,
				label: span.dataset.label,
				path: span.dataset.span,
				starred: span.dataset.starred ? true : false,
				subfolder_of: span.dataset.parent,
				updated_at: span.dataset.updated,
				user_id: span.dataset.userid
			};

			const exists = chosenFolders.some(
				(obj) => obj.folder_id === folder.folder_id && obj.folder_name === folder.folder_name
			);

			if (!exists) {
				addFolderToChosenFolders(folder);
			}

			collectionName = '';

			foldersMatchingQuery = [];
		}
	}

	function addFolderToChosenFolders(folder: Folder) {
		chosenFolders = [];
		chosenFolders = [folder, ...chosenFolders];
	}

	function removeFolderFromChosenFolders(folder: Folder) {
		chosenFolders = chosenFolders.filter((obj) => {
			return obj.folder_id !== folder.folder_id;
		});

		chosenFolders = chosenFolders;
	}

	function handleFolderInputKeydown(e: KeyboardEvent) {
		if (e.code === 'Space' && collectionName === '') {
			e.preventDefault();
			return;
		}

		if (e.code === 'Enter') {
			e.preventDefault();

			const exists = foldersMatchingQuery.some((obj) => obj.folder_name === collectionName);

			if (exists) {
				const folder = foldersMatchingQuery.filter((obj) => {
					return obj.folder_name === collectionName;
				});

				const exists = chosenFolders.some(
					(obj) =>
						obj.folder_id === folder[0].folder_id && obj.folder_name === folder[0].folder_name
				);

				if (!exists) {
					addFolderToChosenFolders(folder[0]);
				}

				collectionName = '';

				foldersMatchingQuery = [];
			} else {
				createFolder(collectionName);
			}
		}
	}

	async function createFolder(folderName: string) {
		const response = await fetch(`${apiURL}/authenticated/collections/createCollection`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${accessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				name: folderName,
				parent_folder_id: '',
				description: ''
			})
		});

		const result = await response.json();

		if (result.message) {
			alert(result.message);
			collectionName = '';
			foldersMatchingQuery = [];
			return;
		}

		addFolderToChosenFolders(result[0]);

		collectionName = '';

		foldersMatchingQuery = [];
	}

	function RemoveFolderFromChosenFolders(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const div = target.closest('div.collection') as HTMLDivElement | null;

		if (div) {
			const folder: Folder = {
				created_at: div.dataset.created,
				deleted_at: div.dataset.deleted,
				folder_description: div.dataset.description,
				folder_id: div.dataset.id,
				folder_name: div.dataset.name,
				label: div.dataset.label,
				path: div.dataset.span,
				starred: div.dataset.starred ? true : false,
				subfolder_of: div.dataset.parent,
				updated_at: div.dataset.updated,
				user_id: div.dataset.userid
			};

			const exists = chosenFolders.some(
				(obj) => obj.folder_id === folder.folder_id && obj.folder_name === folder.folder_name
			);

			if (exists) {
				removeFolderFromChosenFolders(folder);
			}
		}
	}

	async function savePage() {
		const response = await fetch(`${apiURL}/authenticated/bookmarks/savePageFromChromeExtension`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${accessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				url: url,
				title: title,
				tags: chosenTags,
				folders: chosenFolders
			})
		});

		const result = await response.json();

		//alert(result[0].title);

		pageSaved.set(true);
	}
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://maxst.icons8.com/vue-static/landings/line-awesome/line-awesome/1.3.0/css/line-awesome.min.css"
	/>
</svelte:head>

<div class="popup">
	<div class="top">
		<img src={logo} alt="logo" />
		<div class="details">
			<p>BookmarkMonster</p>
			<span>Cloud Bookmark Manager</span>
		</div>
	</div>
	<form
		on:submit|preventDefault={() => {
			alert('submit form');
		}}
	>
		<input type="url" name="url" id="url" bind:value={url} disabled />
		<input type="text" name="title" id="title" bind:value={title} disabled />
		<div class="tags">
			{#if chosenTags.length >= 1}
				<div class="chosenTags">
					{#each chosenTags as { added, deleted, id, name, updated, user_id }}
						<div
							role="none"
							class="tag"
							data-added={added}
							data-deleted={deleted}
							data-id={id}
							data-name={name}
							data-updated={updated}
							data-userid={user_id}
						>
							<i class="las la-hashtag" />
							<span>{name}</span>
							<i class="las la-times" role="none" on:click={RemoveTagFromChosenTags} />
						</div>
					{/each}
				</div>
			{/if}

			<input
				type="text"
				name=""
				id="tags"
				placeholder="Choose tags"
				autocomplete="off"
				bind:value={tagName}
				on:input={searchTagsMatchingQuery}
				on:keydown={handleTagInputKeydown}
				on:pointerdown={() => {
					foldersMatchingQuery = [];
					if (tagName !== '') {
						searchTagsMatchingQuery();
					}
				}}
			/>
			{#if tagsMatchingQuery.length >= 1}
				<div class="tagsMatchingQuery">
					{#each tagsMatchingQuery as { added, deleted, id, name, updated, user_id }}
						<span
							role="none"
							data-name={name}
							data-added={added}
							data-deleted={deleted}
							data-updated={updated}
							data-userid={user_id}
							data-id={id}
							on:click|stopPropagation={chooseTag}
						>
							{name}
						</span>
					{/each}
				</div>
			{/if}
		</div>
		<div class="collections">
			{#if chosenFolders.length >= 1}
				<div class="chosenCollections">
					{#each chosenFolders as { created_at, deleted_at, folder_description, folder_id, folder_name, label, path, starred, subfolder_of, updated_at, user_id }}
						<div
							class="collection"
							data-created={created_at}
							data-deleted={deleted_at}
							data-description={folder_description}
							data-id={folder_id}
							data-name={folder_name}
							data-label={label}
							data-path={path}
							data-starred={starred}
							data-parent={subfolder_of}
							data-updated={updated_at}
							data-userid={user_id}
						>
							<i class="las la-folder-open" />
							<span>{folder_name}</span>
							<i class="las la-times" role="none" on:click={RemoveFolderFromChosenFolders} />
						</div>
					{/each}
				</div>
			{/if}
			<input
				type="text"
				name=""
				id="collection"
				placeholder="Choose collection"
				autocomplete="off"
				bind:value={collectionName}
				on:input={searchUserFolders}
				on:pointerdown={() => {
					tagsMatchingQuery = [];
					if (collectionName !== '') {
						searchUserFolders();
					}
				}}
				on:keydown={handleFolderInputKeydown}
			/>
			{#if foldersMatchingQuery.length >= 1}
				<div class="foldersMatchingQuery">
					{#each foldersMatchingQuery as { created_at, deleted_at, folder_description, folder_id, folder_name, label, path, starred, subfolder_of, updated_at, user_id }}
						<span
							data-created={created_at}
							data-deleted={deleted_at}
							data-description={folder_description}
							data-id={folder_id}
							data-name={folder_name}
							data-label={label}
							data-path={path}
							data-starred={starred}
							data-parent={subfolder_of}
							data-updated={updated_at}
							data-userid={user_id}
							role="none"
							on:click={chooseFolder}>{folder_name}</span
						>
					{/each}
				</div>
			{/if}
		</div>
		<div class="buttons">
			<button on:click|preventDefault={savePage}>Bookmark this page</button>
			<div class="importChromeBookmarks">
				<span>Sync chrome bookmarks</span>
			</div>
		</div>
	</form>
	<small>Version: 0.0.1</small>

	<!-- allerts -->
	<Success />
	<!-- end of allerts -->
</div>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Bricolage+Grotesque:opsz,wght@12..96,200;12..96,300;12..96,400;12..96,500;12..96,600;12..96,700;12..96,800&display=swap');

	.popup {
		min-width: 500px;
		height: max-content;
		max-height: 600px;
		overflow-y: auto;
		display: flex;
		flex-direction: column;
		gap: 3em;
		background-color: #f0f0f0;

		& .top {
			min-height: 10vh;
			display: flex;
			gap: 2em;
			align-items: center;
			padding: 1em;
			border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
			background-color: #f1efef;

			& img {
				width: 3rem;
				height: 3rem;
				object-fit: contain;
			}

			& .details {
				display: flex;
				flex-direction: column;
				justify-content: center;
				gap: 0.3em;

				& p {
					font-size: 1.5rem;
					font-family: 'Bricolage Grotesque', sans-serif;
					font-weight: 700;
				}

				& span {
					font-family: 'Bricolage Grotesque', sans-serif;
					font-size: 1.2rem;
					color: #9ba4b5;
				}
			}
		}

		& form {
			display: flex;
			flex-direction: column;
			height: max-content;
			padding: 1em;
			width: 100%;
			gap: 3em;

			& input {
				width: inherit;
				padding: 0.5em;
				height: 4rem;
				font-family: 'Bricolage Grotesque', sans-serif;
				border: 0.1rem solid rgb(0, 0, 0, 0.1);
				outline: none;
				border-radius: 0.3rem;
			}

			& .tags {
				width: 100%;
				display: flex;
				flex-direction: column;
				gap: 1em;
				position: relative;

				& .chosenTags {
					display: flex;
					flex-flow: row wrap;
					align-content: start;
					gap: 0.5em;
					height: max-content;
					max-height: 15rem;
					overflow-y: auto;

					& .tag {
						background-color: #eeeeee;
						padding: 0.7em;
						border-radius: 0.3rem;
						display: flex;
						align-items: center;
						gap: 0.5em;

						& span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							text-transform: lowercase;
						}

						& i {
							font-size: 1.3rem;
						}

						& i.la-times {
							cursor: pointer;
						}
					}
				}

				& input {
					width: inherit;
					padding: 0.7em;
					font-family: 'Bricolage Grotesque', sans-serif;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					outline: none;
					border-radius: 0.3rem;
				}

				& .tagsMatchingQuery {
					position: absolute;
					top: 100%;
					left: 0;
					right: 0;
					width: inherit;
					height: max-content;
					max-height: 15rem;
					overflow-y: auto;
					padding: 0.5em;
					box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em,
						rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em, rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
					display: flex;
					flex-flow: row wrap;
					align-content: start;
					gap: 0.5em;
					background-color: white;
					z-index: 1;

					& span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						background-color: #eeeeee;
						padding: 0.7em;
						border-radius: 0.3rem;
						cursor: pointer;
					}
				}
			}

			& .collections {
				width: 100%;
				display: flex;
				flex-direction: column;
				gap: 1em;
				position: relative;

				& .chosenCollections {
					display: flex;
					flex-flow: row wrap;
					align-content: start;
					gap: 0.5em;
					height: max-content;
					max-height: 15rem;
					overflow-y: auto;

					& .collection {
						background-color: #eeeeee;
						padding: 0.7em;
						border-radius: 0.3rem;
						display: flex;
						align-items: center;
						gap: 0.5em;

						& span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							text-transform: lowercase;
						}

						& i {
							font-size: 1.3rem;
						}

						& i.la-times {
							cursor: pointer;
						}
					}
				}

				& input {
					width: inherit;
					padding: 0.7em;
					font-family: 'Bricolage Grotesque', sans-serif;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					outline: none;
					border-radius: 0.3rem;
				}

				& .foldersMatchingQuery {
					position: absolute;
					top: 100%;
					left: 0;
					right: 0;
					width: inherit;
					height: max-content;
					max-height: 15rem;
					overflow-y: auto;
					padding: 0.5em;
					box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em,
						rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em, rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
					display: flex;
					flex-flow: row wrap;
					align-content: start;
					gap: 0.5em;
					background-color: white;
					z-index: 1;

					& span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						background-color: #eeeeee;
						padding: 0.7em;
						border-radius: 0.3rem;
						cursor: pointer;
					}
				}
			}

			& .buttons {
				display: flex;
				align-items: center;
				gap: 1em;

				& button {
					flex: 1;
					height: 4rem;
					font-family: 'Bricolage Grotesque', sans-serif;
					font-size: 1.3rem;
					border: none;
					color: white;
					cursor: pointer;
					border-radius: 0.3rem;
					display: flex;
					align-items: center;
					justify-content: center;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					transition: all 200ms ease;
					background-color: #272829;

					&:hover {
						box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px, #1a9047 0px 0px 0px 2px;
					}
				}

				& .importChromeBookmarks {
					flex: 1;
					height: 4rem;
					display: flex;
					align-items: center;
					justify-content: center;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					cursor: pointer;
					border-radius: 0.3rem;
					transition: all 200ms ease;

					& span {
						font-size: 1.3rem;
						font-family: 'Bricolage Grotesque', sans-serif;
					}

					&:hover {
						background-color: #272829;
						box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px, #1a9047 0px 0px 0px 2px;

						& span {
							color: white;
						}
					}
				}
			}
		}

		& small {
			font-size: 1.2rem;
			font-family: 'Arial CE', sans-serif;
			padding: 0.5em;
			color: #b4b4b3;
		}
	}
</style>
