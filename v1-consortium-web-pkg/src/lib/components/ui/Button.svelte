<script lang="ts">
    interface Props {
        type?: 'button' | 'submit' | 'reset';
        variant?: 'filled' | 'outline' | 'soft' | 'ghost' | 'ring';
        color?: 'primary' | 'secondary' | 'tertiary' | 'success' | 'warning' | 'error' | 'surface';
        size?: 'sm' | 'md' | 'lg';
        disabled?: boolean;
        loading?: boolean;
        base?: string;
        background?: string;
        border?: string;
        rounded?: string;
        padding?: string;
        classes?: string;
        onclick?: () => void;
        children?: import('svelte').Snippet;
    }

    let {
        type = 'button',
        variant = 'filled',
        color = 'primary',
        size = 'md',
        disabled = false,
        loading = false,
        base = '',
        background = '',
        border = '',
        rounded = '',
        padding = '',
        classes = '',
        onclick,
        children
    }: Props = $props();

    let buttonClass = $derived(() => {
        // Base button classes following Skeleton UI patterns
        const baseClasses = base || 'btn inline-flex items-center justify-center font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2';
        
        // Variant and color combinations using Skeleton's preset system
        const variantColorClass = (() => {
            switch (variant) {
                case 'filled':
                    return background || `preset-filled-${color}-500`;
                case 'outline':
                    return background || `preset-outlined-${color}-500`;
                case 'soft':
                    return background || `preset-soft-${color}-500`;
                case 'ghost':
                    return background || `preset-ghost-${color}-500`;
                case 'ring':
                    return background || `preset-ring-${color}-500`;
                default:
                    return background || `btn-${color}-500`;
            }
        })();

        // Size classes
        const sizeClasses = {
            sm: 'btn-sm',
            md: 'btn-md',
            lg: 'btn-lg'
        };
        const sizeClass = sizeClasses[size] || 'btn-md';

        // Disabled state
        const disabledClass = disabled || loading ? 'opacity-50 cursor-not-allowed' : '';

        // Border and spacing
        const borderClass = border || '';
        const roundedClass = rounded || '';
        const paddingClass = padding || '';
        
        return [
            baseClasses,
            variantColorClass,
            sizeClass,
            borderClass,
            roundedClass,
            paddingClass,
            disabledClass,
            classes
        ]
            .filter(Boolean)
            .join(' ');
    });
</script>

<button
    {type}
    class={buttonClass()}
    disabled={disabled || loading}
    onclick={onclick}
    aria-busy={loading}
>
    {#if loading}
        <svg class="w-4 h-4 mr-2 animate-spin" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
    {/if}
    {#if children}
        {@render children()}
    {/if}
</button>
