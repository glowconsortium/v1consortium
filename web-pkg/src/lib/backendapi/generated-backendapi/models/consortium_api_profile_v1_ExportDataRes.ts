/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Result data for certain request according API definition
 */
export type consortium_api_profile_v1_ExportDataRes = {
    /**
     * Whether data export was initiated successfully
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * Export request ID
     */
    exportId?: string;
    /**
     * Download URL (when ready)
     */
    downloadUrl?: string;
};

