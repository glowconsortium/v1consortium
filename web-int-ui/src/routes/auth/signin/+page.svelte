<script lang="ts">
	import { onMount } from 'svelte';
	import { authStore, isAuthenticated, authLoading, authError } from '@movsm/v1-consortium-web-pkg';
	import { goto } from '$app/navigation';
	import { Button } from '@movsm/v1-consortium-web-pkg';
	import { PUBLIC_AUTH0_AUDIENCE, PUBLIC_AUTH0_CLIENT_ID, PUBLIC_AUTH0_DOMAIN } from '$env/static/public';

	onMount(async () => {
		// Initialize Auth0 if not already done
		if (!$isAuthenticated) {
			try {
				await authStore.initialize({
					domain: PUBLIC_AUTH0_DOMAIN,
					clientId: PUBLIC_AUTH0_CLIENT_ID,
					audience: PUBLIC_AUTH0_AUDIENCE,
					scope: 'openid profile email',
					redirectUri: `${window.location.origin}/auth/callback`
				});
			} catch (error) {
				console.error('Auth initialization failed:', error);
			}
		}

		// Redirect if already authenticated
		if ($isAuthenticated) {
			goto('/dashboard');
		}
	});

	async function handleLogin() {
		try {
			console.log('Starting login process...');
			await authStore.login({
				redirectUri: `${window.location.origin}/auth/callback`
			});
			console.log('Login method completed - user should be redirected to Auth0');
		} catch (error) {
			console.error('Login failed:', error);
		}
	}

	function clearError() {
		authStore.clearError();
	}
</script>

<svelte:head>
	<title>Sign In - Consortium</title>
	<meta name="description" content="Sign in to your Consortium account." />
</svelte:head>

<div class="flex min-h-screen flex-col justify-center bg-gray-50 py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<!-- Logo -->
		<div class="text-center">
			<h1 class="mb-8 text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<!-- Sign in form -->
		<div class="bg-white px-4 py-8 shadow sm:rounded-lg sm:px-10">
			<div class="mb-6 text-center">
				<h2 class="text-2xl font-semibold text-gray-900">Welcome Back</h2>
				<p class="mt-2 text-gray-600">Sign in to your account</p>
			</div>

			<!-- Sign up link -->
			<div class="mt-6 text-center">
				<p class="text-sm text-gray-600">
					Don't have an account?
					<a href="/auth/signup" class="font-medium text-blue-600 hover:text-blue-500">
						Sign up for free
					</a>
				</p>
			</div>

			<div class="mt-8 space-y-6">
				{#if $authError}
					<div class="rounded-md bg-red-50 p-4">
						<div class="flex">
							<div class="flex-shrink-0">
								<svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
									<path
										fill-rule="evenodd"
										d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
										clip-rule="evenodd"
									/>
								</svg>
							</div>
							<div class="ml-3">
								<h3 class="text-sm font-medium text-red-800">Authentication Error</h3>
								<div class="mt-2 text-sm text-red-700">
									<p>{$authError}</p>
								</div>
								<div class="mt-4">
									<Button
										classes="rounded-md bg-red-50 px-2 py-1 text-sm font-medium text-red-800 hover:bg-red-100 focus:ring-2 focus:ring-red-500 focus:ring-offset-2 focus:outline-none"
										onclick={clearError}
									>
										Dismiss
									</Button>
										
								</div>
							</div>
						</div>
					</div>
				{/if}

				<div>
					<Button
						classes="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-3 text-sm font-medium text-white hover:bg-indigo-700 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
						disabled={$authLoading}
						onclick={handleLogin}
					>
						{#if $authLoading}
							<svg
								class="mr-3 -ml-1 h-5 w-5 animate-spin text-white"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
							>
								<circle
									class="opacity-25"
									cx="12"
									cy="12"
									r="10"
									stroke="currentColor"
									stroke-width="4"
								></circle>
								<path
									class="opacity-75"
									fill="currentColor"
									d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
								></path>
							</svg>
							Signing in...
						{:else}
							Sign in with Auth0
						{/if}
					</Button>
				</div>

				<div class="text-center">
					<p class="text-sm text-gray-600">Secure authentication powered by Auth0</p>
				</div>
			</div>

			<!-- Security notice -->
			<div class="mt-6 text-center">
				<p class="flex items-center justify-center text-xs text-gray-500">
					<svg class="mr-1 h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
						></path>
					</svg>
					Your data is protected with enterprise-grade security
				</p>
			</div>
		</div>
	</div>
</div>
