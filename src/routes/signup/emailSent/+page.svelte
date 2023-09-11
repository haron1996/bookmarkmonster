<script lang="ts">
	import { onMount } from 'svelte';
	import gmail_logo from '$lib/images/gmail_logo.png';
	import outlook_logo from '$lib/images/outlook_logo.png';
	import { apiHost } from '../../../stores/stores';

	let email: string = '';
	let loading: boolean = false;

	onMount(() => {
		const detailsString = localStorage.getItem('emailVerificationDetails');

		if (detailsString === null) return;

		const details = JSON.parse(detailsString);

		email = details.email;
	});

	async function resendConfirmationLink() {
		loading = true;

		const response = await fetch(`${$apiHost}/auth/resendEmailConfirmationLink`, {
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
				email: email
			})
		});

		const result = await response.json();

		alert(result[0]);

		loading = false;
	}
</script>

<div class="emailSent">
	<div class="innerContent">
		<h1>Check your email address</h1>
		<span
			>We've emailed <b>haronkibetrutoh@gmail.com</b> a confirmation link that will expire shortly. Click
			it to verify your email.</span
		>
		<div class="mailClients">
			<a href="https://mail.google.com/" class="openGmail">
				<img src={gmail_logo} alt="gmail" />
				<span>Open Gmail</span>
			</a>
			<a href="https://outlook.live.com/mail/" class="openOutlook">
				<img src={outlook_logo} alt="outlook" />
				<span>Open Outlook</span>
			</a>
		</div>
		<button on:click|preventDefault={resendConfirmationLink}>
			{#if loading}
				<span>Resending confimation link...</span>
			{:else}
				<span>Resend confimation link</span>
			{/if}
		</button>
	</div>
</div>

<style lang="scss">
	.emailSent {
		height: 100vh;
		width: 100vw;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: white;

		.innerContent {
			height: 60vh;
			width: 60rem;
			background-color: inherit;
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 3em;
			text-align: center;

			h1 {
				font-size: 2rem;
				font-family: 'Product Sans Medium', sans-serif;
			}

			span {
				font-family: 'Poppins', sans-serif;
				font-size: 1.5rem;
			}

			.mailClients {
				display: flex;
				align-items: center;
				gap: 2em;

				a {
					display: flex;
					align-items: center;
					gap: 0.7em;
					text-decoration: none;
					border: 0.1rem solid rgb(0, 0, 0, 0.1);
					padding: 0.3em 1em;
					border-radius: 0.3rem;
					color: #272829;
					transition: all 300ms ease-in-out;

					img {
						height: 3rem;
						width: 3rem;
						object-fit: contain;
						max-inline-size: 100%;
					}

					span {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
					}

					&:hover {
						background-color: aliceblue;
					}
				}
			}

			button {
				width: 100%;
				height: 4rem;
				border: none;
				cursor: pointer;
				background-color: #0079ff;
				border-radius: 0.3rem;
				transition: all 300ms ease-in-out;

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.3rem;
					color: white;
				}

				&:hover {
					transform: translateY(-0.5rem);
				}
			}
		}
	}
</style>
