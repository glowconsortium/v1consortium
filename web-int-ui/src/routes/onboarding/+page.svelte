<script lang="ts">
    import { onMount } from 'svelte';
    import { onboardingStore } from '$lib/stores/onboardingStore.js';
    import { goto } from '$app/navigation';
    import { OnboardingStep, InvitedOnboardingStep, type OnboardingState, type OnboardingProgress, UserType, OnboardingType, type InvitedOnboardingState } from '$lib/types/onboarding.js';

    // Import your UI library components
    import { Button, Alert, Spinner, Card } from '@movsm/v1-consortium-web-pkg';

    const { subscribe, progress, loading, error } = onboardingStore;
    
    let currentState = $state<OnboardingState | InvitedOnboardingState>();
    let progressData = $state<OnboardingProgress>();
    
    // Subscribe to store updates
    $effect(() => {
        const unsubscribe = subscribe((state) => {
            currentState = state;
        });
        return unsubscribe;
    });
    
    $effect(() => {
        const unsubscribe = progress.subscribe((prog) => {
            progressData = prog;
        });
        return unsubscribe;
    });
    
    onMount(async () => {
        await onboardingStore.loadOnboardingState();
        
        // Redirect if already completed
        if (currentState?.isCompleted) {
            goto('/dashboard');
            return;
        }
        
        // Redirect to specific step if onboarding is in progress
        if (currentState?.currentStep && currentState.currentStep !== (currentState.userType === UserType.INVITED ? InvitedOnboardingStep.WELCOME : OnboardingStep.WELCOME)) {
            redirectToCurrentStep(currentState.currentStep);
        }
    });
    
    function redirectToCurrentStep(step: OnboardingStep | InvitedOnboardingStep) {
        switch (step) {
            case OnboardingStep.ORGANIZATION:
                if (currentState?.userType === UserType.SIGNUP) {
                    goto('/onboarding/organization');
                }
                break;
            case OnboardingStep.PLAN_SELECTION:
                if (currentState?.userType === UserType.SIGNUP) {
                    goto('/onboarding/plans');
                }
                break;
            case OnboardingStep.TUTORIAL:
            case InvitedOnboardingStep.TUTORIAL:
                goto('/onboarding/tutorial');
                break;
            case OnboardingStep.COMPLETED:
            case InvitedOnboardingStep.COMPLETED:
                goto('/onboarding/complete');
                break;
            default:
                // Stay on welcome page
                break;
        }
    }
    
    function startOnboarding() {
        console.log('Starting onboarding process...');
        if (currentState?.userType === UserType.INVITED) {
            goto('/onboarding/tutorial', { replaceState: true });
        } else {
            goto('/onboarding/organization', { replaceState: true });
        }
    }
    
    function skipOnboarding() {
        goto('/dashboard');
    }
    
    function continueOnboarding() {
        console.log('Continuing onboarding...');
        const welcomeStep = currentState?.userType === UserType.INVITED ? InvitedOnboardingStep.WELCOME : OnboardingStep.WELCOME;
        if (currentState?.currentStep && currentState?.currentStep !== welcomeStep) {
            redirectToCurrentStep(currentState.currentStep);
        } else {
            startOnboarding();
        }
    }

    function getCurrentStepLabel(step: OnboardingStep | InvitedOnboardingStep): string {
        switch (step) {
            case OnboardingStep.WELCOME:
            case InvitedOnboardingStep.WELCOME:
                return 'Getting Started';
            case OnboardingStep.ORGANIZATION:
                return 'Organization Setup';
            case OnboardingStep.PLAN_SELECTION:
                return 'Plan Selection';
            case OnboardingStep.TUTORIAL:
            case InvitedOnboardingStep.TUTORIAL:
                return 'Quick Tutorial';

            case OnboardingStep.COMPLETED:
            case InvitedOnboardingStep.COMPLETED:
                return 'Completed';
            default:
                return 'Unknown Step';
        }
    }

    function getOnboardingSteps() {
        if (currentState?.userType === UserType.INVITED) {
            return [
                {
                    step: 1,
                    title: 'Quick Tutorial',
                    description: 'Learn how to use FormApp effectively'
                },

            ];
        }

        return [
            {
                step: 1,
                title: 'Create Organization',
                description: 'Set up your workspace for team collaboration'
            },
            {
                step: 2,
                title: 'Choose Your Plan',
                description: 'Select the plan that fits your needs (start free!)'
            },

        ];
    }
</script>

<svelte:head>
    <title>Welcome to FormApp - Get Started</title>
    <meta name="description" content="Welcome to FormApp! Let's get you set up to start collecting data." />
</svelte:head>

<div class="min-h-screen flex flex-col justify-center items-center p-4 sm:p-8 bg-gradient-to-br from-blue-500 to-purple-600 text-white">
    {#if $loading.isLoading}
        <div class="flex flex-col items-center gap-4">
            <Spinner />
            <p class="text-white/80">{$loading.operation || 'Loading...'}</p>
        </div>
    {:else}
        <div class="max-w-4xl w-full text-center">
            <div class="mb-12">
                <div class="w-20 h-20 mx-auto mb-8 p-4 bg-white/10 rounded-full backdrop-blur-lg">
                    <svg class="w-full h-full text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                        <polyline points="14,2 14,8 20,8"></polyline>
                        <line x1="16" y1="13" x2="8" y2="13"></line>
                        <line x1="16" y1="17" x2="8" y2="17"></line>
                        <polyline points="10,9 9,9 8,9"></polyline>
                    </svg>
                </div>
                <h1 class="text-4xl sm:text-5xl font-bold mb-4">Welcome to FormApp! 🎉</h1>
                <p class="text-lg sm:text-xl text-white/90 max-w-2xl mx-auto">
                    {currentState?.userType === UserType.INVITED ? 
                        "You've been invited to join an organization. Let's get you set up!" :
                        "The easiest way to collect form submissions and API data without managing servers."
                    }
                </p>
            </div>
            
            {#if currentState && !currentState.isCompleted}
                <Card classes="bg-white/10 backdrop-blur-lg p-8 rounded-2xl mb-8">
                    <h2 class="text-2xl font-semibold mb-4">Continue Your Setup</h2>
                    <p class="text-lg mb-4">You're {progressData?.progress ?? 0}% complete</p>
                    
                    <div class="w-full h-2 bg-white/20 rounded-full overflow-hidden mb-6">
                        <div 
                            class="h-full bg-emerald-500 transition-all duration-300 ease-in-out"
                            style="width: {progressData?.progress ?? 0}%"
                        ></div>
                    </div>
                    
                    <div class="mb-6">
                        <p class="text-lg">Current step: <strong>{currentState?.currentStep ? getCurrentStepLabel(currentState.currentStep) : 'Getting Started'}</strong></p>
                    </div>
                    
                    <div class="flex flex-col sm:flex-row gap-4 justify-center">
                        <Button onclick={continueOnboarding} color="primary">
                            Continue Setup
                        </Button>
                        <Button onclick={skipOnboarding} variant="outline" color="secondary">
                            Skip for now
                        </Button>
                    </div>
                </Card>
            {:else}
                <div class="mb-12">
                    <h2 class="text-2xl font-semibold mb-8">What you'll get with FormApp:</h2>
                    
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

                        <Card classes="bg-white/10 backdrop-blur-lg p-6 rounded-xl text-left">
                            <div class="text-4xl mb-4">🔒</div>
                            <h3 class="text-xl font-semibold mb-2">Enterprise Security</h3>
                            <p class="text-white/80">Built-in spam protection, CORS handling, and data validation</p>
                        </Card>
                        <Card classes="bg-white/10 backdrop-blur-lg p-6 rounded-xl text-left">
                            <div class="text-4xl mb-4">🔗</div>
                            <h3 class="text-xl font-semibold mb-2">Flexible Integrations</h3>
                            <p class="text-white/80">Connect with webhooks, email notifications, and third-party services</p>
                        </Card>
                        <Card classes="bg-white/10 backdrop-blur-lg p-6 rounded-xl text-left">
                            <div class="text-4xl mb-4">📊</div>
                            <h3 class="text-xl font-semibold mb-2">Real-time Analytics</h3>
                            <p class="text-white/80">Monitor submissions, track usage, and analyze your data collection</p>
                        </Card>
                    </div>
                </div>
                
                <div class="mb-12">
                    <h2 class="text-2xl font-semibold mb-8">Let's get you set up in {currentState?.userType === UserType.INVITED ? '2' : '3'} easy steps:</h2>
                    
                    <div class="max-w-xl mx-auto space-y-4">
                        {#each getOnboardingSteps() as { step, title, description }}
                            <Card classes="bg-white/10 backdrop-blur-lg p-6 rounded-xl flex items-center gap-6">
                                <div class="w-10 h-10 rounded-full bg-emerald-500 flex items-center justify-center font-bold flex-shrink-0">
                                    {step}
                                </div>
                                <div>
                                    <h3 class="text-lg font-semibold mb-1">{title}</h3>
                                    <p class="text-white/80">{description}</p>
                                </div>
                            </Card>
                        {/each}
                    </div>
                </div>
                
                <div class="flex flex-col sm:flex-row gap-4 justify-center mb-8">
                    <Button onclick={startOnboarding} color="primary">
                        Get Started ({currentState?.userType === UserType.INVITED ? '1' : '2'} min setup)
                    </Button>
                    <Button onclick={skipOnboarding} color="secondary">
                        Skip Setup
                    </Button>
                </div>
            {/if}
            
            <div class="text-sm text-white/70">
                <p>Need help? Check out our <a href="/docs" target="_blank" class="underline hover:text-white">documentation</a> or <a href="/support" target="_blank" class="underline hover:text-white">contact support</a></p>
            </div>
        </div>
    {/if}
    
    {#if $error}
        <div class="fixed top-8 left-1/2 -translate-x-1/2 max-w-md w-full bg-red-500 text-white p-4 rounded-lg shadow-lg flex items-center justify-between">
            <div class="flex items-center gap-2">
                <strong>Setup Error:</strong> 
                <span>{$error.message}</span>
            </div>
            {#if $error.retryable}
                <Button onclick={() => onboardingStore.clearError()} variant="outline" color="secondary">
                    Try Again
                </Button>
            {/if}
        </div>
    {/if}
</div>