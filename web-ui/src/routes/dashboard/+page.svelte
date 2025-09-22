<script lang="ts">
	import { onMount } from 'svelte';
	import { authStore } from '$lib/stores/auth.js';
	import { Button } from '$lib/components/ui/index.js';
	import { formatDate, formatNumber } from '$lib/utils/format.js';

	// Mock data for carrier compliance dashboard - will be replaced with real API calls
	let complianceStats = $state({
		totalDrivers: 45,
		activeDrivers: 42,
		pendingTests: 8,
		upcomingRandomTests: 3,
		expiringCertificates: 5,
		subscriptionStatus: 'active',
		subscriptionRenewal: '2026-01-01'
	});

	let recentOrders = $state([
		{
			id: '1',
			driver_name: 'John Smith',
			driver_cdl: 'ABC123456',
			test_type: 'Drug Test',
			vendor: 'Quest Diagnostics',
			status: 'completed',
			ordered_at: new Date(Date.now() - 3600000).toISOString(),
			completed_at: new Date().toISOString()
		},
		{
			id: '2',
			driver_name: 'Jane Doe',
			driver_cdl: 'DEF789012',
			test_type: 'DOT Physical',
			vendor: 'DOT Medical',
			status: 'processing',
			ordered_at: new Date(Date.now() - 7200000).toISOString(),
			completed_at: null
		},
		{
			id: '3',
			driver_name: 'Mike Johnson',
			driver_cdl: 'GHI345678',
			test_type: 'MVR Check',
			vendor: 'MVR Provider',
			status: 'pending',
			ordered_at: new Date(Date.now() - 86400000).toISOString(),
			completed_at: null
		}
	]);

	let upcomingCompliance = $state([
		{
			id: '1',
			driver_name: 'Sarah Wilson',
			driver_cdl: 'JKL901234',
			item_type: 'DOT Physical',
			expires_at: new Date(Date.now() + 2592000000).toISOString(), // 30 days
			days_remaining: 30
		},
		{
			id: '2',
			driver_name: 'Tom Brown',
			driver_cdl: 'MNO567890',
			item_type: 'Medical Certificate',
			expires_at: new Date(Date.now() + 1296000000).toISOString(), // 15 days
			days_remaining: 15
		},
		{
			id: '3',
			driver_name: 'Lisa Davis',
			driver_cdl: 'PQR234567',
			item_type: 'Random Test Due',
			expires_at: new Date(Date.now() + 604800000).toISOString(), // 7 days
			days_remaining: 7
		}
	]);

	let randomTestingPool = $state({
		total_drivers: 42,
		pool_percentage: 25,
		next_selection: new Date(Date.now() + 2592000000).toISOString(), // 30 days
		last_selection: new Date(Date.now() - 7776000000).toISOString(), // 90 days ago
		frequency: 'quarterly'
	});
</script>

<svelte:head>
	<title>Compliance Dashboard - Carrier Consortium</title>
</svelte:head>

<div class="space-y-6">
	<!-- Welcome Header -->
	<div class="flex justify-between items-start">
		<div>
			<h1 class="text-3xl font-bold text-on-surface-token">
				Compliance Dashboard
			</h1>
			<p class="text-surface-500-400-token mt-2">
				Manage your drivers and DOT/FMCSA compliance requirements.
			</p>
		</div>
		
		<div class="flex gap-3">
			<Button variant="outline" onclick={() => {}}>
				Download Certificates
			</Button>
			<Button variant="filled" onclick={() => {}}>
				Order Tests
			</Button>
		</div>
	</div>

	<!-- Subscription Status Alert -->
	{#if complianceStats.subscriptionStatus === 'active'}
		<div class="card variant-filled-success p-4">
			<div class="flex items-center gap-3">
				<span class="text-2xl">✅</span>
				<div>
					<p class="font-semibold text-success-900">Subscription Active</p>
					<p class="text-sm text-success-800">
						Your consortium membership renews on {formatDate(complianceStats.subscriptionRenewal)}
					</p>
				</div>
			</div>
		</div>
	{/if}

	<!-- Stats Cards -->
	<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
		<div class="card p-6">
			<div class="flex items-center justify-between">
				<div>
					<p class="text-sm font-medium text-surface-500-400-token">Total Drivers</p>
					<p class="text-2xl font-bold text-on-surface-token">{formatNumber(complianceStats.totalDrivers)}</p>
					<p class="text-xs text-success-500 mt-1">{complianceStats.activeDrivers} active</p>
				</div>
				<div class="w-12 h-12 bg-primary-100-800-token rounded-lg flex items-center justify-center">
					<span class="text-2xl">�</span>
				</div>
			</div>
		</div>

		<div class="card p-6">
			<div class="flex items-center justify-between">
				<div>
					<p class="text-sm font-medium text-surface-500-400-token">Pending Tests</p>
					<p class="text-2xl font-bold text-on-surface-token">{formatNumber(complianceStats.pendingTests)}</p>
					<p class="text-xs text-warning-500 mt-1">Awaiting completion</p>
				</div>
				<div class="w-12 h-12 bg-warning-100-800-token rounded-lg flex items-center justify-center">
					<span class="text-2xl">🧪</span>
				</div>
			</div>
		</div>

		<div class="card p-6">
			<div class="flex items-center justify-between">
				<div>
					<p class="text-sm font-medium text-surface-500-400-token">Random Tests Due</p>
					<p class="text-2xl font-bold text-on-surface-token">{formatNumber(complianceStats.upcomingRandomTests)}</p>
					<p class="text-xs text-tertiary-500 mt-1">Next 30 days</p>
				</div>
				<div class="w-12 h-12 bg-tertiary-100-800-token rounded-lg flex items-center justify-center">
					<span class="text-2xl">🎯</span>
				</div>
			</div>
		</div>

		<div class="card p-6">
			<div class="flex items-center justify-between">
				<div>
					<p class="text-sm font-medium text-surface-500-400-token">Expiring Soon</p>
					<p class="text-2xl font-bold text-on-surface-token">{formatNumber(complianceStats.expiringCertificates)}</p>
					<p class="text-xs text-error-500 mt-1">Certificates & physicals</p>
				</div>
				<div class="w-12 h-12 bg-error-100-800-token rounded-lg flex items-center justify-center">
					<span class="text-2xl">⚠️</span>
				</div>
			</div>
		</div>
	</div>

	<!-- Main Content Grid -->
	<div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
		<!-- Recent Test Orders -->
		<div class="card p-6">
			<div class="flex items-center justify-between mb-6">
				<h2 class="text-xl font-semibold text-on-surface-token">Recent Test Orders</h2>
				<a href="/dashboard/orders" class="text-primary-500 hover:text-primary-400 text-sm font-medium">
					View all orders
				</a>
			</div>
			
			{#if recentOrders.length > 0}
				<div class="space-y-4">
					{#each recentOrders as order}
						<div class="flex items-start justify-between p-4 border border-surface-300-600-token rounded-lg">
							<div class="flex-1">
								<div class="flex items-center gap-2 mb-2">
									<h3 class="font-medium text-on-surface-token">{order.driver_name}</h3>
									<span class="badge variant-filled-{order.status === 'completed' ? 'success' : order.status === 'processing' ? 'warning' : 'surface'} text-xs">
										{order.status}
									</span>
								</div>
								<p class="text-sm text-surface-500-400-token">
									{order.test_type} • CDL: {order.driver_cdl}
								</p>
								<p class="text-xs text-surface-500-400-token mt-1">
									Ordered {formatDate(order.ordered_at)}
									{#if order.completed_at}
										• Completed {formatDate(order.completed_at)}
									{/if}
								</p>
							</div>
							<Button variant="ghost" size="sm">
								{order.status === 'completed' ? 'View Results' : 'Track'}
							</Button>
						</div>
					{/each}
				</div>
			{:else}
				<div class="text-center py-8">
					<div class="w-16 h-16 bg-surface-200-700-token rounded-lg flex items-center justify-center mx-auto mb-4">
						<span class="text-2xl">🧪</span>
					</div>
					<p class="text-surface-500-400-token">No test orders yet</p>
					<p class="text-sm text-surface-500-400-token mt-1">
						Order tests for your drivers to maintain compliance
					</p>
				</div>
			{/if}
		</div>

		<!-- Upcoming Compliance Items -->
		<div class="card p-6">
			<div class="flex items-center justify-between mb-6">
				<h2 class="text-xl font-semibold text-on-surface-token">Upcoming Compliance</h2>
				<a href="/dashboard/compliance" class="text-primary-500 hover:text-primary-400 text-sm font-medium">
					View all items
				</a>
			</div>
			
			{#if upcomingCompliance.length > 0}
				<div class="space-y-4">
					{#each upcomingCompliance as item}
						<div class="flex items-start justify-between p-4 border border-surface-300-600-token rounded-lg">
							<div class="flex-1">
								<div class="flex items-center gap-2 mb-2">
									<h3 class="font-medium text-on-surface-token">{item.driver_name}</h3>
									<span class="badge variant-filled-{item.days_remaining <= 7 ? 'error' : item.days_remaining <= 30 ? 'warning' : 'surface'} text-xs">
										{item.days_remaining} days
									</span>
								</div>
								<p class="text-sm text-surface-500-400-token">
									{item.item_type} • CDL: {item.driver_cdl}
								</p>
								<p class="text-xs text-surface-500-400-token mt-1">
									Due {formatDate(item.expires_at)}
								</p>
							</div>
							<Button variant="ghost" size="sm">
								{item.item_type.includes('Test') ? 'Schedule' : 'Renew'}
							</Button>
						</div>
					{/each}
				</div>
			{:else}
				<div class="text-center py-8">
					<div class="w-16 h-16 bg-surface-200-700-token rounded-lg flex items-center justify-center mx-auto mb-4">
						<span class="text-2xl">✅</span>
					</div>
					<p class="text-surface-500-400-token">All compliance items up to date</p>
					<p class="text-sm text-surface-500-400-token mt-1">
						Great job maintaining your drivers' compliance!
					</p>
				</div>
			{/if}
		</div>
	</div>

	<!-- Random Testing Pool Status -->
	<div class="card p-6">
		<div class="flex items-center justify-between mb-6">
			<h2 class="text-xl font-semibold text-on-surface-token">Random Testing Pool</h2>
			<a href="/dashboard/random-testing" class="text-primary-500 hover:text-primary-400 text-sm font-medium">
				Manage pool
			</a>
		</div>
		
		<div class="grid grid-cols-1 md:grid-cols-4 gap-6">
			<div class="text-center">
				<p class="text-2xl font-bold text-on-surface-token">{randomTestingPool.total_drivers}</p>
				<p class="text-sm text-surface-500-400-token">Drivers in Pool</p>
			</div>
			<div class="text-center">
				<p class="text-2xl font-bold text-on-surface-token">{randomTestingPool.pool_percentage}%</p>
				<p class="text-sm text-surface-500-400-token">Selection Rate</p>
			</div>
			<div class="text-center">
				<p class="text-2xl font-bold text-on-surface-token">{randomTestingPool.frequency}</p>
				<p class="text-sm text-surface-500-400-token">Frequency</p>
			</div>
			<div class="text-center">
				<p class="text-sm font-medium text-surface-500-400-token">Next Selection</p>
				<p class="text-sm text-on-surface-token">{formatDate(randomTestingPool.next_selection)}</p>
			</div>
		</div>
		
		<div class="mt-6 p-4 bg-surface-100-800-token rounded-lg">
			<div class="flex items-center gap-3">
				<span class="text-xl">ℹ️</span>
				<div>
					<p class="text-sm font-medium text-on-surface-token">Last Random Selection</p>
					<p class="text-xs text-surface-500-400-token">
						{formatDate(randomTestingPool.last_selection)} - Next automated selection in {Math.ceil((new Date(randomTestingPool.next_selection).getTime() - Date.now()) / (1000 * 60 * 60 * 24))} days
					</p>
				</div>
			</div>
		</div>
	</div>

	<!-- Quick Actions -->
	<div class="card p-6">
		<h2 class="text-xl font-semibold text-on-surface-token mb-6">Quick Actions</h2>
		
		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
			<Button variant="filled" classes="p-6 h-auto flex-col items-start text-left">
				<div class="w-10 h-10 bg-primary-500 rounded-lg flex items-center justify-center mb-3">
					<span class="text-white text-xl">�</span>
				</div>
				<h3 class="font-medium mb-1">Manage Drivers</h3>
				<p class="text-sm text-surface-500-400-token">Add, edit, or view driver profiles</p>
			</Button>

			<Button variant="filled" classes="p-6 h-auto flex-col items-start text-left">
				<div class="w-10 h-10 bg-success-500 rounded-lg flex items-center justify-center mb-3">
					<span class="text-white text-xl">🧪</span>
				</div>
				<h3 class="font-medium mb-1">Order Tests</h3>
				<p class="text-sm text-surface-500-400-token">Drug tests, DOT physicals, MVR checks</p>
			</Button>

			<Button variant="filled" classes="p-6 h-auto flex-col items-start text-left">
				<div class="w-10 h-10 bg-tertiary-500 rounded-lg flex items-center justify-center mb-3">
					<span class="text-white text-xl">�</span>
				</div>
				<h3 class="font-medium mb-1">Download Certificates</h3>
				<p class="text-sm text-surface-500-400-token">Consortium membership & compliance docs</p>
			</Button>

			<Button variant="filled" classes="p-6 h-auto flex-col items-start text-left">
				<div class="w-10 h-10 bg-warning-500 rounded-lg flex items-center justify-center mb-3">
					<span class="text-white text-xl">🎯</span>
				</div>
				<h3 class="font-medium mb-1">Random Testing</h3>
				<p class="text-sm text-surface-500-400-token">Manage pools and view selections</p>
			</Button>
		</div>
	</div>
</div>
