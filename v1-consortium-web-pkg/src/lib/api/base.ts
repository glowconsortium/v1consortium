import type { Transport, Client } from "@connectrpc/connect";
import { createClient } from "@connectrpc/connect";
import type { GenService } from "@bufbuild/protobuf/codegenv2";

export interface ServiceConfig {
  transport: Transport;
  baseUrl: string;
  getToken: () => string | null;
  setToken: (token: string | null) => void;
  onAuthError?: () => void;
}

/**
 * Base service class that provides common functionality for all service clients
 */
export abstract class BaseService {
  protected transport: Transport;
  protected baseUrl: string;
  protected getToken: () => string | null;
  protected setToken: (token: string | null) => void;
  protected onAuthError?: () => void;

  constructor(config: ServiceConfig) {
    this.transport = config.transport;
    this.baseUrl = config.baseUrl;
    this.getToken = config.getToken;
    this.setToken = config.setToken;
    this.onAuthError = config.onAuthError;
  }

  /**
   * Create a client for the given service
   */
  protected createClient<T extends GenService<any>>(service: T): Client<T> {
    return createClient(service, this.transport);
  }

  /**
   * Handle common API errors
   */
  protected handleError(error: unknown, operation: string): never {
    console.error(`${operation} failed:`, error);
    
    if (error instanceof Error) {
      // Check for authentication errors
      if (error.message.includes('401') || error.message.includes('Unauthorized')) {
        this.onAuthError?.();
        throw new Error(`Authentication failed during ${operation}. Please login again.`);
      }
      
      // Check for authorization errors
      if (error.message.includes('403') || error.message.includes('Forbidden')) {
        throw new Error(`You don't have permission to perform this ${operation}.`);
      }
      
      // Check for network errors
      if (error.message.includes('fetch') || error.message.includes('network')) {
        throw new Error(`Network error during ${operation}. Please check your connection.`);
      }
      
      throw new Error(`${operation} failed: ${error.message}`);
    }
    
    throw new Error(`${operation} failed: ${String(error)}`);
  }

  /**
   * Retry a request with token refresh if it fails with auth error
   */
  protected async withRetry<T>(
    operation: () => Promise<T>,
    operationName: string,
    refreshToken?: () => Promise<boolean>
  ): Promise<T> {
    try {
      return await operation();
    } catch (error) {
      // If we get an auth error and have a refresh function, try refreshing
      if (
        refreshToken && 
        error instanceof Error && 
        (error.message.includes('401') || error.message.includes('Unauthorized'))
      ) {
        const refreshed = await refreshToken();
        if (refreshed) {
          // Retry the operation with the new token
          try {
            return await operation();
          } catch (retryError) {
            this.handleError(retryError, `${operationName} (retry after token refresh)`);
          }
        }
      }
      
      this.handleError(error, operationName);
    }
  }

  /**
   * Get current authentication token
   */
  protected getCurrentToken(): string | null {
    return this.getToken();
  }

  /**
   * Check if we have a valid token
   */
  protected hasValidToken(): boolean {
    const token = this.getCurrentToken();
    return !!token && !this.isTokenExpired(token);
  }

  /**
   * Check if a token is expired
   */
  protected isTokenExpired(token: string): boolean {
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      let exp: number;
      if (typeof payload.exp === 'bigint') {
        exp = Number(payload.exp) * 1000;
      } else {
        exp = payload.exp * 1000;
      }
      return Date.now() >= exp;
    } catch (error) {
      console.error('Error checking token expiration:', error);
      return true;
    }
  }

  /**
   * Require authentication for an operation
   */
  protected requireAuth(operationName: string): void {
    if (!this.hasValidToken()) {
      throw new Error(`${operationName} requires authentication. Please login first.`);
    }
  }
}