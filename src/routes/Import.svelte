<script lang="ts">
	import { apiHost, session, showImportData } from '../stores/stores';
	import pocketLogo from '$lib/images/pocket.svg';
	import raindropLogo from '$lib/images/raindropLogo.png';

	async function doPocketShit() {
		const response = await fetch(`${$apiHost}/authenticated/requestPocketAuthorizationCode`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		window.open(result[0], '_self');
	}
</script>

{#if $showImportData}
	<div
		class="container"
		on:click={() => {
			showImportData.set(false);
		}}
		role="none"
	>
		<div class="card" on:click|stopPropagation role="none">
			<div class="title">
				<p>Import data</p>
			</div>
			<div class="options">
				<div class="csv">
					<i class="las la-file-csv" />
					<span>Import CSV</span>
				</div>
				<div class="html">
					<i class="las la-code" />
					<span>Import HTML</span>
				</div>
				<div class="pocket">
					<img src={pocketLogo} alt="pocket" />
					<span role="none" on:click={doPocketShit}>Import from Pocket</span>
				</div>
				<div class="raindrop">
					<img src={raindropLogo} alt="raindrop" />
					<span>Import from Raindrop</span>
				</div>
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	.container {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: rgb(0, 0, 0, 0.3);
		z-index: 10;

		.card {
			min-height: 10rem;
			height: max-content;
			width: max-content;
			background-color: white;
			display: flex;
			flex-direction: column;
			gap: 2em;
			padding: 2em;
			border-radius: 0.3rem;
			box-shadow: rgba(0, 0, 0, 0.07) 0px 1px 2px, rgba(0, 0, 0, 0.07) 0px 2px 4px,
				rgba(0, 0, 0, 0.07) 0px 4px 8px, rgba(0, 0, 0, 0.07) 0px 8px 16px,
				rgba(0, 0, 0, 0.07) 0px 16px 32px, rgba(0, 0, 0, 0.07) 0px 32px 64px;

			.title {
				display: flex;
				align-items: center;

				p {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.5rem;
					color: #272829;
				}
			}

			.options {
				display: grid;
				grid-template-columns: repeat(2, 1fr);
				gap: 1em;

				div {
					display: flex;
					align-items: center;
					gap: 1em;
					padding: 1em;
					cursor: pointer;
					transition: all 300ms ease;
					border: 0.1rem solid #d8d9da;
					border-radius: 0.3rem;
					text-align: center;

					i {
						font-size: 2rem;
					}

					span {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
					}

					&:hover {
						box-shadow: rgba(9, 30, 66, 0.25) 0px 1px 1px, rgba(9, 30, 66, 0.13) 0px 0px 1px 1px;
					}

					@media only screen and (width <= 425px) {
						flex-direction: column;
					}
				}

				.pocket {
					img {
						width: 2rem;
						height: 2rem;
					}
				}

				.raindrop {
					img {
						width: 2.4rem;
						height: 2.4rem;
					}
				}

				div:not(.pocket) {
					pointer-events: none;
					opacity: 0.5;
				}
			}

			@media only screen and (width <= 600px) {
				width: 95%;
			}
		}
	}
</style>
