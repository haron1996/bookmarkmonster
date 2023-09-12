<script lang="ts">
	import { apiHost } from '../../stores/stores';

	let name: string = '';
	let email: string = '';
	let company_name: string = '';
	let comment: string = '';

	async function joinWaitList() {
		if (email === '' || name === '') {
			alert('email and name required');
			return;
		}

		const response = await fetch(`${$apiHost}/care/joinWaitlist`, {
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
				name: name,
				email: email,
				company_name: company_name,
				comment: comment
			})
		});

		const result = await response.json();

		console.log(result[0]);
	}
</script>

<section>
	<div class="heading">
		<h1 data-id="💎">Get notified when we launch a pro version</h1>
		<p>You will be among the first to know as soon as we launch</p>
	</div>
	<form autocomplete="off">
		<div class="title formItem">
			<p data-id="required">What's your full name?</p>
			<span>This helps us personalize our emails to you.</span>
			<input
				type="text"
				name="title"
				id="title"
				autocomplete="name"
				spellcheck="false"
				required
				bind:value={name}
			/>
		</div>
		<div class="title formItem">
			<p data-id="required">What's your email address?</p>
			<span>This is where we’ll get back to you. Double check that it’s right.</span>
			<input
				type="email"
				name="title"
				id="title"
				spellcheck="false"
				autocomplete="email"
				required
				bind:value={email}
			/>
		</div>
		<div class="title formItem">
			<p data-id="optional">What's your company name?</p>
			<span>This helps us better understand our user base so we can build a better product.</span>
			<input
				type="email"
				name="title"
				id="title"
				spellcheck="false"
				autocomplete="off"
				bind:value={company_name}
			/>
		</div>
		<div class="description formItem">
			<p data-id="optional">Any question, comment, or issue?</p>
			<span>This helps us know if you have other concerns that you'd like us to know</span>
			<textarea name="description" id="description" autocomplete="off" bind:value={comment} />
		</div>
		<button on:click|preventDefault|stopPropagation={joinWaitList}>Join the pro waiting list</button
		>
	</form>
</section>

<style lang="scss">
	section {
		width: 100%;
		height: max-content;
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: flex-start;
		flex-direction: column;
		background-color: rgb(255, 255, 255);
		gap: 3em;
		padding: 3em 0;

		.heading {
			display: flex;
			flex-direction: column;
			align-items: center;
			width: max-content;
			text-align: center;
			gap: 0.5em;
			max-width: fit-content;

			h1 {
				color: #1e1919;
				font-size: 3rem;
				font-weight: 900;
				max-width: 80%;
				line-height: 1.3;
				font-family: 'Poppins', sans-serif;
				max-width: fit-content;
			}

			p {
				font-size: 1.8rem;
				color: rgb(43, 58, 71);
				max-width: 80%;
				font-weight: 500;
				line-height: 1.3;
				max-width: fit-content;
			}
		}

		form {
			display: flex;
			flex-direction: column;
			border: 0.2rem dashed #025464;
			width: max-content;
			width: 90%;
			padding: 3em;
			gap: 3em;
			background-color: rgb(252, 244, 242);
			border-radius: 1rem;

			.formItem {
				display: flex;
				flex-direction: column;
				gap: 0.5em;

				p {
					font-family: 'Poppins', sans-serif;
					font-size: 2rem;
					font-weight: 600;
					position: relative;
					width: max-content;
					max-width: fit-content;

					&::after {
						content: attr(data-id);
						position: absolute;
						top: -1rem;
						left: 101%;
						background: black;
						background-color: #f4d160;
						font-size: 1.2rem;
						padding: 0.1em 0.5em;
						border-radius: 0.6rem;
						font-weight: 500;

						@media only screen and (max-width: 600px) {
							display: none;
						}
					}
				}

				span {
					font-family: 'Poppins', sans-serif;
					font-size: 1.8rem;
					width: max-content;
					max-width: fit-content;
				}

				input {
					padding: 1em 0.5em;
					font-family: 'Roboto', sans-serif;
					font-size: 1.6rem;
					color: #445069;
					border-radius: 0.5rem;
					border: 0.2rem solid #445069;
					background-color: white;
				}

				textarea {
					padding: 1em 0.5em;
					font-family: 'Roboto', sans-serif;
					font-size: 1.6rem;
					color: #445069;
					border-radius: 0.5rem;
					border: 0.2rem solid #445069;
					width: 100%;
					min-height: 30rem;
					resize: vertical;
					transition: all ease 300ms;
					background-color: white;
				}
			}

			button {
				border: none;
				border-radius: 0.3rem;
				width: 100%;
				padding: 1em 0;
				background-color: rgb(45, 165, 48);
				color: white;
				font-family: 'Roboto', sans-serif;
				font-size: 1.6rem;
				cursor: pointer;
				transition: all ease 300ms;
				font-weight: 600;

				&:hover {
					transform: translateY(-5%);
				}
			}

			@media only screen and (max-width: 992px) {
				width: 97%;
				padding: 1em;
			}
		}
	}
</style>
