import { apiClient } from './client.js';
import { OpenAPI, AuthService as GeneratedAuthService } from '../backendapi/generated-backendapi/index.js';
import { tokenManager } from '$lib/utils/tokenManager.js';
import type {
	consortium_api_auth_v1_LoginReq,
	consortium_api_auth_v1_LoginRes,
	consortium_api_auth_v1_SignupReq,
	consortium_api_auth_v1_SignupRes,
	consortium_api_auth_v1_LogoutReq,
	consortium_api_auth_v1_RefreshTokenReq,
	consortium_api_auth_v1_VerifyEmailReq,
	consortium_api_auth_v1_ResetPasswordReq
} from '../backendapi/generated-backendapi/index.js';
import { AuthError, handleApiResponse } from './util.js';

export class AuthService {
	
	async login(credentials: consortium_api_auth_v1_LoginReq) {
		const response = await handleApiResponse(GeneratedAuthService.postApiPublicV1Login(credentials));

		// Store tokens in tokenManager for persistence
		if (response?.data?.accessToken) {
			// Store tokens for persistence
			tokenManager.setTokens({
				accessToken: response.data.accessToken,
				refreshToken: response.data.refreshToken,
				// Add expiration if available, otherwise default to 1 hour
				expiresAt: Date.now() + (60 * 60 * 1000),
				tokenType: 'Bearer'
			});

			// Update OpenAPI configuration with new token
			OpenAPI.TOKEN = response.data.accessToken;
		}

		return response;
	}

	async register(userData: consortium_api_auth_v1_SignupReq) {
		console.log('AuthService: Calling register API with:', userData);
		const response = await handleApiResponse(GeneratedAuthService.postApiPublicV1Signup(userData));
		console.log('AuthService: Register response:', response);

		// If register returns access_token, store it
		if (response?.data?.accessToken) {
			tokenManager.setTokens({
				accessToken: response.data.accessToken,
				refreshToken: response.data.refreshToken,
				expiresAt: Date.now() + (60 * 60 * 1000),
				tokenType: 'Bearer'
			});

			OpenAPI.TOKEN = response.data.accessToken;
		}

		return response;
	}

	async logout(request?: consortium_api_auth_v1_LogoutReq) {
		try {
			const response = await handleApiResponse(GeneratedAuthService.postApiPublicV1Logout(request));
			return response;
		} finally {
			// Always clear tokens locally regardless of API response
			tokenManager.clearTokens();
			OpenAPI.TOKEN = undefined;
		}
	}

	async refreshToken(request: consortium_api_auth_v1_RefreshTokenReq) {
		const refreshToken = tokenManager.getRefreshToken();

		if (!refreshToken) {
			throw new Error('No refresh token available');
		}

		try {
			const response = await handleApiResponse(GeneratedAuthService.postApiPublicV1Refresh(request));

			const data = response.data;

			// Store new tokens
			if (data?.accessToken) {
				tokenManager.setTokens({
					accessToken: data.accessToken,
					refreshToken: data.refreshToken || refreshToken,
					expiresAt: Date.now() + (60 * 60 * 1000),
					tokenType: 'Bearer'
				});

				// Update OpenAPI configuration with new token
				OpenAPI.TOKEN = data.accessToken;
			}

			return { success: true, data };
		} catch (error) {
			// Clear invalid tokens
			tokenManager.clearTokens();
			OpenAPI.TOKEN = undefined;
			throw error;
		}
	}

	async resetPassword(request: consortium_api_auth_v1_ResetPasswordReq) {
		return await handleApiResponse(GeneratedAuthService.postApiPublicV1ResetPassword(request));
	}

	async verifyEmail(request: consortium_api_auth_v1_VerifyEmailReq) {
		return await handleApiResponse(GeneratedAuthService.postApiPublicV1VerifyEmail(request));
	}

	// Helper method to check if user is authenticated
	isAuthenticated(): boolean {
		const tokens = tokenManager.getTokens();
		return !!(tokens?.accessToken && !tokenManager.isTokenExpired());
	}

	// Helper method to get current tokens
	getTokens() {
		return tokenManager.getTokens();
	}

	// Helper method to clear authentication
	clearAuth() {
		tokenManager.clearTokens();
		OpenAPI.TOKEN = undefined;
	}
}


// Create a singleton instance
export const authService = new AuthService();

export { AuthError };
