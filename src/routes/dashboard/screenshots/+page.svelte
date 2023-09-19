<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { apiHost, screenshots, session } from '../../../stores/stores';

	let showFullPageScreenshots: boolean = false;
	let showAboveFoldScreenshots: boolean = false;
	let showElementScreenshots: boolean = false;

	afterNavigate(async () => {
		showFullPageScreenshots = true;
	});

	async function getFullPageScreenshots() {
		const promise = await fetch(`${$apiHost}/authenticated/bookmarks/getFullpageScreenshots`, {
			method: 'get',
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

		const result = await promise.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			return;
		}

		screenshots.set(result[0]);
	}

	async function getAboveFoldScreenshots() {
		const promise = await fetch(`${$apiHost}/authenticated/bookmarks/getAboveFoldScreenshots`, {
			method: 'get',
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

		const result = await promise.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			return;
		}

		screenshots.set(result[0]);
	}

	async function getElementScreenshots() {
		console.log('get element screenshots');
		screenshots.set([]);
	}

	function toggleOfAllBooleans() {
		showFullPageScreenshots = false;
		showAboveFoldScreenshots = false;
		showElementScreenshots = false;
	}

	$: showFullPageScreenshots ? getFullPageScreenshots() : () => {};
	$: showAboveFoldScreenshots ? getAboveFoldScreenshots() : () => {};
	$: showElementScreenshots ? getElementScreenshots() : () => {};
</script>

<div class="container">
	<div class="switch">
		<p
			class:active={showFullPageScreenshots}
			role="none"
			on:click={() => {
				toggleOfAllBooleans();
				showFullPageScreenshots = true;
			}}
		>
			Full page screenshots
		</p>
		<p
			class:active={showAboveFoldScreenshots}
			role="none"
			on:click={() => {
				toggleOfAllBooleans();
				showAboveFoldScreenshots = true;
			}}
		>
			Above the fold screenshots
		</p>
		<p
			class:active={showElementScreenshots}
			role="none"
			on:click={() => {
				toggleOfAllBooleans();
				showElementScreenshots = true;
			}}
		>
			Element screenshots
		</p>
	</div>
	<div class="screenshots">
		{#each $screenshots as { bookmark_id, created, deleted, fullpage, id, screenshot_location, user_id }}
			<div class="screenshot">
				<img src={screenshot_location} alt="screenshot" loading="lazy" />
			</div>
		{/each}
	</div>
</div>

<svelte:head>
	<title>Bookmark Screenshots | Bookmarkmonster</title>
</svelte:head>

<style lang="scss">
	.container {
		width: 100%;
		max-width: 100%;
		height: 100vh;
		display: flex;
		flex-direction: column;
		background-color: white;

		.switch {
			display: flex;
			align-items: center;
			padding: 1em;
			justify-content: start;
			gap: 3em;
			min-height: 5rem;
			background-color: #eeeeee;

			p {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.3rem;
				cursor: pointer;
				text-decoration-line: underline;
				text-underline-offset: 0.5em;
				text-decoration-color: transparent;
				text-decoration-thickness: 0.2rem;
			}

			.active {
				text-decoration-color: #040d12;
			}
		}

		.screenshots {
			height: calc(calc(100vh - 5vh) - 1em);
			width: 100%;
			overflow-y: auto;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			gap: 1em;
			padding: 1em;

			.screenshot {
				height: 30rem;
				width: 30rem;
				border-radius: 0.3rem;
				border: 0.2rem solid #040d12;

				img {
					width: 100%;
					height: 100%;
					max-inline-size: 100%;
					max-width: 100%;
					max-height: 100%;
					object-fit: cover;
					border-radius: inherit;
				}
			}
		}
	}
</style>
