/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_UpdateApiKeyReq = {
    /**
     * API key ID
     */
    keyId: string;
    /**
     * API key name
     */
    name: string;
    /**
     * API key scopes
     */
    scopes: Array<string>;
    /**
     * Whether key is active
     */
    isActive?: boolean;
};

