<script lang="ts">
	interface Props {
		type?: 'success' | 'warning' | 'error' | 'info';
		message: string;
		title?: string;
		dismissible?: boolean;
		base?: string;
		background?: string;
		border?: string;
		rounded?: string;
		padding?: string;
		classes?: string;
		ondismiss?: () => void;
		    children?: import('svelte').Snippet;
	}

	let {
		children,
		type = 'info',
		message,
		title = '',
		dismissible = true,
		base = '',
		background = '',
		border = '',
		rounded = '',
		padding = '',
		classes = '',
		
		ondismiss
	}: Props = $props();

	let visible = $state(true);

	function dismiss() {
		visible = false;
		ondismiss?.();
	}

	let alertClass = $derived(() => {
		const baseClasses = base || 'alert flex items-start gap-4 p-4 transition-all duration-200';
		
		// Background and border based on type using Skeleton's color system
		const typeStyles = {
			success: {
				bg: background || 'bg-success-100 dark:bg-success-900/20',
				border: border || 'border-success-500/20',
				text: 'text-success-800 dark:text-success-200'
			},
			warning: {
				bg: background || 'bg-warning-100 dark:bg-warning-900/20',
				border: border || 'border-warning-500/20',
				text: 'text-warning-800 dark:text-warning-200'
			},
			error: {
				bg: background || 'bg-error-100 dark:bg-error-900/20',
				border: border || 'border-error-500/20',
				text: 'text-error-800 dark:text-error-200'
			},
			info: {
				bg: background || 'bg-primary-100 dark:bg-primary-900/20',
				border: border || 'border-primary-500/20',
				text: 'text-primary-800 dark:text-primary-200'
			}
		};

		const style = typeStyles[type];
		const roundedClass = rounded || 'rounded-md';
		const borderClass = `border ${style.border}`;
		
		return [baseClasses, style.bg, borderClass, roundedClass, style.text, classes]
			.filter(Boolean)
			.join(' ');
	});

	let iconContent = $derived(() => {
		switch (type) {
			case 'success':
				return {
					svg: `<svg class="w-5 h-5 text-success-500" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
					</svg>`,
					color: 'text-success-500'
				};
			case 'warning':
				return {
					svg: `<svg class="w-5 h-5 text-warning-500" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
					</svg>`,
					color: 'text-warning-500'
				};
			case 'error':
				return {
					svg: `<svg class="w-5 h-5 text-error-500" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
					</svg>`,
					color: 'text-error-500'
				};
			case 'info':
			default:
				return {
					svg: `<svg class="w-5 h-5 text-primary-500" fill="currentColor" viewBox="0 0 20 20">
						<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
					</svg>`,
					color: 'text-primary-500'
				};
		}
	});
</script>

{#if visible}
	<div class={alertClass()} role="alert">
		<!-- Icon -->
		<div class="flex-shrink-0">
			{@html iconContent().svg}
		</div>
		
		<!-- Content -->
		<div class="flex-1 min-w-0">
			{#if title}
				<h4 class="text-sm font-semibold mb-1">{title}</h4>
			{/if}
			<p class="text-sm">{message}</p>
			{#if children}
				<div class="mt-2">
					{@render children()}
				</div>
			{/if}
		</div>
		
		<!-- Dismiss button -->
		{#if dismissible}
			<button
				type="button"
				class="btn-sm variant-ghost-surface -m-1.5 p-1.5 hover:bg-surface-200 dark:hover:bg-surface-700 transition-colors"
				onclick={dismiss}
				aria-label="Dismiss alert"
			>
				<svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
					<path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
				</svg>
			</button>
		{/if}
	</div>
{/if}
