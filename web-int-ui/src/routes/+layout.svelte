<script lang="ts">
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import '../app.css';
	import {toaster}  from '@movsm/v1-consortium-web-pkg';
	import {Toaster} from '@skeletonlabs/skeleton-svelte'
	import { ErrorBoundary } from '@movsm/v1-consortium-web-pkg';
	import { browser } from '$app/environment';
	import { PUBLIC_AUTH0_CLIENT_ID, PUBLIC_AUTH0_DOMAIN,PUBLIC_AUTH0_AUDIENCE } from '$env/static/public';

	let { children } = $props();
	let isInitialized = $state(false);

	// Initialize API client and auth store only once
	$effect(() => {
		if (browser && !isInitialized) {
			const initializeApp = async () => {
				// Set API client base URL
				// Initialize auth store
					await authStore.initialize({
					domain: PUBLIC_AUTH0_DOMAIN , // Replace with your Auth0 domain
					clientId: PUBLIC_AUTH0_CLIENT_ID , // Replace with your Auth0 client ID
					audience: PUBLIC_AUTH0_AUDIENCE , // Replace with your API audience (optional)
					scope: 'openid profile email',
					redirectUri: `${window.location.origin}/auth/callback`
				});
				
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





