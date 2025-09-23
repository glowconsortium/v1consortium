<script lang="ts">
    interface Props {
        value?: string;
        placeholder?: string;
        disabled?: boolean;
        readonly?: boolean;
        required?: boolean;
        rows?: number;
        cols?: number;
        maxlength?: number;
        minlength?: number;
        resize?: 'none' | 'both' | 'horizontal' | 'vertical';
        error?: string;
        label?: string;
        description?: string;
        id?: string;
        name?: string;
        base?: string;
        background?: string;
        border?: string;
        rounded?: string;
        padding?: string;
        classes?: string;
        labelClasses?: string;
        descriptionClasses?: string;
        errorClasses?: string;
        'aria-label'?: string;
        'aria-describedby'?: string;
        oninput?: (value: string) => void;
        onchange?: (value: string) => void;
        onfocus?: (event: FocusEvent) => void;
        onblur?: (event: FocusEvent) => void;
    }

    let {
        value = $bindable(''),
        placeholder = '',
        disabled = false,
        readonly = false,
        required = false,
        rows = 3,
        cols,
        maxlength,
        minlength,
        resize = 'vertical',
        error = '',
        label = '',
        description = '',
        id = '',
        name = '',
        base = '',
        background = '',
        border = '',
        rounded = '',
        padding = '',
        classes = '',
        labelClasses = '',
        descriptionClasses = '',
        errorClasses = '',
        'aria-label': ariaLabel = '',
        'aria-describedby': ariaDescribedby = '',
        oninput,
        onchange,
        onfocus,
        onblur
    }: Props = $props();

    let textareaId = $derived(id || `textarea-${Math.random().toString(36).substring(2, 9)}`);

    let textareaClass = $derived(() => {
        const baseClasses = base || 'textarea w-full transition-colors focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent';
        const backgroundClass = background || (error 
            ? 'bg-error-50 dark:bg-error-950' 
            : 'bg-surface-50 dark:bg-surface-900'
        );
        const borderClass = border || (error 
            ? 'border-error-500' 
            : 'border-surface-300 dark:border-surface-600'
        );
        const roundedClass = rounded || 'rounded-md';
        const paddingClass = padding || 'px-3 py-2';
        const disabledClass = disabled ? 'opacity-50 cursor-not-allowed' : '';
        const readonlyClass = readonly ? 'cursor-default' : '';
        const resizeClass = resize ? `resize-${resize}` : '';

        return [
            baseClasses, 
            backgroundClass, 
            borderClass, 
            roundedClass, 
            paddingClass, 
            disabledClass, 
            readonlyClass,
            resizeClass,
            classes
        ].filter(Boolean).join(' ');
    });

    let labelClass = $derived(() => {
        const baseLabelClasses = 'label text-sm font-medium text-surface-700 dark:text-surface-300';
        return [baseLabelClasses, labelClasses].filter(Boolean).join(' ');
    });

    let descriptionClass = $derived(() => {
        const baseDescriptionClasses = 'text-sm text-surface-600 dark:text-surface-400 mt-1';
        return [baseDescriptionClasses, descriptionClasses].filter(Boolean).join(' ');
    });

    let errorClass = $derived(() => {
        const baseErrorClasses = 'text-error-500 text-sm mt-1 flex items-center gap-1';
        return [baseErrorClasses, errorClasses].filter(Boolean).join(' ');
    });

    function handleInput(event: Event) {
        const target = event.target as HTMLTextAreaElement;
        value = target.value;
        oninput?.(value);
    }

    function handleChange(event: Event) {
        const target = event.target as HTMLTextAreaElement;
        value = target.value;
        onchange?.(value);
    }

    function handleFocus(event: FocusEvent) {
        onfocus?.(event);
    }

    function handleBlur(event: FocusEvent) {
        onblur?.(event);
    }
</script>

<div class="space-y-1">
    {#if label}
        <label for={textareaId} class={labelClass()}>
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

    <textarea
        id={textareaId}
        {name}
        bind:value
        {placeholder}
        {disabled}
        {readonly}
        {required}
        {rows}
        {cols}
        {maxlength}
        {minlength}
        class={textareaClass()}
        oninput={handleInput}
        onchange={handleChange}
        onfocus={handleFocus}
        onblur={handleBlur}
        aria-label={ariaLabel}
        aria-describedby={ariaDescribedby || (error ? `${textareaId}-error` : undefined)}
        aria-invalid={error ? 'true' : 'false'}
    ></textarea>

    {#if error}
        <div id="{textareaId}-error" class={errorClass()} role="alert">
            <svg class="w-4 h-4 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
            <span>{error}</span>
        </div>
    {/if}
</div>

<style>
    .resize-none {
        resize: none;
    }
    
    .resize-both {
        resize: both;
    }
    
    .resize-horizontal {
        resize: horizontal;
    }
    
    .resize-vertical {
        resize: vertical;
    }
</style>