/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_profile_v1_ActivityLogEntry } from '../models/consortium_api_profile_v1_ActivityLogEntry';
import type { consortium_api_profile_v1_ApiKey } from '../models/consortium_api_profile_v1_ApiKey';
import type { consortium_api_profile_v1_ChangePasswordReq } from '../models/consortium_api_profile_v1_ChangePasswordReq';
import type { consortium_api_profile_v1_CreateApiKeyReq } from '../models/consortium_api_profile_v1_CreateApiKeyReq';
import type { consortium_api_profile_v1_ExportDataReq } from '../models/consortium_api_profile_v1_ExportDataReq';
import type { consortium_api_profile_v1_UpdateNotificationPrefsReq } from '../models/consortium_api_profile_v1_UpdateNotificationPrefsReq';
import type { consortium_api_profile_v1_UpdateProfileReq } from '../models/consortium_api_profile_v1_UpdateProfileReq';
import type { consortium_api_profile_v1_UploadAvatarReq } from '../models/consortium_api_profile_v1_UploadAvatarReq';
import type { consortium_api_profile_v1_UserProfile } from '../models/consortium_api_profile_v1_UserProfile';
import type { consortium_api_profile_v1_UserSession } from '../models/consortium_api_profile_v1_UserSession';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class ProfileService {
    /**
     * Get current logged-in user profile
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1Profile(): CancelablePromise<{
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
             * Whether request was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * User profile data
             */
            profile?: consortium_api_profile_v1_UserProfile;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/profile',
        });
    }
    /**
     * Update current user profile
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1Profile(
        requestBody?: consortium_api_profile_v1_UpdateProfileReq,
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
             * Whether profile update was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/private/v1/profile',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get user activity log
     * @param page Page number (default: 1)
     * @param pageSize Page size (default: 20)
     * @param action Filter by action type
     * @param dateFrom Filter activities from date (YYYY-MM-DD)
     * @param dateTo Filter activities to date (YYYY-MM-DD)
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1ProfileActivity(
        page?: number,
        pageSize?: number,
        action?: string,
        dateFrom?: string,
        dateTo?: string,
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
             * Whether request was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * List of activity log entries
             */
            activities?: Array<consortium_api_profile_v1_ActivityLogEntry>;
            /**
             * Total number of activities
             */
            totalCount?: number;
            /**
             * Current page
             */
            page?: number;
            /**
             * Page size
             */
            pageSize?: number;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/profile/activity',
            query: {
                'page': page,
                'pageSize': pageSize,
                'action': action,
                'dateFrom': dateFrom,
                'dateTo': dateTo,
            },
        });
    }
    /**
     * Get user API keys
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1ProfileApiKeys(): CancelablePromise<{
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
             * Whether request was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * List of API keys
             */
            apiKeys?: Array<consortium_api_profile_v1_ApiKey>;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/profile/api-keys',
        });
    }
    /**
     * Create new API key
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1ProfileApiKeys(
        requestBody?: consortium_api_profile_v1_CreateApiKeyReq,
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
             * Whether API key was created successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Created API key ID
             */
            keyId?: string;
            /**
             * Generated API key (shown only once)
             */
            apiKey?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/profile/api-keys',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete API key
     * @param keyId API key ID
     * @returns any
     * @throws ApiError
     */
    public static deleteApiPrivateV1ProfileApiKeys(
        keyId: string,
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
             * Whether API key was deleted successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/private/v1/profile/api-keys/{keyId}',
            path: {
                'keyId': keyId,
            },
        });
    }
    /**
     * Update API key
     * @param keyId API key ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1ProfileApiKeys(
        keyId: string,
        requestBody?: {
            /**
             * Whether key is active
             */
            isActive?: boolean;
            /**
             * API key name
             */
            name: string;
            /**
             * API key scopes
             */
            scopes: Array<string>;
        },
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
             * Whether API key was updated successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/private/v1/profile/api-keys/{keyId}',
            path: {
                'keyId': keyId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete user avatar
     * @returns any
     * @throws ApiError
     */
    public static deleteApiPrivateV1ProfileAvatar(): CancelablePromise<{
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
             * Whether avatar deletion was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/private/v1/profile/avatar',
        });
    }
    /**
     * Upload user avatar
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1ProfileAvatar(
        requestBody?: consortium_api_profile_v1_UploadAvatarReq,
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
             * Whether avatar upload was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Avatar image URL
             */
            avatarUrl?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/profile/avatar',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete user account
     * @param password Current password for confirmation
     * @param reason Reason for account deletion
     * @returns any
     * @throws ApiError
     */
    public static deleteApiPrivateV1ProfileDelete(
        password: string,
        reason?: string,
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
             * Whether account deletion was initiated successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/private/v1/profile/delete',
            query: {
                'password': password,
                'reason': reason,
            },
        });
    }
    /**
     * Export user data
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1ProfileExport(
        requestBody?: consortium_api_profile_v1_ExportDataReq,
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
             * Whether data export was initiated successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Export request ID
             */
            exportId?: string;
            /**
             * Download URL (when ready)
             */
            downloadUrl?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/profile/export',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Update notification preferences
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1ProfileNotifications(
        requestBody?: consortium_api_profile_v1_UpdateNotificationPrefsReq,
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
             * Whether notification preferences were updated
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/private/v1/profile/notifications',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Change user password
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1ProfilePassword(
        requestBody?: consortium_api_profile_v1_ChangePasswordReq,
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
             * Whether password change was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/private/v1/profile/password',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Revoke all user sessions except current
     * @returns any
     * @throws ApiError
     */
    public static deleteApiPrivateV1ProfileSessions(): CancelablePromise<{
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
             * Whether all sessions were revoked successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Number of sessions revoked
             */
            revokedCount?: number;
        };
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/private/v1/profile/sessions',
        });
    }
    /**
     * Get active user sessions
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1ProfileSessions(): CancelablePromise<{
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
             * Whether request was successful
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * List of active sessions
             */
            sessions?: Array<consortium_api_profile_v1_UserSession>;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/profile/sessions',
        });
    }
    /**
     * Revoke user session
     * @param sessionId Session ID to revoke
     * @returns any
     * @throws ApiError
     */
    public static deleteApiPrivateV1ProfileSessions1(
        sessionId: string,
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
             * Whether session was revoked successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/private/v1/profile/sessions/{sessionId}',
            path: {
                'sessionId': sessionId,
            },
        });
    }
}
