<script lang="ts">
    import type { SvelteComponent } from 'svelte';
    import * as lucide from 'lucide-svelte';

    interface Props {
        name: string;
        size?: string | number;
        strokeWidth?: string | number;
        classes?: string;
    }

    let { 
        name,
        size = '24',
        strokeWidth = '2',
        classes: className = ''
    }: Props = $props();

    let IconComponent = $state<typeof SvelteComponent | null>(null);

    $effect(() => {
        // Convert kebab-case to PascalCase for Lucide icon names
        const pascalCase = name.split('-')
            .map((word: string) => word.charAt(0).toUpperCase() + word.slice(1))
            .join('');
        IconComponent = (lucide as unknown as Record<string, typeof SvelteComponent>)[pascalCase];
    });
</script>

{#if IconComponent}
    <IconComponent
        {size}
        {strokeWidth}
        class={className}
    />
{/if}