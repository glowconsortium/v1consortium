/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_auth_v1_LoginReq } from '../models/consortium_api_auth_v1_LoginReq';
import type { consortium_api_auth_v1_LogoutReq } from '../models/consortium_api_auth_v1_LogoutReq';
import type { consortium_api_auth_v1_RefreshTokenReq } from '../models/consortium_api_auth_v1_RefreshTokenReq';
import type { consortium_api_auth_v1_ResetPasswordReq } from '../models/consortium_api_auth_v1_ResetPasswordReq';
import type { consortium_api_auth_v1_SignupReq } from '../models/consortium_api_auth_v1_SignupReq';
import type { consortium_api_auth_v1_UpdatePasswordReq } from '../models/consortium_api_auth_v1_UpdatePasswordReq';
import type { consortium_api_auth_v1_VerifyEmailReq } from '../models/consortium_api_auth_v1_VerifyEmailReq';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class AuthService {
    /**
     * User login
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1Login(
        requestBody?: consortium_api_auth_v1_LoginReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether login was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * User ID
             */
            userId?: string;
            /**
             * Associated carrier ID
             */
            carrierId?: string;
            /**
             * User role
             */
            role?: string;
            /**
             * JWT access token
             */
            accessToken?: string;
            /**
             * JWT refresh token
             */
            refreshToken?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * User logout
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1Logout(
        requestBody?: consortium_api_auth_v1_LogoutReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether logout was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/logout',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Refresh access token
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1Refresh(
        requestBody?: consortium_api_auth_v1_RefreshTokenReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether token refresh was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * New JWT access token
             */
            accessToken?: string;
            /**
             * New JWT refresh token
             */
            refreshToken?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/refresh',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Request password reset
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1ResetPassword(
        requestBody?: consortium_api_auth_v1_ResetPasswordReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether password reset email was sent
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/reset-password',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Signup with carrier creation and subscription
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1Signup(
        requestBody?: consortium_api_auth_v1_SignupReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether signup was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Created user ID
             */
            userId?: string;
            /**
             * Created carrier ID
             */
            carrierId?: string;
            /**
             * JWT access token
             */
            accessToken?: string;
            /**
             * JWT refresh token
             */
            refreshToken?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/signup',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Update password with reset token
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1UpdatePassword(
        requestBody?: consortium_api_auth_v1_UpdatePasswordReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether password update was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/update-password',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Verify user email
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPublicV1VerifyEmail(
        requestBody?: consortium_api_auth_v1_VerifyEmailReq,
    ): CancelablePromise<{
        /**
         * Error code
         */
        code?: number;
        /**
         * Error message
         */
        message?: string;
        /**
         * Result data for certain request according API definition
         */
        data?: {
            /**
             * Whether email verification was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/public/v1/verify-email',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
