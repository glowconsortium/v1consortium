<script lang="ts">
	interface Props {
		size?: 'sm' | 'md' | 'lg' | 'xl';
		color?: 'primary' | 'secondary' | 'tertiary' | 'success' | 'warning' | 'error' | 'surface';
		variant?: 'circular' | 'dots' | 'pulse';
		base?: string;
		classes?: string;
	}

	let {
		size = 'md',
		color = 'primary',
		variant = 'circular',
		base = '',
		classes = ''
	}: Props = $props();

	let spinnerClass = $derived(() => {
		const sizeClasses = {
			sm: 'w-4 h-4',
			md: 'w-6 h-6',
			lg: 'w-8 h-8',
			xl: 'w-12 h-12'
		};

		const colorClasses = {
			primary: 'text-primary-500',
			secondary: 'text-secondary-500',
			tertiary: 'text-tertiary-500',
			success: 'text-success-500',
			warning: 'text-warning-500',
			error: 'text-error-500',
			surface: 'text-surface-500'
		};

		if (variant === 'circular') {
			const baseClasses = base || 'animate-spin';
			return [baseClasses, sizeClasses[size], colorClasses[color], classes]
				.filter(Boolean)
				.join(' ');
		} else if (variant === 'dots') {
			const baseClasses = base || 'flex space-x-1';
			return [baseClasses, classes].filter(Boolean).join(' ');
		} else if (variant === 'pulse') {
			const baseClasses = base || 'animate-pulse rounded-full';
			return [baseClasses, sizeClasses[size], `bg-${color}-500`, classes]
				.filter(Boolean)
				.join(' ');
		}

		return classes;
	});

	let dotClass = $derived(() => {
		const sizeClasses = {
			sm: 'w-1 h-1',
			md: 'w-1.5 h-1.5',
			lg: 'w-2 h-2',
			xl: 'w-3 h-3'
		};
		
		return `${sizeClasses[size]} bg-${color}-500 rounded-full animate-bounce`;
	});
</script>

{#if variant === 'circular'}
	<svg class={spinnerClass()} fill="none" viewBox="0 0 24 24">
		<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
		<path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
	</svg>
{:else if variant === 'dots'}
	<div class={spinnerClass()}>
		<div class={dotClass()} style="animation-delay: -0.32s"></div>
		<div class={dotClass()} style="animation-delay: -0.16s"></div>
		<div class={dotClass()}></div>
	</div>
{:else if variant === 'pulse'}
	<div class={spinnerClass()}></div>
{/if}
