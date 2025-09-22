<script lang="ts">
    import { onMount } from 'svelte';
    import { authService } from '$lib/api/auth.js';
    import { userStore } from '$lib/stores/user.js';
    import Button from '$lib/components/ui/Button.svelte';
    import Input from '$lib/components/ui/Input.svelte';
    import Card from '$lib/components/ui/Card.svelte';
    import Icon from '$lib/components/ui/Icon.svelte';
    import Badge from '$lib/components/ui/Badge.svelte';
	import { toaster } from '$lib/utils/toaster';

    // Runes for reactive state
    let loading = $state(false);
    let error = $state('');
    let success = $state('');
    
    // Use userStore instead of local user state
    const user = $derived($userStore.profile);

    // Profile form data
    let profileForm = $state({
        name: '',
        email: '',
        company: '',
        timezone: '',
        avatar_url: ''
    });

    // Email change form
    let emailForm = $state({
        new_email: '',
        password: ''
    });

    // Password change form
    let passwordForm = $state({
        current_password: '',
        new_password: '',
        confirm_password: ''
    });

    // Form states
    let editingProfile = $state(false);
    let changingEmail = $state(false);
    let changingPassword = $state(false);
    let deletingAccount = $state(false);

    // Delete account confirmation
    let deleteForm = $state({
        password: '',
        confirm: '',
        confirmText: ''
    });

    // Available timezones
    const timezones = [
        'UTC',
        'America/New_York',
        'America/Chicago',
        'America/Denver',
        'America/Los_Angeles',
        'Europe/London',
        'Europe/Paris',
        'Europe/Berlin',
        'Asia/Tokyo',
        'Asia/Shanghai',
        'Australia/Sydney'
    ];

    // Computed values using runes
    const isValidEmail = $derived(
        emailForm.new_email && /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(emailForm.new_email)
    );

    const passwordsMatch = $derived(
        passwordForm.new_password === passwordForm.confirm_password
    );

    const isValidPassword = $derived(
        passwordForm.new_password && passwordForm.new_password.length >= 8
    );

    const canChangePassword = $derived(
        passwordForm.current_password && 
        isValidPassword && 
        passwordsMatch && 
        !loading
    );

    const canChangeEmail = $derived(
        isValidEmail && 
        emailForm.password && 
        !loading
    );

    const canDeleteAccount = $derived(
        deleteForm.password && 
        deleteForm.confirm === deleteForm.password && 
        deleteForm.confirmText === 'DELETE' && 
        !loading
    );

    const hasProfileChanges = $derived(() => {
        if (!user) return false;
        return (
            profileForm.name !== (user.name || '') ||
            profileForm.company !== (user.company || '') ||
            profileForm.timezone !== (user.timezone || '') ||
            profileForm.avatar_url !== (user.profile_picture || '')
        );
    });

    onMount(async () => {
        await loadUserProfile();
    });

    async function loadUserProfile() {
        try {
            loading = true;
            await userStore.loadProfile();
            const currentUser = $userStore.profile;
            
            if (currentUser) {
                profileForm = {
                    name: currentUser.name || '',
                    email: currentUser.email || '',
                    company: currentUser.company || '',
                    timezone: currentUser.timezone || 'UTC',
                    avatar_url: currentUser.profile_picture || ''
                };
            }
        } catch (err: any) {
            error = err.message || 'Failed to load profile';
        } finally {
            loading = false;
        }
    }

    async function handleUpdateProfile() {
        try {
            loading = true;
            error = '';

            const updateData = {
                name: profileForm.name.trim() || undefined,
                company: profileForm.company.trim() || undefined,
                timezone: profileForm.timezone || undefined,
                avatar_url: profileForm.avatar_url.trim() || undefined
            };

            await userStore.updateProfile(updateData);
            success = 'Profile updated successfully';
            editingProfile = false;
        } catch (err: any) {
            error = err.message || 'Failed to update profile';
        } finally {
            loading = false;
        }
    }

    async function handleChangeEmail() {
        try {
            loading = true;
            error = '';

            await authService.changeEmail({
                new_email: emailForm.new_email.trim(),
                password: emailForm.password
            });

            success = 'Email change request sent. Please check your email to verify the new address.';
            changingEmail = false;
            emailForm = { new_email: '', password: '' };
        } catch (err: any) {
            error = err.message || 'Failed to change email';
        } finally {
            loading = false;
        }
    }

    async function handleChangePassword() {
        try {
            loading = true;
            error = '';

            await authService.changePassword({
                current_password: passwordForm.current_password,
                new_password: passwordForm.new_password
            });

            success = 'Password changed successfully';
            changingPassword = false;
            passwordForm = {
                current_password: '',
                new_password: '',
                confirm_password: ''
            };
        } catch (err: any) {
            error = err.message || 'Failed to change password';
        } finally {
            loading = false;
        }
    }

    async function handleDeleteAccount() {
        try {
            loading = true;
            error = '';

            await authService.deleteProfile(deleteForm.password, deleteForm.confirm);
            
            // Redirect to login or home page after successful deletion
            window.location.href = '/';
        } catch (err: any) {
            error = err.message || 'Failed to delete account';
        } finally {
            loading = false;
        }
    }

    function cancelProfileEdit() {
        editingProfile = false;
        if (user) {
            profileForm = {
                name: user.name || '',
                email: user.email || '',
                company: user.company || '',
                timezone: user.timezone || 'UTC',
                avatar_url: user.profile_picture || ''
            };
        }
    }

    function cancelEmailChange() {
        changingEmail = false;
        emailForm = { new_email: '', password: '' };
    }

    function cancelPasswordChange() {
        changingPassword = false;
        passwordForm = {
            current_password: '',
            new_password: '',
            confirm_password: ''
        };
    }

    function cancelAccountDeletion() {
        deletingAccount = false;
        deleteForm = {
            password: '',
            confirm: '',
            confirmText: ''
        };
    }

    async function handleResendVerification() {
        try {
            loading = true;
            await authService.resendVerification({ email: user?.email || '' });
            success = 'Verification email sent';
        } catch (err: any) {
            error = 'Failed to send verification email';
        } finally {
            loading = false;
        }
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

    // Sync form data when user changes
    $effect(() => {
        if (user && !editingProfile) {
            profileForm = {
                name: user.name || '',
                email: user.email || '',
                company: user.company || '',
                timezone: user.timezone || 'UTC',
                avatar_url: user.profile_picture || ''
            };
        }
    });
</script>

<svelte:head>
    <title>Profile - FormApp</title>
</svelte:head>

<div class="max-w-4xl mx-auto space-y-6">
    <!-- Header -->
    <div>
        <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
            Profile Settings
        </h1>
        <p class="text-gray-600 dark:text-gray-400 mt-1">
            Manage your account settings and preferences
        </p>
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
    {#if loading && !user}
        <div class="flex items-center justify-center py-12">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>
    {:else}
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <!-- Left Column - Profile Info -->
            <div class="lg:col-span-2 space-y-6">
                <!-- Basic Profile Information -->
                <Card>
                    <div class="flex items-center justify-between mb-6">
                        <div>
                            <h2 class="text-lg font-medium text-gray-900 dark:text-white">
                                Profile Information
                            </h2>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Update your account's profile information
                            </p>
                        </div>
                        
                        {#if !editingProfile}
                            <Button 
                                variant="ghost" 
                                onclick={() => editingProfile = true}
                                classes="flex items-center gap-2"
                            >
                                <Icon name="edit" size="4" />
                                Edit
                            </Button>
                        {/if}
                    </div>

                    {#if editingProfile}
                        <form onsubmit={handleUpdateProfile} class="space-y-4">
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div>
                                    <label for="name" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                        Full Name
                                    </label>
                                    <Input
                                        id="name"
                                        bind:value={profileForm.name}
                                        placeholder="Enter your full name"
                                        disabled={loading}
                                    />
                                </div>

                                <div>
                                    <label for="company" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                        Company
                                    </label>
                                    <Input
                                        id="company"
                                        bind:value={profileForm.company}
                                        placeholder="Enter your company name"
                                        disabled={loading}
                                    />
                                </div>
                            </div>

                            <div>
                                <label for="timezone" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Timezone
                                </label>
                                <select
                                    id="timezone"
                                    bind:value={profileForm.timezone}
                                    disabled={loading}
                                    class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                >
                                    {#each timezones as tz}
                                        <option value={tz}>{tz}</option>
                                    {/each}
                                </select>
                            </div>

                            <div>
                                <label for="avatar_url" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Avatar URL
                                </label>
                                <Input
                                    id="avatar_url"
                                    type="url"
                                    bind:value={profileForm.avatar_url}
                                    placeholder="https://example.com/avatar.jpg"
                                    disabled={loading}
                                />
                            </div>

                            <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200 dark:border-gray-600">
                                <Button
                                    type="button"
                                    variant="ghost"
                                    onclick={cancelProfileEdit}
                                    disabled={loading}
                                >
                                    Cancel
                                </Button>
                                <Button
                                    type="submit"
                                    disabled={!hasProfileChanges || loading}
                                    loading={loading}
                                >
                                    Save Changes
                                </Button>
                            </div>
                        </form>
                    {:else}
                        <div class="space-y-4">
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div>
                                    <label for="name" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                        Full Name
                                    </label>
                                    <p class="text-gray-900 dark:text-white">
                                        {user?.name || 'Not provided'}
                                    </p>
                                </div>

                                <div>
                                    <label for="email" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                        Email
                                    </label>
                                    <div class="flex items-center gap-2">
                                        <p class="text-gray-900 dark:text-white">
                                            {user?.email}
                                        </p>
                                        {#if user?.email_verified}
                                            <Badge color="success" size="sm">Verified</Badge>
                                        {:else}
                                            <Badge color="warning" size="sm">Unverified</Badge>
                                        {/if}
                                    </div>
                                </div>
                            </div>

                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div>
                                    <label for="company" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                        Company
                                    </label>
                                    <p class="text-gray-900 dark:text-white">
                                        {user?.company || 'Not provided'}
                                    </p>
                                </div>

                                <div>
                                    <label for="timezone" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                        Timezone
                                    </label>
                                    <p class="text-gray-900 dark:text-white">
                                        {user?.timezone || 'UTC'}
                                    </p>
                                </div>
                            </div>

                            <div>
                                <label for="member_since" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Member Since
                                </label>
                                <p class="text-gray-900 dark:text-white">
                                    {user?.created_at ? new Date(user.created_at).toLocaleDateString() : 'Unknown'}
                                </p>
                            </div>
                        </div>
                    {/if}
                </Card>

                <!-- Change Email -->
                <Card>
                    <div class="flex items-center justify-between mb-6">
                        <div>
                            <h2 class="text-lg font-medium text-gray-900 dark:text-white">
                                Email Address
                            </h2>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Change your account's email address
                            </p>
                        </div>
                        
                        {#if !changingEmail}
                            <Button 
                                variant="ghost" 
                                onclick={() => changingEmail = true}
                                classes="flex items-center gap-2"
                            >
                                <Icon name="mail" size="4" />
                                Change Email
                            </Button>
                        {/if}
                    </div>

                    {#if changingEmail}
                        <form onsubmit={handleChangeEmail} class="space-y-4">
                            <div>
                                <label for="new_email" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    New Email Address
                                </label>
                                <Input
                                    id="new_email"
                                    type="email"
                                    bind:value={emailForm.new_email}
                                    placeholder="Enter new email address"
                                    disabled={loading}
                                    required
                                />
                            </div>

                            <div>
                                <label for="email_password" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Current Password
                                </label>
                                <Input
                                    id="email_password"
                                    type="password"
                                    bind:value={emailForm.password}
                                    placeholder="Enter your current password"
                                    disabled={loading}
                                    required
                                />
                            </div>

                            <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200 dark:border-gray-600">
                                <Button
                                    type="button"
                                    variant="ghost"
                                    onclick={cancelEmailChange}
                                    disabled={loading}
                                >
                                    Cancel
                                </Button>
                                <Button
                                    type="submit"
                                    disabled={!canChangeEmail}
                                    loading={loading}
                                >
                                    Change Email
                                </Button>
                            </div>
                        </form>
                    {:else}
                        <div class="flex items-center justify-between p-4 bg-gray-50 dark:bg-gray-700 rounded-lg">
                            <div>
                                <p class="text-gray-900 dark:text-white font-medium">
                                    Current: {user?.email}
                                </p>
                                <p class="text-sm text-gray-600 dark:text-gray-400">
                                    Click "Change Email" to update your email address
                                </p>
                            </div>
                        </div>
                    {/if}
                </Card>

                <!-- Change Password -->
                <Card>
                    <div class="flex items-center justify-between mb-6">
                        <div>
                            <h2 class="text-lg font-medium text-gray-900 dark:text-white">
                                Password
                            </h2>
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Update your account password
                            </p>
                        </div>
                        
                        {#if !changingPassword}
                            <Button 
                                variant="ghost" 
                                onclick={() => changingPassword = true}
                                classes="flex items-center gap-2"
                            >
                                <Icon name="lock" size="4" />
                                Change Password
                            </Button>
                        {/if}
                    </div>

                    {#if changingPassword}
                        <form onsubmit={handleChangePassword} class="space-y-4">
                            <div>
                                <label for="current_password" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Current Password
                                </label>
                                <Input
                                    id="current_password"
                                    type="password"
                                    bind:value={passwordForm.current_password}
                                    placeholder="Enter your current password"
                                    disabled={loading}
                                    required
                                />
                            </div>

                            <div>
                                <label for="new_password" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    New Password
                                </label>
                                <Input
                                    id="new_password"
                                    type="password"
                                    bind:value={passwordForm.new_password}
                                    placeholder="Enter new password (min 8 characters)"
                                    disabled={loading}
                                    required
                                />
                                {#if passwordForm.new_password && !isValidPassword}
                                    <p class="text-sm text-red-600 mt-1">Password must be at least 8 characters long</p>
                                {/if}
                            </div>

                            <div>
                                <label for="confirm_password" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Confirm New Password
                                </label>
                                <Input
                                    id="confirm_password"
                                    type="password"
                                    bind:value={passwordForm.confirm_password}
                                    placeholder="Confirm your new password"
                                    disabled={loading}
                                    required
                                />
                                {#if passwordForm.confirm_password && !passwordsMatch}
                                    <p class="text-sm text-red-600 mt-1">Passwords do not match</p>
                                {/if}
                            </div>

                            <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200 dark:border-gray-600">
                                <Button
                                    type="button"
                                    variant="ghost"
                                    onclick={cancelPasswordChange}
                                    disabled={loading}
                                >
                                    Cancel
                                </Button>
                                <Button
                                    type="submit"
                                    disabled={!canChangePassword}
                                    loading={loading}
                                >
                                    Change Password
                                </Button>
                            </div>
                        </form>
                    {:else}
                        <div class="flex items-center justify-between p-4 bg-gray-50 dark:bg-gray-700 rounded-lg">
                            <div>
                                <p class="text-gray-900 dark:text-white font-medium">
                                    Password
                                </p>
                                <p class="text-sm text-gray-600 dark:text-gray-400">
                                </p>
                            </div>
                        </div>
                    {/if}
                </Card>
            </div>

            <!-- Right Column - Account Actions -->
            <div class="space-y-6">
                <!-- Account Status -->
                <Card>
                    <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">
                        Account Status
                    </h3>
                    
                    <div class="space-y-3">
                        <div class="flex items-center justify-between">
                            <span class="text-sm text-gray-600 dark:text-gray-400">Email Verified</span>
                            {#if user?.email_verified}
                                <Badge color="success">Verified</Badge>
                            {:else}
                                <Badge color="error">Pending</Badge>
                            {/if}
                        </div>



                    </div>

                    {#if !user?.email_verified}
                        <div class="mt-4 pt-4 border-t border-gray-200 dark:border-gray-600">
                            <Button 
                                size="sm" 
                                variant="outline"
                                classes="w-full"
                                onclick={handleResendVerification}
                                disabled={loading}
                            >
                                Resend Verification Email
                            </Button>
                        </div>
                    {/if}
                </Card>

                <!-- Quick Actions -->
                <Card>
                    <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">
                        Quick Actions
                    </h3>
                    
                    <div class="space-y-3">
                        <Button 
                            variant="outline" 
                            classes="w-full justify-start"
                            onclick={() => window.location.href = '/dashboard/profile/security'}
                        >
                            <Icon name="shield" size="4" />
                            Security Settings
                        </Button>

                        <Button 
                            variant="outline" 
                            classes="w-full justify-start"
                            onclick={() => window.location.href = '/dashboard/profile/sessions'}
                        >
                            <Icon name="monitor" size="4"  />
                            Active Sessions
                        </Button>

                        <Button 
                            variant="outline" 
                            classes="w-full justify-start"
                            onclick={() => window.location.href = '/dashboard/billing'}
                        >
                            <Icon name="credit-card" size="4" />
                            Billing & Usage
                        </Button>
                    </div>
                </Card>

                <!-- Danger Zone -->
                <Card classes="border-red-200 dark:border-red-800">
                    <h3 class="text-lg font-medium text-red-600 dark:text-red-400 mb-4">
                        Danger Zone
                    </h3>
                    
                    {#if !deletingAccount}
                        <div class="space-y-3">
                            <p class="text-sm text-gray-600 dark:text-gray-400">
                                Once you delete your account, there is no going back. Please be certain.
                            </p>
                            
                            <Button 
                                variant="filled"
                                classes="w-full"
                                onclick={() => deletingAccount = true}
                            >
                                <Icon name="trash" size="4" />
                                Delete Account
                            </Button>
                        </div>
                    {:else}
                        <form onsubmit={handleDeleteAccount} class="space-y-4">
                            <div class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
                                <h4 class="text-red-800 dark:text-red-400 font-medium mb-2">
                                    ⚠️ This action cannot be undone
                                </h4>
                                <p class="text-sm text-red-700 dark:text-red-300">
                                    This will permanently delete your account and all associated data.
                                </p>
                            </div>

                            <div>
                                <label for="delete_password" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Enter your password
                                </label>
                                <Input
                                    id="delete_password"
                                    type="password"
                                    bind:value={deleteForm.password}
                                    placeholder="Enter your password"
                                    disabled={loading}
                                    required
                                />
                            </div>

                            <div>
                                <label for="delete_confirm" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Confirm your password
                                </label>
                                <Input
                                    id="delete_confirm"
                                    type="password"
                                    bind:value={deleteForm.confirm}
                                    placeholder="Confirm your password"
                                    disabled={loading}
                                    required
                                />
                            </div>

                            <div>
                                <label for="delete_text" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                                    Type "DELETE" to confirm
                                </label>
                                <Input
                                    id="delete_text"
                                    bind:value={deleteForm.confirmText}
                                    placeholder="DELETE"
                                    disabled={loading}
                                    required
                                />
                            </div>

                            <div class="flex space-x-3">
                                <Button
                                    type="button"
                                    variant="ghost"
                                    classes="flex-1"
                                    onclick={cancelAccountDeletion}
                                    disabled={loading}
                                >
                                    Cancel
                                </Button>
                                <Button
                                    type="submit"
                                    variant="filled"
                                    classes="flex-1"
                                    disabled={!canDeleteAccount}
                                    loading={loading}
                                >
                                    Delete Forever
                                </Button>
                            </div>
                        </form>
                    {/if}
                </Card>
            </div>
        </div>
    {/if}
</div>