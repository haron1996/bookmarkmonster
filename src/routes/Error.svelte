<script lang="ts">
	import { error } from '../stores/stores';
</script>

<div class="error" id="error" class:showError={$error !== ''}>
	<i class="las la-exclamation-triangle" />
	<span>{$error}</span>
</div>

<style lang="scss">
	.error {
		position: fixed;
		top: 2%;
		left: 0;
		right: 0;
		margin: 0 auto;
		border: 0.2rem solid #f45050;
		background-color: rgb(255, 255, 255);
		width: max-content;
		padding: 1em;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 1em;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		border-radius: 0.3rem;
		transition: all ease 0.5s;
		color: red;
		transform: translateY(-200%);
		opacity: 0;

		i {
			font-size: 2rem;
		}

		span {
			font-family: 'Arial CE', sans-serif;
			font-size: 1.3rem;
		}
	}

	.showError {
		transform: translateX(0);
		opacity: 1;
	}
</style>
