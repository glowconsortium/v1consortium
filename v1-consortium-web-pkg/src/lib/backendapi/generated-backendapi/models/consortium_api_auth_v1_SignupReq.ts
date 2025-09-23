/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_auth_v1_SignupReq = {
    /**
     * User email address
     */
    email: string;
    /**
     * User password
     */
    password: string;
    /**
     * User first name
     */
    firstName: string;
    /**
     * User last name
     */
    lastName: string;
    /**
     * Carrier company name
     */
    carrierName: string;
    /**
     * DOT number
     */
    dotNumber: string;
    /**
     * MC number (optional)
     */
    mcNumber?: string;
    /**
     * Carrier phone number
     */
    phone: string;
    /**
     * Address line 1
     */
    addressLine1: string;
    /**
     * Address line 2 (optional)
     */
    addressLine2?: string;
    /**
     * City
     */
    city: string;
    /**
     * State code (2 letters)
     */
    state: string;
    /**
     * ZIP code
     */
    zipCode: string;
    /**
     * Subscription plan type
     */
    subscriptionPlan: consortium_api_auth_v1_SignupReq.subscriptionPlan;
    /**
     * Stripe payment method ID
     */
    stripePaymentMethodId: string;
    /**
     * Account type
     */
    accountType: consortium_api_auth_v1_SignupReq.accountType;
};
export namespace consortium_api_auth_v1_SignupReq {
    /**
     * Subscription plan type
     */
    export enum subscriptionPlan {
        BASIC = 'basic',
        PREMIUM = 'premium',
    }
    /**
     * Account type
     */
    export enum accountType {
        INDIVIDUAL = 'individual',
        COMPANY = 'company',
    }
}

