<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Button, Input, Alert } from '@movsm/v1-consortium-web-pkg';
	import { validatePassword, validateRequired } from '@movsm/v1-consortium-web-pkg';

	let token = $state('');
	let invitationDetails = $state<{ email?: string; organization?: string; role?: string } | null>(null);
	let isValidating = $state(true);
	let isInvalid = $state(false);
	
	// Form state
	let firstName = $state('');
	let lastName = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let errors = $state<Record<string, string>>({});
	let isSubmitting = $state(false);

	onMount(() => {
		// Get token from URL
		token = $page.url.searchParams.get('token') || '';
		
		if (!token) {
			isInvalid = true;
			isValidating = false;
			return;
		}

		validateInvitation();
	});

	async function validateInvitation() {
		try {
			// Simulate API call to validate invitation token
			await new Promise(resolve => setTimeout(resolve, 1500));
			
			// In a real implementation, this would call:
			// const invitation = await authStore.validateInvitation(token);
			
			// For now, simulate invitation details
			invitationDetails = {
				email: 'john.doe@abctrucking.com',
				organization: 'ABC Trucking Company',
				role: 'Driver'
			};
			
			isValidating = false;
		} catch (error) {
			isInvalid = true;
			isValidating = false;
		}
	}

	function validateForm(): boolean {
		const newErrors: Record<string, string> = {};

		const firstNameError = validateRequired(firstName, 'First name');
		if (firstNameError) newErrors.firstName = firstNameError;

		const lastNameError = validateRequired(lastName, 'Last name');
		if (lastNameError) newErrors.lastName = lastNameError;

		const passwordError = validatePassword(password);
		if (passwordError) newErrors.password = passwordError;

		if (password !== confirmPassword) {
			newErrors.confirmPassword = 'Passwords do not match';
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function handleSubmit() {
		if (!validateForm()) return;

		isSubmitting = true;
		authStore.clearError();

		try {
			// Simulate API call delay
			await new Promise(resolve => setTimeout(resolve, 2000));
			
			// In a real implementation, this would call:
			// await authStore.acceptInvitation({
			//   token,
			//   firstName,
			//   lastName,
			//   password
			// });
			
			// Simulate success and redirect to dashboard
			goto('/dashboard');
		} catch (error) {
			authStore.setError('Failed to accept invitation. Please try again.');
		} finally {
			isSubmitting = false;
		}
	}

	function createFieldHandler(field: string) {
		return (event: Event) => {
			const value = (event.target as HTMLInputElement).value;
			if (field === 'firstName') firstName = value;
			else if (field === 'lastName') lastName = value;
			else if (field === 'password') password = value;
			else if (field === 'confirmPassword') confirmPassword = value;
			
			// Clear field error
			if (errors[field]) {
				errors = { ...errors, [field]: '' };
			}
		};
	}
</script>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			
			{#if isValidating}
				<!-- Loading State -->
				<div class="text-center">
					<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto mb-4"></div>
					<h2 class="text-xl font-semibold text-gray-900 mb-2">Validating Invitation</h2>
					<p class="text-gray-600">Please wait while we verify your invitation...</p>
				</div>

			{:else if isInvalid}
				<!-- Invalid Token State -->
				<div class="text-center">
					<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-red-100 mb-4">
						<svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
						</svg>
					</div>

					<h2 class="text-xl font-semibold text-gray-900 mb-4">Invalid Invitation</h2>
					
					<p class="text-gray-600 mb-6">
						This invitation link is invalid or has expired. Please contact your organization 
						administrator for a new invitation.
					</p>

					<div class="bg-yellow-50 border border-yellow-200 rounded-md p-4 mb-6">
						<div class="flex">
							<div class="flex-shrink-0">
								<svg class="h-5 w-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
									<path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
								</svg>
							</div>
							<div class="ml-3">
								<h3 class="text-sm font-medium text-yellow-800">Common Issues</h3>
								<ul class="mt-2 text-sm text-yellow-700 list-disc list-inside">
									<li>The invitation link has expired</li>
									<li>The link has already been used</li>
									<li>The link was copied incorrectly</li>
								</ul>
							</div>
						</div>
					</div>

					<div class="space-y-3">
						<a
							href="/auth/signin"
							class="block w-full text-center bg-blue-600 border border-transparent rounded-md py-2 px-4 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
						>
							Go to Sign In
						</a>
					</div>
				</div>

			{:else if invitationDetails}
				<!-- Accept Invitation Form -->
				<div class="mb-6">
					<div class="text-center">
						<h2 class="text-2xl font-semibold text-gray-900">Accept Invitation</h2>
						<p class="text-gray-600 mt-2">Complete your account setup</p>
					</div>
				</div>

				<!-- Invitation Details -->
				<div class="bg-blue-50 border border-blue-200 rounded-md p-4 mb-6">
					<h3 class="text-sm font-medium text-blue-800 mb-2">Invitation Details</h3>
					<div class="text-sm text-blue-700 space-y-1">
						<p><span class="font-medium">Email:</span> {invitationDetails.email}</p>
						<p><span class="font-medium">Organization:</span> {invitationDetails.organization}</p>
						<p><span class="font-medium">Role:</span> {invitationDetails.role}</p>
					</div>
				</div>

				<form onsubmit={handleSubmit} class="space-y-6">
					<div class="grid grid-cols-1 gap-6 sm:grid-cols-2">
						<div>
							<label for="firstName" class="block text-sm font-medium text-gray-700 mb-1">
								First Name *
							</label>
							<Input
								id="firstName"
								type="text"
								value={firstName}
								oninput={createFieldHandler('firstName')}
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
								oninput={createFieldHandler('lastName')}
								placeholder="Doe"
								error={errors.lastName}
								required
							/>
						</div>
					</div>

					<div>
						<label for="password" class="block text-sm font-medium text-gray-700 mb-1">
							Password *
						</label>
						<Input
							id="password"
							type="password"
							value={password}
							oninput={createFieldHandler('password')}
							placeholder="Create a strong password"
							error={errors.password}
							required
						/>
						<p class="mt-1 text-xs text-gray-500">
							Must be at least 8 characters with uppercase, lowercase, number, and special character
						</p>
					</div>

					<div>
						<label for="confirmPassword" class="block text-sm font-medium text-gray-700 mb-1">
							Confirm Password *
						</label>
						<Input
							id="confirmPassword"
							type="password"
							value={confirmPassword}
							oninput={createFieldHandler('confirmPassword')}
							placeholder="Confirm your password"
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
							{isSubmitting ? 'Accepting Invitation...' : 'Accept Invitation & Join'}
						</Button>
					</div>

					<div class="text-center">
						<p class="text-xs text-gray-500">
							By accepting this invitation, you agree to the Consortium
							<a href="/terms" class="text-blue-600 hover:text-blue-500">Terms of Service</a>
							and <a href="/privacy" class="text-blue-600 hover:text-blue-500">Privacy Policy</a>.
						</p>
					</div>
				</form>
			{/if}
		</div>
	</div>

</div>	