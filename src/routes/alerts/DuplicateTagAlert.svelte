<div class="container" id="duplicateTag">
	<span>Duplicate tag!</span>
</div>

<style lang="scss">
	.container {
		position: fixed;
		top: 2%;
		transform: translateX(-50%);
		left: 50%;
		min-width: max-content;
		z-index: 2;
		padding: 0 1em;
		min-height: 3.5rem;
		display: flex;
		align-items: center;
		gap: 1em;
		border-radius: 0.3rem;
		background-color: rgb(255, 102, 102);
		box-shadow: rgba(0, 0, 0, 0.15) 0px 5px 15px 0px;
		transition: all linear 300ms;
		opacity: 0;
		transform: translateY(-100%);

		span {
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
		}
	}
</style>
