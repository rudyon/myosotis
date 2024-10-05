<script>
	import { pb } from '$lib/pocketbase';

	let frontText = 'Loading...';
	let backText = 'Loading...';
	let isFlipped = false;
	let cards = [];
	let cardIndex = 0;
	let userId = pb.authStore.model?.id;
	let loading = false;

	async function fetchCards() {
		try {
			const now = Math.floor(Date.now() / 1000);
			const res = await pb.collection('cards').getFullList({
				filter: `owner = "${userId}"`
			});

			cards = res.filter((card) => card.lastReview + card.interval <= now);

			if (cards.length > 0) {
				loadCard();
			} else {
				frontText = 'You have no cards to review.';
				backText = '';
			}
		} catch (error) {
			console.error('Error fetching cards:', error);
			frontText = 'Error fetching cards. Please try again.';
			backText = '';
		}
	}

	function loadCard() {
		loading = true;
		setTimeout(() => {
			frontText = cards[cardIndex]?.front || 'No front text available.';
			backText = cards[cardIndex]?.back || 'No back text available.';
			isFlipped = false;
			loading = false;
		}, 300);
	}

	async function handleRemember() {
		await updateCard(true);
		loadNextCard();
	}

	async function handleForget() {
		await updateCard(false);
		loadNextCard();
	}

	async function updateCard(remembered) {
		const now = Math.floor(Date.now() / 1000);
		let newInterval = remembered ? cards[cardIndex].interval * 2 : 86400;
		await pb.collection('cards').update(cards[cardIndex].id, {
			lastReview: now,
			interval: newInterval
		});
	}

	function loadNextCard() {
		if (cardIndex < cards.length - 1) {
			cardIndex++;
			loadCard();
		} else {
			frontText = 'No more cards to review.';
			backText = '';
		}
	}

	function flipCard() {
		isFlipped = !isFlipped;
	}

	fetchCards();
</script>

<h1>Myosotis</h1>

<a href="/browse">Browse your cards</a>

<div id="cards">
	<div id="current_card">
		<div class="scale_on_hover" on:click={flipCard}>
			{#if loading}
				<div class="card hidden">
					<p>Loading...</p>
				</div>
			{:else}
				<div id="flip-container" flip={isFlipped ? 'yes' : 'no'}>
					<div class="flipper">
						<div id="card_front" class="card">
							<p>{frontText}</p>
						</div>
						<div id="card_back" class="card">
							<p>{backText}</p>
						</div>
					</div>
				</div>
			{/if}
		</div>
	</div>
</div>

<br />
<br />
<br />
<br />

<button on:click={handleForget}>Forgot</button>
<button on:click={handleRemember}>Remembered</button>

<style>
	#cards {
		position: relative;
		left: 100px;
		top: 35px;
	}

	#current_card {
		width: 400px;
		height: 240px;
		position: relative;
	}

	.card {
		width: 396px;
		height: 236px;
		border-radius: 20px;
		font-size: 40px;
		border: 2px inset #000;
		text-align: center;
		backface-visibility: hidden;
		position: absolute;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.scale_on_hover {
		width: 100%;
		height: 100%;
		transition: 0.15s;
		cursor: pointer;
	}

	.scale_on_hover:hover {
		transform: scale(1.03);
	}

	#flip-container {
		width: 100%;
		height: 100%;
		perspective: 1200px;
	}

	#flip-container[flip='yes'] .flipper {
		transform: rotateY(180deg);
	}

	#flip-container[flip='no'] .flipper {
		transform: rotateY(0deg);
	}

	#flip-container .flipper {
		width: 100%;
		height: 100%;
		transition: 0.5s;
		transform-style: preserve-3d;
		position: relative;
	}

	#card_front {
		z-index: 2;
		transform: rotateY(0deg);
		background: #fff;
	}

	#card_back {
		transform: rotateY(180deg);
		background: #fff;
	}

	.hidden {
		display: none;
	}
</style>
