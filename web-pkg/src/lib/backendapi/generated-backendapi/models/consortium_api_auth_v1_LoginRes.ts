/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Result data for certain request according API definition
 */
export type consortium_api_auth_v1_LoginRes = {
    /**
     * Whether login was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * User ID
     */
    userId?: string;
    /**
     * Associated carrier ID
     */
    carrierId?: string;
    /**
     * User role
     */
    role?: string;
    /**
     * JWT access token
     */
    accessToken?: string;
    /**
     * JWT refresh token
     */
    refreshToken?: string;
};

