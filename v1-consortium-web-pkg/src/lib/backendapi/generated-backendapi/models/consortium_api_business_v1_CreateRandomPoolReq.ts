/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_CreateRandomPoolReq = {
    /**
     * Pool name
     */
    name: string;
    /**
     * Testing frequency
     */
    frequency: consortium_api_business_v1_CreateRandomPoolReq.frequency;
    /**
     * Percentage of drivers to select
     */
    percentage: number;
    /**
     * Driver IDs to include in pool
     */
    driverIds: Array<string>;
};
export namespace consortium_api_business_v1_CreateRandomPoolReq {
    /**
     * Testing frequency
     */
    export enum frequency {
        MONTHLY = 'monthly',
        QUARTERLY = 'quarterly',
        ANNUALLY = 'annually',
    }
}

