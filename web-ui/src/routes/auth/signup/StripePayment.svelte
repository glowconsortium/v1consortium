<script lang="ts">
	import { Button, Alert } from '$lib/components/ui/index.js';

    let {success, error, cancel, selectedPlan, companyName, email, isProcessing}:{
         success: (paymentMethodId: string, customerId: string) => void;
          error: (message: string) => void;
          cancel: () => void;
          selectedPlan: 'basic' | 'premium';
          companyName: string;
            email: string;
            isProcessing: boolean;
        } = $props()

	// Stripe variables
	let stripe: any = null;
	let elements: any = null;
	let cardElement: any = null;
	let stripeError = $state('');
	let isStripeLoaded = $state(false);
	let paymentProcessing = $state(false);
	let cardElementContainer: HTMLElement;

	// Plan details
	const plans = {
		basic: {
			name: 'Basic Plan',
			price: 99,
			annualPrice: 1188, // $99 * 12
			features: [
				'Up to 50 drivers',
				'Drug & alcohol testing',
				'DOT physicals',
				'MVR checks',
				'Basic reporting',
				'Email support'
			]
		},
		premium: {
			name: 'Premium Plan',
			price: 199,
			annualPrice: 2388, // $199 * 12
			features: [
				'Up to 500 drivers',
				'All Basic features',
				'Advanced analytics',
				'Custom reporting',
				'API integrations',
				'Priority support',
				'Dedicated success manager'
			]
		}
	};

	const currentPlan = $derived(plans[selectedPlan]);
	const annualDiscount = $derived(Math.round((currentPlan.price * 12 - currentPlan.annualPrice) / (currentPlan.price * 12) * 100));
	
	// Helper to determine if we're in demo mode
	const isDemoMode = $derived(!import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY || (!stripe && isStripeLoaded));

	// Initialize Stripe when the component and DOM are ready
	$effect(() => {
		// Wait for card element container to be available
		if (cardElementContainer && !isStripeLoaded && !stripe) {
			initializeStripe();
		}
		// If no Stripe key is set, enable demo mode immediately
		else if (!import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY && !isStripeLoaded) {
			console.warn('⚠️ VITE_STRIPE_PUBLISHABLE_KEY not set. Using demo mode.');
			isStripeLoaded = true;
		}
	});

	async function loadStripeScript(): Promise<void> {
		return new Promise((resolve, reject) => {
			if ((window as any).Stripe) {
				resolve();
				return;
			}

			const script = document.createElement('script');
			script.src = 'https://js.stripe.com/v3/';
			script.onload = () => resolve();
			script.onerror = () => reject(new Error('Failed to load Stripe script'));
			document.head.appendChild(script);
		});
	}

	async function initializeStripe() {
		try {
			// Check if we should use demo mode
			const stripeKey = import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY;
			
			// For development, provide a helpful message if no key is set
			if (!stripeKey) {
				console.warn('⚠️ VITE_STRIPE_PUBLISHABLE_KEY not set. Using demo mode.');
				// In demo mode, we'll simulate payment without actually processing
				isStripeLoaded = true;
				return;
			}

			// Ensure Stripe script is loaded
			await loadStripeScript();

			// Check if container exists
			if (!cardElementContainer) {
				throw new Error('Card element container not found');
			}
			
			stripe = (window as any).Stripe(stripeKey);
			
			if (!stripe) {
				throw new Error('Failed to initialize Stripe');
			}
			
			elements = stripe.elements({
				appearance: {
					theme: 'stripe',
					variables: {
						colorPrimary: '#2563eb',
						colorBackground: '#ffffff',
						colorText: '#374151',
						colorDanger: '#dc2626',
						fontFamily: 'system-ui, sans-serif',
						spacingUnit: '4px',
						borderRadius: '6px'
					}
				}
			});

			cardElement = elements.create('card', {
				style: {
					base: {
						fontSize: '16px',
						color: '#374151',
						'::placeholder': {
							color: '#9ca3af'
						}
					}
				}
			});

			// Mount to the container directly
			cardElement.mount(cardElementContainer);
			
			cardElement.on('change', (event: any) => {
				if (event.error) {
					stripeError = event.error.message;
				} else {
					stripeError = '';
				}
			});

			isStripeLoaded = true;
		} catch (err) {
			console.error('Stripe initialization failed:', err);
			// Fallback to demo mode if Stripe fails to initialize
			console.warn('🎭 Falling back to demo mode due to Stripe initialization failure');
			stripeError = '';
			isStripeLoaded = true;
		}
	}

	async function handlePayment() {
		paymentProcessing = true;
		stripeError = '';

		try {
			// Demo mode (when no Stripe key is configured or Stripe failed to initialize)
			if (isDemoMode) {
				console.log('🎭 Demo mode: Simulating payment...');
				await new Promise(resolve => setTimeout(resolve, 2000));
				success('pm_demo_' + Date.now(), 'cus_demo_' + Date.now());
				return;
			}

			// Real Stripe processing
			if (!stripe || !cardElement) {
				stripeError = 'Payment system not ready. Please wait a moment and try again.';
				return;
			}

			// Create payment method
			const { error: paymentMethodError, paymentMethod } = await stripe.createPaymentMethod({
				type: 'card',
				card: cardElement,
				billing_details: {
					name: companyName || 'Customer',
					email: email
				}
			});

			if (paymentMethodError) {
				throw new Error(paymentMethodError.message);
			}

			// In a real implementation, you would:
			// 1. Send payment method to your backend
			// 2. Create customer and subscription
			// 3. Handle payment confirmation if needed

			// For now, simulate success with real payment method
			await new Promise(resolve => setTimeout(resolve, 1000));
            success(paymentMethod.id, 'cus_simulated_' + Date.now());

		} catch (error: any) {
			stripeError = error.message || 'Payment failed. Please try again.';
            error(stripeError);
		} finally {
			paymentProcessing = false;
		}
	}

	function handleCancel() {
		//dispatch('cancel');
        cancel();
	}
</script>

<div class="max-w-md mx-auto">
	<!-- Plan Summary -->
	<div class="bg-gray-50 rounded-lg p-6 mb-6">
		<div class="flex items-center justify-between mb-4">
			<h3 class="text-lg font-semibold text-gray-900">{currentPlan.name}</h3>
			<div class="text-right">
				<div class="text-2xl font-bold text-gray-900">${currentPlan.annualPrice}/year</div>
				<div class="text-sm text-gray-500 line-through">${currentPlan.price * 12}/year</div>
				<div class="text-sm text-green-600 font-medium">Save {annualDiscount}%</div>
			</div>
		</div>
		
		<div class="space-y-2">
			{#each currentPlan.features as feature}
				<div class="flex items-center text-sm text-gray-600">
					<svg class="w-4 h-4 text-green-500 mr-2" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
					</svg>
					{feature}
				</div>
			{/each}
		</div>

		<div class="mt-4 pt-4 border-t border-gray-200">
			<div class="flex justify-between text-sm text-gray-600 mb-1">
				<span>Billing cycle:</span>
				<span>Annual (January - December)</span>
			</div>
			<div class="flex justify-between text-sm text-gray-600 mb-1">
				<span>Next billing:</span>
				<span>January 1, 2026</span>
			</div>
			<div class="flex justify-between text-sm text-gray-600">
				<span>Payment method:</span>
				<span>Credit card</span>
			</div>
		</div>
	</div>

	<!-- Payment Form -->
	<div class="bg-white border border-gray-200 rounded-lg p-6">
		<h4 class="text-lg font-medium text-gray-900 mb-4">Payment Information</h4>
		
		{#if !isStripeLoaded}
			<div class="flex items-center justify-center py-8">
				<div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
				<span class="ml-3 text-gray-600">Loading payment form...</span>
			</div>
		{:else if isDemoMode}
			<!-- Demo mode payment form -->
			<div class="space-y-4">
				<div class="bg-blue-50 border border-blue-200 rounded-md p-4">
					<div class="flex">
						<svg class="w-5 h-5 text-blue-400 mr-2" fill="currentColor" viewBox="0 0 20 20">
							<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
						</svg>
						<div class="text-sm text-blue-800">
							<p class="font-medium">Demo Mode</p>
							<p>Stripe payment processing is not configured or unavailable. This is a demonstration of the payment flow.</p>
						</div>
					</div>
				</div>
				
				<!-- Mock payment form -->
				<div>
					<label for="card-info-demo" class="block text-sm font-medium text-gray-700 mb-2">
						Card Information (Demo)
					</label>
					<div class="p-3 border border-gray-300 rounded-md bg-gray-50 min-h-[44px] flex items-center">
						<span class="text-gray-500 italic">Demo: Payment processing will be simulated</span>
					</div>
				</div>
			</div>
		{:else}
			<div class="space-y-4">
				<!-- Card Element -->
				<div>
					<label for="card-info" class="block text-sm font-medium text-gray-700 mb-2">
						Card Information
					</label>
					<div bind:this={cardElementContainer} class="p-3 border border-gray-300 rounded-md bg-white min-h-[44px]"></div>
				</div>

				<!-- Security notice -->
				<div class="bg-blue-50 border border-blue-200 rounded-md p-3">
					<div class="flex items-start">
						<svg class="w-5 h-5 text-blue-500 mr-2 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
							<path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd"></path>
						</svg>
						<div class="text-sm text-blue-700">
							<p class="font-medium">Secure Payment</p>
							<p>Your payment information is encrypted and secure. We use Stripe for payment processing.</p>
						</div>
					</div>
				</div>

				<!-- Billing terms -->
				<div class="text-xs text-gray-500 space-y-1">
					<p>• Your subscription will begin immediately upon payment</p>
					<p>• Annual billing cycle runs January 1 - December 31</p>
					<p>• Subscription automatically renews unless cancelled</p>
					<p>• 30-day money-back guarantee</p>
				</div>
			</div>
		{/if}

		<!-- Error Display -->
		{#if stripeError}
			<div class="mt-4">
				<Alert type="error" message={stripeError} />
			</div>
		{/if}

		<!-- Action Buttons -->
		<div class="flex space-x-3 mt-6">
			<Button
				variant="outline"
				size="lg"
				classes="flex-1"
				onclick={handleCancel}
				disabled={paymentProcessing}
			>
				Back
			</Button>
			
			<Button
				color="primary"
				size="lg"
				classes="flex-1"
				onclick={handlePayment}
				disabled={!isStripeLoaded || paymentProcessing}
			>
				{#if paymentProcessing}
					<svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
						<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
						<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
					</svg>
					{#if isDemoMode}
						Simulating Payment...
					{:else}
						Processing...
					{/if}
				{:else}
					{#if isDemoMode}
						Demo Pay ${currentPlan.annualPrice}
					{:else}
						Pay ${currentPlan.annualPrice}
					{/if}
				{/if}
			</Button>
		</div>

		<!-- Trust indicators -->
		<div class="mt-6 pt-4 border-t border-gray-200">
			<div class="flex items-center justify-center space-x-6 text-xs text-gray-500">
				<div class="flex items-center">
					<svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd"></path>
					</svg>
					SSL Encrypted
				</div>
				<div class="flex items-center">
					<svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
					</svg>
					PCI Compliant
				</div>
				<div class="flex items-center">
					<svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
					</svg>
					Powered by Stripe
				</div>
			</div>
		</div>
	</div>
</div>
