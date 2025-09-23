<script lang="ts">
	// import { Button } from '$lib/components/ui/index.js';
	import { authStore } from '@movsm/v1-consortium-web-pkg';
	import { Button } from '@movsm/v1-consortium-web-pkg';
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
					<a href="/pricing">
						<Button variant="outline" size="lg">
							View Pricing
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
				
				<div>
					<h4 class="font-semibold text-on-surface-token mb-4">Services</h4>
					<ul class="space-y-2 text-sm text-surface-500-400-token">
						<li><a href="/services/drug-testing" class="hover:text-primary-500">Drug & Alcohol Testing</a></li>
						<li><a href="/services/dot-physicals" class="hover:text-primary-500">DOT Physicals</a></li>
						<li><a href="/services/mvr" class="hover:text-primary-500">Motor Vehicle Records</a></li>
						<li><a href="/services/random-testing" class="hover:text-primary-500">Random Testing Programs</a></li>
					</ul>
				</div>
				
				<div>
					<h4 class="font-semibold text-on-surface-token mb-4">Resources</h4>
					<ul class="space-y-2 text-sm text-surface-500-400-token">
						<li><a href="/compliance-guide" class="hover:text-primary-500">Compliance Guide</a></li>
						<li><a href="/fmcsa-requirements" class="hover:text-primary-500">FMCSA Requirements</a></li>
						<li><a href="/api-docs" class="hover:text-primary-500">API Documentation</a></li>
						<li><a href="/support" class="hover:text-primary-500">Support Center</a></li>
					</ul>
				</div>
				
				<div>
					<h4 class="font-semibold text-on-surface-token mb-4">Company</h4>
					<ul class="space-y-2 text-sm text-surface-500-400-token">
						<li><a href="/about" class="hover:text-primary-500">About Us</a></li>
						<li><a href="/pricing" class="hover:text-primary-500">Pricing</a></li>
						<li><a href="/privacy" class="hover:text-primary-500">Privacy Policy</a></li>
						<li><a href="/terms" class="hover:text-primary-500">Terms of Service</a></li>
					</ul>
				</div>
			</div>
			
			<div class="border-t border-surface-300-600-token mt-8 pt-8 text-center text-sm text-surface-500-400-token">
				<p>&copy; 2025 Consortium Compliance. All rights reserved. | DOT/FMCSA Compliant Testing Platform</p>
			</div>
		</div>
	</footer>
</div>
