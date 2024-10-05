<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';
	let errorMessage = '';
	let successMessage = '';
	let isLoggedIn = false;

	if (pb.authStore.isValid) {
		isLoggedIn = true;
	}

	async function handleLogin() {
		try {
			const authData = await pb.collection('users').authWithPassword(email, password);
			goto('/app');
		} catch (error) {
			errorMessage = 'Login failed. Please try again.';
			console.error('Error during login:', error);
		}
	}

	async function handleRegister() {
		try {
			const userData = {
				email: email,
				password: password,
				passwordConfirm: password
			};

			await pb.collection('users').create(userData);
			successMessage = 'Registration successful! You can now log in.';
		} catch (error) {
			errorMessage = 'Registration failed. Please try again.';
			console.error('Error during registration:', error);
		}
	}

	// Handle form submission
	async function handleSubmit(event, action) {
		event.preventDefault();
		errorMessage = '';
		successMessage = '';

		if (action === 'login') {
			await handleLogin();
		} else if (action === 'register') {
			await handleRegister();
		}
	}
</script>

{#if isLoggedIn}
	<p>You are already logged in! Click below to go to the app:</p>
	<a href="/app">Go to App</a>
{:else}
	<form>
		<label for="email">Email</label><br />
		<input type="email" id="email" bind:value={email} required /><br />

		<label for="password">Password</label><br />
		<input type="password" id="password" bind:value={password} required /><br />

		<button on:click={(e) => handleSubmit(e, 'login')}>Login</button>
		<button on:click={(e) => handleSubmit(e, 'register')}>Register</button>
	</form>

	{#if errorMessage}
		<p style="color: red;">{errorMessage}</p>
	{/if}
	{#if successMessage}
		<p style="color: green;">{successMessage}</p>
	{/if}
{/if}
