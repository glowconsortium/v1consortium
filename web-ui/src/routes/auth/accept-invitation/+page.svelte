<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { goto } from '$app/navigation';
    import { invitationService } from '$lib/api/invitation';
    import { Alert, Button, Card, Input } from '$lib/components/ui';
    import { validateEmail, validateRequired, validatePassword, validateConfirmPassword } from '$lib/utils/validation.js';
    import { authStore } from '$lib/stores/auth.js';

    let loading = $state(false);
    let error = $state('');
    let success = $state('');
    let invitation = $state<any>(null);

    let name = $state('');
    let email = $state('');
    let password = $state('');
    let confirmPassword = $state('');
    let formErrors = $state<Record<string, string>>({});
    let isSubmitting = $state(false);

    let showSignupForm = $state(false);

    let token = $state<string | null>(null);
    let organizationId = $state<string | null>(null); 
    let invitationId = $state<string | null>(null);

    onMount(async () => {
        token = $page.url.searchParams.get('token');
        organizationId = $page.url.searchParams.get('organization');
        invitationId = $page.url.searchParams.get('invitation');

        if (!token || !organizationId || !invitationId) {
            error = 'Invalid invitation link. Please check your email for the correct link.';
            return;
        }

        try {
            loading = true;
            const response = await invitationService.getInvitationByToken(organizationId, invitationId, token);
            invitation = response.data;
            if (invitation) {
                email = invitation.email || '';
                showSignupForm = true;
            } else {
                error = 'Invitation not found or already accepted.';
            }
        } catch (err: any) {
            error = err.message || 'Failed to load invitation';
        } finally {
            loading = false;
        }
    });

    function validateForm(): boolean {
        const newErrors: Record<string, string> = {};

        const nameError = validateRequired(name, 'Name');
        if (nameError) newErrors.name = nameError;

        const emailError = validateRequired(email, 'Email');
        if (emailError) newErrors.email = emailError;
        else if (!validateEmail(email)) newErrors.email = 'Please enter a valid email address';

        const passwordError = validateRequired(password, 'Password');
        if (passwordError) newErrors.password = passwordError;
        else {
            const passwordValidation = validatePassword(password);
            if (!passwordValidation.isValid) newErrors.password = passwordValidation.errors[0];
        }

        const confirmPasswordError = validateConfirmPassword(password, confirmPassword);
        if (confirmPasswordError) newErrors.confirmPassword = confirmPasswordError;

        formErrors = newErrors;
        return Object.keys(newErrors).length === 0;
    }

    async function handleSignup(e: Event) {
        e.preventDefault();
        if (!validateForm()) return;
        isSubmitting = true;
        error = '';
        try {
            await authStore.register(email, password, name);
            const response = await invitationService.acceptInvitation(
                organizationId as string,
                invitationId as string,
                {
                    id: organizationId as string,
                    invitation_id: invitationId as string,
                    token: token as string,
                    email: email as string,
                    password: password as string,
                    name: name as string
                }
            );
            invitation = response.data;
            success = 'Invitation accepted and account created successfully!';
            setTimeout(() => goto('/dashboard'), 2000);
        } catch (err: any) {
            error = err.message || 'Failed to accept invitation or create account';
        } finally {
            isSubmitting = false;
        }
    }
</script>

<div class="min-h-screen flex items-center justify-center p-4">
    <Card classes="max-w-md w-full">
        <div class="text-center space-y-6 p-6">
            <h1 class="h2"><svelte:head>
	<title>Accept Invitation - Consortium</title>
	<meta name="description" content="Accept your invitation to join Consortium." />
</svelte:head>

<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { authStore } from '$lib/stores/auth.js';
	import { Button, Input, Alert } from '$lib/components/ui/index.js';
	import { validatePassword, validateRequired } from '$lib/utils/validation.js';

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
</div></h1>

            {#if loading}
                <div class="flex justify-center">
                    <span class="loading loading-spinner loading-lg" />
                </div>
            {:else if error}
                <Alert type="error" title="Error" message={error} />
                <Button variant="ghost" onclick={() => goto('/')}>
                    <span>Return to Home</span>
                </Button>
            {:else if success}
                <Alert type="success" title="Success" message={success} />
                <p class="text-surface-600-300-token">
                    Redirecting to dashboard...
                </p>
            {:else if showSignupForm}
                <form on:submit={handleSignup} class="space-y-4">
                    <Input
                        type="text"
                        label="Full name"
                        placeholder="Enter your full name"
                        value={name}
                        required
                        error={formErrors.name}
                        oninput={(e) => name = (e.target as HTMLInputElement).value}
                    />
                    <Input
                        type="email"
                        label="Email address"
                        placeholder="Enter your email"
                        value={email}
                        required
                        error={formErrors.email}
                        oninput={(e) => email = (e.target as HTMLInputElement).value}
                        disabled
                    />
                    <Input
                        type="password"
                        label="Password"
                        placeholder="Create a password"
                        value={password}
                        required
                        error={formErrors.password}
                        oninput={(e) => password = (e.target as HTMLInputElement).value}
                    />
                    <Input
                        type="password"
                        label="Confirm password"
                        placeholder="Confirm your password"
                        value={confirmPassword}
                        required
                        error={formErrors.confirmPassword}
                        oninput={(e) => confirmPassword = (e.target as HTMLInputElement).value}
                    />
                    <Button
                        type="submit"
                        variant="filled"
                        size="lg"
                        classes="w-full"
                        loading={isSubmitting}
                        disabled={isSubmitting}
                    >
                        {isSubmitting ? 'Creating account...' : 'Accept & Create Account'}
                    </Button>
                </form>
            {/if}
        </div>
    </Card>
</div>