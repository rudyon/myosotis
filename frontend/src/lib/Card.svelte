<script>
	import {
		GetTotalCards,
		GetCardBack,
		GetCardFront,
		UpdateCard,
	} from "../../wailsjs/go/main/Cards.js";

	let flipped = false;
	let cardFront = "card_front";
	let cardBack = "card_back";
	let currentCardId = 0;
	let totalCards = 0;

	function flip() {
		flipped = !flipped;
	}

	function getTotalCards() {
		GetTotalCards().then((result) => {
			totalCards = result;
			// Check if there are cards before getting the current card
			if (totalCards > 0) {
				getCard(currentCardId);
			}
		});
	}

	function getCard(id) {
		GetCardFront(id).then((result) => {
			cardFront = result;
		});

		GetCardBack(id).then((result) => {
			cardBack = result;
		});
	}

	function recalledCard() {
		UpdateCard(currentCardId, true);
		nextCard();
	}

	function forgotCard() {
		UpdateCard(currentCardId, false);
		nextCard();
	}

	function nextCard() {
		if (currentCardId < totalCards - 1) {
			currentCardId++;
			getCard(currentCardId);
		} else {
			// All cards completed
			currentCardId = totalCards;
		}
	}

	getTotalCards();
</script>

<div class="center-vertical center-horizontal">
	{#if currentCardId >= totalCards}
		<p>Completed</p>
	{:else}
		<div class="card-container" on:click={flip} on:keypress={flip}>
			<div class="card front {flipped ? 'flipped_front' : 'front'}">
				<p>{cardFront}</p>
			</div>
			<div class="card back {flipped ? 'flipped_back' : 'back'}">
				<p>{cardBack}</p>
			</div>
		</div>
	{/if}
</div>

<div class="center-horizontal">
	<button on:click={forgotCard}>bad</button>
	<button on:click={recalledCard}>good</button>
</div>

<style>
	button {
		display: inline-block;
		min-width: 140px;
		height: 40px;
		margin: 5px;
		left: 20px;
		border-radius: 12px;
		border: none;
		font-size: 20px;
		line-height: 40px;
		letter-spacing: 0.5px;
		cursor: pointer;
		transition: 0.1s;
		position: relative;
	}

	.center-horizontal {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 85lvw;
	}

	.center-vertical {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 80lvh;
	}

	.card-container {
		perspective: 1000px;
		width: 396px;
		height: 236px;
		position: relative;
		cursor: pointer;
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
