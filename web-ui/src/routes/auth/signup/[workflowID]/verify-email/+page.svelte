<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { signupStore } from '$lib/store/signupStore';
	import { Button, Input } from '@movsm/v1-consortium-web-pkg';

	// Get workflow ID from URL params
	let workflowId = $page.params.workflowID;
	let verificationCode = $state('');
	let isSubmitting = $state(false);
	let error = $state<string | null>(null);
	let resendLoading = $state(false);
	let showManualEntry = $state(false);

	onMount(async () => {
		if (!workflowId) {
			error = 'Invalid workflow ID';
			return;
		}

		// Check if we have a verification code in URL (from email link)
		const urlCode = $page.url.searchParams.get('code');
		if (urlCode) {
			// Auto-verify if code is in URL
			await handleVerification(urlCode);
		} else {
			// Load workflow state
			try {
				await signupStore.loadWorkflow(workflowId);
			} catch (err) {
				console.error('Failed to load workflow:', err);
				error = 'Invalid verification link';
			}
		}
	});

	async function handleVerification(code: string = verificationCode) {
		if (!workflowId) {
			error = 'Invalid workflow ID';
			return;
		}

		if (!code.trim()) {
			error = 'Please enter the verification code';
			return;
		}

		isSubmitting = true;
		error = null;

		try {
			await signupStore.verifyEmail(workflowId, code.trim());
			// Redirect to subscription step
			goto(`/auth/signup/${workflowId}/subscribe`);
		} catch (err) {
			console.error('Email verification failed:', err);
			error = err instanceof Error ? err.message : 'Verification failed';
		} finally {
			isSubmitting = false;
		}
	}

	async function handleResendEmail() {
		if (!workflowId) {
			error = 'Invalid workflow ID';
			return;
		}

		resendLoading = true;
		
		try {
			await signupStore.resendVerification(workflowId);
			// Show success message (you could add a toast notification here)
			alert('Verification email sent! Please check your inbox.');
		} catch (err) {
			console.error('Failed to resend email:', err);
			error = err instanceof Error ? err.message : 'Failed to resend email. Please try again.';
		} finally {
			resendLoading = false;
		}
	}

	function handleCodeInput(event: Event) {
		verificationCode = (event.target as HTMLInputElement).value;
		if (error) error = null;
	}
</script>

<svelte:head>
	<title>Verify Email - Consortium</title>
</svelte:head>

<div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-md">
		<div class="text-center mb-8">
			<h1 class="text-2xl font-bold text-blue-600">Consortium</h1>
		</div>

		<div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
			<div class="text-center mb-6">
				<div class="mx-auto flex items-center justify-center w-12 h-12 rounded-full bg-blue-100 mb-4">
					<svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
					</svg>
				</div>
				<h2 class="text-2xl font-semibold text-gray-900">Verify Your Email</h2>
				<p class="text-gray-600 mt-2">
					We've sent a verification email to:<br>
					<span class="font-medium">{$signupStore.email}</span>
				</p>
			</div>

			<div class="space-y-6">
				<p class="text-sm text-gray-600 text-center">
					Click the link in your email to verify your account, or enter the verification code below.
				</p>

				{#if !showManualEntry}
					<div class="text-center">
						<button 
							class="text-blue-600 hover:text-blue-500 text-sm font-medium"
							onclick={() => showManualEntry = true}
						>
							Enter verification code manually
						</button>
					</div>
				{:else}
					<form onsubmit={() => handleVerification()}>
						<div class="mb-4">
							<label for="verificationCode" class="block text-sm font-medium text-gray-700">
								Verification Code
							</label>
							<div class="mt-1">
								<Input
									id="verificationCode"
									type="text"
									value={verificationCode}
									oninput={handleCodeInput}
									placeholder="Enter 6-digit code"
									required
								/>
							</div>
						</div>

						{#if error}
							<div class="mb-4 p-3 bg-red-50 border border-red-200 rounded-md">
								<p class="text-red-600 text-sm">{error}</p>
							</div>
						{/if}

						<Button
							type="submit"
							color="primary"
							size="lg"
							classes="w-full"
							loading={isSubmitting}
							disabled={!verificationCode.trim()}
						>
							Verify Email
						</Button>
					</form>
				{/if}

				<div class="space-y-3">
					<div class="text-center">
						<span class="text-sm text-gray-600">Didn't receive the email?</span>
					</div>
					
					<Button
						color="secondary"
						size="md"
						classes="w-full"
						loading={resendLoading}
						onclick={handleResendEmail}
					>
						Resend Email
					</Button>
				</div>

				<div class="bg-gray-50 rounded-lg p-4">
					<h4 class="font-medium text-gray-900 mb-2">Check your spam folder</h4>
					<p class="text-sm text-gray-600">
						If you don't see the email in your inbox, please check your spam or junk folder.
					</p>
				</div>
			</div>

			<!-- Help and support -->
			<div class="mt-6 text-center">
				<p class="text-xs text-gray-500">
					Need help? 
					<a href="/contact" class="text-blue-600 hover:text-blue-500">Contact Support</a> or 
					<a href="/auth/signin" class="text-blue-600 hover:text-blue-500">Back to Sign In</a>
				</p>
			</div>
		</div>
	</div>
</div>
