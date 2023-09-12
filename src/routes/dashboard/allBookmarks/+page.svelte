<script lang="ts">
	import { page } from '$app/stores';
	import { allItems, bookmarks, loadingItems } from '../../../stores/stores';
	import ActionBarV2 from '../../ActionBarV2.svelte';
	import PlainBookmark from '../../PlainBookmark.svelte';
	import TopBar from '../../TopBar.svelte';
	import nodata from '$lib/images/no-data.jpg';

	$: $page.data.bookmarks,
		$page.data.bookmarks.length >= 1 ? bookmarks.set($page.data.bookmarks) : bookmarks.set([]);
</script>

<svelte:head>
	<title>All Bookmarks | BookmarkMonster</title>
</svelte:head>

<div class="allBookmarks">
	<ActionBarV2 />
	<TopBar />
	<div class="bookmarks">
		{#if $loadingItems}
			<div class="loadingItems">
				<p>Loading...</p>
				<div class="loader" />
			</div>
		{:else if !$loadingItems && $allItems.length >= 1}
			<PlainBookmark />
		{:else if !$loadingItems && $allItems.length < 1}
			<div class="noItems">
				<img src={nodata} alt="no data" />
				<p>You have no bookmarks yet</p>
			</div>
		{/if}
	</div>
</div>

<style lang="scss">
	.bookmarks {
		padding: 1em;
		display: flex;
		flex-flow: row wrap;
		align-content: start;
		gap: 1em;
		height: calc(100vh - calc(8vh));
		overflow-y: auto;
		width: 100%;
		max-width: 100%;
		background-color: white;

		.loadingItems {
			width: 100%;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			flex-direction: column;
			gap: 1em;

			p {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
			}

			.loader {
				border: 2px solid #f3f3f3; /* Light grey */
				border-top: 2px solid #3498db; /* Blue */
				border-radius: 50%;
				width: 3rem;
				height: 3rem;
				animation: spin 0.5s linear infinite;

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

		.noItems {
			width: 100%;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			flex-direction: column;

			img {
				// max-inline-size: 100%;
				// object-fit: contain;
				width: 50%;
				height: 50%;
				object-fit: contain;
			}

			p {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
			}
		}
	}
</style>
