<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { apiHost } from '../../../stores/stores';
	import { goto } from '$app/navigation';

	let email: string = '';
	let password: string = '';
	let token: string = '';

	onMount(() => {
		const t: string | null = $page.url.searchParams.get('token');

		alert(t);

		if (t === null) {
			goto('/signup');
			return;
		}

		token = t;

		const detailsString = localStorage.getItem('emailVerificationDetails');

		if (detailsString === null) {
			goto('/signup');
			return;
		}

		const details = JSON.parse(detailsString);

		email = details.email;

		password = details.user_password;

		createUser();
	});

	async function createUser() {
		const response = await fetch(`${$apiHost}/auth/createUser`, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				email: email,
				password: password,
				token: token
			})
		});

		const result = await response.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			// goto('/signup');
			// return;
		}

		const session = result[0];

		localStorage.setItem('session', JSON.stringify(session));

		localStorage.removeItem('emailVerificationDetails');

		goto('/dashboard');
	}
</script>

<div class="container">
	<div class="innerWrapper">
		<div class="loader" />
		<span>Verifying your email...</span>
	</div>
</div>

<style lang="scss">
	.container {
		height: 100vh;
		width: 100vw;
		display: flex;
		align-items: center;
		justify-content: center;

		.innerWrapper {
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			gap: 2em;

			.loader {
				border: 2px solid #f3f3f3; /* Light grey */
				border-top: 2px solid #252b48; /* Blue */
				border-radius: 50%;
				width: 5rem;
				height: 5rem;
				animation: spin 0.5s linear infinite;

				@keyframes spin {
					0% {
						transform: rotate(0deg);
					}
					100% {
						transform: rotate(360deg);
					}
				}
			}

			span {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.3rem;
			}
		}
	}
</style>
