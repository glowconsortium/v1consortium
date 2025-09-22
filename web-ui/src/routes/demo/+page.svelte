<script lang="ts">
	import { Button, Input, Alert, Spinner, Card } from '$lib/components/ui';
	
	let formData = $state({
		name: '',
		email: '',
		password: ''
	});
	
	let showAlert = $state(false);
	let alertType = $state<'success' | 'warning' | 'error' | 'info'>('success');
	let isLoading = $state(false);

	function handleSubmit() {
		isLoading = true;
		setTimeout(() => {
			isLoading = false;
			showAlert = true;
			alertType = 'success';
		}, 2000);
	}

	function showAlertType(type: typeof alertType) {
		alertType = type;
		showAlert = true;
	}
</script>

<svelte:head>
	<title>Component Demo - Skeleton UI Extensions</title>
</svelte:head>

<div class="container mx-auto px-4 py-8 max-w-4xl">
	<h1 class="text-4xl font-bold text-surface-900 dark:text-surface-100 mb-8">
		Skeleton UI Component Demo
	</h1>
	
	<div class="space-y-12">
		<!-- Button Components -->
		<section class="space-y-6">
			<h2 class="text-2xl font-semibold text-surface-800 dark:text-surface-200">Buttons</h2>
			
			<!-- Button Variants -->
			<div class="space-y-4">
				<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">Variants</h3>
				<div class="flex flex-wrap gap-3">
					<Button variant="filled" color="primary">Filled Primary</Button>
					<Button variant="outline" color="primary">Outline Primary</Button>
					<Button variant="soft" color="primary">Soft Primary</Button>
					<Button variant="ghost" color="primary">Ghost Primary</Button>
					<Button variant="ring" color="primary">Ring Primary</Button>
				</div>
			</div>

			<!-- Button Colors -->
			<div class="space-y-4">
				<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">Colors</h3>
				<div class="flex flex-wrap gap-3">
					<Button color="primary">Primary</Button>
					<Button color="secondary">Secondary</Button>
					<Button color="tertiary">Tertiary</Button>
					<Button color="success">Success</Button>
					<Button color="warning">Warning</Button>
					<Button color="error">Error</Button>
					<Button color="surface">Surface</Button>
				</div>
			</div>

			<!-- Button Sizes -->
			<div class="space-y-4">
				<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">Sizes</h3>
				<div class="flex flex-wrap items-center gap-3">
					<Button size="sm">Small</Button>
					<Button size="md">Medium</Button>
					<Button size="lg">Large</Button>
				</div>
			</div>

			<!-- Button States -->
			<div class="space-y-4">
				<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">States</h3>
				<div class="flex flex-wrap gap-3">
					<Button disabled>Disabled</Button>
					<Button loading>Loading</Button>
					<Button onclick={() => isLoading = !isLoading}>
						Toggle Loading
					</Button>
				</div>
			</div>
		</section>

		<!-- Input Components -->
		<section class="space-y-6">
			<h2 class="text-2xl font-semibold text-surface-800 dark:text-surface-200">Inputs</h2>
			
			<div class="grid md:grid-cols-2 gap-6">
				<div class="space-y-4">
					<Input 
						label="Name" 
						placeholder="Enter your name" 
						bind:value={formData.name}
						required
					/>
					
					<Input 
						type="email"
						label="Email" 
						placeholder="Enter your email" 
						bind:value={formData.email}
						required
					/>
					
					<Input 
						type="password"
						label="Password" 
						placeholder="Enter your password" 
						bind:value={formData.password}
						required
					/>
					
					<Input 
						label="Disabled Input" 
						placeholder="This is disabled" 
						disabled
					/>
				</div>
				
				<div class="space-y-4">
					<Input 
						label="Input with Error" 
						placeholder="This has an error" 
						error="This field is required"
					/>
					
					<Input 
						type="number"
						label="Number Input" 
						placeholder="Enter a number"
					/>
					
					<Input 
						type="tel"
						label="Phone Number" 
						placeholder="+1 (555) 123-4567"
					/>
					
					<Input 
						type="url"
						label="Website URL" 
						placeholder="https://example.com"
					/>
				</div>
			</div>
			
			<div class="flex gap-3">
				<Button onclick={handleSubmit} loading={isLoading}>
					{isLoading ? 'Submitting...' : 'Submit Form'}
				</Button>
				<Button variant="outline" onclick={() => {
					formData.name = '';
					formData.email = '';
					formData.password = '';
				}}>
					Clear Form
				</Button>
			</div>
		</section>

		<!-- Alert Components -->
		<section class="space-y-6">
			<h2 class="text-2xl font-semibold text-surface-800 dark:text-surface-200">Alerts</h2>
			
			<div class="space-y-4">
				<div class="flex flex-wrap gap-3 mb-4">
					<Button variant="soft" color="success" onclick={() => showAlertType('success')}>
						Show Success
					</Button>
					<Button variant="soft" color="warning" onclick={() => showAlertType('warning')}>
						Show Warning
					</Button>
					<Button variant="soft" color="error" onclick={() => showAlertType('error')}>
						Show Error
					</Button>
					<Button variant="soft" color="primary" onclick={() => showAlertType('info')}>
						Show Info
					</Button>
				</div>

				{#if showAlert}
					<Alert 
						type={alertType}
						title={alertType.charAt(0).toUpperCase() + alertType.slice(1)}
						message={`This is a ${alertType} alert message with additional context and information.`}
						ondismiss={() => showAlert = false}
					/>
				{/if}

				<!-- Static alerts for demonstration -->
				<Alert 
					type="success"
					title="Success!"
					message="Your form has been submitted successfully."
					dismissible={false}
				/>
				
				<Alert 
					type="warning"
					message="Please review your information before proceeding."
				/>
				
				<Alert 
					type="error"
					title="Error"
					message="There was an error processing your request. Please try again."
				/>
				
				<Alert 
					type="info"
					message="This is some helpful information for the user."
				/>
			</div>
		</section>

		<!-- Spinner Components -->
		<section class="space-y-6">
			<h2 class="text-2xl font-semibold text-surface-800 dark:text-surface-200">Spinners</h2>
			
			<div class="space-y-6">
				<!-- Spinner Variants -->
				<div class="space-y-4">
					<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">Variants</h3>
					<div class="flex items-center gap-8">
						<div class="text-center space-y-2">
							<Spinner variant="circular" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Circular</p>
						</div>
						<div class="text-center space-y-2">
							<Spinner variant="dots" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Dots</p>
						</div>
						<div class="text-center space-y-2">
							<Spinner variant="pulse" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Pulse</p>
						</div>
					</div>
				</div>

				<!-- Spinner Sizes -->
				<div class="space-y-4">
					<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">Sizes</h3>
					<div class="flex items-center gap-8">
						<div class="text-center space-y-2">
							<Spinner size="sm" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Small</p>
						</div>
						<div class="text-center space-y-2">
							<Spinner size="md" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Medium</p>
						</div>
						<div class="text-center space-y-2">
							<Spinner size="lg" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Large</p>
						</div>
						<div class="text-center space-y-2">
							<Spinner size="xl" />
							<p class="text-sm text-surface-600 dark:text-surface-400">Extra Large</p>
						</div>
					</div>
				</div>

				<!-- Spinner Colors -->
				<div class="space-y-4">
					<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300">Colors</h3>
					<div class="flex items-center gap-6">
						<Spinner color="primary" />
						<Spinner color="secondary" />
						<Spinner color="success" />
						<Spinner color="warning" />
						<Spinner color="error" />
					</div>
				</div>
			</div>
		</section>

		<!-- Card Components -->
		<section class="space-y-6">
			<h2 class="text-2xl font-semibold text-surface-800 dark:text-surface-200">Cards</h2>
			
			<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
				<!-- Basic Card -->
				<Card>
					{#snippet header()}
						<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-100">Basic Card</h3>
					{/snippet}
					
					<p class="text-surface-700 dark:text-surface-300">
						This is a basic card with header content. Perfect for displaying information in a structured way.
					</p>
					
					{#snippet footer()}
						<div class="flex justify-end">
							<Button size="sm" variant="outline">Action</Button>
						</div>
					{/snippet}
				</Card>

				<!-- Primary Card with hover -->
				<Card variant="filled" color="primary" hover>
					{#snippet header()}
						<h3 class="text-lg font-semibold">Primary Card</h3>
					{/snippet}
					
					<p class="text-primary-800 dark:text-primary-200">
						A primary colored card with hover effects. Try hovering over this card to see the animation.
					</p>
				</Card>

				<!-- Outline Card -->
				<Card variant="outline" color="secondary">
					<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-100 mb-3">Outline Card</h3>
					<p class="text-surface-700 dark:text-surface-300">
						An outline variant card without header/footer structure. Clean and minimal.
					</p>
				</Card>

				<!-- Success Card -->
				<Card variant="soft" color="success" padding="lg">
					{#snippet header()}
						<div class="flex items-center gap-2">
							<svg class="w-5 h-5 text-success-500" fill="currentColor" viewBox="0 0 20 20">
								<path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
							</svg>
							<h3 class="text-lg font-semibold text-success-800 dark:text-success-200">Success</h3>
						</div>
					{/snippet}
					
					<p class="text-success-700 dark:text-success-300">
						Everything is working perfectly! Your components are ready to use.
					</p>
				</Card>

				<!-- Interactive Card -->
				<Card variant="ghost" hover classes="border-2 border-dashed border-surface-300 dark:border-surface-600 hover:border-primary-400">
					<div class="text-center py-8">
						<svg class="w-12 h-12 mx-auto text-surface-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
						</svg>
						<h3 class="text-lg font-medium text-surface-700 dark:text-surface-300 mb-2">Add New Item</h3>
						<p class="text-surface-500 text-sm">Click to add something new</p>
					</div>
				</Card>

				<!-- Form Card -->
				<Card variant="outline" padding="lg">
					{#snippet header()}
						<h3 class="text-lg font-semibold text-surface-900 dark:text-surface-100">Quick Form</h3>
					{/snippet}
					
					<div class="space-y-4">
						<Input 
							label="Name" 
							placeholder="Enter name"
						/>
						<Input 
							type="email"
							label="Email" 
							placeholder="Enter email"
						/>
					</div>
					
					{#snippet footer()}
						<div class="flex justify-between">
							<Button variant="ghost" size="sm">Cancel</Button>
							<Button size="sm">Submit</Button>
						</div>
					{/snippet}
				</Card>
			</div>
		</section>

		<!-- Theme Demonstration -->
		<section class="space-y-6">
			<h2 class="text-2xl font-semibold text-surface-800 dark:text-surface-200">Theme Integration</h2>
			<div class="card p-6 bg-surface-100 dark:bg-surface-800 rounded-lg">
				<p class="text-surface-700 dark:text-surface-300 mb-4">
					All components automatically adapt to Skeleton's theme system and respond to light/dark mode changes.
					They use Skeleton's color palette and support color pairings for consistent theming.
				</p>
				<div class="flex gap-3">
					<Button variant="filled" color="primary">Primary Button</Button>
					<Button variant="outline" color="secondary">Secondary Outline</Button>
					<Button variant="soft" color="success">Success Soft</Button>
				</div>
			</div>
		</section>
	</div>
</div>

<a href="/demo/paraglide" class="fixed bottom-4 right-4 btn variant-outline-primary">
	Paraglide Demo →
</a>


<button class ="btn preset-outlined-primary-500">Test Outlined</button>