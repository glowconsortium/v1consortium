/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_RandomSelection } from './consortium_api_business_v1_RandomSelection';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_business_v1_GetRandomSelectionsRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * List of random selections
     */
    selections?: Array<consortium_api_business_v1_RandomSelection>;
    /**
     * Total number of selections
     */
    totalCount?: number;
    /**
     * Current page
     */
    page?: number;
    /**
     * Page size
     */
    pageSize?: number;
};

