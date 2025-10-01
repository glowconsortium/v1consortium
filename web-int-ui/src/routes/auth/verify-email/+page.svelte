<script lang="ts">
	// import { onMount } from 'svelte';
	// import { goto } from '$app/navigation';
	// import { page } from '$app/stores';
	// import { authStore } from '@movsm/v1-consortium-web-pkg';
	// import { Button, Alert, Spinner } from '@movsm/v1-consortium-web-pkg';

	// let isVerifying = $state(true);
	// let isSuccess = $state(false);
	// let error = $state('');
	// let token = $state('');

	// // Get token from URL parameters and attempt verification
	// $effect(() => {
	// 	const urlToken = $page.url.searchParams.get('token');
	// 	if (urlToken) {
	// 		token = urlToken;
	// 		verifyEmail();
	// 	} else {
	// 		// No token, show error
	// 		isVerifying = false;
	// 		error = 'Invalid verification link. Please check your email and try again.';
	// 	}
	// });

	// async function verifyEmail() {
	// 	if (!token) return;

	// 	isVerifying = true;
	// 	error = '';

	// 	try {
	// 		await authStore.verifyEmail(token);
	// 		isSuccess = true;
	// 	} catch (err) {
	// 		if (err instanceof Error) {
	// 			error = err.message;
	// 		} else {
	// 			error = 'Email verification failed. Please try again.';
	// 		}
	// 	} finally {
	// 		isVerifying = false;
	// 	}
	// }

	// async function resendVerification() {
	// 	try {
	// 		// For resending, we'll need to ask for email
	// 		goto('/auth/resend-verification');
	// 	} catch (err) {
	// 		console.error('Redirect failed:', err);
	// 	}
	// }

	// function goToDashboard() {
	// 	goto('/dashboard');
	// }

	// function goToSignIn() {
	// 	goto('/auth/signin');
	// }
</script>



<svelte:head>
	<title>Verify Email - FormApp</title>
</svelte:head>
<!-- 
<div class="min-h-screen flex items-center justify-center bg-surface-50-900-token py-12 px-4 sm:px-6 lg:px-8">
	<div class="max-w-md w-full space-y-8">
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
</div> -->
