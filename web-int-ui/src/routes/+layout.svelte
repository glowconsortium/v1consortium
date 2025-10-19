<script lang="ts">
	import '../app.css';
	import { toaster } from '$lib/utils/toaster';
	import { Toaster } from '@skeletonlabs/skeleton-svelte';
	import { ErrorBoundary } from '@movsm/v1-consortium-web-pkg';
	import { authStore, authError, authLoading, isAuthenticated } from '@movsm/v1-consortium-web-pkg';
	import { onMount } from 'svelte';
	import { PUBLIC_API_URL } from '$env/static/public';

	let { children } = $props();
	let isInitialized = $state(false);

	// Initialize auth store and application
	onMount(async () => {
		try {
			// Initialize the auth store with API configuration
			await authStore.initialize({
				baseUrl: PUBLIC_API_URL || "http://localhost:8000"
			});
			isInitialized = true;
		} catch (error) {
			console.error('Failed to initialize auth store:', error);
			isInitialized = true; // Still show the app even if auth fails
		}
	});

</script>

<ErrorBoundary>
	<div class="app">
		{#if !isInitialized}
			<!-- Loading state while initializing -->
			<div class="flex items-center justify-center min-h-screen">
				<div class="text-center">
					<div class="animate-spin rounded-full h-32 w-32 border-b-2 border-indigo-500 mx-auto"></div>
					<p class="mt-4 text-lg">Initializing application...</p>
				</div>
			</div>
		{:else if $authError}
			<!-- Auth error state -->
			<div class="flex items-center justify-center min-h-screen">
				<div class="text-center">
					<h1 class="text-2xl font-bold text-red-500 mb-4">Authentication Error</h1>
					<p class="text-lg mb-4">{$authError}</p>
					<button 
						class="btn btn-primary"
						onclick={() => authStore?.clearError()}
					>
						Try Again
					</button>
				</div>
			</div>
		{:else}
			<!-- Main application -->
			{@render children()}
		{/if}
	</div>
</ErrorBoundary>

<!-- Toast notifications -->
<Toaster {toaster} />

<!-- Global loading indicator for auth operations -->
{#if $authLoading}
	<div class="fixed top-4 right-4 z-50">
		<div class="bg-white p-3 rounded-lg shadow-lg flex items-center gap-2">
			<div class="animate-spin rounded-full h-4 w-4 border-b-2 border-indigo-500"></div>
			<span class="text-sm">Authenticating...</span>
		</div>
	</div>
{/if}





