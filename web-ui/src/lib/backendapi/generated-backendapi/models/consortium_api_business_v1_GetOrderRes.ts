/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_Order } from './consortium_api_business_v1_Order';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_business_v1_GetOrderRes = {
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

