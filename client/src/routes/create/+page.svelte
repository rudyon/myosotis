<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let front = '';
	let back = '';
	let errorMessage = '';
	let successMessage = '';
	let isLoggedIn = pb.authStore.isValid;
	let isSubmitting = false;

	async function handleSubmit(event) {
		event.preventDefault();
		errorMessage = '';
		successMessage = '';

		if (!isLoggedIn) {
			errorMessage = 'You must be logged in to create a card.';
			return;
		}

		if (!front.trim() || !back.trim()) {
			errorMessage = 'Please fill out both fields before submitting.';
			return;
		}

		try {
			isSubmitting = true;
			const now = Math.floor(Date.now() / 1000);
			const cardData = {
				front: front.trim(),
				back: back.trim(),
				owner: pb.authStore.model.id,
				lastReview: now,
				interval: 0
			};

			await pb.collection('cards').create(cardData);
			successMessage = 'Card created successfully!';

			// Reset form
			front = '';
			back = '';

			// Automatically redirect after a short delay
			setTimeout(() => {
				goto('/browse');
			}, 1500);
		} catch (error) {
			errorMessage = 'Failed to create card. Please try again.';
			console.error('Error creating card:', error);
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div class="create-container">
	<header>
		<h1 class="gradient-text">Create a New Card</h1>
		<a href="/browse" class="button secondary">Back to Browse</a>
	</header>

	<main>
		<form on:submit={handleSubmit} class="card-form">
			<div class="form-group">
				<label for="front">Front of the Card</label>
				<textarea
					id="front"
					bind:value={front}
					placeholder="Enter the question or prompt"
					rows="4"
					required
				></textarea>
			</div>

			<div class="form-group">
				<label for="back">Back of the Card</label>
				<textarea
					id="back"
					bind:value={back}
					placeholder="Enter the answer or explanation"
					rows="4"
					required
				></textarea>
			</div>

			<button type="submit" class="button primary" disabled={isSubmitting}>
				{#if isSubmitting}
					Creating...
				{:else}
					Create Card
				{/if}
			</button>
		</form>

		{#if errorMessage}
			<div class="alert error">
				<p>{errorMessage}</p>
			</div>
		{/if}
		{#if successMessage}
			<div class="alert success">
				<p>{successMessage}</p>
			</div>
		{/if}
	</main>
</div>

<style>
	.create-container {
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

	.card-form {
		background: var(--card-bg);
		border: 1px solid var(--border-color);
		border-radius: 0.5rem;
		padding: 2rem;
		margin-bottom: 1rem;
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
		border-radius: 0.375rem;
		background-color: rgba(255, 255, 255, 0.05);
		color: var(--text-color);
		font-size: 1rem;
		resize: vertical;
		transition: border-color 0.3s;
	}

	textarea:focus {
		outline: none;
		border-color: var(--accent-color);
	}

	button {
		width: 100%;
	}

	button:disabled {
		opacity: 0.7;
		cursor: not-allowed;
	}

	.alert {
		padding: 1rem;
		border-radius: 0.375rem;
		margin-top: 1rem;
	}

	.alert p {
		margin: 0;
	}

	.alert.error {
		background-color: rgba(239, 68, 68, 0.2);
		border: 1px solid #ef4444;
	}

	.alert.success {
		background-color: rgba(34, 197, 94, 0.2);
		border: 1px solid #22c55e;
	}

	@media (max-width: 600px) {
		header {
			flex-direction: column;
			gap: 1rem;
		}

		header .button {
			width: 100%;
			text-align: center;
		}

		.card-form {
			padding: 1rem;
		}
	}
</style>
