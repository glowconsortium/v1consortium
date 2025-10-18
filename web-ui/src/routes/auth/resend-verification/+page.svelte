<script lang="ts">
	import { goto } from '$app/navigation';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Button, Input, Alert } from '@movsm/v1-consortium-web-pkg';
	import { validateEmail, validateRequired } from '@movsm/v1-consortium-web-pkg';

	let email = $state('');
	let errors = $state<Record<string, string>>({});
	let isSubmitting = $state(false);
	let isSubmitted = $state(false);

	// Redirect if already authenticated
	$effect(() => {
		if ($authStore.isAuthenticated) {
			goto('/dashboard');
		}
	});

	function validateForm(): boolean {
		const newErrors: Record<string, string> = {};

		// Email validation
		const emailError = validateRequired(email, 'Email');
		if (emailError) {
			newErrors.email = emailError;
		} else if (!validateEmail(email)) {
			newErrors.email = 'Please enter a valid email address';
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (!validateForm()) return;

		isSubmitting = true;
		authStore.clearError();

		try {
			await authStore.resendVerification(email);
			isSubmitted = true;
		} catch (error) {
			// Error is handled by the store
		} finally {
			isSubmitting = false;
		}
	}

	function handleEmailInput(event: Event) {
		email = (event.target as HTMLInputElement).value;
		if (errors.email) {
			errors = { ...errors, email: '' };
		}
	}
</script>

<svelte:head>
	<title>Resend Verification - Consortium</title>
	<meta name="description" content="Resend verification email for your Consortium account." />
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			
			{#if emailSent}
				<!-- Email Sent Confirmation -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-blue-100 mb-4">
						<svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
						</svg>
					</div>

					<h2 class="text-2xl font-semibold text-gray-900 mb-4">Verification Email Sent</h2>
					
					<p class="text-gray-600 mb-2">We've sent a new verification email to:</p>
					<p class="font-medium text-gray-900 mb-4">{email}</p>
					
					<p class="text-sm text-gray-600 mb-6">
						Please check your email and click the verification link to activate your account.
						If you don't see it, check your spam folder.
					</p>

					<div class="bg-yellow-50 border border-yellow-200 rounded-md p-3 mb-6">
						<p class="text-sm text-yellow-800">
							The verification link will expire in 24 hours for security.
						</p>
					</div>

					<div class="space-y-3">
						<button
							onclick={resendEmail}
							class="w-full bg-white border border-gray-300 rounded-md py-2 px-4 text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
							disabled={isSubmitting}
						>
							{isSubmitting ? 'Sending...' : "Didn't receive it? Send again"}
						</button>

						<a
							href="/auth/signin"
							class="block w-full text-center bg-blue-600 border border-transparent rounded-md py-2 px-4 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
						>
							Back to Sign In
						</a>
					</div>
				</div>

			{:else}
				<!-- Resend Verification Form -->
				<div class="mb-6">
					<a href="/auth/signin" class="flex items-center text-sm text-gray-600 hover:text-gray-900 mb-6">
						<svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
						</svg>
						Back to Sign In
					</a>

					<div class="text-center">
						<h2 class="text-2xl font-semibold text-gray-900">Resend Verification Email</h2>
						<p class="text-gray-600 mt-2">Enter your email to receive a new verification link</p>
					</div>
				</div>

				<form on:submit|preventDefault={handleSubmit} class="space-y-6">
					<div>
						<label for="email" class="block text-sm font-medium text-gray-700 mb-1">
							Email Address
						</label>
						<Input
							id="email"
							type="email"
							value={email}
							oninput={handleFieldInput()}
							placeholder="john@abctrucking.com"
							error={errors.email}
							required
						/>
					</div>

					<!-- Error alert -->
					{#if $authStore.error}
						<div>
							<Alert type="error" message={$authStore.error} />
						</div>
					{/if}

					<div>
						<Button
							type="submit"
							color="primary"
							size="lg"
							classes="w-full"
							disabled={isSubmitting}
						>
							{isSubmitting ? 'Sending Verification Email...' : 'Send Verification Email'}
						</Button>
					</div>

					<div class="text-center">
						<p class="text-sm text-gray-600">
							We'll send you a secure link to verify your email address.
							The link will expire in 24 hours for security.
						</p>
					</div>

					<div class="text-center pt-4 border-t border-gray-200">
						<p class="text-sm text-gray-600">
							Already verified your email?
							<a href="/auth/signin" class="font-medium text-blue-600 hover:text-blue-500">
								Sign In
							</a>
						</p>
					</div>
				</form>
			{/if}
		</div>
	</div>
</div>