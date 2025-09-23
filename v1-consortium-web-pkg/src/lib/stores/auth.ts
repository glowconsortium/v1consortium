import { writable, derived } from 'svelte/store';
import { browser } from '$app/environment';
import { AuthService } from '../backendapi/generated-backendapi/index.js';
import { tokenManager } from '../utils/tokenManager.js';
import type { 
    User, 
    LoginRequest, 
    RegisterRequest, 
    LoginResponse, 
    RegisterResponse 
} from '../types/auth.js';

interface AuthState {
    user: User | null;
    isAuthenticated: boolean;
    isLoading: boolean;
    error: string | null;
    isInitialized: boolean;
}

// Initial state
const initialState: AuthState = {
    user: null,
    isAuthenticated: false,
    isLoading: false,
    error: null,
    isInitialized: false
};

// Create the internal store
const _authStore = writable<AuthState>(initialState);

// Create a custom store with methods
function createAuthStore() {
    const { subscribe, set, update } = _authStore;

    return {
        subscribe,
        
        // Initialize auth state on app start
        async initialize() {
            if (!browser) return;

            update((state: AuthState) => ({ ...state, isLoading: true }));

            try {
                const tokens = tokenManager.getTokens();
                
                if (tokens?.accessToken && !tokenManager.isTokenExpired()) {
                    // Try to get current user with existing token
                    await this.getCurrentUser();
                }
            } catch (error) {
                console.warn('Failed to initialize auth state:', error);
                tokenManager.clearTokens();
            } finally {
                update((state: AuthState) => ({ 
                    ...state, 
                    isLoading: false, 
                    isInitialized: true 
                }));
            }
        },

        // Login user
        async login(email: string, password: string): Promise<LoginResponse> {
            update((state: AuthState) => ({ ...state, isLoading: true, error: null }));

            try {
                const loginRequest: LoginRequest = { email, password };
                const response = await AuthService.postApiPublicV1Login(loginRequest);

                if (response.data?.accessToken) {
                    // Store tokens
                    tokenManager.setTokens({
                        accessToken: response.data.accessToken,
                        refreshToken: response.data.refreshToken,
                        expiresAt: Date.now() + (60 * 60 * 1000), // Default 1 hour
                        tokenType: 'Bearer'
                    });

                    // Create user object from response
                    const user: User = {
                        id: response.data.userId || '',
                        email: email,
                        name: email.split('@')[0], // Fallback name
                        role: response.data.role,
                        carrierId: response.data.carrierId
                    };

                    update((state: AuthState) => ({
                        ...state,
                        user,
                        isAuthenticated: true,
                        isLoading: false,
                        error: null
                    }));

                    return response;
                } else {
                    throw new Error(response.message || 'Login failed');
                }
            } catch (error: any) {
                const errorMessage = error.message || 'Login failed. Please check your credentials.';
                update((state: AuthState) => ({
                    ...state,
                    isLoading: false,
                    error: errorMessage,
                    isAuthenticated: false,
                    user: null
                }));
                throw error;
            }
        },

        // Register new user
        async register(email: string, password: string, name?: string, companyData?: any): Promise<RegisterResponse> {
            update((state: AuthState) => ({ ...state, isLoading: true, error: null }));

            try {
                const registerRequest: RegisterRequest = {
                    email,
                    password,
                    name,
                    ...companyData
                };
                
                const response = await AuthService.postApiPublicV1Signup(registerRequest);

                if (response.data?.success) {
                    // If signup returns tokens, log the user in automatically
                    if (response.data.accessToken) {
                        tokenManager.setTokens({
                            accessToken: response.data.accessToken,
                            refreshToken: response.data.refreshToken,
                            expiresAt: Date.now() + (60 * 60 * 1000),
                            tokenType: 'Bearer'
                        });

                        const user: User = {
                            id: response.data.userId || '',
                            email: email,
                            name: name || email.split('@')[0],
                            carrierId: response.data.carrierId
                        };

                        update((state: AuthState) => ({
                            ...state,
                            user,
                            isAuthenticated: true,
                            isLoading: false,
                            error: null
                        }));
                    } else {
                        // Registration successful but no auto-login
                        update((state: AuthState) => ({
                            ...state,
                            isLoading: false,
                            error: null
                        }));
                    }

                    return response;
                } else {
                    throw new Error(response.message || 'Registration failed');
                }
            } catch (error: any) {
                const errorMessage = error.message || 'Registration failed. Please try again.';
                update((state: AuthState) => ({
                    ...state,
                    isLoading: false,
                    error: errorMessage
                }));
                throw error;
            }
        },

        // Logout user
        async logout(): Promise<void> {
            update((state: AuthState) => ({ ...state, isLoading: true }));

            try {
                // Call logout endpoint if we have a token
                const tokens = tokenManager.getTokens();
                if (tokens?.refreshToken) {
                    await AuthService.postApiPublicV1Logout({
                        refreshToken: tokens.refreshToken
                    });
                }
            } catch (error) {
                console.warn('Logout API call failed:', error);
            } finally {
                // Always clear local state regardless of API call result
                tokenManager.clearTokens();
                update((state: AuthState) => ({
                    ...state,
                    user: null,
                    isAuthenticated: false,
                    isLoading: false,
                    error: null
                }));
            }
        },

        // Refresh session
        async refreshSession(): Promise<boolean> {
            try {
                const tokens = tokenManager.getTokens();
                if (!tokens?.refreshToken) {
                    throw new Error('No refresh token available');
                }

                const response = await AuthService.postApiPublicV1Refresh({
                    refreshToken: tokens.refreshToken
                });

                if (response.data?.accessToken) {
                    tokenManager.setTokens({
                        accessToken: response.data.accessToken,
                        refreshToken: response.data.refreshToken || tokens.refreshToken,
                        expiresAt: Date.now() + (60 * 60 * 1000),
                        tokenType: 'Bearer'
                    });

                    return true;
                } else {
                    throw new Error('Failed to refresh token');
                }
            } catch (error) {
                console.warn('Token refresh failed:', error);
                await this.logout();
                return false;
            }
        },

        // Get current user
        async getCurrentUser(): Promise<User | null> {
            try {
                const tokens = tokenManager.getTokens();
                if (!tokens?.accessToken) {
                    throw new Error('No access token available');
                }

                // This would call a profile endpoint to get user details
                // For now, we'll use stored user data or create minimal user from token
                // If we don't have user data, we might need to implement a getCurrentUser API call
                // For now, return null and let the user re-authenticate
                return null;
            } catch (error) {
                console.warn('Failed to get current user:', error);
                return null;
            }
        },

        // Verify email
        async verifyEmail(token: string): Promise<void> {
            update((state: AuthState) => ({ ...state, isLoading: true, error: null }));

            try {
                const response = await AuthService.postApiPublicV1VerifyEmail({ token });
                
                if (response.data?.success) {
                    update((state: AuthState) => ({
                        ...state,
                        isLoading: false,
                        error: null
                    }));
                } else {
                    throw new Error(response.message || 'Email verification failed');
                }
            } catch (error: any) {
                const errorMessage = error.message || 'Email verification failed';
                update((state: AuthState) => ({
                    ...state,
                    isLoading: false,
                    error: errorMessage
                }));
                throw error;
            }
        },

        // Reset password
        async resetPassword(email: string): Promise<void> {
            update((state: AuthState) => ({ ...state, isLoading: true, error: null }));

            try {
                const response = await AuthService.postApiPublicV1ResetPassword({ email });
                
                update((state: AuthState) => ({
                    ...state,
                    isLoading: false,
                    error: null
                }));
            } catch (error: any) {
                const errorMessage = error.message || 'Password reset failed';
                update((state: AuthState) => ({
                    ...state,
                    isLoading: false,
                    error: errorMessage
                }));
                throw error;
            }
        },

        // Clear error
        clearError(): void {
            update((state: AuthState) => ({ ...state, error: null }));
        }
    };
}

// Create and export the auth store
export const authStore = createAuthStore();

// Derived stores for easy access to authentication status
export const isAuthenticated = derived(_authStore, $auth => $auth.isAuthenticated);
export const currentUser = derived(_authStore, $auth => $auth.user);
export const authError = derived(_authStore, $auth => $auth.error);
export const authLoading = derived(_authStore, $auth => $auth.isLoading);

// Initialize auth state when the module loads (client-side only)
if (browser) {
    authStore.initialize();
}
