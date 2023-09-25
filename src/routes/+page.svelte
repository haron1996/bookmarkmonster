<script lang="ts">
	import hero from '$lib/images/dashboard.png';
	import Footer from './Footer.svelte';
	import Header from './Header.svelte';
	import signature from '$lib/images/signature.png';
</script>

<svelte:head>
	<title>BookmarkMonster - Online bookmark manager | Save websites you can visit later</title>
	<meta name="description" content="Free online bookmark manager" />
</svelte:head>

<Header />

<section id="section1">
	<img
		src={hero}
		alt="hero"
		class="hero"
		id="hero"
		on:contextmenu|preventDefault={() => {}}
		draggable="false"
	/>
	<div class="offer">
		<h1>FREE Bookmark Manager for Individuals, researchers, and Teams.</h1>
		<p>
			Save, organize, and share links with team... has tags, nested categories, and a built-in to-do
			list. 🔥
		</p>
	</div>

	<a href="/signup">Get my free account</a>
</section>

<section id="features">
	<div class="title">
		<h2>Bookmarking re-imagined</h2>
		<p>Built to make managing and sharing bookmarks easy and fun</p>
	</div>
	<div class="innerWrapper">
		<div class="featurePlusBenefit">
			<i class="las la-folder" />
			<p>Nested categories</p>
			<span> Organize your bookmarks into folders for easier access. </span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="lab la-slack" />
			<p>Tags</p>
			<span> Tag and filter bookmarks by tags for easier access </span>
			<span class="status beta">Beta</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="lab la-searchengin" />
			<p>Robust search</p>
			<span>
				Find the link or folder you're looking for with our powerful search-as-you-type feature
			</span>
			<span class="status beta">Beta</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-share-alt" />
			<p>Share links</p>
			<span>
				Emailing links is too cumbersome. Want to share bookmarks with team? Create a folder, add
				links, and invite your teammates! It's that easy!
			</span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-comments" />
			<p>Collaboration</p>
			<span> Discuss and comment on any link or folder with your team. </span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-edit" />
			<p>Link notes</p>
			<span> Add little notes to any link you save for context </span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-image" />
			<p>Thumbnails and favicons</p>
			<span>
				Our system automatically captures the screenshot and favicon of each link you save. Edit
				bookmark thumbnail and favicon whenever you want.
			</span>
			<span class="status beta">Beta</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-pencil-alt" />
			<p>Bulk editing</p>
			<span>
				Save time by editting items in builk: update, move to trash, restore from trash,
				drag-and-drop to move, and more! No need to manage each bookmark one by one
			</span>
			<span class="status beta">Beta</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-list-alt" />
			<p>To-Do list</p>
			<span>
				Save time and increase your productivity with our built-in to-do list. You don't have to
				switch between your bookmark manager and to-do list apps.
			</span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="lab la-chrome" />
			<p>Browser extensions</p>
			<span>
				Stay in the flow by saving saving links right from your browser. No need to visit our web
				app just to save a link!
			</span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-shield-alt" />
			<p>Duplicate checker</p>
			<span>
				Our system will stop you from adding duplicate links in the same folder so you never have to
				deal with duplicates
			</span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-sort" />
			<p>Sort alphabetically</p>
			<span> Sort bookmarks by alphabetical order so you can easily find bookmarks. </span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-lock" />
			<p>Password protected folders</p>
			<span>
				Increase bookmark privacy by protecting your folders with passwords. Only users with correct
				folder password will be allowed to access.
			</span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit">
			<i class="las la-recycle" />
			<p>Powerful recycle bin</p>
			<span>
				Accidentally deleted a bookmark or folder? No worries! Easily restore any deleted item
				within 30 days.
			</span>
			<span class="status cooking">Cooking</span>
		</div>
		<div class="featurePlusBenefit requestFeature">
			<i class="las la-lightbulb" />
			<p>Your favorite feature</p>
			<span>
				We are committed to building a product you actually want to use. Any feature that you want
				added? Send us a request. We'll prioritize the most request ones
			</span>
			<span class="status requestFeature">
				<a href="/support">Request feature</a>
			</span>
		</div>
	</div>
</section>

<section id="pricing">
	<div class="title">
		<h2>Free forever, Upgrade any time.</h2>
		<p>Simple pricing. Advanced features.</p>
	</div>

	<div class="pricingCards">
		<div class="priceCard free">
			<div class="top">
				<div class="head">
					<p>FREE</p>
					<span class="price"><strong>$0</strong>/month</span>
				</div>
				<!-- <a href="/">Get a free account</a> -->
				<div class="benefit">
					<i class="las la-check" />
					<span>Unlimited tags</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Unlimited bookmarks</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Basic search</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Bulk editing</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Link notes</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Recycle bin</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>browser extensions</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Basic sorting</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Share folder with up to 3 people</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Add supported</span>
				</div>
			</div>
			<div class="bottom">
				<a href="/signup">Get my free account</a>
			</div>
		</div>
		<div class="priceCard pro">
			<div class="top">
				<div class="head">
					<p>PRO</p>
					<span class="price">Launching soon</span>
				</div>
				<!-- <a href="/">Let me know when you launch</a> -->
				<div class="benefit">
					<i class="las la-check" />
					<span>Everything in free</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Unlimited nested collections</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Robust search</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Share collection with unlimited users</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Comment and discuss on collections and bookmarks</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Duplicate checker</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Protect collections with password</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>To-do list</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Simple notes</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>Advance sorting</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>No ads</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>URL to PDF</span>
				</div>
				<div class="benefit">
					<i class="las la-check" />
					<span>URL to screenshot</span>
				</div>
			</div>
			<div class="bottom">
				<a href="/pro-waiting-list">Notify me when you launch</a>
			</div>
		</div>
	</div>
</section>

<div class="about" id="about">
	<div class="card">
		<p>Hey there 👋</p>

		<p>I'm Haron, founder here at Bookmark Monster.</p>

		<p>
			Bookmarks can become overwhelming, and most existing solutions are complex or lack core
			bookmarking features.
		</p>

		<p>
			That's why we are on a mission to build a simple online bookmark manager to make saving,
			organizing, and sharing bookmarks fun.
		</p>

		<p>Happy bookmarking!</p>

		<img src={signature} alt="founder" />
	</div>
</div>

<div class="signup">
	<a href="/signup">Get a free account</a>
</div>

<Footer />

<style lang="scss">
	#section1 {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		width: 100%;
		min-height: max-content;
		min-height: 100vh;
		gap: 2em;
		background-color: #e6fffd;

		img.hero {
			width: 60%;
			height: 60%;
			max-inline-size: 60%;
			object-fit: fill;
			object-position: center;
			border-radius: 0.5rem;
			box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em, rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em,
				rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;

			@media only screen and (max-width: 600px) {
				width: 95%;
				max-inline-size: 95%;
			}
		}

		.offer {
			width: 50%;
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 1em;
			text-align: center;

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

			@media only screen and (max-width: 992px) {
				width: 100%;
				padding: 0 0.5em;

				h1 {
					font-weight: 600;
					font-size: 2rem;
				}
			}

			// @media only screen and (min-width: 601px) and (max-width: 992px) {
			// 	background-color: orange;
			// }
		}

		a {
			min-width: max-content;
			padding: 1em;
			background-color: #0079ff;
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			color: white;
			text-decoration: none;
			border-radius: 0.3rem;
			display: flex;
			align-items: center;
			justify-content: center;
			min-width: 30rem;

			@media only screen and (max-width: 600px) {
				min-width: 80%;
			}
		}
	}

	#features {
		width: 100%;
		min-height: 100vh;
		max-height: max-content;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 2em;
		padding: 5em 0;
		background-color: inherit;

		.title {
			display: flex;
			flex-direction: column;
			gap: 1em;
			text-align: center;
			align-items: center;

			h2 {
				font-family: 'Poppins', sans-serif;
				font-size: 2.5rem;
				color: #1e1919;
				font-weight: 900;
			}

			p {
				font-size: 1.6rem;
				font-family: 'Poppins', sans-serif;
				font-size: rgb(43, 58, 71);
			}

			@media only screen and (max-width: 600px) {
				padding: 0.5em;
				gap: 0.5em;

				h2 {
					font-size: 2rem;
				}
			}
		}

		.innerWrapper {
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			align-items: center;
			gap: 3em;
			width: 80%;

			.featurePlusBenefit {
				display: flex;
				flex-direction: column;
				min-height: max-content;
				padding: 1em;
				gap: 1em;
				border: 0.1rem solid rgb(4, 13, 18, 0.2);

				i {
					font-size: 2.5rem;
				}

				p {
					font-size: 1.5rem;
					font-family: 'Poppins', sans-serif;
					font-weight: 600;
				}

				span {
					font-size: 1.4rem;
					font-family: 'Poppins', sans-serif;
					line-height: 1.6;
				}

				span.status {
					font-size: 1.3rem;
					max-width: max-content;
					padding: 0.1em 1em;
					border-radius: 0.3rem;

					&.cooking {
						background-color: #f8fdcf;
					}

					&.beta {
						background-color: #e2f6ca;
					}

					&.requestFeature {
						background-color: #ffd24c;
						cursor: pointer;

						a {
							color: #0079ff;
							text-decoration: none;
							color: #1e1919;
						}
					}
				}

				&.requestFeature {
					border: 0.2rem solid #ffd24c;
					box-shadow: rgba(0, 0, 0, 0.09) 0px 3px 6px;
				}
			}

			@media only screen and (max-width: 600px) {
				width: 100%;
				grid-template-columns: repeat(1, 1fr);
				padding: 0.5em;
			}

			@media screen and (min-width: 601px) and (max-width: 992px) {
				width: 95%;
			}
		}
	}

	#pricing {
		width: 100%;
		min-height: max-content;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 2em;
		background-color: inherit;
		padding: 2em 0;
		display: none;

		.title {
			display: flex;
			flex-direction: column;
			gap: 1em;
			text-align: center;
			align-items: center;

			h2 {
				font-family: 'Poppins', sans-serif;
				font-size: 2.5rem;
				color: #1e1919;
				font-weight: 900;
			}

			p {
				font-size: 1.6rem;
				font-family: 'Poppins', sans-serif;
				font-size: rgb(43, 58, 71);
			}

			@media only screen and (max-width: 600px) {
				h2 {
					font-size: 2rem;
				}
			}
		}

		.pricingCards {
			width: max-content;
			display: flex;
			justify-content: center;
			align-items: flex-start;
			gap: 2em;
			min-height: max-content;

			.priceCard {
				width: 40rem;
				display: flex;
				flex-direction: column;
				padding: 1em;
				justify-content: space-between;
				box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
				border-radius: 0.5rem;
				gap: 3em;
				min-height: max-content;

				.top {
					display: flex;
					flex-direction: column;
					gap: 2em;

					.head {
						display: flex;
						flex-direction: column;
						gap: 0.5em;

						p {
							font-family: 'Poppins', sans-serif;
							color: #1e1919;
							font-size: 1.8rem;
							font-weight: 500;
						}

						span.price {
							font-family: 'Poppins', sans-serif;
							font-size: 1.5rem;
							color: #0079ff;
						}
					}

					a {
						width: 100%;
						padding: 0.5em 0;
						color: rgb(253, 253, 253);
						background-color: #0079ff;
						display: flex;
						align-items: center;
						justify-content: center;
						border-radius: 0.3rem;
						font-family: 'Poppins', sans-serif;
						font-size: 1.3rem;
						text-decoration: none;
					}

					.benefit {
						display: flex;
						align-items: center;
						gap: 1em;

						i {
							font-size: 1.5rem;
						}

						span {
							font-size: 1.3rem;
							font-family: 'Poppins', sans-serif;
						}
					}
				}

				.bottom {
					a {
						width: 100%;
						padding: 1em;
						color: rgb(253, 253, 253);
						background-color: #0079ff;
						display: flex;
						align-items: center;
						justify-content: center;
						border-radius: 0.3rem;
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						font-weight: 600;
						text-decoration: none;
						box-shadow: rgba(50, 50, 93, 0.25) 0px 2px 5px -1px, rgba(0, 0, 0, 0.3) 0px 1px 3px -1px;

						&:hover {
							text-decoration: underline;
						}
					}
				}

				&.pro {
					box-shadow: rgb(0, 121, 255) 0px 0px 0px 3px;
				}
			}

			@media only screen and (max-width: 600px) {
				flex-direction: column;
				width: 100%;
				padding: 0.7em;
				align-items: center;

				.priceCard {
					max-width: 95%;
				}
			}

			@media only screen and (min-width: 600px) and (max-width: 899px) {
				flex-direction: column;
				width: 100%;
				padding: 0em;
				align-items: center;
				min-height: max-content;

				.priceCard {
					max-width: 95%;
				}
			}
		}
	}

	#about {
		width: 100%;
		min-height: 100vh;
		max-height: max-content;
		padding: 2em 1em;
		display: flex;
		align-items: center;
		justify-content: center;
		display: none;

		.card {
			display: flex;
			flex-direction: column;
			width: 50%;
			padding: 2em;
			box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
			gap: 3em;
			border-radius: 1rem;

			p {
				font-size: 2rem;
				font-family: 'Poppins', sans-serif;
				line-height: 1.6;
			}

			@media only screen and (max-width: 600px) {
				width: 100%;
			}

			@media only screen and (min-width: 601px) and (max-width: 992px) {
				width: 90%;
			}
		}
	}

	.signup {
		padding: 1em;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;

		a {
			text-decoration: none;
			padding: 1em;
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			background-color: #0079ff;
			color: white;
			border-radius: 0.3rem;
			width: max-content;
		}
	}
</style>
