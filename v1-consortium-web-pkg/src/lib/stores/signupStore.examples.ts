/**
 * Example usage of the new signup store
 */

import { 
	signupStore,
	signupCurrentStep,
	signupEmail,
	signupError,
	signupLoading,
	signupMessage,
	signupWorkflowId,
	signupRequiresEmailVerification,
	signupType
} from '../stores/signupStore.js';
import type { 
	SignupCredentials,
	SocialSignupCredentials,
	CompleteRegistrationCredentials
} from '../stores/signupStore.js';

// Import auth store for proper initialization order
import { authStore } from '../stores/auth.js';

// Initialize the signup store (assumes auth store is already initialized)
export async function initializeSignup() {
	await signupStore.initialize();
}

// Initialize both auth and signup stores in the correct order
export async function initializeAppStores(config: { baseUrl: string }) {
	try {
		// Step 1: Initialize auth store with base URL (this sets up the API client)
		await authStore.initialize(config);
		
		// Step 2: Initialize signup store (this just verifies auth is ready)
		await initializeSignup();
		
		console.log('All stores initialized successfully');
	} catch (error) {
		console.error('Store initialization failed:', error);
		throw error;
	}
}

// Example: Regular email/password signup
export async function handleRegularSignup(credentials: SignupCredentials) {
	try {
		console.log('Starting regular signup...');
		
		const response = await signupStore.signup(credentials);
		
		console.log('Signup initiated:', {
			workflowId: response.workflowId,
			message: response.message,
			requiresEmailVerification: response.requiresEmailVerification
		});

		// Listen to store updates
		signupCurrentStep.subscribe(step => {
			console.log('Current step:', step);
		});

		signupMessage.subscribe(message => {
			if (message) console.log('Message:', message);
		});

		return response;
	} catch (error) {
		console.error('Regular signup failed:', error);
		throw error;
	}
}

// Example: Social signup with Google
export async function handleGoogleSignup(userInfo: {
	firstName: string;
	lastName: string;
	email: string;
	googleToken: string;
	invitationToken?: string;
}) {
	try {
		console.log('Starting Google signup...');

		const credentials: SocialSignupCredentials = {
			provider: 'google',
			providerToken: userInfo.googleToken,
			firstName: userInfo.firstName,
			lastName: userInfo.lastName,
			email: userInfo.email,
			invitationToken: userInfo.invitationToken
		};

		const response = await signupStore.socialSignup(credentials);

		console.log('Social signup initiated:', {
			workflowId: response.workflowId,
			message: response.message,
			requiresEmailVerification: response.requiresEmailVerification
		});

		return response;
	} catch (error) {
		console.error('Google signup failed:', error);
		throw error;
	}
}

// Example: Complete registration
export async function handleCompleteRegistration(
	workflowId: string,
	emailVerificationToken: string,
	options?: {
		organizationId?: string;
		role?: string;
		subscriptionPlan?: string;
	}
) {
	try {
		console.log('Completing registration...');

		const credentials: CompleteRegistrationCredentials = {
			workflowId,
			emailVerificationToken,
			organizationId: options?.organizationId,
			role: options?.role,
			subscriptionPlan: options?.subscriptionPlan
		};

		const response = await signupStore.completeRegistration(credentials);

		console.log('Registration completed:', {
			userId: response.user?.userId,
			accessToken: response.accessToken ? 'Present' : 'Missing',
			sessionId: response.sessionId
		});

		return response;
	} catch (error) {
		console.error('Complete registration failed:', error);
		throw error;
	}
}

// Example: Handle email verification
export async function handleEmailVerification(
	workflowId: string,
	verificationToken: string
) {
	try {
		console.log('Verifying email...');
		
		await signupStore.verifyEmail(workflowId, verificationToken);
		
		console.log('Email verified successfully');
	} catch (error) {
		console.error('Email verification failed:', error);
		throw error;
	}
}

// Example: Get signup status
export async function checkSignupStatus(workflowId: string) {
	try {
		console.log('Checking signup status...');
		
		const status = await signupStore.getSignupStatus(workflowId);
		
		console.log('Signup status:', {
			workflowId: status.workflowId,
			email: status.email,
			emailVerified: status.emailVerified,
			status: status.status
		});

		return status;
	} catch (error) {
		console.error('Failed to get signup status:', error);
		throw error;
	}
}

// Example: Resend verification email
export async function resendVerificationEmail(workflowId: string) {
	try {
		console.log('Resending verification email...');
		
		await signupStore.resendVerification(workflowId);
		
		console.log('Verification email sent successfully');
	} catch (error) {
		console.error('Failed to resend verification email:', error);
		throw error;
	}
}

// Example: Reactive signup flow component logic
export function createSignupFlowSubscriptions() {
	const subscriptions = [
		// Monitor current step
		signupCurrentStep.subscribe(step => {
			switch (step) {
				case 'signup':
					console.log('Ready for signup');
					break;
				case 'verify-email':
					console.log('Waiting for email verification');
					break;
				case 'complete-registration':
					console.log('Ready to complete registration');
					break;
				case 'completed':
					console.log('Signup completed successfully');
					break;
			}
		}),

		// Monitor loading state
		signupLoading.subscribe(isLoading => {
			console.log('Loading:', isLoading);
		}),

		// Monitor errors
		signupError.subscribe(error => {
			if (error) {
				console.error('Signup error:', error);
				// Show error notification to user
			}
		}),

		// Monitor messages
		signupMessage.subscribe(message => {
			if (message) {
				console.log('Signup message:', message);
				// Show info message to user
			}
		})
	];

	// Return cleanup function
	return () => {
		subscriptions.forEach(unsubscribe => unsubscribe());
	};
}

// Example: Complete signup flow (assumes auth store is already initialized)
export async function completeSignupFlow() {
	try {
		// Step 1: Initialize signup store (auth must be initialized first)
		await initializeSignup();

		// Step 2: Regular signup
		const signupResponse = await handleRegularSignup({
			email: 'user@example.com',
			password: 'securePassword123',
			firstName: 'John',
			lastName: 'Doe'
		});

		// Step 3: Wait for email verification (this would be handled by user clicking email link)
		// In a real app, this would happen on a verification page
		console.log('Please check your email and click the verification link');

		// Step 4: Complete registration (after email verification)
		// This would be called from the email verification page
		// await handleCompleteRegistration(
		//   signupResponse.workflowId,
		//   'verification-token-from-email'
		// );

	} catch (error) {
		console.error('Signup flow failed:', error);
	}
}

// Example: OAuth signup flow with Google (assumes auth store is already initialized)
export async function completeGoogleSignupFlow() {
	try {
		// Step 1: Initialize signup store (auth must be initialized first)
		await initializeSignup();

		// Step 2: Handle OAuth callback (this would come from your OAuth library)
		const mockGoogleUserInfo = {
			firstName: 'Jane',
			lastName: 'Smith',
			email: 'jane.smith@gmail.com',
			googleToken: 'google-oauth-token-here'
		};

		// Step 3: Social signup
		const socialSignupResponse = await handleGoogleSignup(mockGoogleUserInfo);

		// Step 4: Complete registration (might skip email verification for verified OAuth emails)
		if (!socialSignupResponse.requiresEmailVerification) {
			await handleCompleteRegistration(
				socialSignupResponse.workflowId,
				'' // Empty token for OAuth signups that don't need email verification
			);
		} else {
			console.log('Please check your email and click the verification link');
		}

	} catch (error) {
		console.error('Google signup flow failed:', error);
	}
}

// Utility functions for managing signup state
export const signupUtils = {
	// Clear any error messages
	clearError: () => signupStore.clearError(),
	
	// Clear any info messages
	clearMessage: () => signupStore.clearMessage(),
	
	// Reset the entire signup flow
	resetFlow: () => signupStore.reset(),
	
	// Set current step manually (if needed)
	setStep: (step: 'signup' | 'verify-email' | 'complete-registration' | 'completed') => 
		signupStore.setCurrentStep(step),
	
	// Update workflow state manually (if needed)
	updateState: (updates: any) => signupStore.updateWorkflowState(updates)
};