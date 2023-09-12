<script lang="ts">
	import { page } from '$app/stores';
	import { loadingItems, recentUserBookmarks, session } from '../../stores/stores';

	// let links = [
	// 	'Learning Sales Funnel Conversion Tactics From 3 Print Ads',
	// 	'Launching A High Converting Sales Funnel Out Of The Box',
	// 	'How To Validate A Market For Your Next Sales Funnel',
	// 	'How To Take The Training Wheels Off Your Sales Funnel',
	// 	'7 Classic Marketing Books You Must Read Before Funnel Hacking',
	// 	'Marketer’s Guide to Including Social Proof in Sales Funnel',
	// 	'How to Advertise on Google Search: Leveraging the Traffic You Control',
	// 	'Huge Funnel Hack Exposed! Manscaped 6 Step Recipe To Success {Steal This Formula!}',
	// 	'Traffic Triage: How to Save Your Ad Money!',
	// 	'How to Create Successful Email Marketing Campaigns',
	// 	'5 Reasons To Replace Your Stock Standard Website With A Sales Funnel',
	// 	'Copywriting Hacks in the ClickFunnels Editor to Increase Conversions',
	// 	'The fastest and easiest way to check and copy CSS'
	// ];

	$: $page.data.recentUserBookmarks
		? recentUserBookmarks.set($page.data.recentUserBookmarks)
		: recentUserBookmarks.set([]);
</script>

<div class="dashboard">
	<div class="top">
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
			// display: flex;
			// align-items: center;
			padding: 1em;
			gap: 1em;
			display: flex;
			flex-direction: column;

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

			// p {
			// 	font-family: 'Product Sans Medium', sans-serif;
			// 	font-size: 2rem;
			// 	color: #495464;
			// }

			// .loader {
			// 	border: 2px solid #f3f3f3; /* Light grey */
			// 	border-top: 2px solid #3498db; /* Blue */
			// 	border-radius: 50%;
			// 	width: 2rem;
			// 	height: 2rem;
			// 	animation: spin 0.5s linear infinite;

			// 	@keyframes spin {
			// 		0% {
			// 			transform: rotate(0deg);
			// 		}
			// 		100% {
			// 			transform: rotate(360deg);
			// 		}
			// 	}
			// }

			span {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
				color: #61677a;
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
		}
	}
</style>
