<script>
	import { GetCardBack, GetCardFront } from "../../wailsjs/go/main/Cards.js";

	let flipped = false;
	let cardFront = "card_front";
	let cardBack = "card_back";
	let currentCardId = 1;

	function flip() {
		flipped = !flipped;
	}

	function getCard(id) {
		GetCardFront(id).then((result) => {
			cardFront = result;
		});

		GetCardBack(id).then((result) => {
			cardBack = result;
		});
	}

	function nextCard() {
		if (currentCardId < 2) {
			currentCardId++;
			getCard(currentCardId);
		}
	}

	function previousCard() {
		if (currentCardId > 0) {
			currentCardId--;
			getCard(currentCardId);
		}
	}
</script>

<button on:click={previousCard}>Previous Card</button>
<button on:click={nextCard}>Next Card</button>

<div class="center">
	<div class="card-container" on:click={flip} on:keypress={flip}>
		<div class="card front {flipped ? 'flipped_front' : 'front'}">
			<p>{cardFront}</p>
		</div>
		<div class="card back {flipped ? 'flipped_back' : 'back'}">
			<p>{cardBack}</p>
		</div>
	</div>
</div>

<style>
	.center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	.card-container {
		perspective: 1000px;
		width: 440px;
		height: 300px;
		position: relative;
		text-align: center;
	}

	.card {
		position: absolute;
		width: 100%;
		height: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: transform 250ms;
		background-color: white;
		border-radius: 20px;
		color: black;
		font-size: 32px;
		backface-visibility: hidden;
		padding: 32px;
	}

	.front {
		transform: rotateY(0deg);
	}

	.back {
		transform: rotateY(-180deg);
	}

	.flipped_back {
		transform: rotateY(0deg);
	}

	.flipped_front {
		transform: rotateY(180deg);
	}
</style>
