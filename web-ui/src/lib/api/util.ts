import { ApiError } from "./client";


// Response wrapper interface that matches the API structure
export interface ApiResponse<T> {
    code?: number;
    message?: string;
    data?: T;
}

// Custom error class for API errors
export class AuthError extends Error {
    public readonly code?: number;
    public readonly statusCode?: number;
    public readonly errorCode?: string;
    public readonly errorId?: string;

    constructor(message: string, code?: number, statusCode?: number, errorCode?: string, errorId?: string) {
        super(message);
        this.name = 'AuthError';
        this.code = code;
        this.statusCode = statusCode;
        this.errorCode = errorCode;
        this.errorId = errorId;
    }
}

// Helper function to handle API responses and errors
export function handleApiResponse<T>(responsePromise: Promise<T>): Promise<T> {
    return responsePromise
        .then((response: unknown) => {
            // Check if the response has an error structure even with 200 status
            if (response && typeof response === 'object' && (response as { code?: number; data?: unknown; message: string }).code !== undefined) {
                // If code is not 200 or data is null, it's an error
                if ((response as { code: number; data?: unknown }).code > 0 || (response as { data?: unknown }).data === null) {
                    let errorMessage = (response as { message?: string }).message || 'Unknown error occurred';
                    let errorCode: string | undefined;
                    let errorId: string | undefined;
                    
                    // Try to extract nested error details from the message
                    if ((response as { message?: string }).message && typeof (response as { message?: string }).message === 'string') {
                        try {
                            // Look for JSON structure in the message
                            const jsonMatch = (response as { message: string }).message.match(/\{[^}]+\}/);
                            if (jsonMatch) {
                                const nestedError = JSON.parse(jsonMatch[0]);
                                errorMessage = nestedError.msg || nestedError.message || errorMessage;
                                errorCode = nestedError.error_code;
                                errorId = nestedError.error_id;
                            }
                        } catch {
                            // If parsing fails, use the original message
                        }
                    }

                    throw new AuthError(errorMessage, (response as { code?: number }).code, undefined, errorCode, errorId);
                }
            }
            
            return response as T;
        })
        .catch((error) => {
            // If it's already an AuthError, re-throw it
            if (error instanceof AuthError) {
                throw error;
            }
            
            if (error instanceof ApiError) {
                // Handle OpenAPI client errors
                let errorMessage = error.message;
                let errorCode: string | undefined;
                let errorId: string | undefined;
                const statusCode = error.status;

                // Try to extract error details from the response body
                if (error.body) {
                    try {
                        const body = typeof error.body === 'string' ? JSON.parse(error.body) : error.body;
                        errorMessage = body.msg || body.message || errorMessage;
                        errorCode = body.error_code;
                        errorId = body.error_id;
                    } catch {
                        // If parsing fails, use the original error message
                    }
                }

                throw new AuthError(errorMessage, undefined, statusCode, errorCode, errorId);
            }

            // Handle other types of errors
            throw new AuthError(error.message || 'Unknown error occurred');
        });
}