<script lang="ts">
	import { afterNavigate, goto } from '$app/navigation';
	import { page } from '$app/stores';
	import heroSrc from '$lib/images/dashboard.png';
	import googleLogoSrc from '$lib/images/google-logo.png';
	import { apiHost, session } from '../stores/stores';

	afterNavigate(() => {
		loadUserSession();
	});

	const loadUserSession = () => {
		const sessionString = localStorage.getItem('session') as string | null;

		if (sessionString != null) {
			goto(`${$page.url.origin}/dashboard`);
			return;
		}
	};

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

<svelte:head>
	<title>BookmarkMonster - Free online bookmark manager | Save websites you can visit later</title>
	<meta name="description" content="Free online bookmark manager" />
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<!-- Hotjar Tracking Code for BookmarkMonster -->
	<script>
		(function (h, o, t, j, a, r) {
			h.hj =
				h.hj ||
				function () {
					(h.hj.q = h.hj.q || []).push(arguments);
				};
			h._hjSettings = { hjid: 3619057, hjsv: 6 };
			a = o.getElementsByTagName('head')[0];
			r = o.createElement('script');
			r.async = 1;
			r.src = t + h._hjSettings.hjid + j + h._hjSettings.hjsv;
			a.appendChild(r);
		})(window, document, 'https://static.hotjar.com/c/hotjar-', '.js?sv=');
	</script>
</svelte:head>

<section>
	<img
		src={heroSrc}
		alt="hero"
		class="hero"
		id="hero"
		on:contextmenu|preventDefault={() => {}}
		draggable="false"
	/>
	<div class="offer">
		<h1>Free online bookmark manager</h1>
		<p>Save and tag websites so you can visit later.</p>
	</div>
	<div class="buttons">
		<button class="google" on:click|stopPropagation={getGoogleLoginUrl}>
			<img src={googleLogoSrc} alt="google logo" class="google-logo" draggable="false" />
			<span>Continue with Google</span>
		</button>
		<div class="or" />
		<div class="magic-code">
			<input
				type="email"
				name="email"
				id="email"
				autocomplete="off"
				spellcheck="false"
				placeholder="Continue with email (coming soon)"
			/>
			<button type="submit">
				<span>Get magic code</span>
			</button>
		</div>
	</div>
</section>

<style lang="scss">
	section {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
		height: 100%;
		min-height: inherit;
		padding-top: 2em;
		padding-bottom: 1em;
		gap: 2em;
		background-color: hsl(209, 36%, 86%);
		background-attachment: fixed;
		background-size: 100% 100%;
		background-image: radial-gradient(
				50% 50% at 50% 50%,
				rgba(255, 255, 255, 0.75) 0%,
				rgba(255, 255, 255, 0) 100%
			),
			linear-gradient(180deg, rgb(202, 216, 228) 0%, hsl(209, 36%, 86%) 15%, hsl(224, 44%, 95%) 50%);

		img.hero {
			width: 60%;
			height: 60%;
			max-inline-size: 60%;
			object-fit: fill;
			object-position: center;
			border-radius: 0.5rem;

			@media only screen and (max-width: 425px) {
				width: 99%;
				max-inline-size: 99%;
			}

			@media only screen and (min-width: 426px) and (max-width: 768px) {
				width: 90%;
				max-inline-size: 90%;
			}
		}

		.offer {
			width: 50%;
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 0.5em;
			text-align: center;
			line-height: 1.6;

			h1 {
				color: #020202;
				font-size: 3rem;
				font-weight: 800;
				font-family: 'Inter', sans-serif !important;
			}

			p {
				font-size: 1.5rem;
				color: #020202;
				max-width: 50%;
			}

			@media only screen and (max-width: 768px) {
				width: 95%;

				h1 {
					font-weight: 500;
					font-size: 2rem;
				}

				p {
					max-width: 75%;
				}
			}
		}

		.buttons {
			width: 50%;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			gap: 1em;

			.magic-code {
				display: flex;
				align-items: center;
				width: 60%;
				height: 4rem;
				padding: 0.5em;
				background-color: rgba(255, 255, 255);
				border-radius: 0.3rem;
				border: 0.1rem solid transparent;
				gap: 0.1em;
				pointer-events: none;

				input[type='email'] {
					width: 70%;
					height: 100%;
					font-size: 1.3rem;
					outline: none;
					border: none;
					background-color: inherit;
					font-family: 'Arial CE', sans-serif;
				}

				button {
					width: 30%;
					min-width: max-content;
					height: 100%;
					background-color: rgb(6, 143, 255, 0.5);
					border: none;
					border-radius: 0.3rem;
					cursor: pointer;

					span {
						font-size: 1.3rem;
						color: rgba(255, 255, 255);
						font-family: 'Arial CE', sans-serif;
					}
				}

				&:hover {
					border-color: rgb(6, 143, 255);
				}

				&:focus-within {
					border-color: rgb(6, 143, 255);
				}

				@media only screen and (max-width: 425px) {
					width: 99%;
				}

				@media only screen and (min-width: 426px) and (max-width: 768px) {
					width: 85%;
				}
			}

			.or {
				width: 60%;
				height: 0.1rem;
				background-color: rgb(0, 0, 0, 0.1);
			}

			button.google {
				display: flex;
				align-items: center;
				justify-content: center;
				width: 60%;
				background-color: rgb(0, 121, 255);
				border: 0.1rem solid rgb(0, 0, 0, 0.1);
				cursor: pointer;
				height: 4rem;
				gap: 1em;
				border-radius: 0.3rem;

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

				@media only screen and (max-width: 425px) {
					width: 98%;
				}

				@media only screen and (min-width: 426px) and (max-width: 768px) {
					width: 85%;
				}
			}

			@media only screen and (max-width: 768px) {
				width: 100%;
			}
		}

		// @media only screen and (max-width: 320px) {

		// }
	}
</style>
