<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
	import type { PageProps } from './$types';

    import {ApiClientPkg} from "@movsm/v1-consortium-web-pkg"

	let { data }: PageProps = $props();
  let workflowID = data.workflowID;

  let loading = $state(true);
  let error: string | null = $state(null);
  let signupStatus: any = $state(null);
  let statusMessage = $state('Checking signup status...');
  let pollCount = $state(0);

  let apiClient = ApiClientPkg.apiClient;
  let pollInterval: NodeJS.Timeout | null = null;

  async function checkSignupStatus() {
    try {
      loading = true;
      error = null;
      pollCount++;
      statusMessage = `Polling signup status... (attempt ${pollCount})`;
      
      const response = await apiClient.auth.getSignupStatus(workflowID);
      signupStatus = response;
      
      if (response) {
        statusMessage = `Status received: ${JSON.stringify(response, null, 2)}`;
        loading = false;
      } else {
        statusMessage = 'No status data received';
      }
    } catch (err: any) {
      error = err.message || 'Failed to get signup status';
      statusMessage = `Error occurred: ${error}`;
      loading = false;
    }
  }

  onMount(() => {
    // Initial check
    checkSignupStatus();
    
    // Set up polling every 3 seconds
    pollInterval = setInterval(() => {
      checkSignupStatus();
    }, 3000);
  });

  onDestroy(() => {
    if (pollInterval) {
      clearInterval(pollInterval);
    }
  });
</script>

{#if loading}
  <div class="status-container">
    <p class="status-message">{statusMessage}</p>
    <div class="loading-indicator">
      <span>●</span><span>●</span><span>●</span>
    </div>
  </div>
{:else if error}
  <div class="error-container">
    <p class="error">Error: {error}</p>
    <p class="status-message">{statusMessage}</p>
    <p class="poll-info">Poll count: {pollCount}</p>
  </div>
{:else}
  <div class="success-container">
    <h1>Signup Status for Workflow: {workflowID}</h1>
    <p class="status-message">{statusMessage}</p>
    <p class="poll-info">Poll count: {pollCount}</p>
    
    {#if signupStatus}
      <div class="status-data">
        <h2>Status Data:</h2>
        <pre>{JSON.stringify(signupStatus, null, 2)}</pre>
      </div>
    {/if}
  </div>
{/if}

<style>
  .status-container {
    text-align: center;
    padding: 2rem;
  }
  
  .status-message {
    font-size: 1.1rem;
    margin-bottom: 1rem;
    color: #666;
  }
  
  .loading-indicator {
    display: flex;
    justify-content: center;
    gap: 0.5rem;
  }
  
  .loading-indicator span {
    animation: pulse 1.5s infinite;
    font-size: 1.5rem;
  }
  
  .loading-indicator span:nth-child(2) {
    animation-delay: 0.5s;
  }
  
  .loading-indicator span:nth-child(3) {
    animation-delay: 1s;
  }
  
  @keyframes pulse {
    0%, 100% { opacity: 0.3; }
    50% { opacity: 1; }
  }
  
  .error-container {
    padding: 2rem;
    border: 2px solid #ff4444;
    border-radius: 8px;
    background-color: #fff5f5;
  }
  
  .error {
    color: #ff4444;
    font-weight: bold;
    margin-bottom: 1rem;
  }
  
  .success-container {
    padding: 2rem;
  }
  
  .poll-info {
    font-size: 0.9rem;
    color: #888;
    margin-top: 0.5rem;
  }
  
  .status-data {
    margin-top: 2rem;
    padding: 1rem;
    background-color: #f5f5f5;
    border-radius: 4px;
  }
  
  .status-data pre {
    background-color: #fff;
    padding: 1rem;
    border-radius: 4px;
    overflow-x: auto;
    font-size: 0.9rem;
  }
</style>