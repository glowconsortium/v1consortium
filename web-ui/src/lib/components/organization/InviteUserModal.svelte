<script lang="ts">
    import { organizationsStore } from '$lib/stores/organization.js';
    import Modal from '$lib/components/ui/Modal.svelte';
    import Button from '$lib/components/ui/Button.svelte';
    import Input from '$lib/components/ui/Input.svelte';
    import Select from '$lib/components/ui/Select.svelte';
    import Checkbox from '$lib/components/ui/Checkbox.svelte';
    import { toaster } from '$lib/utils/toaster';

    let {organizationId, cancel, success} : {organizationId: string, cancel: () => void, success: () => void} = $props();

    // Form state
    let email = $state('');
    let role = $state('viewer');
    let sendEmail = $state(true);
    let customMessage = $state('');
    let permissions = $state<string[]>([]);
    let error = $state('');

    // Subscribe to store loading state
    let loading = $state(false);
    
    organizationsStore.subscribe(state => {
        loading = state.invitationsLoading;
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
        { value: 'admin', label: 'Admin', description: 'User management and full resource access' }
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

    // Form validation
    const isValid = $derived(email && email.includes('@') && role);

    async function handleSubmit() {
        if (!isValid || loading) return;

        try {
            error = '';

            const inviteData = {
                email: email.trim(),
                role,
                send_email: sendEmail,
                custom_message: customMessage.trim() || undefined,
                permissions: permissions.length > 0 ? permissions : undefined
            };

            await organizationsStore.inviteUser(organizationId, inviteData);
            success();
        } catch (err: any) {
            error = err.message || 'Failed to send invitation';
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
</script>

<Modal title="Invite User" size="lg" oncancel={handleCancel}>
    {#snippet content()}
        
    <form onsubmit={handleSubmit} class="space-y-6">
        <!-- Error Toast -->
        {#if error}
            {toaster.create({
                type: 'error',
                title: error,
                duration: 5000,
            })}
        {/if}

        <!-- Email Input -->
        <div>
            <label for="email" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                Email Address *
            </label>
            <Input
                id="email"
                type="email"
                bind:value={email}
                placeholder="user@example.com"
                required
                disabled={loading}
            />
        </div>

        <!-- Role Selection -->
        <div>
            <label for="role" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                Role *
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
            </div>
        {/if}

        <!-- Email Options -->
        <div class="space-y-3">
            <Checkbox
                id="sendEmail"
                bind:checked={sendEmail}
                label="Send invitation email"
                disabled={loading}
            />

            {#if sendEmail}
                <div>
                    <label for="customMessage" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                        Custom Message (Optional)
                    </label>
                    <textarea
                        id="customMessage"
                        bind:value={customMessage}
                        placeholder="Add a personal message to the invitation email..."
                        rows="3"
                        class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent bg-white dark:bg-gray-700 text-gray-900 dark:text-white resize-none"
                        disabled={loading}
                    ></textarea>
                </div>
            {/if}
        </div>

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
                disabled={!isValid || loading}
                loading={loading}
            >
                {#if loading}
                    Sending Invitation...
                {:else}
                    Send Invitation
                {/if}
            </Button>
        </div>
    </form>
        {/snippet}

</Modal>