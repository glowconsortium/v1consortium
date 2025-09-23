/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_profile_v1_NotificationPreferences } from './consortium_api_profile_v1_NotificationPreferences';
export type consortium_api_profile_v1_UserProfile = {
    /**
     * User ID
     */
    userId?: string;
    /**
     * User email address
     */
    email?: string;
    /**
     * User first name
     */
    firstName?: string;
    /**
     * User last name
     */
    lastName?: string;
    /**
     * User role (admin, staff, viewer)
     */
    role?: string;
    /**
     * Whether email is verified
     */
    emailVerified?: boolean;
    /**
     * Last login timestamp
     */
    lastLoginAt?: string;
    /**
     * Account creation timestamp
     */
    createdAt?: string;
    /**
     * Last profile update timestamp
     */
    updatedAt?: string;
    /**
     * Associated carrier ID
     */
    carrierId?: string;
    /**
     * Associated carrier name
     */
    carrierName?: string;
    /**
     * Role within the carrier organization
     */
    carrierRole?: string;
    /**
     * Whether profile is complete
     */
    profileComplete?: boolean;
    /**
     * List of missing required fields
     */
    missingFields?: Array<string>;
    /**
     * User timezone
     */
    timezone?: string;
    /**
     * Preferred language
     */
    language?: string;
    /**
     * Notification preferences
     */
    notificationPrefs?: consortium_api_profile_v1_NotificationPreferences;
};

