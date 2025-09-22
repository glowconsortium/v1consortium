/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_CreateApiKeyReq = {
    /**
     * API key name
     */
    name: string;
    /**
     * API key scopes
     */
    scopes: Array<string>;
    /**
     * Key expiration date (optional)
     */
    expiresAt?: string;
};

