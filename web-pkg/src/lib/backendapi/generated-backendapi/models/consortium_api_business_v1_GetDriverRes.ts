/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_business_v1_Driver } from './consortium_api_business_v1_Driver';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_business_v1_GetDriverRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * Driver details
     */
    driver?: consortium_api_business_v1_Driver;
};

