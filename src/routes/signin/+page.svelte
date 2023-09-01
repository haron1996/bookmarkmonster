<section>
	<form>
		<div class="top">
			<h1>Welcome back</h1>
		</div>
		<div class="inputs">
			<div class="email">
				<label for="email">Email address</label>
				<input type="email" name="email" id="email" autocomplete="email" />
			</div>
			<div class="password">
				<label for="password">Password</label>
				<div class="passwordInput">
					<input type="password" name="password" id="password" autocomplete="on" />
					<div class="showPassword">
						<i class="las la-eye" />
						<span>Show</span>
					</div>
				</div>
			</div>
		</div>
		<button type="submit">
			<i class="las la-sign-in-alt" />
			<span>Sign in</span>
		</button>
		<h2><span>or</span></h2>
		<button class="google">
			<i class="lab la-google-plus" />
			<span>Sign in with Google</span>
		</button>
		<span>Don't have an account yet? <a href="/signup">Sign up</a></span>
	</form>
</section>

<style lang="scss">
	section {
		width: 100%;
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;

		form {
			width: 50rem;
			height: max-content;
			display: flex;
			flex-direction: column;
			gap: 2em;
			padding: 3em;
			box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
			border-radius: 0.5rem;

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
							width: 85%;
							border: none;
							outline: none;
							background-color: inherit;
							padding: 0 0.5em;
							height: 100%;
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

							i {
								font-size: 1.6rem;
							}

							span {
								font-family: 'Roboto', sans-serif;
								font-size: 1.3rem;

								@media only screen and (max-width: 425px) {
									display: none;
								}
							}
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
				background-color: #0079ff;
				border: none;
				cursor: pointer;
				transition: all ease 300ms;
				border-radius: 0.3rem;
				gap: 1em;

				i {
					font-size: 2rem;
					color: white;
				}

				span {
					color: white;
					font-size: 1.3rem;
					font-family: 'Roboto', sans-serif;
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

				&:hover {
					transform: translateY(-5%);
				}
			}

			h2 {
				width: 100%;
				text-align: center;
				border-bottom: 1px solid #000;
				line-height: 0.1em;
				margin: 10px 0 20px;

				span {
					background: #fff;
					padding: 0 10px;
					font-size: 1.3rem;
					text-transform: uppercase;
					color: #4d4d4d;
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

		@media only screen and (max-width: 600px) {
			align-items: flex-start;
			padding: 1em 1em 0 1em;

			form {
				max-width: 98%;

				.top {
					h1 {
						font-size: 2.5rem;
					}
				}
			}
		}
	}
</style>
