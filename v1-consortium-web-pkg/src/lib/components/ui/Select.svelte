<script lang="ts">
    interface Option {
        value: string;
        label: string;
        disabled?: boolean;
    }

    interface Props {
        value?: string;
        options: Option[];
        placeholder?: string;
        disabled?: boolean;
        required?: boolean;
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
        onchange?: (value: string) => void;
    }

    let {
        value = $bindable(''),
        options,
        placeholder = 'Select an option...',
        disabled = false,
        required = false,
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
        onchange
    }: Props = $props();

    let selectId = $derived(id || `select-${Math.random().toString(36).substring(2, 9)}`);

    let selectClass = $derived(() => {
        const baseClasses = base || 'select w-full transition-colors focus:outline-none focus:ring-2 focus:ring-primary-500';
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

    function handleChange(event: Event) {
        const target = event.target as HTMLSelectElement;
        value = target.value;
        onchange?.(value);
    }
</script>

<div class="space-y-2">
    {#if label}
        <label for={selectId} class={labelClass()}>
            {label}
            {#if required}
                <span class="text-error-500 ml-1">*</span>
            {/if}
        </label>
    {/if}

    <select
        id={selectId}
        bind:value
        {disabled}
        {required}
        class={selectClass()}
        onchange={handleChange}
        aria-invalid={error ? 'true' : 'false'}
        aria-describedby={error ? `${selectId}-error` : undefined}
    >
        {#if placeholder}
            <option value="" disabled>{placeholder}</option>
        {/if}
        
        {#each options as option}
            <option value={option.value} disabled={option.disabled}>
                {option.label}
            </option>
        {/each}
    </select>

    {#if error}
        <div id="{selectId}-error" class={errorClass()} role="alert">
            <svg class="w-4 h-4 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
            <span>{error}</span>
        </div>
    {/if}
</div>