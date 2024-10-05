<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let front = '';
	let back = '';
	let errorMessage = '';
	let successMessage = '';
	let isLoggedIn = false;

	async function handleSubmit(event) {
		event.preventDefault();

		try {
			const cardData = {
				front: front,
				back: back,
				owner: pb.authStore.model.id
			};

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
