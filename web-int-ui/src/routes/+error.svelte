<script lang="ts">
	import { page } from '$app/stores';
	
	let errorTitle = $derived(
		$page.status === 404 
			? "Page Not Found" 
			: "Something went wrong"
	);
	
	let errorMessage = $derived(
		$page.status === 404 
			? "The page you're looking for doesn't exist or has been moved."
			: "We encountered an unexpected error. Please try again later."
	);
	
	let errorCode = $derived($page.status || 500);
</script>

<div class="min-h-[80vh] flex items-center justify-center p-4">
	<div class="text-center space-y-6">
		<div class="flex flex-col items-center gap-4">
			<!-- Error Icon -->
			<div class="w-20 h-20 rounded-full bg-error-500/20 flex items-center justify-center">
				<span class="material-symbols-outlined text-4xl text-error-500">
					{errorCode === 404 ? 'search_off' : 'error'}
				</span>
			</div>
			
			<!-- Error Code -->
			<p class="text-8xl font-bold text-error-500">{errorCode}</p>
		</div>

		<!-- Error Details -->
		<div class="max-w-md space-y-2">
			<h1 class="h2">{errorTitle}</h1>
			<p class="text-surface-600-300-token">{errorMessage}</p>
			{#if $page.error?.message}
				<p class="text-sm text-surface-500-400-token italic">
					Error details: {$page.error.message}
				</p>
			{/if}
		</div>

		<!-- Action Buttons -->
		<div class="flex flex-col sm:flex-row gap-4 justify-center mt-8">
			<a 
				href="/"
				class="btn variant-filled"
			>
				Go to Home
			</a>
			<button 
				class="btn variant-ghost"
				onclick={() => history.back()}
			>
				Go Back
			</button>
		</div>
	</div>
</div> 