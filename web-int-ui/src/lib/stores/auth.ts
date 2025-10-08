import { writable, derived } from 'svelte/store';
import { browser } from '$app/environment';
import { auth0Store, type Auth0Config, type LoginOptions, type LogoutOptions } from '../authproviders/auth0/auth0auth.js';
import type { User } from '@auth0/auth0-spa-js';

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
    let stateUpdateInterval: NodeJS.Timeout | null = null;

    return {
        subscribe,
        
        // Initialize auth state on app start
        async initialize(config: Auth0Config) {
            if (!browser) return;

            update((state: AuthState) => ({ ...state, isLoading: true }));

            try {
                // Initialize Auth0
                await auth0Store.init(config);
                
                // Sync state with auth0Store
                this.syncWithAuth0Store();
                
                // Set up periodic state sync to catch auth0Store changes
                if (stateUpdateInterval) {
                    clearInterval(stateUpdateInterval);
                }
                stateUpdateInterval = setInterval(() => {
                    this.syncWithAuth0Store();
                }, 1000);
                
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

        // Sync local state with auth0Store state
        syncWithAuth0Store() {
            update((state: AuthState) => ({
                ...state,
                user: auth0Store.user,
                isAuthenticated: auth0Store.isAuthenticated,
                isLoading: auth0Store.isLoading,
                error: auth0Store.error
            }));
        },

        // Login user (Auth0 uses redirect-based login)
        async login(options: LoginOptions = {}): Promise<void> {
            update((state: AuthState) => ({ ...state, isLoading: true, error: null }));

            try {
                // Ensure Auth0 is initialized before attempting login
                if (!auth0Store.isInitialized) {
                    throw new Error('Auth0 client not initialized. Please call initialize() first.');
                }

                await auth0Store.login({
                    ...options,
                    appState: options.appState ?? { 
                        returnTo: window.location.pathname + window.location.search 
                    },
                    redirectUri: options.redirectUri || (window.location.origin + '/auth/callback')
                });
                // Note: After successful login, user will be redirected to Auth0
                // and then back to the callback URL
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
            }
        },

        // Logout user
        async logout(options: LogoutOptions = {}): Promise<void> {
            update((state: AuthState) => ({ ...state, isLoading: true }));

            try {
                await auth0Store.logout(options);
                // Auth0 will handle the logout and redirect
                // Local state will be cleared by the auth0Store
                this.syncWithAuth0Store();
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

        // Refresh session
        async refreshSession(): Promise<boolean> {
            try {
                await auth0Store.refreshToken();
                this.syncWithAuth0Store();
                return true;
            } catch (error) {
                console.warn('Token refresh failed:', error);
                await this.logout();
                return false;
            }
        },

        // Get access token
        async getAccessToken(): Promise<string | null> {
            try {
                return await auth0Store.getAccessToken();
            } catch (error) {
                console.warn('Failed to get access token:', error);
                return null;
            }
        },

        // Get ID token
        async getIdToken(): Promise<string | null> {
            try {
                return await auth0Store.getIdToken() || null;
            } catch (error) {
                console.warn('Failed to get ID token:', error);
                return null;
            }
        },

        // Get current user (Auth0 manages this internally)
        async getCurrentUser(): Promise<User | null> {
            this.syncWithAuth0Store();
            return auth0Store.user;
        },

        // Expose appState captured during redirect
        getAppState(): any | null {
            return auth0Store.getAppState();
        },

        // Clear error
        clearError(): void {
            auth0Store.clearError();
            update((state: AuthState) => ({ ...state, error: null }));
        },

        // Check if user has role
        hasRole(role: string): boolean {
            return auth0Store.hasRole(role);
        },

        // Check if user has permission
        hasPermission(permission: string): boolean {
            return auth0Store.hasPermission(permission);
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

// Note: Initialize auth state manually by calling authStore.initialize(config) 
// where config is your Auth0 configuration object
