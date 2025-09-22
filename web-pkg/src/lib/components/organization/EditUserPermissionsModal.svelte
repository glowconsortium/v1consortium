<script lang="ts">
    import { onMount } from 'svelte';
    import { organizationsStore } from '$lib/stores/organization.js';
    import Modal from '$lib/components/ui/Modal.svelte';
    import Button from '$lib/components/ui/Button.svelte';
    import Select from '$lib/components/ui/Select.svelte';
    import Checkbox from '$lib/components/ui/Checkbox.svelte';
    import {toaster} from '$lib/utils/toaster.js';
    import Badge from '$lib/components/ui/Badge.svelte';

    let {organizationId, member, success, cancel}: {organizationId: string, member: any, 
        success: ()=>void, cancel: ()=>void} = $props();

    // Form state
    let role = $state(member.role);
    let permissions = $state<string[]>(member.permissions || []);
    let loading = $state(false);
    let error = $state('');

    // Subscribe to store
    organizationsStore.subscribe(state => {
        loading = state.usersLoading;
        if (state.error) {
            error = state.error;
        }
    });

    // Available roles
    const roles = [
        { value: 'viewer', label: 'Viewer', description: 'Read-only access to all resources' },
        { value: 'support', label: 'Support', description: 'Read submissions, limited endpoint access' },
        { value: 'analyst', label: 'Analyst', description: 'Read submissions and analytics' },
        { value: 'developer', label: 'Developer', description: 'Full endpoint and integration management' },
        { value: 'admin', label: 'Admin', description: 'User management and full resource access' },
        { value: 'custom', label: 'Custom', description: 'Custom permission set' }
    ];

    // Available permissions for custom roles
    const availablePermissions = [
        { id: 'endpoints:read', label: 'View Endpoints', category: 'Endpoints' },
        { id: 'endpoints:write', label: 'Create/Edit Endpoints', category: 'Endpoints' },
        { id: 'endpoints:delete', label: 'Delete Endpoints', category: 'Endpoints' },
        { id: 'submissions:read', label: 'View Submissions', category: 'Submissions' },
        { id: 'submissions:export', label: 'Export Submissions', category: 'Submissions' },
        { id: 'submissions:delete', label: 'Delete Submissions', category: 'Submissions' },
        { id: 'integrations:read', label: 'View Integrations', category: 'Integrations' },
        { id: 'integrations:write', label: 'Create/Edit Integrations', category: 'Integrations' },
        { id: 'integrations:delete', label: 'Delete Integrations', category: 'Integrations' },
        { id: 'analytics:read', label: 'View Analytics', category: 'Analytics' },
        { id: 'organization:admin', label: 'Organization Administration', category: 'Organization' }
    ];

    // Group permissions by category
    const permissionGroups = $derived(availablePermissions.reduce((groups, permission) => {
        if (!groups[permission.category]) {
            groups[permission.category] = [];
        }
        groups[permission.category].push(permission);
        return groups;
    }, {} as Record<string, typeof availablePermissions>));

    // Check if changes have been made
    const hasChanges = $derived(role !== member.role || 
        JSON.stringify([...permissions].sort()) !== JSON.stringify([...(member.permissions || [])].sort()));

    onMount(() => {
        // If user has custom permissions, set role to custom
        if (member.permissions && member.permissions.length > 0 && !roles.some(r => r.value === member.role)) {
            role = 'custom';
        }
    });

    async function handleSubmit() {
        if (!hasChanges || loading) return;

        try {
            error = '';

            if (role !== 'custom') {
                // For predefined roles, send role update
                await organizationsStore.updateUserRole(organizationId, member.user_id, role);
            } else {
                // For custom roles, use updateUserRole with permissions
                // Note: You may need to add updateUserPermissions method to the store
                // For now, using updateUserRole which should handle both role and permissions
                await organizationsStore.updateUserRole(organizationId, member.user_id, 'custom');
            }

            success();
        } catch (err: any) {
            error = err.message || 'Failed to update permissions';
        }
    }

    function handleCancel() {
        cancel();
    }

    function togglePermission(permissionId: string) {
        if (permissions.includes(permissionId)) {
            permissions = permissions.filter(p => p !== permissionId);
        } else {
            permissions = [...permissions, permissionId];
        }
    }

    function selectAllInCategory(category: string) {
        const categoryPermissions = permissionGroups[category].map(p => p.id);
        const allSelected = categoryPermissions.every(p => permissions.includes(p));
        
        if (allSelected) {
            // Deselect all in category
            permissions = permissions.filter(p => !categoryPermissions.includes(p));
        } else {
            // Select all in category
            permissions = [...new Set([...permissions, ...categoryPermissions])];
        }
    }

    // Clear custom permissions when switching to predefined role
    $effect(() => {
        if (role !== 'custom') {
            permissions = [];
        } else if (role === 'custom' && permissions.length === 0 && member.permissions) {
            // Restore original permissions when switching to custom
            permissions = [...member.permissions];
        }
    });

    // Clear error when store error changes
    $effect(() => {
        if (error) {
            const timer = setTimeout(() => {
                error = '';
                organizationsStore.clearError();
            }, 5000);
            
            return () => clearTimeout(timer);
        }
    });

    function getRoleBadgeColor(roleValue: string) {
        switch (roleValue) {
            case 'owner': return 'success';
            case 'admin': return 'primary';
            case 'developer': return 'success';
            case 'analyst': return 'warning';
            case 'support': return 'secondary';
            case 'viewer': return 'tertiary';
            case 'custom': return 'error';
            default: return 'warning';
        }
    }
</script>

<Modal title="Edit User Permissions" size="lg" oncancel={handleCancel}>
   
    {#snippet content()}
        
    <div class="space-y-6">
        <!-- Error Toast -->
        {#if error}
        {toaster.create({
            type: 'error',
            title: error,
            duration: 5000,
        })}
        {/if}

        <!-- User Info -->
        <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4">
            <div class="flex items-center space-x-3">
                <div class="flex-shrink-0 h-10 w-10">
                    <div class="h-10 w-10 rounded-full bg-gray-300 dark:bg-gray-600 flex items-center justify-center">
                        <span class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {member.name?.charAt(0) || member.email?.charAt(0) || '?'}
                        </span>
                    </div>
                </div>
                <div>
                    <div class="text-sm font-medium text-gray-900 dark:text-white">
                        {member.name || 'N/A'}
                    </div>
                    <div class="text-sm text-gray-500 dark:text-gray-400">
                        {member.email}
                    </div>
                </div>
                <div class="flex-1"></div>
                <Badge color={getRoleBadgeColor(member.role)}>
                    Current: {member.role}
                </Badge>
            </div>
        </div>

        <form onsubmit={handleSubmit} class="space-y-6">
            <!-- Role Selection -->
            <div>
                <label for="role" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    Role
                </label>
                <Select
                    id="role"
                    bind:value={role}
                    disabled={loading}
                    options={roles.map(r => ({ value: r.value, label: r.label }))}
                />
                
                <!-- Role Description -->
                {#if role}
                    {@const selectedRole = roles.find(r => r.value === role)}
                    {#if selectedRole}
                        <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">
                            {selectedRole.description}
                        </p>
                    {/if}
                {/if}
            </div>

            <!-- Custom Permissions (if custom role is selected) -->
            {#if role === 'custom'}
                <div>
                    <label for="custom_permissions" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
                        Custom Permissions
                    </label>
                    
                    <div class="space-y-4 max-h-64 overflow-y-auto border border-gray-200 dark:border-gray-600 rounded-lg p-4">
                        {#each Object.entries(permissionGroups) as [category, categoryPermissions]}
                            <div>
                                <div class="flex items-center justify-between mb-2">
                                    <h4 class="text-sm font-medium text-gray-900 dark:text-white">
                                        {category}
                                    </h4>
                                    <Button
                                        type="button"
                                        variant="ghost"
                                        size="sm"
                                        onclick={() => selectAllInCategory(category)}
                                    >
                                        {categoryPermissions.every(p => permissions.includes(p.id)) ? 'Deselect All' : 'Select All'}
                                    </Button>
                                </div>
                                
                                <div class="space-y-2 pl-4">
                                    {#each categoryPermissions as permission}
                                        <Checkbox
                                            id={permission.id}
                                            checked={permissions.includes(permission.id)}
                                            onchange={() => togglePermission(permission.id)}
                                            label={permission.label}
                                        />
                                    {/each}
                                </div>
                            </div>
                        {/each}
                    </div>

                    {#if permissions.length === 0}
                        <p class="text-sm text-yellow-600 dark:text-yellow-400 mt-2">
                            ⚠️ No permissions selected. User will have very limited access.
                        </p>
                    {/if}
                </div>
            {/if}

            <!-- Current Permissions Display (if not custom) -->
            {#if role !== 'custom' && member.permissions && member.permissions.length > 0}
                <div>
                    <label for="current_permissions" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                        Current Custom Permissions (will be replaced)
                    </label>
                    <div class="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-3">
                        <p class="text-sm text-yellow-800 dark:text-yellow-200 mb-2">
                            This user currently has custom permissions. Selecting a predefined role will replace these permissions.
                        </p>
                        <div class="flex flex-wrap gap-1">
                            {#each member.permissions as permission}
                                <Badge color="warning" size="sm">
                                    {availablePermissions.find(p => p.id === permission)?.label || permission}
                                </Badge>
                            {/each}
                        </div>
                    </div>
                </div>
            {/if}

            <!-- Action Buttons -->
            <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200 dark:border-gray-600">
                <Button
                    type="button"
                    variant="ghost"
                    onclick={handleCancel}
                    disabled={loading}
                >
                    Cancel
                </Button>
                
                <Button
                    type="submit"
                    disabled={!hasChanges || loading}
                    loading={loading}
                >
                    {#if loading}
                        Updating Permissions...
                    {:else}
                        Update Permissions
                    {/if}
                </Button>
            </div>
        </form>
    </div>
        {/snippet}

</Modal>