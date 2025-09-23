/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_AddDriverReq } from '../models/consortium_api_business_v1_AddDriverReq';
import type { consortium_api_business_v1_CancelSubscriptionReq } from '../models/consortium_api_business_v1_CancelSubscriptionReq';
import type { consortium_api_business_v1_Driver } from '../models/consortium_api_business_v1_Driver';
import type { consortium_api_business_v1_UpdateSubscriptionReq } from '../models/consortium_api_business_v1_UpdateSubscriptionReq';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class BusinessProfileService {
    /**
     * Add new driver
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1Drivers(
        requestBody?: consortium_api_business_v1_AddDriverReq,
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
             * Whether driver was added successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Created driver ID
             */
            driverId?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/drivers',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Delete driver
     * @param driverId Driver ID
     * @returns any
     * @throws ApiError
     */
    public static deleteApiPrivateV1Drivers(
        driverId: string,
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
             * Whether driver was deleted successfully
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
            url: '/api/private/v1/drivers/{driverId}',
            path: {
                'driverId': driverId,
            },
        });
    }
    /**
     * Get driver details
     * @param driverId Driver ID
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1Drivers(
        driverId: string,
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
             * Driver details
             */
            driver?: consortium_api_business_v1_Driver;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/drivers/{driverId}',
            path: {
                'driverId': driverId,
            },
        });
    }
    /**
     * Update driver
     * @param driverId Driver ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1Drivers(
        driverId: string,
        requestBody?: {
            /**
             * Hire date (YYYY-MM-DD)
             */
            hireDate: string;
            /**
             * Driver first name
             */
            firstName: string;
            /**
             * Driver last name
             */
            lastName: string;
            /**
             * Date of birth (YYYY-MM-DD)
             */
            dateOfBirth: string;
            /**
             * Phone number
             */
            phone: string;
            /**
             * Email address (optional)
             */
            email?: string;
            /**
             * Driver status
             */
            status: 'active' | 'inactive' | 'suspended';
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
             * Whether driver was updated successfully
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
            url: '/api/private/v1/drivers/{driverId}',
            path: {
                'driverId': driverId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get subscription details
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1Subscription(): CancelablePromise<{
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
             * Stripe subscription ID
             */
            subscriptionId?: string;
            /**
             * Subscription status
             */
            subscriptionStatus?: string;
            /**
             * Subscription plan type
             */
            subscriptionPlan?: string;
            /**
             * Current period start date
             */
            currentPeriodStart?: string;
            /**
             * Current period end date
             */
            currentPeriodEnd?: string;
            /**
             * Renewal month (1-12)
             */
            renewalMonth?: number;
            /**
             * Next billing date
             */
            nextBillingDate?: string;
            /**
             * Whether subscription is active
             */
            isActive?: boolean;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/subscription',
        });
    }
    /**
     * Update subscription plan
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1Subscription(
        requestBody?: consortium_api_business_v1_UpdateSubscriptionReq,
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
             * Whether subscription was updated successfully
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
            url: '/api/private/v1/subscription',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Cancel subscription
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1SubscriptionCancel(
        requestBody?: consortium_api_business_v1_CancelSubscriptionReq,
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
             * Whether subscription was cancelled successfully
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
            url: '/api/private/v1/subscription/cancel',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
