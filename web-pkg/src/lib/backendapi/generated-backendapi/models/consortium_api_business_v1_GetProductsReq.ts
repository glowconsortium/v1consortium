/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_GetProductsReq = {
    /**
     * Filter by product type
     */
    type?: consortium_api_business_v1_GetProductsReq.type;
    /**
     * Filter by active status
     */
    isActive?: boolean;
};
export namespace consortium_api_business_v1_GetProductsReq {
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

