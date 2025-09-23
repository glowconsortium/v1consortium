/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_GetRandomSelectionsReq = {
    /**
     * Page number (default: 1)
     */
    page?: number;
    /**
     * Page size (default: 20)
     */
    pageSize?: number;
    /**
     * Filter by pool ID
     */
    poolId?: string;
    /**
     * Filter selections from date (YYYY-MM-DD)
     */
    dateFrom?: string;
    /**
     * Filter selections to date (YYYY-MM-DD)
     */
    dateTo?: string;
};

