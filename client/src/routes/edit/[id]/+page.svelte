<script>
	import { pb } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	import { error } from '@sveltejs/kit';

	let cardId;
	let frontText = '';
	let backText = '';
	let loading = true;
	let errorMessage = '';
	let successMessage = '';
	let saving = false;

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
		} finally {
			loading = false;
		}
	}

	async function saveCard() {
		if (!cardId) {
			errorMessage = 'Cannot save card, card ID is undefined.';
			return;
		}
		if (!frontText.trim() || !backText.trim()) {
			errorMessage = 'Please fill out both fields before saving.';
			return;
		}

		saving = true;
		errorMessage = '';
		successMessage = '';

		try {
			await pb.collection('cards').update(cardId, {
				front: frontText.trim(),
				back: backText.trim()
			});
			successMessage = 'Card updated successfully!';
			setTimeout(() => {
				window.location.href = '/browse';
			}, 1500);
		} catch (err) {
			console.error('Error updating card:', err);
			errorMessage = 'An error occurred while updating the card. Please try again.';
		} finally {
			saving = false;
		}
	}

	function handleKeyDown(event) {
		if (event.ctrlKey && event.key === 'Enter') {
			saveCard();
		}
	}
</script>

<div class="edit-container">
	<header>
		<h1 class="gradient-text">Edit Card</h1>
		<div class="header-buttons">
			<button class="button secondary" on:click={() => (window.location.href = '/browse')}>
				Cancel
			</button>
			<button class="button" on:click={saveCard} disabled={saving || loading}>
				{#if saving}
					Saving...
				{:else}
					Save Changes
				{/if}
			</button>
		</div>
	</header>

	<main>
		{#if loading}
			<div class="loading">
				<div class="loading-spinner"></div>
				<p>Loading card...</p>
			</div>
		{:else}
			<form on:submit|preventDefault={saveCard}>
				<div class="card-preview">
					<div class="preview-side">
						<h3>Front</h3>
						<div class="card preview">
							{frontText || 'Front preview'}
						</div>
					</div>
					<div class="preview-side">
						<h3>Back</h3>
						<div class="card preview">
							{backText || 'Back preview'}
						</div>
					</div>
				</div>

				<div class="form-group">
					<label for="front">Front Text</label>
					<textarea
						id="front"
						bind:value={frontText}
						rows="4"
						placeholder="Enter the front text of your card"
						required
						on:keydown={handleKeyDown}
					></textarea>
				</div>

				<div class="form-group">
					<label for="back">Back Text</label>
					<textarea
						id="back"
						bind:value={backText}
						rows="4"
						placeholder="Enter the back text of your card"
						required
						on:keydown={handleKeyDown}
					></textarea>
				</div>

				{#if errorMessage}
					<div class="message error">
						{errorMessage}
					</div>
				{/if}

				{#if successMessage}
					<div class="message success">
						{successMessage}
					</div>
				{/if}

				<p class="hint">Tip: Press Ctrl + Enter to quickly save your changes</p>
			</form>
		{/if}
	</main>
</div>

<style>
	.edit-container {
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

	.card-preview {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 1rem;
		margin-bottom: 2rem;
	}

	.preview-side h3 {
		margin-bottom: 0.5rem;
		color: #94a3b8;
		font-size: 0.875rem;
	}

	.card.preview {
		background: var(--card-bg);
		border: 1px solid var(--border-color);
		border-radius: 0.5rem;
		padding: 1rem;
		min-height: 100px;
		display: flex;
		align-items: center;
		justify-content: center;
		text-align: center;
		font-size: 0.875rem;
		color: #94a3b8;
	}

	.form-group {
		margin-bottom: 1.5rem;
	}

	label {
		display: block;
		margin-bottom: 0.5rem;
		font-weight: 500;
	}

	textarea {
		width: 100%;
		padding: 0.75rem;
		border: 1px solid var(--border-color);
		border-radius: 0.5rem;
		background-color: var(--card-bg);
		color: var(--text-color);
		font-family: inherit;
		resize: vertical;
		min-height: 100px;
		transition: border-color 0.3s;
	}

	textarea:focus {
		outline: none;
		border-color: var(--accent-color);
	}

	.message {
		padding: 1rem;
		border-radius: 0.5rem;
		margin-bottom: 1rem;
	}

	.message.error {
		background-color: #fef2f2;
		border: 1px solid #ef4444;
		color: #ef4444;
	}

	.message.success {
		background-color: #f0fdf4;
		border: 1px solid #22c55e;
		color: #22c55e;
	}

	.hint {
		text-align: center;
		color: #94a3b8;
		font-size: 0.875rem;
		margin-top: 1rem;
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
		}

		.header-buttons {
			width: 100%;
		}

		.header-buttons button {
			flex: 1;
		}

		.card-preview {
			grid-template-columns: 1fr;
		}
	}
</style>
