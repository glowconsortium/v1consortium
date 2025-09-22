/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Result data for certain request according API definition
 */
export type consortium_api_business_v1_GetSubscriptionRes = {
    /**
     * Whether request was successful
     */
    success?: boolean;
    /**
     * Response message
     */
    message?: string;
    /**
     * Stripe subscription ID
     */
    subscriptionId?: string;
    /**
     * Subscription status
     */
    subscriptionStatus?: string;
    /**
     * Subscription plan type
     */
    subscriptionPlan?: string;
    /**
     * Current period start date
     */
    currentPeriodStart?: string;
    /**
     * Current period end date
     */
    currentPeriodEnd?: string;
    /**
     * Renewal month (1-12)
     */
    renewalMonth?: number;
    /**
     * Next billing date
     */
    nextBillingDate?: string;
    /**
     * Whether subscription is active
     */
    isActive?: boolean;
};

