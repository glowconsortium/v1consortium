import { authStore } from '$lib/stores/auth.js';
import { tokenManager } from '$lib/utils/tokenManager.js';
import { browser } from '$app/environment';

// Create a fetch wrapper that automatically handles token refresh
export async function authenticatedFetch(url: string, options: RequestInit = {}): Promise<Response> {
    if (!browser) {
        return fetch(url, options);
    }

    // Add auth headers if we have tokens
    const authHeaders = tokenManager.getAuthHeader();
    const headers = {
        ...options.headers,
        ...authHeaders,
    };

    // First attempt
    let response = await fetch(url, {
        ...options,
        headers,
        credentials: 'include',
    });

    // If unauthorized and we have a refresh token, try to refresh
    if (response.status === 401 && tokenManager.hasRefreshToken()) {
        console.log('Interceptor: Attempting token refresh due to 401');
        
        try {
            const refreshed = await authStore.refreshSession();
            
            if (refreshed) {
                // Retry the original request with new tokens
                const newAuthHeaders = tokenManager.getAuthHeader();
                response = await fetch(url, {
                    ...options,
                    headers: {
                        ...options.headers,
                        ...newAuthHeaders,
                    },
                    credentials: 'include',
                });
            }
        } catch (error) {
            console.error('Interceptor: Token refresh failed:', error);
        }
    }

    return response;
}