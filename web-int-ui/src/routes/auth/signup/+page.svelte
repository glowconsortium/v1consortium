<script lang="ts">
	import { goto } from '$app/navigation';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Button, Input, Alert } from '@movsm/v1-consortium-web-pkg';
	import { validateEmail, validateRequired, validatePassword, validateConfirmPassword } from '@movsm/v1-consortium-web-pkg';
	import type { RegisterRequest } from '@movsm/v1-consortium-web-pkg';
	// import { consortium_api_auth_v1_SignupReq } from '@movsm/v1-consortium-web-pkg';
	import StripePayment from './StripePayment.svelte';

	// State management
	let currentStep = $state(1);
	let isSubmitting = $state(false);
	let errors = $state<Record<string, string>>({});
	let showVerifyEmail = $state(false);
	let paymentCompleted = $state(false);

	// Step 1: Account Type Selection
	let accountType = $state<'company' | 'individual' | null>(null);

	// Step 2: Company Information (for carriers)
	let companyName = $state('');
	let dotNumber = $state('');
	let mcNumber = $state('');
	let phoneNumber = $state('');

	// Step 3: Address Information
	let addressLine1 = $state('');
	let addressLine2 = $state('');
	let city = $state('');
	let addressstate = $state('');
	let zipCode = $state('');

	// Step 4: Account & Plan
	let firstName = $state('');
	let lastName = $state('');
	let email = $state('');
	let password = $state('');
	let selectedPlan = $state<'basic' | 'premium'>('basic');
	let agreeToTerms = $state(false);
	let showPassword = $state(false);

	// Step 5: Payment (for companies)
	let paymentMethodId = $state('');
	let customerId = $state('');

	// Redirect if already authenticated
	$effect(() => {
		if ($authStore.isAuthenticated) {
			goto('/dashboard');
		}
	});

	// Navigation functions
	function nextStep() {
		if (validateCurrentStep()) {
			currentStep++;
		}
	}

	function previousStep() {
		currentStep--;
		authStore.clearError();
	}

	function validateCurrentStep(): boolean {
		const newErrors: Record<string, string> = {};

		switch (currentStep) {
			case 1:
				if (!accountType) {
					newErrors.accountType = 'Please select an account type';
				}
				break;

			case 2:
				if (accountType === 'company') {
					if (!companyName.trim()) {
						newErrors.companyName = 'Company name is required';
					}
					if (!dotNumber.trim()) {
						newErrors.dotNumber = 'DOT number is required';
					}
					if (!phoneNumber.trim()) {
						newErrors.phoneNumber = 'Phone number is required';
					}
				}
				break;

			case 3:
				if (!addressLine1.trim()) {
					newErrors.addressLine1 = 'Address is required';
				}
				if (!city.trim()) {
					newErrors.city = 'City is required';
				}
				if (!addressstate.trim()) {
					newErrors.addressstate = 'State is required';
				}
				if (!zipCode.trim()) {
					newErrors.zipCode = 'ZIP code is required';
				}
				break;

			case 4:
				if (!firstName.trim()) {
					newErrors.firstName = 'First name is required';
				}
				if (!lastName.trim()) {
					newErrors.lastName = 'Last name is required';
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
				if (!agreeToTerms) {
					newErrors.agreeToTerms = 'You must agree to the terms and privacy policy';
				}
				break;
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (accountType === 'company' && currentStep === 4) {
			// For companies, go to payment step
			if (!validateCurrentStep()) return;
			currentStep = 5;
			return;
		}

		// For individuals or after payment completion, proceed with registration
		if (!validateCurrentStep()) return;
		if (!accountType) return; // Early return if accountType is null

		isSubmitting = true;
		authStore.clearError();


		try {
			const signupData = {
				email,
				password,
				firstName,
				lastName,
				// accountType: accountType === 'company'? consortium_api_auth_v1_SignupReq.accountType.COMPANY : consortium_api_auth_v1_SignupReq.accountType.INDIVIDUAL,
				addressLine1,
				city,
				state: addressstate,
				zipCode,
				...(accountType === 'company' && {
					carrierName: companyName,
					dotNumber: dotNumber,
					mcNumber: mcNumber,
					phone: phoneNumber,
					address: {
						line1: addressLine1,
						line2: addressLine2,
						city: city,
						state: addressstate,
						zipCode: zipCode
					},
					...(paymentMethodId && { paymentMethodId }),
					...(customerId && { customerId })
				}),
				stripePaymentMethodId: paymentMethodId,
				// subscriptionPlan:selectedPlan === 'basic' ?  consortium_api_auth_v1_SignupReq.subscriptionPlan.BASIC : consortium_api_auth_v1_SignupReq.subscriptionPlan.PREMIUM
			};

			const result = await authStore.register(email, password, `${firstName} ${lastName}`.trim(), signupData);
			showVerifyEmail = true;
		} catch (error) {
			console.error('Registration failed:', error);
		} finally {
			isSubmitting = false;
		}
	}

	function handleFieldInput(field: string) {
		return (event: Event) => {
			const value = (event.target as HTMLInputElement).value;
			
			switch (field) {
				case 'companyName': companyName = value; break;
				case 'dotNumber': dotNumber = value; break;
				case 'mcNumber': mcNumber = value; break;
				case 'phoneNumber': phoneNumber = value; break;
				case 'addressLine1': addressLine1 = value; break;
				case 'addressLine2': addressLine2 = value; break;
				case 'city': city = value; break;
				case 'addressstate': addressstate = value; break;
				case 'zipCode': zipCode = value; break;
				case 'firstName': firstName = value; break;
				case 'lastName': lastName = value; break;
				case 'email': email = value; break;
				case 'password': password = value; break;
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

	function selectAccountType(type: 'company' | 'individual') {
		accountType = type;
		if (errors.accountType) {
			errors = { ...errors, accountType: '' };
		}
	}

	function selectPlan(plan: 'basic' | 'premium') {
		selectedPlan = plan;
	}

	// Payment handlers
	function handlePaymentSuccess(paymentMethodId:string , customerId:string) {
		paymentMethodId = paymentMethodId;
		customerId = customerId;
		paymentCompleted = true;
		// Auto-proceed to complete registration
		handleSubmit();
	}

	function handlePaymentError(message: string) {
		console.error('Payment failed:', message);
		// Error is already displayed by StripePayment component
	}

	function handlePaymentCancel() {
		currentStep = 4; // Go back to account creation step
	}


</script>

<svelte:head>
	<title>Sign Up - Consortium</title>
	<meta name="description" content="Create your Consortium account for DOT compliance management." />
</svelte:head>

{#if showVerifyEmail}
	<!-- Email Verification Screen -->
	<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
		<div class="sm:mx-auto sm:w-full sm:max-w-md">
			<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10 text-center">
				<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-green-100 mb-4">
					<svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
					</svg>
				</div>
				<h2 class="text-2xl font-semibold text-gray-900 mb-2">Account Created!</h2>
				<h3 class="text-lg font-medium text-gray-700 mb-4">Welcome to Consortium{companyName ? `, ${companyName}` : ''}!</h3>
				
				<p class="text-gray-600 mb-4">
					We've sent a verification email to:<br>
					<span class="font-medium">{email}</span>
				</p>
				
				<p class="text-sm text-gray-500 mb-6">
					Please check your email and click the verification link to activate your account.
				</p>

				<div class="bg-gray-50 rounded-lg p-4 mb-6">
					<h4 class="font-medium text-gray-900 mb-2">Next Steps:</h4>
					<ol class="text-sm text-gray-600 space-y-1 text-left">
						<li class="flex items-center">
							<span class="text-blue-600 mr-2">1.</span>
							<span>📧 Verify your email address</span>
						</li>
						<li class="flex items-center">
							<span class="text-blue-600 mr-2">2.</span>
							<span>👥 Add your first drivers</span>
						</li>
						<li class="flex items-center">
							<span class="text-blue-600 mr-2">3.</span>
							<span>🧪 Set up random testing pools</span>
						</li>
						<li class="flex items-center">
							<span class="text-blue-600 mr-2">4.</span>
							<span>📋 Order compliance tests</span>
						</li>
					</ol>
				</div>

				<div class="space-y-3">
					<Button color="primary" size="lg" classes="w-full" onclick={() => goto('/dashboard')}>
						Go to Dashboard
					</Button>
					<button class="text-blue-600 hover:text-blue-500 text-sm font-medium">
						Resend Email
					</button>
				</div>

				<p class="text-xs text-gray-500 mt-6">
					Need help? <a href="/contact" class="text-blue-600 hover:text-blue-500">Contact Support</a> or 
					<a href="/onboarding" class="text-blue-600 hover:text-blue-500">Schedule Onboarding</a>
				</p>
			</div>
		</div>
	</div>
{:else}
	<!-- Multi-step Signup Form -->
	<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
		<div class="sm:mx-auto sm:w-full sm:max-w-md">
			<!-- Header with back button and progress -->
			<div class="flex items-center justify-between mb-6">
				{#if currentStep > 1}
					<button onclick={previousStep} class="flex items-center text-gray-600 hover:text-gray-900">
						<svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
						</svg>
						Back
					</button>
				{:else}
					<div></div>
				{/if}
				
				<div class="text-center">
					<span class="text-sm text-gray-500">{currentStep} of {accountType === 'company' ? 5 : 4} Steps</span>
					<div class="flex space-x-1 mt-1">
						{#each Array(accountType === 'company' ? 5 : 4) as _, i}
							<div class="h-2 w-8 rounded-full {i < currentStep ? 'bg-blue-600' : 'bg-gray-200'}"></div>
						{/each}
					</div>
				</div>
				
				<div></div>
			</div>

			<!-- Logo -->
			<div class="text-center mb-8">
				<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
			</div>

			<!-- Form Container -->
			<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
				
				{#if currentStep === 1}
					<!-- Step 1: Account Type Selection -->
					<div class="text-center mb-6">
						<h2 class="text-2xl font-semibold text-gray-900">Get Started with Consortium</h2>
						<p class="text-gray-600 mt-2">Choose your account type to begin</p>
					</div>

					<div class="space-y-4">
						<button
							onclick={() => selectAccountType('company')}
							class="w-full p-6 border-2 rounded-lg text-left transition-colors {accountType === 'company' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
						>
							<div class="flex items-start">
								<div class="text-2xl mr-4">🚛</div>
								<div>
									<h3 class="text-lg font-medium text-gray-900">Motor Carrier</h3>
									<p class="text-sm text-gray-600 mt-1">DOT-regulated trucking company</p>
								</div>
							</div>
						</button>

						<button
							onclick={() => selectAccountType('individual')}
							class="w-full p-6 border-2 rounded-lg text-left transition-colors {accountType === 'individual' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
						>
							<div class="flex items-start">
								<div class="text-2xl mr-4">👤</div>
								<div>
									<h3 class="text-lg font-medium text-gray-900">Individual</h3>
									<p class="text-sm text-gray-600 mt-1">Single driver or owner-operator</p>
								</div>
							</div>
						</button>

						{#if errors.accountType}
							<p class="text-red-600 text-sm">{errors.accountType}</p>
						{/if}
					</div>

					<div class="mt-8">
						<Button 
							color="primary" 
							size="lg" 
							classes="w-full" 
							onclick={nextStep}
							disabled={!accountType}
						>
							Continue
						</Button>
					</div>

				{:else if currentStep === 2}
					<!-- Step 2: Company Information -->
					<div class="text-center mb-6">
						<h2 class="text-2xl font-semibold text-gray-900">
							{accountType === 'company' ? 'Company Information' : 'Personal Information'}
						</h2>
						<p class="text-gray-600 mt-2">
							{accountType === 'company' ? 'Tell us about your company' : 'Tell us about yourself'}
						</p>
					</div>

					{#if accountType === 'company'}
						<div class="space-y-6">
							<div>
								<label for="companyName" class="block text-sm font-medium text-gray-700 mb-1">
									Company Name *
								</label>
								<Input
									id="companyName"
									type="text"
									value={companyName}
									oninput={handleFieldInput('companyName')}
									placeholder="ABC Trucking LLC"
									error={errors.companyName}
									required
								/>
							</div>

							<div class="grid grid-cols-2 gap-4">
								<div>
									<label for="dotNumber" class="block text-sm font-medium text-gray-700 mb-1">
										DOT Number *
									</label>
									<Input
										id="dotNumber"
										type="text"
										value={dotNumber}
										oninput={handleFieldInput('dotNumber')}
										placeholder="123456"
										error={errors.dotNumber}
										required
									/>
									<p class="text-xs text-gray-500 mt-1">❓ Your DOT number is required for FMCSA compliance</p>
								</div>

								<div>
									<label for="mcNumber" class="block text-sm font-medium text-gray-700 mb-1">
										MC Number
									</label>
									<Input
										id="mcNumber"
										type="text"
										value={mcNumber}
										oninput={handleFieldInput('mcNumber')}
										placeholder="654321"
									/>
								</div>
							</div>

							<div>
								<label for="phoneNumber" class="block text-sm font-medium text-gray-700 mb-1">
									Phone Number *
								</label>
								<Input
									id="phoneNumber"
									type="tel"
									value={phoneNumber}
									oninput={handleFieldInput('phoneNumber')}
									placeholder="+1 (555) 123-4567"
									error={errors.phoneNumber}
									required
								/>
							</div>
						</div>
					{:else}
						<p class="text-center text-gray-600">Individual account setup is simplified. Click continue to proceed.</p>
					{/if}

					<div class="mt-8">
						<Button color="primary" size="lg" classes="w-full" onclick={nextStep}>
							Continue
						</Button>
					</div>

				{:else if currentStep === 3}
					<!-- Step 3: Address Information -->
					<div class="text-center mb-6">
						<h2 class="text-2xl font-semibold text-gray-900">
							{accountType === 'company' ? 'Business Address' : 'Address'}
						</h2>
						<p class="text-gray-600 mt-2">
							{accountType === 'company' ? 'Required for DOT registration' : 'Your contact address'}
						</p>
					</div>

					<div class="space-y-6">
						<div>
							<label for="addressLine1" class="block text-sm font-medium text-gray-700 mb-1">
								Address Line 1 *
							</label>
							<Input
								id="addressLine1"
								type="text"
								value={addressLine1}
								oninput={handleFieldInput('addressLine1')}
								placeholder="123 Main Street"
								error={errors.addressLine1}
								required
							/>
						</div>

						<div>
							<label for="addressLine2" class="block text-sm font-medium text-gray-700 mb-1">
								Address Line 2
							</label>
							<Input
								id="addressLine2"
								type="text"
								value={addressLine2}
								oninput={handleFieldInput('addressLine2')}
								placeholder="Suite 100"
							/>
						</div>

						<div class="grid grid-cols-6 gap-4">
							<div class="col-span-3">
								<label for="city" class="block text-sm font-medium text-gray-700 mb-1">
									City *
								</label>
								<Input
									id="city"
									type="text"
									value={city}
									oninput={handleFieldInput('city')}
									placeholder="Dallas"
									error={errors.city}
									required
								/>
							</div>

							<div class="col-span-2">
								<label for="state" class="block text-sm font-medium text-gray-700 mb-1">
									State *
								</label>
								<select
									id="state"
									bind:value={addressstate}
									class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
									required
								>
									<option value="">Select</option>
									<option value="TX">TX</option>
									<option value="CA">CA</option>
									<option value="FL">FL</option>
									<!-- Add more states as needed -->
								</select>
								{#if errors.state}
									<p class="text-red-600 text-sm mt-1">{errors.state}</p>
								{/if}
							</div>

							<div class="col-span-1">
								<label for="zipCode" class="block text-sm font-medium text-gray-700 mb-1">
									ZIP *
								</label>
								<Input
									id="zipCode"
									type="text"
									value={zipCode}
									oninput={handleFieldInput('zipCode')}
									placeholder="75201"
									error={errors.zipCode}
									required
								/>
							</div>
						</div>
					</div>

					<div class="mt-8">
						<Button color="primary" size="lg" classes="w-full" onclick={nextStep}>
							Continue
						</Button>
					</div>

				{:else if currentStep === 4}
					<!-- Step 4: Account & Plan Selection -->
					<div class="text-center mb-6">
						<h2 class="text-2xl font-semibold text-gray-900">Create Your Admin Account</h2>
					</div>

					<div class="space-y-6">
						<div class="grid grid-cols-2 gap-4">
							<div>
								<label for="firstName" class="block text-sm font-medium text-gray-700 mb-1">
									First Name *
								</label>
								<Input
									id="firstName"
									type="text"
									value={firstName}
									oninput={handleFieldInput('firstName')}
									placeholder="John"
									error={errors.firstName}
									required
								/>
							</div>

							<div>
								<label for="lastName" class="block text-sm font-medium text-gray-700 mb-1">
									Last Name *
								</label>
								<Input
									id="lastName"
									type="text"
									value={lastName}
									oninput={handleFieldInput('lastName')}
									placeholder="Doe"
									error={errors.lastName}
									required
								/>
							</div>
						</div>

						<div>
							<label for="email" class="block text-sm font-medium text-gray-700 mb-1">
								Email Address *
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

						<div>
							<label for="password" class="block text-sm font-medium text-gray-700 mb-1">
								Password *
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

						{#if accountType === 'company'}
							<!-- Plan Selection -->
							<div>
								<label for="plan" class="block text-sm font-medium text-gray-700 mb-3">Choose Your Plan</label>
								<div class="grid grid-cols-1 gap-4">
									<button
										type="button"
										onclick={() => selectPlan('basic')}
										class="p-4 border-2 rounded-lg text-left transition-colors {selectedPlan === 'basic' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
									>
										<div class="flex items-center justify-between mb-2">
											<div class="flex items-center">
												<div class="w-4 h-4 rounded-full border-2 mr-3 {selectedPlan === 'basic' ? 'border-blue-600 bg-blue-600' : 'border-gray-300'}">
													{#if selectedPlan === 'basic'}
														<div class="w-2 h-2 bg-white rounded-full mx-auto mt-0.5"></div>
													{/if}
												</div>
												<h4 class="font-medium">Basic Plan</h4>
											</div>
											<div class="text-right">
												<p class="text-lg font-semibold text-gray-900">$1,188/year</p>
												<p class="text-sm text-gray-500">$99/month</p>
											</div>
										</div>
										<div class="ml-7">
											<p class="text-sm text-gray-600 mb-2">Perfect for small carriers</p>
											<ul class="text-xs text-gray-500 space-y-1">
												<li>• Up to 50 drivers</li>
												<li>• Basic reporting</li>
												<li>• Email support</li>
											</ul>
										</div>
									</button>

									<button
										type="button"
										onclick={() => selectPlan('premium')}
										class="p-4 border-2 rounded-lg text-left transition-colors {selectedPlan === 'premium' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
									>
										<div class="flex items-center justify-between mb-2">
											<div class="flex items-center">
												<div class="w-4 h-4 rounded-full border-2 mr-3 {selectedPlan === 'premium' ? 'border-blue-600 bg-blue-600' : 'border-gray-300'}">
													{#if selectedPlan === 'premium'}
														<div class="w-2 h-2 bg-white rounded-full mx-auto mt-0.5"></div>
													{/if}
												</div>
												<h4 class="font-medium">Premium Plan</h4>
											</div>
											<div class="text-right">
												<p class="text-lg font-semibold text-gray-900">$2,388/year</p>
												<p class="text-sm text-gray-500">$199/month</p>
											</div>
										</div>
										<div class="ml-7">
											<p class="text-sm text-gray-600 mb-2">For growing operations</p>
											<ul class="text-xs text-gray-500 space-y-1">
												<li>• Up to 500 drivers</li>
												<li>• Advanced analytics</li>
												<li>• Priority support</li>
												<li>• API integrations</li>
											</ul>
										</div>
									</button>
								</div>
							</div>
						{/if}

						<!-- Terms checkbox -->
						<div class="flex items-start">
							<input
								id="terms"
								name="terms"
								type="checkbox"
								bind:checked={agreeToTerms}
								class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded mt-1"
								required
							/>
							<label for="terms" class="ml-3 block text-sm text-gray-900">
								I agree to the 
								<a href="/terms" class="font-medium text-blue-600 hover:text-blue-500">Terms of Service</a>
								 and 
								<a href="/privacy" class="font-medium text-blue-600 hover:text-blue-500">Privacy Policy</a>
							</label>
						</div>
						{#if errors.agreeToTerms}
							<p class="text-red-600 text-sm">{errors.agreeToTerms}</p>
						{/if}
					</div>

					<!-- Error alert -->
					{#if $authStore.error}
						<div class="mt-6">
							<Alert type="error" message={$authStore.error} />
						</div>
					{/if}

					<div class="mt-8">
						<Button
							color="primary"
							size="lg"
							classes="w-full"
							onclick={handleSubmit}
							disabled={isSubmitting}
						>
							{#if accountType === 'company'}
								{isSubmitting ? 'Processing...' : 'Continue to Payment'}
							{:else}
								{isSubmitting ? 'Creating Account...' : 'Create Account'}
							{/if}
						</Button>
					</div>

				{:else if currentStep === 5}
					<!-- Step 5: Payment (Company accounts only) -->
					<div class="text-center mb-6">
						<h2 class="text-2xl font-semibold text-gray-900">Complete Your Subscription</h2>
						<p class="text-gray-600 mt-2">Secure payment to activate your account</p>
					</div>

					<StripePayment
						{selectedPlan}
						{companyName}
						{email}
						isProcessing={isSubmitting}
						success={handlePaymentSuccess}
						error={handlePaymentError}
						cancel={handlePaymentCancel}
					/>
				{/if}

				<!-- Already have account link -->
				<div class="mt-6 text-center">
					<p class="text-sm text-gray-600">
						Already have an account?
						<a href="/auth/signin" class="font-medium text-blue-600 hover:text-blue-500">
							Sign In
						</a>
					</p>
				</div>
			</div>
		</div>
	</div>
{/if}
			

