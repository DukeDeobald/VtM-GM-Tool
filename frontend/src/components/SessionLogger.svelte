<script>
  import { AddSessionLog, GetSessionLogs, UpdateSessionLog, DeleteSessionLog } from '../../wailsjs/go/main/App.js';
  import { onMount } from 'svelte';

  let newLogContent = '';
  let sessionLogs = [];
  let editingLog = null;

  async function addSessionLog() {
    try {
      if (newLogContent.trim() === '') {
        alert("Session log cannot be empty.");
        return;
      }
      await AddSessionLog(newLogContent);
      newLogContent = '';
      await getSessionLogs();
    } catch (error) {
      console.error("Error adding session log:", error);
      alert("Failed to add session log: " + error.message);
    }
  }

  async function getSessionLogs() {
    try {
      sessionLogs = await GetSessionLogs();
    } catch (error) {
      console.error("Error getting session logs:", error);
      alert("Failed to load session logs: " + error.message);
    }
  }

  function editSessionLog(logEntry) {
    editingLog = { ...logEntry };
  }

  async function updateSessionLog() {
    try {
      if (editingLog.content.trim() === '') {
        alert("Session log cannot be empty.");
        return;
      }
      await UpdateSessionLog(editingLog);
      editingLog = null;
      await getSessionLogs();
    } catch (error) {
      console.error("Error updating session log:", error);
      alert("Failed to update session log: " + error.message);
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
      alert("Failed to delete session log: " + error.message);
    }
  }

  onMount(() => {
    getSessionLogs();
  });
</script>

<div>
  <h2>Session Log</h2>

  <h3>Add New Log Entry</h3>
  <textarea bind:value={newLogContent} placeholder="Enter session notes here..." rows="10"></textarea>
  <button on:click={addSessionLog}>Add Log Entry</button>

  {#if editingLog}
    <h3>Edit Log Entry (ID: {editingLog.id})</h3>
    <textarea bind:value={editingLog.content} rows="10"></textarea>
    <button on:click={updateSessionLog}>Update Log</button>
    <button type="button" on:click={() => (editingLog = null)}>Cancel</button>
  {/if}

  <h3>Previous Log Entries</h3>
  {#if sessionLogs.length > 0}
    {#each sessionLogs as logEntry (logEntry.id)}
      <div class="log-entry">
        <p>{logEntry.content}</p>
        <small>Log ID: {logEntry.id}</small>
        <div class="log-actions">
          <button on:click={() => editSessionLog(logEntry)}>Edit</button>
          <button on:click={() => deleteSessionLog(logEntry.id)}>Delete</button>
        </div>
      </div>
    {/each}
  {:else}
    <p>No session logs yet.</p>
  {/if}
</div>

<style>
  textarea {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    box-sizing: border-box;
  }

  button {
    display: block;
    margin-top: 10px;
    padding: 10px 15px;
    border: none;
    cursor: pointer;
    margin-bottom: 20px;
  }

  .log-entry {
    border: 1px solid #444;
    padding: 10px;
    margin-bottom: 10px;
    background-color: #2a2a2a;
    color: #e0e0e0;
  }

  .log-entry p {
    margin-top: 0;
    white-space: pre-wrap;
  }

  .log-entry small {
    color: #aaa;
  }

  .log-actions button {
    margin-right: 10px;
    margin-bottom: 0;
    display: inline-block;
  }
</style>