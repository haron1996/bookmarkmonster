<script lang="ts">
	import {
		apiHost,
		selectedBookmarks,
		selectedFolders,
		selectedItems,
		session,
		showMoveItemsPopup,
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

	async function handleClickOnQuickView() {
		const response = await fetch(`${$apiHost}/authenticated/bookmarks/updateBookmarkHtml`, {
			method: 'PATCH',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				bookmark_id: $selectedBookmarks[0].id
			})
		});

		const result = await response.json();

		console.log(result[0]);
	}
</script>

<div
	class="actionBar animate__animated animate__backInDown"
	id="actionBar"
	class:showActionBar={$selectedItems.length >= 1}
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
		<div
			class="move"
			role="none"
			on:click={() => {
				showMoveItemsPopup.set(true);
			}}
		>
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
		<div class="quickView" role="none" on:click={handleClickOnQuickView}>
			<i class="las la-eye" />
			<span>quick view</span>
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
		transition: all 200ms ease-in-out;
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
				min-width: max-content;
				padding: 0.5em 1em;
				display: flex;
				align-items: center;
				justify-content: center;
				background-color: #f0f0f0;
				cursor: pointer;
				gap: 0.7em;
				border-radius: 0.3rem;
				position: relative;
				transition: all 200ms ease-in-out;

				i {
					font-size: 2rem;
					transition: all 300ms ease;
				}

				span {
					font-size: 1.2rem;
					font-family: 'Segoe UI', sans-serif;
					text-transform: uppercase;
					transition: all 300ms ease;
				}

				.dis {
					position: absolute;
					left: 0;
					top: 0;
					cursor: not-allowed;
					height: inherit;
					width: inherit;
					background-color: transparent;

					&:hover {
						background-color: transparent;
					}
				}

				&:hover {
					background-color: #040d12;

					i {
						color: white;
					}

					span {
						color: white;
					}
				}
			}

			.disabled {
				opacity: 0.5;
			}
		}

		.count {
			padding: 0.7em 1em;
			min-width: max-content;
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
		height: 5rem;

		.actions {
			display: flex;
		}

		.count {
			display: flex;
		}
	}
</style>
