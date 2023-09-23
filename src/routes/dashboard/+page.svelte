<script lang="ts">
	import { page } from '$app/stores';
	import { hideSideBar, loadingItems, recentUserBookmarks, session } from '../../stores/stores';

	$: $page.data.recentUserBookmarks
		? recentUserBookmarks.set($page.data.recentUserBookmarks)
		: recentUserBookmarks.set([]);
</script>

<div class="dashboard">
	<div class="top">
		<div class="left">
			<div class="salutation">
				<p>Hey,</p>
				{#if $session.User?.name}
					<p>{$session.User.name}</p>
				{:else}
					<div class="loader" />
				{/if}
			</div>
			<span>Here's your brief bookmarking overview</span>
		</div>
		<i
			class="las la-bars"
			role="none"
			on:click={() => {
				$hideSideBar = !$hideSideBar;
			}}
		/>
	</div>
	<div class="cards">
		<div class="recentlyAdded card">
			<div class="title">
				<p>Recently added</p>
			</div>
			<li class="links">
				{#if $loadingItems}
					<div class="loadingRecentBookmarks">
						<div class="loader" />
					</div>
				{:else if !$loadingItems && $recentUserBookmarks.length >= 1}
					{#each $recentUserBookmarks as { added, bookmark, deleted, favicon, host, id, notes, thumbnail, title, updated, user_id }}
						<a target="_blank" href={bookmark}>{title}</a>
					{/each}
				{:else if !$loadingItems && $recentUserBookmarks.length < 1}
					<span>Your recent bookmarks will appear here</span>
				{/if}
			</li>
		</div>
		<div class="mostVisited card">
			<div class="title">
				<p>Most visited</p>
			</div>
			<li class="links">
				<span>Coming soon</span>
			</li>
		</div>
		<div class="leastVisited card">
			<div class="title">
				<p>Least visited</p>
			</div>
			<li class="links">
				<span>Coming soon</span>
			</li>
		</div>
		<div class="suggestion card">
			<div class="title">
				<p>Suggest widget</p>
			</div>
			<div class="content">
				<span>What would you like to see on your dashboard?</span>
				<a href="/support">suggest widget</a>
			</div>
		</div>
	</div>
</div>

<style lang="scss">
	.dashboard {
		display: flex;
		flex-direction: column;
		width: 100%;
		height: 100vh;
		max-height: 100vh;
		overflow-y: auto;

		.top {
			height: 10vh;
			width: 100%;
			padding: 1em;
			display: flex;
			justify-content: space-between;

			.left {
				display: flex;
				flex-direction: column;
				gap: 0.7em;

				.salutation {
					display: flex;
					align-items: center;
					gap: 0.3em;

					p {
						font-family: 'Product Sans Medium', sans-serif;
						font-size: 2rem;
						color: #495464;
					}

					.loader {
						border: 2px solid #f3f3f3; /* Light grey */
						border-top: 2px solid #3498db; /* Blue */
						border-radius: 50%;
						width: 2rem;
						height: 2rem;
						animation: spin 0.5s linear infinite;

						@keyframes spin {
							0% {
								transform: rotate(0deg);
							}
							100% {
								transform: rotate(360deg);
							}
						}
					}
				}

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					color: #61677a;
				}
			}

			i {
				font-size: 2rem;
				cursor: pointer;
			}
		}

		.cards {
			height: 90vh;
			width: 100%;
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			gap: 2em;
			padding: 1em;
			overflow-y: auto;

			.card {
				box-shadow: rgba(9, 30, 66, 0.25) 0px 1px 1px, rgba(9, 30, 66, 0.13) 0px 0px 1px 1px;
				height: 40rem;
				border-radius: 0.3rem;

				.title {
					height: 10%;
					display: flex;
					align-items: center;
					border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);
					padding: 0.7em;

					p {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
						color: #252b48;
						font-weight: 600;
					}
				}

				.links {
					height: 90%;
					display: flex;
					flex-direction: column;
					overflow-y: auto;
					gap: 1em;
					padding: 1em;
					list-style: none;
					word-break: break-word;

					a {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						color: #61677a;
						transition: all 300ms ease-in-out;

						&:hover {
							color: red;
						}
					}

					span {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						color: #61677a;
					}

					.loadingRecentBookmarks {
						width: 100%;
						height: 100%;
						display: flex;
						flex-direction: column;
						align-items: center;
						justify-content: center;

						.loader {
							border: 2px solid #f3f3f3; /* Light grey */
							border-top: 2px solid #3498db; /* Blue */
							border-radius: 50%;
							width: 3rem;
							height: 3rem;
							animation: spin 0.5s linear infinite;

							@keyframes spin {
								0% {
									transform: rotate(0deg);
								}
								100% {
									transform: rotate(360deg);
								}
							}
						}
					}
				}

				&.mostVisited,
				&.leastVisited {
					.links {
						span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							color: #61677a;
						}
					}
				}

				&.suggestion {
					.content {
						display: flex;
						flex-direction: column;
						padding: 1em;
						gap: 1em;

						span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							color: #61677a;
						}

						a {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							color: #0079ff;
						}
					}
				}
			}

			@media only screen and (width <= 600px) {
				grid-template-columns: repeat(1, 1fr);
			}

			@media only screen and (601px <= width <= 1200px) {
				grid-template-columns: repeat(2, 1fr);
			}
		}
	}
</style>
