<script lang="ts">
  import { onMount } from 'svelte';
  import { authStore, isAuthenticated, authLoading, authError } from '@movsm/v1-consortium-web-pkg';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
	import { PUBLIC_AUTH0_AUDIENCE, PUBLIC_AUTH0_CLIENT_ID, PUBLIC_AUTH0_DOMAIN } from '$env/static/public';

  let processing = true;
  let errorMessage = '';

  onMount(async () => {
    try {
      console.log('Callback page mounted, URL:', window.location.href);
      
      // Check if we have Auth0 configuration
      if (!PUBLIC_AUTH0_DOMAIN || !PUBLIC_AUTH0_CLIENT_ID) {
        throw new Error('Auth0 configuration missing. Please check your environment variables.');
      }

      // Initialize Auth0 to handle the callback
      await authStore.initialize({
        domain: PUBLIC_AUTH0_DOMAIN,
        clientId: PUBLIC_AUTH0_CLIENT_ID,
        audience: PUBLIC_AUTH0_AUDIENCE,
        scope: 'openid profile email',
        redirectUri: `${window.location.origin}/auth/callback`
      });

      console.log('Auth0 initialized in callback');

      // Wait for auth state to be determined
      await new Promise(resolve => setTimeout(resolve, 2000));

      console.log('Auth state after initialization:', {
        isAuthenticated: $isAuthenticated,
        error: $authError,
        loading: $authLoading
      });

      if ($isAuthenticated) {
        // Get the intended destination from URL params or default to dashboard
        const returnTo = $page.url.searchParams.get('returnTo') || '/dashboard';
        console.log('User authenticated, redirecting to:', returnTo);
        goto(returnTo);
      } else if ($authError) {
        console.error('Authentication error:', $authError);
        errorMessage = $authError;
        processing = false;
      } else {
        // If not authenticated and no error, redirect to signin
        console.log('Not authenticated, redirecting to signin');
        goto('/auth/signin');
      }
    } catch (error) {
      console.error('Callback processing failed:', error);
      errorMessage = error instanceof Error ? error.message : 'Authentication failed';
      processing = false;
    }
  });

  function handleRetry() {
    goto('/auth/signin');
  }
</script>

<svelte:head>
  <title>Authenticating... - V1 Consortium</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-gray-50">
  <div class="max-w-md w-full text-center">
    {#if processing && !errorMessage}
      <div class="space-y-4">
        <div class="flex justify-center">
          <svg class="animate-spin h-12 w-12 text-indigo-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
        </div>
        <h2 class="text-2xl font-bold text-gray-900">Completing sign in...</h2>
        <p class="text-gray-600">Please wait while we authenticate you.</p>
      </div>
    {:else if errorMessage}
      <div class="space-y-4">
        <div class="flex justify-center">
          <svg class="h-12 w-12 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.728-.833-2.498 0L4.346 15.5c-.77.833.192 2.5 1.732 2.5z" />
          </svg>
        </div>
        <h2 class="text-2xl font-bold text-gray-900">Authentication Failed</h2>
        <p class="text-gray-600">{errorMessage}</p>
        <button
          type="button"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          on:click={handleRetry}
        >
          Try Again
        </button>
      </div>
    {/if}
  </div>
</div>