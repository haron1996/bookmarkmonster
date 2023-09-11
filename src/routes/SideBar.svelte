<script lang="ts">
	import { page } from '$app/stores';
	import {
		currentTagID,
		loadingItems,
		selectedTags,
		session,
		sideBarVisible,
		tags
	} from '../stores/stores';
	import { closeTagMenu } from '../utils/closeTagMenu';
	import type { Tag } from '../types/tag';
	import { showCreateTagComponent } from '../utils/showCreateTagComponent';
	import { toggleSideBar } from '../utils/toggleSideBar';
	import logo from '$lib/images/logo.png';
	import { goto } from '$app/navigation';

	function toggleUserMenu() {
		const userMenu = document.getElementById('userMenu') as HTMLDivElement | null;

		if (userMenu === null) return;

		userMenu.classList.toggle('toggleUserMenu');
	}

	// function logUserOut() {
	// 	const session = localStorage.getItem('session') as string | null;

	// 	if (session === null) return;

	// 	localStorage.removeItem('session');

	// 	window.location.reload();
	// }

	// const handleClickOnTag = (e: MouseEvent) => {
	// 	closeTagMenu();

	// 	const clickedEl = e.target as HTMLElement;

	// 	const clickedTag = clickedEl.closest('.tag') as HTMLSpanElement | null;

	// 	if (clickedTag === null) return;

	// 	const tags = document.querySelectorAll('.tag') as NodeListOf<HTMLSpanElement>;

	// 	tags.forEach((tag) => {
	// 		if (tag.classList.contains('active-tag')) {
	// 			tag.classList.remove('active-tag');
	// 		}
	// 	});

	// 	clickedTag.classList.add('active-tag');

	// 	if (clickedTag.closest('.all-tags') === null) {
	// 		if (clickedTag.dataset.id) {
	// 			currentTagID.set(clickedTag.dataset.id);
	// 		}
	// 	} else if (clickedTag.closest('.all-tags')) {
	// 		if (clickedTag.dataset.id) {
	// 			currentTagID.set(clickedTag.dataset.id);
	// 		}
	// 	}
	// };

	// function openTagMenu(e: MouseEvent) {
	// 	const target = e.currentTarget as HTMLElement;

	// 	const tag = target.closest('.tag') as HTMLDivElement | null;

	// 	if (tag === null) return;

	// 	const t: Tag = {
	// 		id: tag.dataset.id,
	// 		name: tag.dataset.name
	// 	};

	// 	selectedTags.set([]);

	// 	selectedTags.update((values) => [t, ...values]);

	// 	const menu = document.getElementById('tagMenu') as HTMLElement | null;

	// 	if (menu === null) return;

	// 	menu.style.display = 'flex';

	// 	menu.style.top = `${e.clientY}px`;

	// 	menu.style.left = `${e.clientX}px`;
	// }

	// function handleClickOnRecycleBin() {
	// 	const url: URL = new URL(`${$page.url.origin}/dashboard/my_bookmarks`);
	// 	url.searchParams.append('id', 'root');
	// 	url.searchParams.append('which', 'recycle_bin');
	// 	goto(url);
	// }

	// function handleClickOnMyBookmarks() {
	// 	const url: URL = new URL(`${$page.url.origin}/dashboard/my_bookmarks`);
	// 	url.searchParams.append('id', 'root');
	// 	goto(url);
	// }
</script>

<div
	class="sidebar"
	id="sideBar"
	class:sideBarVisible={$sideBarVisible}
	on:click={closeTagMenu}
	role="none"
>
	<div class="top">
		<a data-sveltekit-preload-data="tap" href={`${$page.url.origin}/dashboard`} class="logo">
			<img src={logo} alt="logo" draggable="false" />
		</a>

		<div class="profile" role="none" on:click={toggleUserMenu}>
			{#if $loadingItems}
				<div class="loader" />
			{:else if $session && $session.User && $session.User.picture && $session.User.name && $session.User.email}
				<img src={$session.User.picture} alt="profile" draggable="false" />

				<div class="name-and-email">
					<h3>{$session.User.name}</h3>
					<span>{$session.User.email}</span>
				</div>
				<i class="las la-angle-down" />
			{/if}
		</div>

		<ul>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard">
					<i class="las la-home" />
					<span>Dashboard</span></a
				>
			</li>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard/allBookmarks">
					<i class="las la-external-link-square-alt" />
					<span>Bookmarks</span></a
				>
			</li>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard/my_collections?id=root">
					<i class="las la-folder-open" />
					<span>Folders</span>
				</a>
			</li>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard/tags">
					<i class="las la-tags" />
					<span>Tags</span>
				</a>
			</li>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard/screenshots">
					<i class="las la-image" />
					<span>Screenshots</span>
				</a>
			</li>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard/PDFs">
					<i class="las la-file-pdf" />
					<span>PDFs</span>
				</a>
			</li>
			<li>
				<a data-sveltekit-preload-data="tap" href="/dashboard/recycle_bin">
					<i class="las la-recycle" />
					<span>Recycle bin</span></a
				>
			</li>
		</ul>
	</div>

	<div class="signout">
		<i class="las la-sign-out-alt" />
	</div>
</div>

<div
	id="closeSideBar"
	class:closeSideBarVisible={$sideBarVisible}
	on:click={toggleSideBar}
	role="none"
/>

<style lang="scss">
	.sidebar {
		height: 100vh;
		display: flex;
		flex-direction: column;
		border-right: 0.1rem solid rgb(0, 0, 0, 0.1);
		transition: all ease 0.5s;
		position: relative;
		z-index: 3;
		min-width: 9rem;
		align-items: center;
		justify-content: space-between;

		.top {
			display: flex;
			flex-direction: column;
			align-items: center;
			width: 100%;
			padding-top: 2em;
			gap: 2em;
			min-height: 90vh;

			.logo {
				height: 3.5rem;
				width: 3.5rem;
				align-items: center;
				justify-content: center;

				img {
					max-inline-size: 100%;
					object-fit: cover;
				}
			}

			.profile {
				display: flex;
				align-items: center;
				justify-content: space-evenly;
				width: 100%;
				padding: 1em 0.5em;
				cursor: pointer;
				transition: 0.3s;
				border-radius: 0.3rem;
				border-bottom: 0.1rem solid transparent;
				min-height: 10vh;

				.loader {
					border: 0.2rem solid #f3f3f3;
					border-radius: 50%;
					border-top: 0.2rem solid #3498db;
					width: 3.5rem;
					height: 3.5rem;
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
					display: none;

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

				i {
					font-size: 1.5rem;
					display: none;
				}
			}

			ul {
				display: flex;
				flex-direction: column;
				gap: 1em;
				max-height: 80vh;
				overflow-y: auto;

				&::-webkit-scrollbar {
					display: none;
				}

				li {
					list-style: none;
					border-radius: 100vh;

					a {
						text-decoration: none;
						display: flex;
						align-items: center;
						justify-content: center;
						gap: 0.7em;
						padding: 1em 2em;
						cursor: pointer;
						transition: all ease 300ms;
						height: 5rem;
						width: 5rem;
						padding: 0.7em;
						color: #272829;
						border-radius: inherit;

						i {
							font-size: 2.4rem;
						}

						span {
							font-size: 1.5rem;
							font-family: 'Segoe UI', sans-serif;
							display: none;
						}
					}

					&:hover {
						background-color: rgb(238, 238, 238);
					}
				}
			}
		}

		.signout {
			width: 100%;
			min-height: 10vh;
			display: flex;
			align-items: center;
			justify-content: center;

			i {
				height: 5rem;
				width: 5rem;
				font-size: 2.4rem;
				display: flex;
				align-items: center;
				justify-content: center;
				transition: background-color ease 0.5s;
				cursor: pointer;
				border-radius: 100vh;

				&:hover {
					background-color: rgb(238, 238, 238);
				}
			}
		}
	}

	.sideBarVisible {
		transform: translateX(0);
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}

	#closeSideBar {
		position: fixed;
		left: 27.5rem;
		top: 0;
		height: 100vh;
		width: calc(100vw - 27.5rem);
		background-color: rgb(0, 0, 0, 0.4);
		z-index: 5;
		opacity: 0;
		display: none;
		transition: opacity 0.5s;
	}

	.closeSideBarVisible {
		display: flex !important;
	}
</style>
