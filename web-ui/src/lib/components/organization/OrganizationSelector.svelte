<script lang="ts">
    import { organizationsStore, currentOrganization, organizations, organizationsLoading } from '$lib/stores/organization';
    import { onMount } from 'svelte';

    let loading = false;

    onMount(async () => {
        await organizationsStore.loadOrganizations();
        // If only one organization, set it as current
        if ($organizations.length === 1) {
            organizationsStore.switchOrganization($organizations[0].id);
        }
    });

    async function handleOrganizationChange(event: Event) {
        const select = event.target as HTMLSelectElement;
        const organizationId = select.value;
        
        if (organizationId) {
            loading = true;
            try {
                await organizationsStore.switchOrganization(organizationId);
            } catch (error) {
                console.error('Failed to switch organization:', error);
            } finally {
                loading = false;
            }
        }
    }
</script>

<div class="relative">
    <select
        class="w-48 px-3 py-2 bg-white border border-slate-200 rounded-lg text-sm text-slate-900 focus:outline-none focus:ring-2 focus:ring-primary-500 disabled:opacity-50"
        onchange={handleOrganizationChange}
        disabled={$organizationsLoading || loading}
    >
        {#if $organizations.length === 0}
            <option value="">No organizations</option>
        {:else}
            {#each $organizations as org}
                <option value={org.id} selected={org.id === $currentOrganization?.id}>
                    {org.name}
                </option>
            {/each}
        {/if}
    </select>

    {#if $organizationsLoading || loading}
        <div class="absolute right-2 top-1/2 -translate-y-1/2">
            <span class="material-symbols-outlined animate-spin text-slate-400 text-sm">
                sync
            </span>
        </div>
    {/if}
</div>

<style>
    select {
        appearance: none;
        background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e");
        background-position: right 0.5rem center;
        background-repeat: no-repeat;
        background-size: 1.5em 1.5em;
        padding-right: 2.5rem;
    }
</style>