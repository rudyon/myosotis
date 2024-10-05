<script>
	import { pb } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	import { error } from '@sveltejs/kit';

	let cardId;
	let frontText = '';
	let backText = '';
	let loading = true;
	let errorMessage = '';

	onMount(async () => {
		cardId = window.location.pathname.split('/').pop();
		if (!cardId) {
			console.error('Card ID is not defined.');
			throw error(404, 'Card not found');
		}
		await loadCard(cardId);
	});

	async function loadCard(id) {
		try {
			const card = await pb.collection('cards').getOne(id);
			frontText = card.front;
			backText = card.back;
		} catch (err) {
			console.error('Error loading card:', err);
			errorMessage = 'Card not found or unable to load card';
			return;
		} finally {
			loading = false;
		}
	}

	async function saveCard() {
		if (!cardId) {
			console.error('Cannot save card, card ID is undefined.');
			return;
		}
		if (!frontText || !backText) {
			errorMessage = 'Please fill out both fields before saving.';
			return;
		}
		try {
			const response = await pb.collection('cards').update(cardId, {
				front: frontText,
				back: backText
			});
			console.log('Update Response:', response);
			alert('Card updated successfully!');
			window.location.href = '/browse';
		} catch (err) {
			console.error('Error updating card:', err);
			errorMessage = 'An error occurred while updating the card. Please try again.'; // Set error message
		}
	}
</script>

<h1>Edit Card</h1>

{#if loading}
	<p>Loading card...</p>
{:else}
	{#if errorMessage}
		<p style="color: red;">{errorMessage}</p>
	{/if}

	<label for="front">Front Text:</label>
	<textarea id="front" bind:value={frontText} rows="4" required></textarea>

	<label for="back">Back Text:</label>
	<textarea id="back" bind:value={backText} rows="4" required></textarea>

	<button on:click={saveCard}>Save Changes</button>
	<button on:click={() => (window.location.href = '/app')}>Cancel</button>
{/if}

<style>
	input,
	textarea {
		display: block;
		margin-bottom: 10px;
		width: 100%;
	}
	textarea {
		resize: none;
	}
</style>
