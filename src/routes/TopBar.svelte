<script lang="ts">
	import { page } from '$app/stores';
	import { hideSideBar, query, showAddNewBookmark } from '../stores/stores';
	import { openCreateBookmark } from '../utils/openCreateBookmark';
	import { openCreateFolder } from '../utils/openCreateFolder';

	let addMenuShown: boolean = false;
	let searchPlaceHolder: string = '';

	$: if ($page.url.pathname === '/dashboard/my_collections') {
		searchPlaceHolder = 'Search collections';
	} else if ($page.url.pathname === '/dashboard/allBookmarks') {
		searchPlaceHolder = 'Search bookmarks';
	}
</script>

<div class="top">
	<i
		class="las la-bars"
		role="none"
		on:click={() => {
			$hideSideBar = !$hideSideBar;
		}}
	/>
	<div class="search">
		<input
			type="search"
			name="search"
			id="search"
			placeholder={searchPlaceHolder}
			autocomplete="off"
			spellcheck="false"
		/>
		{#if $query}
			<div class="clearOrLoadIcon">
				<i class="las la-times" />
			</div>
		{/if}
	</div>

	{#if $page.url.pathname === '/dashboard/my_collections'}
		<div
			class="new"
			on:click={() => {
				addMenuShown = !addMenuShown;
			}}
			role="none"
		>
			<i class="las la-plus" />
			<!-- <span class="addNew">add new</span>
			<i class="las la-angle-down" /> -->

			{#if addMenuShown}
				<div class="addMenu">
					<div class="createBookmark" on:click={openCreateBookmark} role="none">
						<svg
							width="24px"
							height="24px"
							stroke-width="1.5"
							viewBox="0 0 24 24"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							color="#000000"
							><path
								d="M14 11.998C14 9.506 11.683 7 8.857 7H7.143C4.303 7 2 9.238 2 11.998c0 2.378 1.71 4.368 4 4.873a5.3 5.3 0 001.143.124"
								stroke="#000000"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/><path
								d="M10 11.998c0 2.491 2.317 4.997 5.143 4.997h1.714c2.84 0 5.143-2.237 5.143-4.997 0-2.379-1.71-4.37-4-4.874A5.304 5.304 0 0016.857 7"
								stroke="#000000"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/></svg
						>
						<span>bookmark</span>
					</div>
					<div class="createCollection" on:click={openCreateFolder} role="none">
						<svg
							width="24px"
							height="24px"
							stroke-width="1.5"
							viewBox="0 0 24 24"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							color="#000000"
							><path
								d="M18 6h2m2 0h-2m0 0V4m0 2v2M21.4 20H2.6a.6.6 0 01-.6-.6V11h19.4a.6.6 0 01.6.6v7.8a.6.6 0 01-.6.6zM2 11V4.6a.6.6 0 01.6-.6h6.178a.6.6 0 01.39.144l3.164 2.712a.6.6 0 00.39.144H14"
								stroke="#000000"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/></svg
						>
						<span>collection</span>
					</div>
				</div>
			{/if}
		</div>
	{/if}

	{#if $page.url.pathname === '/dashboard/allBookmarks'}
		<div
			class="newBookmark"
			on:click|preventDefault={() => {
				showAddNewBookmark.set(true);
			}}
			role="none"
		>
			<i class="las la-plus" />
			<!-- <span class="addNew">add new</span>
			<i class="las la-angle-down" /> -->
		</div>
	{/if}
</div>

<style lang="scss">
	.top {
		width: 100%;
		padding: 0 1em;
		display: flex;
		z-index: 1;
		align-items: center;
		justify-content: space-between;
		background-color: white;
		height: 5rem;

		i.la-bars {
			font-size: 2rem;
			cursor: pointer;
		}

		.search {
			display: flex;
			align-items: center;
			width: 40rem;
			height: 3rem;
			border-radius: 0.3rem;
			background-color: white;

			input {
				border: none;
				outline: none;
				width: 90%;
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
				height: 100%;
				padding: 0 0.5em;
				background-color: inherit;

				&::placeholder {
					color: #d8d9da;
				}
			}

			::-webkit-search-cancel-button {
				display: none;
			}

			.clearOrLoadIcon {
				min-width: 10%;
				display: flex;
				align-items: center;
				justify-content: center;
				height: 100%;
				cursor: pointer;

				i {
					font-size: 1.8rem;
				}
			}

			@media only screen and (width <= 425px) {
				width: 50%;
			}
		}

		.new {
			min-width: max-content;
			display: flex;
			align-items: center;
			gap: 0.7em;
			background-color: inherit;
			color: #040d12;
			cursor: pointer;
			transition: all 300ms ease;
			position: relative;

			i.la-plus {
				font-size: 2rem;
			}

			.addMenu {
				position: absolute;
				top: 100%;
				right: 0;
				background-color: white;
				width: max-content;
				height: max-content;
				border-radius: 0.5rem;
				box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px, rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
				padding: 1.5em;
				display: flex;
				gap: 1em;
				color: #001c30;

				div {
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					gap: 1em;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					padding: 1em;
					border-radius: 0.5rem;
					height: 10rem;
					width: 10rem;
					transition: all 300ms ease;

					svg {
						height: 2rem;
						width: 2rem;
					}

					span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						text-transform: capitalize;
					}

					&:hover {
						background-color: #eeeeee;
					}
				}
			}
		}

		.newBookmark {
			min-width: max-content;
			display: flex;
			align-items: center;
			gap: 0.7em;
			background-color: inherit;
			border-radius: 0.3rem;
			padding: 1em 2em;
			color: #040d12;
			cursor: pointer;
			transition: all 300ms ease;
			position: relative;

			i.la-plus {
				font-size: 2rem;
			}
		}

		@media only screen and (width <= 600px) {
			justify-content: space-between;
		}
	}
</style>
