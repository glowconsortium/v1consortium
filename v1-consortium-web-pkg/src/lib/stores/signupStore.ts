import { writable, derived } from 'svelte/store';
import { browser } from '$app/environment';
import { apiClient } from '../api/api.js';
import type { 
	SignupResponse, 
	SocialSignupResponse, 
	CompleteRegistrationResponse 
} from '../gen/auth/v1/auth_pb.js';

export interface SignupCredentials {
	email: string;
	password: string;
	firstName: string;
	lastName: string;
	companyName?: string;
	isDotCompany?: boolean;
	dotNumber?: string;
	invitationToken?: string;
}

export interface SocialSignupCredentials {
	provider: 'google' | 'facebook' | 'github' | 'linkedin';
	providerToken: string;
	firstName: string;
	lastName: string;
	email: string;
	invitationToken?: string;
}

export interface CompleteRegistrationCredentials {
	workflowId: string;
	emailVerificationToken: string;
	organizationId?: string;
	role?: string;
	subscriptionPlan?: string;
}

export interface SubscriptionData {
	accountType: 'company' | 'individual';
	userInfo: {
		firstName: string;
		lastName: string;
	};
	companyInfo?: {
		name: string;
		dotNumber?: string;
		mcNumber?: string;
		phone?: string;
	};
	addressInfo: {
		line1: string;
		line2?: string;
		city: string;
		state: string;
		zipCode: string;
	};
	planInfo: {
		selectedPlan: 'basic' | 'premium';
	};
	paymentInfo?: {
		paymentMethodId: string;
		customerId: string;
	};
}

export interface SignupWorkflow {
	workflowId: string | null;
	email: string;
	isEmailVerified: boolean;
	currentStep: 'signup' | 'verify-email' | 'complete-registration' | 'subscription' | 'completed';
	signupType: 'regular' | 'social' | null;
	requiresEmailVerification: boolean;
	message: string | null;
	error: string | null;
	isLoading: boolean;
	isInitialized: boolean;
	accountType?: 'company' | 'individual';
	firstName?: string;
	lastName?: string;
	companyName?: string;
	subscriptionPlan?: 'basic' | 'premium';
}

// Initial state
const initialState: SignupWorkflow = {
	workflowId: null,
	email: '',
	isEmailVerified: false,
	currentStep: 'signup',
	signupType: null,
	requiresEmailVerification: false,
	message: null,
	error: null,
	isLoading: false,
	isInitialized: false,
	accountType: undefined,
	firstName: undefined,
	lastName: undefined,
	companyName: undefined,
	subscriptionPlan: undefined
};

// Create the internal store
const _signupStore = writable<SignupWorkflow>(initialState);

// Create a custom store with methods
function createSignupStore() {
	const { subscribe, set, update } = _signupStore;

	const store = {
		subscribe,

		// Initialize signup store (assumes auth store has already initialized API client)
		async initialize(): Promise<void> {
			if (!browser) return;

			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true 
			}));

			try {
				// Check if API client is initialized by auth store
				const authState = apiClient.getAuthState();
				
				// Verify API client is working
				if (typeof apiClient.getToken !== 'function') {
					throw new Error('API client not properly initialized. Please initialize auth store first.');
				}
				
				update((state: SignupWorkflow) => ({ 
					...state, 
					isLoading: false, 
					isInitialized: true 
				}));
			} catch (error) {
				console.error('Failed to initialize signup store:', error);
				update((state: SignupWorkflow) => ({ 
					...state, 
					error: error instanceof Error ? error.message : 'Signup store initialization failed - auth store must be initialized first',
					isLoading: false,
					isInitialized: true
				}));
			}
		},

		// Initiate signup (simplified method expected by web-ui)
		async initiateSignup(email: string, password: string): Promise<string> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				const response = await apiClient.signup(
					email,
					password,
					'', // firstName - will be collected later
					'', // lastName - will be collected later
					undefined // invitationToken
				);

				update((state: SignupWorkflow) => ({
					...state,
					workflowId: response.workflowId,
					email: email,
					signupType: 'regular',
					requiresEmailVerification: response.requiresEmailVerification,
					message: response.message,
					currentStep: response.requiresEmailVerification ? 'verify-email' : 'subscription',
					isLoading: false,
					error: null
				}));

				return response.workflowId;
			} catch (error: any) {
				const errorMessage = error.message || 'Signup failed';
				console.error('Signup error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Handle social OAuth callback
		async handleSocialCallback(provider: string, code: string): Promise<string> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				// Convert the OAuth code to a provider token (this would typically involve
				// exchanging the code for a token with the OAuth provider)
				// For now, we'll use the code as the token
				const response = await apiClient.socialSignup(
					provider as 'google' | 'facebook' | 'github' | 'linkedin',
					code, // Using code as token - in real implementation, exchange for token first
					'', // firstName - will be collected later
					'', // lastName - will be collected later  
					'', // email - will be extracted from provider
					undefined // invitationToken
				);

				update((state: SignupWorkflow) => ({
					...state,
					workflowId: response.workflowId,
					signupType: 'social',
					requiresEmailVerification: response.requiresEmailVerification,
					message: response.message,
					currentStep: response.requiresEmailVerification ? 'verify-email' : 'subscription',
					isLoading: false,
					error: null
				}));

				return response.workflowId;
			} catch (error: any) {
				const errorMessage = error.message || 'Social signup failed';
				console.error('Social signup error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Load workflow state by ID
		async loadWorkflow(workflowId: string): Promise<void> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				// Use the getSignupStatus method from auth service
				const response = await apiClient.auth.getSignupStatus(workflowId);

				update((state: SignupWorkflow) => ({
					...state,
					workflowId: response.workflowId,
					email: response.email,
					isEmailVerified: response.emailVerified,
					requiresEmailVerification: !response.emailVerified,
					currentStep: response.emailVerified ? 'subscription' : 'verify-email',
					isLoading: false,
					error: null
				}));
			} catch (error: any) {
				const errorMessage = error.message || 'Failed to load workflow';
				console.error('Load workflow error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Complete subscription step
		async completeSubscription(workflowId: string, subscriptionData: SubscriptionData): Promise<void> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				// Complete registration with subscription data
				const response = await apiClient.completeRegistration(
					workflowId,
					'', // emailVerificationToken - already verified at this point
					undefined, // organizationId
					subscriptionData.accountType === 'company' ? 'company_admin' : 'individual_user',
					subscriptionData.planInfo.selectedPlan
				);

				update((state: SignupWorkflow) => ({
					...state,
					currentStep: 'completed',
					isEmailVerified: true,
					accountType: subscriptionData.accountType,
					firstName: subscriptionData.userInfo.firstName,
					lastName: subscriptionData.userInfo.lastName,
					companyName: subscriptionData.companyInfo?.name,
					subscriptionPlan: subscriptionData.planInfo.selectedPlan,
					isLoading: false,
					error: null
				}));
			} catch (error: any) {
				const errorMessage = error.message || 'Subscription completion failed';
				console.error('Complete subscription error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Start regular email/password signup (with full credentials)
		async signup(credentials: SignupCredentials): Promise<SignupResponse> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				const response = await apiClient.signup(
					credentials.email,
					credentials.password,
					credentials.firstName,
					credentials.lastName,
					credentials.invitationToken,
					credentials.companyName,
					credentials.isDotCompany,
					credentials.dotNumber
				);

				update((state: SignupWorkflow) => ({
					...state,
					workflowId: response.workflowId,
					email: credentials.email,
					firstName: credentials.firstName,
					lastName: credentials.lastName,
					signupType: 'regular',
					requiresEmailVerification: response.requiresEmailVerification,
					message: response.message,
					currentStep: response.requiresEmailVerification ? 'verify-email' : 'subscription',
					isLoading: false,
					error: null
				}));

				return response;
			} catch (error: any) {
				const errorMessage = error.message || 'Signup failed';
				console.error('Signup error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Start social signup (OAuth) with full credentials
		async socialSignup(credentials: SocialSignupCredentials): Promise<SocialSignupResponse> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				const response = await apiClient.socialSignup(
					credentials.provider,
					credentials.providerToken,
					credentials.firstName,
					credentials.lastName,
					credentials.email,
					credentials.invitationToken
				);

				update((state: SignupWorkflow) => ({
					...state,
					workflowId: response.workflowId,
					email: credentials.email,
					firstName: credentials.firstName,
					lastName: credentials.lastName,
					signupType: 'social',
					requiresEmailVerification: response.requiresEmailVerification,
					message: response.message,
					currentStep: response.requiresEmailVerification ? 'verify-email' : 'subscription',
					isLoading: false,
					error: null
				}));

				return response;
			} catch (error: any) {
				const errorMessage = error.message || 'Social signup failed';
				console.error('Social signup error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Complete registration (final step)
		async completeRegistration(credentials: CompleteRegistrationCredentials): Promise<CompleteRegistrationResponse> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				const response = await apiClient.completeRegistration(
					credentials.workflowId,
					credentials.emailVerificationToken,
					credentials.organizationId,
					credentials.role,
					credentials.subscriptionPlan
				);

				update((state: SignupWorkflow) => ({
					...state,
					currentStep: 'completed',
					isEmailVerified: true,
					isLoading: false,
					error: null
				}));

				return response;
			} catch (error: any) {
				const errorMessage = error.message || 'Registration completion failed';
				console.error('Complete registration error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Verify email using API (if you have a separate endpoint)
		async verifyEmail(workflowId: string, verificationToken: string): Promise<void> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				// Use the verify email method from auth service
				await apiClient.auth.verifyEmail(verificationToken);

				update((state: SignupWorkflow) => ({
					...state,
					isEmailVerified: true,
					currentStep: 'complete-registration',
					isLoading: false,
					error: null
				}));
			} catch (error: any) {
				const errorMessage = error.message || 'Email verification failed';
				console.error('Email verification error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Get signup status (check workflow state)
		async getSignupStatus(workflowId: string) {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				// Use the getSignupStatus method from auth service
				const response = await apiClient.auth.getSignupStatus(workflowId);

				update((state: SignupWorkflow) => ({
					...state,
					workflowId: response.workflowId,
					email: response.email,
					isEmailVerified: response.emailVerified,
					requiresEmailVerification: !response.emailVerified,
					currentStep: response.emailVerified ? 'complete-registration' : 'verify-email',
					isLoading: false,
					error: null
				}));

				return response;
			} catch (error: any) {
				const errorMessage = error.message || 'Failed to get signup status';
				console.error('Get signup status error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Resend verification email
		async resendVerification(workflowId: string): Promise<void> {
			update((state: SignupWorkflow) => ({ 
				...state, 
				isLoading: true, 
				error: null 
			}));

			try {
				// Use the resendVerification method from auth service
				await apiClient.auth.resendVerification(workflowId);

				update((state: SignupWorkflow) => ({
					...state,
					message: 'Verification email sent successfully',
					isLoading: false,
					error: null
				}));
			} catch (error: any) {
				const errorMessage = error.message || 'Failed to resend verification email';
				console.error('Resend verification error:', error);
				update((state: SignupWorkflow) => ({
					...state,
					isLoading: false,
					error: errorMessage
				}));
				throw error;
			}
		},

		// Update workflow state manually
		updateWorkflowState(updates: Partial<SignupWorkflow>) {
			update((state: SignupWorkflow) => ({ ...state, ...updates }));
		},

		// Set current step
		setCurrentStep(step: SignupWorkflow['currentStep']) {
			update((state: SignupWorkflow) => ({ ...state, currentStep: step }));
		},

		// Clear error
		clearError(): void {
			update((state: SignupWorkflow) => ({ ...state, error: null }));
		},

		// Clear message
		clearMessage(): void {
			update((state: SignupWorkflow) => ({ ...state, message: null }));
		},

		// Reset store to initial state
		reset(): void {
			set(initialState);
		}
	};

	return store;
}

// Create and export the signup store
export const signupStore = createSignupStore();

// Derived stores for easy access to signup state
export const signupWorkflowId = derived(_signupStore, $signup => $signup.workflowId);
export const signupCurrentStep = derived(_signupStore, $signup => $signup.currentStep);
export const signupEmail = derived(_signupStore, $signup => $signup.email);
export const signupError = derived(_signupStore, $signup => $signup.error);
export const signupLoading = derived(_signupStore, $signup => $signup.isLoading);
export const signupMessage = derived(_signupStore, $signup => $signup.message);
export const signupRequiresEmailVerification = derived(_signupStore, $signup => $signup.requiresEmailVerification);
export const signupType = derived(_signupStore, $signup => $signup.signupType);

// Note: Auth store must be initialized first with base URL before using signup store
// Initialize auth store: await authStore.initialize({ baseUrl: 'your-api-url' })
// Then initialize signup store: await signupStore.initialize()
// 
// Usage patterns:
// - signupStore.initiateSignup(email, password) - starts signup with minimal info
// - signupStore.handleSocialCallback(provider, code) - handles OAuth callbacks  
// - signupStore.loadWorkflow(workflowId) - loads existing workflow state
// - signupStore.verifyEmail(workflowId, token) - verifies email
// - signupStore.completeSubscription(workflowId, subscriptionData) - completes signup
// 
// For more complete signup flows:
// - signupStore.signup(credentials) - full signup with all user details
// - signupStore.socialSignup(credentials) - full social signup
// - signupStore.completeRegistration(credentials) - complete registration step

// Usage for simplified signup flow (web-ui):
// 1. Use signupStore.initiateSignup(email, password) for quick signup
// 2. Use signupStore.loadWorkflow(workflowId) to load workflow state
// 3. Use signupStore.verifyEmail(workflowId, token) for email verification
// 4. Use signupStore.completeSubscription(workflowId, subscriptionData) to complete signup

// Usage for full signup flow:
// 1. Use signupStore.signup(credentials) for regular email/password signup
// 2. Use signupStore.socialSignup(credentials) for OAuth-based signup
// 3. Use signupStore.completeRegistration(credentials) to finalize registration
