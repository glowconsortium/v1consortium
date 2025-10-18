export interface TokenStorage {
  accessToken: string | null;
  refreshToken: string | null;
  user: any | null;
}

export interface TokenData {
  accessToken: string;
  refreshToken?: string;
  user?: any;
  expiresAt?: Date;
}

/**
 * Service for managing authentication tokens and user session storage
 */
export class TokenService {
  private static readonly ACCESS_TOKEN_KEY = 'auth_token';
  private static readonly REFRESH_TOKEN_KEY = 'refresh_token';
  private static readonly USER_SESSION_KEY = 'user_session';

  /**
   * Get current tokens from storage
   */
  getTokens(): TokenStorage {
    if (typeof window === 'undefined') {
      return { accessToken: null, refreshToken: null, user: null };
    }

    const accessToken = localStorage.getItem(TokenService.ACCESS_TOKEN_KEY);
    const refreshToken = localStorage.getItem(TokenService.REFRESH_TOKEN_KEY);
    const userJson = localStorage.getItem(TokenService.USER_SESSION_KEY);

    let user = null;
    if (userJson) {
      try {
        user = JSON.parse(userJson);
      } catch (error) {
        console.error('Failed to parse stored user session:', error);
        this.clearUserSession();
      }
    }

    return { accessToken, refreshToken, user };
  }

  /**
   * Get only the access token
   */
  getAccessToken(): string | null {
    if (typeof window === 'undefined') return null;
    return localStorage.getItem(TokenService.ACCESS_TOKEN_KEY);
  }

  /**
   * Get only the refresh token
   */
  getRefreshToken(): string | null {
    if (typeof window === 'undefined') return null;
    return localStorage.getItem(TokenService.REFRESH_TOKEN_KEY);
  }

  /**
   * Get stored user session
   */
  getUser(): any | null {
    if (typeof window === 'undefined') return null;
    
    const userJson = localStorage.getItem(TokenService.USER_SESSION_KEY);
    if (!userJson) return null;

    try {
      return JSON.parse(userJson);
    } catch (error) {
      console.error('Failed to parse stored user session:', error);
      this.clearUserSession();
      return null;
    }
  }

  /**
   * Store authentication tokens and user data
   */
  saveTokens(data: TokenData): void {
    if (typeof window === 'undefined') return;

    localStorage.setItem(TokenService.ACCESS_TOKEN_KEY, data.accessToken);
    
    if (data.refreshToken) {
      localStorage.setItem(TokenService.REFRESH_TOKEN_KEY, data.refreshToken);
    }

    if (data.user) {
      try {
        // Use a custom replacer to handle BigInt values
        const userJson = JSON.stringify(data.user, (key, value) =>
          typeof value === 'bigint' ? value.toString() : value
        );
        localStorage.setItem(TokenService.USER_SESSION_KEY, userJson);
      } catch (error) {
        console.error('Failed to save user session to storage:', error);
      }
    }
  }

  /**
   * Update only the access token (useful for token refresh)
   */
  updateAccessToken(accessToken: string): void {
    if (typeof window === 'undefined') return;
    localStorage.setItem(TokenService.ACCESS_TOKEN_KEY, accessToken);
  }

  /**
   * Update user session data
   */
  updateUser(user: any): void {
    if (typeof window === 'undefined') return;

    try {
      const userJson = JSON.stringify(user, (key, value) =>
        typeof value === 'bigint' ? value.toString() : value
      );
      localStorage.setItem(TokenService.USER_SESSION_KEY, userJson);
    } catch (error) {
      console.error('Failed to save user session to storage:', error);
    }
  }

  /**
   * Clear all stored authentication data
   */
  clearAll(): void {
    if (typeof window === 'undefined') return;

    localStorage.removeItem(TokenService.ACCESS_TOKEN_KEY);
    localStorage.removeItem(TokenService.REFRESH_TOKEN_KEY);
    localStorage.removeItem(TokenService.USER_SESSION_KEY);
  }

  /**
   * Clear only the access token
   */
  clearAccessToken(): void {
    if (typeof window === 'undefined') return;
    localStorage.removeItem(TokenService.ACCESS_TOKEN_KEY);
  }

  /**
   * Clear only the user session
   */
  clearUserSession(): void {
    if (typeof window === 'undefined') return;
    localStorage.removeItem(TokenService.USER_SESSION_KEY);
  }

  /**
   * Check if we have valid authentication
   */
  isAuthenticated(): boolean {
    const accessToken = this.getAccessToken();
    return !!accessToken && !this.isTokenExpired(accessToken);
  }

  /**
   * Check if a token is expired
   */
  isTokenExpired(token: string): boolean {
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      let exp: number;
      if (typeof payload.exp === 'bigint') {
        exp = Number(payload.exp) * 1000;
      } else {
        exp = payload.exp * 1000;
      }
      return Date.now() >= exp;
    } catch (error) {
      console.error('Error checking token expiration:', error);
      return true;
    }
  }

  /**
   * Get token expiration time
   */
  getTokenExpiration(token: string): Date | null {
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      let exp: number;
      if (typeof payload.exp === 'bigint') {
        exp = Number(payload.exp) * 1000;
      } else {
        exp = payload.exp * 1000;
      }
      return new Date(exp);
    } catch (error) {
      console.error('Error getting token expiration:', error);
      return null;
    }
  }

  /**
   * Check if token expires soon (within 5 minutes)
   */
  isTokenExpiringSoon(token: string): boolean {
    const expiration = this.getTokenExpiration(token);
    if (!expiration) return true;

    const fiveMinutesFromNow = new Date(Date.now() + 5 * 60 * 1000);
    return expiration <= fiveMinutesFromNow;
  }
}