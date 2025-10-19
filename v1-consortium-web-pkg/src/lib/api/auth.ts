import { BaseService, type ServiceConfig } from './base.js';
import { TokenService, type TokenData } from './token.js';
import { create } from "@bufbuild/protobuf";
import type { Client } from "@connectrpc/connect";

// Import from the Auth namespace we created in index.ts
import type {
  UserSession,
  LoginRequest,
  LoginResponse,
  GetUserRequest,
  GetUserResponse,
  LogoutRequest,
  LogoutResponse,
  ForgotPasswordRequest,
  ForgotPasswordResponse,
  ResetPasswordRequest,
  ResetPasswordResponse,
  ChangePasswordRequest,
  ChangePasswordResponse,
  VerifyEmailRequest,
  VerifyEmailResponse,
  RefreshTokenRequest,
  RefreshTokenResponse,
  SignupRequest,
  SignupResponse,
  SocialSignupRequest,
  SocialSignupResponse,
  CompleteRegistrationRequest,
  CompleteRegistrationResponse,
  EnableMFARequest,
  EnableMFAResponse,
  VerifyMFARequest,
  VerifyMFAResponse,
  DisableMFARequest,
  DisableMFAResponse,
  UpdateUserRequest,
  UpdateUserResponse,
  CheckPermissionRequest,
  CheckPermissionResponse,
  GetUserPermissionsRequest,
  GetUserPermissionsResponse,
  GetSignupStatusRequest,
  GetSignupStatusResponse,
  ResendVerificationRequest,
  ResendVerificationResponse
} from '../gen/auth/v1/auth_pb.js';

import {
  AuthService as AuthServiceDef,
  LoginRequestSchema,
  GetUserRequestSchema,
  LogoutRequestSchema,
  ForgotPasswordRequestSchema,
  ResetPasswordRequestSchema,
  ChangePasswordRequestSchema,
  VerifyEmailRequestSchema,
  RefreshTokenRequestSchema,
  SignupRequestSchema,
  SocialSignupRequestSchema,
  CompleteRegistrationRequestSchema,
  EnableMFARequestSchema,
  VerifyMFARequestSchema,
  DisableMFARequestSchema,
  UpdateUserRequestSchema,
  CheckPermissionRequestSchema,
  GetUserPermissionsRequestSchema,
  GetSignupStatusRequestSchema,
  ResendVerificationRequestSchema
} from '../gen/auth/v1/auth_pb.js';

export interface AuthState {
  user: UserSession | null;
  token: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
}

/**
 * Authentication service that wraps the generated AuthService protobuf client
 */
export class AuthService extends BaseService {
  private client: Client<typeof AuthServiceDef>;
  private tokenService: TokenService;
  private refreshPromise: Promise<boolean> | null = null;

  constructor(config: ServiceConfig) {
    super(config);
    this.client = this.createClient(AuthServiceDef);
    this.tokenService = new TokenService();
  }

  /**
   * Authenticate user with email and password
   */
  async login(email: string, password: string, rememberMe: boolean = false): Promise<LoginResponse> {
    try {
      const request = create(LoginRequestSchema, { email, password, rememberMe });
      const response = await this.client.login(request);

      if (response.accessToken) {
        this.tokenService.saveTokens({
          accessToken: response.accessToken,
          refreshToken: response.refreshToken,
          user: response.user,
          expiresAt: response.expiresAt ? new Date(Number(response.expiresAt.seconds) * 1000) : undefined
        });
      }

      return response;
    } catch (error) {
      this.handleError(error, 'login');
    }
  }

  /**
   * Start new two-step signup process
   */
  async signup(email: string, password: string, firstName: string, lastName: string, invitationToken?: string, companyName?: string, isDotCompany?: boolean, dotNumber?: string): Promise<SignupResponse> {
    try {
      const request = create(SignupRequestSchema, {
        email,
        password,
        firstName,
        lastName,
        invitationToken: invitationToken || '',
        companyName: companyName || '',
        isDotCompany: isDotCompany || false,
        dotNumber: dotNumber || ''

      });
      const response = await this.client.signup(request);
      return response;
    } catch (error) {
      this.handleError(error, 'signup');
    }
  }

  /**
   * Start social signup process with OAuth provider
   */
  async socialSignup(
    provider: string,
    providerToken: string,
    firstName: string,
    lastName: string,
    email: string,
    invitationToken?: string
  ): Promise<SocialSignupResponse> {
    try {
      const request = create(SocialSignupRequestSchema, {
        provider,
        providerToken,
        firstName,
        lastName,
        email,
        invitationToken: invitationToken || ''
      });
      const response = await this.client.socialSignup(request);
      return response;
    } catch (error) {
      this.handleError(error, 'social signup');
    }
  }

  /**
   * Complete registration process
   */
  async completeRegistration(
    workflowId: string,
    emailVerificationToken: string,
    organizationId?: string,
    role?: string,
    subscriptionPlan?: string
  ): Promise<CompleteRegistrationResponse> {
    try {
      const request = create(CompleteRegistrationRequestSchema, {
        workflowId,
        emailVerificationToken,
        organizationId: organizationId || '',
        role: role || '',
        subscriptionPlan: subscriptionPlan || ''
      });
      const response = await this.client.completeRegistration(request);

      if (response.accessToken) {
        this.tokenService.saveTokens({
          accessToken: response.accessToken,
          refreshToken: response.refreshToken,
          user: response.user,
          expiresAt: response.expiresAt ? new Date(Number(response.expiresAt.seconds) * 1000) : undefined
        });
      }

      return response;
    } catch (error) {
      this.handleError(error, 'complete registration');
    }
  }

  /**
   * Refresh access token using refresh token
   */
  async refreshAccessToken(): Promise<boolean> {
    // Prevent multiple simultaneous refresh attempts
    if (this.refreshPromise) {
      return this.refreshPromise;
    }

    const refreshToken = this.tokenService.getRefreshToken();
    if (!refreshToken) {
      return false;
    }

    this.refreshPromise = this.performTokenRefresh(refreshToken);
    const result = await this.refreshPromise;
    this.refreshPromise = null;
    return result;
  }

  private async performTokenRefresh(refreshToken: string): Promise<boolean> {
    try {
      const request = create(RefreshTokenRequestSchema, { refreshToken });
      const response = await this.client.refreshToken(request);

      if (response.accessToken) {
        this.tokenService.saveTokens({
          accessToken: response.accessToken,
          refreshToken: response.refreshToken || refreshToken,
          user: this.tokenService.getUser(),
          expiresAt: response.expiresAt ? new Date(Number(response.expiresAt.seconds) * 1000) : undefined
        });
        return true;
      }

      return false;
    } catch (error) {
      console.error('Token refresh failed:', error);
      this.tokenService.clearAll();
      return false;
    }
  }

  /**
   * Logout user and clear tokens
   */
  async logout(): Promise<void> {
    const token = this.tokenService.getAccessToken();
    if (token) {
      try {
        const request = create(LogoutRequestSchema, { accessToken: token });
        await this.client.logout(request);
      } catch (error) {
        console.error('Logout request failed:', error);
      }
    }
    this.tokenService.clearAll();
  }

  /**
   * Get current user information
   */
  async getCurrentUser(userId?: string): Promise<UserSession | null> {
    // Return cached user if no userId specified and we have one
    if (!userId) {
      const cachedUser = this.tokenService.getUser();
      if (cachedUser) return cachedUser;
    }

    if (!this.hasValidToken()) {
      return null;
    }

    // Use stored user ID if no userId provided
    const requestUserId = userId || this.tokenService.getUser()?.userId || '';
    if (!requestUserId) {
      return null;
    }

    const operation = async () => {
      const request = create(GetUserRequestSchema, { userId: requestUserId });
      const response = await this.client.getUser(request);

      if (response.user) {
        this.tokenService.updateUser(response.user);
      }

      return response.user || null;
    };

    return this.withRetry(operation, 'get current user', () => this.refreshAccessToken());
  }

  /**
   * Update user profile
   */
  async updateUser(userId: string, firstName?: string, lastName?: string, phone?: string): Promise<UserSession | null> {
    this.requireAuth('update user');

    const operation = async () => {
      const request = create(UpdateUserRequestSchema, {
        userId,
        firstName: firstName || '',
        lastName: lastName || '',
        phone: phone || ''
      });
      const response = await this.client.updateUser(request);

      if (response.user) {
        this.tokenService.updateUser(response.user);
      }

      return response.user || null;
    };

    return this.withRetry(operation, 'update user', () => this.refreshAccessToken());
  }

  /**
   * Change user password
   */
  async changePassword(currentPassword: string, newPassword: string): Promise<void> {
    this.requireAuth('change password');

    const operation = async () => {
      const request = create(ChangePasswordRequestSchema, {
        currentPassword,
        newPassword
      });
      await this.client.changePassword(request);
    };

    return this.withRetry(operation, 'change password', () => this.refreshAccessToken());
  }

  /**
   * Request password reset
   */
  async forgotPassword(email: string): Promise<void> {
    try {
      const request = create(ForgotPasswordRequestSchema, { email });
      await this.client.forgotPassword(request);
    } catch (error) {
      this.handleError(error, 'forgot password');
    }
  }

  /**
   * Reset password with token
   */
  async resetPassword(token: string, newPassword: string): Promise<void> {
    try {
      const request = create(ResetPasswordRequestSchema, { token, newPassword });
      await this.client.resetPassword(request);
    } catch (error) {
      this.handleError(error, 'reset password');
    }
  }

  /**
   * Verify email with token
   */
  async verifyEmail(token: string): Promise<void> {
    try {
      const request = create(VerifyEmailRequestSchema, { token });
      await this.client.verifyEmail(request);
    } catch (error) {
      this.handleError(error, 'verify email');
    }
  }

  /**
   * Enable MFA for user
   */
  async enableMFA(userId: string, method: string): Promise<EnableMFAResponse> {
    this.requireAuth('enable MFA');

    const operation = async () => {
      const request = create(EnableMFARequestSchema, { userId, method });
      return await this.client.enableMFA(request);
    };

    return this.withRetry(operation, 'enable MFA', () => this.refreshAccessToken());
  }

  /**
   * Verify MFA code
   */
  async verifyMFA(userId: string, code: string, method: string): Promise<VerifyMFAResponse> {
    this.requireAuth('verify MFA');

    const operation = async () => {
      const request = create(VerifyMFARequestSchema, { userId, code, method });
      return await this.client.verifyMFA(request);
    };

    return this.withRetry(operation, 'verify MFA', () => this.refreshAccessToken());
  }

  /**
   * Disable MFA for user
   */
  async disableMFA(userId: string, password: string): Promise<void> {
    this.requireAuth('disable MFA');

    const operation = async () => {
      const request = create(DisableMFARequestSchema, { userId, password });
      await this.client.disableMFA(request);
    };

    return this.withRetry(operation, 'disable MFA', () => this.refreshAccessToken());
  }

  /**
   * Check user permission
   */
  async checkPermission(userId: string, permission: string, resourceId?: string): Promise<boolean> {
    this.requireAuth('check permission');

    const operation = async () => {
      const request = create(CheckPermissionRequestSchema, {
        userId,
        permission,
        resourceId: resourceId || ''
      });
      const response = await this.client.checkPermission(request);
      return response.allowed;
    };

    return this.withRetry(operation, 'check permission', () => this.refreshAccessToken());
  }

  /**
   * Get user permissions
   */
  async getUserPermissions(userId: string): Promise<{ permissions: string[]; role: string }> {
    this.requireAuth('get user permissions');

    const operation = async () => {
      const request = create(GetUserPermissionsRequestSchema, { userId });
      const response = await this.client.getUserPermissions(request);
      return {
        permissions: response.permissions,
        role: response.role
      };
    };

    return this.withRetry(operation, 'get user permissions', () => this.refreshAccessToken());
  }

  /**
   * Get signup status by workflow ID
   */
  async getSignupStatus(workflowId: string): Promise<GetSignupStatusResponse> {
    try {
      const request = create(GetSignupStatusRequestSchema, { workflowId });
      const response = await this.client.getSignupStatus(request);
      return response;
    } catch (error) {
      this.handleError(error, 'get signup status');
    }
  }

  /**
   * Resend verification email
   */
  async resendVerification(workflowId: string): Promise<ResendVerificationResponse> {
    try {
      const request = create(ResendVerificationRequestSchema, { workflowId });
      const response = await this.client.resendVerification(request);
      return response;
    } catch (error) {
      this.handleError(error, 'resend verification');
    }
  }

  /**
   * Get current authentication state
   */
  getAuthState(): AuthState {
    const tokens = this.tokenService.getTokens();
    return {
      user: tokens.user,
      token: tokens.accessToken,
      refreshToken: tokens.refreshToken,
      isAuthenticated: this.tokenService.isAuthenticated(),
    };
  }

  /**
   * Get current access token
   */
  getCurrentAccessToken(): string | null {
    return this.tokenService.getAccessToken();
  }

  /**
   * Get current user from storage
   */
  getUser(): UserSession | null {
    return this.tokenService.getUser();
  }

  /**
   * Check if current token is expired
   */
  isTokenExpired(): boolean {
    const token = this.tokenService.getAccessToken();
    return !token || this.tokenService.isTokenExpired(token);
  }
}