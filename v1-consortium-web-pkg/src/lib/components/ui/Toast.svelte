<script lang="ts">
    import { fade, fly } from 'svelte/transition';

    interface Props {
        type?: 'success' | 'warning' | 'error' | 'info';
        message: string;
        title?: string;
        duration?: number;
        position?: 'top-right' | 'top-left' | 'bottom-right' | 'bottom-left';
        base?: string;
        background?: string;
        border?: string;
        rounded?: string;
        padding?: string;
        classes?: string;
        ondismiss?: () => void;
    }

    let {
        type = 'info',
        message,
        title = '',
        duration = 5000,
        position = 'top-right',
        base = '',
        background = '',
        border = '',
        rounded = '',
        padding = '',
        classes = '',
        ondismiss
    }: Props = $props();

    let isVisible = $state(true);

    $effect(() => {
        if (duration > 0) {
            const timer = setTimeout(() => {
                isVisible = false;
                ondismiss?.();
            }, duration);

            return () => clearTimeout(timer);
        }
    });

    let toastClass = $derived(() => {
        const baseClasses = base || 'toast flex items-start gap-4';
        
        // Background and border based on type using Skeleton's color system
        const variantClass = background || `variant-filled-${type}`;
        const borderClass = border || '';
        const roundedClass = rounded || 'rounded-container-token';
        const paddingClass = padding || 'p-4';
        
        return [
            baseClasses,
            variantClass,
            borderClass,
            roundedClass,
            paddingClass,
            classes
        ]
            .filter(Boolean)
            .join(' ');
    });

    let positionClass = $derived(() => {
        const positions = {
            'top-right': 'top-4 right-4',
            'top-left': 'top-4 left-4',
            'bottom-right': 'bottom-4 right-4',
            'bottom-left': 'bottom-4 left-4'
        };
        return `fixed z-[888] ${positions[position]}`;
    });

    function dismiss() {
        isVisible = false;
        ondismiss?.();
    }
</script>

{#if isVisible}
    <div
        class={positionClass()}
        transition:fly={{ y: position.startsWith('top') ? -20 : 20, duration: 300 }}
    >
        <div
            class={toastClass()}
            role="alert"
        >
            <!-- Content -->
            <div class="flex-1">
                {#if title}
                    <h4 class="font-bold">{title}</h4>
                {/if}
                <p>{message}</p>
            </div>

            <!-- Dismiss button -->
            <button
                type="button"
                class="btn-icon variant-ghost"
                aria-label="Dismiss"
                onclick={dismiss}
            >
                <svg class="w-4 h-4" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
            </button>
        </div>
    </div>
{/if}