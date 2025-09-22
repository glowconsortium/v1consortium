<script lang="ts">
	import { Button } from '$lib/components/ui/index.js';
	import { authStore } from '$lib/stores/auth.js';
</script>

<svelte:head>
	<title>Consortium Compliance - DOT/FMCSA Compliance Management</title>
	<meta name="description" content="Complete compliance management for motor carriers: drug testing, DOT physicals, MVR checks, and random testing programs" />
</svelte:head>

<div class="min-h-screen bg-surface-50-900-token">
	
	<!-- Hero Section -->
	<section class="bg-gradient-to-br from-primary-500 to-secondary-500 text-white">
		<div class="container mx-auto px-4 py-20 text-center">
			<h1 class="text-5xl md:text-6xl font-bold mb-6">
				Stay DOT Compliant.<br />
				<span class="text-primary-200">Drive with Confidence.</span>
			</h1>
			<p class="text-xl md:text-2xl mb-8 text-primary-100 max-w-3xl mx-auto">
				Complete FMCSA compliance management for motor carriers. Drug testing, DOT physicals, 
				MVR checks, and automated random testing programs - all in one platform.
			</p>
			
			<div class="flex flex-col sm:flex-row gap-4 justify-center">
				{#if $authStore.isAuthenticated}
					<a href="/dashboard">
						<Button variant="filled" size="lg" classes="bg-white text-primary-600 hover:bg-gray-100">
							Go to Dashboard
						</Button>
					</a>
				{:else}
					<a href="/auth/signup">
						<Button variant="outline" size="lg" classes="bg-white text-primary-600 hover:bg-gray-100">
							Start Free Trial
						</Button>
					</a>
					<a href="/auth/signin">
						<Button variant="ghost" size="lg" classes="border-white text-white hover:bg-white/10">
							Sign In
						</Button>
					</a>
				{/if}
			</div>
			
			<div class="mt-12 text-sm text-primary-200">
				✓ Annual subscription (Jan-Dec)  ✓ FMCSA compliant  ✓ Instant consortium certificates
			</div>
		</div>
	</section>

	<!-- Features Section -->
	<section class="py-20">
		<div class="container mx-auto px-4">
			<div class="text-center mb-16">
				<h2 class="text-3xl md:text-4xl font-bold text-on-surface-token mb-4">
					Complete DOT compliance management
				</h2>
				<p class="text-xl text-surface-500-400-token max-w-2xl mx-auto">
					Everything motor carriers need to stay FMCSA compliant: automated testing, random selection programs, and instant certificate downloads.
				</p>
			</div>

			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
				<!-- Feature 1 -->
				<div class="card p-8 text-center">
					<div class="w-16 h-16 bg-primary-500 rounded-xl flex items-center justify-center mx-auto mb-6">
						<span class="text-2xl text-white">🧪</span>
					</div>
					<h3 class="text-xl font-semibold mb-4 text-on-surface-token">Drug & Alcohol Testing</h3>
					<p class="text-surface-500-400-token">
						Automated ordering through Quest Diagnostics and other certified labs. Track results and maintain compliance records automatically.
					</p>
				</div>

				<!-- Feature 2 -->
				<div class="card p-8 text-center">
					<div class="w-16 h-16 bg-success-500 rounded-xl flex items-center justify-center mx-auto mb-6">
						<span class="text-2xl text-white">🏥</span>
					</div>
					<h3 class="text-xl font-semibold mb-4 text-on-surface-token">DOT Physicals</h3>
					<p class="text-surface-500-400-token">
						Schedule and track DOT medical examinations. Store medical certificates and get automated reminders for renewals.
					</p>
				</div>

				<!-- Feature 3 -->
				<div class="card p-8 text-center">
					<div class="w-16 h-16 bg-tertiary-500 rounded-xl flex items-center justify-center mx-auto mb-6">
						<span class="text-2xl text-white">�</span>
					</div>
					<h3 class="text-xl font-semibold mb-4 text-on-surface-token">Motor Vehicle Records</h3>
					<p class="text-surface-500-400-token">
						Pull MVRs from all 50 states. Automated monitoring for violations and license status changes to keep your fleet compliant.
					</p>
				</div>

				<!-- Feature 4 -->
				<div class="card p-8 text-center">
					<div class="w-16 h-16 bg-warning-500 rounded-xl flex items-center justify-center mx-auto mb-6">
						<span class="text-2xl text-white">🎯</span>
					</div>
					<h3 class="text-xl font-semibold mb-4 text-on-surface-token">Random Testing Programs</h3>
					<p class="text-surface-500-400-token">
						FMCSA-compliant random selection algorithms. Automated quarterly selections with full audit trails and selection reports.
					</p>
				</div>

				<!-- Feature 5 -->
				<div class="card p-8 text-center">
					<div class="w-16 h-16 bg-error-500 rounded-xl flex items-center justify-center mx-auto mb-6">
						<span class="text-2xl text-white">📜</span>
					</div>
					<h3 class="text-xl font-semibold mb-4 text-on-surface-token">Instant Certificates</h3>
					<p class="text-surface-500-400-token">
						Download consortium membership certificates, test completion certificates, and compliance reports instantly. Audit-ready documentation.
					</p>
				</div>

				<!-- Feature 6 -->
				<div class="card p-8 text-center">
					<div class="w-16 h-16 bg-surface-500 rounded-xl flex items-center justify-center mx-auto mb-6">
						<span class="text-2xl text-white">⚙️</span>
					</div>
					<h3 class="text-xl font-semibold mb-4 text-on-surface-token">API Integrations</h3>
					<p class="text-surface-500-400-token">
						Connect with your existing carrier management systems. Token-based API access for seamless data synchronization.
					</p>
				</div>
			</div>
		</div>
	</section>

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
