<script lang="ts">
	import { toolTipText } from '../stores/stores';
</script>

<div id="toolTip">
	<span>{$toolTipText}</span>
</div>

<style lang="scss">
	div {
		width: max-content;
		padding: 0.5em 1em;
		background-color: #0f2c59;
		position: fixed;
		z-index: 10;
		cursor: pointer;
		border-radius: 0.3rem;
		box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em, rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em,
			rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
		visibility: hidden;
		transition: visibility 300ms ease;

		span {
			font-size: 1.3rem;
			color: white;
			font-family: 'Arial CE', sans-serif;
		}
	}
</style>
