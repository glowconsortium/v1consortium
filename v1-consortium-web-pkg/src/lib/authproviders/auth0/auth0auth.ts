import { createAuth0Client, type Auth0Client, type User } from '@auth0/auth0-spa-js';

interface Auth0Config {
  domain: string;
  clientId: string;
  redirectUri?: string;
  audience?: string;
  scope?: string;
}

interface LoginOptions {
  appState?: any;
  fragment?: string;
  redirectUri?: string;
}

interface LogoutOptions {
  returnTo?: string;
  federated?: boolean;
}

interface TokenOptions {
  audience?: string;
  scope?: string;
  ignoreCache?: boolean;
}

class Auth0Store {
  private client: Auth0Client | null = $state(null);
  public isAuthenticated = $state(false);
  public isLoading = $state(true);
  public user: User | null = $state(null);
  public error: string | null = $state(null);

  async init(config: Auth0Config): Promise<void> {
    try {
      this.isLoading = true;
      this.error = null;

      this.client = await createAuth0Client({
        domain: config.domain,
        clientId: config.clientId,
        authorizationParams: {
          redirect_uri: config.redirectUri || window.location.origin,
          audience: config.audience,
          scope: config.scope || 'openid profile email'
        }
      });

      // Handle redirect callback
      await this.handleRedirectCallback();

      // Update authentication state
      await this.updateAuthState();
    } catch (err) {
      this.handleError('Auth0 initialization failed', err);
    } finally {
      this.isLoading = false;
    }
  }

  private async handleRedirectCallback(): Promise<void> {
    const urlParams = new URLSearchParams(window.location.search);
    if (urlParams.has('code') && urlParams.has('state')) {
      try {
        await this.client?.handleRedirectCallback();
        // Clean up URL
        const url = new URL(window.location.href);
        url.search = '';
        window.history.replaceState({}, document.title, url.toString());
      } catch (err) {
        this.handleError('Failed to handle redirect callback', err);
      }
    }
  }

  private async updateAuthState(): Promise<void> {
    if (!this.client) return;

    try {
      this.isAuthenticated = await this.client.isAuthenticated();
      
      if (this.isAuthenticated) {
        const user = await this.client.getUser();
        this.user = user ?? null;
      } else {
        this.user = null;
      }
    } catch (err) {
      this.handleError('Failed to update authentication state', err);
    }
  }

  async login(options: LoginOptions = {}): Promise<void> {
    if (!this.client) {
      this.handleError('Auth0 client not initialized', new Error('Client not initialized'));
      return;
    }

    try {
      this.error = null;
      await this.client.loginWithRedirect({
        appState: options.appState,
        fragment: options.fragment,
        authorizationParams: {
          redirect_uri: options.redirectUri
        }
      });
    } catch (err) {
      this.handleError('Login failed', err);
    }
  }

  async logout(options: LogoutOptions = {}): Promise<void> {
    if (!this.client) {
      this.handleError('Auth0 client not initialized', new Error('Client not initialized'));
      return;
    }

    try {
      this.error = null;
      await this.client.logout({
        logoutParams: {
          returnTo: options.returnTo || window.location.origin,
          federated: options.federated
        }
      });
      
      // Clear local state
      this.isAuthenticated = false;
      this.user = null;
    } catch (err) {
      this.handleError('Logout failed', err);
    }
  }

  async getAccessToken(options: TokenOptions = {}): Promise<string> {
    if (!this.client) {
      throw new Error('Auth0 client not initialized');
    }

    if (!this.isAuthenticated) {
      throw new Error('User not authenticated');
    }

    try {
      this.error = null;
      return await this.client.getTokenSilently({
        authorizationParams: {
          audience: options.audience,
          scope: options.scope
        },
        cacheMode: options.ignoreCache ? 'off' : 'on'
      });
    } catch (err) {
      this.handleError('Failed to get access token', err);
      throw err;
    }
  }

  async getIdToken(): Promise<string | undefined> {
    if (!this.client) {
      throw new Error('Auth0 client not initialized');
    }

    if (!this.isAuthenticated) {
      throw new Error('User not authenticated');
    }

    try {
      const claims = await this.client.getIdTokenClaims();
      return claims?.__raw;
    } catch (err) {
      this.handleError('Failed to get ID token', err);
      throw err;
    }
  }

  async refreshToken(): Promise<void> {
    if (!this.client) {
      throw new Error('Auth0 client not initialized');
    }

    try {
      this.error = null;
      await this.client.getTokenSilently({ cacheMode: 'off' });
      await this.updateAuthState();
    } catch (err) {
      this.handleError('Failed to refresh token', err);
      throw err;
    }
  }

  private handleError(message: string, error: any): void {
    const errorMessage = error instanceof Error ? error.message : String(error);
    this.error = `${message}: ${errorMessage}`;
    console.error(message, error);
  }

  // Utility methods
  hasRole(role: string): boolean {
    if (!this.user) return false;
    const roles = this.user['https://example.com/roles'] || this.user.roles || [];
    return Array.isArray(roles) && roles.includes(role);
  }

  hasPermission(permission: string): boolean {
    if (!this.user) return false;
    const permissions = this.user['https://example.com/permissions'] || this.user.permissions || [];
    return Array.isArray(permissions) && permissions.includes(permission);
  }

  clearError(): void {
    this.error = null;
  }
}

export const auth0Store = new Auth0Store();
export type { Auth0Config, LoginOptions, LogoutOptions, TokenOptions };