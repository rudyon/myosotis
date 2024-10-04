<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';
	let errorMessage = '';
	let isLoggedIn = false;

	if (pb.authStore.isValid) {
		isLoggedIn = true;
	}

	async function handleLogin(event) {
		event.preventDefault();

		try {
			const authData = await pb.collection('users').authWithPassword(email, password);
			goto('/app');
		} catch (error) {
			errorMessage = 'Login failed. Please try again.';
			console.error('Error during login:', error);
		}
	}
</script>

{#if isLoggedIn}
	<p>You are already logged in! Click below to go to the app:</p>
	<a href="/app">Go to App</a>
{:else}
	<form on:submit={handleLogin}>
		<label for="email">Email:</label><br />
		<input type="email" id="email" bind:value={email} required /><br />

		<label for="password">Password:</label><br />
		<input type="password" id="password" bind:value={password} required /><br />

		<input type="submit" value="Login" />
	</form>

	{#if errorMessage}
		<p style="color: red;">{errorMessage}</p>
	{/if}
{/if}
