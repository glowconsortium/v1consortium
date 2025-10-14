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
  isAuthenticated: boolean;
}

export class ApiClient {
  private token: string | null = null;
  private transport: Transport | null = null;
  private client: Client<typeof AuthService> | null = null;

  constructor() {
    this.loadTokenFromStorage();
  }

  initialize(transport: Transport): void {
    this.transport = transport;
    this.client = createClient(AuthService, this.transport);
  }

  private loadTokenFromStorage(): void {
    if (typeof window !== 'undefined') {
      this.token = localStorage.getItem('auth_token');
    }
  }

  private saveTokenToStorage(token: string): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem('auth_token', token);
    }
    this.token = token;
  }

  private removeTokenFromStorage(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('auth_token');
    }
    this.token = null;
  }

  async login(email: string, password: string): Promise<LoginResponse> {
    if (!this.client) {
      throw new Error('ApiClient not initialized. Call initialize() first.');
    }

    try {
      const request = create(LoginRequestSchema, { email, password, rememberMe: false });
      const response = await this.client.login(request);

      if (response.accessToken) {
        this.saveTokenToStorage(response.accessToken);
      }

      return response;
    } catch (error) {
      throw new Error(`Login failed: ${error}`);
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
    this.removeTokenFromStorage();
  }

  async getCurrentUser(userId: string): Promise<UserSession | null> {
    if (!this.token || !this.client) {
      return null;
    }

    try {
      // For getting current user, we'll use "me" as the user ID
      const request = create(GetUserRequestSchema, { userId: userId});
      const response = await this.client.getUser(request);
      return response.user || null;
    } catch (error) {
      console.error('Error getting current user:', error);
      if (error instanceof Error && error.message.includes('401')) {
        this.removeTokenFromStorage();
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
      user: null, // Will be populated by the auth store
      token: this.token,
      isAuthenticated: !!this.token,
    };
  }

  getToken(): string | null {
    return this.token;
  }
}

export const apiClient = new ApiClient();