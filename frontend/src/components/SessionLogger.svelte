
<script>
  import { AddSessionLog, GetSessionLogs, UpdateSessionLog, DeleteSessionLog } from '../../wailsjs/go/main/App.js';
  import { onMount } from 'svelte';

  let newLogContent = '';
  let sessionLogs = [];
  let editingLog = null;

  async function addSessionLog() {
    try {
      if (newLogContent.trim() === '') {
        return;
      }
      await AddSessionLog(newLogContent);
      newLogContent = '';
      await getSessionLogs();
    } catch (error) {
      console.error("Error adding session log:", error);
    }
  }

  async function getSessionLogs() {
    try {
      sessionLogs = await GetSessionLogs();
    } catch (error) {
      console.error("Error getting session logs:", error);
    }
  }

  function editSessionLog(logEntry) {
    editingLog = { ...logEntry };
  }

  async function updateSessionLog() {
    try {
      if (editingLog.content.trim() === '') {
        return;
      }
      await UpdateSessionLog(editingLog);
      editingLog = null;
      await getSessionLogs();
    } catch (error) {
      console.error("Error updating session log:", error);
    }
  }

  async function deleteSessionLog(id) {
    if (!confirm("Are you sure you want to delete this session log?")) {
      return;
    }
    try {
      await DeleteSessionLog(id);
      await getSessionLogs();
    } catch (error) {
      console.error("Error deleting session log:", error);
    }
  }

  onMount(() => {
    getSessionLogs();
  });
</script>

<div class="card">
  <h2>Session Log</h2>

  <form on:submit|preventDefault={addSessionLog}>
    <textarea bind:value={newLogContent} placeholder="Enter session notes here..." rows="8"></textarea>
    <button type="submit">Add Log Entry</button>
  </form>

  {#if editingLog}
    <div class="card editing-card">
        <h3>Edit Log Entry</h3>
        <textarea bind:value={editingLog.content} rows="8"></textarea>
        <button on:click={updateSessionLog}>Update Log</button>
        <button type="button" class="secondary" on:click={() => (editingLog = null)}>Cancel</button>
    </div>
  {/if}

  <h3>Previous Log Entries</h3>
  <div class="log-list">
      {#if sessionLogs.length > 0}
        {#each sessionLogs as logEntry (logEntry.id)}
          <div class="log-entry card">
            <p>{logEntry.content}</p>
            <div class="log-meta">
                <small>Log ID: {logEntry.id}</small>
                <div class="log-actions">
                  <button class="small" on:click={() => editSessionLog(logEntry)}>Edit</button>
                  <button class="small secondary" on:click={() => deleteSessionLog(logEntry.id)}>Delete</button>
                </div>
            </div>
          </div>
        {/each}
      {:else}
        <p>No session logs yet.</p>
      {/if}
  </div>
</div>

<style>
    form {
        margin-bottom: 30px;
    }

    textarea {
        width: 100%;
        margin-bottom: 10px;
    }

    .editing-card {
        background-color: var(--background-tertiary);
        margin-bottom: 20px;
    }

    .log-list {
        display: flex;
        flex-direction: column;
        gap: 15px;
    }

    .log-entry p {
        white-space: pre-wrap;
        margin-bottom: 15px;
    }

    .log-meta {
        display: flex;
        justify-content: space-between;
        align-items: center;
        color: var(--text-secondary);
        font-size: 0.9em;
    }

    .log-actions {
        display: flex;
        gap: 10px;
    }

    button.small {
        padding: 5px 10px;
        font-size: 0.9em;
    }
</style>
