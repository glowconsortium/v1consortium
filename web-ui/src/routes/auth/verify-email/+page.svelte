<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { authStore } from '$lib/stores/auth.js';
	import { Button, Alert, Spinner } from '$lib/components/ui/index.js';

	let isVerifying = $state(true);
	let isSuccess = $state(false);
	let error = $state('');
	let token = $state('');

	// Get token from URL parameters and attempt verification
	$effect(() => {
		const urlToken = $page.url.searchParams.get('token');
		if (urlToken) {
			token = urlToken;
			verifyEmail();
		} else {
			// No token, show error
			isVerifying = false;
			error = 'Invalid verification link. Please check your email and try again.';
		}
	});

	async function verifyEmail() {
		if (!token) return;

		isVerifying = true;
		error = '';

		try {
			await authStore.verifyEmail(token);
			isSuccess = true;
		} catch (err) {
			if (err instanceof Error) {
				error = err.message;
			} else {
				error = 'Email verification failed. Please try again.';
			}
		} finally {
			isVerifying = false;
		}
	}

	async function resendVerification() {
		try {
			// For resending, we'll need to ask for email
			goto('/auth/resend-verification');
		} catch (err) {
			console.error('Redirect failed:', err);
		}
	}

	function goToDashboard() {
		goto('/dashboard');
	}

	function goToSignIn() {
		goto('/auth/signin');
	}
</script>

<svelte:head>
	<title>Verify Email - Consortium</title>
	<meta name="description" content="Verify your Consortium account email address." />
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			
			{#if isVerifying}
				<!-- Verifying State -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-blue-100 mb-4">
						<div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
					</div>

					<h2 class="text-2xl font-semibold text-gray-900 mb-4">Verifying Your Email</h2>
					
					<p class="text-gray-600 mb-6">
						Please wait while we verify your email address...
					</p>
				</div>

			{:else if isSuccess}
				<!-- Success State -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-green-100 mb-4">
						<svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
						</svg>
					</div>

					<h2 class="text-2xl font-semibold text-gray-900 mb-4">Email Verified Successfully!</h2>
					
					<p class="text-gray-600 mb-6">
						Your email address has been verified. You can now access all features of your Consortium account.
					</p>

					<div class="bg-green-50 border border-green-200 rounded-md p-4 mb-6">
						<div class="flex">
							<div class="flex-shrink-0">
								<svg class="h-5 w-5 text-green-400" viewBox="0 0 20 20" fill="currentColor">
									<path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
								</svg>
							</div>
							<div class="ml-3">
								<h3 class="text-sm font-medium text-green-800">
									Welcome to Consortium!
								</h3>
								<p class="text-sm text-green-700 mt-1">
									Your account is now fully activated and ready to use.
								</p>
							</div>
						</div>
					</div>

					<div class="space-y-3">
						<Button variant="primary" size="lg" class="w-full" onclick={goToDashboard}>
							Go to Dashboard
						</Button>
						
						<a
							href="/auth/signin"
							class="block w-full text-center bg-white border border-gray-300 rounded-md py-2 px-4 text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
						>
							Sign In
						</a>
					</div>
				</div>

			{:else if error}
				<!-- Error State -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-red-100 mb-4">
						<svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
						</svg>
					</div>

					<h2 class="text-2xl font-semibold text-gray-900 mb-4">Verification Failed</h2>
					
					<div class="mb-6">
						<Alert type="error" message={error} />
					</div>

					<p class="text-sm text-gray-600 mb-6">
						This could happen if:
					</p>
					
					<ul class="text-sm text-gray-600 text-left mb-6 space-y-1">
						<li>• The verification link has expired</li>
						<li>• The link has already been used</li>
						<li>• The link was copied incorrectly</li>
					</ul>

					<div class="space-y-3">
						<button
							onclick={resendVerification}
							class="w-full bg-blue-600 border border-transparent rounded-md py-2 px-4 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
						>
							Request New Verification Email
						</button>

						<a
							href="/auth/signin"
							class="block w-full text-center bg-white border border-gray-300 rounded-md py-2 px-4 text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
						>
							Back to Sign In
						</a>
					</div>
				</div>
			{/if}

			<!-- Help section -->
			<div class="mt-6 pt-6 border-t border-gray-200 text-center">
				<p class="text-xs text-gray-500">
					Need help? <a href="/support" class="text-blue-600 hover:text-blue-500">Contact Support</a>
				</p>
			</div>
		</div>
	</div>
</div>
			console.error('Failed to redirect:', err);
		}
	}
</script>

<svelte:head>
	<title>Verify Email - FormApp</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-surface-50-900-token py-12 px-4 sm:px-6 lg:px-8">
	<div class="max-w-md w-full space-y-8">
		<!-- Header -->
		<div class="text-center">
			<div class="mx-auto w-12 h-12 bg-primary-500 rounded-lg flex items-center justify-center mb-6">
				<span class="text-white font-bold text-xl">F</span>
			</div>
			{#if isVerifying}
				<h2 class="text-3xl font-bold text-on-surface-token">Verifying your email</h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					Please wait while we verify your email address...
				</p>
			{:else if isSuccess}
				<h2 class="text-3xl font-bold text-on-surface-token">Email verified!</h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					Your email has been successfully verified.
				</p>
			{:else}
				<h2 class="text-3xl font-bold text-on-surface-token">Verification failed</h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					We couldn't verify your email address.
				</p>
			{/if}
		</div>

		<!-- Content -->
		<div class="card p-8">
			{#if isVerifying}
				<div class="text-center space-y-4">
					<Spinner size="lg" />
					<p class="text-surface-500-400-token">Verifying your email address...</p>
				</div>
			{:else if isSuccess}
				<div class="text-center space-y-4">
					<Alert type="success" message="Your email has been successfully verified!" />
					
					<div class="space-y-4">
						<p class="text-sm text-surface-500-400-token">
							You can now sign in to your account and start using FormApp.
						</p>
						
						<a href="/auth/signin">
							<Button variant="filled" classes="w-full">
								Sign in to your account
							</Button>
						</a>
						
						<a href="/dashboard">
							<Button variant="ghost" classes="w-full">
								Go to dashboard
							</Button>
						</a>
					</div>
				</div>
			{:else}
				<div class="text-center space-y-4">
					<Alert type="error" message={error} />
					
					<div class="space-y-4">
						<p class="text-sm text-surface-500-400-token">
							This could happen if the verification link has expired or has already been used.
						</p>
						
						<div class="flex flex-col space-y-2">
							<Button variant="filled" onclick={resendVerification}>
								Request new verification email
							</Button>
							
							<a href="/auth/signin">
								<Button variant="ghost" classes="w-full">
									Back to sign in
								</Button>
							</a>
						</div>
					</div>
				</div>
			{/if}
		</div>
	</div>
</div>
