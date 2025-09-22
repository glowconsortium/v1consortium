<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { authStore } from '$lib/stores/auth.js';
	import { Button, Input, Alert } from '$lib/components/ui/index.js';
	import { validatePassword, validateRequired, validateConfirmPassword } from '$lib/utils/validation.js';

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
</script>

<svelte:head>
	<title>Reset Password - Consortium</title>
	<meta name="description" content="Create a new password for your Consortium account." />
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			
			{#if isSuccess}
				<!-- Success Screen -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-green-100 mb-4">
						<svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
						</svg>
					</div>

					<h2 class="text-2xl font-semibold text-gray-900 mb-4">Password Reset Successful</h2>
					
					<p class="text-gray-600 mb-6">
						Your password has been successfully updated. You can now sign in with your new password.
					</p>

					<div>
						<a
							href="/auth/signin"
							class="block w-full text-center bg-blue-600 border border-transparent rounded-md py-2 px-4 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
						>
							Sign In
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
						<h2 class="text-2xl font-semibold text-gray-900">Create New Password</h2>
						<p class="text-gray-600 mt-2">Enter your new password below</p>
					</div>
				</div>

				<form on:submit|preventDefault={handleSubmit} class="space-y-6">
					<div>
						<label for="password" class="block text-sm font-medium text-gray-700 mb-1">
							New Password *
						</label>
						<div class="relative">
							<Input
								id="password"
								type={showPassword ? 'text' : 'password'}
								value={password}
								oninput={handleFieldInput('password')}
								placeholder="••••••••••••••••"
								error={errors.password}
								required
							/>
							<button
								type="button"
								class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-gray-600"
								onclick={togglePasswordVisibility}
								aria-label={showPassword ? 'Hide password' : 'Show password'}
							>
								{#if showPassword}
									<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
									</svg>
								{:else}
									<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
									</svg>
								{/if}
							</button>
						</div>
						
						<!-- Password strength indicator -->
						{#if password}
							<div class="mt-2 text-xs space-y-1">
								<div class="flex items-center space-x-2">
									<span class="{password.length >= 8 ? 'text-green-600' : 'text-gray-400'}">✓</span>
									<span class="{password.length >= 8 ? 'text-green-600' : 'text-gray-400'}">8+ characters</span>
								</div>
								<div class="flex items-center space-x-2">
									<span class="{/[A-Z]/.test(password) ? 'text-green-600' : 'text-gray-400'}">✓</span>
									<span class="{/[A-Z]/.test(password) ? 'text-green-600' : 'text-gray-400'}">Uppercase</span>
								</div>
								<div class="flex items-center space-x-2">
									<span class="{/[0-9]/.test(password) ? 'text-green-600' : 'text-gray-400'}">✓</span>
									<span class="{/[0-9]/.test(password) ? 'text-green-600' : 'text-gray-400'}">Number</span>
								</div>
								<div class="flex items-center space-x-2">
									<span class="{/[^A-Za-z0-9]/.test(password) ? 'text-green-600' : 'text-gray-400'}">✓</span>
									<span class="{/[^A-Za-z0-9]/.test(password) ? 'text-green-600' : 'text-gray-400'}">Special char</span>
								</div>
							</div>
						{/if}
					</div>

					<div>
						<label for="confirmPassword" class="block text-sm font-medium text-gray-700 mb-1">
							Confirm New Password *
						</label>
						<Input
							id="confirmPassword"
							type="password"
							value={confirmPassword}
							oninput={handleFieldInput('confirmPassword')}
							placeholder="••••••••••••••••"
							error={errors.confirmPassword}
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
							{isSubmitting ? 'Updating Password...' : 'Update Password'}
						</Button>
					</div>

					<div class="text-center">
						<p class="text-sm text-gray-600">
							Your new password must be different from your previous password and meet our security requirements.
						</p>
					</div>
				</form>
			{/if}
		</div>
	</div>
</div>
			newErrors.confirmPassword = confirmError;
		} else {
			const confirmValidationError = validateConfirmPassword(password, confirmPassword);
			if (confirmValidationError) {
				newErrors.confirmPassword = confirmValidationError;
			}
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (!validateForm()) return;

		isSubmitting = true;
		authStore.clearError();

		try {
			await authStore.resetPassword('', token, password);
			isSuccess = true;
		} catch (error) {
			// Error is handled by the store
		} finally {
			isSubmitting = false;
		}
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
