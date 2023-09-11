<script lang="ts">
	import { page } from '$app/stores';
	import { bookmarks } from '../../../stores/stores';
	import ActionBarV2 from '../../ActionBarV2.svelte';
	import PlainBookmark from '../../PlainBookmark.svelte';
	import TopBar from '../../TopBar.svelte';

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
		<PlainBookmark />
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
	}
</style>
