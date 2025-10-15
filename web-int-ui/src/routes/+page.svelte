<script lang="ts">
	import { authStore } from '@movsm/v1-consortium-web-pkg/stores';
	import { Button } from '@movsm/v1-consortium-web-pkg/ui';
</script>

<svelte:head>
	<title>Consortium Compliance - DOT/FMCSA Compliance Management</title>
	<meta name="description" content="Complete compliance management for motor carriers: drug testing, DOT physicals, MVR checks, and random testing programs" />
</svelte:head>

<div class="min-h-screen bg-surface-50-900-token">
	

	<!-- CTA Section -->
	<section class="bg-surface-100-800-token py-20">
		<div class="container mx-auto px-4 text-center">
			<h2 class="text-3xl md:text-4xl font-bold text-on-surface-token mb-4">
				Ready to streamline your compliance?
			</h2>
			<p class="text-xl text-surface-500-400-token mb-8 max-w-2xl mx-auto">
				Join motor carriers nationwide who trust our platform for complete DOT/FMCSA compliance management.
			</p>
			
			{#if !$authStore.isAuthenticated}
				<div class="flex flex-col sm:flex-row gap-4 justify-center">
					<a href="/auth/signup">
						<Button variant="filled" size="lg">
							Start Annual Subscription
						</Button>
					</a>
					<a href="/auth/signin">
						<Button variant="outline" size="lg">
							Sign In
						</Button>
					</a>
				</div>
			{:else}
				<a href="/dashboard">
					<Button variant="filled" size="lg">
						Go to Dashboard
					</Button>
				</a>
			{/if}
		</div>
	</section>

	<!-- Footer -->
	<footer class="border-t border-surface-300-600-token py-12">
		<div class="container mx-auto px-4">
			<div class="grid grid-cols-1 md:grid-cols-4 gap-8">
				<div>
					<div class="flex items-center space-x-2 mb-4">
						<div class="w-8 h-8 bg-primary-500 rounded-lg flex items-center justify-center">
							<span class="text-white font-bold text-sm">C</span>
						</div>
						<span class="text-xl font-bold text-on-surface-token">Consortium Compliance</span>
					</div>
					<p class="text-surface-500-400-token text-sm">
						Complete DOT/FMCSA compliance management for motor carriers and owner-operators.
					</p>
				</div>
				

				

				

			</div>
			
			<div class="border-t border-surface-300-600-token mt-8 pt-8 text-center text-sm text-surface-500-400-token">
				<p>&copy; 2025 Consortium Compliance. All rights reserved. | DOT/FMCSA Compliant Testing Platform</p>
			</div>
		</div>
	</footer>
</div>
