/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_RandomPool } from './consortium_api_business_v1_RandomPool';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_business_v1_GetRandomPoolsRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * List of random pools
     */
    pools?: Array<consortium_api_business_v1_RandomPool>;
};

