<script>
	import { pb } from '$lib/pocketbase';

	let cards = [];
	let userId = pb.authStore.model?.id;

	async function fetchCards() {
		try {
			const res = await pb.collection('cards').getFullList({
				filter: `owner = "${userId}"`
			});
			cards = res;
		} catch (error) {
			console.error('Error fetching cards:', error);
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

<h1>Browse Cards</h1>

<a href="/app">Go back to app</a>

<ul>
	{#each cards as card}
		<li>
			<strong>{card.front}</strong> |
			<details><summary>Back</summary><em>{card.back}</em></details>
			| {card.lastReview} | {card.interval}
			<button on:click={() => editCard(card.id)}>Edit</button>
			<button on:click={() => deleteCard(card.id)}>Delete</button>
		</li>
	{/each}
</ul>

<a href="/create">Create new card</a>

<style>
	ul {
		list-style-type: none;
		padding: 0;
	}

	li {
		border: 1px solid #000;
		border-radius: 5px;
		padding: 10px;
		margin: 5px 0;
		display: flex;
		justify-content: space-between;
	}

	button {
		margin-left: 10px;
	}
</style>
