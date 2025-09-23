/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_OrderItem } from './consortium_api_business_v1_OrderItem';
export type consortium_api_business_v1_Order = {
    /**
     * Order ID
     */
    orderId?: string;
    /**
     * Carrier ID
     */
    carrierId?: string;
    /**
     * Order status
     */
    status?: string;
    /**
     * Order priority
     */
    priority?: string;
    /**
     * Order total amount
     */
    total?: number;
    /**
     * Order items
     */
    orderItems?: Array<consortium_api_business_v1_OrderItem>;
    /**
     * Order notes
     */
    notes?: string;
    /**
     * Creation timestamp
     */
    createdAt?: string;
    /**
     * Last update timestamp
     */
    updatedAt?: string;
    /**
     * Completion timestamp
     */
    completedAt?: string;
    /**
     * Temporal workflow ID
     */
    workflowId?: string;
    /**
     * Vendor order ID
     */
    vendorOrderId?: string;
    /**
     * Results download URL
     */
    resultsUrl?: string;
};

