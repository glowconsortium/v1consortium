<script lang="ts">
	interface ErrorDetails {
		message: string;
		stack?: string;
	}

	let error = $state<ErrorDetails | null>(null);
	let { children } = $props();

	function handleError(event: ErrorEvent) {
		error = {
			message: event.error.message || 'An unexpected error occurred',
			stack: event.error.stack
		};
	}

	$effect.root(() => {
		window.addEventListener('error', handleError);
		return () => window.removeEventListener('error', handleError);
	});
</script>

{#if error}
	<div class="rounded-container-token bg-error-500/10 border-error-500/20 border p-4">
		<div class="flex items-start gap-4">
			<div
				class="bg-error-500/20 flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-full"
			>
				<span class="material-symbols-outlined text-error-500">warning</span>
			</div>
			<div class="flex-1 space-y-2">
				<h2 class="h3 text-error-500">Something went wrong</h2>
				<p class="text-surface-600-300-token">{error.message}</p>
				{#if error.stack && import.meta.env.DEV}
					<details class="mt-4">
						<summary class="text-surface-500-400-token cursor-pointer text-sm">
							View error details
						</summary>
						<pre
							class="rounded-container-token bg-surface-100-800-token mt-2 overflow-x-auto p-4 text-xs">
							{error.stack}
						</pre>
					</details>
				{/if}
				<div class="mt-4 flex gap-2">
					<button class="btn variant-soft-error" onclick={() => window.location.reload()}>
						Reload Page
					</button>
					<button class="btn variant-ghost" onclick={() => (error = null)}> Dismiss </button>
				</div>
			</div>
		</div>
	</div>
{:else}
	{@render children()}
{/if}
