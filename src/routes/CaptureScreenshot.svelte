<script lang="ts">
	import {
		apiHost,
		screenshot,
		selectedBookmarks,
		selectedFolders,
		selectedItems,
		session,
		showCaptureScreeshot
	} from '../stores/stores';
	import { removeSelectedClassFromAllDomBookmarks } from '../utils/removeSelectedClassFromAllSelectedDomBookmarks';
	import { removeSlideInAnimationFromActionBarV2 } from '../utils/removeSlideInAnimationFromActionBarV2';

	let gettingScreenshot: boolean = false;

	async function captureFullpageScreenshot() {
		gettingScreenshot = true;

		if (
			$selectedItems.length > 1 ||
			$selectedBookmarks.length < 1 ||
			$selectedFolders.length >= 1
		) {
			gettingScreenshot = false;
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/bookmarks/captureFullpageScreeshot`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ bookmark_id: $selectedBookmarks[0].id })
		});

		const result = await response.json();

		const msg = result.msg;

		if (msg) {
			alert(msg);
			gettingScreenshot = false;
			return;
		}

		screenshot.set(result[0]);

		gettingScreenshot = false;
	}

	async function captureAboveTheFoldScreenshot() {
		gettingScreenshot = true;

		if (
			$selectedItems.length > 1 ||
			$selectedBookmarks.length < 1 ||
			$selectedFolders.length >= 1
		) {
			gettingScreenshot = false;
			return;
		}

		const response = await fetch(
			`${$apiHost}/authenticated/bookmarks/captureAboveTheFoldScreenshot`,
			{
				method: 'POST',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${$session.AccessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer',
				body: JSON.stringify({ bookmark_id: $selectedBookmarks[0].id })
			}
		);

		const result = await response.json();

		const msg = result.msg;

		if (msg) {
			alert(msg);
			gettingScreenshot = false;
			return;
		}

		screenshot.set(result[0]);

		gettingScreenshot = false;
	}

	async function downloadScreenshot() {
		if ($screenshot.screenshot_location === undefined) return;

		window.open($screenshot.screenshot_location, '_blank');

		screenshot.set({});

		showCaptureScreeshot.set(false);

		selectedBookmarks.set([]);

		removeSelectedClassFromAllDomBookmarks();

		removeSlideInAnimationFromActionBarV2();
	}
</script>

{#if $showCaptureScreeshot}
	<div
		class="wrapper"
		role="none"
		on:click={() => {
			showCaptureScreeshot.set(false);
		}}
	>
		<div class="card" role="none" on:click|stopPropagation={() => {}}>
			<div class="title">
				<p>Capture screenshot</p>
			</div>
			<div class="options">
				{#if gettingScreenshot}
					<div class="capturing">
						<div class="loader" />
						<p>Capturing screenshot...</p>
					</div>
				{:else if !gettingScreenshot && $screenshot.screenshot_location !== undefined}
					<div class="download">
						<button on:click|preventDefault={downloadScreenshot}> Download screenshot </button>
					</div>
				{:else if !gettingScreenshot && $screenshot.screenshot_location === undefined}
					<span role="none" on:click={captureFullpageScreenshot}>Capture full page</span>
					<span role="none" on:click={captureAboveTheFoldScreenshot}>Capture above the fold</span>
					<span>Capture element screenshot</span>
				{/if}
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	.wrapper {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-color: rgb(0, 0, 0, 0.4);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 10;

		.card {
			background-color: white;
			border-radius: 0.3rem;
			height: max-content;
			min-height: 10rem;
			min-width: 40rem;
			display: flex;
			flex-direction: column;
			padding: 2em;
			gap: 2em;
			box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;

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

				.capturing {
					min-height: inherit;
					min-width: inherit;
					display: flex;
					align-items: center;
					justify-content: center;
					gap: 1em;

					.loader {
						border: 0.2rem solid #f3f3f3;
						border-top: 0.2rem solid #040d12;
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

					p {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
					}
				}

				.download {
					min-height: inherit;
					min-width: inherit;
					display: flex;
					align-items: center;
					justify-content: center;

					button {
						border: none;
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						padding: 0.7em 1em;
						cursor: pointer;
						transition: all 300ms ease;

						&:hover {
							box-shadow: rgba(9, 30, 66, 0.25) 0px 1px 1px, rgba(9, 30, 66, 0.13) 0px 0px 1px 1px;
						}
					}
				}

				span {
					padding: 0.7em 1em;
					border: 0.1rem solid #d8d9da;
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					cursor: pointer;
					border-radius: 0.2rem;
					transition: all 300ms ease;

					&:hover {
						box-shadow: rgba(9, 30, 66, 0.25) 0px 1px 1px, rgba(9, 30, 66, 0.13) 0px 0px 1px 1px;
					}
				}
			}
		}
	}
</style>
