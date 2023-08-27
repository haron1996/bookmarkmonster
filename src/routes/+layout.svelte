<script lang="ts">
	import 'animate.css';
	import Error from './Error.svelte';
	import Header from './Header.svelte';
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>

	<footer />

	<Error />
</div>

<style lang="scss" global>
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@600;700;800;900&family=Roboto:wght@100;300;400;500;700;900&display=swap');
	@import url('https://fonts.cdnfonts.com/css/google-sans');
	@import url('https://fonts.cdnfonts.com/css/arial');
	@import url('https://fonts.cdnfonts.com/css/segoe-ui-4');

	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
		font-family: 'Arial CE', sans-serif;
	}

	main {
		display: flex;
		flex-direction: column;
		width: 100%;
		min-height: calc(100vh - 7vh);
		height: auto;
		box-sizing: border-box;
		overflow-y: auto;
		font-family: 'Arial CE', sans-serif;
	}

	// footer {
	// 	display: flex;
	// 	flex-direction: column;
	// 	justify-content: center;
	// 	align-items: center;
	// 	height: 5vh;
	// }
</style>
