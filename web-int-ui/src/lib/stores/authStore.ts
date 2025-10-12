import { writable, derived, type Writable, type Readable } from 'svelte/store';
import { browser } from '$app/environment';
import type { Transport } from '@connectrpc/connect';
import { createClient } from '@connectrpc/connect';
import {
  AuthService,
  LoginRequestSchema,
  LogoutRequestSchema,
  RefreshTokenRequestSchema,
  GetUserRequestSchema,
  CheckPermissionRequestSchema,
  GetUserPermissionsRequestSchema,
  type UserSession
} from '$lib/gen/auth/v1/auth_pb';
import { create } from '@bufbuild/protobuf';
import type { Timestamp } from '@bufbuild/protobuf/wkt';

// Type alias for User using the generated UserSession
export type User = UserSession;

// Helper function to convert Timestamp to Date
function timestampToDate(timestamp: Timestamp): Date {
  const milliseconds = Number(timestamp.seconds) * 1000 + Math.floor(timestamp.nanos / 1000000);
  return new Date(milliseconds);
}

// Auth state interface
export interface AuthState {
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

// Auth store interface
export interface AuthStore extends Readable<AuthState> {
    login: (email: string, password: string, rememberMe?: boolean) => Promise<void>;
    logout: () => Promise<void>;
    refreshToken: () => Promise<void>;
    getCurrentUser: () => Promise<User | null>;
    getIdToken: () => string | null;
    hasRole: (role: string) => boolean;
    hasPermission: (permission: string, resourceId?: string) => Promise<boolean>;
    clearError: () => void;
}

// Storage keys
const STORAGE_KEYS = {
    ACCESS_TOKEN: 'auth_access_token',
    REFRESH_TOKEN: 'auth_refresh_token',
    USER: 'auth_user',
    EXPIRES_AT: 'auth_expires_at'
} as const;

// Utility functions for token storage
const TokenStorage = {
    setTokens: (accessToken: string, refreshToken: string, expiresAt?: Date) => {
        if (!browser) return;
        
        localStorage.setItem(STORAGE_KEYS.ACCESS_TOKEN, accessToken);
        localStorage.setItem(STORAGE_KEYS.REFRESH_TOKEN, refreshToken);
        if (expiresAt) {
            localStorage.setItem(STORAGE_KEYS.EXPIRES_AT, expiresAt.toISOString());
        }
    },
    
    getAccessToken: (): string | null => {
        if (!browser) return null;
        return localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN);
    },
    
    getRefreshToken: (): string | null => {
        if (!browser) return null;
        return localStorage.getItem(STORAGE_KEYS.REFRESH_TOKEN);
    },
    
    getExpiresAt: (): Date | null => {
        if (!browser) return null;
        const expiresAt = localStorage.getItem(STORAGE_KEYS.EXPIRES_AT);
        return expiresAt ? new Date(expiresAt) : null;
    },
    
    setUser: (user: User) => {
        if (!browser) return;
        localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(user));
    },
    
    getUser: (): User | null => {
        if (!browser) return null;
        const userData = localStorage.getItem(STORAGE_KEYS.USER);
        return userData ? JSON.parse(userData) : null;
    },
    
    clearAll: () => {
        if (!browser) return;
        Object.values(STORAGE_KEYS).forEach(key => {
            localStorage.removeItem(key);
        });
    },
    
    isTokenExpired: (): boolean => {
        const expiresAt = TokenStorage.getExpiresAt();
        if (!expiresAt) return true;
        return new Date() >= expiresAt;
    }
};

// Create auth store factory function
export function createAuthStore(transport: Transport, args: Record<string, any> = {}): AuthStore {
    const { subscribe, set, update } = writable<AuthState>(initialState);
    
    // Create Connect RPC client
    const client = createClient(AuthService, transport);
    
    // Initialize store with persisted data
    const initializeStore = () => {
        if (!browser) return;
        
        const savedUser = TokenStorage.getUser();
        const accessToken = TokenStorage.getAccessToken();
        
        if (savedUser && accessToken && !TokenStorage.isTokenExpired()) {
            update(state => ({
                ...state,
                user: savedUser,
                isAuthenticated: true,
                isInitialized: true
            }));
        } else {
            // Clear expired tokens
            TokenStorage.clearAll();
            update(state => ({
                ...state,
                isInitialized: true
            }));
        }
    };
    
    // Login method
    const login = async (email: string, password: string, rememberMe = false): Promise<void> => {
        update(state => ({ ...state, isLoading: true, error: null }));
        
        try {
            const request = create(LoginRequestSchema, {
                email,
                password,
                rememberMe
            });
            
            const response = await client.login(request);
            
            if (response.user && response.accessToken && response.refreshToken) {
                // Store tokens
                const expiresAt = response.expiresAt ? timestampToDate(response.expiresAt) : undefined;
                TokenStorage.setTokens(response.accessToken, response.refreshToken, expiresAt);
                TokenStorage.setUser(response.user);
                
                update(state => ({
                    ...state,
                    user: response.user!,
                    isAuthenticated: true,
                    isLoading: false,
                    error: null
                }));
            } else {
                throw new Error('Invalid login response');
            }
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Login failed';
            update(state => ({
                ...state,
                isLoading: false,
                error: errorMessage
            }));
            throw error;
        }
    };
    
    // Logout method
    const logout = async (): Promise<void> => {
        update(state => ({ ...state, isLoading: true }));
        
        try {
            const accessToken = TokenStorage.getAccessToken();
            if (accessToken) {
                const request = create(LogoutRequestSchema, {
                    accessToken
                });
                
                await client.logout(request);
            }
        } catch (error) {
            console.warn('Logout request failed:', error);
        } finally {
            // Clear local storage regardless of API call success
            TokenStorage.clearAll();
            update(state => ({
                ...initialState,
                isInitialized: true
            }));
        }
    };
    
    // Refresh token method
    const refreshToken = async (): Promise<void> => {
        const refreshTokenValue = TokenStorage.getRefreshToken();
        if (!refreshTokenValue) {
            throw new Error('No refresh token available');
        }
        
        update(state => ({ ...state, isLoading: true, error: null }));
        
        try {
            const request = create(RefreshTokenRequestSchema, {
                refreshToken: refreshTokenValue
            });
            
            const response = await client.refreshToken(request);
            
            if (response.accessToken && response.refreshToken) {
                const expiresAt = response.expiresAt ? timestampToDate(response.expiresAt) : undefined;
                TokenStorage.setTokens(response.accessToken, response.refreshToken, expiresAt);
                
                update(state => ({
                    ...state,
                    isLoading: false,
                    error: null
                }));
            } else {
                throw new Error('Invalid refresh token response');
            }
        } catch (error) {
            // Clear tokens on refresh failure
            TokenStorage.clearAll();
            const errorMessage = error instanceof Error ? error.message : 'Token refresh failed';
            update(state => ({
                ...initialState,
                isInitialized: true,
                error: errorMessage
            }));
            throw error;
        }
    };
    
    // Get current user method
    const getCurrentUser = async (): Promise<User | null> => {
        const savedUser = TokenStorage.getUser();
        if (!savedUser) return null;
        
        try {
            const request = create(GetUserRequestSchema, {
                userId: savedUser.userId
            });
            
            const response = await client.getUser(request);
            
            if (response.user) {
                TokenStorage.setUser(response.user);
                update(state => ({
                    ...state,
                    user: response.user!
                }));
                return response.user;
            }
            
            return null;
        } catch (error) {
            console.warn('Failed to fetch current user:', error);
            return savedUser; // Return cached user on error
        }
    };
    
    // Get ID token method
    const getIdToken = (): string | null => {
        return TokenStorage.getAccessToken();
    };
    
    // Has role method
    const hasRole = (role: string): boolean => {
        const user = TokenStorage.getUser();
        return user?.role === role;
    };
    
    // Has permission method
    const hasPermission = async (permission: string, resourceId?: string): Promise<boolean> => {
        const user = TokenStorage.getUser();
        if (!user) return false;
        
        // Check if permission exists in user's permissions array
        if (user.permissions?.includes(permission)) {
            return true;
        }
        
        // If resource-specific, check with API
        if (resourceId) {
            try {
                const request = create(CheckPermissionRequestSchema, {
                    userId: user.userId,
                    permission,
                    resourceId
                });
                
                const response = await client.checkPermission(request);
                return response.allowed;
            } catch (error) {
                console.warn('Permission check failed:', error);
                return false;
            }
        }
        
        return false;
    };
    
    // Clear error method
    const clearError = (): void => {
        update(state => ({ ...state, error: null }));
    };
    
    // Initialize store when created
    initializeStore();
    
    return {
        subscribe,
        login,
        logout,
        refreshToken,
        getCurrentUser,
        getIdToken,
        hasRole,
        hasPermission,
        clearError
    };
}

// Global auth store instance
let _authStore: AuthStore | null = null;

// Function to set the global auth store instance
export function setAuthStore(store: AuthStore): void {
    _authStore = store;
}

// Create auth store factory with default state store for derived stores
const createDefaultStore = () => writable<AuthState>(initialState);
const defaultStore = createDefaultStore();

// Create and export the auth store factory
// Export the main createAuthStore function as authStore for backward compatibility
export { createAuthStore as authStore };

// Derived stores for easy access to authentication status
export const isAuthenticated = derived(
    _authStore || defaultStore, 
    $auth => $auth.isAuthenticated
);

export const currentUser = derived(
    _authStore || defaultStore, 
    $auth => $auth.user
);

export const authError = derived(
    _authStore || defaultStore, 
    $auth => $auth.error
);

export const authLoading = derived(
    _authStore || defaultStore, 
    $auth => $auth.isLoading
);

// Example usage:
/* 
import { createAuthStore } from '$lib/stores/authStore';
import { createConnectTransport } from '@connectrpc/connect-web';

const transport = createConnectTransport({
    baseUrl: 'http://localhost:8000',
    useBinaryFormat: true,
});

const authStore = createAuthStore(transport);

// Use the store
await authStore.login('user@example.com', 'password');
console.log(authStore.hasRole('admin'));
const hasPermission = await authStore.hasPermission('read', 'resource-id');
*/
