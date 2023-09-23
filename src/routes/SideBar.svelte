<script lang="ts">
	import { page } from '$app/stores';
	import {
		currentTag,
		hideSideBar,
		loadingItems,
		session,
		sideBarVisible,
		toolTipText
	} from '../stores/stores';
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
	class:hideSideBar={$hideSideBar}
	class="sidebar"
	id="sideBar"
	class:sideBarVisible={$sideBarVisible}
	on:click={closeTagMenu}
	role="none"
>
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
			role="none"
			on:click|preventDefault={() => {
				goto('/dashboard/tags');
				currentTag.set({});
				$hideSideBar = !$hideSideBar;
			}}
		>
			<a data-sveltekit-preload-data="tap" href="/dashboard/tags">
				<i class="las la-tags" />
			</a>
		</li>
		<li
			style="display:none"
			class:active={$page.url.pathname === '/dashboard/screenshots'}
			data-id="Screenshots"
			on:pointerover={showToolTip}
			on:pointerout={hideToolTip}
		>
			<a data-sveltekit-preload-data="tap" href="/dashboard/screenshots">
				<i class="las la-file-image" />
			</a>
		</li>
		<li
			style="display:none"
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

	<div class="profile" role="none" on:click={toggleUserMenu}>
		<i
			class="las la-angle-double-left"
			role="none"
			on:click={() => {
				$hideSideBar = !$hideSideBar;
			}}
		/>
		{#if $loadingItems}
			<div class="loader" />
		{:else if $session && $session.User && $session.User.picture && $session.User.name && $session.User.email}
			<img src={$session.User.picture} alt="profile" draggable="false" />
		{/if}
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
		width: 7rem;
		align-items: center;
		justify-content: space-between;
		background-color: rgb(240, 240, 240);

		ul {
			display: flex;
			flex-direction: column;
			gap: 1em;
			height: calc(100vh - 10rem);
			overflow-y: auto;
			width: 100%;
			align-items: center;
			justify-content: center;
			padding-top: 0.5em;

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
						font-size: 2rem;
					}
				}

				&:hover {
					a {
						color: #0079ff;
					}
				}

				&.active {
					a {
						color: #0079ff;
					}
				}
			}
		}

		.profile {
			display: flex;
			align-items: center;
			justify-content: space-between;
			width: 100%;
			transition: 0.3s;
			border-radius: 0.3rem;
			border-bottom: 0.1rem solid transparent;
			min-height: 10rem;
			height: max-content;
			flex-direction: column;
			//gap: 1em;
			padding: 0.5em 0;

			i {
				font-size: 1.5rem;
				cursor: pointer;
			}

			.loader {
				border: 0.2rem solid #f3f3f3;
				border-radius: 50%;
				border-top: 0.2rem solid #3498db;
				width: 3.5rem;
				height: 3.5rem;
				-webkit-animation: spin 2s linear infinite;
				animation: spin 0.5s linear infinite;

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
				width: 3rem;
				height: 3rem;
				object-fit: cover;
				border-radius: 100vh;
				cursor: pointer;
				transition: all 200ms ease-in-out;

				&:hover {
					box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
				}
			}
		}
	}

	.hideSideBar {
		display: none;
	}
</style>
