
<script lang="ts">
    interface Props {
        checked?: boolean;
        disabled?: boolean;
        required?: boolean;
        indeterminate?: boolean;
        value?: string;
        label?: string;
        description?: string;
        error?: string;
        id?: string;
        name?: string;
        base?: string;
        background?: string;
        border?: string;
        rounded?: string;
        classes?: string;
        labelClasses?: string;
        descriptionClasses?: string;
        errorClasses?: string;
        'aria-label'?: string;
        'aria-describedby'?: string;
        onchange?: (checked: boolean) => void;
    }

    let {
        checked = $bindable(false),
        disabled = false,
        required = false,
        indeterminate = false,
        value = '',
        label = '',
        description = '',
        error = '',
        id = '',
        name = '',
        base = '',
        background = '',
        border = '',
        rounded = '',
        classes = '',
        labelClasses = '',
        descriptionClasses = '',
        errorClasses = '',
        'aria-label': ariaLabel = '',
        'aria-describedby': ariaDescribedby = '',
        onchange
    }: Props = $props();

    let checkboxElement: HTMLInputElement;
    let checkboxId = $derived(id || `checkbox-${Math.random().toString(36).substring(2, 9)}`);

    let checkboxClass = $derived(() => {
        const baseClasses = base || 'checkbox h-4 w-4 transition-colors focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2';
        const backgroundClass = background || (checked 
            ? 'bg-primary-600 border-primary-600' 
            : 'bg-surface-50 dark:bg-surface-900 border-surface-300 dark:border-surface-600'
        );
        const borderClass = border || '';
        const roundedClass = rounded || 'rounded';
        const disabledClass = disabled ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer';
        const errorClass = error ? 'border-error-500 focus:ring-error-500' : '';

        return [baseClasses, backgroundClass, borderClass, roundedClass, disabledClass, errorClass, classes]
            .filter(Boolean)
            .join(' ');
    });

    let labelClass = $derived(() => {
        const baseLabelClasses = 'text-sm font-medium text-surface-700 dark:text-surface-300';
        const disabledClass = disabled ? 'opacity-50' : 'cursor-pointer';
        return [baseLabelClasses, disabledClass, labelClasses].filter(Boolean).join(' ');
    });

    let descriptionClass = $derived(() => {
        const baseDescriptionClasses = 'text-sm text-surface-600 dark:text-surface-400';
        return [baseDescriptionClasses, descriptionClasses].filter(Boolean).join(' ');
    });

    let errorClass = $derived(() => {
        const baseErrorClasses = 'text-error-500 text-sm mt-1 flex items-center gap-1';
        return [baseErrorClasses, errorClasses].filter(Boolean).join(' ');
    });

    function handleChange(event: Event) {
        if (disabled) return;
        
        const target = event.target as HTMLInputElement;
        checked = target.checked;
        onchange?.(checked);
    }

    // Handle indeterminate state
    $effect(() => {
        if (checkboxElement) {
            checkboxElement.indeterminate = indeterminate;
        }
    });
</script>

<div class="flex items-start space-x-3">
    <div class="flex items-center h-5">
        <input
            bind:this={checkboxElement}
            type="checkbox"
            id={checkboxId}
            {name}
            {value}
            bind:checked
            {disabled}
            {required}
            class={checkboxClass()}
            onchange={handleChange}
            aria-label={ariaLabel}
            aria-describedby={ariaDescribedby || (error ? `${checkboxId}-error` : undefined)}
            aria-invalid={error ? 'true' : 'false'}
        />
    </div>

    {#if label || description}
        <div class="flex-1 min-w-0">
            {#if label}
                <label for={checkboxId} class={labelClass()}>
                    {label}
                    {#if required}
                        <span class="text-error-500 ml-1">*</span>
                    {/if}
                </label>
            {/if}
            
            {#if description}
                <p class={descriptionClass()}>
                    {description}
                </p>
            {/if}
        </div>
    {/if}
</div>

{#if error}
    <div id="{checkboxId}-error" class={errorClass()} role="alert">
        <svg class="w-4 h-4 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
        </svg>
        <span>{error}</span>
    </div>
{/if}