<script lang="ts">
	import { apiHost, deletedTag, indexOfDeletedTag, session, tags } from '../stores/stores';
	import type { Tag } from '../types/tag';

	async function handleClickOnUndo() {
		const response = await fetch(`${$apiHost}/authenticated/tags/restoreTag`, {
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
			body: JSON.stringify({ tag: $deletedTag })
		});

		const result = await response.json();

		const rt: Tag = result[0];

		if ($indexOfDeletedTag !== null) {
			$tags.splice($indexOfDeletedTag, 0, rt);

			tags.set($tags);
		}

		deletedTag.set({});

		indexOfDeletedTag.set(null);

		setTimeout(() => {
			const tagNodes = document.querySelectorAll(
				'.sideBarTag'
			) as NodeListOf<HTMLDivElement> | null;

			if (tagNodes === null) return;

			tagNodes.forEach((tn) => {
				if (tn.dataset.id === rt.id) {
					tn.click();
				}
			});
		}, 50);
	}

	$: $deletedTag,
		$deletedTag.name !== undefined
			? setTimeout(() => {
					deletedTag.set({});
			  }, 4000)
			: () => {};

	$: $indexOfDeletedTag,
		$indexOfDeletedTag !== null
			? setTimeout(() => {
					indexOfDeletedTag.set(null);
			  }, 10000)
			: () => {};
</script>

{#if $deletedTag.name}
	<div class="container" id="tagDeleted">
		<span>{$deletedTag.name} removed from your tags</span>
		<span class="undo" role="none" on:click={handleClickOnUndo}>undo</span>
	</div>
{/if}

<style lang="scss">
	.container {
		position: fixed;
		bottom: 2%;
		right: 0;
		left: 0;
		margin-right: auto;
		margin-left: auto;
		max-width: max-content;
		z-index: 6;
		padding: 1.5em;
		display: flex;
		align-items: center;
		gap: 1em;
		border-radius: 0.3rem;
		background-color: #040d12;
		box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em, rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em,
			rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
		transition: all ease 500ms;

		span {
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			color: rgb(253, 253, 253);
		}

		span.undo {
			text-transform: uppercase;
			cursor: pointer;
		}
	}
</style>
