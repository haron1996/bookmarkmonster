<script lang="ts">
	import { bookmarks } from '../stores/stores';
	import { selectBookmark } from '../utils/selectBookmark';
	import GlobePNG from '$lib/images/globe.png';
	import { showBookmarkDetails } from '../utils/showBookmarkDetails';

	function openBookmark(e: MouseEvent) {
		const el = e.currentTarget as HTMLElement;

		const b = el.closest('.bookmark') as HTMLDivElement | null;

		if (b === null) return;

		const bookmarkLink: string | undefined = b.dataset.bookmark;

		if (bookmarkLink === undefined) return;

		window.open(bookmarkLink, '_blank');
	}

	function handleClickOnBookmarkInfoIcon(e: MouseEvent) {
		selectBookmark(e);
		showBookmarkDetails();
	}
</script>

{#each $bookmarks as { id, bookmark, title, thumbnail, notes, user_id, host, updated, favicon, added, deleted }}
	<div
		class="bookmark"
		data-id={id}
		data-bookmark={bookmark}
		data-title={title}
		data-thumbnail={thumbnail}
		data-notes={notes}
		data-userid={user_id}
		data-favicon={favicon}
		data-updated={updated}
		data-added={added}
		data-deleted={deleted}
		on:click|stopPropagation={selectBookmark}
		role="none"
	>
		<div class="thumbnail">
			<img src={thumbnail} alt="bookmark thumbnail" loading="lazy" draggable="false" />
		</div>
		<div class="title-favicon-and-domain">
			<div class="title">
				<a href={bookmark} target="_blank" on:click|preventDefault|stopPropagation={openBookmark}
					>{title}</a
				>
			</div>
			<div class="bottom">
				<div class="favicon-and-domain">
					{#if favicon === ''}
						<img src={GlobePNG} alt="bookmark favicon" draggable="false" />
					{:else}
						<img src={favicon} alt="bookmark favicon" draggable="false" />
					{/if}
					<span>{host}</span>
				</div>
				<i class="las la-info" role="none" on:click={handleClickOnBookmarkInfoIcon} />
			</div>
		</div>
	</div>
{/each}

<style lang="scss">
	.bookmark {
		max-width: 35rem;
		height: 35rem;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		gap: 1em;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		padding: 0.5em;
		border: 0.1rem solid #025464;
		outline: 0.2rem solid #025464;
		transition: all ease 0.5s;
		word-break: break-word;
		flex-grow: 1;

		@media only screen and (max-width: 568px) {
			max-width: 50rem;
		}

		@media only screen and (min-width: 728px) and (max-width: 425px) {
			max-width: 30rem;
		}

		.thumbnail {
			height: 70%;
			width: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			border-radius: 0.3rem;

			img {
				height: 100%;
				width: 100%;
				max-inline-size: 100%;
				object-fit: fill;
				border-radius: inherit;
			}
		}

		.title-favicon-and-domain {
			height: 30%;
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 0.5em;

			.title {
				height: 50%;
				width: 100%;
				display: flex;
				align-items: center;

				a {
					display: -webkit-box;
					-webkit-line-clamp: 2;
					-webkit-box-orient: vertical;
					overflow: hidden;
					text-overflow: ellipsis;
					font-size: 1.3rem;
					line-height: 1.6;
					color: rgb(24, 23, 40);
					font-family: 'Arial CE', sans-serif;

					&:hover {
						color: rgb(78, 79, 235);
						text-decoration-color: rgb(78, 79, 235);
					}
				}
			}

			.bottom {
				display: flex;
				align-items: center;
				justify-content: space-between;
				height: 50%;
				width: 100%;

				.favicon-and-domain {
					width: 100%;
					display: flex;
					align-items: center;
					gap: 1em;
					//justify-content: space-between;

					img {
						height: 2rem;
						width: 2rem;
						object-fit: cover;
						border-radius: 100vh;
					}

					span {
						color: rgb(24, 23, 40);
						font-size: 1.3rem;
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
						max-width: 90%;
						font-family: 'Arial CE', sans-serif;
					}
				}

				i {
					font-size: 2rem;
					cursor: pointer;
				}
			}
		}

		&:hover {
			background-color: #aebdca;
			box-shadow: 0px 5px 5px -3px rgba(0, 0, 0, 0.1), 0px 8px 8px 1px rgba(0, 0, 0, 0.07),
				0px 3px 8px 2px rgba(0, 0, 0, 0.08), 0px 0px 0px 2px;
		}
	}
</style>
