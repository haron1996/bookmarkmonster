<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let code: string | null;

	onMount(async () => {
		code = $page.url.searchParams.get('code');

		await exchangeCodeForToken();
	});

	const exchangeCodeForToken = async () => {
		if (code) {
			const response = await fetch(`http://localhost:5000/auth/register-google-user`, {
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

			goto('http://localhost:5173/dashboard');
		}
	};
</script>
