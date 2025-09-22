/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Result data for certain request according API definition
 */
export type consortium_api_profile_v1_CreateApiKeyRes = {
    /**
     * Whether API key was created successfully
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * Created API key ID
     */
    keyId?: string;
    /**
     * Generated API key (shown only once)
     */
    apiKey?: string;
};

