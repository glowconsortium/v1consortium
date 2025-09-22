<script lang="ts">
    import { onMount } from 'svelte';
    import { onboardingStore } from '$lib/stores/onboardingStore.js';
    import { goto } from '$app/navigation';
    import OnboardingProgress from '$lib/components/onboarding/OnboardingProgress.svelte';
    import { OnboardingStep } from '$lib/types/onboarding.js';
	import { authStore } from '$lib/stores/auth.js';

    let currentState = $state($onboardingStore);
    
    onMount(async () => {
        await onboardingStore.loadOnboardingState();

		if ($authStore.isInitialized) {
			if (!$authStore.isAuthenticated) {
				goto('/auth/signin');
			} else if (!$onboardingStore.isCompleted) {
				goto('/onboarding');
			}
		}
        
        // Redirect if already completed
        if (currentState?.isCompleted) {
            goto('/dashboard');
            return;
        }
    });
    
    $effect(() => {
        const unsubscribe = onboardingStore.subscribe((state) => {
            currentState = state;
        });
        
        return unsubscribe;
    });
</script>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900 px-4 sm:px-8 py-8 sm:py-12">
    <div class="max-w-7xl mx-auto">
        <OnboardingProgress />
        
        <div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-4 sm:p-8">
            <slot />
        </div>
    </div>
</div> 