<div class="comingSoon">
	<button on:click|preventDefault|stopPropagation={() => {}}>
		<span>Save your bookmarks as PDFs</span>
	</button>
	<p>Stay tuned. Coming soon</p>
</div>

<svelte:head>
	<title>Bookmark PDFS | Bookmarkmonster</title>
</svelte:head>

<style lang="scss">
	.comingSoon {
		width: 100%;
		max-width: 100%;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		gap: 2em;

		button {
			span {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
			}
		}

		p {
			font-family: 'Arial CE', sans-serif;
			font-size: 1.3rem;
		}
	}
</style>
