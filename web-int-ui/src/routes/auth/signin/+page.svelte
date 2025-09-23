<script lang="ts">

	import { goto } from '$app/navigation';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Button, Input, Alert } from '@movsm/v1-consortium-web-pkg';
	import { validateEmail, validateRequired } from '@movsm/v1-consortium-web-pkg';

	let email = $state('');
	let password = $state('');
	let showPassword = $state(false);
	let rememberMe = $state(false);
	let errors = $state<Record<string, string>>({});
	let isSubmitting = $state(false);

	function validateForm(): boolean {
		const newErrors: Record<string, string> = {};

		const emailError = validateRequired(email, 'Email');
		if (emailError) {
			newErrors.email = emailError;
		} else if (!validateEmail(email)) {
			newErrors.email = 'Please enter a valid email address';
		}

		const passwordError = validateRequired(password, 'Password');
		if (passwordError) newErrors.password = passwordError;

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (!validateForm()) return;

		isSubmitting = true;
		authStore.clearError();

		try {
			await authStore.login(email, password);
			goto('/dashboard');
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

	function handlePasswordInput(event: Event) {
		password = (event.target as HTMLInputElement).value;
		if (errors.password) {
			errors = { ...errors, password: '' };
		}
	}

	function togglePasswordVisibility() {
		showPassword = !showPassword;
	}

	function handleFieldInput(field: string) {
		return (event: Event) => {
			const value = (event.target as HTMLInputElement).value;

			if (field === 'email') {
				email = value;
			} else if (field === 'password') {
				password = value;
			}

			// Clear specific field error
			if (errors[field]) {
				errors = { ...errors, [field]: '' };
			}
		};
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

			<form onsubmit={handleSubmit} class="space-y-6">
				<!-- Email field -->
				<div>
					<label for="email" class="mb-1 block text-sm font-medium text-gray-700">
						Email Address
					</label>
					<Input
						id="email"
						type="email"
						value={email}
						oninput={handleFieldInput('email')}
						placeholder="john@abctrucking.com"
						error={errors.email}
						required
					/>
				</div>

				<!-- Password field -->
				<div>
					<label for="password" class="mb-1 block text-sm font-medium text-gray-700">
						Password
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
							class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-400 hover:text-gray-600"
							onclick={togglePasswordVisibility}
							aria-label={showPassword ? 'Hide password' : 'Show password'}
						>
							{#if showPassword}
								<!-- Eye slash icon -->
								<svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"
									></path>
								</svg>
							{:else}
								<!-- Eye icon -->
								<svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
									></path>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
									></path>
								</svg>
							{/if}
						</button>
					</div>
				</div>

				<!-- Remember me and forgot password -->
				<div class="flex items-center justify-between">
					<div class="flex items-center">
						<input
							id="remember-me"
							name="remember-me"
							type="checkbox"
							bind:checked={rememberMe}
							class="h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
						/>
						<label for="remember-me" class="ml-2 block text-sm text-gray-900"> Remember me </label>
					</div>

					<div class="text-sm">
						<a href="/auth/forgot-password" class="font-medium text-blue-600 hover:text-blue-500">
							Forgot password?
						</a>
					</div>
				</div>

				<!-- Error alert -->
				{#if $authStore.error}
					<div class="mt-4">
						<Alert type="error" message={$authStore.error} />
					</div>
				{/if}

				<!-- Submit button -->
				<div>
					<Button type="submit" color="primary" size="lg" classes="w-full" disabled={isSubmitting}>
						{isSubmitting ? 'Signing In...' : 'Sign In'}
					</Button>
				</div>

				<!-- Divider -->
				<div class="mt-6">
					<div class="relative">
						<div class="absolute inset-0 flex items-center">
							<div class="w-full border-t border-gray-300" />
						</div>
						<div class="relative flex justify-center text-sm">
							<span class="bg-white px-2 text-gray-500">or</span>
						</div>
					</div>
				</div>

				<!-- Social login buttons -->
				<div class="mt-6 grid grid-cols-2 gap-3">
					<button
						type="button"
						class="inline-flex w-full justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 shadow-sm hover:bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:outline-none"
					>
						<svg class="h-5 w-5" viewBox="0 0 24 24" fill="currentColor">
							<path
								d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
								fill="#4285F4"
							/>
							<path
								d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
								fill="#34A853"
							/>
							<path
								d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
								fill="#FBBC05"
							/>
							<path
								d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
								fill="#EA4335"
							/>
						</svg>
						<span class="ml-2">Google</span>
					</button>

					<button
						type="button"
						class="inline-flex w-full justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 shadow-sm hover:bg-gray-50 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:outline-none"
					>
						<svg class="h-5 w-5" viewBox="0 0 24 24" fill="currentColor">
							<path
								d="M23.5 12.1c0-1.2-.1-2.4-.3-3.6h-10v6.8h5.7c-.2 1.2-1 2.3-2.1 3v2.5h3.4c2-1.8 3.3-4.4 3.3-7.7z"
								fill="#4285f4"
							/>
							<path
								d="M12.1 23c2.8 0 5.1-.9 6.8-2.5l-3.4-2.5c-.9.6-2.1 1-3.4 1-2.6 0-4.8-1.7-5.6-4.1H2.8v2.6C4.5 20.4 8.1 23 12.1 23z"
								fill="#34a853"
							/>
							<path
								d="M6.5 14.9c-.2-.6-.3-1.2-.3-1.9s.1-1.3.3-1.9V8.5H2.8C2.1 9.8 1.7 11.3 1.7 12.9s.4 3.1 1.1 4.4l3.7-2.4z"
								fill="#fbbc05"
							/>
							<path
								d="M12.1 4.8c1.5 0 2.8.5 3.8 1.5l2.9-2.9C17.2 1.8 14.9.9 12.1.9 8.1.9 4.5 3.5 2.8 7.4l3.7 2.6c.8-2.4 3-4.2 5.6-4.2z"
								fill="#ea4335"
							/>
						</svg>
						<span class="ml-2">Microsoft</span>
					</button>
				</div>
			</form>

			<!-- Sign up link -->
			<div class="mt-6 text-center">
				<p class="text-sm text-gray-600">
					Don't have an account?
					<a href="/auth/signup" class="font-medium text-blue-600 hover:text-blue-500">
						Sign up for free
					</a>
				</p>
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


