<script lang="ts">
    import { onMount } from 'svelte';
    import { authService } from '@movsm/v1-consortium-web-pkg';
    import { Button } from '@movsm/v1-consortium-web-pkg';
    import { Card } from '@movsm/v1-consortium-web-pkg';
    import { Icon } from '@movsm/v1-consortium-web-pkg';
    import { Badge } from '@movsm/v1-consortium-web-pkg';
    import { Modal } from '@movsm/v1-consortium-web-pkg';
    import type { v2oneglobe_api_profile_v1_Session } from '$lib/backendapi/generated-backendapi';
	import { toaster } from '@movsm/v1-consortium-web-pkg';

    // Runes for reactive state
    let loading = $state(false);
    let error = $state('');
    let success = $state('');
    let sessions: v2oneglobe_api_profile_v1_Session[] = $state([]);
    let showRevokeAllModal = $state(false);

    onMount(async () => {
        await loadSessions();
    });

    async function loadSessions() {
        try {
            loading = true;
            const response = await authService.getSessions();
            // Fix: Access nested data structure - try multiple possible paths
            sessions = response.data?.sessions || response.data?.sessions || response.data?.sessions || [];
            console.log('Sessions loaded:', sessions);
        } catch (err: any) {
            console.error('Failed to load sessions:', err);
            error = err.message || 'Failed to load sessions';
        } finally {
            loading = false;
        }
    }

    async function revokeSession(sessionId: string) {
        try {
            loading = true;
            error = '';

            await authService.revokeSession(sessionId);
            success = 'Session revoked successfully';
            await loadSessions();
        } catch (err: any) {
            error = err.message || 'Failed to revoke session';
        } finally {
            loading = false;
        }
    }

    async function revokeAllSessions() {
        try {
            loading = true;
            error = '';

            await authService.revokeAllSessions();
            success = 'All sessions revoked successfully';
            showRevokeAllModal = false;
            await loadSessions();
        } catch (err: any) {
            error = err.message || 'Failed to revoke all sessions';
        } finally {
            loading = false;
        }
    }

    // Helper function to parse user agent for device info
    function parseUserAgent(userAgent: string) {
        const ua = userAgent || '';
        
        // Detect OS
        let os = 'Unknown';
        if (ua.includes('Windows')) os = 'Windows';
        else if (ua.includes('Mac OS')) os = 'macOS';
        else if (ua.includes('Linux')) os = 'Linux';
        else if (ua.includes('Android')) os = 'Android';
        else if (ua.includes('iPhone') || ua.includes('iPad')) os = 'iOS';

        // Detect browser
        let browser = 'Unknown';
        if (ua.includes('Chrome') && !ua.includes('Edg')) browser = 'Chrome';
        else if (ua.includes('Firefox')) browser = 'Firefox';
        else if (ua.includes('Safari') && !ua.includes('Chrome')) browser = 'Safari';
        else if (ua.includes('Edg')) browser = 'Edge';

        // Detect device type
        let deviceType = 'desktop';
        if (ua.includes('Mobile') || ua.includes('Android')) deviceType = 'mobile';
        else if (ua.includes('Tablet') || ua.includes('iPad')) deviceType = 'tablet';

        return { os, browser, deviceType };
    }

    function getDeviceIcon(userAgent: string) {
        const { deviceType } = parseUserAgent(userAgent);
        switch (deviceType) {
            case 'mobile': return 'smartphone';
            case 'tablet': return 'tablet';
            case 'desktop': return 'monitor';
            default: return 'monitor';
        }
    }

    function getBrowserIcon(userAgent: string) {
        const { browser } = parseUserAgent(userAgent);
        switch (browser?.toLowerCase()) {
            case 'chrome': return 'chrome';
            case 'firefox': return 'firefox';
            case 'safari': return 'safari';
            case 'edge': return 'edge';
            default: return 'globe';
        }
    }

    function formatLastActivity(timestamp: string) {
        if (!timestamp) return 'Unknown';
        
        const date = new Date(timestamp);
        const now = new Date();
        const diffInMinutes = Math.floor((now.getTime() - date.getTime()) / (1000 * 60));
        
        if (diffInMinutes < 1) return 'Just now';
        if (diffInMinutes < 60) return `${diffInMinutes} minutes ago`;
        if (diffInMinutes < 1440) return `${Math.floor(diffInMinutes / 60)} hours ago`;
        return date.toLocaleDateString();
    }

    // Clear messages after 5 seconds
    $effect(() => {
        if (error || success) {
            const timer = setTimeout(() => {
                error = '';
                success = '';
            }, 5000);
            
            return () => clearTimeout(timer);
        }
    });
</script>

<svelte:head>
    <title>Active Sessions - FormApp</title>
</svelte:head>

<div class="max-w-4xl mx-auto space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
        <div>
            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
                Active Sessions
            </h1>
            <p class="text-gray-600 dark:text-gray-400 mt-1">
                Manage your active login sessions across different devices
            </p>
        </div>

        {#if sessions.length > 1}
            <Button 
                variant="outline"
                onclick={() => showRevokeAllModal = true}
                classes="flex items-center gap-2"
            >
                <Icon name="x-circle" size="4" />
                Revoke All Sessions
            </Button>
        {/if}
    </div>

    <!-- Toast Messages -->
    {#if error}
        {toaster.create({
            type: 'error',
            title: error
        })}
    {/if}

    {#if success}
        {toaster.create({
            type: 'success',
            title: success
        })}
    {/if}

    <!-- Loading State -->
    {#if loading && sessions.length === 0}
        <div class="flex items-center justify-center py-12">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>
    {:else if sessions.length === 0}
        <Card>
            <div class="text-center py-12">
                <Icon name="monitor" size="12" classes="mx-auto text-gray-400 mb-4" />
                <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">
                    No active sessions
                </h3>
                <p class="text-gray-600 dark:text-gray-400">
                    You don't have any active sessions at the moment.
                </p>
            </div>
        </Card>
    {:else}
        <div class="space-y-4">
            {#each sessions as session (session.id)}
                {@const { os, browser, deviceType } = parseUserAgent(session.user_agent || '')}
                <Card classes={session.is_current ? 'border-blue-200 dark:border-blue-800 bg-blue-50 dark:bg-blue-900/20' : ''}>
                    <div class="flex items-center justify-between">
                        <div class="flex items-center space-x-4">
                            <!-- Device Icon -->
                            <div class="flex-shrink-0">
                                <div class="h-12 w-12 bg-gray-100 dark:bg-gray-700 rounded-lg flex items-center justify-center">
                                    <Icon 
                                        name={getDeviceIcon(session.user_agent || '')} 
                                        size="6" 
                                    />
                                </div>
                            </div>

                            <!-- Session Info -->
                            <div class="flex-1 min-w-0">
                                <div class="flex items-center space-x-2 mb-1">
                                    <h3 class="text-sm font-medium text-gray-900 dark:text-white truncate">
                                        {`${os} ${browser}`}
                                    </h3>
                                    
                                    {#if session.is_current}
                                        <Badge color="success" size="sm">Current Session</Badge>
                                    {/if}
                                </div>

                                <div class="flex items-center space-x-4 text-sm text-gray-600 dark:text-gray-400">
                                    <div class="flex items-center space-x-1">
                                        <Icon name={getBrowserIcon(session.user_agent || '')} size="4" />
                                        <span>{browser} on {os}</span>
                                    </div>
                                    
                                    {#if session.location}
                                        <div class="flex items-center space-x-1">
                                            <Icon name="map-pin" size="4" />
                                            <span>{session.location}</span>
                                        </div>
                                    {/if}
                                </div>

                                <div class="flex items-center space-x-4 text-xs text-gray-500 dark:text-gray-500 mt-1">
                                    {#if session.ip}
                                        <span>IP: {session.ip}</span>
                                    {/if}
                                    <span>Last active: {formatLastActivity(session.last_used || '')}</span>
                                    {#if session.created_at}
                                        <span>Created: {new Date(session.created_at).toLocaleDateString()}</span>
                                    {/if}
                                </div>
                            </div>
                        </div>

                        <!-- Actions -->
                        <div class="flex items-center space-x-2">
                            {#if !session.is_current}
                                <Button
                                    variant="ghost"
                                    size="sm"
                                    onclick={() => revokeSession(session.id || '')}
                                    disabled={loading}
                                    classes="text-red-600 hover:text-red-900"
                                >
                                    <Icon name="x" size="4" />
                                    Revoke
                                </Button>
                            {:else}
                                <span class="text-xs text-gray-500 dark:text-gray-500 px-2">
                                    This session
                                </span>
                            {/if}
                        </div>
                    </div>
                </Card>
            {/each}
        </div>

        <!-- Session Security Tips -->
        <Card>
            <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">
                Session Security Tips
            </h3>
            
            <div class="space-y-3 text-sm text-gray-600 dark:text-gray-400">
                <div class="flex items-start">
                    <Icon name="shield" size="4" classes="text-blue-500 mr-2 mt-0.5" />
                    <p>Always log out from shared or public computers</p>
                </div>
                
                <div class="flex items-start">
                    <Icon name="eye" size="4" classes="text-blue-500 mr-2 mt-0.5" />
                    <p>Review your active sessions regularly and revoke any you don't recognize</p>
                </div>
                
                <div class="flex items-start">
                    <Icon name="alert-triangle" size="4" classes="text-blue-500 mr-2 mt-0.5" />
                    <p>If you see suspicious activity, change your password immediately</p>
                </div>
                
                <div class="flex items-start">
                    <Icon name="lock" size="4" classes="text-blue-500 mr-2 mt-0.5" />
                    <p>Enable two-factor authentication for additional security</p>
                </div>
            </div>
        </Card>
    {/if}
</div>

<!-- Revoke All Sessions Modal -->
{#if showRevokeAllModal}
    <Modal 
        title="Revoke All Sessions" 
        oncancel={() => showRevokeAllModal = false}
    >
    {#snippet  content()}
        
        <div class="space-y-4">
            <div class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
                <div class="flex items-start">
                    <Icon name="alert-triangle" size="5" classes="text-red-600 dark:text-red-400 mr-2 mt-0.5" />
                    <div>
                        <h4 class="text-red-800 dark:text-red-400 font-medium">
                            This will revoke all sessions
                        </h4>
                        <p class="text-sm text-red-700 dark:text-red-300 mt-1">
                            You will be logged out from all devices except this one. You'll need to log in again on other devices.
                        </p>
                    </div>
                </div>
            </div>

            <p class="text-gray-700 dark:text-gray-300">
                Are you sure you want to revoke all active sessions? This action cannot be undone.
            </p>

            <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200 dark:border-gray-600">
                <Button
                    variant="ghost"
                    onclick={() => showRevokeAllModal = false}
                    disabled={loading}
                >
                    Cancel
                </Button>
                <Button
                    variant="filled"
                    onclick={revokeAllSessions}
                    loading={loading}
                >
                    Revoke All Sessions
                </Button>
            </div>
        </div>
        {/snippet}
    </Modal>
{/if}