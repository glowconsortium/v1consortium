<script lang="ts">
    interface Props {
        color?: 'primary' | 'secondary' | 'tertiary' | 'success' | 'warning' | 'error' | 'surface';
        size?: 'sm' | 'md' | 'lg';
        variant?: 'filled' | 'outline' | 'soft' | 'ghost' | 'ring';
        base?: string;
        background?: string;
        border?: string;
        rounded?: string;
        padding?: string;
        classes?: string;
        children?: import('svelte').Snippet;
    }

    let {
        color = 'primary',
        size = 'md',
        variant = 'filled',
        base = '',
        background = '',
        border = '',
        rounded = '',
        padding = '',
        classes = '',
        children
    }: Props = $props();

    let badgeClass = $derived(() => {
        const baseClasses = base || 'badge inline-flex items-center font-medium';
        
        // Variant and color combinations using Skeleton's preset system
        const variantColorClass = (() => {
            switch (variant) {
                case 'filled':
                    return background || `variant-filled-${color}`;
                case 'outline':
                    return background || `variant-outline-${color}`;
                case 'soft':
                    return background || `variant-soft-${color}`;
                case 'ghost':
                    return background || `variant-ghost-${color}`;
                case 'ring':
                    return background || `variant-ring-${color}`;
                default:
                    return background || `variant-filled-${color}`;
            }
        })();

        // Size classes
        const sizeClasses = {
            sm: 'badge-sm',
            md: 'badge-md',
            lg: 'badge-lg'
        };
        const sizeClass = sizeClasses[size];

        // Border and spacing
        const borderClass = border || '';
        const roundedClass = rounded || 'rounded-token';
        const paddingClass = padding || '';
        
        return [
            baseClasses,
            variantColorClass,
            sizeClass,
            borderClass,
            roundedClass,
            paddingClass,
            classes
        ]
            .filter(Boolean)
            .join(' ');
    });
</script>

<span class={badgeClass()}>
    {@render children?.()}
</span> 