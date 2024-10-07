<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';
	let key = ''; // Added key field
	let errorMessage = '';
	let successMessage = '';
	let isLoggedIn = false;
	let isLoading = false;
	let isRegisterMode = false;

	if (pb.authStore.isValid) {
		isLoggedIn = true;
	}

	async function handleLogin() {
		try {
			isLoading = true;
			const authData = await pb.collection('users').authWithPassword(email, password);
			goto('/app');
		} catch (error) {
			errorMessage = 'Login failed. Please check your credentials and try again.';
			console.error('Error during login:', error);
		} finally {
			isLoading = false;
		}
	}

	async function handleRegister() {
		try {
			isLoading = true;

			// Check if the key exists in the 'keys' collection and if it is unused
			const keyRecords = await pb.collection('keys').getList(1, 1, {
				filter: `key = "${key}"`
			});

			if (keyRecords.items.length === 0) {
				throw new Error('Invalid key. Please try again.');
			}

			const keyRecord = keyRecords.items[0];

			if (keyRecord.used) {
				throw new Error('This key has already been used.');
			}

			const userData = {
				email: email,
				password: password,
				passwordConfirm: password
			};

			await pb.collection('users').create(userData);

			// Mark the key as used
			await pb.collection('keys').update(keyRecord.id, { used: true });

			successMessage = 'Registration successful! You can now log in.';
			isRegisterMode = false;
		} catch (error) {
			errorMessage = error.message || 'Registration failed. Please try again.';
			console.error('Error during registration:', error);
		} finally {
			isLoading = false;
		}
	}

	async function handleSubmit(event) {
		event.preventDefault();
		errorMessage = '';
		successMessage = '';

		if (isRegisterMode) {
			await handleRegister();
		} else {
			await handleLogin();
		}
	}
</script>

<div class="auth-container">
	{#if isLoggedIn}
		<div class="card text-center">
			<h2>Welcome Back! 👋</h2>
			<p>You're already logged in. Ready to continue your learning journey?</p>
			<a href="/app" class="button">Go to App</a>
		</div>
	{:else}
		<div class="auth-card card">
			<h2 class="gradient-text">{isRegisterMode ? 'Create Account' : 'Welcome Back'}</h2>
			<p class="auth-subtitle">
				{isRegisterMode
					? 'Start your learning journey today!'
					: 'Sign in to continue your progress'}
			</p>

			<form on:submit={handleSubmit}>
				<div class="form-group">
					<label for="email">Email</label>
					<input type="email" id="email" bind:value={email} required placeholder="your@email.com" />
				</div>

				<div class="form-group">
					<label for="password">Password</label>
					<input
						type="password"
						id="password"
						bind:value={password}
						required
						placeholder="••••••••"
					/>
				</div>

				{#if isRegisterMode}
					<div class="form-group">
						<label for="key">Key</label>
						<input type="text" id="key" bind:value={key} required placeholder="Enter your key" />
					</div>
				{/if}

				<button type="submit" class="button-primary" disabled={isLoading}>
					{#if isLoading}
						Loading...
					{:else}
						{isRegisterMode ? 'Create Account' : 'Sign In'}
					{/if}
				</button>
			</form>

			{#if errorMessage}
				<p class="error-message">{errorMessage}</p>
			{/if}
			{#if successMessage}
				<p class="success-message">{successMessage}</p>
			{/if}

			<div class="auth-footer">
				<button class="button-text" on:click={() => (isRegisterMode = !isRegisterMode)}>
					{isRegisterMode ? 'Already have an account? Sign In' : 'Need an account? Register'}
				</button>
			</div>
		</div>
	{/if}
</div>

<style>
	.auth-container {
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 100vh;
		padding: 1rem;
	}

	.auth-card {
		width: 100%;
		max-width: 400px;
	}

	.auth-card h2 {
		margin-bottom: 0.5rem;
		font-size: 1.75rem;
	}

	.auth-subtitle {
		color: #94a3b8;
		margin-bottom: 1.5rem;
	}

	.form-group {
		margin-bottom: 1rem;
	}

	.form-group label {
		display: block;
		margin-bottom: 0.5rem;
		font-weight: 500;
	}

	input {
		width: 100%;
		padding: 0.75rem;
		border: 1px solid var(--border-color);
		border-radius: 6px;
		background-color: rgba(255, 255, 255, 0.05);
		color: var(--text-color);
		transition: border-color 0.3s;
	}

	input:focus {
		outline: none;
		border-color: var(--accent-color);
	}

	.button-primary {
		width: 100%;
		background: linear-gradient(45deg, var(--accent-color), var(--accent-color-2));
		color: white;
		font-weight: 500;
		margin-top: 1rem;
	}

	.button-primary:disabled {
		opacity: 0.7;
		cursor: not-allowed;
	}

	.button-text {
		background: none;
		border: none;
		color: var(--accent-color);
		padding: 0;
		font-size: 0.875rem;
	}

	.button-text:hover {
		color: var(--accent-color-2);
		background: none;
		border: none;
	}

	.auth-footer {
		margin-top: 1.5rem;
		text-align: center;
	}

	.error-message {
		color: #ef4444;
		margin-top: 1rem;
		font-size: 0.875rem;
	}

	.success-message {
		color: #22c55e;
		margin-top: 1rem;
		font-size: 0.875rem;
	}

	.text-center {
		text-align: center;
	}
</style>
