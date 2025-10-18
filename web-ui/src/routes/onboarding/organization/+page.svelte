<script lang="ts">
    // import { onboardingStore } from '$lib/stores/onboardingStore.js';
    // import { goto } from '$app/navigation';
    // import type { CreateOrganizationData } from '$lib/types/onboarding.js';
    // import { OnboardingStep } from '$lib/types/onboarding.js';

    // const { loading, error, clearError } = onboardingStore;
    
    // let orgData = $state<CreateOrganizationData>({
    //     name: '',
    //     slug: '',
    //     description: ''
    // });
    
    // let formErrors = $state<Record<string, string>>({});
    
    // // Generate slug from name
    // function generateSlug(name: string): string {
    //     return name
    //         .toLowerCase()
    //         .replace(/[^a-z0-9]+/g, '-') // Replace non-alphanumeric chars with hyphens
    //         .replace(/^-+|-+$/g, '') // Remove leading/trailing hyphens
    //         .substring(0, 50); // Limit length
    // }
    
    // // Update slug when name changes
    // $effect(() => {
    //     if (orgData.name) {
    //         orgData.slug = generateSlug(orgData.name);
    //     }
    // });
    
    // async function handleSubmit() {
    //     formErrors = {};
    //     clearError();
        
    //     // Basic validation
    //     if (!orgData.name?.trim()) {
    //         formErrors.name = 'Organization name is required';
    //         return;
    //     }
        
    //     if (!orgData.slug?.trim()) {
    //         formErrors.name = 'Invalid organization name. Please use alphanumeric characters.';
    //         return;
    //     }
        
    //     try {
    //         const response = await onboardingStore.createOrganization(orgData);
    //         if (response?.data) {
    //             goto('/onboarding/plans');
    //         } else {
    //             formErrors.name = 'Failed to create organization. Please try again.';
    //         }
    //     } catch (error: any) {
    //         console.error('Failed to create organization:', error);
    //         formErrors.name = error?.message || 'Failed to create organization. Please try again.';
    //     }
    // }
    
    // function handleBack() {
    //     goto('/onboarding');
    // }
</script>

<svelte:head>
    <title>Create Organization - FormApp Onboarding</title>
</svelte:head>

<!-- <div class="max-w-2xl mx-auto px-4 sm:px-6 py-8">
    <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">Create Your Organization</h1>
        <p class="text-lg text-gray-600 dark:text-gray-400">Set up your workspace to start collecting data</p>
    </div>
    
    <form 
        onsubmit={(e) => {
            e.preventDefault();
            handleSubmit();
        }} 
        class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 sm:p-8"
    >
        <div class="space-y-6">
            <div>
                <label 
                    for="orgName" 
                    class="block font-medium text-gray-700 dark:text-gray-300 mb-2"
                >
                    Organization Name *
                </label>
                <input
                    id="orgName"
                    type="text"
                    bind:value={orgData.name}
                    placeholder="Enter your organization name"
                    class="w-full px-4 py-2 rounded-lg border {formErrors.name ? 'border-red-300 dark:border-red-700' : 'border-gray-300 dark:border-gray-600'} 
                           bg-white dark:bg-gray-700 text-gray-900 dark:text-white
                           focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent"
                    disabled={$loading.isLoading}
                    required
                />
                {#if formErrors.name}
                    <p class="mt-2 text-sm text-red-600 dark:text-red-400">{formErrors.name}</p>
                {/if}
            </div>
            
            <div>
                <label 
                    for="orgDescription" 
                    class="block font-medium text-gray-700 dark:text-gray-300 mb-2"
                >
                    Description (Optional)
                </label>
                <textarea
                    id="orgDescription"
                    bind:value={orgData.description}
                    placeholder="Brief description of your organization"
                    rows="3"
                    class="w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600
                           bg-white dark:bg-gray-700 text-gray-900 dark:text-white
                           focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent"
                    disabled={$loading.isLoading}
                ></textarea>
            </div>
            
            <div class="flex gap-4 pt-6 border-t border-gray-200 dark:border-gray-700 sm:flex-row flex-col-reverse">
                <button 
                    type="button" 
                    onclick={handleBack} 
                    class="w-full sm:w-auto px-6 py-2 rounded-lg border border-gray-300 dark:border-gray-600
                           text-gray-700 dark:text-gray-300 font-medium
                           hover:bg-gray-50 dark:hover:bg-gray-700
                           disabled:opacity-50 disabled:cursor-not-allowed
                           transition-colors duration-200"
                    disabled={$loading.isLoading}
                >
                    Back
                </button>
                <button 
                    type="submit" 
                    disabled={$loading.isLoading}
                    class="w-full sm:w-auto px-6 py-2 rounded-lg bg-blue-600 text-white font-medium
                           hover:bg-blue-700 
                           disabled:opacity-50 disabled:cursor-not-allowed
                           transition-colors duration-200"
                >
                    {$loading.isLoading ? 'Creating...' : 'Create Organization'}
                </button>
            </div>
        </div>
    </form>
    
    {#if $error}
        <div class="mt-8 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4" role="alert">
            <div class="flex items-center gap-4">
                <strong class="text-red-700 dark:text-red-400">Error:</strong>
                <p class="text-red-600 dark:text-red-300">{$error.message}</p>
                {#if $error.retryable}
                    <button 
                        onclick={clearError} 
                        class="text-red-600 dark:text-red-400 underline hover:text-red-800 dark:hover:text-red-300"
                    >
                        Try Again
                    </button>
                {/if}
            </div>
        </div>
    {/if}
</div> -->