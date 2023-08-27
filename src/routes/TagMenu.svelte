<script lang="ts">
	import { selectedTags } from '../stores/stores';
	import { closeTagMenu } from '../utils/closeTagMenu';

	function handleTagMenuMouseOver() {
		const menu = document.getElementById('tagMenu') as HTMLElement | null;

		if (menu === null) return;

		menu.style.display = 'flex';
	}

	function handleTagMenuMouseLeave() {
		const menu = document.getElementById('tagMenu') as HTMLElement | null;

		if (menu === null) return;

		menu.style.display = 'none';
	}

	function handleClickOnRenameTag() {
		const bookmarkNodes = document.querySelectorAll('.tag') as NodeListOf<HTMLDivElement> | null;
		if (bookmarkNodes === null) return;
		bookmarkNodes.forEach((bm) => {
			if (bm.dataset.id === $selectedTags[0].id) {
				const span = bm.childNodes[2] as HTMLSpanElement;
				span.contentEditable = 'true';
				selectCurrentTagName(span);
			}
		});

		closeTagMenu();
	}

	function selectCurrentTagName(span: HTMLSpanElement) {
		var range: Range = document.createRange();

		range.selectNodeContents(span);

		var selection: Selection | null = window.getSelection();

		if (selection === null) return;

		selection.removeAllRanges();

		selection.addRange(range);
	}
</script>

<div
	class="wrapper"
	id="tagMenu"
	on:mouseover={handleTagMenuMouseOver}
	on:focus={() => {}}
	on:mouseleave={handleTagMenuMouseLeave}
	role="none"
>
	<div class="rename" role="none" on:click={handleClickOnRenameTag}>
		<i class="las la-pen" />
		<span>Rename</span>
	</div>
	<div class="remove">
		<i class="las la-trash-alt" />
		<span>Delete</span>
	</div>
</div>

<style lang="scss">
	.wrapper {
		display: flex;
		flex-direction: column;
		position: absolute;
		z-index: 2;
		background-color: rgb(248, 246, 244);
		min-width: 20rem;
		border-radius: 0.1rem;
		box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
		display: none;

		.rename,
		.remove {
			min-height: 4rem;
			display: flex;
			align-items: center;
			gap: 1em;
			border-bottom: 0.1rem solid rgb(2, 84, 100, 0.1);
			padding: 0.5em;
			cursor: pointer;

			i {
				font-size: 1.8rem;
			}

			span {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
			}
		}
	}
</style>
