/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_Product } from './consortium_api_business_v1_Product';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_business_v1_GetProductsRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * List of products
     */
    products?: Array<consortium_api_business_v1_Product>;
};

