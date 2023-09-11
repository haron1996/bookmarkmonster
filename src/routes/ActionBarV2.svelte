<script lang="ts">
	import {
		selectedBookmarks,
		selectedFolders,
		selectedItems,
		showUpdateBookmark,
		showUpdateFolder
	} from '../stores/stores';
	import { removeSelectedClassFromAllDomFolders } from '../utils/removeSelectedClassFromAllDomFolders';
	import { removeSelectedClassFromAllDomBookmarks } from '../utils/removeSelectedClassFromAllSelectedDomBookmarks';
	import { removeSlideInAnimationFromActionBarV2 } from '../utils/removeSlideInAnimationFromActionBarV2';

	function handleClickOnUpdate() {
		if ($selectedItems.length >= 1) {
			if ($selectedFolders.length >= 1) {
				showUpdateFolder.set(true);
			} else if ($selectedBookmarks.length >= 1) {
				showUpdateBookmark.set(true);
			}
		}
	}
</script>

<div
	class="actionBar animate__animated animate__backInDown"
	class:showActionBar={$selectedItems.length >= 1}
	id="actionBar"
>
	<div class="actions">
		<div class="share">
			<i class="las la-share" />
			<span>Share</span>
		</div>
		<div class="delete">
			<i class="las la-trash" />
			<span>Delete</span>
		</div>
		<div class="move">
			<i class="las la-arrows-alt" />
			<span>Move</span>
		</div>
		<div
			class="rename"
			on:click={handleClickOnUpdate}
			role="none"
			class:disabled={$selectedItems.length > 1}
		>
			<i class="las la-edit" />
			<span>Update</span>
			{#if $selectedItems.length > 1}
				<div class="dis" on:click|stopPropagation|preventDefault={() => {}} role="none" />
			{/if}
		</div>
		<div class="beautify" class:disabled={$selectedFolders.length >= 1}>
			<i class="las la-magic" />
			<span>Beautify</span>
			{#if $selectedFolders.length >= 1}
				<div class="dis" on:click|stopPropagation|preventDefault={() => {}} role="none" />
			{/if}
		</div>
	</div>
	<div class="count">
		<span>{$selectedItems.length} selected</span>
		<i
			class="las la-times"
			role="none"
			on:click={() => {
				selectedBookmarks.set([]);
				selectedFolders.set([]);
				removeSelectedClassFromAllDomFolders();
				removeSelectedClassFromAllDomBookmarks();
				removeSlideInAnimationFromActionBarV2();
			}}
		/>
	</div>
</div>

<style lang="scss">
	.actionBar {
		padding: 0 1em;
		display: flex;
		align-items: center;
		justify-content: space-between;
		transition: all 1s ease-in-out;
		height: 0%;
		position: fixed;
		left: 9rem;
		width: calc(100vw - 9rem);
		z-index: 2;
		background-color: white;

		.actions {
			align-items: center;
			gap: 1em;
			display: none;

			div {
				height: 3.5rem;
				width: 10rem;
				display: flex;
				align-items: center;
				justify-content: center;
				background-color: #f0f0f0;
				cursor: pointer;
				gap: 0.7em;
				border-radius: 0.3rem;
				position: relative;

				i {
					font-size: 2rem;
				}

				span {
					font-size: 1.2rem;
					font-family: 'Segoe UI', sans-serif;
					text-transform: uppercase;
				}

				.dis {
					position: absolute;
					left: 0;
					top: 0;
					cursor: not-allowed;
					height: inherit;
					width: inherit;
					background-color: transparent;
				}
			}

			.disabled {
				opacity: 0.5;
			}
		}

		.count {
			padding: 0.7em 1em;
			align-items: center;
			background-color: #f0f0f0;
			gap: 0.7em;
			border-radius: 0.3rem;
			transition: all 300ms ease;
			display: none;

			i {
				font-size: 2rem;
				cursor: pointer;
			}

			span {
				font-size: 1.2rem;
				font-family: 'Segoe UI', sans-serif;
				text-transform: uppercase;
			}
		}
	}

	.showActionBar {
		display: flex;
		min-height: 8vh;

		.actions {
			display: flex;
		}

		.count {
			display: flex;
		}
	}
</style>
