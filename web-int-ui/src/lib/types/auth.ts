// Authentication related types that extend the generated API types
import type {
	consortium_api_auth_v1_LoginReq,
	consortium_api_auth_v1_LoginRes,
	consortium_api_auth_v1_SignupReq,
	consortium_api_auth_v1_SignupRes,
	consortium_api_auth_v1_ResetPasswordReq,
	consortium_api_auth_v1_RefreshTokenReq,
	consortium_api_auth_v1_VerifyEmailReq,
	consortium_api_profile_v1_UserProfile,
	consortium_api_profile_v1_UpdateProfileReq,
	consortium_api_profile_v1_ChangePasswordReq
} from '../backendapi/generated-backendapi/index.js';

// Extended User type that combines auth user and profile user
export interface User extends Partial<consortium_api_profile_v1_UserProfile> {
	id: string;
	email: string;
	name?: string;
	role?: string;
	carrierId?: string;
}

// Re-export generated types with cleaner names
export type LoginRequest = consortium_api_auth_v1_LoginReq;
export type RegisterRequest = consortium_api_auth_v1_SignupReq;
export type LoginResponse = consortium_api_auth_v1_LoginRes;
export type RegisterResponse = consortium_api_auth_v1_SignupRes;
//export type ForgotPasswordRequest = consortium_api_auth_v1_ForgotPasswordReq;
export type ResetPasswordRequest = consortium_api_auth_v1_ResetPasswordReq;
export type RefreshTokenRequest = consortium_api_auth_v1_RefreshTokenReq;
export type VerifyEmailRequest = consortium_api_auth_v1_VerifyEmailReq;
//export type ResendVerificationRequest = consortium_api_auth_v1_ResendVerificationReq;
export type UserProfile = consortium_api_profile_v1_UserProfile;
export type UpdateProfileRequest = consortium_api_profile_v1_UpdateProfileReq;
export type ChangePasswordRequest = consortium_api_profile_v1_ChangePasswordReq;
//export type ChangeEmailRequest = consortium_api_profile_v1_ChangeEmailReq;

// Legacy types for backward compatibility
export interface AuthResponse {
	user: User;
	access_token: string;
	refresh_token?: string;
	expires_in?: number;
}

export interface PasswordResetRequest {
	email: string;
}

export interface PasswordResetConfirm {
	token: string;
	password: string;
}
