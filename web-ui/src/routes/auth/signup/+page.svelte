<script lang="ts">
	import { goto } from '$app/navigation';
	import { authStore, type SignupCredentials } from '@movsm/v1-consortium-web-pkg';
	import { Button, Input, Alert } from '@movsm/v1-consortium-web-pkg';
	import { validateEmail, validatePassword } from '@movsm/v1-consortium-web-pkg';
	import {signupStore} from '$lib/store/signupStore';

	// Form state
	let email = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let fullName = $state('');	
	let companyName = $state('');
	let isDotCompany = $state(false);
	let DOTNumber = $state('');
	let agreeToTerms = $state(false);
	let showPassword = $state(false);
	let showConfirmPassword = $state(false);
	let errors = $state<Record<string, string>>({});
	let isLoading = $state(false);

	// Redirect if already authenticated
	$effect(() => {
		if ($authStore.isAuthenticated) {
			goto('/dashboard');
		}
	});

	function validateForm(): boolean {
		const newErrors: Record<string, string> = {};

		if (!fullName.trim()) {
			newErrors.fullName = 'Full name is required';
		}

		if (!companyName.trim()) {
			newErrors.companyName = 'Company name is required';
		}

		if (isDotCompany && !DOTNumber.trim()) {
			newErrors.DOTNumber = 'DOT number is required for DOT regulated companies';
		}

		if (!email.trim()) {
			newErrors.email = 'Email is required';
		} else if (!validateEmail(email)) {
			newErrors.email = 'Please enter a valid email address';
		}

		if (!password.trim()) {
			newErrors.password = 'Password is required';
		} else {
			const passwordValidation = validatePassword(password);
			if (!passwordValidation.isValid) {
				newErrors.password = passwordValidation.errors[0];
			}
		}

		if (!confirmPassword.trim()) {
			newErrors.confirmPassword = 'Please confirm your password';
		} else if (password !== confirmPassword) {
			newErrors.confirmPassword = 'Passwords do not match';
		}

		if (!agreeToTerms) {
			newErrors.agreeToTerms = 'You must agree to the terms and privacy policy';
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (!validateForm()) return;

		isLoading = true;
		// Clear previous errors
		errors = { ...errors, general: '' };

		try {
			// For now, redirect to onboarding flow with email and password as query params
			// This allows the user to continue with the full multi-step signup process
			// const params = new URLSearchParams({
			// 	email: email.trim(),
			// 	password: password
			// });
			
			// Redirect to onboarding flow
			//goto(`/onboarding?${params.toString()}`);
			let signupCredentials:SignupCredentials = { 
				email: email.trim(), 
				password: password, 
				firstName: fullName.trim().split(' ')[0], 
				lastName: fullName.trim().split(' ')[1],
				companyName: companyName.trim(),
				isDotCompany: isDotCompany,
				dotNumber: isDotCompany ? DOTNumber.trim() : undefined
			};
		 const signupresp = await signupStore.signup(signupCredentials);
		 goto(`/auth/signup/${signupresp.workflowId}`);
			
		} catch (error: any) {
			console.error('Signup failed:', error);
			// Show the error to the user
			errors = { ...errors, general: error.message || 'Signup failed. Please try again.' };
		} finally {
			isLoading = false;
		}
	}

	function handleFieldInput(field: string) {
		return (event: Event) => {
			const value = (event.target as HTMLInputElement).value;
			
			switch (field) {
				case 'fullName': fullName = value; break;
				case 'companyName': companyName = value; break;
				case 'DOTNumber': DOTNumber = value; break;
				case 'email': email = value; break;
				case 'password': password = value; break;
				case 'confirmPassword': confirmPassword = value; break;
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

	function toggleConfirmPasswordVisibility() {
		showConfirmPassword = !showConfirmPassword;
	}

</script>

<svelte:head>
	<title>Sign Up - Consortium</title>
	<meta name="description" content="Create your Consortium account for DOT compliance management." />
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<!-- Form Container -->
		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			<div class="text-center mb-6">
				<h2 class="text-2xl font-semibold text-gray-900">Create Your Account</h2>
				<p class="text-gray-600 mt-2">Get started with Consortium</p>
			</div>

			<form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="space-y-6">
				<!-- Full Name -->
				<div>
					<label for="fullName" class="block text-sm font-medium text-gray-700">Full Name</label>
					<div class="mt-1">
						<Input
							id="fullName"
							type="text"
							value={fullName}
							oninput={handleFieldInput('fullName')}
							required
							placeholder="Enter your full name"
							error={errors.fullName}
						/>
					</div>
				</div>
				
				<!-- Company Name -->
				<div>
					<label for="companyName" class="block text-sm font-medium text-gray-700">Company Name</label>
					<div class="mt-1">
						<Input
							id="companyName"
							type="text"
							value={companyName}
							oninput={handleFieldInput('companyName')}
							placeholder="Enter your company name"
							error={errors.companyName}
						/>
					</div>
				</div>

				<!-- DOT Company Checkbox -->
				<div class="flex items-start">
					<div class="flex items-center h-5">
						<input
							id="isDotCompany"
							type="checkbox"
							bind:checked={isDotCompany}
							class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
						/>
					</div>
					<div class="ml-3 text-sm">
						<label for="isDotCompany" class="text-gray-900">
							This is a DOT regulated company
						</label>
					</div>
				</div>

				<!-- DOT Number (conditional) -->
				{#if isDotCompany}
					<div>
						<label for="DOTNumber" class="block text-sm font-medium text-gray-700">DOT Number</label>
						<div class="mt-1">
							<Input
								id="DOTNumber"
								type="text"
								value={DOTNumber}
								oninput={handleFieldInput('DOTNumber')}
								placeholder="Enter your DOT number"
								error={errors.DOTNumber}
							/>
						</div>
					</div>
				{/if}

				<!-- Email -->
				<div>
					<label for="email" class="block text-sm font-medium text-gray-700">Email address</label>
					<div class="mt-1">
						<Input
							id="email"
							type="email"
							value={email}
							oninput={handleFieldInput('email')}
							required
							placeholder="Enter your email"
							error={errors.email}
						/>
					</div>
				</div>

				<!-- Password -->
				<div>
					<label for="password" class="block text-sm font-medium text-gray-700">Password</label>
					<div class="mt-1 relative">
						<Input
							id="password"
							type={showPassword ? 'text' : 'password'}
							value={password}
							oninput={handleFieldInput('password')}
							required
							placeholder="Enter your password"
							error={errors.password}
						/>
						<button
							type="button"
							class="absolute inset-y-0 right-0 pr-3 flex items-center"
							onclick={togglePasswordVisibility}
						>
							<svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								{#if showPassword}
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L8.464 8.464M14.12 14.12l1.414 1.414M9.878 9.878l-3.29-3.29M14.12 14.12l3.29 3.29"></path>
								{:else}
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
								{/if}
							</svg>
						</button>
					</div>
				</div>

				<!-- Confirm Password -->
				<div>
					<label for="confirmPassword" class="block text-sm font-medium text-gray-700">Confirm Password</label>
					<div class="mt-1 relative">
						<Input
							id="confirmPassword"
							type={showConfirmPassword ? 'text' : 'password'}
							value={confirmPassword}
							oninput={handleFieldInput('confirmPassword')}
							required
							placeholder="Confirm your password"
							error={errors.confirmPassword}
						/>
						<button
							type="button"
							class="absolute inset-y-0 right-0 pr-3 flex items-center"
							onclick={toggleConfirmPasswordVisibility}
						>
							<svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								{#if showConfirmPassword}
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L8.464 8.464M14.12 14.12l1.414 1.414M9.878 9.878l-3.29-3.29M14.12 14.12l3.29 3.29"></path>
								{:else}
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
								{/if}
							</svg>
						</button>
					</div>
				</div>

				<!-- Terms and Conditions -->
				<div class="flex items-start">
					<div class="flex items-center h-5">
						<input
							id="agreeToTerms"
							type="checkbox"
							bind:checked={agreeToTerms}
							class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
						/>
					</div>
					<div class="ml-3 text-sm">
						<label for="agreeToTerms" class="text-gray-900">
							I agree to the 
							<a href="/terms" class="text-blue-600 hover:text-blue-500">Terms of Service</a> and 
							<a href="/privacy" class="text-blue-600 hover:text-blue-500">Privacy Policy</a>
						</label>
						{#if errors.agreeToTerms}
							<p class="text-red-600 text-sm mt-1">{errors.agreeToTerms}</p>
						{/if}
					</div>
				</div>

				<!-- Error alert -->
				{#if errors.general}
					<Alert type="error" message={errors.general} ondismiss={() => errors = { ...errors, general: '' }} />
				{/if}

				<!-- Submit Button -->
				<div>
					<Button
						type="submit"
						color="primary"
						size="lg"
						classes="w-full"
						loading={isLoading}
						disabled={!email || !password || !confirmPassword || !fullName || !companyName || !agreeToTerms || isLoading}
					>
						Create Account
					</Button>
				</div>
			</form>

			<!-- Social signup options -->
			<div class="mt-6">
				<div class="relative">
					<div class="absolute inset-0 flex items-center">
						<div class="w-full border-t border-gray-300"></div>
					</div>
					<div class="relative flex justify-center text-sm">
						<span class="px-2 bg-white text-gray-500">Or continue with</span>
					</div>
				</div>

				<div class="mt-6 grid grid-cols-2 gap-3">
					<a
						href="/auth/oauth/google"
						class="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
					>
						<svg class="w-5 h-5" viewBox="0 0 24 24">
							<path fill="currentColor" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
							<path fill="currentColor" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
							<path fill="currentColor" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
							<path fill="currentColor" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
						</svg>
						<span class="ml-2">Google</span>
					</a>

					<a
						href="/auth/oauth/microsoft"
						class="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
					>
						<svg class="w-5 h-5" viewBox="0 0 24 24">
							<path fill="currentColor" d="M11.4 24H0V12.6h11.4V24zM24 24H12.6V12.6H24V24zM11.4 11.4H0V0h11.4v11.4zM24 11.4H12.6V0H24v11.4z"/>
						</svg>
						<span class="ml-2">Microsoft</span>
					</a>
				</div>
			</div>

			<!-- Already have account link -->
			<div class="mt-6 text-center">
				<p class="text-sm text-gray-600">
					Already have an account? 
					<a href="/auth/signin" class="text-blue-600 hover:text-blue-500 font-medium">
						Sign in here
					</a>
				</p>
			</div>
		</div>
	</div>
</div>