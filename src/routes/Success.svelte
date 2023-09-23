<script lang="ts">
	import checkIcon from '$lib/images/check.png';
	import cancelIcon from '$lib/images/cancel.png';
	import { successMessage } from '../stores/stores';
</script>

{#if $successMessage !== ''}
	<div class="wrapper">
		<div class="success animate__animated animate__slideInUp">
			<img src={checkIcon} alt="check icon" />
			<span>{$successMessage}</span>
			<img src={cancelIcon} alt="cancel icon" />
		</div>
	</div>
{/if}

<style lang="scss">
	.wrapper {
		position: fixed;
		left: 0;
		top: 0;
		z-index: 10;
		width: 100vw;
		height: 100vh;
		background-color: transparent;
		display: flex;
		align-items: flex-end;
		justify-content: center;
		padding: 1em;

		.success {
			display: flex;
			align-items: center;
			gap: 1em;
			background-color: #ddffbb;
			width: max-content;
			padding: 1em;
			border: 0.2rem solid green;
			border-radius: 0.3rem;
			box-shadow: rgba(0, 0, 0, 0.15) 0px 5px 15px 0px;

			span {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
			}
		}
	}
</style>
