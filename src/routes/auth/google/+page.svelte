<script lang="ts">
	import { afterNavigate, goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { apiHost } from '../../../stores/stores';

	let code: string | null;

	afterNavigate(async () => {
		code = $page.url.searchParams.get('code');

		await exchangeCodeForToken();
	});

	const exchangeCodeForToken = async () => {
		if (code) {
			const response = await fetch(`${$apiHost}/auth/register-google-user`, {
				method: 'POST',
				mode: 'cors',
				cache: 'no-cache',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json'
				},
				redirect: 'follow',
				referrerPolicy: 'no-referrer',
				body: JSON.stringify({ code: code })
			});

			const result = await response.json();

			const session = result[0];

			localStorage.setItem('session', JSON.stringify(session));

			if ($page.url.origin === 'http://localhost:5173') {
				goto('http://localhost:5173/dashboard');
			} else {
				goto('https://bookmarkmonster.xyz/dashboard');
			}
		}
	};
</script>

<svelte:head>
	<title>Continue with Google | BookmarkMonster</title>
</svelte:head>

<div class="container">
	<div class="loader" />
</div>

<style lang="scss">
	.container {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-color: rgb(255, 255, 255);
		display: flex;
		align-items: center;
		justify-content: center;

		.loader {
			border: 0.2rem solid #f3f3f3;
			border-radius: 50%;
			border-top: 0.2rem solid #3498db;
			width: 3rem;
			height: 3rem;
			-webkit-animation: spin 2s linear infinite; /* Safari */
			animation: spin 0.5s linear infinite;

			/* Safari */
			@-webkit-keyframes spin {
				0% {
					-webkit-transform: rotate(0deg);
				}
				100% {
					-webkit-transform: rotate(360deg);
				}
			}

			@keyframes spin {
				0% {
					transform: rotate(0deg);
				}
				100% {
					transform: rotate(360deg);
				}
			}
		}
	}
</style>
