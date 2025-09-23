/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_profile_v1_ActivityLogEntry } from './consortium_api_profile_v1_ActivityLogEntry';
/**
 * Result data for certain request according API definition
 */
export type consortium_api_profile_v1_GetActivityLogRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * List of activity log entries
     */
    activities?: Array<consortium_api_profile_v1_ActivityLogEntry>;
    /**
     * Total number of activities
     */
    totalCount?: number;
    /**
     * Current page
     */
    page?: number;
    /**
     * Page size
     */
    pageSize?: number;
};

