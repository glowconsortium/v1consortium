import type { 
  UserSession,
  LoginRequest,
  LoginResponse,
  GetUserRequest,
  GetUserResponse,
  RegisterRequest,
  RegisterResponse,
  LogoutRequest
} from '../gen/auth/v1/auth_pb.js';
import { 
  AuthService,
  LoginRequestSchema,
  GetUserRequestSchema,
  RegisterRequestSchema,
  LogoutRequestSchema
} from '../gen/auth/v1/auth_pb.js';
import { createClient, type Transport, type Client } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { create } from "@bufbuild/protobuf";

export function createApiTransport(baseUrl: string, getToken: () => string | null): Transport {
  return createConnectTransport({
    baseUrl,
    // Include cookies in cross-origin requests:
    fetch: (input, init) => fetch(input, {...init, credentials: "include"}),
    interceptors: [
      (next) => async (req) => {
        const token = getToken();
        if (token) {
          req.header.set("Authorization", `Bearer ${token}`);
        }
        return next(req);
      }, 
    ],
  });
}

export interface AuthState {
  user: UserSession | null;
  token: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
}

export class ApiClient {
  private token: string | null = null;
  private refreshToken: string | null = null;
  private user: UserSession | null = null;
  private transport: Transport | null = null;
  private client: Client<typeof AuthService> | null = null;
  private refreshPromise: Promise<boolean> | null = null;
  private baseUrl: string = '';

  constructor() {
    this.loadFromStorage();
  }

  initialize(transport: Transport, baseUrl: string): void {
    this.transport = transport;
    this.baseUrl = baseUrl;
    this.client = createClient(AuthService, this.transport);
  }

  private loadFromStorage(): void {
    if (typeof window !== 'undefined') {
      this.token = localStorage.getItem('auth_token');
      this.refreshToken = localStorage.getItem('refresh_token');
      const userJson = localStorage.getItem('user_session');
      if (userJson) {
        try {
          this.user = JSON.parse(userJson);
        } catch (error) {
          console.error('Failed to parse stored user session:', error);
          this.user = null;
        }
      }
    }
  }

  private saveToStorage(token: string, refreshToken?: string, user?: UserSession): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem('auth_token', token);
      if (refreshToken) {
        localStorage.setItem('refresh_token', refreshToken);
      }
      if (user) {
        try {
          // Use a custom replacer to handle BigInt values
          const userJson = JSON.stringify(user, (key, value) =>
            typeof value === 'bigint' ? value.toString() : value
          );
          localStorage.setItem('user_session', userJson);
          this.user = user;
        } catch (error) {
          console.error('Failed to save user session to storage:', error);
        }
      }
    }
    this.token = token;
    if (refreshToken) {
      this.refreshToken = refreshToken;
    }
  }

  private removeFromStorage(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('auth_token');
      localStorage.removeItem('refresh_token');
      localStorage.removeItem('user_session');
    }
    this.token = null;
    this.refreshToken = null;
    this.user = null;
  }

  async login(email: string, password: string): Promise<LoginResponse> {
    if (!this.client) {
      throw new Error('ApiClient not initialized. Call initialize() first.');
    }

    try {
      const request = create(LoginRequestSchema, { email, password, rememberMe: false });
      const response = await this.client.login(request);

      if (response.accessToken) {
        this.saveToStorage(
          response.accessToken, 
          response.refreshToken, 
          response.user || undefined
        );
      }

      return response;
    } catch (error) {
      throw new Error(`Login failed: ${error}`);
    }
  }

  async refreshAccessToken(): Promise<boolean> {
    // Prevent multiple simultaneous refresh attempts
    if (this.refreshPromise) {
      return this.refreshPromise;
    }

    if (!this.refreshToken || !this.client) {
      return false;
    }

    this.refreshPromise = this.performTokenRefresh();
    const result = await this.refreshPromise;
    this.refreshPromise = null;
    return result;
  }

  private async performTokenRefresh(): Promise<boolean> {
    try {
      // Note: You'll need to add a refresh endpoint to your AuthService
      // This is a placeholder - adjust based on your actual refresh endpoint
      const response = await fetch(`${this.baseUrl}/auth/refresh`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ refreshToken: this.refreshToken }),
        credentials: 'include'
      });

      if (!response.ok) {
        throw new Error('Token refresh failed');
      }

      const data = await response.json();
      
      if (data.accessToken) {
        this.saveToStorage(
          data.accessToken, 
          data.refreshToken || this.refreshToken,
          data.user || this.user
        );
        return true;
      }
      
      return false;
    } catch (error) {
      console.error('Token refresh failed:', error);
      this.removeFromStorage();
      return false;
    }
  }

  async logout(): Promise<void> {
    if (this.client && this.token) {
      try {
        const request = create(LogoutRequestSchema, { accessToken: this.token });
        await this.client.logout(request);
      } catch (error) {
        console.error('Logout request failed:', error);
      }
    }
    this.removeFromStorage();
  }

  async getCurrentUser(userId?: string): Promise<UserSession | null> {
    // Return cached user if no userId specified and we have one
    if (!userId && this.user) {
      return this.user;
    }

    if (!this.token || !this.client) {
      return null;
    }

    // Use stored user ID if no userId provided
    const requestUserId = userId || this.user?.userId || '';
    if (!requestUserId) {
      return null;
    }

    try {
      const request = create(GetUserRequestSchema, { userId: requestUserId });
      const response = await this.client.getUser(request);
      
      if (response.user) {
        // Update cached user
        this.user = response.user;
        if (typeof window !== 'undefined') {
          try {
            const userJson = JSON.stringify(response.user, (key, value) =>
              typeof value === 'bigint' ? value.toString() : value
            );
            localStorage.setItem('user_session', userJson);
          } catch (error) {
            console.error('Failed to save user session to storage:', error);
          }
        }
      }
      
      return response.user || null;
    } catch (error) {
      console.error('Error getting current user:', error);
      if (error instanceof Error && error.message.includes('401')) {
        // Try to refresh token before giving up
        const refreshed = await this.refreshAccessToken();
        if (refreshed) {
          // Retry the request with new token
          try {
            const request = create(GetUserRequestSchema, { userId: requestUserId });
            const response = await this.client.getUser(request);
            
            if (response.user) {
              this.user = response.user;
              if (typeof window !== 'undefined') {
                try {
                  const userJson = JSON.stringify(response.user, (key, value) =>
                    typeof value === 'bigint' ? value.toString() : value
                  );
                  localStorage.setItem('user_session', userJson);
                } catch (error) {
                  console.error('Failed to save user session to storage:', error);
                }
              }
            }
            
            return response.user || null;
          } catch (retryError) {
            console.error('Retry failed after token refresh:', retryError);
            this.removeFromStorage();
            return null;
          }
        } else {
          this.removeFromStorage();
        }
      }
      return null;
    }
  }

  async register(
    email: string, 
    password: string, 
    firstName: string, 
    lastName: string,
    organizationId?: string,
    role?: string,
    invitationToken?: string
  ): Promise<RegisterResponse> {
    if (!this.client) {
      throw new Error('ApiClient not initialized. Call initialize() first.');
    }

    try {
      const request = create(RegisterRequestSchema, { 
        email, 
        password, 
        firstName, 
        lastName,
        organizationId: organizationId || '',
        role: role || '',
        invitationToken: invitationToken || ''
      });
      const response = await this.client.registerUser(request);
      return response;
    } catch (error) {
      throw new Error(`Registration failed: ${error}`);
    }
  }

  getAuthState(): AuthState {
    return {
      user: this.user,
      token: this.token,
      refreshToken: this.refreshToken,
      isAuthenticated: !!this.token,
    };
  }

  getToken(): string | null {
    return this.token;
  }

  getUser(): UserSession | null {
    return this.user;
  }

  isTokenExpired(): boolean {
    if (!this.token) return true;
    
    try {
      const payload = JSON.parse(atob(this.token.split('.')[1]));
      // Handle both regular numbers and BigInt values
      let exp: number;
      if (typeof payload.exp === 'bigint') {
        exp = Number(payload.exp) * 1000; // Convert BigInt to number, then to milliseconds
      } else {
        exp = payload.exp * 1000; // Convert to milliseconds
      }
      return Date.now() >= exp;
    } catch (error) {
      console.error('Error checking token expiration:', error);
      return true;
    }
  }
}

export const apiClient = new ApiClient();