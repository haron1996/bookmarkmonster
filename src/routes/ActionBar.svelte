<script>
	import { selectedBookmarks } from '../stores/stores';
	import { openTrashBookmarkConfirmation } from '../utils/openTrashBookmarkConfirmation';
	import { showBookmarkDetails } from '../utils/showBookmarkDetails';
	import { showRenameBookmark } from '../utils/showRenameBookmark';
	import { showTagBookmarkComp } from '../utils/showTagBookmarkComp';
</script>

{#if $selectedBookmarks.length >= 1}
	<div class="actionBar" id="actionBar">
		<i
			class="las la-keyboard"
			class:disabled={$selectedBookmarks.length > 1}
			on:click={showRenameBookmark}
			role="none"
		/>
		<i
			class="las la-hashtag"
			class:disabled={$selectedBookmarks.length > 1}
			on:click={showTagBookmarkComp}
			role="none"
		/>
		<i
			class="las la-info"
			class:disabled={$selectedBookmarks.length > 1}
			on:click={showBookmarkDetails}
			role="none"
		/>
		<i class="las la-trash" on:click={openTrashBookmarkConfirmation} role="none" />
	</div>
{/if}

<style lang="scss">
	#actionBar {
		width: calc(100vw - 29rem);
		position: fixed;
		left: 27.5rem;
		width: calc(100vw - 27.5rem);
		z-index: 1;
		background-color: rgb(240, 240, 240);
		min-height: 7vh;
		display: flex;
		align-items: center;
		justify-content: flex-end;
		padding: 0 1em 0 1em;
		gap: 1em;

		i {
			font-size: 2.4rem;
			cursor: pointer;
			border-radius: 0.5rem;
			height: 3rem;
			width: 3rem;
			display: flex;
			align-items: center;
			justify-content: center;

			&:hover {
				color: rgb(96, 1, 255);
				background-color: rgb(250, 250, 250);
			}
		}

		.disabled {
			opacity: 0.5;
			pointer-events: none;
		}

		@media only screen and (max-width: 768px) {
			left: 0;
			width: 100vw;
		}
	}
</style>
