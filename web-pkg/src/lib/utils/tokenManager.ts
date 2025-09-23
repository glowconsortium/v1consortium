import { browser } from '$app/environment';

export interface TokenData {
    accessToken?: string;
    refreshToken?: string;
    expiresAt?: number;
    tokenType?: string;
}

class TokenManager {
    private readonly ACCESS_TOKEN_KEY = 'formapp_access_token';
    private readonly REFRESH_TOKEN_KEY = 'formapp_refresh_token';
    private readonly EXPIRES_AT_KEY = 'formapp_token_expires_at';
    private readonly TOKEN_TYPE_KEY = 'formapp_token_type';

    // Store tokens
    setTokens(tokens: TokenData): void {
        if (!browser) return;

        try {
            if (tokens.accessToken) {
                localStorage.setItem(this.ACCESS_TOKEN_KEY, tokens.accessToken);
            }
            if (tokens.refreshToken) {
                localStorage.setItem(this.REFRESH_TOKEN_KEY, tokens.refreshToken);
            }
            if (tokens.expiresAt) {
                localStorage.setItem(this.EXPIRES_AT_KEY, tokens.expiresAt.toString());
            }
            if (tokens.tokenType) {
                localStorage.setItem(this.TOKEN_TYPE_KEY, tokens.tokenType);
            }
        } catch (error) {
            console.error('Failed to store tokens:', error);
        }
    }

    // Get stored tokens
    getTokens(): TokenData {
        if (!browser) return {};

        try {
            const accessToken = localStorage.getItem(this.ACCESS_TOKEN_KEY);
            const refreshToken = localStorage.getItem(this.REFRESH_TOKEN_KEY);
            const expiresAt = localStorage.getItem(this.EXPIRES_AT_KEY);
            const tokenType = localStorage.getItem(this.TOKEN_TYPE_KEY);

            return {
                accessToken: accessToken || undefined,
                refreshToken: refreshToken || undefined,
                expiresAt: expiresAt ? parseInt(expiresAt, 10) : undefined,
                tokenType: tokenType || undefined,
            };
        } catch (error) {
            console.error('Failed to retrieve tokens:', error);
            return {};
        }
    }

    // Get access token
    getAccessToken(): string | null {
        if (!browser) return null;
        return localStorage.getItem(this.ACCESS_TOKEN_KEY);
    }

    // Get refresh token
    getRefreshToken(): string | null {
        if (!browser) return null;
        return localStorage.getItem(this.REFRESH_TOKEN_KEY);
    }

    // Check if access token is expired
    isAccessTokenExpired(): boolean {
        if (!browser) return true;

        const accessToken = this.getAccessToken();
        if (!accessToken) return true;

        const expiresAt = localStorage.getItem(this.EXPIRES_AT_KEY);
        if (!expiresAt) {
            // If no expiration time stored, assume 1 hour from now
            return false;
        }

        const expirationTime = parseInt(expiresAt, 10);
        const currentTime = Date.now();
        
        // Add 5 minutes buffer to prevent edge cases
        return currentTime >= (expirationTime - 5 * 60 * 1000);
    }

    // Check if we have valid tokens
    hasValidTokens(): boolean {
        const tokens = this.getTokens();
        return !!(tokens.accessToken && !this.isAccessTokenExpired());
    }

    // Check if we have a refresh token
    hasRefreshToken(): boolean {
        return !!this.getRefreshToken();
    }

    // Check if token is expired (alias for isAccessTokenExpired)
    isTokenExpired(): boolean {
        return this.isAccessTokenExpired();
    }

    // Clear all tokens
    clearTokens(): void {
        if (!browser) return;

        try {
            localStorage.removeItem(this.ACCESS_TOKEN_KEY);
            localStorage.removeItem(this.REFRESH_TOKEN_KEY);
            localStorage.removeItem(this.EXPIRES_AT_KEY);
            localStorage.removeItem(this.TOKEN_TYPE_KEY);
        } catch (error) {
            console.error('Failed to clear tokens:', error);
        }
    }

    // Get authorization header
    getAuthHeader(): Record<string, string> {
        const accessToken = this.getAccessToken();
        const tokenType = localStorage.getItem(this.TOKEN_TYPE_KEY) || 'Bearer';
        
        if (accessToken) {
            return {
                'Authorization': `${tokenType} ${accessToken}`
            };
        }
        
        return {};
    }
}

export const tokenManager = new TokenManager();