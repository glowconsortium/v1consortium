<script lang="ts">
	import { onMount } from 'svelte';
	import { invitationStore, invitations, loading, error } from '@movsm/v1-consortium-web-pkg';
	import { currentOrganization } from '@movsm/v1-consortium-web-pkg';
	import type { Invitation } from '@movsm/v1-consortium-web-pkg';
	import { Alert, Button, Card, Table, Modal, Select } from '@movsm/v1-consortium-web-pkg';
	import { toaster } from '@movsm/v1-consortium-web-pkg';

	let showInviteModal = false;
	let inviteEmail = '';
	let selectedRoleId = '';
	//let roles = $currentOrganization?.role ? [{ id: $currentOrganization.role.id, name: $currentOrganization.role.name }] : [];
	let roles =  ['admin', 'editor', 'viewer'];
	onMount(() => {
		console.log('Loading invitations for organization:', $currentOrganization);
		invitationStore.loadInvitations();
	});

	async function handleInvite() {
		try {
			await invitationStore.inviteMember({
				id: $currentOrganization?.id || '',
				email: inviteEmail,
				role_id: selectedRoleId,
				role: selectedRoleId
			});
			showInviteModal = false;
			inviteEmail = '';
			selectedRoleId = '';
			toaster.create({ type: 'success', title: 'Invitation sent successfully' });
		} catch (error: any) {
			toaster.create({ type: 'error', title: error.message || 'Failed to send invitation' });
		}
	}

	async function handleCancel(invitation: Invitation) {
		try {
			await invitationStore.cancelInvitation(invitation.id);
			toaster.create({ type: 'success', title: 'Invitation cancelled successfully' });
		} catch (error: any) {
			toaster.create({ type: 'error', title: error.message || 'Failed to cancel invitation' });
		}
	}

	async function handleResend(invitation: Invitation) {
		try {
			await invitationStore.resendInvitation(invitation.id);
			toaster.create({ type: 'success', title: 'Invitation resent successfully' });
		} catch (error: any) {
			toaster.create({ type: 'error', title: error.message || 'Failed to resend invitation' });
		}
	}

	function formatDate(dateString: string) {
		return new Date(dateString).toLocaleDateString();
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<h1 class="h2">Team Invitations</h1>
		<Button variant="filled" onclick={() => showInviteModal = true}>
			<span>Invite Member</span>
		</Button>
	</div>

	{#if $error}
		<Alert type="error" title="Error" message={$error} />
	{/if}

	<Card>
		{#if $loading}
			<div class="flex justify-center p-4">
				<span class="loading loading-spinner loading-lg" />
			</div>
		{:else if $invitations.length === 0}
			<div class="text-center p-8">
				<p class="text-surface-600-300-token">No invitations found</p>
			</div>
		{:else}
			<Table>
				{#snippet header()}
					
					<tr>
						<th>Email</th>
						<th>Role</th>
						<th>Status</th>
						<th>Sent</th>
						<th>Expires</th>
						<th>Actions</th>
					</tr>
									{/snippet}
				{#snippet body()}
					{#each $invitations as invitation}
						<tr>
							<td>{invitation.email}</td>
							<td>{invitation.role || 'N/A'}</td>
							<td>
								<span class="badge" class:badge-success={invitation.status === 'accepted'} class:badge-warning={invitation.status === 'pending'} class:badge-error={invitation.status === 'expired' || invitation.status === 'cancelled'}>
									{invitation.status}
								</span>
							</td>
							<td>{formatDate(invitation.created_at)}</td>
							<td>{formatDate(invitation.expires_at)}</td>
							<td class="space-x-2">
								{#if invitation.status === 'pending'}
									<Button size="sm" variant="ghost" onclick={() => handleResend(invitation)}>
										<span>Resend</span>
									</Button>
									<Button size="sm" variant="ghost" color="error" onclick={() => handleCancel(invitation)}>
										<span>Cancel</span>
									</Button>
								{/if}
							</td>
						</tr>
					{/each}
				{/snippet}
				</Table>
		{/if}
	</Card>
</div>

<Modal open={showInviteModal} title="Invite Team Member">
	{#snippet content()}
	<form on:submit|preventDefault={handleInvite} class="space-y-4">
		<div class="form-control">
			<label for="email" class="label">Email</label>
			<input
				type="email"
				id="email"
				bind:value={inviteEmail}
				class="input"
				placeholder="Enter email address"
				required
			/>
		</div>


		<div class="form-control">
			<label for="role" class="label">Role</label>
			<Select
				id="role"
				bind:value={selectedRoleId}
				options={roles.map(role => ({
					value: role ?? '',
					label: role ?? ''
				}))}
				placeholder="Select a role"
				required
			/>
		</div>

		<div class="flex justify-end space-x-2">
			<Button variant="ghost" onclick={() => showInviteModal = false}>
				<span>Cancel</span>
			</Button>
			<Button type="submit" variant="filled" disabled={$loading}>
				<span>Send Invitation</span>
			</Button>
		</div>
	</form>
		{/snippet}

</Modal> 