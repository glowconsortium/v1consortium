/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_UpdateDriverReq = {
    /**
     * Driver ID
     */
    driverId: string;
    /**
     * Driver first name
     */
    firstName: string;
    /**
     * Driver last name
     */
    lastName: string;
    /**
     * Date of birth (YYYY-MM-DD)
     */
    dateOfBirth: string;
    /**
     * Phone number
     */
    phone: string;
    /**
     * Email address (optional)
     */
    email?: string;
    /**
     * Driver status
     */
    status: consortium_api_business_v1_UpdateDriverReq.status;
    /**
     * Hire date (YYYY-MM-DD)
     */
    hireDate: string;
};
export namespace consortium_api_business_v1_UpdateDriverReq {
    /**
     * Driver status
     */
    export enum status {
        ACTIVE = 'active',
        INACTIVE = 'inactive',
        SUSPENDED = 'suspended',
    }
}

