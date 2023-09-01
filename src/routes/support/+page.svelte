<section>
	<div class="heading">
		<h1>Friendly folks, standing by.</h1>
		<p>There's no stupid question. Anything you want to ask, we are ready to answer.</p>
	</div>
	<form>
		<div class="title formItem">
			<p data-id="required">What do you need help with?</p>
			<span>This helps make sure you get the right answer fast.</span>
			<input type="text" name="title" id="title" spellcheck="false" autocomplete="off" />
		</div>
		<div class="description formItem">
			<p data-id="required">What's your question, comment, or issue?</p>
			<span>Share all the details. The more we know, the better we can help you.</span>
			<textarea name="description" id="description" autocomplete="off" spellcheck="false" />
		</div>
		<div class="title formItem">
			<p data-id="required">What's your email address?</p>
			<span>This is where we’ll get back to you. Double check that it’s right.</span>
			<input type="email" name="title" id="title" spellcheck="false" autocomplete="off" />
		</div>
		<button>Send support request</button>
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
			max-width: fit-content;
			text-align: center;
			gap: 0.5em;

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
						max-width: fit-content;
					}

					@media only screen and (max-width: 600px) {
						&::after {
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
					width: 100%;
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
				font-weight: 500;

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
