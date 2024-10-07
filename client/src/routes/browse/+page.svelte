<script>
	import { pb } from '$lib/pocketbase';
	let cards = [];
	let userId = pb.authStore.model?.id;
	let loading = true;

	function formatDate(timestamp) {
		return new Date(timestamp * 1000).toLocaleString();
	}

	function getNextReviewDate(lastReview, interval) {
		const nextReview = (lastReview + interval) * 1000;
		return new Date(nextReview).toLocaleString();
	}

	function getDueStatus(lastReview, interval) {
		const now = Date.now();
		const nextReview = (lastReview + interval) * 1000;
		if (now >= nextReview) {
			return 'due';
		} else if (now >= nextReview - 24 * 60 * 60 * 1000) {
			return 'soon';
		}
		return 'future';
	}

	async function fetchCards() {
		try {
			loading = true;
			const res = await pb.collection('cards').getFullList({
				filter: `owner = "${userId}"`,
				sort: 'created'
			});
			cards = res;
		} catch (error) {
			console.error('Error fetching cards:', error);
		} finally {
			loading = false;
		}
	}

	async function deleteCard(cardId) {
		if (confirm('Are you sure you want to delete this card?')) {
			try {
				await pb.collection('cards').delete(cardId);
				cards = cards.filter((card) => card.id !== cardId);
			} catch (error) {
				console.error('Error deleting card:', error);
			}
		}
	}

	function editCard(cardId) {
		window.location.href = `/edit/${cardId}`;
	}

	fetchCards();
</script>

<div class="browse-container">
	<header>
		<h1 class="gradient-text">Browse Cards</h1>
		<div class="header-buttons">
			<a href="/app" class="button secondary">Back to App</a>
			<a href="/create" class="button">Create New Card</a>
		</div>
	</header>

	<main>
		{#if loading}
			<div class="loading">
				<div class="loading-spinner"></div>
				<p>Loading cards...</p>
			</div>
		{:else if cards.length === 0}
			<div class="empty-state">
				<p>You haven't created any cards yet.</p>
				<a href="/create" class="button">Create Your First Card</a>
			</div>
		{:else}
			<div class="cards-grid">
				{#each cards as card}
					<div
						class="card"
						class:due={getDueStatus(card.lastReview, card.interval) === 'due'}
						class:soon={getDueStatus(card.lastReview, card.interval) === 'soon'}
					>
						<div class="card-content">
							<div class="card-front">{card.front}</div>
							<details class="card-back">
								<summary>Show Answer</summary>
								<p>{card.back}</p>
							</details>
						</div>
						<div class="card-info">
							<div class="info-item">
								<span class="label">Last Review</span>
								<span class="value">{formatDate(card.lastReview)}</span>
							</div>
							<div class="info-item">
								<span class="label">Next Review</span>
								<span class="value">{getNextReviewDate(card.lastReview, card.interval)}</span>
							</div>
						</div>
						<div class="card-actions">
							<button class="button secondary" on:click={() => editCard(card.id)}>Edit</button>
							<button class="button danger" on:click={() => deleteCard(card.id)}>Delete</button>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</main>
</div>

<style>
	.browse-container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 2rem 1rem;
	}

	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 2rem;
	}

	.header-buttons {
		display: flex;
		gap: 1rem;
	}

	.loading {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		min-height: 200px;
	}

	.loading-spinner {
		width: 40px;
		height: 40px;
		border: 4px solid var(--border-color);
		border-top-color: var(--accent-color);
		border-radius: 50%;
		animation: spin 1s linear infinite;
		margin-bottom: 1rem;
	}

	.empty-state {
		text-align: center;
		padding: 3rem 1rem;
	}

	.cards-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 1.5rem;
	}

	.card {
		background: var(--card-bg);
		border: 1px solid var(--border-color);
		border-radius: 0.5rem;
		padding: 1.5rem;
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.card.due {
		border-color: #ef4444;
	}

	.card.soon {
		border-color: #eab308;
	}

	.card-front {
		font-size: 1.25rem;
		font-weight: 500;
		margin-bottom: 0.5rem;
	}

	.card-back details {
		margin-top: 0.5rem;
	}

	.card-back summary {
		cursor: pointer;
		color: var(--accent-color);
	}

	.card-back summary:hover {
		color: var(--accent-color-2);
	}

	.card-info {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		font-size: 0.875rem;
	}

	.info-item {
		display: flex;
		justify-content: space-between;
	}

	.info-item .label {
		color: #94a3b8;
	}

	.card-actions {
		display: flex;
		gap: 0.5rem;
		margin-top: auto;
	}

	.button.secondary {
		background-color: var(--card-bg);
		border-color: var(--accent-color);
	}

	.button.danger {
		background-color: #ef4444;
		border-color: #dc2626;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	@media (max-width: 600px) {
		header {
			flex-direction: column;
			gap: 1rem;
			align-items: stretch;
		}

		.header-buttons {
			flex-direction: column;
		}
	}
</style>
