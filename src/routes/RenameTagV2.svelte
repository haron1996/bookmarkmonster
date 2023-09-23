<script lang="ts">
	import { apiHost, selectedTag, session, tags } from '../stores/stores';
	import type { Tag } from '../types/tag';

	function handleClickOnClose() {
		selectedTag.set({});

		const form = document.getElementById('renameTagV2') as HTMLFormElement | null;

		if (form === null) return;

		form.style.display = 'none';

		const tags = document.getElementById('tags') as HTMLDivElement | null;

		if (tags === null) return;

		tags.style.overflowY = 'auto';
	}

	async function handleRenameTagFormSubmit() {
		if ($selectedTag.name === '' || $selectedTag.name === undefined) {
			alert('tag name is required');
			return;
		}

		const response = await fetch(`${$apiHost}/authenticated/tags/renameTag`, {
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
			body: JSON.stringify({ tag: $selectedTag })
		});

		const result = await response.json();

		const t: Tag = result[0];

		const index: number = $tags.findIndex((value) => {
			return value.id === t.id;
		});

		$tags.splice(index, 1, t);

		tags.set($tags);

		const form = document.getElementById('renameTagV2') as HTMLFormElement | null;

		if (form === null) return;

		form.style.display = 'none';

		selectedTag.set({});
	}
</script>

<form id="renameTagV2" on:submit|preventDefault={handleRenameTagFormSubmit}>
	<input type="text" name="renameTag" id="renameTag" bind:value={$selectedTag.name} />
	<div class="close">
		<i class="las la-times" role="none" on:click|stopPropagation={handleClickOnClose} />
	</div>
</form>

<style lang="scss">
	form {
		display: flex;
		align-items: center;
		background-color: inherit;
		position: absolute;
		z-index: 10;
		background-color: white;
		border-bottom: 0.1rem solid rgb(4, 13, 18, 0.1);
		display: none;

		input {
			width: 85%;
			height: 4rem;
			padding: 0em 0.5em;
			font-family: 'Arial CE', sans-serif;
			font-size: 1.3rem;
			background-color: inherit;
			outline: none;
			border: none;
			transition: all 300ms ease-in-out;
			text-transform: lowercase;
		}

		.close {
			height: 4rem;
			width: 15%;
			display: flex;
			align-items: center;
			justify-content: center;
			background-color: inherit;

			i {
				font-size: 1.5rem;
				cursor: pointer;
			}
		}
	}
</style>
