<script lang="ts">
	import {
		apiHost,
		deletedTag,
		indexOfDeletedTag,
		selectedTags,
		session,
		tags
	} from '../stores/stores';
	import { closeTagMenu } from '../utils/closeTagMenu';
	import { openRenameTagModal } from '../utils/openRenameTagModal';
	import type { Tag } from '../types/tag';

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
		closeTagMenu();

		openRenameTagModal();
	}

	function selectCurrentTagName(span: HTMLSpanElement) {
		var range: Range = document.createRange();

		range.selectNodeContents(span);

		var selection: Selection | null = window.getSelection();

		if (selection === null) return;

		selection.removeAllRanges();

		selection.addRange(range);
	}

	async function handleClickOnDeleteTag() {
		const index: number = $tags.findIndex((value) => {
			return value.id === $selectedTags[0].id;
		});

		indexOfDeletedTag.set(index);

		const response = await fetch(`${$apiHost}/authenticated/tags/trashTag`, {
			method: 'PATCH',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ tag: $selectedTags[0] })
		});

		const result = await response.json();

		const dt: Tag = result[0];

		deletedTag.set(dt);

		$tags = $tags.filter((value) => {
			return value.id !== dt.id;
		});

		closeTagMenu();
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
	<div class="remove" role="none" on:click={handleClickOnDeleteTag}>
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
