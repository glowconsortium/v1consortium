/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type consortium_api_profile_v1_UserSession = {
    /**
     * Session ID
     */
    sessionId?: string;
    /**
     * IP address
     */
    ipAddress?: string;
    /**
     * User agent
     */
    userAgent?: string;
    /**
     * Approximate location
     */
    location?: string;
    /**
     * Whether this is the current session
     */
    isCurrentSession?: boolean;
    /**
     * Session start timestamp
     */
    createdAt?: string;
    /**
     * Last activity timestamp
     */
    lastActiveAt?: string;
};

