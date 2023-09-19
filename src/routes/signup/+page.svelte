<script lang="ts">
	import { apiHost } from '../../stores/stores';
	import { goto } from '$app/navigation';

	let email: string = '';
	let password: string = '';
	let requestingEmailVerificationToken: boolean = false;
	let requestingGoogleAuthLink: boolean = false;
	let passwordShown: boolean = false;

	const getGoogleLoginUrl = async () => {
		requestingGoogleAuthLink = true;

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

	async function requestEmailVerification() {
		if (email === '' || password === '') {
			alert('email and password required');
			return;
		}

		requestingEmailVerificationToken = true;

		const response = await fetch(`${$apiHost}/auth/requestEmailVerificationCode`, {
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
				password: password
			})
		});

		const result = await response.json();

		const msg = result.message;

		if (msg) {
			alert(msg);
			requestingEmailVerificationToken = false;
			return;
		}

		localStorage.setItem('emailVerificationDetails', JSON.stringify(result[0]));

		requestingEmailVerificationToken = false;

		goto('/signup/emailSent');
	}

	function showPassword() {
		const input = document.getElementById('password') as HTMLInputElement | null;

		if (input === null) return;

		input.type === 'password' ? (input.type = 'text') : (input.type = 'password');

		passwordShown = !passwordShown;
	}
</script>

<section>
	<div class="innerWrapper">
		<div class="features">
			<div class="title">
				<h1>Join 1000+ happy beta users</h1>
				<p>Quick and core features that help you unleash your bookmarking powers</p>
			</div>
			<p>Free plan includes</p>
			<div class="list">
				<div>
					<i class="las la-check" />
					<span>Unlimited tags</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Unlimited bookmarks</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Bulk editing</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Recycle bin</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>share folder with up to 3 users</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Unlimited tags</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Import from pocket</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Bookmark descriptions</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Category descriptions</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Basic sorting</span>
				</div>
				<div>
					<i class="las la-check" />
					<span>Browser extensions</span>
				</div>
			</div>
		</div>
		<form on:submit={requestEmailVerification}>
			<div class="top">
				<h1>Get your FREE account</h1>
				<p>No credit card required.</p>
			</div>
			<div class="inputs">
				<div class="email">
					<label for="email">Email address</label>
					<input type="email" name="email" id="email" autocomplete="email" bind:value={email} />
				</div>
				<div class="password">
					<label for="password">Password</label>
					<div class="passwordInput">
						<input
							type="password"
							name="password"
							id="password"
							autocomplete="on"
							bind:value={password}
						/>
						<div class="showPassword" on:click={showPassword} role="none" class:passwordShown>
							<i class="las la-eye" />
							<span>Show</span>
						</div>
					</div>
				</div>
			</div>
			<button
				type="submit"
				on:click|preventDefault={requestEmailVerification}
				class:btnDisabled={requestingEmailVerificationToken || email === '' || password === ''}
			>
				<i class="las la-envelope" />
				<span>Sign up with email</span>
				{#if requestingEmailVerificationToken || email === '' || password === ''}
					<div
						class="buttonBlocked"
						on:click|preventDefault|stopPropagation={() => {
							alert('blocked');
						}}
						role="none"
					/>
				{/if}
			</button>
			<h2><span>or</span></h2>
			<button
				class="google"
				on:click|preventDefault={getGoogleLoginUrl}
				class:btnDisabled={requestingGoogleAuthLink}
			>
				<i class="lab la-google-plus" />
				<span>Sign up with Google</span>
				{#if requestingGoogleAuthLink}
					<div
						class="buttonBlocked"
						on:click|preventDefault|stopPropagation={() => {
							alert('blocked');
						}}
						role="none"
					/>
				{/if}
			</button>
			<span>Already have an account? <a href="/signin">Log In</a></span>
		</form>
	</div>
</section>

<style lang="scss">
	section {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		min-height: 100vh;

		.innerWrapper {
			display: flex;
			width: 75%;
			min-height: max-content;
			min-height: inherit;
			align-items: center;
			justify-content: center;

			.features {
				width: 50%;
				display: flex;
				flex-direction: column;
				gap: 2em;
				padding: 2em;

				.title {
					display: flex;
					flex-direction: column;
					gap: 1em;

					h1 {
						color: #1e1919;
						font-size: 2rem;
						font-weight: 600;
						max-width: 80%;
						line-height: 1.3;
						font-family: 'Poppins', sans-serif;
					}

					p {
						font-size: 1.5rem;
						color: rgb(43, 58, 71);
						max-width: 80%;
						font-weight: 500;
						line-height: 1.3;
						font-family: 'Poppins', sans-serif;
					}
				}

				p {
					font-size: 1.4rem;
					color: rgb(43, 58, 71);
					max-width: 80%;
					font-weight: 600;
					line-height: 1.3;
					font-family: 'Poppins', sans-serif;
				}

				.list {
					display: flex;
					flex-direction: column;
					gap: 1em;

					div {
						display: flex;
						align-items: center;
						gap: 1em;

						i {
							background-color: #01d28e;
							height: 1.8rem;
							width: 1.8rem;
							border-radius: 100vh;
							display: flex;
							align-items: center;
							justify-content: center;
							font-size: 1.3rem;
							color: white;
						}

						span {
							font-family: 'Roboto', sans-serif;
							font-size: 1.3rem;
						}
					}
				}
			}

			form {
				width: 50rem;
				display: flex;
				flex-direction: column;
				gap: 2em;
				padding: 2em;
				box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
				border-radius: 0.5rem;
				background-color: aliceblue;

				.top {
					display: flex;
					flex-direction: column;

					h1 {
						color: #1e1919;
						font-size: 3rem;
						font-weight: 600;
						max-width: 80%;
						line-height: 1.3;
						font-family: 'Poppins', sans-serif;

						@media only screen and (max-width: 600px) {
							font-size: 2.4rem;
						}
					}

					p {
						font-size: 1.5rem;
						color: rgb(43, 58, 71);
						max-width: 80%;
						font-weight: 500;
						line-height: 1.6;
						font-family: 'Poppins', sans-serif;
					}
				}

				.inputs {
					display: flex;
					flex-direction: column;
					gap: 1em;

					.email {
						display: flex;
						flex-direction: column;
						display: flex;
						flex-direction: column;
						gap: 0.5em;

						label {
							font-size: 1.3rem;
							max-width: 80%;
							font-family: 'Poppins', sans-serif;
						}

						input {
							padding: 0 0.5em;
							font-family: 'Roboto', sans-serif;
							font-size: 1.3rem;
							border-radius: 0.3rem;
							border: 0.2rem solid #445069;
							background-color: white;
							min-height: 4rem;
						}
					}

					.password {
						display: flex;
						flex-direction: column;
						display: flex;
						flex-direction: column;
						gap: 0.5em;

						label {
							font-size: 1.3rem;
							max-width: 80%;
							font-family: 'Poppins', sans-serif;
						}

						.passwordInput {
							display: flex;
							align-items: center;
							width: 100%;
							border: 0.2rem solid #445069;
							height: 4rem;
							border-radius: 0.3rem;
							background-color: white;

							input {
								font-family: 'Roboto', sans-serif;
								font-size: 1.3rem;
								height: 100%;
								width: 85%;
								border: none;
								outline: none;
								background-color: inherit;
								padding: 0 0.5em;
							}

							.showPassword {
								min-width: 15%;
								height: 100%;
								display: flex;
								align-items: center;
								justify-content: center;
								border-left: 0.1rem solid rgb(0, 0, 0, 0.1);
								gap: 0.7em;
								cursor: pointer;
								padding: 0 0.2em;
								background-color: #e8f0fe;

								i {
									font-size: 1.6rem;
								}

								span {
									font-family: 'Roboto', sans-serif;
									font-size: 1.3rem;
								}

								@media only screen and (max-width: 425px) {
									span {
										display: none;
									}
								}
							}

							.passwordShown {
								color: #0079ff;
							}
						}
					}
				}

				button {
					display: flex;
					align-items: center;
					width: 100%;
					justify-content: center;
					height: 4rem;
					font-family: 'Roboto', sans-serif;
					background-color: #0079ff;
					border: none;
					cursor: pointer;
					transition: all ease 300ms;
					border-radius: 0.3rem;
					gap: 1em;
					position: relative;

					i {
						font-size: 2rem;
						color: white;
					}

					span {
						color: white;
						font-size: 1.3rem;
					}

					&.google {
						background-color: white;
						border: 0.1rem solid #1e1919;

						i {
							color: #1e1919;
						}

						span {
							color: #1e1919;
						}
					}

					.buttonBlocked {
						position: absolute;
						top: 0;
						left: 0;
						bottom: 0;
						right: 0;
						height: inherit;
						width: inherit;
						cursor: not-allowed;
						background-color: transparent;
					}
				}

				.btnDisabled {
					opacity: 0.5;
				}

				h2 {
					width: 100%;
					text-align: center;
					border-bottom: 1px solid #000;
					line-height: 0.1em;
					margin: 10px 0 20px;

					span {
						background: aliceblue;
						padding: 0 10px;
						font-size: 1.3rem;
						text-transform: uppercase;
						color: #4d4d4d;
						font-family: 'Roboto', sans-serif;
					}
				}

				span {
					font-size: 1.3rem;
					font-family: 'Roboto', sans-serif;

					a {
						color: #0079ff;
					}
				}
			}

			@media only screen and (max-width: 768px) {
				flex-direction: column-reverse;
				width: 100%;
				padding: 1em;
				gap: 3em;

				.features {
					width: 60rem;
					max-width: 98%;
					padding: 1em;
				}

				form {
					width: 60rem;
					max-width: 98%;
					padding: 1em;
				}
			}

			@media only screen and (min-width: 769px) and (max-width: 992px) {
				width: 100%;
				flex-direction: column-reverse;
				gap: 3em;
				justify-content: flex-end;
				padding: 3em 1em 1em 1em;

				.features {
					width: 60rem;
					max-width: 98%;
					padding: 1em;
				}

				form {
					width: 60rem;
					max-width: 98%;
					padding: 1em;
				}
			}
		}
	}
</style>
