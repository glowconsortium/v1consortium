import { writable, derived } from 'svelte/store';
import { browser } from '$app/environment';
import { apiClient } from '../api/api.js';
import type { UserSession, LoginResponse } from '../gen/auth/v1/auth_pb.js';

interface AuthState {
    user: UserSession | null;
    isAuthenticated: boolean;
    isLoading: boolean;
    error: string | null;
    isInitialized: boolean;
}

interface LoginCredentials {
    email: string;
    password: string;
    rememberMe?: boolean;
}

interface ApiConfig {
    baseUrl: string;
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
    let checkUserInterval: NodeJS.Timeout | null = null;
    let tokenRefreshInterval: NodeJS.Timeout | null = null;

    // Sync local state with API client state
    const syncWithApiClient = async () => {
        try {
            const apiState = apiClient.getAuthState();
            
            // Get user from cache first, then try to fetch if needed
            let currentUser = apiState.user;
            if (!currentUser && apiState.isAuthenticated) {
                currentUser = await apiClient.getCurrentUser();
            }
            
            update((state: AuthState) => ({
                ...state,
                user: currentUser,
                isAuthenticated: apiState.isAuthenticated,
                error: null
            }));
        } catch (error) {
            console.error('Error syncing with API client:', error);
            update((state: AuthState) => ({
                ...state,
                user: null,
                isAuthenticated: false,
                error: error instanceof Error ? error.message : 'Sync failed'
            }));
        }
    };

    // Set up automatic token refresh
    const setupTokenRefresh = () => {
        if (tokenRefreshInterval) {
            clearInterval(tokenRefreshInterval);
        }

        tokenRefreshInterval = setInterval(async () => {
            const apiState = apiClient.getAuthState();
            if (apiState.isAuthenticated && apiClient.isTokenExpired()) {
                try {
                    const refreshed = await apiClient.refreshAccessToken();
                    if (refreshed) {
                        await syncWithApiClient();
                    } else {
                        await store.logout();
                    }
                } catch (error) {
                    console.error('Automatic token refresh failed:', error);
                    await store.logout();
                }
            }
        }, 60000); // Check every minute
    };

    const store = {
        subscribe,
        
        // Initialize auth state on app start
        async initialize(config: ApiConfig) {
            if (!browser) return;

            update((state: AuthState) => ({ ...state, isLoading: true }));

            try {
                // Initialize API client
                apiClient.initialize(config.baseUrl);
                
                // Sync initial state
                await syncWithApiClient();
                
                // Set up periodic user state check
                if (checkUserInterval) {
                    clearInterval(checkUserInterval);
                }
                checkUserInterval = setInterval(() => {
                    syncWithApiClient();
                }, 30000); // Check every 30 seconds

                // Set up automatic token refresh
                setupTokenRefresh();
                
            } catch (error) {
                console.error('Failed to initialize auth state:', error);
                update((state: AuthState) => ({ 
                    ...state, 
                    error: error instanceof Error ? error.message : 'Initialization failed'
                }));
            } finally {
                update((state: AuthState) => ({ 
                    ...state, 
                    isLoading: false, 
                    isInitialized: true 
                }));
            }
        },

        // Login user
        async login(credentials: LoginCredentials): Promise<LoginResponse> {
            update((state: AuthState) => ({ ...state, isLoading: true, error: null }));

            try {
                const response = await apiClient.login(credentials.email, credentials.password);
                
                // Sync state after successful login
                await syncWithApiClient();
                
                // Restart token refresh monitoring
                setupTokenRefresh();
                
                return response;
            } catch (error: any) {
                const errorMessage = error.message || 'Login failed';
                console.error('Login error:', error);
                update((state: AuthState) => ({
                    ...state,
                    isLoading: false,
                    error: errorMessage,
                    isAuthenticated: false,
                    user: null
                }));
                throw error;
            } finally {
                update((state: AuthState) => ({ ...state, isLoading: false }));
            }
        },

        // Logout user
        async logout(): Promise<void> {
            update((state: AuthState) => ({ ...state, isLoading: true }));

            try {
                await apiClient.logout();
                
                // Clear intervals
                if (tokenRefreshInterval) {
                    clearInterval(tokenRefreshInterval);
                    tokenRefreshInterval = null;
                }
                
                // Clear local state
                update((state: AuthState) => ({
                    ...state,
                    user: null,
                    isAuthenticated: false,
                    isLoading: false,
                    error: null
                }));
            } catch (error) {
                console.warn('Logout failed:', error);
                // Even if logout fails, clear local state
                update((state: AuthState) => ({
                    ...state,
                    user: null,
                    isAuthenticated: false,
                    isLoading: false,
                    error: null
                }));
            }
        },

        // Refresh session - API client handles tokens internally
        async refreshSession(userId: string): Promise<boolean> {
            try {
                // Try to get current user to verify token is still valid
                const user = await apiClient.getCurrentUser(userId);
                if (user) {
                    await syncWithApiClient();
                    return true;
                } else {
                    await store.logout();
                    return false;
                }
            } catch (error) {
                console.warn('Token refresh failed:', error);
                await store.logout();
                return false;
            }
        },

        // Get access token
        async getAccessToken(): Promise<string | null> {
            return apiClient.getToken();
        },

        // Get current user
        async getCurrentUser(): Promise<UserSession | null> {
            return await apiClient.getCurrentUser();
        },

        // Clear error
        clearError(): void {
            update((state: AuthState) => ({ ...state, error: null }));
        },

        // Clean up interval when needed
        destroy(): void {
            if (checkUserInterval) {
                clearInterval(checkUserInterval);
                checkUserInterval = null;
            }
            if (tokenRefreshInterval) {
                clearInterval(tokenRefreshInterval);
                tokenRefreshInterval = null;
            }
        }
    };

    return store;
}

// Create and export the auth store
export const authStore = createAuthStore();

// Derived stores for easy access to authentication status
export const isAuthenticated = derived(_authStore, $auth => $auth.isAuthenticated);
export const currentUser = derived(_authStore, $auth => $auth.user);
export const authError = derived(_authStore, $auth => $auth.error);
export const authLoading = derived(_authStore, $auth => $auth.isLoading);

// Note: Initialize auth state manually by calling authStore.initialize(config) 
// where config is your Auth0 configuration object
