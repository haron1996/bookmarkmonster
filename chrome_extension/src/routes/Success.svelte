<script lang="ts">
	import check from '../../static/icons8-check.png';
	import { alertMessage } from '../../store/stores';

	$: $alertMessage,
		$alertMessage !== ''
			? setTimeout(() => {
					alertMessage.set('');
			  }, 3000)
			: () => {};
</script>

{#if $alertMessage !== ''}
	<div class="success">
		<img src={check} alt="check" />
		<span>{$alertMessage}</span>
		<small
			role="none"
			on:click={() => {
				alertMessage.set('');
			}}>ok</small
		>
	</div>
{/if}

<style>
	.success {
		display: flex;
		align-items: center;
		width: max-content;
		background-color: #272829;
		height: 4rem;
		padding: 0.5em 1em;
		z-index: 1;
		position: fixed;
		right: 0;
		bottom: 0;
		gap: 1em;
		box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px, #1a9047 0px 0px 0px 2px;

		& img {
			height: 2.4rem;
			width: 2.4rem;
			object-fit: contain;
		}

		& span {
			font-family: 'Arial CE', sans-serif;
			font-size: 1.3rem;
			color: white;
		}

		& small {
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			color: white;
			border: 0.1rem dotted white;
			border-radius: 0.3rem;
			cursor: pointer;
		}
	}
</style>
