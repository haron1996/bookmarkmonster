<script lang="ts">
	import { page } from '$app/stores';
	import logo from '$lib/images/logo.png';
	import googleLogoSrc from '$lib/images/google-logo.png';
	import { apiHost } from '../stores/stores';

	const getGoogleLoginUrl = async () => {
		const response = await fetch(`${$apiHost}/auth/get-google-login-url`, {
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
	<div class="innerContainer">
		<a class="logo" href="/"><img src={logo} alt="logo" draggable="false" /></a>

		<nav>
			<a href="#features">Features</a>
			<a href="#pricing">Pricing</a>
			<a href="/signin">Log In</a>
		</nav>

		<a class="btn" href="/signup">Get my free account</a>
	</div>
</header>

<style lang="scss">
	header {
		min-height: 7vh;
		display: flex;
		align-content: center;
		justify-content: center;
		background-color: #e6fffd;
		padding: 1em;

		.innerContainer {
			width: 60%;
			display: flex;
			align-items: center;
			justify-content: space-between;

			a.logo {
				height: 4rem;
				width: 4rem;

				img {
					object-fit: contain;
					max-inline-size: 100%;
				}
			}

			nav {
				min-width: max-content;
				display: flex;
				align-items: center;
				justify-content: center;
				gap: 1em;

				a {
					color: #025464;
					font-size: 1.3rem;
					font-weight: 600;
					text-decoration-color: transparent;
					transition: all ease 300ms;
					font-family: 'Roboto', sans-serif;

					&:hover {
						color: #0079ff;
						text-decoration-color: #0079ff;
					}
				}

				@media only screen and (max-width: 600px) {
					display: none;
				}
			}

			a.btn {
				min-width: max-content;
				padding: 0.5em 1em;
				background-color: #0079ff;
				font-size: 1.3rem;
				font-family: 'Poppins', sans-serif;
				color: white;
				text-decoration: none;
				border-radius: 0.3rem;
				font-weight: 600;
			}

			@media only screen and (max-width: 768px) {
				width: 100%;
			}

			@media only screen and (min-width: 601px) and (max-width: 992px) {
				width: 90%;
			}
		}
	}
</style>
