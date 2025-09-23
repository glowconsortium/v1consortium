/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_GetOrdersReq = {
    /**
     * Page number (default: 1)
     */
    page?: number;
    /**
     * Page size (default: 20)
     */
    pageSize?: number;
    /**
     * Filter by order status
     */
    status?: consortium_api_business_v1_GetOrdersReq.status;
    /**
     * Filter by product type
     */
    type?: consortium_api_business_v1_GetOrdersReq.type;
    /**
     * Filter by driver ID
     */
    driverId?: string;
    /**
     * Filter orders from date (YYYY-MM-DD)
     */
    dateFrom?: string;
    /**
     * Filter orders to date (YYYY-MM-DD)
     */
    dateTo?: string;
};
export namespace consortium_api_business_v1_GetOrdersReq {
    /**
     * Filter by order status
     */
    export enum status {
        PENDING = 'pending',
        PROCESSING = 'processing',
        COMPLETED = 'completed',
        FAILED = 'failed',
        CANCELLED = 'cancelled',
    }
    /**
     * Filter by product type
     */
    export enum type {
        DRUG = 'drug',
        PHYSICAL = 'physical',
        MVR = 'mvr',
        RANDOM = 'random',
    }
}

