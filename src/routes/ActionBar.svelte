<script lang="ts">
	import { selectedBookmarks } from '../stores/stores';
	import { openTrashBookmarkConfirmation } from '../utils/openTrashBookmarkConfirmation';
	import { showBookmarkDetails } from '../utils/showBookmarkDetails';
	import { showRenameBookmark } from '../utils/showRenameBookmark';
	import { showTagBookmarkComp } from '../utils/showTagBookmarkComp';
	import { unselectBookmarkNodes } from '../utils/unselectBookmarkNodes';

	function clearSelectedBookmarks() {
		selectedBookmarks.set([]);

		unselectBookmarkNodes();
	}
</script>

{#if $selectedBookmarks.length >= 1}
	<div class="actionBar" id="actionBar">
		<div class="left">
			<div class="selectCount">
				{#if $selectedBookmarks.length >= 1}
					<span>{$selectedBookmarks.length} selected</span>
					<i class="lar la-trash-alt" role="none" on:click={clearSelectedBookmarks} />
				{/if}
			</div>
		</div>
		<div class="actions">
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
			<i class="lar la-trash-alt" on:click={openTrashBookmarkConfirmation} role="none" />
		</div>
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
		justify-content: space-between;
		padding: 0 1em 0 1em;

		.left {
			min-height: inherit;
			min-width: max-content;
			display: flex;
			align-items: center;

			.selectCount {
				padding: 0.5em 1em;
				display: flex;
				align-items: center;
				gap: 1em;
				border-radius: 0.3rem;
				background-color: rgb(255, 255, 255);
				border: 0.1rem solid rgb(2, 84, 100);

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
				}

				i {
					font-size: 2rem;
					color: #f45050;
					cursor: pointer;
					transform: scale(1);
					transition: all ease 0.2s;

					&:hover {
						transform: scale(1.1);
					}
				}
			}
		}

		.actions {
			display: flex;
			align-items: center;
			gap: 1em;
			min-height: inherit;
			min-width: max-content;

			i {
				font-size: 2rem;
				cursor: pointer;
				transform: scale(1);
				transition: all ease 0.2s;

				&:hover {
					color: rgb(2, 84, 100);
					transform: scale(1.1);
				}
			}

			.disabled {
				opacity: 0.5;
				pointer-events: none;
			}
		}

		@media only screen and (max-width: 768px) {
			left: 0;
			width: 100vw;
		}
	}
</style>
