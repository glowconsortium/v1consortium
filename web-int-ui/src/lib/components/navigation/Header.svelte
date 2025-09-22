<script lang="ts">
	import { authStore } from '$lib/stores/auth.js';
	import { Button } from '$lib/components/ui/index.js';
	import { goto } from '$app/navigation';
	// import OrganizationSelector from '../organization/OrganizationSelector.svelte';

	let showMobileMenu = $state(false);
	let showUserMenu = $state(false);

	async function handleLogout() {
		await authStore.logout();
		goto('/');
	}

	function toggleMobileMenu() {
		showMobileMenu = !showMobileMenu;
	}

	function toggleUserMenu() {
		showUserMenu = !showUserMenu;
	}
</script>

<header class="border-b border-slate-200 bg-white px-6 py-4">
	<div class="flex items-center justify-between">
		<!-- Mobile menu button (hidden on desktop) -->
		<button
			class="rounded-md p-2 text-slate-600 transition-colors hover:bg-slate-100 md:hidden"
			onclick={toggleMobileMenu}
			aria-label="Toggle mobile menu"
		>
			<svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 6h16M4 12h16M4 18h16"
				/>
			</svg>
		</button>

		<!-- Page Title and Organization Selector Section -->
		<div class="flex flex-1 items-center gap-4">
			<div>
				<h1 class="text-2xl font-bold text-slate-900">Dashboard</h1>
				<p class="mt-1 text-sm text-slate-600">Manage your company</p>
			</div>
			<div class="hidden md:block">
				<!-- <OrganizationSelector /> -->
			</div>
		</div>

		<!-- User Menu -->
		<div class="relative">
			<button
				class="flex items-center space-x-3 rounded-lg p-2 transition-colors hover:bg-slate-100"
				onclick={toggleUserMenu}
			>
				<!-- User Avatar -->
				<div
					class="flex h-8 w-8 items-center justify-center rounded-full bg-gradient-to-br from-green-400 to-blue-500"
				>
					<span class="text-sm font-semibold text-white">
						{$authStore.user?.email?.charAt(0).toUpperCase() || 'U'}
					</span>
				</div>
				<div class="hidden text-left md:block">
					<div class="text-sm font-medium text-slate-900">
						{$authStore.user?.name || 'User'}
					</div>
					<div class="text-xs text-slate-600">
						{$authStore.user?.email || 'user@example.com'}
					</div>
				</div>
				<svg class="h-4 w-4 text-slate-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M19 9l-7 7-7-7"
					/>
				</svg>
			</button>

			<!-- User Dropdown Menu -->
			{#if showUserMenu}
				<div
					class="absolute right-0 z-50 mt-2 w-48 rounded-lg border border-slate-200 bg-white py-2 shadow-lg"
				>
					<!-- Show organization selector in mobile menu -->
					<div class="border-b border-slate-200 px-4 py-2 md:hidden">
						<!-- <OrganizationSelector /> -->
					</div>
					<a href="/dashboard/profile" class="dropdown-item">
						<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
							/>
						</svg>
						Profile
					</a>
					<a href="/dashboard/settings" class="dropdown-item">
						<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
							/>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
							/>
						</svg>
						Settings
					</a>
					<hr class="my-2 border-slate-200" />
					<button onclick={handleLogout} class="dropdown-item w-full text-left text-red-600">
						<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
							/>
						</svg>
						Sign Out
					</button>
				</div>
			{/if}
		</div>
	</div>
</header>

<style>
	.dropdown-item {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		padding: 0.5rem 1rem;
		font-size: 0.875rem;
		color: #475569; /* slate-600 */
		text-decoration: none;
		transition: background-color 150ms ease-in-out;
	}

	.dropdown-item:hover {
		background-color: #f8fafc; /* slate-50 */
	}
</style>
