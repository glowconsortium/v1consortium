/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_UpdateNotificationPrefsReq = {
    /**
     * Enable email notifications
     */
    emailNotifications?: boolean;
    /**
     * Enable SMS notifications
     */
    smsNotifications?: boolean;
    /**
     * Notify on order status updates
     */
    orderUpdates?: boolean;
    /**
     * Notify on random testing selections
     */
    randomSelections?: boolean;
    /**
     * Notify on expiring documents
     */
    expiringDocuments?: boolean;
    /**
     * Notify on system maintenance
     */
    systemMaintenance?: boolean;
};

