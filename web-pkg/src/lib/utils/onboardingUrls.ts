import { OnboardingStep } from '../types/onboarding.js';
import { goto } from '$app/navigation';

export const ONBOARDING_ROUTES = {
    [OnboardingStep.WELCOME]: '/onboarding/welcome',
    [OnboardingStep.ORGANIZATION]: '/onboarding/organization',
    [OnboardingStep.PLAN_SELECTION]: '/onboarding/plan',
    [OnboardingStep.TUTORIAL]: '/onboarding/tutorial',
    [OnboardingStep.FIRST_ENDPOINT]: '/onboarding/endpoint',
    [OnboardingStep.COMPLETED]: '/dashboard'
} as const;

export const ONBOARDING_STEP_NAMES = {
    [OnboardingStep.WELCOME]: 'Welcome',
    [OnboardingStep.ORGANIZATION]: 'Create Organization',
    [OnboardingStep.PLAN_SELECTION]: 'Choose Plan',
    [OnboardingStep.TUTORIAL]: 'Quick Tutorial',
    [OnboardingStep.FIRST_ENDPOINT]: 'Create Endpoint',
    [OnboardingStep.COMPLETED]: 'Complete'
} as const;

export function getOnboardingStepFromUrl(pathname: string): OnboardingStep | null {
    const routeEntries = Object.entries(ONBOARDING_ROUTES);
    for (const [step, route] of routeEntries) {
        if (pathname === route) {
            return step as OnboardingStep;
        }
    }
    return null;
}

export function getUrlForStep(step: OnboardingStep): string {
    return ONBOARDING_ROUTES[step] || '/dashboard';
}

export function isOnboardingRoute(pathname: string): boolean {
    return pathname.startsWith('/onboarding');
}

export async function navigateToStep(step: OnboardingStep, replace = false) {
    const url = getUrlForStep(step);
    if (replace) {
        await goto(url, { replaceState: true });
    } else {
        await goto(url);
    }
}

export function getNextStepUrl(currentStep: OnboardingStep): string {
    const steps = Object.values(OnboardingStep);
    const currentIndex = steps.indexOf(currentStep);
    const nextStep = steps[currentIndex + 1] || OnboardingStep.COMPLETED;
    return getUrlForStep(nextStep);
}

export function getPreviousStepUrl(currentStep: OnboardingStep): string {
    const steps = Object.values(OnboardingStep);
    const currentIndex = steps.indexOf(currentStep);
    const previousStep = steps[currentIndex - 1] || OnboardingStep.WELCOME;
    return getUrlForStep(previousStep);
}

export function buildOnboardingQueryParams(params: {
    utm_source?: string;
    utm_medium?: string;
    utm_campaign?: string;
    referral?: string;
}): string {
    const searchParams = new URLSearchParams();
    
    Object.entries(params).forEach(([key, value]) => {
        if (value) {
            searchParams.set(key, value);
        }
    });

    return searchParams.toString();
}

export function parseOnboardingQueryParams(search: string) {
    const params = new URLSearchParams(search);
    return {
        utm_source: params.get('utm_source'),
        utm_medium: params.get('utm_medium'),
        utm_campaign: params.get('utm_campaign'),
        referral: params.get('referral')
    };
}

export function shouldRedirectToOnboarding(user: any, currentPath: string): boolean {
    // Don't redirect if already on onboarding
    if (isOnboardingRoute(currentPath)) {
        return false;
    }

    // Don't redirect if onboarding is completed
    if (user?.onboarding_completed) {
        return false;
    }

    // Don't redirect from auth routes
    if (currentPath.startsWith('/auth')) {
        return false;
    }

    // Don't redirect from public routes
    if (currentPath === '/' || currentPath.startsWith('/public')) {
        return false;
    }

    return true;
}