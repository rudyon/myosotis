<script>
	import { pb } from '$lib/pocketbase';
	let email = '';
	let password = '';
	let authData = null;

	async function handleSubmit(event) {
		event.preventDefault();

		try {
			authData = await pb.collection('users').authWithPassword(email, password);
		} catch (error) {
			console.error('Login failed:', error);
			authData = null;
		}
	}
</script>

<form on:submit={handleSubmit}>
	<label for="email">email</label><br />
	<input type="text" id="email" bind:value={email} name="email" /><br />

	<label for="password">password</label><br />
	<input type="password" id="password" bind:value={password} name="password" /><br />

	<input type="submit" value="Submit" />
</form>

{#if authData}
	<p>logged in as {pb.authStore.model.username}</p>
{:else}
	<p>not logged in</p>
{/if}
