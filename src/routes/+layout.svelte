<script lang="ts">
	import 'animate.css';
	import Error from './Error.svelte';
	import Success from './Success.svelte';
	import { errorMessage, successMessage } from '../stores/stores';
	import ErrorV2 from './ErrorV2.svelte';

	$: $successMessage !== ''
		? setTimeout(() => {
				successMessage.set('');
		  }, 6000)
		: () => {};
	$: $errorMessage !== ''
		? setTimeout(() => {
				errorMessage.set('');
		  }, 6000)
		: () => {};
</script>

<div class="app">
	<!-- <Header /> -->

	<main>
		<slot />
	</main>

	<!-- <Footer /> -->

	<Error />
	<Success />
	<ErrorV2 />
</div>

<style lang="scss" global>
	@import url('https://fonts.googleapis.com/css2?family=Lora:wght@700&family=Montserrat:wght@300;400;500;600;700;800;900&family=Poppins:wght@100;200;300;400;500;700;800;900&family=Roboto:wght@300;400;500;700;900&display=swap');
	@import url('https://fonts.cdnfonts.com/css/google-sans');
	@import url('https://fonts.cdnfonts.com/css/arial');
	@import url('https://fonts.cdnfonts.com/css/segoe-ui-4');

	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		display: flex;
		flex-direction: column;
		width: 100%;
		min-height: calc(100vh - 7vh);
		height: auto;
		box-sizing: border-box;
		overflow-y: auto;
	}
</style>
