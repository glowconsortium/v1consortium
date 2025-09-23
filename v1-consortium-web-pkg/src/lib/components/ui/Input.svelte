
<script lang="ts">
	interface Props {
		type?: 'text' | 'email' | 'password' | 'number' | 'tel' | 'url' | 'date';
		placeholder?: string;
		value?: string | number;
		required?: boolean;
		disabled?: boolean;
		error?: string;
		label?: string;
		id?: string;
		base?: string;
		background?: string;
		border?: string;
		rounded?: string;
		padding?: string;
		classes?: string;
		labelClasses?: string;
		errorClasses?: string;
		oninput?: (event: Event) => void;
		onblur?: (event: Event) => void;
		onfocus?: (event: Event) => void;
	}

	let {
		type = 'text',
		placeholder = '',
		value = $bindable(''),
		required = false,
		disabled = false,
		error = '',
		label = '',
		id = '',
		base = '',
		background = '',
		border = '',
		rounded = '',
		padding = '',
		classes = '',
		labelClasses = '',
		errorClasses = '',
		oninput,
		onblur,
		onfocus
	}: Props = $props();

	let inputId = $derived(id || `input-${Math.random().toString(36).substring(2, 9)}`);
	
	let inputClass = $derived(() => {
		const baseClasses = base || 'input w-full transition-colors focus:outline-none focus:ring-2 focus:ring-primary-500';
		const backgroundClass = background || (error ? 'bg-error-50 dark:bg-error-950' : 'bg-surface-50 dark:bg-surface-900');
		const borderClass = border || (error ? 'border-error-500' : 'border-surface-300 dark:border-surface-600');
		const roundedClass = rounded || 'rounded-md';
		const paddingClass = padding || 'px-3 py-2';
		const disabledClass = disabled ? 'opacity-50 cursor-not-allowed' : '';
		
		return [baseClasses, backgroundClass, borderClass, roundedClass, paddingClass, disabledClass, classes]
			.filter(Boolean)
			.join(' ');
	});

	let labelClass = $derived(() => {
		const baseLabelClasses = 'label text-sm font-medium text-surface-700 dark:text-surface-300';
		return [baseLabelClasses, labelClasses].filter(Boolean).join(' ');
	});

	let errorClass = $derived(() => {
		const baseErrorClasses = 'text-error-500 text-sm mt-1 flex items-center gap-1';
		return [baseErrorClasses, errorClasses].filter(Boolean).join(' ');
	});
</script>

<div class="space-y-2">
	{#if label}
		<label for={inputId} class={labelClass()}>
			{label}
			{#if required}
				<span class="text-error-500 ml-1">*</span>
			{/if}
		</label>
	{/if}
	
	<input
		{type}
		{placeholder}
		bind:value
		{required}
		{disabled}
		id={inputId}
		class={inputClass()}
		{oninput}
		{onblur}
		{onfocus}
		aria-invalid={error ? 'true' : 'false'}
		aria-describedby={error ? `${inputId}-error` : undefined}
	/>
	
	{#if error}
		<div id="{inputId}-error" class={errorClass()} role="alert">
			<svg class="w-4 h-4 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
				<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
			</svg>
			<span>{error}</span>
		</div>
	{/if}
</div>
