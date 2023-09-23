<script lang="ts">
	import { bookmarks, selectedBookmarks } from '../stores/stores';
	import type { Bookmark } from '../types/bookmark';

	function selectBookmark(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const li = t.closest('.plainBookmark') as HTMLLIElement | null;

		if (li === null) return;

		const div = document.getElementById('actionBar') as HTMLDivElement | null;

		if (div === null) return;

		const b: Bookmark = {
			added: li.dataset.added,
			bookmark: li.dataset.bookmark,
			deleted: li.dataset.deleted,
			favicon: li.dataset.favicon,
			host: li.dataset.host,
			id: li.dataset.id,
			notes: li.dataset.notes,
			thumbnail: li.dataset.thumbnail,
			title: li.dataset.title,
			updated: li.dataset.updated,
			user_id: li.dataset.userid
		};

		const exists = $selectedBookmarks
			.map((value) => {
				return value.id;
			})
			.includes(b.id);

		if (exists) {
			$selectedBookmarks = $selectedBookmarks.filter((value) => {
				return value.id !== b.id;
			});

			selectedBookmarks.set($selectedBookmarks);

			li.classList.remove('bookmarkSelected');

			div.classList.remove('animate__animated');

			div.classList.remove('animate__fadeInDown');
		} else {
			selectedBookmarks.update((values) => [b, ...values]);

			li.classList.add('bookmarkSelected');

			div.classList.add('animate__animated');

			div.classList.add('animate__fadeInDown');
		}
	}
</script>

{#each $bookmarks as { added, bookmark, deleted, favicon, host, id, notes, thumbnail, title, updated, user_id }}
	<div
		class="plainBookmark bookmark"
		id="bookmark"
		data-title={title}
		data-id={id}
		data-bookmark={bookmark}
		data-deleted={deleted}
		data-host={host}
		data-notes={notes}
		data-thumbnail={thumbnail}
		data-favicon={favicon}
		data-updated={updated}
		data-userid={user_id}
	>
		<div class="top">
			<i class="las la-slash" />
			<img
				src={`https://continuing-blush-basilisk.faviconkit.com/${host}/256`}
				alt="profile"
				draggable="false"
			/>
			<span>{host}</span>
		</div>
		<div class="titleAndTags">
			<h2>
				<a href={bookmark} target="_blank">{title}</a>
			</h2>
			<div class="tags">
				<span>untagged</span>
			</div>
		</div>
		<div class="selector" id="folderSelector" role="none" on:click|stopPropagation={selectBookmark}>
			<i class="las la-check" />
		</div>
	</div>
{/each}

<style lang="scss">
	.plainBookmark {
		display: flex;
		flex-direction: column;
		padding: 4rem;
		gap: 3em;
		border: 0.1rem solid #e7ebed;
		transition: all 300ms ease-in-out;
		word-break: break-word;
		border-radius: 0.5rem;
		width: 30rem;
		height: 25rem;
		position: relative;

		.top {
			display: flex;
			gap: 0.7em;
			align-items: center;

			i {
				font-size: 3rem;
				color: #4f709c;
				display: none;
			}

			img {
				width: 3.5rem;
				height: 3.5rem;
				object-fit: cover;
				border-radius: 100vh;
			}

			span {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.3rem;
				color: #263238;
			}
		}

		.titleAndTags {
			display: flex;
			flex-direction: column;
			gap: 2em;

			h2 {
				font-size: 1.7rem;
				line-height: 1.6;
				font-family: 'Product Sans Medium', sans-serif;

				a {
					text-decoration: none;
					color: inherit;
					display: -webkit-box;
					-webkit-line-clamp: 3;
					-webkit-box-orient: vertical;
					overflow: hidden;
					text-overflow: ellipsis;
					color: #040d12;
				}
			}

			.tags {
				display: flex;
				flex-flow: row wrap;
				gap: 1em;

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					text-transform: uppercase;
					background-color: #f4f4f2;
					padding: 0.2em;
					border-radius: 0.3rem;
					color: #546e7a;
				}
			}
		}

		.selector {
			height: 1.8rem;
			width: 1.8rem;
			position: absolute;
			top: 0.5rem;
			right: 0.5rem;
			background-color: white;
			display: flex;
			align-items: center;
			justify-content: center;
			border: 0.1rem solid #9ba4b5;
			transition: all 300ms ease;
			opacity: 0;
			border-radius: 100vh;
			cursor: default;
			cursor: pointer;

			i {
				font-size: 1.3rem;
				opacity: 0;
			}

			&:hover {
				i {
					opacity: 1;
				}
			}

			@media only screen and (width <= 768px) {
				opacity: 1;
			}
		}

		&:hover {
			box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px, rgba(17, 17, 26, 0.1) 0px 0px 8px;

			.selector {
				opacity: 1;
			}
		}

		@media only screen and (width <= 425px) {
			width: 100%;
		}

		@media only screen and (426px <= width <= 768px) {
			width: 30rem;
		}
	}

	// global
	:global(.bookmarkSelected) {
		box-shadow: rgb(4, 13, 18) 0px 0px 0px 3px;

		:global(.selector) {
			background-color: rgb(4, 13, 18) !important;
			box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
			opacity: 1 !important;
			border-color: white !important;

			:global(i) {
				opacity: 1 !important;
				color: white !important;
			}
		}

		&:hover {
			box-shadow: rgb(4, 13, 18) 0px 0px 0px 3px !important;
		}
	}
</style>
