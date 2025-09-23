/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_profile_v1_UserSession } from './consortium_api_profile_v1_UserSession';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_profile_v1_GetSessionsRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * List of active sessions
     */
    sessions?: Array<consortium_api_profile_v1_UserSession>;
};

