<script lang="ts">
	import { page } from '$app/stores';
	import { loadingItems, session, sideBarVisible, toolTipText } from '../stores/stores';
	import { closeTagMenu } from '../utils/closeTagMenu';
	import { toggleSideBar } from '../utils/toggleSideBar';
	import logo from '$lib/images/logo.png';
	import { goto } from '$app/navigation';
	import ToolTip from './ToolTip.svelte';

	function toggleUserMenu() {
		const userMenu = document.getElementById('userMenu') as HTMLDivElement | null;

		if (userMenu === null) return;

		userMenu.classList.toggle('toggleUserMenu');
	}

	function logUserOut() {
		localStorage.removeItem('session');
		alert('log out successful');
		goto('/signin');
		return;
	}

	function showToolTip(e: PointerEvent) {
		const t = e.currentTarget as HTMLElement;
		const li = t.closest('li') as HTMLLIElement | null;
		if (li === null) return;
		const text: string | undefined = li.dataset.id;
		if (text === undefined) return;
		toolTipText.set(text);
		const tooltip = document.getElementById('toolTip') as HTMLDivElement | null;

		if (tooltip === null) return;

		tooltip.style.top = `${e.clientY}px`;

		tooltip.style.left = `${e.clientX + 30}px`;

		tooltip.style.visibility = 'visible';
	}

	function hideToolTip() {
		const tooltip = document.getElementById('toolTip') as HTMLDivElement | null;
		if (tooltip === null) return;
		tooltip.style.visibility = 'hidden';
	}
</script>

<ToolTip />

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
			<li
				class:active={$page.url.pathname === '/dashboard'}
				data-id="Dashboard"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard">
					<i class="las la-home" />
				</a>
			</li>
			<li
				class:active={$page.url.pathname === '/dashboard/allBookmarks'}
				data-id="All bookmarks"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard/allBookmarks">
					<i class="las la-external-link-square-alt" />
				</a>
			</li>
			<li
				class:active={$page.url.searchParams.get('id')}
				data-id="Collections"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard/my_collections?id=root">
					<i class="las la-folder-open" />
				</a>
			</li>
			<li
				class:active={$page.url.pathname === '/dashboard/tags'}
				data-id="Tags"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard/tags">
					<i class="las la-tags" />
				</a>
			</li>
			<li
				class:active={$page.url.pathname === '/dashboard/screenshots'}
				data-id="Screenshots"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard/screenshots">
					<i class="las la-image" />
				</a>
			</li>
			<li
				class:active={$page.url.pathname === '/dashboard/PDFs'}
				data-id="PDFs"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard/PDFs">
					<i class="las la-file-pdf" />
				</a>
			</li>
			<li
				class:active={$page.url.pathname === '/dashboard/recycle_bin'}
				data-id="Recycle bin"
				on:pointerover={showToolTip}
				on:pointerout={hideToolTip}
			>
				<a data-sveltekit-preload-data="tap" href="/dashboard/recycle_bin">
					<i class="las la-recycle" />
				</a>
			</li>
		</ul>
	</div>

	<li class="signout" data-id="Log Out">
		<i
			class="las la-sign-out-alt"
			role="none"
			on:click={logUserOut}
			on:pointerover={showToolTip}
			on:pointerout={hideToolTip}
		/>
	</li>
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
			gap: 2em;
			min-height: 90vh;
			padding-top: 0.5em;
			z-index: inherit;

			.logo {
				height: 3.5rem;
				width: 3.5rem;
				align-items: center;
				justify-content: center;
				display: none;

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
				cursor: pointer;
				transition: 0.3s;
				border-radius: 0.3rem;
				border-bottom: 0.1rem solid transparent;
				min-height: 10vh;
				min-height: max-content;

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
					width: 4rem;
					height: 4rem;
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
				height: 90vh;
				overflow-y: auto;
				width: 100%;
				display: flex;
				flex-direction: column;
				align-items: center;

				&::-webkit-scrollbar {
					display: none;
				}

				li {
					list-style: none;
					border-radius: 100vh;
					position: relative;
					height: 4rem;
					width: 4rem;
					display: flex;
					align-items: center;
					justify-content: center;

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
					}

					&:hover {
						background-color: rgb(238, 238, 238);
					}

					&.active {
						background-color: #040d12;

						i {
							color: white;
						}
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
				height: 4rem;
				width: 4rem;
				font-size: 2.4rem;
				display: flex;
				align-items: center;
				justify-content: center;
				transition: background-color ease 0.5s;
				cursor: pointer;
				border-radius: 100vh;

				&:hover {
					background-color: #040d12;
					color: white;
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
