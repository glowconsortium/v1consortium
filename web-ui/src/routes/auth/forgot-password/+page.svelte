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
			//await authStore.forgotPassword(email);
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
	<title>Reset Password - Consortium</title>
	<meta name="description" content="Reset your Consortium account password." />
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">


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
