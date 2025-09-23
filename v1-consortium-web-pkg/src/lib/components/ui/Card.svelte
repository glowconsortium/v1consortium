<script lang="ts">
	interface Props {
		variant?: 'filled' | 'outline' | 'soft' | 'ghost';
		color?: 'primary' | 'secondary' | 'tertiary' | 'success' | 'warning' | 'error' | 'surface';
		hover?: boolean;
		padding?: 'sm' | 'md' | 'lg' | 'xl';
		base?: string;
		background?: string;
		border?: string;
		rounded?: string;
		shadow?: string;
		classes?: string;
		
		header?: import('svelte').Snippet;
		footer?: import('svelte').Snippet;
		children?: import('svelte').Snippet;
	}

	let {
		variant = 'filled',
		color = 'surface',
		hover = false,
		padding = 'md',
		base = '',
		background = '',
		border = '',
		rounded = '',
		shadow = '',
		classes = '',
		header,
		footer,
		children
	}: Props = $props();

	let cardClass = $derived(() => {
		const baseClasses = base || 'card overflow-hidden transition-all duration-200';
		
		// Variant styles using Skeleton's color system
		const variantStyles = {
			filled: {
				bg: background || `bg-${color}-100 dark:bg-${color}-900/20`,
				border: border || `border border-${color}-200 dark:border-${color}-700`
			},
			outline: {
				bg: background || 'bg-transparent',
				border: border || `border border-${color}-300 dark:border-${color}-600`
			},
			soft: {
				bg: background || `bg-${color}-50 dark:bg-${color}-950/50`,
				border: border || 'border border-transparent'
			},
			ghost: {
				bg: background || 'bg-transparent',
				border: border || 'border border-transparent'
			}
		};

		const style = variantStyles[variant];
		
		// Padding classes
		const paddingClasses = {
			sm: 'p-3',
			md: 'p-4',
			lg: 'p-6',
			xl: 'p-8'
		};

		const paddingClass = paddingClasses[padding];
		const roundedClass = rounded || 'rounded-lg';
		const shadowClass = shadow || (variant === 'filled' ? 'shadow-sm hover:shadow-md' : '');
		const hoverClass = hover ? 'hover:scale-[1.02] hover:shadow-lg cursor-pointer' : '';

		return [
			baseClasses,
			style.bg,
			style.border,
			paddingClass,
			roundedClass,
			shadowClass,
			hoverClass,
			classes
		]
			.filter(Boolean)
			.join(' ');
	});

	let headerClass = $derived(() => {
		const paddingClasses = {
			sm: 'px-3 py-2',
			md: 'px-4 py-3',
			lg: 'px-6 py-4',
			xl: 'px-8 py-5'
		};
		
		return `${paddingClasses[padding]} border-b border-surface-200 dark:border-surface-700 bg-surface-50 dark:bg-surface-800/50`;
	});

	let footerClass = $derived(() => {
		const paddingClasses = {
			sm: 'px-3 py-2',
			md: 'px-4 py-3',
			lg: 'px-6 py-4',
			xl: 'px-8 py-5'
		};
		
		return `${paddingClasses[padding]} border-t border-surface-200 dark:border-surface-700 bg-surface-50 dark:bg-surface-800/50`;
	});

	let contentClass = $derived(() => {
		const paddingClasses = {
			sm: 'p-3',
			md: 'p-4',
			lg: 'p-6',
			xl: 'p-8'
		};
		
		return paddingClasses[padding];
	});
</script>

<article class={cardClass()}>
	{#if header}
		<header class={headerClass()}>
			{@render header()}
		</header>
	{/if}

	{#if children}
		<div class={header || footer ? contentClass() : ''}>
			{@render children()}
		</div>
	{/if}

	{#if footer}
		<footer class={footerClass()}>
			{@render footer()}
		</footer>
	{/if}
</article>
