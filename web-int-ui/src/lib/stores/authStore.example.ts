// Example usage of the authStore

import { createAuthStore, setAuthStore, isAuthenticated, currentUser } from '$lib/stores/authStore';
import { createConnectTransport } from '@connectrpc/connect-web';

// Create transport
const transport = createConnectTransport({
    baseUrl: 'http://localhost:8000',
    useBinaryFormat: true,
});

// Create auth store instance
const authStoreInstance = createAuthStore(transport, {
    // Additional configuration args can be passed here
});

// Set the global instance for derived stores
setAuthStore(authStoreInstance);

// Example usage in a Svelte component:
/*
<script lang="ts">
    import { authStore, isAuthenticated, currentUser, authError, authLoading } from '$lib/stores/authStore';
    import { createConnectTransport } from '@connectrpc/connect-web';
    
    // Create and initialize auth store
    const transport = createConnectTransport({
        baseUrl: 'http://localhost:8000',
        useBinaryFormat: true,
    });
    
    const auth = createAuthStore(transport);
    
    let email = '';
    let password = '';
    
    async function handleLogin() {
        try {
            await auth.login(email, password, true);
            console.log('Login successful');
        } catch (error) {
            console.error('Login failed:', error);
        }
    }
    
    async function handleLogout() {
        try {
            await auth.logout();
            console.log('Logout successful');
        } catch (error) {
            console.error('Logout failed:', error);
        }
    }
    
    function checkPermissions() {
        // Synchronous role check
        const isAdmin = auth.hasRole('admin');
        
        // Asynchronous permission check
        auth.hasPermission('read', 'resource-123').then(canRead => {
            console.log('Can read:', canRead);
        });
    }
</script>

{#if $isAuthenticated}
    <div>
        <h1>Welcome, {$currentUser?.firstName} {$currentUser?.lastName}!</h1>
        <p>Email: {$currentUser?.email}</p>
        <p>Role: {$currentUser?.role}</p>
        <p>Organization: {$currentUser?.organizationName}</p>
        
        <button on:click={handleLogout}>Logout</button>
        <button on:click={checkPermissions}>Check Permissions</button>
    </div>
{:else}
    <form on:submit|preventDefault={handleLogin}>
        <input type="email" bind:value={email} placeholder="Email" required />
        <input type="password" bind:value={password} placeholder="Password" required />
        <button type="submit" disabled={$authLoading}>
            {$authLoading ? 'Logging in...' : 'Login'}
        </button>
        
        {#if $authError}
            <p class="error">{$authError}</p>
        {/if}
    </form>
{/if}
*/