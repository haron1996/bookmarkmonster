<script lang="ts">
	import { page } from '$app/stores';
	import logo from '$lib/images/logo.png';
	import googleLogoSrc from '$lib/images/google-logo.png';

	const getGoogleLoginUrl = async () => {
		const response = await fetch('http://localhost:5000/auth/get-google-login-url', {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		window.open(result, '_self');
	};
</script>

<header>
	<img src={logo} alt="logo" class="logo" draggable="false" />

	<!-- <nav /> -->

	<button class="google" on:click={getGoogleLoginUrl}>
		<img src={googleLogoSrc} alt="google logo" class="google-logo" draggable="false" />
		<span>Continue with Google</span>
	</button>
</header>

<style lang="scss">
	header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		min-height: 7vh;
		border-bottom: 0.1rem solid rgb(0, 0, 0, 0.1);

		img {
			height: 4.5rem;
			width: 4.5rem;
			object-fit: contain;
			margin-left: 1em;
		}

		button.google {
			display: flex;
			align-items: center;
			justify-content: center;
			min-width: max-content;
			background-color: rgb(0, 121, 255);
			border: none;
			cursor: pointer;
			height: 3.5rem;
			width: 15%;
			gap: 0.5em;
			border-radius: 0.3rem;
			margin-right: 0.5em;
			padding: 0 0.5em;

			img.google-logo {
				height: 2rem;
				width: 2rem;
				object-fit: contain;
			}

			span {
				font-family: 'Arial CE', sans-serif;
				font-size: 1.3rem;
				color: rgb(255, 255, 255);
			}

			&:hover {
				background-color: rgb(6, 143, 255);
			}
		}
	}
</style>
