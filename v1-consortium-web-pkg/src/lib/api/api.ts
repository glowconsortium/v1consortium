import { createConnectTransport } from "@connectrpc/connect-web";
import type { Transport } from "@connectrpc/connect";
import { AuthService } from './auth.js';
import { TokenService } from './token.js';
import type { ServiceConfig } from './base.js';
import type { 
  UserSession, 
  LoginResponse, 
  SignupResponse,
  SocialSignupResponse,
  CompleteRegistrationResponse
} from '../gen/auth/v1/auth_pb.js';

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

export interface ApiClientState {
  user: UserSession | null;
  token: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
}

/**
 * Main API client that coordinates all services
 */
export class ApiClient {
  private authService: AuthService | null = null;
  private tokenService: TokenService;
  private baseUrl: string = '';

  constructor() {
    this.tokenService = new TokenService();
  }

  /**
   * Initialize the API client with transport and base URL
   */
  initialize(baseUrl: string): void {
    this.baseUrl = baseUrl;
    
    const transport = createApiTransport(baseUrl, () => this.tokenService.getAccessToken());
    
    const serviceConfig: ServiceConfig = {
      transport,
      baseUrl,
      getToken: () => this.tokenService.getAccessToken(),
      setToken: (token: string | null) => {
        if (token) {
          this.tokenService.updateAccessToken(token);
        } else {
          this.tokenService.clearAccessToken();
        }
      },
      onAuthError: () => {
        this.tokenService.clearAll();
      }
    };

    this.authService = new AuthService(serviceConfig);
  }

  /**
   * Get the auth service instance
   */
  get auth(): AuthService {
    if (!this.authService) {
      throw new Error('ApiClient not initialized. Call initialize() first.');
    }
    return this.authService;
  }

  // Convenience methods that delegate to auth service
  async login(email: string, password: string, rememberMe: boolean = false): Promise<LoginResponse> {
    return this.auth.login(email, password, rememberMe);
  }

  async logout(): Promise<void> {
    return this.auth.logout();
  }

  async getCurrentUser(userId?: string): Promise<UserSession | null> {
    return this.auth.getCurrentUser(userId);
  }

  async signup(
    email: string, 
    password: string, 
    firstName: string, 
    lastName: string,
    invitationToken?: string,
    companyName?: string,
    isDotCompany?: boolean,
    dotNumber?: string
  ): Promise<SignupResponse> {
    return this.auth.signup(email, password, firstName, lastName, invitationToken, companyName, isDotCompany, dotNumber);
  }

  async socialSignup(
    provider: string,
    providerToken: string,
    firstName: string,
    lastName: string,
    email: string,
    invitationToken?: string
  ): Promise<SocialSignupResponse> {
    return this.auth.socialSignup(provider, providerToken, firstName, lastName, email, invitationToken);
  }

  async completeRegistration(
    workflowId: string,
    emailVerificationToken: string,
    organizationId?: string,
    role?: string,
    subscriptionPlan?: string
  ): Promise<CompleteRegistrationResponse> {
    return this.auth.completeRegistration(workflowId, emailVerificationToken, organizationId, role, subscriptionPlan);
  }

  async refreshAccessToken(): Promise<boolean> {
    return this.auth.refreshAccessToken();
  }

  getAuthState(): ApiClientState {
    const tokens = this.tokenService.getTokens();
    return {
      user: tokens.user,
      token: tokens.accessToken,
      refreshToken: tokens.refreshToken,
      isAuthenticated: this.tokenService.isAuthenticated(),
    };
  }

  getToken(): string | null {
    return this.tokenService.getAccessToken();
  }

  getUser(): UserSession | null {
    return this.tokenService.getUser();
  }

  isTokenExpired(): boolean {
    const token = this.tokenService.getAccessToken();
    return !token || this.tokenService.isTokenExpired(token);
  }
}

export const apiClient = new ApiClient();