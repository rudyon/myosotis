import { redirect } from '@sveltejs/kit';
import { pb } from '$lib/pocketbase';

export async function load() {
	if (!pb.authStore.isValid) {
		throw redirect(302, '/login');
	}

	return {};
}
