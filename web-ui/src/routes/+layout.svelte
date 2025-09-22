<script lang="ts">
	import { onMount } from 'svelte';
	import { authStore } from '$lib/stores/auth.js';
	import { apiClient } from '$lib/api/client.js';
	import '../app.css';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { toaster } from '$lib/utils/toaster.js';
	import {Toaster} from '@skeletonlabs/skeleton-svelte'
	import { ErrorBoundary } from '$lib/components/ui';
	import { browser } from '$app/environment';

	let { children } = $props();
	let isInitialized = $state(false);

	// Initialize API client and auth store only once
	$effect(() => {
		if (browser && !isInitialized) {
			const initializeApp = async () => {
				// Set API client base URL
				// Initialize auth store
				await authStore.initialize();
				
				isInitialized = true;
			};
			
			initializeApp();
		}
	});
</script>

<ErrorBoundary>
	<div class="app">
		{@render children()}
	</div>
</ErrorBoundary>

<Toaster {toaster} />





