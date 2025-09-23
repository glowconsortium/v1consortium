<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Button, Input, Alert } from '@movsm/v1-consortium-web-pkg';
	import { validatePassword, validateRequired, validateConfirmPassword } from '@movsm/v1-consortium-web-pkg';

	let password = $state('');
	let confirmPassword = $state('');
	let errors = $state<Record<string, string>>({});
	let isSubmitting = $state(false);
	let isSuccess = $state(false);
	let token = $state('');
	let showPassword = $state(false);

	// Get token from URL parameters
	$effect(() => {
		const urlToken = $page.url.searchParams.get('token');
		if (urlToken) {
			token = urlToken;
		} else {
			// No token, redirect to forgot password
			goto('/auth/forgot-password');
		}
	});

	// Redirect if already authenticated
	$effect(() => {
		if ($authStore.isAuthenticated) {
			goto('/dashboard');
		}
	});

	function validateForm(): boolean {
		const newErrors: Record<string, string> = {};

		// Password validation
		const passwordError = validateRequired(password, 'Password');
		if (passwordError) {
			newErrors.password = passwordError;
		} else {
			const passwordValidation = validatePassword(password);
			if (!passwordValidation.isValid) {
				newErrors.password = passwordValidation.errors[0];
			}
		}

		// Confirm password validation
		const confirmError = validateConfirmPassword(password, confirmPassword);
		if (confirmError) {
			newErrors.confirmPassword = confirmError;
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (!validateForm()) return;

		isSubmitting = true;
		authStore.clearError();

		try {
			// Here you would call the reset password with token API
			// For now, using the basic reset password method
			await authStore.resetPassword(password);
			isSuccess = true;
		} catch (error) {
			// Error is handled by the store
		} finally {
			isSubmitting = false;
		}
	}

	function handleFieldInput(field: string) {
		return (event: Event) => {
			const value = (event.target as HTMLInputElement).value;
			
			if (field === 'password') {
				password = value;
			} else if (field === 'confirmPassword') {
				confirmPassword = value;
			}

			// Clear specific field error
			if (errors[field]) {
				errors = { ...errors, [field]: '' };
			}
		};
	}

	function togglePasswordVisibility() {
		showPassword = !showPassword;
	}
		function handlePasswordInput(event: Event) {
		password = (event.target as HTMLInputElement).value;
		if (errors.password) {
			errors = { ...errors, password: '' };
		}
	}

	function handleConfirmPasswordInput(event: Event) {
		confirmPassword = (event.target as HTMLInputElement).value;
		if (errors.confirmPassword) {
			errors = { ...errors, confirmPassword: '' };
		}
	}
</script>



<svelte:head>
	<title>Reset Password - FormApp</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-surface-50-900-token py-12 px-4 sm:px-6 lg:px-8">
	<div class="max-w-md w-full space-y-8">
		<!-- Header -->
		<div class="text-center">
			<div class="mx-auto w-12 h-12 bg-primary-500 rounded-lg flex items-center justify-center mb-6">
				<span class="text-white font-bold text-xl">F</span>
			</div>
			{#if isSuccess}
				<h2 class="text-3xl font-bold text-on-surface-token">Password reset successful</h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					Your password has been updated successfully.
				</p>
			{:else}
				<h2 class="text-3xl font-bold text-on-surface-token">Reset your password</h2>
				<p class="mt-2 text-sm text-surface-500-400-token">
					Enter your new password below.
				</p>
			{/if}
		</div>

		<!-- Form or success message -->
		<div class="card p-8">
			{#if isSuccess}
				<div class="text-center space-y-4">
					<Alert type="success" message="Your password has been successfully reset!" />
					
					<div class="space-y-4">
						<p class="text-sm text-surface-500-400-token">
							You can now sign in with your new password.
						</p>
						
						<a href="/auth/signin">
							<Button variant="filled" classes="w-full">
								Sign in to your account
							</Button>
						</a>
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
						type="password"
						label="New password"
						placeholder="Enter your new password"
						value={password}
						required
						error={errors.password}
						oninput={handlePasswordInput}
					/>

					<Input
						type="password"
						label="Confirm new password"
						placeholder="Confirm your new password"
						value={confirmPassword}
						required
						error={errors.confirmPassword}
						oninput={handleConfirmPasswordInput}
					/>

					<div class="text-xs text-surface-500-400-token space-y-1">
						<p>Password requirements:</p>
						<ul class="list-disc list-inside space-y-1">
							<li>At least 8 characters long</li>
							<li>Contains uppercase and lowercase letters</li>
							<li>Contains at least one number</li>
						</ul>
					</div>

					<Button
						type="submit"
						variant="filled"
						size="lg"
						classes="w-full"
						loading={isSubmitting}
						disabled={isSubmitting}
					>
						{isSubmitting ? 'Resetting...' : 'Reset password'}
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
