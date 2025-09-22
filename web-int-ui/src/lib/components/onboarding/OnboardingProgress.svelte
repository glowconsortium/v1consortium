<script lang="ts">
	import { onboardingStore } from '$lib/stores/onboardingStore.js';
	import {
		OnboardingStep,
		type OnboardingProgress,
		type OnboardingState,
		type InvitedOnboardingState
	} from '$lib/types/onboarding.js';
	import { onDestroy } from 'svelte';

	const STEP_LABELS = {
		[OnboardingStep.WELCOME]: 'Welcome',
		[OnboardingStep.ORGANIZATION]: 'Create Organization',
		[OnboardingStep.PLAN_SELECTION]: 'Choose Plan',
		[OnboardingStep.TUTORIAL]: 'Quick Tutorial',

		[OnboardingStep.COMPLETED]: 'Complete'
	} as const;

	// Use Svelte store auto-subscription
	let currentState: OnboardingState | InvitedOnboardingState;
	let progressData: OnboardingProgress;

	const unsubscribeState = onboardingStore.subscribe((state) => (currentState = state));
	const unsubscribeProgress = onboardingStore.progress.subscribe((prog) => (progressData = prog));

	onDestroy(() => {
		unsubscribeState();
		unsubscribeProgress();
	});

	// Only show the main steps (not WELCOME or COMPLETED)
	const steps: OnboardingStep[] = [
		OnboardingStep.ORGANIZATION,
		OnboardingStep.PLAN_SELECTION,
		OnboardingStep.TUTORIAL,
	].filter((step) => step in STEP_LABELS);
</script>

<div class="mb-8 rounded-lg bg-white p-6 dark:bg-gray-800">
	<div class="mb-4 flex items-center justify-between">
		<h3 class="m-0 text-lg font-semibold">Setup Progress</h3>
		<span class="text-sm text-gray-600 dark:text-gray-400"
			>{progressData?.progress ?? 0}% Complete</span
		>
	</div>

	<div class="mb-6 h-2 overflow-hidden rounded bg-gray-100 dark:bg-gray-700">
		<div
			class="h-full bg-blue-600 transition-all duration-300 ease-in-out dark:bg-blue-500"
			style="width: {progressData?.progress ?? 0}%"
		></div>
	</div>

	<div class="flex flex-col gap-4">
		{#if currentState && currentState.userType === 'SIGNUP'}
			{#each steps as step, index}
				{@const isCompleted = (currentState.completedSteps as OnboardingStep[])?.includes(step)}
				{@const isCurrent = (currentState.currentStep as OnboardingStep) === step}
				{@const stepLabel = STEP_LABELS[step as keyof typeof STEP_LABELS] ?? step}

				<div
					class="flex items-center gap-4 rounded p-2 {isCurrent
						? 'bg-gray-50 dark:bg-gray-700'
						: ''} transition-all duration-200"
				>
					<div
						class="
							flex h-6 w-6 items-center justify-center rounded-full text-sm
							{isCompleted
								? 'bg-green-500 text-white'
								: isCurrent
									? 'bg-blue-600 text-white'
									: 'bg-gray-100 text-gray-600 dark:bg-gray-600 dark:text-gray-400'}
						"
					>
						{#if isCompleted}
							<svg class="h-4 w-4 stroke-[3]" viewBox="0 0 24 24" fill="none" stroke="currentColor">
								<polyline points="20,6 9,17 4,12"></polyline>
							</svg>
						{:else}
							<span>{index + 1}</span>
						{/if}
					</div>
					<span class="flex-1 font-medium {isCompleted ? 'text-gray-600 dark:text-gray-400' : ''}"
						>{stepLabel}</span
					>
					{#if isCurrent}
						<span class="text-xs font-medium text-blue-600 dark:text-blue-500">Current Step</span>
					{/if}
				</div>
			{/each}
		{/if}
	</div>
</div>
