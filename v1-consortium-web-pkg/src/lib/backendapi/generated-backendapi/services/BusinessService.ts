/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_CreateOrderReq } from '../models/consortium_api_business_v1_CreateOrderReq';
import type { consortium_api_business_v1_CreateRandomPoolReq } from '../models/consortium_api_business_v1_CreateRandomPoolReq';
import type { consortium_api_business_v1_Order } from '../models/consortium_api_business_v1_Order';
import type { consortium_api_business_v1_Product } from '../models/consortium_api_business_v1_Product';
import type { consortium_api_business_v1_RandomPool } from '../models/consortium_api_business_v1_RandomPool';
import type { consortium_api_business_v1_RandomSelection } from '../models/consortium_api_business_v1_RandomSelection';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class BusinessService {
    /**
     * Get orders for carrier
     * @param page Page number (default: 1)
     * @param pageSize Page size (default: 20)
     * @param status Filter by order status
     * @param type Filter by product type
     * @param driverId Filter by driver ID
     * @param dateFrom Filter orders from date (YYYY-MM-DD)
     * @param dateTo Filter orders to date (YYYY-MM-DD)
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1CarrierOrders(
        page?: number,
        pageSize?: number,
        status?: 'pending' | 'processing' | 'completed' | 'failed' | 'cancelled',
        type?: 'drug' | 'physical' | 'mvr' | 'random',
        driverId?: string,
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
             * List of orders
             */
            orders?: Array<consortium_api_business_v1_Order>;
            /**
             * Total number of orders
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
            url: '/api/private/v1/carrier/orders',
            query: {
                'page': page,
                'pageSize': pageSize,
                'status': status,
                'type': type,
                'driverId': driverId,
                'dateFrom': dateFrom,
                'dateTo': dateTo,
            },
        });
    }
    /**
     * Create new compliance order
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1Orders(
        requestBody?: consortium_api_business_v1_CreateOrderReq,
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
             * Whether order was created successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Created order ID
             */
            orderId?: string;
            /**
             * Order total amount
             */
            total?: number;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/orders',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get order details
     * @param orderId Order ID
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1Orders(
        orderId: string,
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
             * Order details
             */
            order?: consortium_api_business_v1_Order;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/orders/{orderId}',
            path: {
                'orderId': orderId,
            },
        });
    }
    /**
     * Update order
     * @param orderId Order ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1Orders(
        orderId: string,
        requestBody?: {
            /**
             * Order priority
             */
            priority?: 'normal' | 'urgent' | 'rush';
            /**
             * Order notes
             */
            notes?: string;
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
             * Whether order was updated successfully
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
            url: '/api/private/v1/orders/{orderId}',
            path: {
                'orderId': orderId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Cancel order
     * @param orderId Order ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1OrdersCancel(
        orderId: string,
        requestBody?: {
            /**
             * Cancellation reason
             */
            reason?: string;
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
             * Whether order was cancelled successfully
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
            url: '/api/private/v1/orders/{orderId}/cancel',
            path: {
                'orderId': orderId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get order results
     * @param orderId Order ID
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1OrdersResults(
        orderId: string,
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
             * Results download URL
             */
            resultsUrl?: string;
            /**
             * Results status
             */
            status?: string;
            /**
             * Last update timestamp
             */
            updatedAt?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/orders/{orderId}/results',
            path: {
                'orderId': orderId,
            },
        });
    }
    /**
     * Retry failed order
     * @param orderId Order ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1OrdersRetry(
        orderId: string,
        requestBody?: any,
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
             * Whether order retry was initiated successfully
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
            url: '/api/private/v1/orders/{orderId}/retry',
            path: {
                'orderId': orderId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get available compliance products
     * @param type Filter by product type
     * @param isActive Filter by active status
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1Products(
        type?: 'drug' | 'physical' | 'mvr' | 'random',
        isActive?: boolean,
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
             * List of products
             */
            products?: Array<consortium_api_business_v1_Product>;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/products',
            query: {
                'type': type,
                'isActive': isActive,
            },
        });
    }
    /**
     * Get random testing pools
     * @param isActive Filter by active status
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1RandomPools(
        isActive?: boolean,
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
             * List of random pools
             */
            pools?: Array<consortium_api_business_v1_RandomPool>;
        };
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/private/v1/random-pools',
            query: {
                'isActive': isActive,
            },
        });
    }
    /**
     * Create random testing pool
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1RandomPools(
        requestBody?: consortium_api_business_v1_CreateRandomPoolReq,
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
             * Whether pool was created successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Created pool ID
             */
            poolId?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/random-pools',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Update random testing pool
     * @param poolId Pool ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static putApiPrivateV1RandomPools(
        poolId: string,
        requestBody?: {
            /**
             * Pool name
             */
            name: string;
            /**
             * Testing frequency
             */
            frequency: 'monthly' | 'quarterly' | 'annually';
            /**
             * Percentage of drivers to select
             */
            percentage: number;
            /**
             * Driver IDs to include in pool
             */
            driverIds: Array<string>;
            /**
             * Whether pool is active
             */
            isActive?: boolean;
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
             * Whether pool was updated successfully
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
            url: '/api/private/v1/random-pools/{poolId}',
            path: {
                'poolId': poolId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Trigger manual random selection
     * @param poolId Pool ID
     * @param requestBody
     * @returns any
     * @throws ApiError
     */
    public static postApiPrivateV1RandomPoolsSelect(
        poolId: string,
        requestBody?: {
            /**
             * Reason for manual selection
             */
            reason?: string;
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
             * Whether selection was triggered successfully
             */
            success?: boolean;
            /**
             * Response message
             */
            message?: string;
            /**
             * Created selection ID
             */
            selectionId?: string;
            /**
             * Temporal workflow ID
             */
            workflowId?: string;
        };
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/private/v1/random-pools/{poolId}/select',
            path: {
                'poolId': poolId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get random testing selections
     * @param page Page number (default: 1)
     * @param pageSize Page size (default: 20)
     * @param poolId Filter by pool ID
     * @param dateFrom Filter selections from date (YYYY-MM-DD)
     * @param dateTo Filter selections to date (YYYY-MM-DD)
     * @returns any
     * @throws ApiError
     */
    public static getApiPrivateV1RandomSelections(
        page?: number,
        pageSize?: number,
        poolId?: string,
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
             * List of random selections
             */
            selections?: Array<consortium_api_business_v1_RandomSelection>;
            /**
             * Total number of selections
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
            url: '/api/private/v1/random-selections',
            query: {
                'page': page,
                'pageSize': pageSize,
                'poolId': poolId,
                'dateFrom': dateFrom,
                'dateTo': dateTo,
            },
        });
    }
}
