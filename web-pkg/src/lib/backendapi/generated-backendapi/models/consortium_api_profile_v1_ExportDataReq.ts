/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_ExportDataReq = {
    /**
     * Export format
     */
    format: consortium_api_profile_v1_ExportDataReq.format;
    /**
     * Data types to export
     */
    dataTypes: Array<string>;
};
export namespace consortium_api_profile_v1_ExportDataReq {
    /**
     * Export format
     */
    export enum format {
        JSON = 'json',
        CSV = 'csv',
    }
}

