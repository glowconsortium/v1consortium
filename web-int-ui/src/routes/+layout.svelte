<script lang="ts">
	import { createAuthStore, setAuthStore, isAuthenticated, currentUser, authError, authLoading, type AuthStore } from '$lib/stores/authStore';
	import '../app.css';
	import { toaster } from '$lib/utils/toaster';
	import { Toaster } from '@skeletonlabs/skeleton-svelte';
	import { ErrorBoundary } from '@movsm/v1-consortium-web-pkg';
	import { browser } from '$app/environment';
	import { setContext, onMount } from 'svelte';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { createConnectTransport } from "@connectrpc/connect-web";

	let { children } = $props();
	let isInitialized = $state(false);
	let authStoreInstance = $state<AuthStore | null>(null);

	// Create the main transport
	const maintransport = createConnectTransport({
		baseUrl: PUBLIC_API_URL || "http://localhost:8000",
		useBinaryFormat: true,
		// Add interceptors for authentication if needed
		interceptors: [
			(next) => async (req) => {
				// Add auth token to requests if available
				if (authStoreInstance) {
					const token = authStoreInstance.getIdToken();
					if (token) {
						req.header.set('Authorization', `Bearer ${token}`);
					}
				}
				return await next(req);
			}
		]
	});

	setContext("maintransport", maintransport);

	// Initialize auth store and application
	onMount(() => {
		if (browser && !isInitialized) {
			try {
				// Create auth store instance with the transport
				authStoreInstance = createAuthStore(maintransport, {
					// Optional configuration
					autoRefresh: true,
					refreshThreshold: 5 * 60 * 1000, // Refresh token 5 minutes before expiry
				});

				// Set the global auth store instance for derived stores
				setAuthStore(authStoreInstance);
				
				// Also set auth store in context for components to access
				setContext("authStore", authStoreInstance);
				
				isInitialized = true;
			} catch (error) {
				console.error('Failed to initialize auth store:', error);
				toaster.create({ type: 'error', title: 'Failed to initialize authentication' });
			}
		}
	});

	// Auto-refresh token when it's about to expire
	$effect(() => {
		if (authStoreInstance && $isAuthenticated && browser) {
			const interval = setInterval(async () => {
				try {
					if (authStoreInstance) {
						await authStoreInstance.refreshToken();
					}
				} catch (error) {
					console.warn('Token refresh failed:', error);
					// Token refresh failed, user might need to log in again
				}
			}, 15 * 60 * 1000); // Check every 15 minutes

			return () => clearInterval(interval);
		}
	});
</script>

<ErrorBoundary>
	<div class="app">
		{#if !isInitialized}
			<!-- Loading state while initializing -->
			<div class="flex items-center justify-center min-h-screen">
				<div class="text-center">
					<div class="animate-spin rounded-full h-32 w-32 border-b-2 border-primary-500 mx-auto"></div>
					<p class="mt-4 text-lg">Initializing application...</p>
				</div>
			</div>
		{:else if $authError}
			<!-- Auth error state -->
			<div class="flex items-center justify-center min-h-screen">
				<div class="text-center">
					<h1 class="text-2xl font-bold text-error-500 mb-4">Authentication Error</h1>
					<p class="text-lg mb-4">{$authError}</p>
					<button 
						class="btn variant-filled-primary"
						onclick={() => authStoreInstance?.clearError()}
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
		<div class="bg-surface-100-800-token p-3 rounded-lg shadow-lg flex items-center gap-2">
			<div class="animate-spin rounded-full h-4 w-4 border-b-2 border-primary-500"></div>
			<span class="text-sm">Authenticating...</span>
		</div>
	</div>
{/if}





