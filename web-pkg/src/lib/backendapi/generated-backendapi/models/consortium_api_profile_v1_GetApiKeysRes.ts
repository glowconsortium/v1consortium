/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_profile_v1_ApiKey } from './consortium_api_profile_v1_ApiKey';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_profile_v1_GetApiKeysRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * List of API keys
     */
    apiKeys?: Array<consortium_api_profile_v1_ApiKey>;
};

