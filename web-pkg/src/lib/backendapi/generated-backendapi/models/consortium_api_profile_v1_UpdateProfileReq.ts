/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_UpdateProfileReq = {
    /**
     * User first name
     */
    firstName: string;
    /**
     * User last name
     */
    lastName: string;
    /**
     * User timezone
     */
    timezone?: string;
    /**
     * Preferred language
     */
    language?: consortium_api_profile_v1_UpdateProfileReq.language;
};
export namespace consortium_api_profile_v1_UpdateProfileReq {
    /**
     * Preferred language
     */
    export enum language {
        EN = 'en',
        ES = 'es',
    }
}

