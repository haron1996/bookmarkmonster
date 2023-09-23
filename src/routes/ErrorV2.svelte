<script lang="ts">
	import cancelCircleRedIcon from '$lib/images/cancelCircleRed.png';
	import cancelRedIcon from '$lib/images/cancelRed.png';
	import { errorMessage } from '../stores/stores';
</script>

{#if $errorMessage !== ''}
	<div class="wrapper">
		<div class="error animate__animated animate__slideInUp">
			<img src={cancelCircleRedIcon} alt="check icon" />
			<span>{$errorMessage}</span>
			<img src={cancelRedIcon} alt="cancel icon" />
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

		.error {
			display: flex;
			align-items: center;
			gap: 1em;
			background-color: #ffeadd;
			width: max-content;
			padding: 1em;
			border: 0.2rem solid red;
			border-radius: 0.3rem;
			box-shadow: rgba(0, 0, 0, 0.15) 0px 5px 15px 0px;

			span {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
			}
		}
	}
</style>
