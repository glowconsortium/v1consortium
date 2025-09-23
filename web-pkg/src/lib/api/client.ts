import { PUBLIC_API_URL } from '$env/static/public';
import {
    OpenAPI,
    HelloService,
    AuthService,
    ProfileService,
    BusinessService,
    ApiError as FormApiError,


} from '../backendapi/generated-backendapi/index.js';
import { authenticatedFetch } from './interceptor.js';



/**
 * API Client Factory that provides access to all generated services
 * Uses the OpenAPI generated client as the underlying implementation
 */
export class ApiClientFactory {
    private baseUrl: string;
    private accessToken: string | null = null;

    constructor(baseUrl: string = '') {
        this.baseUrl = baseUrl;
        this.configure();
    }

    /**
     * Configure the OpenAPI client with current settings
     */
    private configure() {
        OpenAPI.BASE = this.baseUrl;
        OpenAPI.TOKEN = this.accessToken || undefined;
        OpenAPI.WITH_CREDENTIALS = true;
        OpenAPI.CREDENTIALS = 'include';
    }

    /**
     * Set the access token for authenticated requests
     */
    setAccessToken(token: string | null) {
        this.accessToken = token;
        this.configure();
    }

    /**
     * Set the base URL for API requests
     */
    setBaseUrl(baseUrl: string) {
        this.baseUrl = baseUrl;
        this.configure();
    }

    /**
     * Get the authentication service
     */
    get auth() {
        return AuthService;
    }


    /**
     * Get the hello service (for health checks)
     */
    get hello() {
        return HelloService;
    }

    /**
     * Get the business service
     */
    get business() {
        return BusinessService;
    }

    /**
     * Get the profile service
     */
    get profile() {
        return ProfileService;
    }

    // /**
    //  * Get the TControl service
    //  */
    // get tcontrol() {
    //     return TControlService;
    // }

    // get subscription() {
    //     return SubscriptionService;
    // }
}

// Re-export the generated ApiError for convenience
export { FormApiError as ApiError };

// Create a singleton instance of the generated client
export const generatedApiClient = new ApiClientFactory(PUBLIC_API_URL || '');


// Create a custom apiClient object that uses authenticatedFetch
export const apiClient = {
    baseUrl: PUBLIC_API_URL || '',

    setBaseUrl(url: string) {
        this.baseUrl = url;
        // Also update the generated client
        generatedApiClient.setBaseUrl(url);
    },

    setAccessToken(token: string | null) {
        generatedApiClient.setAccessToken(token);
    },

    // Use authenticated fetch for all requests
    async request(endpoint: string, options: RequestInit = {}) {
        const url = `${this.baseUrl}${endpoint}`;
        return authenticatedFetch(url, options);
    },

    // Provide access to generated services
    get auth() {
        return generatedApiClient.auth;
    },



    get hello() {
        return generatedApiClient.hello;
    },

    // get organization() {
    //     return generatedApiClient.organization;
    // },

    // get profile() {
    //     return generatedApiClient.profile;
    // },

    // get tcontrol() {
    //     return generatedApiClient.tcontrol;
    // },

    // get subscription() {
    //     return generatedApiClient.subscription;
    // }
};

