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
				frontText = 'You have completed all reviews for now!';
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
			frontText = cards[cardIndex]?.front || 'No cards available.';
			backText = cards[cardIndex]?.back || 'No content available.';
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
			frontText = 'All done for now! Great job! 🎉';
			backText = '';
		}
	}

	function flipCard() {
		isFlipped = !isFlipped;
	}

	fetchCards();
</script>

<div class="app-container">
	<header>
		<h1 class="gradient-text">Myosotis</h1>
		<a href="/browse" class="button">Browse Cards</a>
	</header>

	<main>
		<div class="stats">
			<div class="stat">
				<span class="stat-label">Remaining</span>
				<span class="stat-value">{cards.length - cardIndex}</span>
			</div>
			<div class="stat">
				<span class="stat-label">Reviewed</span>
				<span class="stat-value">{cardIndex}</span>
			</div>
		</div>

		<div class="card-container">
			{#if loading}
				<div class="card loading">
					<div class="loading-spinner"></div>
				</div>
			{:else}
				<div class="flip-container" class:flipped={isFlipped} on:click={flipCard}>
					<div class="flipper">
						<div class="card front">
							<p>{frontText}</p>
						</div>
						<div class="card back">
							<p>{backText}</p>
						</div>
					</div>
				</div>
			{/if}
		</div>

		<div class="button-container">
			<button class="button forget" on:click={handleForget}> Forgot </button>
			<button class="button remember" on:click={handleRemember}> Remembered </button>
		</div>
	</main>
</div>

<style>
	.app-container {
		max-width: 800px;
		margin: 0 auto;
		padding: 2rem 1rem;
	}

	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 2rem;
	}

	h1 {
		font-size: 2rem;
		margin: 0;
	}

	.stats {
		display: flex;
		justify-content: center;
		gap: 2rem;
		margin-bottom: 2rem;
	}

	.stat {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.stat-label {
		font-size: 0.875rem;
		color: #94a3b8;
	}

	.stat-value {
		font-size: 1.5rem;
		font-weight: 600;
	}

	.card-container {
		perspective: 1000px;
		width: 100%;
		max-width: 500px;
		height: 300px;
		margin: 0 auto 2rem;
	}

	.flip-container {
		width: 100%;
		height: 100%;
		position: relative;
		cursor: pointer;
		transition: transform 0.1s;
	}

	.flip-container:hover {
		transform: scale(1.02);
	}

	.flip-container.flipped .flipper {
		transform: rotateY(180deg);
	}

	.flipper {
		width: 100%;
		height: 100%;
		transition: transform 0.6s;
		transform-style: preserve-3d;
		position: relative;
	}

	.card {
		width: 100%;
		height: 100%;
		padding: 2rem;
		position: absolute;
		backface-visibility: hidden;
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 1rem;
		background: var(--card-bg);
		border: 1px solid var(--border-color);
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
	}

	.card p {
		font-size: 1.5rem;
		margin: 0;
		text-align: center;
	}

	.card.back {
		transform: rotateY(180deg);
	}

	.card.loading {
		justify-content: center;
		align-items: center;
	}

	.loading-spinner {
		width: 40px;
		height: 40px;
		border: 4px solid var(--border-color);
		border-top-color: var(--accent-color);
		border-radius: 50%;
		animation: spin 1s linear infinite;
	}

	.button-container {
		display: flex;
		justify-content: center;
		gap: 1rem;
		margin-top: 2rem;
	}

	button {
		min-width: 120px;
	}

	.forget {
		background-color: #ef4444;
		border-color: #dc2626;
	}

	.remember {
		background-color: #22c55e;
		border-color: #16a34a;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	@media (max-width: 600px) {
		.card-container {
			height: 250px;
		}

		.card p {
			font-size: 1.25rem;
		}
	}
</style>
