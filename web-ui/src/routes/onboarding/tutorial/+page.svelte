<script lang="ts">
    import { onMount } from 'svelte';
    import { onboardingStore } from '$lib/stores/onboardingStore.js';
    import { goto } from '$app/navigation';
    import type { TutorialStep } from '$lib/types/onboarding.js';

    // Example UI components, replace with your actual imports
    import Button from '$lib/components/ui/Button.svelte';
    import Spinner from '$lib/components/ui/Spinner.svelte';
    import Alert from '$lib/components/ui/Alert.svelte';
    import Card from '$lib/components/ui/Card.svelte';

    const { tutorialSteps, loading, error, clearError } = onboardingStore;
    let currentStepIndex = $state(0);
    let currentState = $state($onboardingStore);

    onMount(() => {
        if (!currentState?.organizationData?.id || !currentState?.selectedPlan) {
            goto('/onboarding/plans');
            return;
        }
    });

    $effect(() => {
        const allCompleted = $tutorialSteps.every(step => step.isCompleted);
        if (allCompleted) {
            completeTutorial();
        }
    });

    async function completeTutorial() {
        clearError();
        try {
            await onboardingStore.completeTutorial();
            goto('/onboarding/complete');
        } catch (error: any) {
            console.error('Failed to complete tutorial:', error);
        }
    }

    function nextStep() {
        if (currentStepIndex < $tutorialSteps.length - 1) {
            currentStepIndex++;
        }
    }

    function prevStep() {
        if (currentStepIndex > 0) {
            currentStepIndex--;
        }
    }

    function handleStepAction(step: TutorialStep) {
        if (step.action?.type === 'navigate') {
            goto(`${step.action.target}`);
        }
    }

    function handleBack() {
        goto('/onboarding/plans');
    }
</script>

<svelte:head>
    <title>Quick Tutorial - FormApp Onboarding</title>
</svelte:head>

<div class="max-w-5xl mx-auto py-8">
    <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Quick Tutorial</h1>
        <p class="text-gray-500 text-lg">Learn the basics of FormApp in just a few minutes</p>
    </div>

    {#if $loading.isLoading}
        <div class="flex flex-col items-center py-8">
            <Spinner classes="w-10 h-10 mb-4" />
            <p class="text-gray-500">{$loading.operation || 'Loading tutorial...'}</p>
        </div>
    {:else if $error}
            <div class="flex items-center gap-2">
                <strong>Error:</strong> {$error.message}
                {#if $error.retryable}
                    <Button color="primary" onclick={clearError}>Try Again</Button>
                {/if}
            </div>
    {:else}
        <Card classes="grid grid-cols-1 md:grid-cols-[300px_1fr] gap-6 bg-white rounded-xl shadow overflow-hidden">
            <aside class="bg-gray-50 p-6 border-r border-gray-200">
                <ul class="space-y-2">
                    {#each $tutorialSteps as step, index}
                        <li>
                            <Button
                                color={index === currentStepIndex ? 'primary' : 'tertiary'}
                                classes="w-full flex items-center gap-3 justify-start px-4 py-3 rounded-lg"
                                onclick={() => currentStepIndex = index}
                                disabled={$loading.isLoading}
                            >
                                <span class="w-6 h-6 flex items-center justify-center rounded-full bg-gray-100 text-gray-700 font-semibold">{index + 1}</span>
                                <span class="flex-1 font-medium">{step.title}</span>
                                {#if step.isCompleted}
                                    <span class="text-green-600 font-bold">✓</span>
                                {/if}
                            </Button>
                        </li>
                    {/each}
                </ul>
            </aside>

            <section class="p-8">
                {#if $tutorialSteps[currentStepIndex]}
                    {@const currentStep = $tutorialSteps[currentStepIndex]}
                    <div class="mb-8">
                        <h2 class="text-xl font-semibold text-gray-900 mb-2">{currentStep.title}</h2>
                        <p class="text-gray-700 mb-4">{currentStep.description}</p>
                        {#if currentStep.action}
                            <Button
                                color="primary"
                                onclick={() => handleStepAction(currentStep)}
                                disabled={$loading.isLoading}
                            >
                                {currentStep.action.type === 'navigate'
                                    ? `Go to ${currentStep.action.target.replace(/^\//, '')}`
                                    : 'Continue'}
                            </Button>
                        {/if}
                    </div>
                {/if}

                <div class="flex flex-wrap gap-3 border-t pt-6 mt-6">
                    <Button
                        color="secondary"
                        onclick={prevStep}
                        disabled={currentStepIndex === 0 || $loading.isLoading}
                    >
                        Previous
                    </Button>
                    <Button
                        color="secondary"
                        onclick={nextStep}
                        disabled={currentStepIndex === $tutorialSteps.length - 1 || $loading.isLoading}
                    >
                        Next
                    </Button>
                    <Button
                        color="primary"
                        onclick={completeTutorial}
                        disabled={$loading.isLoading}
                    >
                        {$loading.isLoading ? 'Completing...' : 'Skip Tutorial'}
                    </Button>
                </div>
            </section>
        </Card>

        <div class="mt-8 text-center">
            <Button
                color="secondary"
                type="button"
                onclick={handleBack}
                disabled={$loading.isLoading}
            >
                Back
            </Button>
        </div>
    {/if}
</div>
