<script lang="ts">
	import {
		apiHost,
		bookmarks,
		currentTagID,
		lastAddedBookmark,
		processingBookmark,
		session,
		sideBarWidth,
		tags
	} from '../../stores/stores';
	import { afterNavigate, goto } from '$app/navigation';
	import Overlay from '../Overlay.svelte';
	import AddBookmark from '../AddBookmark.svelte';
	import CreateTag from '../CreateTag.svelte';
	import { showCreateTagComponent } from '../../utils/showCreateTagComponent';
	import TagCreated from '../alerts/TagCreated.svelte';
	import DuplicateTagAlert from '../alerts/DuplicateTagAlert.svelte';
	import { showCreateBookmarkComponent } from '../../utils/showAddBookmarkComponent';
	import GlobePNG from '$lib/images/globe.png';
	import { page } from '$app/stores';
	import { browser } from '$app/environment';
	import TagAddedBookmark from '../TagAddedBookmark.svelte';
	import { showTagCreatedBookmarkForm } from '../../utils/showTagCreatedBookmarkForm';

	let sideBarWidthFromStore: number;
	let sidebarVisible: boolean = false;

	afterNavigate(async () => {
		loadUserSession();

		await getUserTags();

		//await getUserBookmarks();

		checkUrlTagAndFetchAppropriateBookmarks();
	});

	const loadUserSession = () => {
		const sessionString = localStorage.getItem('session') as string | null;

		if (sessionString === null) {
			goto(`${$page.url.origin}`);
			console.log('session is empty');
			return;
		}

		session.set(JSON.parse(sessionString));
	};

	function checkUrlTagAndFetchAppropriateBookmarks() {
		const currentTagNameinURL: string | null = $page.url.searchParams.get('tag');
		if (currentTagNameinURL === null) {
			// not tag search param found
			const newURL: URL = new URL($page.url.href);

			newURL.searchParams.set('tag', 'all-tags');

			window.history.pushState({}, '', newURL);

			currentTagID.set('all-tags');

			return;
		}
		const currentUserTags = document.querySelectorAll('.tag') as NodeListOf<HTMLDivElement> | null;
		if (currentUserTags === null) return;
		currentUserTags.forEach((t) => {
			t.classList.remove('active-tag');
			if (t.dataset.name) {
				if (t.dataset.name === currentTagNameinURL) {
					if (t.dataset.id === undefined) return;
					t.classList.add('active-tag');
					currentTagID.set(t.dataset.id);
				}
			}
		});
		console.log();
	}

	const getUserTags = async () => {
		if (browser) {
			const sessionString = localStorage.getItem('session') as string | null;

			if (sessionString === null) {
				goto(`${$page.url.origin}`);
				console.log('session is empty');
				return;
			}

			session.set(JSON.parse(sessionString));

			const response = await fetch(`${$apiHost}/authenticated/tags`, {
				method: 'GET',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${$session.AccessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer'
			});

			const result = await response.json();

			tags.set(result[0]);
		}
	};

	const getUserBookmarks = async () => {
		if (browser) {
			const sessionString = localStorage.getItem('session') as string | null;

			if (sessionString === null) {
				goto(`${$page.url.origin}`);
				console.log('session is empty');
				return;
			}

			session.set(JSON.parse(sessionString));

			const response = await fetch(`${$apiHost}/authenticated/bookmarks`, {
				method: 'GET',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${$session.AccessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer'
			});

			const result = await response.json();

			bookmarks.set(result[0]);
		}
	};

	const handleClickOnTag = (e: MouseEvent) => {
		const clickedEl = e.target as HTMLElement;

		const clickedTag = clickedEl.closest('.tag') as HTMLSpanElement | null;

		if (clickedTag === null) return;

		const tags = document.querySelectorAll('.tag') as NodeListOf<HTMLSpanElement>;

		tags.forEach((tag) => {
			if (tag.classList.contains('active-tag')) {
				tag.classList.remove('active-tag');
			}
		});

		clickedTag.classList.add('active-tag');

		if (clickedTag.closest('.all-tags') === null) {
			if (clickedTag.dataset.id) {
				currentTagID.set(clickedTag.dataset.id);
			}

			if (clickedTag.dataset.name === undefined) return;

			$page.url.searchParams.delete('tag');

			const newURL: URL = new URL($page.url.href);

			newURL.searchParams.set('tag', clickedTag.dataset.name);

			window.history.pushState({}, '', newURL);

			//window.history.pushState({}, '', `${$page.url.href}/?tag=${clickedTag.dataset.name}`);
		} else if (clickedTag.closest('.all-tags')) {
			if (clickedTag.dataset.id) {
				currentTagID.set(clickedTag.dataset.id);
			}

			if (clickedTag.dataset.name === undefined) return;

			$page.url.searchParams.delete('tag');

			const newURL: URL = new URL($page.url.href);

			newURL.searchParams.set('tag', clickedTag.dataset.name);

			window.history.replaceState({}, '', newURL);
		}
	};

	function openBookmark(e: MouseEvent) {
		const el = e.currentTarget as HTMLElement;

		const b = el.closest('.bookmark') as HTMLDivElement | null;

		if (b === null) return;

		const bookmarkLink: string | undefined = b.dataset.bookmark;

		if (bookmarkLink === undefined) return;

		window.open(bookmarkLink, '_blank');
	}

	async function getUserBookmarksByTagID() {
		if (browser) {
			if ($session.AccessToken) {
				const response = await fetch(`${$apiHost}/authenticated/bookmarks/${$currentTagID}`, {
					method: 'GET',
					mode: 'cors',
					cache: 'no-cache',
					credentials: 'include',
					headers: {
						'Content-Type': 'application/json',
						authorization: `Bearer${$session.AccessToken}`
					},
					redirect: 'follow',
					referrerPolicy: 'no-referrer'
				});

				if (response.ok) {
					const result = await response.json();

					if (result[0] === null) {
						bookmarks.set([]);
						return;
					}

					bookmarks.set(result[0]);
				} else {
					console.log(response.status, response.statusText);
				}
			}
		}
	}

	function handleWindPopstate() {
		window.history.forward();
		return;
	}

	function toggleSideBar() {
		const sideBar = document.getElementById('sideBar') as HTMLDivElement | null;

		if (sideBar === null) return;

		sidebarVisible = !sidebarVisible;
	}

	function toggleUserMenu() {
		const userMenu = document.getElementById('userMenu') as HTMLDivElement | null;

		if (userMenu === null) return;

		userMenu.classList.toggle('toggleUserMenu');
	}

	function logUserOut() {
		const session = localStorage.getItem('session') as string | null;

		if (session === null) return;

		localStorage.removeItem('session');

		window.location.reload();
	}

	$: sideBarWidthFromStore = $sideBarWidth;

	$: $currentTagID === 'all-tags' ? getUserBookmarks() : getUserBookmarksByTagID();

	$: $lastAddedBookmark,
		$lastAddedBookmark && $lastAddedBookmark.id !== undefined
			? showTagCreatedBookmarkForm()
			: () => {};
</script>

<svelte:head>
	<title>BookmarkMonster | Dashboard</title>
</svelte:head>

<svelte:window on:popstate|preventDefault={handleWindPopstate} />

<div class="app">
	<div class="sidebar" id="sideBar" class:toggleSidebar={sidebarVisible}>
		<div class="profile" role="none" on:click={toggleUserMenu}>
			{#if $session && $session.User && $session.User.picture && $session.User.name && $session.User.email}
				<img src={$session.User.picture} alt="profile" />
				<div class="name-and-email">
					<h3>{$session.User.name}</h3>
					<span>{$session.User.email}</span>
				</div>
			{/if}
			<div class="userMenu" id="userMenu">
				<div role="none" on:click|stopPropagation={logUserOut}>
					<svg
						width="24px"
						height="24px"
						stroke-width="1.5"
						viewBox="0 0 24 24"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						color="#000000"
						><path
							d="M12 12h7m0 0l-3 3m3-3l-3-3M19 6V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2h10a2 2 0 002-2v-1"
							stroke="#000000"
							stroke-width="1.5"
							stroke-linecap="round"
							stroke-linejoin="round"
						/>
					</svg>
					<span>Log Out</span>
				</div>
				<div style="display: none;">
					<svg
						width="24px"
						height="24px"
						stroke-width="1.5"
						viewBox="0 0 24 24"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						color="#000000"
						><path
							d="M12 15a3 3 0 100-6 3 3 0 000 6z"
							stroke="#000000"
							stroke-width="1.5"
							stroke-linecap="round"
							stroke-linejoin="round"
						/><path
							d="M19.622 10.395l-1.097-2.65L20 6l-2-2-1.735 1.483-2.707-1.113L12.935 2h-1.954l-.632 2.401-2.645 1.115L6 4 4 6l1.453 1.789-1.08 2.657L2 11v2l2.401.655L5.516 16.3 4 18l2 2 1.791-1.46 2.606 1.072L11 22h2l.604-2.387 2.651-1.098C16.697 18.831 18 20 18 20l2-2-1.484-1.75 1.098-2.652 2.386-.62V11l-2.378-.605z"
							stroke="#000000"
							stroke-width="1.5"
							stroke-linecap="round"
							stroke-linejoin="round"
						/>
					</svg>
					<span>Settings</span>
				</div>
			</div>
		</div>

		<div class="tags" id="userTags">
			{#if $tags.length > 0}
				<div
					class="tag active-tag all-tags"
					id="allTagsDiv"
					on:click|stopPropagation={handleClickOnTag}
					data-id="all-tags"
					data-name="all-tags"
					role="none"
				>
					<i class="las la-hashtag" />
					<span>all tags</span>
				</div>
				{#each $tags as { id, name, user_id, added, updated, deleted }}
					<div
						class="tag"
						id="tag"
						on:click|stopPropagation={handleClickOnTag}
						on:keyup
						role="none"
						data-id={id}
						data-userid={user_id}
						data-name={name}
					>
						<i class="las la-hashtag" />
						<span>{name}</span>
					</div>
				{/each}
			{/if}
		</div>
		<div class="create-tag" on:click|stopPropagation={showCreateTagComponent} role="none">
			<div class="new-tag">
				<i class="las la-plus" />
				<span>create tag</span>
			</div>
		</div>
	</div>
	<div
		class="main"
		id="main"
		on:click={() => {
			if (sidebarVisible) {
				sidebarVisible = false;
			}
		}}
		role="none"
	>
		<div class="search-or-add">
			<i class="las la-bars" role="none" on:click|stopPropagation={toggleSideBar} />
			<input
				type="search"
				name="search"
				id="search"
				placeholder="Type to start searching (coming soon)"
				autocomplete="off"
			/>
			<button
				on:click|stopPropagation={showCreateBookmarkComponent}
				class:disabled={$processingBookmark}
			>
				{#if $processingBookmark}
					<div class="loader" />
				{:else}
					<i class="las la-plus" />
				{/if}
			</button>
		</div>
		<div class="bookmarks">
			{#each $bookmarks as { id, bookmark, title, thumbnail, notes, user_id, host, updated, favicon, added, deleted }}
				<div
					class="bookmark"
					data-id={id}
					data-bookmark={bookmark}
					data-title={title}
					data-thumbnail={thumbnail}
					data-notes={notes}
					data-userid={user_id}
					data-favicon={favicon}
					data-updated={updated}
					data-added={added}
					data-deleted={deleted}
				>
					<div class="thumbnail">
						<img src={thumbnail} alt="bookmark thumbnail" loading="lazy" draggable="false" />
					</div>
					<div class="title-favicon-and-domain">
						<div class="title">
							<a
								href={bookmark}
								target="_blank"
								on:click|preventDefault|stopPropagation={openBookmark}>{title}</a
							>
						</div>
						<div class="favicon-and-domain">
							{#if favicon === ''}
								<img src={GlobePNG} alt="bookmark favicon" draggable="false" />
							{:else}
								<img src={favicon} alt="bookmark favicon" draggable="false" />
							{/if}
							<span>{host}</span>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>

	<!-- overlay for popups -->
	<Overlay />

	<!-- add bookmark component -->
	<AddBookmark />

	<!-- create tag component -->
	<CreateTag />

	<!-- bookmark saved alert -->
	<TagCreated />

	<!-- duplicate tag alert -->
	<DuplicateTagAlert />

	<!-- tag created bookmark -->
	<TagAddedBookmark />
</div>

<style lang="scss">
	.app {
		height: 100vh;
		width: 100vw;
		position: fixed;
		top: 0;
		left: 0;
		display: flex;

		.sidebar {
			width: 27.5rem;
			height: 100%;
			background-color: rgb(255, 255, 255);
			display: flex;
			flex-direction: column;
			border-right: 0.1rem solid rgb(0, 0, 0, 0.1);
			transition: all ease 0.5s;

			.profile {
				height: 10%;
				background-color: rgb(255, 255, 255);
				display: flex;
				align-items: center;
				gap: 0.5em;
				width: 100%;
				padding: 0.5em;
				cursor: pointer;
				transition: 0.3s;
				border-radius: 0.3rem;
				position: relative;
				border-bottom: 0.1rem solid transparent;

				img {
					width: 3.5rem;
					height: 3.5rem;
					object-fit: cover;
					border-radius: 100vh;
				}

				.name-and-email {
					display: flex;
					flex-direction: column;
					gap: 0.3em;
					white-space: nowrap;
					overflow: hidden;
					text-overflow: ellipsis;

					h3 {
						font-size: 1.3rem;
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
						font-family: 'Arial CE', sans-serif;
					}

					span {
						font-size: 1.3rem;
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
						font-family: 'Arial CE', sans-serif;
					}
				}

				.userMenu {
					position: absolute;
					top: 100%;
					left: 0;
					right: 0;
					width: inherit;
					min-height: max-content;
					z-index: 1;
					border-radius: 0.3rem;
					display: flex;
					flex-direction: column;
					background-color: rgb(255, 255, 255);
					min-height: calc(100vh - 10%);
					transition: all ease 0.3s;
					border-top: 0.1rem solid rgb(0, 0, 0, 0.1);
					box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
					transform: translateX(-200%);

					div {
						min-height: 4rem;
						display: flex;
						align-items: center;
						gap: 1em;
						padding: 0.5em;
						background-color: inherit;
						border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);

						svg {
							height: 2rem;
							width: 2rem;
						}

						span {
							font-size: 1.3rem;
							font-family: 'Segoe UI', sans-serif;
						}

						&:hover {
							background-color: rgb(96, 1, 255, 0.1);
						}
					}
				}

				&:hover {
					border-bottom-color: rgb(0, 0, 0, 0.1);
				}
			}

			.tags {
				width: 100%;
				height: 85%;
				background-color: inherit;
				overflow-y: auto;
				padding-top: 1em;
				display: flex;
				flex-direction: column;

				.tag {
					width: 100%;
					padding: 0.5em;
					gap: 0.5em;
					cursor: pointer;
					display: flex;
					align-items: center;

					i {
						font-size: 2rem;
					}

					span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
					}
				}
			}

			.create-tag {
				width: 100%;
				height: 5%;
				border-top: 0.1rem solid rgb(0, 0, 0, 0.1);
				cursor: pointer;
				background-color: rgb(255, 255, 255);

				.new-tag {
					height: 100%;
					width: 100%;
					padding: 0 0.5em;
					display: flex;
					align-items: center;
					gap: 0.5em;
					transition: background-color ease 0.5s;

					i {
						font-size: 2rem;
					}

					span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
					}
				}

				&:hover {
					background-color: rgb(238, 238, 238);
				}
			}

			@media only screen and (max-width: 768px) {
				position: fixed;
				top: 0;
				left: 0;
				transform: translateX(-200%);
				//width: 100vw;
			}
		}

		.toggleSidebar {
			transform: translateX(0);
			box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		}

		.main {
			width: calc(100vw - 27.5rem);
			height: 100%;
			background-color: rgb(255, 255, 255);

			@media only screen and (max-width: 768px) {
				width: 100vw;
			}

			.search-or-add {
				display: flex;
				min-height: 7vh;
				//border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
				display: flex;
				align-items: center;
				justify-content: space-between;
				padding: 0.5em;

				i.la-bars {
					cursor: pointer;

					font-size: 3rem;

					@media only screen and (min-width: 769px) {
						display: none;
					}
				}

				input[type='search'] {
					width: 50%;
					padding: 0.5em;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					outline: none;
					border-radius: 0.3rem;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					pointer-events: none;
					min-height: 3.5rem;

					&:hover {
						border-color: rgb(255, 137, 137);
					}

					@media only screen and (max-width: 425px) {
						width: 70%;
					}
				}

				button {
					padding: 0.5em;
					outline: none;
					border-radius: 0.6rem;
					display: flex;
					align-items: center;
					cursor: pointer;
					gap: 1em;
					background-color: rgb(0, 121, 255);
					margin-right: 0em;
					border: none;
					box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;

					// span {
					// 	font-size: 1.3rem;
					// 	color: rgb(255, 255, 255);
					// 	text-transform: capitalize;
					// 	font-family: 'Arial CE', sans-serif;
					// }

					i {
						color: rgb(253, 253, 253);
						font-size: 2rem;
					}

					&:hover {
						background-color: rgb(6, 143, 255);
					}
				}

				.disabled {
					pointer-events: none;
				}

				.loader {
					border: 0.2rem solid #f3f3f3;
					border-radius: 50%;
					border-top: 0.2rem solid #3498db;
					width: 1.8rem;
					height: 1.8rem;
					-webkit-animation: spin 2s linear infinite; /* Safari */
					animation: spin 0.5s linear infinite;

					/* Safari */
					@-webkit-keyframes spin {
						0% {
							-webkit-transform: rotate(0deg);
						}
						100% {
							-webkit-transform: rotate(360deg);
						}
					}

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

			.bookmarks {
				height: calc(100vh - 7vh);
				padding: 0.5em;
				overflow-y: auto;
				display: flex;
				flex-wrap: wrap;
				align-content: flex-start;
				align-items: stretch;
				justify-content: flex-start;
				flex-flow: row wrap;
				gap: 1em;

				.bookmark {
					max-width: 35rem;
					border: 0.1rem solid rgb(96, 1, 255);
					display: flex;
					flex-direction: column;
					border-radius: 0.3rem;
					gap: 1em;
					box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
					padding: 0.5em;
					outline: 0.2rem solid transparent;
					transition: all ease 0.5s;
					word-break: break-word;
					flex-grow: 1;

					.thumbnail {
						height: 70%;
						width: 100%;
						display: flex;
						align-items: center;
						justify-content: center;
						border-radius: 0.3rem;

						img {
							height: 100%;
							width: 100%;
							max-inline-size: 100%;
							object-fit: fill;
							border-radius: inherit;
						}
					}

					.title-favicon-and-domain {
						height: 30%;
						display: flex;
						flex-direction: column;
						align-items: center;
						gap: 0.5em;

						.title {
							height: 50%;
							width: 100%;
							display: flex;
							align-items: center;

							a {
								display: -webkit-box;
								-webkit-line-clamp: 2;
								-webkit-box-orient: vertical;
								overflow: hidden;
								text-overflow: ellipsis;
								font-size: 1.3rem;
								line-height: 1.6;
								color: rgb(24, 23, 40);
								font-family: 'Arial CE', sans-serif;

								&:hover {
									color: rgb(78, 79, 235);
									text-decoration-color: rgb(78, 79, 235);
								}
							}
						}

						.favicon-and-domain {
							width: 100%;
							height: 50%;
							display: flex;
							align-items: center;
							gap: 1em;
							//justify-content: space-between;

							img {
								height: 2rem;
								width: 2rem;
								object-fit: cover;
								border-radius: 100vh;
							}

							span {
								color: rgb(24, 23, 40);
								font-size: 1.3rem;
								white-space: nowrap;
								overflow: hidden;
								text-overflow: ellipsis;
								max-width: 90%;
								font-family: 'Arial CE', sans-serif;
							}
						}
					}

					&:hover {
						outline-color: rgb(96, 1, 255);
						background-color: rgb(96, 1, 255, 0.2);
						box-shadow: 0px 5px 5px -3px rgba(0, 0, 0, 0.1), 0px 8px 8px 1px rgba(0, 0, 0, 0.07),
							0px 3px 8px 2px rgba(0, 0, 0, 0.08), 0px 0px 0px 2px;
					}
				}
			}
		}
	}

	// GLOBAL STYLES
	:global(.active-tag) {
		i {
			color: red !important;
		}

		span {
			color: red !important;
		}
	}

	:global(.toggleUserMenu) {
		transform: translateX(0) !important;
	}
</style>
