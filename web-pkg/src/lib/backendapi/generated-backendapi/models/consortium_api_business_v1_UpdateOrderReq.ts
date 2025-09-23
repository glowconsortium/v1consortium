/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_UpdateOrderReq = {
    /**
     * Order ID
     */
    orderId: string;
    /**
     * Order priority
     */
    priority?: consortium_api_business_v1_UpdateOrderReq.priority;
    /**
     * Order notes
     */
    notes?: string;
};
export namespace consortium_api_business_v1_UpdateOrderReq {
    /**
     * Order priority
     */
    export enum priority {
        NORMAL = 'normal',
        URGENT = 'urgent',
        RUSH = 'rush',
    }
}

