/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Result data for certain request according API definition
 */
export type consortium_api_auth_v1_SignupRes = {
    /**
     * Whether signup was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * Created user ID
     */
    userId?: string;
    /**
     * Created carrier ID
     */
    carrierId?: string;
    /**
     * JWT access token
     */
    accessToken?: string;
    /**
     * JWT refresh token
     */
    refreshToken?: string;
};

