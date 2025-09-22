/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_GetActivityLogReq = {
    /**
     * Page number (default: 1)
     */
    page?: number;
    /**
     * Page size (default: 20)
     */
    pageSize?: number;
    /**
     * Filter by action type
     */
    action?: string;
    /**
     * Filter activities from date (YYYY-MM-DD)
     */
    dateFrom?: string;
    /**
     * Filter activities to date (YYYY-MM-DD)
     */
    dateTo?: string;
};

