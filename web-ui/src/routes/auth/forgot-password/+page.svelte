<script lang="ts">
	import { goto } from '$app/navigation';
	import { authStore } from '$lib/stores/auth.js';
	import { Button, Input, Alert } from '$lib/components/ui/index.js';
	import { validateEmail, validateRequired } from '$lib/utils/validation.js';

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
			await authStore.forgotPassword(email);
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
	<title>Forgot Password - FormApp</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-surface-50-900-token py-12 px-4 sm:px-6 lg:px-8">
	<div class="max-w-md w-full space-y-8">
		<!-- Header -->
		<div class="text-center">
			<div class="mx-auto w-12 h-12 bg-primary-500 rounded-lg flex items-center justify-center mb-6">
				<span class="text-white font-bold text-xl">F</span>
			</div>
			{#if isSubmitted}
				<h2 class="text-3xl font-bold text-on-surface-token">Check your email</h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					We've sent a password reset link to {email}
				</p>
			{:else}
				<h2 class="text-3xl font-bold text-on-surface-token"><script lang="ts">
	import { goto } from '$app/navigation';
	import { authStore } from '$lib/stores/auth.js';
	import { Button, Input, Alert } from '$lib/components/ui/index.js';
	import { validateEmail, validateRequired } from '$lib/utils/validation.js';

	let email = $state('');
	let errors = $state<Record<string, string>>({});
	let isSubmitting = $state(false);
	let emailSent = $state(false);

	function validateForm(): boolean {
		const newErrors: Record<string, string> = {};

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
			await authStore.resetPassword(email);
			emailSent = true;
		} catch (error) {
			// Error is handled by the store
		} finally {
			isSubmitting = false;
		}
	}

	function handleFieldInput() {
		return (event: Event) => {
			email = (event.target as HTMLInputElement).value;
			// Clear email error
			if (errors.email) {
				errors = { ...errors, email: '' };
			}
		};
	}

	function resendEmail() {
		emailSent = false;
		handleSubmit();
	}
</script>

<svelte:head>
	<title>Reset Password - Consortium</title>
	<meta name="description" content="Reset your Consortium account password." />
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

					<h2 class="text-2xl font-semibold text-gray-900 mb-4">Check Your Email</h2>
					
					<p class="text-gray-600 mb-2">We've sent a password reset link to:</p>
					<p class="font-medium text-gray-900 mb-4">{email}</p>
					
					<p class="text-sm text-gray-600 mb-6">
						Click the link in your email to reset your password.
						If you don't see it, check your spam folder.
					</p>

					<div class="bg-yellow-50 border border-yellow-200 rounded-md p-3 mb-6">
						<p class="text-sm text-yellow-800">
							The link will expire in 24 hours for security.
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
				<!-- Password Reset Form -->
				<div class="mb-6">
					<a href="/auth/signin" class="flex items-center text-sm text-gray-600 hover:text-gray-900 mb-6">
						<svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
						</svg>
						Back to Sign In
					</a>

					<div class="text-center">
						<h2 class="text-2xl font-semibold text-gray-900">Reset Your Password</h2>
						<p class="text-gray-600 mt-2">Enter your email to receive a reset link</p>
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
							variant="primary"
							size="lg"
							class="w-full"
							disabled={isSubmitting}
						>
							{isSubmitting ? 'Sending Reset Link...' : 'Send Reset Link'}
						</Button>
					</div>

					<div class="text-center">
						<p class="text-sm text-gray-600">
							We'll send you a secure link to reset your password.
							The link will expire in 24 hours for security.
						</p>
					</div>

					<div class="text-center pt-4 border-t border-gray-200">
						<p class="text-sm text-gray-600">
							Remember your password?
							<a href="/auth/signin" class="font-medium text-blue-600 hover:text-blue-500">
								Sign In
							</a>
						</p>
					</div>
				</form>
			{/if}
		</div>
	</div>
</div></h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					Enter your email address and we'll send you a link to reset your password.
				</p>
			{/if}
		</div>

		<!-- Form or success message -->
		<div class="card p-8">
			{#if isSubmitted}
				<div class="text-center space-y-4">
					<Alert type="success" message="Password reset email sent successfully!" />
					
					<div class="space-y-4">
						<p class="text-sm text-surface-500-400-token">
							Check your email and click the link to reset your password. 
							If you don't see the email, check your spam folder.
						</p>
						
						<div class="flex flex-col space-y-2">
							<Button variant="filled" onclick={() => { isSubmitted = false; email = ''; }}>
								Send another email
							</Button>
							
							<a href="/auth/signin">
								<Button variant="ghost" classes="w-full">
									Back to sign in
								</Button>
							</a>
						</div>
					</div>
				</div>
			{:else}
				{#if $authStore.error}
					<div class="mb-6">
						<Alert type="error" message={$authStore.error} />
					</div>
				{/if}

				<form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="space-y-6">
					<Input
						type="email"
						label="Email address"
						placeholder="Enter your email"
						value={email}
						required
						error={errors.email}
						oninput={handleEmailInput}
					/>

					<Button
						type="submit"
						variant="filled"
						size="lg"
						classes="w-full"
						loading={isSubmitting}
						disabled={isSubmitting}
					>
						{isSubmitting ? 'Sending...' : 'Send reset link'}
					</Button>
				</form>

				<div class="mt-6 text-center">
					<a href="/auth/signin" class="text-sm font-medium text-primary-500 hover:text-primary-400">
						← Back to sign in
					</a>
				</div>
			{/if}
		</div>
	</div>
</div>
