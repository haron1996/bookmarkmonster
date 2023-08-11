<script>
	import { page } from '$app/stores';
	import logo from '$lib/images/svelte-logo.svg';
	import github from '$lib/images/github.svg';
</script>

<header>
	<div class="logo" />

	<nav />
</header>

<style lang="scss">
	header {
		display: flex;
		justify-content: space-between;
		height: 7vh;

		.logo {
			width: 10%;
		}

		nav {
			display: flex;
			justify-content: center;
			height: inherit;
			width: 90%;
		}
	}
</style>
