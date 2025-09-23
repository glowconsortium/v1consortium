/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_OrderItem } from './consortium_api_business_v1_OrderItem';
export type consortium_api_business_v1_CreateOrderReq = {
    /**
     * List of order items
     */
    orderItems: Array<consortium_api_business_v1_OrderItem>;
    /**
     * Order priority (default: normal)
     */
    priority?: consortium_api_business_v1_CreateOrderReq.priority;
    /**
     * Order notes
     */
    notes?: string;
};
export namespace consortium_api_business_v1_CreateOrderReq {
    /**
     * Order priority (default: normal)
     */
    export enum priority {
        NORMAL = 'normal',
        URGENT = 'urgent',
        RUSH = 'rush',
    }
}

