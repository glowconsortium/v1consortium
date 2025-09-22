<script lang="ts">
	import { page } from '$app/stores';
	import Icon from '../ui/Icon.svelte';

	interface SidebarItem {
		label: string;
		href: string;
		icon: string;
		active?: boolean;
		badge?: string;
	}

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

	let isCollapsed = $state(false);
	let currentPath = $derived($page.url.pathname);

	function toggleSidebar() {
		isCollapsed = !isCollapsed;
	}
</script>

<div class="sidebar-container {isCollapsed ? 'collapsed' : ''} h-full">
	<aside class="bg-surface-100-800-token h-full border-r border-surface-500/30">
		<!-- Brand Section -->
		<div class="flex items-center justify-between border-b border-surface-500/30 p-4">
			<div class="flex items-center gap-3 {isCollapsed ? 'justify-center' : ''}">
				<!-- Brand Icon with Gradient -->
				<div
					class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-primary-500 to-secondary-500"
				>
					<span class="text-sm font-bold text-white">C</span>
				</div>
				{#if !isCollapsed}
					<h2 class="h3">Consortium</h2>
				{/if}
			</div>
			<button
				class="variant-ghost-surface btn-icon"
				onclick={toggleSidebar}
				aria-label={isCollapsed ? 'Expand Sidebar' : 'Collapse Sidebar'}
			>
				<span class="material-symbols-outlined">
					<Icon name={isCollapsed ? 'chevron_right' : 'chevron_left'} size={16}/>
					</span>
			</button>
		</div>

		<!-- Navigation -->
		<nav class="p-4">
			<div class="space-y-2">
				{#each sidebarItems as item}
					<a
						href={item.href}
						class="btn {isCollapsed ? 'justify-center !p-2' : 'justify-start'} {currentPath ===
						item.href
							? 'variant-filled-primary'
							: 'variant-ghost-surface'} w-full"
					>
						<span class="text-lg">{item.icon}</span>
						{#if !isCollapsed}
							<span class="flex-1">{item.label}</span>
							{#if item.badge}
								<span class="variant-filled badge">{item.badge}</span>
							{/if}
						{/if}
					</a>
				{/each}
			</div>
		</nav>

		<!-- Organization Section -->
		<div class="mt-auto border-t border-surface-500/30 p-4">
			<div class="flex items-center gap-3 {isCollapsed ? 'justify-center' : ''}">
				<div class="bg-surface-300-600-token flex h-8 w-8 items-center justify-center rounded-full">
					<span class="text-sm">PW</span>
				</div>
				{#if !isCollapsed}
					<div>
						<div class="font-medium">Personal Workspace</div>
						<div class="text-xs opacity-50">Free Plan</div>
					</div>
				{/if}
			</div>
			{#if !isCollapsed}
				<div class="mt-4 flex items-center justify-between">
					<button class="variant-ghost-surface btn btn-sm">
						<span class="material-symbols-outlined">dark_mode</span>
					</button>
					<a href="/dashboard/profile" class="variant-ghost-surface btn btn-sm"> Profile </a>
				</div>
			{/if}
		</div>
	</aside>
</div>

<style>
	.sidebar-container {
		width: 280px;
		transition: width 200ms ease-in-out;
	}

	.sidebar-container.collapsed {
		width: 72px;
	}

	aside {
		display: flex;
		flex-direction: column;
	}
</style>
