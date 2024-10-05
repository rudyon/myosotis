<script>
	import { pb } from '$lib/pocketbase';

	let front = '';
	let back = '';
	let errorMessage = '';
	let successMessage = '';
	let isLoggedIn = pb.authStore.isValid; // Check if the user is logged in

	async function handleSubmit(event) {
		event.preventDefault();

		if (!isLoggedIn) {
			errorMessage = 'You must be logged in to create a card.';
			return;
		}

		if (!front || !back) {
			errorMessage = 'Please fill out both fields before submitting.';
			return;
		}

		try {
			const now = Math.floor(Date.now() / 1000);
			const cardData = {
				front: front,
				back: back,
				owner: pb.authStore.model.id,
				lastReview: now,
				interval: 86400
			};

			console.log('Card Data:', cardData);

			await pb.collection('cards').create(cardData);
			successMessage = 'Card created successfully!';
			front = '';
			back = '';
		} catch (error) {
			errorMessage = 'Failed to create card. Please try again.';
			console.error('Error creating card:', error);
		}
	}
</script>

<h1>Create a New Card</h1>

<a href="/app">Back</a>
<form on:submit={handleSubmit}>
	<label for="front">Front of the Card</label><br />
	<textarea id="front" bind:value={front} required></textarea><br />

	<label for="back">Back of the Card</label><br />
	<textarea id="back" bind:value={back} required></textarea><br />

	<input type="submit" value="Create Card" />
</form>

{#if errorMessage}
	<p style="color: red;">{errorMessage}</p>
{/if}
{#if successMessage}
	<p style="color: green;">{successMessage}</p>
{/if}
