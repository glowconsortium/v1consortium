<script lang="ts">
    // import { onMount } from 'svelte';
    // import { authService } from '@movsm/v1-consortium-web-pkg';
    // import { Button } from '@movsm/v1-consortium-web-pkg';
    // import { Card } from '@movsm/v1-consortium-web-pkg';
    // import { Icon } from '@movsm/v1-consortium-web-pkg';
    // import { Badge } from '@movsm/v1-consortium-web-pkg';
    // import { Toggle } from '@movsm/v1-consortium-web-pkg';
    // import type { v2oneglobe_api_profile_v1_SecuritySettings } from '@movsm/v1-consortium-web-pkg';
	// import { toaster } from '@movsm/v1-consortium-web-pkg';

    // // Runes for reactive state
    // let loading = $state(false);
    // let error = $state('');
    // let success = $state('');
    // let securitySettings: v2oneglobe_api_profile_v1_SecuritySettings | null = $state(null);

    // onMount(async () => {
    //     await loadSecuritySettings();
    // });

    // async function loadSecuritySettings() {
    //     try {
    //         loading = true;
    //         const response = await authService.getSecuritySettings();
    //         // Fix: Access nested data structure and provide defaults
    //         securitySettings = response.data?.security || response.data?.security || {
    //             two_factor_enabled: false,
    //             email_notifications: true,
    //             security_alerts: true,
    //             session_timeout_minutes: 1440, // 24 hours
    //             allow_multiple_sessions: true,
    //             require_password_for_actions: false
    //         };
    //         console.log('Security settings loaded:', securitySettings);
    //     } catch (err: any) {
    //         console.error('Failed to load security settings:', err);
    //         error = err.message || 'Failed to load security settings';
    //         // Set default settings on error
    //         securitySettings = {
    //             two_factor_enabled: false,
    //             email_notifications: true,
    //             security_alerts: true,
    //             session_timeout_minutes: 1440,
    //             allow_multiple_sessions: true,
    //             require_password_for_actions: false
    //         };
    //     } finally {
    //         loading = false;
    //     }
    // }

    // async function updateSecuritySettings(updates: Partial<v2oneglobe_api_profile_v1_SecuritySettings>) {
    //     try {
    //         loading = true;
    //         error = '';

    //         await authService.updateSecuritySettings(updates);
    //         success = 'Security settings updated successfully';
            
    //         // Update local state
    //         securitySettings = { ...securitySettings, ...updates };
    //     } catch (err: any) {
    //         console.error('Failed to update security settings:', err);
    //         error = err.message || 'Failed to update security settings';
    //     } finally {
    //         loading = false;
    //     }
    // }

    // // Helper function to convert minutes to readable format
    // function getTimeoutLabel(minutes: number): string {
    //     if (minutes === 0) return 'Never';
    //     if (minutes < 60) return `${minutes} minutes`;
    //     if (minutes < 1440) return `${Math.floor(minutes / 60)} hours`;
    //     return `${Math.floor(minutes / 1440)} days`;
    // }

    // // Helper function to get timeout options
    // function getTimeoutOptions() {
    //     return [
    //         { value: 0, label: 'Never' },
    //         { value: 60, label: '1 hour' },
    //         { value: 240, label: '4 hours' },
    //         { value: 480, label: '8 hours' },
    //         { value: 1440, label: '24 hours' },
    //         { value: 10080, label: '7 days' }
    //     ];
    // }

    // // Clear messages after 5 seconds
    // $effect(() => {
    //     if (error || success) {
    //         const timer = setTimeout(() => {
    //             error = '';
    //             success = '';
    //         }, 5000);
            
    //         return () => clearTimeout(timer);
    //     }
    // });
</script>

<svelte:head>
    <title>Security Settings - FormApp</title>
</svelte:head>

<!-- <div class="max-w-4xl mx-auto space-y-6">
    <div>
        <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
            Security Settings
        </h1>
        <p class="text-gray-600 dark:text-gray-400 mt-1">
            Manage your account security and authentication settings
        </p>
    </div>

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

    {#if loading && !securitySettings}
        <div class="flex items-center justify-center py-12">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>
    {:else}
        <div class="space-y-6">
            <Card>
                <div class="flex items-center justify-between mb-6">
                    <div>
                        <h2 class="text-lg font-medium text-gray-900 dark:text-white">
                            Two-Factor Authentication
                        </h2>
                        <p class="text-sm text-gray-600 dark:text-gray-400">
                            Add an extra layer of security to your account
                        </p>
                    </div>
                </div>

                <div class="space-y-4">
                    {#if securitySettings?.two_factor_enabled}
                        <div class="bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 rounded-lg p-4">
                            <div class="flex items-center">
                                <Icon name="shield-check" size="5" classes="text-green-600 dark:text-green-400 mr-3" />
                                <div>
                                    <p class="text-green-800 dark:text-green-400 font-medium">
                                        Two-factor authentication is enabled
                                    </p>
                                    <p class="text-sm text-green-700 dark:text-green-300">
                                        Your account is protected with 2FA
                                    </p>
                                </div>
                            </div>
                        </div>

                        <div class="flex space-x-3">
                            <Button 
                                variant="outline"
                                onclick={() => updateSecuritySettings({ two_factor_enabled: false })}
                                disabled={loading}
                            >
                                Disable 2FA
                            </Button>
                            <Button 
                                variant="ghost"
                                onclick={() => {/* TODO: Show backup codes */}}
                            >
                                View Backup Codes
                            </Button>
                        </div>
                    {:else}
                        <div class="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-4">
                            <div class="flex items-center">
                                <Icon name="shield-exclamation" size="5" classes="text-yellow-600 dark:text-yellow-400 mr-3" />
                                <div>
                                    <p class="text-yellow-800 dark:text-yellow-400 font-medium">
                                        Two-factor authentication is disabled
                                    </p>
                                    <p class="text-sm text-yellow-700 dark:text-yellow-300">
                                        Enable 2FA to better protect your account
                                    </p>
                                </div>
                            </div>
                        </div>

                        <Button 
                            onclick={() => updateSecuritySettings({ two_factor_enabled: true })}
                            disabled={loading}
                            classes="flex items-center gap-2"
                        >
                            <Icon name="shield-plus" size="4" />
                            Enable Two-Factor Authentication
                        </Button>
                    {/if}
                </div>
            </Card>

            <Card>
                <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-6">
                    Login Security
                </h2>

                <div class="space-y-4">
                    <div class="flex items-center justify-between">
                        <div>
                            <h3 class="text-sm font-medium text-gray-900 dark:text-white">
                                Email Notifications
                            </h3>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Get notified when someone logs into your account
                            </p>
                        </div>
                        <Toggle
                            checked={securitySettings?.email_notifications ?? true}
                            onchange={(value) => updateSecuritySettings({ 
                                email_notifications: value 
                            })}
                            disabled={loading}
                        />
                    </div>

                    <div class="flex items-center justify-between">
                        <div>
                            <h3 class="text-sm font-medium text-gray-900 dark:text-white">
                                Multiple Sessions
                            </h3>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Allow multiple active sessions from different devices
                            </p>
                        </div>
                        <Toggle
                            checked={securitySettings?.allow_multiple_sessions ?? true}
                            onchange={(value) => updateSecuritySettings({ 
                                allow_multiple_sessions: value 
                            })}
                            disabled={loading}
                        />
                    </div>

                    <div class="flex items-center justify-between">
                        <div>
                            <h3 class="text-sm font-medium text-gray-900 dark:text-white">
                                Session Timeout
                            </h3>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Automatically log out after period of inactivity
                            </p>
                        </div>

                    </div>
                </div>
            </Card>

            <Card>
                <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-6">
                    Security Actions
                </h2>

                <div class="space-y-4">
                    <div class="flex items-center justify-between">
                        <div>
                            <h3 class="text-sm font-medium text-gray-900 dark:text-white">
                                Password Required for Sensitive Actions
                            </h3>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Require password confirmation for account changes and sensitive operations
                            </p>
                        </div>
                        <Toggle
                            checked={securitySettings?.require_password_for_actions ?? false}
                            onchange={(value) => updateSecuritySettings({ 
                                require_password_for_actions: value 
                            })}
                            disabled={loading}
                        />
                    </div>
                </div>
            </Card>

            <Card>
                <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-6">
                    Account Monitoring
                </h2>

                <div class="space-y-4">
                    <div class="flex items-center justify-between">
                        <div>
                            <h3 class="text-sm font-medium text-gray-900 dark:text-white">
                                Security Alerts
                            </h3>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Get alerts for unusual account activity and security events
                            </p>
                        </div>
                        <Toggle
                            checked={securitySettings?.security_alerts ?? true}
                            onchange={(value) => updateSecuritySettings({ 
                                security_alerts: value 
                            })}
                            disabled={loading}
                        />
                    </div>
                </div>
            </Card>

            <Card>
                <h2 class="text-lg font-medium text-gray-900 dark:text-white mb-6">
                    Current Security Status
                </h2>

                <div class="space-y-4">
                    <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4">
                        <h3 class="text-sm font-medium text-gray-900 dark:text-white mb-3">
                            Security Overview
                        </h3>
                        <div class="grid grid-cols-2 gap-4 text-sm">
                            <div class="flex items-center justify-between">
                                <span class="text-gray-600 dark:text-gray-400">Two-Factor Auth:</span>
                                <Badge color={securitySettings?.two_factor_enabled ? 'success' : 'error'} size="sm">
                                    {securitySettings?.two_factor_enabled ? 'Enabled' : 'Disabled'}
                                </Badge>
                            </div>
                            <div class="flex items-center justify-between">
                                <span class="text-gray-600 dark:text-gray-400">Email Notifications:</span>
                                <Badge color={securitySettings?.email_notifications ? 'success' : 'warning'} size="sm">
                                    {securitySettings?.email_notifications ? 'On' : 'Off'}
                                </Badge>
                            </div>
                            <div class="flex items-center justify-between">
                                <span class="text-gray-600 dark:text-gray-400">Security Alerts:</span>
                                <Badge color={securitySettings?.security_alerts ? 'success' : 'warning'} size="sm">
                                    {securitySettings?.security_alerts ? 'On' : 'Off'}
                                </Badge>
                            </div>
                            <div class="flex items-center justify-between">
                                <span class="text-gray-600 dark:text-gray-400">Session Timeout:</span>
                                <span class="text-gray-900 dark:text-white font-medium">
                                    {getTimeoutLabel(securitySettings?.session_timeout_minutes ?? 1440)}
                                </span>
                            </div>
                        </div>
                    </div>

                    <div class="bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 rounded-lg p-4">
                        <h4 class="text-blue-800 dark:text-blue-400 font-medium mb-2">
                            Security Recommendations
                        </h4>
                        <ul class="text-sm text-blue-700 dark:text-blue-300 space-y-1">
                            {#if !securitySettings?.two_factor_enabled}
                                <li class="flex items-center">
                                    <Icon name="arrow-right" size="4" classes="mr-2" />
                                    Enable two-factor authentication for enhanced security
                                </li>
                            {/if}
                            {#if !securitySettings?.security_alerts}
                                <li class="flex items-center">
                                    <Icon name="arrow-right" size="4" classes="mr-2" />
                                    Turn on security alerts to monitor account activity
                                </li>
                            {/if}
                            {#if securitySettings?.session_timeout_minutes === 0}
                                <li class="flex items-center">
                                    <Icon name="arrow-right" size="4" classes="mr-2" />
                                    Consider setting a session timeout for better security
                                </li>
                            {/if}
                            {#if !securitySettings?.require_password_for_actions}
                                <li class="flex items-center">
                                    <Icon name="arrow-right" size="4" classes="mr-2" />
                                    Require password for sensitive actions
                                </li>
                            {/if}
                        </ul>
                    </div>
                </div>
            </Card>
        </div>
    {/if}
</div> -->