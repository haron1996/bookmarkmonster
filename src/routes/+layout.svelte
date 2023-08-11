<script lang="ts">
	import Header from './Header.svelte';
	import './styles.css';
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>

	<footer />
</div>

<style lang="scss">
	@import url('https://fonts.cdnfonts.com/css/google-sans');
	@import url('https://fonts.cdnfonts.com/css/arial');

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
		height: calc(100vh - calc(7vh + 5vh));
		box-sizing: border-box;
	}

	footer {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		height: 5vh;
	}
</style>
