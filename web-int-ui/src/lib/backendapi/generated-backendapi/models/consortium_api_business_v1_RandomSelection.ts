/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_RandomSelection = {
    /**
     * Selection ID
     */
    selectionId?: string;
    /**
     * Pool ID
     */
    poolId?: string;
    /**
     * Selection run timestamp
     */
    runAt?: string;
    /**
     * Random seed used
     */
    seed?: string;
    /**
     * Selection verification hash
     */
    selectionHash?: string;
    /**
     * Selected driver IDs
     */
    selectedDrivers?: Array<string>;
    /**
     * Selection status
     */
    status?: string;
    /**
     * Selection report URL
     */
    reportUrl?: string;
    /**
     * Creation timestamp
     */
    createdAt?: string;
};

