<script lang="ts">
    import { Modal } from '@skeletonlabs/skeleton-svelte';
    import type { ComponentProps } from 'svelte';

    type SkeletonModalProps = ComponentProps<typeof Modal>;

    interface Props extends Omit<SkeletonModalProps,'open'> {
        open?: boolean;
        title?: string;
        size?: 'sm' | 'md' | 'lg' | 'xl';
        oncancel?: () => void;
        content?: import('svelte').Snippet;
        trigger?: import('svelte').Snippet;
        
        // Custom props for extending classes
        contentBase?: string;
        contentClasses?: string;
        backdropBase?: string;
        backdropBackground?: string;
        backdropClasses?: string;
    }

    let {
        open = true, // Default to open for direct usage, adjust as needed

        ...restProps // Capture any other Skeleton Modal props
    }: Props = $props();

        const baseClasses = 'bg-white dark:bg-surface-800 rounded-lg shadow-xl p-6 w-full max-w-lg space-y-6 overflow-auto max-h-[calc(100vh-4rem)] min-h-[100px]';  

    let contentBase = restProps.contentBase || baseClasses

    let base = 'fixed inset-0 z-50 flex items-center justify-center p-4';
    let backdropBase = restProps.backdropBase || 'bg-black/50 dark:bg-black/80';
    let backdropBackground = restProps.backdropBackground || 'bg-black/50 dark:bg-black/80';
    let backdropClasses = restProps.backdropClasses || 'fixed inset-0 z-40';
</script>

<Modal
    open={open}
    contentBase={contentBase}
    backdropBase={backdropBase}
    backdropBackground={backdropBackground}
    backdropClasses={backdropClasses}

    {...restProps}

    ></Modal>

