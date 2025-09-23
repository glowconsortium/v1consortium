<script lang="ts">
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import '../app.css';
	import { toaster } from '@movsm/v1-consortium-web-pkg';
	import {Toaster} from '@skeletonlabs/skeleton-svelte'
	import { ErrorBoundary } from '@movsm/v1-consortium-web-pkg';
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





