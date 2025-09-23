<script lang="ts">
    interface Props {
        striped?: boolean;
        hoverable?: boolean;
        base?: string;
        classes?: string;
        header?: import('svelte').Snippet;
        body?: import('svelte').Snippet;
        footer?: import('svelte').Snippet;
    }

    let { 
        striped = false, 
        hoverable = true,
        base = '',
        classes = '',
        header,
        body,
        footer
    }: Props = $props();

    let tableClass = $derived(() => {
        const baseClasses = base || 'table w-full';
        const stripedClass = striped ? 'table-striped' : '';
        const hoverClass = hoverable ? 'table-hover' : '';
        
        return [baseClasses, stripedClass, hoverClass, classes]
            .filter(Boolean)
            .join(' ');
    });
</script>

<div class="table-container">
    <table class={tableClass()}>
        {#if header}
            <thead>
                {@render header()}
            </thead>
        {/if}

        <tbody>
            {#if body}
                {@render body()}
            {/if}
        </tbody>

        {#if footer}
            <tfoot>
                {@render footer()}
            </tfoot>
        {/if}
    </table>
</div>