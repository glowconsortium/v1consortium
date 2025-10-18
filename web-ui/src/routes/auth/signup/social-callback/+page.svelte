<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { signupStore } from '$lib/store/signupStore';
	import { Button } from '@movsm/v1-consortium-web-pkg';

	let isLoading = $state(true);
	let error = $state<string | null>(null);

	onMount(async () => {
		try {
			// Extract OAuth parameters from URL
			const urlParams = $page.url.searchParams;
			const code = urlParams.get('code');
			const state = urlParams.get('state');
			const provider = urlParams.get('provider') || 'google'; // default to google
			const errorParam = urlParams.get('error');

			if (errorParam) {
				throw new Error(`OAuth error: ${errorParam}`);
			}

			if (!code) {
				throw new Error('No authorization code received');
			}

			// Handle the social callback and get workflow ID
			const workflowId = await signupStore.handleSocialCallback(provider, code);
			
			// Redirect to email verification with the workflow ID
			goto(`/auth/signup/${workflowId}/verify-email`);
		} catch (err) {
			console.error('Social callback error:', err);
			error = err instanceof Error ? err.message : 'Social signup failed';
			isLoading = false;
		}
	});

	function handleRetry() {
		goto('/auth/signup');
	}
</script>

<svelte:head>
	<title>Social Sign Up - Consortium</title>
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			{#if isLoading}
				<!-- Loading state -->
				<div class="text-center">
					<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto mb-4"></div>
					<h2 class="text-xl font-semibold text-gray-900 mb-2">Setting up your account...</h2>
					<p class="text-gray-600">Please wait while we complete your social signup.</p>
				</div>
			{:else if error}
				<!-- Error state -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-red-100 mb-4">
						<svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
						</svg>
					</div>
					<h2 class="text-xl font-semibold text-gray-900 mb-2">Social Signup Failed</h2>
					<p class="text-gray-600 mb-4">{error}</p>
					
					<div class="space-y-3">
						<Button 
							color="primary" 
							size="lg" 
							classes="w-full"
							onclick={handleRetry}
						>
							Try Again
						</Button>
						<a 
							href="/auth/signin" 
							class="block text-center text-blue-600 hover:text-blue-500 text-sm font-medium"
						>
							Back to Sign In
						</a>
					</div>
				</div>
			{/if}
		</div>
	</div>
</div>
