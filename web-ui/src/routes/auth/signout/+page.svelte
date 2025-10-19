<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Spinner } from '@movsm/v1-consortium-web-pkg';

	let isLoggingOut = $state(true);

	onMount(async () => {
		try {
			await authStore.logout();
		} catch (error) {
			console.warn('Logout failed:', error);
		} finally {
			isLoggingOut = false;
			// Redirect to home page after a brief delay
			setTimeout(() => {
				goto('/');
			}, 1500);
		}
	});
</script>

<svelte:head>
	<title>Signing Out - FormApp</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-surface-50-900-token">
	<div class="text-center">
		{#if isLoggingOut}
			<Spinner size="lg" />
			<p class="mt-4 text-lg text-on-surface-token">Signing you out...</p>
		{:else}
			<div class="w-16 h-16 bg-success-500 rounded-full flex items-center justify-center mx-auto mb-4">
				<svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
				</svg>
			</div>
			<h2 class="text-2xl font-bold text-on-surface-token mb-2">You've been signed out</h2>
			<p class="text-surface-500-400-token mb-6">Thank you for using FormApp. Redirecting to home page...</p>
			<a href="/" class="text-primary-500 hover:text-primary-400 font-medium">
				Go to home page
			</a>
		{/if}
	</div>
</div>
