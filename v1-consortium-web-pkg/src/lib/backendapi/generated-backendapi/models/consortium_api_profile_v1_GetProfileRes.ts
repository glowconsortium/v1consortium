/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_profile_v1_UserProfile } from './consortium_api_profile_v1_UserProfile';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_profile_v1_GetProfileRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * User profile data
     */
    profile?: consortium_api_profile_v1_UserProfile;
};

