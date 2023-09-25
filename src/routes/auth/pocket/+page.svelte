<script lang="ts">
	import { onMount } from 'svelte';
	import { apiHost, errorMessage, session, successMessage } from '../../../stores/stores';
	import { goto } from '$app/navigation';
	import loadingAnimation from '$lib/images/loading_animation.gif';

	onMount(() => {
		convertPocketRequestTokenToAccessToken();
	});

	async function convertPocketRequestTokenToAccessToken() {
		const response = await fetch(
			`${$apiHost}/authenticated/convertPocketRequestTokenToAccessToken`,
			{
				method: 'POST',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json',
					authorization: `Bearer${$session.AccessToken}`
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer'
			}
		);

		const result = await response.json();

		const message = result.message;

		if (message) {
			errorMessage.set(message);
			return;
		}

		const pocketToken = result[0];

		if (pocketToken) {
			importFromPocket(pocketToken.access_token, pocketToken.username);
		}
	}

	async function importFromPocket(pocket_access_token: string, pocket_username: string) {
		const response = await fetch(`${$apiHost}/authenticated/importFromPocket`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${$session.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				pocket_access_token: pocket_access_token,
				pocket_username: pocket_username
			})
		});

		const result = await response.json();

		if (result.message) {
			errorMessage.set(result.message);
			return;
		}

		successMessage.set('Bookmarks imported successfully');

		setTimeout(() => {
			goto('/dashboard/tags');
		}, 6100);
	}
</script>

<div class="container">
	<div class="loading">
		<span>Importing from Pocket....</span>
		<img src={loadingAnimation} alt="loading" />
	</div>
</div>

<style lang="scss">
	.container {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: white;

		.loading {
			display: flex;
			flex-direction: column;
			gap: 2em;
			align-items: center;
			justify-content: center;
			text-align: center;

			span {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.5rem;
			}
		}
	}
</style>
