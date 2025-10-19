// Re-export signupStore from the v1-consortium-web-pkg
export { signupStore } from '@movsm/v1-consortium-web-pkg';
export type { 
	SignupCredentials, 
	SocialSignupCredentials, 
	CompleteRegistrationCredentials,
	SignupWorkflow 
} from '@movsm/v1-consortium-web-pkg';

// Define SubscriptionData locally if not exported
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