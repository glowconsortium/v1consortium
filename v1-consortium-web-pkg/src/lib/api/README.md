# API Service Architecture Changes

## Overview

The API client has been refactored from a monolithic approach to a service-based architecture, and the authentication flow has been updated to use a new two-step signup process.

## Key Changes

### 1. Service-Based Architecture

- **BaseService**: Common functionality for all services (error handling, token management, retry logic)
- **TokenService**: Dedicated service for managing authentication tokens and user session storage
- **AuthService**: Wrapper around the generated AuthService protobuf client
- **ApiClient**: Main client that coordinates all services

### 2. Authentication Flow Changes

#### Old Single-Step Registration (REMOVED)
```typescript
// This is no longer available
await apiClient.register(email, password, firstName, lastName, organizationId, role, invitationToken);
```

#### New Two-Step Signup Process

**Step 1a: Regular Signup**
```typescript
import { apiClient } from '@/lib/api';

// Initialize the client
apiClient.initialize('https://your-api-base-url.com');

// Step 1: Start signup process
const signupResponse = await apiClient.signup(
  'user@example.com',
  'securePassword123',
  'John',
  'Doe',
  'optional-invitation-token'
);

console.log('Workflow ID:', signupResponse.workflowId);
console.log('Message:', signupResponse.message);
console.log('Requires email verification:', signupResponse.requiresEmailVerification);
```

**Step 1b: Social Signup (OAuth)**
```typescript
import { apiClient } from '@/lib/api';

// Initialize the client
apiClient.initialize('https://your-api-base-url.com');

// Step 1: Start social signup process with OAuth provider
const socialSignupResponse = await apiClient.socialSignup(
  'google',              // provider: 'google', 'facebook', 'github', etc.
  'oauth-provider-token', // token from OAuth provider
  'John',                // first name
  'Doe',                 // last name
  'user@example.com',    // email from OAuth provider
  'optional-invitation-token' // optional invitation token
);

console.log('Workflow ID:', socialSignupResponse.workflowId);
console.log('Message:', socialSignupResponse.message);
console.log('Requires email verification:', socialSignupResponse.requiresEmailVerification);
```

**Step 2: Complete Registration**
```typescript
// Step 2: Complete registration after email verification
// This works for both regular signup and social signup
const completionResponse = await apiClient.completeRegistration(
  signupResponse.workflowId, // or socialSignupResponse.workflowId
  'email-verification-token-from-email',
  'organization-id', // optional
  'user-role',       // optional
  'subscription-plan' // optional
);

// User is now logged in with access token
console.log('Access Token:', completionResponse.accessToken);
console.log('User:', completionResponse.user);
```

### 3. Service Usage Examples

#### Using AuthService Directly
```typescript
import { AuthService, TokenService, createApiTransport } from '@/lib/api';

// Create transport and token service
const tokenService = new TokenService();
const transport = createApiTransport(baseUrl, () => tokenService.getAccessToken());

// Create auth service
const authService = new AuthService({
  transport,
  baseUrl,
  getToken: () => tokenService.getAccessToken(),
  setToken: (token) => token ? tokenService.updateAccessToken(token) : tokenService.clearAccessToken(),
  onAuthError: () => tokenService.clearAll()
});

// Use auth service methods
const user = await authService.getCurrentUser();
await authService.changePassword('oldPassword', 'newPassword');
const hasPermission = await authService.checkPermission('user-id', 'read:documents');
```

#### Using ApiClient (Recommended)
```typescript
import { apiClient } from '@/lib/api';

// Initialize once
apiClient.initialize('https://your-api-base-url.com');

// Use convenience methods
await apiClient.login('user@example.com', 'password');

// Regular signup
const signupResponse = await apiClient.signup('user@example.com', 'password', 'John', 'Doe');

// Social signup
const socialSignupResponse = await apiClient.socialSignup('google', 'oauth-token', 'John', 'Doe', 'user@example.com');

// Complete registration (works for both)
await apiClient.completeRegistration(signupResponse.workflowId, 'verification-token');

const user = await apiClient.getCurrentUser();
const authState = apiClient.getAuthState();

// Access specific services
await apiClient.auth.changePassword('oldPassword', 'newPassword');
await apiClient.auth.enableMFA('user-id', 'totp');
```

### 4. Error Handling

The new architecture provides better error handling:

```typescript
try {
  await apiClient.login('user@example.com', 'password');
} catch (error) {
  if (error.message.includes('Authentication failed')) {
    // Handle auth error
  } else if (error.message.includes('Network error')) {
    // Handle network error
  } else {
    // Handle other errors
  }
}
```

### 5. Token Management

Token management is now handled automatically:

```typescript
// Check authentication status
const isAuthenticated = apiClient.getAuthState().isAuthenticated;

// Get current token
const token = apiClient.getToken();

// Check if token is expired
const isExpired = apiClient.isTokenExpired();

// Refresh token automatically happens during API calls
// But you can also do it manually:
const refreshed = await apiClient.refreshAccessToken();
```

## Breaking Changes

1. **Removed `register()` method**: Use the two-step `signup()` + `completeRegistration()` process instead
2. **Service-based architecture**: Direct access to protobuf clients is no longer recommended
3. **Updated imports**: Import from the new service modules

## Migration Guide

### Before
```typescript
import { apiClient } from '@/lib/api';

apiClient.initialize(transport, baseUrl);
await apiClient.register(email, password, firstName, lastName);
```

### After
```typescript
import { apiClient } from '@/lib/api';

apiClient.initialize(baseUrl); // No need to pass transport

// Regular signup
const signupResponse = await apiClient.signup(email, password, firstName, lastName);

// OR Social signup
const socialSignupResponse = await apiClient.socialSignup(provider, providerToken, firstName, lastName, email);

// Handle email verification...
await apiClient.completeRegistration(signupResponse.workflowId, emailVerificationToken);
```

### OAuth Integration Example

For social signup, you'll typically integrate with OAuth providers like this:

```typescript
// Example with Google OAuth
async function handleGoogleSignup() {
  try {
    // Get OAuth token from Google (using your preferred OAuth library)
    const googleResponse = await googleOAuth.getToken();
    
    // Extract user info from Google response
    const { token, user } = googleResponse;
    
    // Start social signup process
    const signupResponse = await apiClient.socialSignup(
      'google',
      token,
      user.given_name,
      user.family_name,
      user.email
    );
    
    // Handle the workflow (email verification might be skipped for verified OAuth emails)
    if (signupResponse.requiresEmailVerification) {
      // Redirect to email verification flow
      window.location.href = `/verify-email?workflow=${signupResponse.workflowId}`;
    } else {
      // Complete registration immediately
      await apiClient.completeRegistration(signupResponse.workflowId, '');
    }
  } catch (error) {
    console.error('Social signup failed:', error);
  }
}
```