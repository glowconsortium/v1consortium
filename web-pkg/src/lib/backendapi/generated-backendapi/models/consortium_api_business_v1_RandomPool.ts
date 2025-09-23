/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_RandomPool = {
    /**
     * Pool ID
     */
    poolId?: string;
    /**
     * Carrier ID
     */
    carrierId?: string;
    /**
     * Pool name
     */
    name?: string;
    /**
     * Testing frequency
     */
    frequency?: string;
    /**
     * Percentage of drivers to select
     */
    percentage?: number;
    /**
     * Driver IDs in the pool
     */
    driverIds?: Array<string>;
    /**
     * Whether pool is active
     */
    isActive?: boolean;
    /**
     * Creation timestamp
     */
    createdAt?: string;
    /**
     * Last update timestamp
     */
    updatedAt?: string;
};

