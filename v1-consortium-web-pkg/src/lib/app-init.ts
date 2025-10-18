/**
 * App initialization example showing the correct order for initializing stores
 */

import { authStore } from './stores/auth.js';
import { signupStore } from './stores/signupStore.js';

interface AppConfig {
	baseUrl: string;
}

/**
 * Initialize the application stores in the correct order
 * This should be called once when your app starts (e.g., in +layout.svelte or app.js)
 */
export async function initializeApp(config: AppConfig) {
	try {
		console.log('Initializing application stores...');

		// Step 1: Initialize auth store first - this sets up the API client with base URL
		console.log('Initializing auth store...');
		await authStore.initialize(config);

		// Step 2: Initialize signup store - this just verifies auth is ready
		console.log('Initializing signup store...');
		await signupStore.initialize();

		console.log('✅ All stores initialized successfully');
		return true;
	} catch (error) {
		console.error('❌ Store initialization failed:', error);
		throw error;
	}
}

/**
 * Example usage in a Svelte app (e.g., in +layout.svelte)
 */
export function createAppInitializer(config: AppConfig) {
	let isInitialized = false;
	let initError: string | null = null;

	const initialize = async () => {
		try {
			await initializeApp(config);
			isInitialized = true;
			initError = null;
		} catch (error) {
			initError = error instanceof Error ? error.message : 'Initialization failed';
			isInitialized = false;
		}
	};

	// Auto-initialize
	initialize();

	return {
		get isInitialized() { return isInitialized; },
		get error() { return initError; },
		retry: initialize
	};
}

/**
 * Example: Usage in +layout.svelte
 * 
 * <script lang="ts">
 *   import { createAppInitializer } from '$lib/app-init.js';
 *   import { authStore } from '$lib/stores/auth.js';
 *   import { isAuthenticated } from '$lib/stores/auth.js';
 *   
 *   const appInit = createAppInitializer({
 *     baseUrl: 'https://your-api.com'
 *   });
 * </script>
 * 
 * {#if !appInit.isInitialized}
 *   {#if appInit.error}
 *     <div class="error">
 *       Failed to initialize app: {appInit.error}
 *       <button on:click={appInit.retry}>Retry</button>
 *     </div>
 *   {:else}
 *     <div class="loading">Initializing app...</div>
 *   {/if}
 * {:else}
 *   <!-- Your app content -->
 *   <main>
 *     <slot />
 *   </main>
 * {/if}
 */

/**
 * Alternative: Manual initialization pattern
 */
export class AppInitializer {
	private _isInitialized = false;
	private _error: string | null = null;
	private config: AppConfig;

	constructor(config: AppConfig) {
		this.config = config;
	}

	async initialize(): Promise<void> {
		try {
			this._error = null;
			
			// Initialize auth store (sets up API client)
			await authStore.initialize(this.config);
			
			// Initialize other stores that depend on auth
			await signupStore.initialize();
			
			this._isInitialized = true;
		} catch (error) {
			this._error = error instanceof Error ? error.message : 'Initialization failed';
			this._isInitialized = false;
			throw error;
		}
	}

	get isInitialized(): boolean {
		return this._isInitialized;
	}

	get error(): string | null {
		return this._error;
	}

	reset(): void {
		this._isInitialized = false;
		this._error = null;
	}
}

/**
 * Environment-based configuration helper
 */
export function getAppConfig(): AppConfig {
	// In a real app, you might get this from environment variables
	const baseUrl = process.env.NODE_ENV === 'production' 
		? 'https://api.v1consortium.com'
		: 'http://localhost:8080';

	return { baseUrl };
}

/**
 * Complete initialization example
 */
export async function setupApp() {
	const config = getAppConfig();
	
	try {
		await initializeApp(config);
		
		// Now you can use any store safely
		const currentUser = await authStore.getCurrentUser();
		console.log('Current user after init:', currentUser);
		
		return true;
	} catch (error) {
		console.error('App setup failed:', error);
		return false;
	}
}