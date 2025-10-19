/**
 * Example usage of the new auth service with both regular and social signup
 */

import { apiClient } from './index.js';

// Initialize the API client
apiClient.initialize('https://your-api-base-url.com');

/**
 * Example: Regular email/password signup
 */
export async function handleRegularSignup(
  email: string,
  password: string,
  firstName: string,
  lastName: string,
  invitationToken?: string
) {
  try {
    // Step 1: Start signup process
    const signupResponse = await apiClient.signup(
      email,
      password,
      firstName,
      lastName,
      invitationToken
    );

    console.log('Signup initiated:', {
      workflowId: signupResponse.workflowId,
      message: signupResponse.message,
      requiresEmailVerification: signupResponse.requiresEmailVerification
    });

    // Return workflow ID for the next step
    return signupResponse.workflowId;
  } catch (error) {
    console.error('Regular signup failed:', error);
    throw error;
  }
}

/**
 * Example: Social (OAuth) signup
 */
export async function handleSocialSignup(
  provider: 'google' | 'facebook' | 'github' | 'linkedin',
  providerToken: string,
  userInfo: {
    firstName: string;
    lastName: string;
    email: string;
  },
  invitationToken?: string
) {
  try {
    // Step 1: Start social signup process
    const socialSignupResponse = await apiClient.socialSignup(
      provider,
      providerToken,
      userInfo.firstName,
      userInfo.lastName,
      userInfo.email,
      invitationToken
    );

    console.log('Social signup initiated:', {
      workflowId: socialSignupResponse.workflowId,
      message: socialSignupResponse.message,
      requiresEmailVerification: socialSignupResponse.requiresEmailVerification
    });

    // Return workflow ID for the next step
    return socialSignupResponse.workflowId;
  } catch (error) {
    console.error('Social signup failed:', error);
    throw error;
  }
}

/**
 * Example: Complete registration (works for both regular and social signup)
 */
export async function completeRegistration(
  workflowId: string,
  emailVerificationToken: string,
  options?: {
    organizationId?: string;
    role?: string;
    subscriptionPlan?: string;
  }
) {
  try {
    const completionResponse = await apiClient.completeRegistration(
      workflowId,
      emailVerificationToken,
      options?.organizationId,
      options?.role,
      options?.subscriptionPlan
    );

    console.log('Registration completed:', {
      userId: completionResponse.user?.userId,
      accessToken: completionResponse.accessToken ? 'Present' : 'Missing',
      sessionId: completionResponse.sessionId
    });

    // User is now logged in and can use authenticated endpoints
    const currentUser = await apiClient.getCurrentUser();
    console.log('Current user:', currentUser);

    return completionResponse;
  } catch (error) {
    console.error('Complete registration failed:', error);
    throw error;
  }
}

/**
 * Example: Google OAuth integration
 */
export async function handleGoogleOAuth() {
  try {
    // This is pseudo-code - you'd use your actual OAuth library
    // const googleAuth = new GoogleAuth();
    // const googleResponse = await googleAuth.signIn();
    
    // Mock Google response for example
    const mockGoogleResponse = {
      token: 'google-oauth-token-here',
      user: {
        given_name: 'John',
        family_name: 'Doe',
        email: 'john.doe@gmail.com'
      }
    };

    // Start social signup with Google
    const workflowId = await handleSocialSignup(
      'google',
      mockGoogleResponse.token,
      {
        firstName: mockGoogleResponse.user.given_name,
        lastName: mockGoogleResponse.user.family_name,
        email: mockGoogleResponse.user.email
      }
    );

    // For social logins, email verification might be skipped
    // Complete registration immediately if no verification needed
    await completeRegistration(workflowId, '');

  } catch (error) {
    console.error('Google OAuth signup failed:', error);
  }
}

/**
 * Example: Traditional signup flow
 */
export async function handleTraditionalSignup() {
  try {
    // Step 1: Regular signup
    const workflowId = await handleRegularSignup(
      'user@example.com',
      'securePassword123',
      'Jane',
      'Smith'
    );

    // Step 2: User receives email and gets verification token
    // This would typically happen on a verification page
    // const verificationToken = 'token-from-email-link';
    
    // Step 3: Complete registration
    // await completeRegistration(workflowId, verificationToken);

    console.log('Check your email for verification instructions');
    
  } catch (error) {
    console.error('Traditional signup failed:', error);
  }
}

/**
 * Example: Login after registration
 */
export async function handleLogin() {
  try {
    const loginResponse = await apiClient.login(
      'user@example.com',
      'securePassword123',
      true // remember me
    );

    console.log('Login successful:', {
      userId: loginResponse.user?.userId,
      accessToken: loginResponse.accessToken ? 'Present' : 'Missing'
    });

    // Get current user info
    const user = await apiClient.getCurrentUser();
    console.log('Current user:', user);

  } catch (error) {
    console.error('Login failed:', error);
  }
}

/**
 * Example: Check authentication status
 */
export function checkAuthStatus() {
  const authState = apiClient.getAuthState();
  
  console.log('Authentication status:', {
    isAuthenticated: authState.isAuthenticated,
    hasToken: !!authState.token,
    hasUser: !!authState.user,
    isTokenExpired: apiClient.isTokenExpired()
  });

  return authState;
}