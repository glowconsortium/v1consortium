<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Header, Sidebar } from '@movsm/v1-consortium-web-pkg';
	import { Spinner } from '@movsm/v1-consortium-web-pkg';

	let { children } = $props();

		let sidebarItems: SidebarItem[] = [
		{
			label: 'Dashboard',
			href: '/dashboard',
			icon: '📊'
		},
		{
			label: 'Members',
			href: '/dashboard/members',
			icon: '👥'
		},
		{
			label: 'Results',
			href: '/dashboard/results',
			icon: '📁'
		},

		{
			label: 'Orders',
			href: '/dashboard/orders',
			icon: '💊'
		},

		{
			label: 'Integrations',
			href: '/dashboard/integrations',
			icon: '🔌'
		},

				{
			label: 'Invitations',
			href: '/dashboard/invitations',
			icon: '🔗'
		},
		{
			label: 'Settings',
			href: '/dashboard/settings',
			icon: '⚙️'
		}
	];

	// Load onboarding state when component mounts
	// onMount(async () => {
	// 	await onboardingStore.loadOnboardingState();
	// });

	// Redirect if not authenticated or onboarding not completed
	$effect(() => {
		if ($authStore.isInitialized) {
			if (!$authStore.isAuthenticated) {
				goto('/auth/signin');
			}
			// } else if (!$onboardingStore.isCompleted) {
			// 	goto('/onboarding');
			// }
		}
	});

	// Show content when authenticated and initialized
	const showContent = $derived($authStore.isInitialized && $authStore.isAuthenticated);
	const isLoading = $derived(!$authStore.isInitialized || $authStore.isLoading);
</script>

{#if isLoading}
	<div class="min-h-screen flex items-center justify-center bg-surface-50">
		<div class="text-center">
			<Spinner size="lg" color="primary" />
			<p class="mt-4 text-surface-600">Loading...</p>
		</div>
	</div>
{:else if showContent}
	<div class="flex h-screen overflow-hidden bg-surface-50-900-token">
		<Sidebar />
		<div class="flex-1 flex flex-col overflow-hidden">
			<Header />
			<main class="flex-1 overflow-auto p-4">
				<div class="container mx-auto">
					{@render children()}
				</div>
			</main>
		</div>
	</div>
{:else}
	<!-- Redirect trigger -->
	<div></div>
{/if}