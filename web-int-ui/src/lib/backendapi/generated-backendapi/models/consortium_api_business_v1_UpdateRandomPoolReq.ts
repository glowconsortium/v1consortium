/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_UpdateRandomPoolReq = {
    /**
     * Pool ID
     */
    poolId: string;
    /**
     * Pool name
     */
    name: string;
    /**
     * Testing frequency
     */
    frequency: consortium_api_business_v1_UpdateRandomPoolReq.frequency;
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
};
export namespace consortium_api_business_v1_UpdateRandomPoolReq {
    /**
     * Testing frequency
     */
    export enum frequency {
        MONTHLY = 'monthly',
        QUARTERLY = 'quarterly',
        ANNUALLY = 'annually',
    }
}

