<script lang="ts">
    // import { onMount } from 'svelte';
    // import { onboardingStore } from '$lib/stores/onboardingStore.js';
    // import { goto } from '$app/navigation';
    // import type { SubscriptionPlan } from '$lib/types/subscription.js';
    // import type { PlanSelectionData, OnboardingState, InvitedOnboardingState } from '$lib/types/onboarding.js';

    // // Import your UI library components
    // import { Button, Spinner, Alert, Card } from '$lib/components';

    // const { plans, loading, error, clearError } = onboardingStore;
    
    // let selectedPlan = $state<SubscriptionPlan | null>(null);
    // let billingCycle = $state<'monthly' | 'yearly'>('monthly');
    // let currentState = $state<OnboardingState | InvitedOnboardingState | null>(null);
    // let validationError = $state<string | null>(null);
    
    // // Subscribe to store updates
    // $effect(() => {
    //     const unsubscribe = onboardingStore.subscribe((state) => {
    //         currentState = state;
    //     });
    //     return unsubscribe;
    // });
    
    // onMount(async () => {
    //     clearError();
    //     await onboardingStore.loadPlans();
        
    //     // Redirect if no organization data
    //     if (!currentState?.organizationData?.id) {
    //         goto('/onboarding/organization');
    //         return;
    //     }
    // });
    
    // async function validatePlanSelection(planId: string, organizationId: string): Promise<boolean> {
    //     try {
    //         const validation = await onboardingStore.validatePlanSelection(organizationId, planId);
    //         if (!validation.valid) {
    //             validationError = validation.reason || 'Invalid plan selection';
    //             return false;
    //         }
    //         return true;
    //     } catch (error) {
    //         validationError = 'Failed to validate plan selection';
    //         return false;
    //     }
    // }
    
    // async function handlePlanSelection() {
    //     clearError();
    //     validationError = null;
        
    //     if (!selectedPlan || !currentState?.organizationData?.id) {
    //         validationError = 'Please select a plan to continue';
    //         return;
    //     }
        
    //     // Validate plan selection first
    //     const isValid = await validatePlanSelection(selectedPlan.id, currentState.organizationData.id);
    //     if (!isValid) {
    //         return;
    //     }
        
    //     const planSelection: PlanSelectionData = {
    //         plan: selectedPlan,
    //         billing_cycle: billingCycle
    //     };
        
    //     try {
    //         const response = await onboardingStore.selectPlan(planSelection, currentState.organizationData.id);
    //         if (response.success) {
    //             goto('/onboarding/tutorial');
    //         } else {
    //             validationError = response.message || 'Failed to select plan';
    //         }
    //     } catch (error: any) {
    //         console.error('Failed to select plan:', error);
    //         validationError = error?.message || 'Failed to select plan. Please try again.';
    //     }
    // }
    
    // function handleBack() {
    //     goto('/onboarding/organization');
    // }
</script>

<svelte:head>
    <title>Choose Your Plan - FormApp</title>
    <meta name="description" content="Select the plan that best fits your needs" />
</svelte:head>

<!-- <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    {#if $loading.isLoading}
        <div class="text-center py-12">
            <Spinner />
            <p class="mt-4 text-gray-600 dark:text-gray-400">{$loading.operation || 'Loading plans...'}</p>
        </div>
    {:else if $error}
        <div class="text-center py-8">
            <p class="text-red-600 dark:text-red-400 mb-4">{$error.message}</p>
            {#if $error.retryable}
                <Button 
                    variant="soft" 
                    color="primary" 
                    onclick={() => onboardingStore.loadPlans()}
                >
                    Try Again
                </Button>
            {/if}
        </div>
    {:else}
        <div class="space-y-8">
            <div class="text-center">
                <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-2">Choose Your Plan</h1>
                <p class="text-xl text-gray-600 dark:text-gray-400">Select the plan that best fits your needs</p>
            </div>
            
            {#if validationError}
                <div class="max-w-2xl mx-auto">
                    <p class="text-red-600 dark:text-red-400 mb-4">{validationError}</p>
                </div>
            {/if}
            
            <div class="flex justify-center">
                <div class="inline-flex rounded-lg bg-gray-100 dark:bg-gray-800 p-1">
                    <Button 
                        variant={billingCycle === 'monthly' ? 'filled' : 'ghost'}
                        color="primary"
                        size="sm"
                        onclick={() => billingCycle = 'monthly'}
                        disabled={$loading.isLoading}
                        classes="rounded-l-lg"
                    >
                        Monthly
                    </Button>
                    <Button 
                        variant={billingCycle === 'yearly' ? 'filled' : 'ghost'}
                        color="primary"
                        size="sm"
                        onclick={() => billingCycle = 'yearly'}
                        disabled={$loading.isLoading}
                        classes="rounded-r-lg"
                    >
                        Yearly (Save 20%)
                    </Button>
                </div>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                {#each $plans as plan}
                    <Card 
                        variant={selectedPlan?.id === plan.id ? 'filled' : 'outline'}
                        color={selectedPlan?.id === plan.id ? 'primary' : 'surface'}
                        hover={true}
                        padding="lg"
                    >
                        <div class="flex justify-between items-start mb-6">
                            <h3 class="text-2xl font-semibold text-gray-900 dark:text-white">{plan.name}</h3>
                            {#if plan.is_popular}
                                <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200">
                                    Most Popular
                                </span>
                            {/if}
                        </div>
                        
                        <div class="mb-6">
                            <div class="flex items-baseline">
                                <span class="text-2xl font-semibold text-gray-900 dark:text-white">$</span>
                                <span class="text-5xl font-bold text-gray-900 dark:text-white">
                                    {billingCycle === 'monthly' 
                                        ? plan.pricing.monthly / 100 
                                        : plan.pricing.yearly / 1200}
                                </span>
                                <span class="ml-2 text-gray-600 dark:text-gray-400">/month</span>
                            </div>
                            {#if billingCycle === 'yearly'}
                                <span class="inline-block mt-2 text-sm text-green-600 dark:text-green-400 bg-green-50 dark:bg-green-900/20 px-2 py-1 rounded">
                                    Save 20%
                                </span>
                            {/if}
                        </div>
                        
                        <ul class="space-y-4 mb-8">
                            {#each plan.features as feature}
                                <li class="flex items-start gap-3">
                                    <span class="text-green-500 dark:text-green-400 mt-1">
                                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                                        </svg>
                                    </span>
                                    <div>
                                        <span class="text-gray-900 dark:text-white">{feature.name}</span>
                                        {#if feature.description}
                                            <p class="text-sm text-gray-600 dark:text-gray-400">{feature.description}</p>
                                        {/if}
                                    </div>
                                </li>
                            {/each}
                        </ul>
                        
                        <Button 
                            variant={selectedPlan?.id === plan.id ? 'filled' : 'outline'}
                            color="primary"
                            classes="w-full"
                            disabled={$loading.isLoading}
                            onclick={() => {
                                selectedPlan = plan;
                                validationError = null; // Clear any previous validation errors
                            }}
                        >
                            {selectedPlan?.id === plan.id ? 'Selected' : 'Select Plan'}
                        </Button>
                    </Card>
                {/each}
            </div>
            
            <div class="flex justify-center gap-4 pt-8 sm:flex-row flex-col">
                <Button 
                    variant="outline"
                    color="surface"
                    onclick={handleBack}
                    disabled={$loading.isLoading}
                    classes="w-full sm:w-auto"
                >
                    Back
                </Button>
                <Button 
                    variant="filled"
                    color="primary"
                    onclick={handlePlanSelection}
                    disabled={!selectedPlan || $loading.isLoading}
                    classes="w-full sm:w-auto"
                >
                    {$loading.isLoading ? 'Processing...' : 'Continue'}
                </Button>
            </div>
        </div>
    {/if}
</div> -->