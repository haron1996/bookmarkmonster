<script lang="ts">
	import { sideBarWidth } from '../../stores/stores';
	import profilePic from '$lib/images/profile-pic.jpg';
	import bookmarkThumbnail from '$lib/images/bookmark-thumbnail.png';
	import favicon from '$lib/images/faviconV2.png';

	let sideBarWidthFromStore: number;

	const tags: string[] = [
		'illustration',
		'school',
		'work',
		'marketing',
		'diet',
		'health',
		'maths',
		'physics',
		'history',
		'climate',
		'animals',
		'cities'
	];

	const handleClickOnTag = (e: MouseEvent) => {
		const clickedEl = e.target as HTMLElement;

		const clickedTag = clickedEl.closest('.tag') as HTMLSpanElement | null;

		if (clickedTag === null) return;

		const tags = document.querySelectorAll('.tag') as NodeListOf<HTMLSpanElement>;

		tags.forEach((tag) => {
			if (tag.classList.contains('active-tag')) {
				tag.classList.remove('active-tag');
			}
		});

		clickedTag.classList.add('active-tag');
	};

	const handleClickOnCreateTag = (e: MouseEvent) => {
		console.log('add tag');
	};

	$: sideBarWidthFromStore = $sideBarWidth;
</script>

<div class="app">
	<div class="sidebar">
		<div class="profile">
			<div class="profile-pic">
				<img src={profilePic} alt="profile" />
			</div>
			<div class="name-and-email">
				<h3>Kwandap Kipchumba</h3>
				<span>kwandapchumba@gmail.com</span>
			</div>
		</div>
		<div class="tags">
			<div
				class="tag active-tag"
				id="tag"
				on:click|stopPropagation={handleClickOnTag}
				on:keyup
				role="none"
			>
				<i class="las la-hashtag" />
				<span>all tags</span>
			</div>
			{#each tags as tag}
				<div class="tag" id="tag" on:click|stopPropagation={handleClickOnTag} on:keyup role="none">
					<i class="las la-hashtag" />
					<span>{tag}</span>
				</div>
			{/each}
		</div>
		<div class="create-tag" on:click|stopPropagation={handleClickOnCreateTag} role="none">
			<div class="new-tag">
				<i class="las la-plus" />
				<span>create tag</span>
			</div>
		</div>
	</div>
	<div class="main">
		<div class="search-or-add">
			<input
				type="search"
				name="search"
				id="search"
				placeholder="Type to search a bookmark..."
				autocomplete="off"
			/>
			<button>
				<span>add bookmark</span>
			</button>
		</div>
		<div class="bookmarks">
			<div class="bookmark">
				<div class="thumbnail">
					<img src={bookmarkThumbnail} alt="bookmark thumbnail" />
				</div>
				<div class="title-favicon-and-domain">
					<div class="title">
						<a href="/">
							Lorem ipsum dolor sit amet consectetur adipisicing elit. Officiis ab quod molestiae
							voluptas veniam asperiores ipsam, assumenda facere mollitia aut explicabo dignissimos
							repudiandae magni beatae eius labore minus facilis quae.
						</a>
					</div>
					<div class="favicon-and-domain">
						<img src={favicon} alt="bookmark favicon" />
						<span>google.com</span>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<style lang="scss">
	.app {
		height: 100vh;
		width: 100vw;
		position: fixed;
		top: 0;
		left: 0;
		display: flex;

		.sidebar {
			width: 20vw;
			height: 100%;
			background-color: rgb(255, 255, 255);
			display: flex;
			flex-direction: column;
			border-right: 0.1rem solid rgb(0, 0, 0, 0.1);

			.profile {
				width: 100%;
				height: 10%;
				background-color: rgb(255, 255, 255);
				display: flex;
				align-items: center;
				gap: 0.5em;
				padding: 0.5em;

				img {
					width: 3.5rem;
					height: 3.5rem;
					object-fit: cover;
					border-radius: 100vh;
				}

				.name-and-email {
					display: flex;
					flex-direction: column;
					gap: 0.3em;
					white-space: nowrap;
					overflow: hidden;
					text-overflow: ellipsis;

					h3 {
						font-size: 1.3rem;
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
					}

					span {
						font-size: 1.3rem;
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
					}
				}
			}

			.tags {
				width: 100%;
				height: 85%;
				background-color: inherit;
				overflow-y: auto;
				padding-top: 1em;
				display: flex;
				flex-direction: column;

				.tag {
					width: 100%;
					padding: 0.5em;
					gap: 0.5em;
					cursor: pointer;
					display: flex;
					align-items: center;

					i {
						font-size: 2rem;
					}

					span {
						font-size: 1.3rem;
					}
				}
			}

			.create-tag {
				width: 100%;
				height: 5%;
				border-top: 0.1rem solid rgb(0, 0, 0, 0.1);
				cursor: pointer;
				background-color: rgb(255, 255, 255);

				.new-tag {
					height: 100%;
					width: 100%;
					padding: 0 0.5em;
					display: flex;
					align-items: center;
					gap: 0.5em;
					transition: background-color ease 0.5s;

					i {
						font-size: 2rem;
					}

					span {
						font-size: 1.3rem;
					}
				}

				&:hover {
					background-color: rgb(238, 238, 238);
				}
			}
		}

		.main {
			width: calc(100vw - 20vw);
			height: 100%;
			background-color: rgb(255, 255, 255);

			.search-or-add {
				display: flex;
				height: 7vh;
				border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
				display: flex;
				align-items: center;
				justify-content: space-between;
				padding: 0.5em;

				input[type='search'] {
					width: 50%;
					padding: 0.5em;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					outline: none;
					border-radius: 0.3rem;
					font-size: 1.3rem;

					&:hover {
						border-color: rgb(47, 88, 205);
					}
				}

				button {
					padding: 0.5em;
					border: 0.1rem solid rgb(47, 88, 205);
					outline: none;
					border-radius: 0.3rem;
					display: flex;
					align-items: center;
					cursor: pointer;
					gap: 0.3em;
					background-color: transparent;

					span {
						font-size: 1.3rem;
						color: rgb(47, 88, 205);
						text-transform: capitalize;
					}

					&:hover {
						background-color: rgb(47, 88, 205);

						span {
							color: rgb(255, 255, 255);
						}
					}
				}
			}

			.bookmarks {
				height: calc(100vh - 7vh);
				padding: 0.5em;

				.bookmark {
					height: 30rem;
					width: 30rem;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					display: flex;
					flex-direction: column;
					border-radius: 0.6rem;
					gap: 1em;
					box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
					padding: 0.5em;

					.thumbnail {
						height: 70%;
						width: 100%;
						display: flex;
						align-items: center;
						justify-content: center;

						img {
							height: 100%;
							width: 100%;
							max-inline-size: 100%;
							object-fit: cover;
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
								font-weight: 600;

								&:hover {
									color: rgb(78, 79, 235);
									text-decoration-color: rgb(78, 79, 235);
								}
							}
						}

						.favicon-and-domain {
							width: 100%;
							height: 50%;
							display: flex;
							align-items: center;
							justify-content: space-between;

							img {
								height: 2.5rem;
								width: 2.5rem;
								object-fit: cover;
								border-radius: 100vh;
							}

							span {
								color: rgb(24, 23, 40);
								font-size: 1.3rem;
								white-space: nowrap;
								overflow: hidden;
								text-overflow: ellipsis;
								max-width: 30%;
							}
						}
					}

					&:hover {
						border-color: rgb(78, 79, 235);
					}
				}
			}
		}
	}

	// GLOBAL STYLES
	:global(.active-tag) {
		i {
			color: red !important;
		}

		span {
			color: red !important;
		}
	}
</style>
