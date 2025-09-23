/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_business_v1_UpdateSubscriptionReq = {
    /**
     * New subscription plan
     */
    subscriptionPlan: consortium_api_business_v1_UpdateSubscriptionReq.subscriptionPlan;
    /**
     * Stripe payment method ID
     */
    stripePaymentMethodId: string;
};
export namespace consortium_api_business_v1_UpdateSubscriptionReq {
    /**
     * New subscription plan
     */
    export enum subscriptionPlan {
        BASIC = 'basic',
        PREMIUM = 'premium',
    }
}

