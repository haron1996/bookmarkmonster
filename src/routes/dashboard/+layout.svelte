<script lang="ts">
	import UpdateFolder from '../UpdateFolder.svelte';
	import SideBar from '../SideBar.svelte';
	import UpdateBookmark from '../UpdateBookmark.svelte';
	import { afterNavigate } from '$app/navigation';
	import { loadingItems } from '../../stores/stores';
	import NewBookmark from '../NewBookmark.svelte';

	afterNavigate(() => {
		loadingItems.set(false);
	});
</script>

<div class="app">
	<NewBookmark />
	<UpdateFolder />
	<UpdateBookmark />
	<div class="bottom">
		<SideBar />
		<div class="content">
			<slot />
		</div>
	</div>
</div>

<style lang="scss">
	@import url('https://fonts.cdnfonts.com/css/arial');
	@import url('https://fonts.cdnfonts.com/css/segoe-ui-4');
	.app {
		position: fixed;
		left: 0;
		top: 0;
		height: 100vh;
		width: 100vw;
		overflow: hidden;
		display: flex;
		flex-direction: column;
		background-color: rgb(250, 250, 250);

		.bottom {
			display: flex;

			.content {
				height: 100%;
				width: calc(100vw - 9rem);
			}
		}
	}
</style>
