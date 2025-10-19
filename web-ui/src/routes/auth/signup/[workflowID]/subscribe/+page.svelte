<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { signupStore } from '$lib/store/signupStore';
	import { Button, Input } from '@movsm/v1-consortium-web-pkg';
	import StripePayment from '../../StripePayment.svelte';

	// Get workflow ID from URL params
	let workflowId = $page.params.workflowID;
	
	// State management
	let currentStep = $state(1);
	let isSubmitting = $state(false);
	let errors = $state<Record<string, string>>({});

	// Step 1: Account Type Selection
	let accountType = $state<'company' | 'individual' | null>(null);

	// Step 2: Personal Information
	let firstName = $state('');
	let lastName = $state('');

	// Step 3: Company Information (for companies)
	let companyName = $state('');
	let dotNumber = $state('');
	let mcNumber = $state('');
	let phoneNumber = $state('');

	// Step 4: Address Information
	let addressLine1 = $state('');
	let addressLine2 = $state('');
	let city = $state('');
	let addressState = $state('');
	let zipCode = $state('');

	// Step 5: Plan Selection
	let selectedPlan = $state<'basic' | 'premium'>('basic');

	// Step 6: Payment (for companies with premium plans)
	let paymentMethodId = $state('');
	let customerId = $state('');

	onMount(async () => {
		if (!workflowId) {
			goto('/auth/signup');
			return;
		}

		try {
			await signupStore.loadWorkflow(workflowId);
			// Check if email is verified
			if (!$signupStore.isEmailVerified) {
				goto(`/auth/signup/${workflowId}/verify-email`);
				return;
			}
		} catch (err) {
			console.error('Failed to load workflow:', err);
			goto('/auth/signup');
		}
	});

	function nextStep() {
		if (validateCurrentStep()) {
			currentStep++;
		}
	}

	function previousStep() {
		currentStep--;
		signupStore.clearError();
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
				if (!firstName.trim()) {
					newErrors.firstName = 'First name is required';
				}
				if (!lastName.trim()) {
					newErrors.lastName = 'Last name is required';
				}
				break;

			case 3:
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

			case 4:
				if (!addressLine1.trim()) {
					newErrors.addressLine1 = 'Address is required';
				}
				if (!city.trim()) {
					newErrors.city = 'City is required';
				}
				if (!addressState.trim()) {
					newErrors.addressState = 'State is required';
				}
				if (!zipCode.trim()) {
					newErrors.zipCode = 'ZIP code is required';
				}
				break;
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleComplete() {
		if (!workflowId) {
			goto('/auth/signup');
			return;
		}

		if (accountType === 'company' && selectedPlan === 'premium' && currentStep === 5) {
			// For companies with premium plans, go to payment step
			if (!validateCurrentStep()) return;
			currentStep = 6;
			return;
		}

		// Complete the subscription
		if (!validateCurrentStep()) return;

		isSubmitting = true;
		signupStore.clearError();

		try {
			const subscriptionData = {
				accountType,
				personalInfo: {
					firstName,
					lastName
				},
				...(accountType === 'company' && {
					companyInfo: {
						name: companyName,
						dotNumber,
						mcNumber,
						phone: phoneNumber
					}
				}),
				addressInfo: {
					line1: addressLine1,
					line2: addressLine2,
					city,
					state: addressState,
					zipCode
				},
				planInfo: {
					selectedPlan
				},
				...(paymentMethodId && {
					paymentInfo: {
						paymentMethodId,
						customerId
					}
				})
			};

			//await signupStore.completeSubscription(workflowId, subscriptionData);
			
			// Redirect to success/dashboard
			goto('/dashboard');
		} catch (error) {
			console.error('Subscription failed:', error);
		} finally {
			isSubmitting = false;
		}
	}

	function handleFieldInput(field: string) {
		return (event: Event) => {
			const value = (event.target as HTMLInputElement).value;
			
			switch (field) {
				case 'firstName': firstName = value; break;
				case 'lastName': lastName = value; break;
				case 'companyName': companyName = value; break;
				case 'dotNumber': dotNumber = value; break;
				case 'mcNumber': mcNumber = value; break;
				case 'phoneNumber': phoneNumber = value; break;
				case 'addressLine1': addressLine1 = value; break;
				case 'addressLine2': addressLine2 = value; break;
				case 'city': city = value; break;
				case 'addressState': addressState = value; break;
				case 'zipCode': zipCode = value; break;
			}

			// Clear specific field error
			if (errors[field]) {
				errors = { ...errors, [field]: '' };
			}
		};
	}

	function selectAccountType(type: 'company' | 'individual') {
		accountType = type;
		if (errors.accountType) {
			errors = { ...errors, accountType: '' };
		}
		
		// Update store
		signupStore.updateWorkflowState({ accountType: type });
	}

	function selectPlan(plan: 'basic' | 'premium') {
		selectedPlan = plan;
	}

	// Payment handlers
	function handlePaymentSuccess(paymentMethodIdParam: string, customerIdParam: string) {
		paymentMethodId = paymentMethodIdParam;
		customerId = customerIdParam;
		// Auto-proceed to complete subscription
		handleComplete();
	}

	function handlePaymentError(message: string) {
		console.error('Payment failed:', message);
	}

	function handlePaymentCancel() {
		currentStep = 5; // Go back to plan selection
	}

	// Calculate total steps
	// $: totalSteps = accountType === 'company' && selectedPlan === 'premium' ? 6 : 5;
</script>

<svelte:head>
	<title>Complete Registration - Consortium</title>
</svelte:head>

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
					<h2 class="text-2xl font-semibold text-gray-900">Choose Account Type</h2>
					<p class="text-gray-600 mt-2">Select the type of account that best fits your needs</p>
				</div>

				<div class="space-y-4">
					<button
						onclick={() => selectAccountType('company')}
						class="w-full p-6 border-2 rounded-lg text-left transition-colors {accountType === 'company' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
					>
						<div class="flex items-center">
							<div class="flex-shrink-0">
								<svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"></path>
								</svg>
							</div>
							<div class="ml-4">
								<h3 class="text-lg font-medium text-gray-900">Company/Carrier</h3>
								<p class="text-gray-600">For trucking companies and carriers with DOT requirements</p>
							</div>
						</div>
					</button>

					<button
						onclick={() => selectAccountType('individual')}
						class="w-full p-6 border-2 rounded-lg text-left transition-colors {accountType === 'individual' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
					>
						<div class="flex items-center">
							<div class="flex-shrink-0">
								<svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
								</svg>
							</div>
							<div class="ml-4">
								<h3 class="text-lg font-medium text-gray-900">Individual Driver</h3>
								<p class="text-gray-600">For individual drivers and owner-operators</p>
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
				<!-- Step 2: Personal Information -->
				<div class="text-center mb-6">
					<h2 class="text-2xl font-semibold text-gray-900">Personal Information</h2>
					<p class="text-gray-600 mt-2">Tell us about yourself</p>
				</div>

				<div class="space-y-6">
					<div class="grid grid-cols-2 gap-4">
						<div>
							<label for="firstName" class="block text-sm font-medium text-gray-700">
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
							<label for="lastName" class="block text-sm font-medium text-gray-700">
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
				</div>

				<div class="mt-8">
					<Button 
						color="primary" 
						size="lg" 
						classes="w-full"
						onclick={nextStep}
					>
						Continue
					</Button>
				</div>

			{:else if currentStep === 3}
				<!-- Step 3: Company Information (if company) -->
				<div class="text-center mb-6">
					<h2 class="text-2xl font-semibold text-gray-900">
						{accountType === 'company' ? 'Company Information' : 'Account Setup'}
					</h2>
					<p class="text-gray-600 mt-2">
						{accountType === 'company' ? 'Tell us about your company' : 'Continue to address information'}
					</p>
				</div>

				{#if accountType === 'company'}
					<div class="space-y-6">
						<div>
							<label for="companyName" class="block text-sm font-medium text-gray-700">
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
								<label for="dotNumber" class="block text-sm font-medium text-gray-700">
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
							</div>

							<div>
								<label for="mcNumber" class="block text-sm font-medium text-gray-700">
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
							<label for="phoneNumber" class="block text-sm font-medium text-gray-700">
								Phone Number *
							</label>
							<Input
								id="phoneNumber"
								type="tel"
								value={phoneNumber}
								oninput={handleFieldInput('phoneNumber')}
								placeholder="(555) 123-4567"
								error={errors.phoneNumber}
								required
							/>
						</div>
					</div>
				{:else}
					<p class="text-center text-gray-600">Individual account setup is simplified. Click continue to proceed.</p>
				{/if}

				<div class="mt-8">
					<Button 
						color="primary" 
						size="lg" 
						classes="w-full"
						onclick={nextStep}
					>
						Continue
					</Button>
				</div>

			{:else if currentStep === 4}
				<!-- Step 4: Address Information -->
				<div class="text-center mb-6">
					<h2 class="text-2xl font-semibold text-gray-900">Address Information</h2>
					<p class="text-gray-600 mt-2">Where are you located?</p>
				</div>

				<div class="space-y-6">
					<div>
						<label for="addressLine1" class="block text-sm font-medium text-gray-700">
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
						<label for="addressLine2" class="block text-sm font-medium text-gray-700">
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

					<div class="grid grid-cols-3 gap-4">
						<div>
							<label for="city" class="block text-sm font-medium text-gray-700">
								City *
							</label>
							<Input
								id="city"
								type="text"
								value={city}
								oninput={handleFieldInput('city')}
								placeholder="Chicago"
								error={errors.city}
								required
							/>
						</div>

						<div>
							<label for="addressState" class="block text-sm font-medium text-gray-700">
								State *
							</label>
							<Input
								id="addressState"
								type="text"
								value={addressState}
								oninput={handleFieldInput('addressState')}
								placeholder="IL"
								error={errors.addressState}
								required
							/>
						</div>

						<div>
							<label for="zipCode" class="block text-sm font-medium text-gray-700">
								ZIP Code *
							</label>
							<Input
								id="zipCode"
								type="text"
								value={zipCode}
								oninput={handleFieldInput('zipCode')}
								placeholder="60601"
								error={errors.zipCode}
								required
							/>
						</div>
					</div>
				</div>

				<div class="mt-8">
					<Button 
						color="primary" 
						size="lg" 
						classes="w-full"
						onclick={nextStep}
					>
						Continue
					</Button>
				</div>

			{:else if currentStep === 5}
				<!-- Step 5: Plan Selection -->
				<div class="text-center mb-6">
					<h2 class="text-2xl font-semibold text-gray-900">Choose Your Plan</h2>
					<p class="text-gray-600 mt-2">Select the plan that works best for you</p>
				</div>

				<div class="space-y-4">
					<div class="grid grid-cols-1 gap-4 lg:grid-cols-2">
						<button
							onclick={() => selectPlan('basic')}
							class="p-6 border-2 rounded-lg text-left transition-colors {selectedPlan === 'basic' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
						>
							<h3 class="text-lg font-medium text-gray-900">Basic Plan</h3>
							<p class="text-gray-600 mt-1">Essential DOT compliance features</p>
							<p class="text-2xl font-bold text-gray-900 mt-2">Free</p>
						</button>

						<button
							onclick={() => selectPlan('premium')}
							class="p-6 border-2 rounded-lg text-left transition-colors {selectedPlan === 'premium' ? 'border-blue-600 bg-blue-50' : 'border-gray-200 hover:border-gray-300'}"
						>
							<h3 class="text-lg font-medium text-gray-900">Premium Plan</h3>
							<p class="text-gray-600 mt-1">Advanced features + priority support</p>
							<p class="text-2xl font-bold text-gray-900 mt-2">$99/month</p>
						</button>
					</div>
				</div>

				<div class="mt-8">
					<Button 
						color="primary" 
						size="lg" 
						classes="w-full"
						onclick={handleComplete}
						loading={isSubmitting}
					>
						{accountType === 'company' && selectedPlan === 'premium' ? 'Continue to Payment' : 'Complete Registration'}
					</Button>
				</div>

			{:else if currentStep === 6}
				<!-- Step 6: Payment (Company Premium only) -->
				<div class="text-center mb-6">
					<h2 class="text-2xl font-semibold text-gray-900">Payment Information</h2>
					<p class="text-gray-600 mt-2">Complete your premium subscription</p>
				</div>

			{/if}

			<!-- Error alert -->
			{#if $signupStore.error}
				<div class="mt-6 p-3 bg-red-50 border border-red-200 rounded-md">
					<p class="text-red-600 text-sm">{$signupStore.error}</p>
				</div>
			{/if}
		</div>
	</div>
</div>
