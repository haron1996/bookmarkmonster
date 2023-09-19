<script lang="ts">
	import { onMount } from 'svelte';
	import { selectedBookmarks, showQuickView } from '../stores/stores';

	// kcb
	// kilimall
	// equity eazzy banking
	// furniture palace
	// quickbooks

	// onMount(() => {
	// 	console.log({ ...localStorage });
	// 	console.log({ ...sessionStorage });
	// });
	// sandbox="allow-forms allow-modals allow-pointer-lock allow-popups allow-popups-to-escape-sandbox allow-scripts allow-top-navigation"
</script>

{#if $showQuickView}
	<div
		class="container"
		id="quickView"
		role="none"
		on:click={() => {
			showQuickView.set(false);
		}}
	>
		{#if $selectedBookmarks[0]}
			<iframe
				src={$selectedBookmarks[0].bookmark}
				frameborder="0"
				title="quick view"
				allowfullscreen
				on:click|stopPropagation={() => {}}
				role="none"
			/>
		{/if}
	</div>
{/if}

<style lang="scss">
	#quickView {
		z-index: 5;
		position: fixed;
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: white;

		iframe {
			height: 100vh;
			width: 60vw;
			background-color: inherit;
			position: relative;
			display: flex;
			flex-direction: column;
			border: none;
			outline: none;
			background-color: white;
		}
	}
</style>
