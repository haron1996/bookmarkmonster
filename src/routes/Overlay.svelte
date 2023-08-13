<script lang="ts">
	import { hideOverlay } from '../utils/hideOverlay';
</script>

<div id="overlay" on:click|stopPropagation={hideOverlay} role="none" />

<style>
	div {
		z-index: 1;
		position: fixed;
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
		background-color: rgb(0, 0, 0, 0.4);
		opacity: 0;
		transition: opacity ease 1s;
		transform: translateX(-200%);
	}
</style>
