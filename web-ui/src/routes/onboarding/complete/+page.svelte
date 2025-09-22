<script lang="ts">
    import { onMount } from 'svelte';
    import { onboardingStore } from '$lib/stores/onboardingStore.js';
    import { goto } from '$app/navigation';

    let countdown = $state(5);
    let autoRedirect = $state(true);
    
    onMount(() => {
        const interval = setInterval(() => {
            if (autoRedirect && countdown > 0) {
                countdown--;
            } else if (autoRedirect && countdown === 0) {
                goto('/dashboard');
            }
        }, 1000);
        
        return () => clearInterval(interval);
    });
    
    function goToDashboard() {
        goto('/dashboard');
    }
    

    
    function cancelAutoRedirect() {
        autoRedirect = false;
    }
</script>

<div class="max-w-2xl mx-auto px-4 py-8 text-center">
    <div class="space-y-8">
        <div class="w-20 h-20 mx-auto text-emerald-500 dark:text-emerald-400 animate-[scale_0.5s_ease-in-out]">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="w-full h-full">
                <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                <polyline points="22,4 12,14.01 9,11.01"></polyline>
            </svg>
        </div>
        
        <div>
            <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">Welcome to FormApp! 🎉</h1>
            <p class="text-lg text-gray-600 dark:text-gray-300">
                Your onboarding is complete! You're all set to start collecting data with your new endpoint.
            </p>
        </div>
        
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
            <div class="bg-white dark:bg-gray-800 rounded-xl p-6 shadow-sm">
                <h3 class="font-semibold text-gray-900 dark:text-white mb-2">Organization Created</h3>
                <p class="text-gray-600 dark:text-gray-400">Your workspace is ready</p>
            </div>
            <div class="bg-white dark:bg-gray-800 rounded-xl p-6 shadow-sm">
                <h3 class="font-semibold text-gray-900 dark:text-white mb-2">Plan Selected</h3>
                <p class="text-gray-600 dark:text-gray-400">Subscription configured</p>
            </div>
            <div class="bg-white dark:bg-gray-800 rounded-xl p-6 shadow-sm">
                <h3 class="font-semibold text-gray-900 dark:text-white mb-2">First Endpoint</h3>
                <p class="text-gray-600 dark:text-gray-400">Ready to receive data</p>
            </div>
        </div>
        
        <div class="bg-blue-50 dark:bg-blue-900/20 rounded-xl p-6 text-left">
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">What's Next?</h2>
            <ul class="space-y-3 text-gray-600 dark:text-gray-300">
                <li class="flex items-center gap-2">
                    <svg class="w-5 h-5 text-blue-500 dark:text-blue-400 flex-shrink-0" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                    </svg>
                    Test your endpoint with a sample submission
                </li>
                <li class="flex items-center gap-2">
                    <svg class="w-5 h-5 text-blue-500 dark:text-blue-400 flex-shrink-0" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                    </svg>
                    Set up integrations (webhooks, email notifications)
                </li>
                <li class="flex items-center gap-2">
                    <svg class="w-5 h-5 text-blue-500 dark:text-blue-400 flex-shrink-0" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                    </svg>
                    Customize your data collection settings
                </li>
                <li class="flex items-center gap-2">
                    <svg class="w-5 h-5 text-blue-500 dark:text-blue-400 flex-shrink-0" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                    </svg>
                    Invite team members to collaborate
                </li>
            </ul>
        </div>
        
        <div class="flex flex-col sm:flex-row gap-4 justify-center">

            <button 
                onclick={goToDashboard}
                class="px-6 py-2 rounded-lg bg-blue-600 text-white hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 
                       transition-colors duration-200"
            >
                Go to Dashboard
            </button>
        </div>
        
        {#if autoRedirect}
            <div class="bg-amber-50 dark:bg-amber-900/20 rounded-lg p-4">
                <p class="text-amber-800 dark:text-amber-200 mb-2">
                    Automatically redirecting to dashboard in {countdown} seconds...
                </p>
                <button 
                    onclick={cancelAutoRedirect}
                    class="text-amber-600 dark:text-amber-400 underline hover:no-underline text-sm"
                >
                    Cancel auto-redirect
                </button>
            </div>
        {/if}
    </div>
</div>

<style>
    @keyframes scale {
        0% { transform: scale(0.5); opacity: 0; }
        100% { transform: scale(1); opacity: 1; }
    }
</style>