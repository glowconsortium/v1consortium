/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_ApiKey = {
    /**
     * API key ID
     */
    keyId?: string;
    /**
     * API key name
     */
    name?: string;
    /**
     * First few characters of the key
     */
    keyPreview?: string;
    /**
     * API key scopes
     */
    scopes?: Array<string>;
    /**
     * Whether key is active
     */
    isActive?: boolean;
    /**
     * Last usage timestamp
     */
    lastUsedAt?: string;
    /**
     * Key expiration timestamp
     */
    expiresAt?: string;
    /**
     * Key creation timestamp
     */
    createdAt?: string;
};

