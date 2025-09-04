<script>
  import { AddNPC, GetNPCs, UpdateNPC, DeleteNPC } from '../../wailsjs/go/main/App.js';
  import { onMount } from 'svelte';

  let newNPC = {
    name: '',
    clan: '',
    generation: '',
    notes: '',
    attributes: { Strength: 1, Dexterity: 1, Stamina: 1, Charisma: 1, Manipulation: 1, Appearance: 1, Perception: 1, Intelligence: 1, Wits: 1 },
  };

  let npcs = [];
  let editingNPC = null;
  let searchTerm = '';

  $: filteredNPCs = npcs.filter(npc =>
    npc.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    npc.clan.toLowerCase().includes(searchTerm.toLowerCase()) ||
    npc.generation.toLowerCase().includes(searchTerm.toLowerCase()) ||
    npc.notes.toLowerCase().includes(searchTerm.toLowerCase()) ||
    Object.entries(npc.attributes || {}).some(([key, value]) =>
      key.toLowerCase().includes(searchTerm.toLowerCase()) ||
      String(value).includes(searchTerm.toLowerCase())
    )
  );

  const npcAttributeNames = ["Strength", "Dexterity", "Stamina", "Charisma", "Manipulation", "Appearance", "Perception", "Intelligence", "Wits"];

  async function addNPC() {
    try {
      await AddNPC(newNPC);
      newNPC = {
        name: '',
        clan: '',
        generation: '',
        notes: '',
        attributes: { Strength: 1, Dexterity: 1, Stamina: 1, Charisma: 1, Manipulation: 1, Appearance: 1, Perception: 1, Intelligence: 1, Wits: 1 },
      };
      await getNPCs();
    } catch (error) {
      console.error("Error adding NPC:", error);
      alert("Failed to add NPC: " + error.message);
    }
  }

  async function getNPCs() {
    try {
      npcs = await GetNPCs();
    } catch (error) {
      console.error("Error getting NPCs:", error);
      alert("Failed to load NPCs: " + error.message);
    }
  }

  function editNPC(npc) {
    editingNPC = { ...npc };
  }

  async function updateNPC() {
    try {
      await UpdateNPC(editingNPC);
      editingNPC = null;
      await getNPCs();
    } catch (error) {
      console.error("Error updating NPC:", error);
      alert("Failed to update NPC: " + error.message);
    }
  }

  async function deleteNPC(id) {
    if (!confirm("Are you sure you want to delete this NPC?")) {
      return;
    }
    try {
      await DeleteNPC(id);
      await getNPCs();
    } catch (error) {
      console.error("Error deleting NPC:", error);
      alert("Failed to delete NPC: " + error.message);
    }
  }

  onMount(() => {
    getNPCs();
  });
</script>

<div>
  <h2>NPC Management</h2>

  <h3>Add New NPC</h3>
  <form on:submit|preventDefault={addNPC} class="npc-form">
    <label for="name">Name:</label>
    <input type="text" id="name" bind:value={newNPC.name} required />

    <label for="clan">Clan:</label>
    <input type="text" id="clan" bind:value={newNPC.clan} />

    <label for="generation">Generation:</label>
    <input type="text" id="generation" bind:value={newNPC.generation} />

    <h4>Attributes</h4>
    <div class="attributes-grid">
      {#each npcAttributeNames as attrName}
        <div class="attribute-item">
          <label for="{attrName}">{attrName}:</label>
          <input type="number" id="{attrName}" bind:value={newNPC.attributes[attrName]} min="1" max="5" />
        </div>
      {/each}
    </div>

    <label for="notes">Notes:</label>
    <textarea id="notes" bind:value={newNPC.notes}></textarea>

    <button type="submit">Add NPC</button>
  </form>

  {#if editingNPC}
    <h3>Edit NPC (ID: {editingNPC.id})</h3>
    <form on:submit|preventDefault={updateNPC} class="npc-form">
      <label for="edit-name">Name:</label>
      <input type="text" id="edit-name" bind:value={editingNPC.name} required />

      <label for="edit-clan">Clan:</label>
      <input type="text" id="edit-clan" bind:value={editingNPC.clan} />

      <label for="edit-generation">Generation:</label>
      <input type="text" id="edit-generation" bind:value={editingNPC.generation} />

      <h4>Attributes</h4>
      <div class="attributes-grid">
        {#each npcAttributeNames as attrName}
          <div class="attribute-item">
            <label for="edit-{attrName}">{attrName}:</label>
            <input type="number" id="edit-{attrName}" bind:value={editingNPC.attributes[attrName]} min="1" max="5" />
          </div>
        {/each}
      </div>

      <label for="edit-notes">Notes:</label>
      <textarea id="edit-notes" bind:value={editingNPC.notes}></textarea>

      <button type="submit">Update NPC</button>
      <button type="button" on:click={() => (editingNPC = null)}>Cancel</button>
    </form>
  {/if}

  <h3>Existing NPCs</h3>
  <input type="text" placeholder="Search NPCs..." bind:value={searchTerm} class="search-input" />

  {#if filteredNPCs.length > 0}
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Clan</th>
          <th>Generation</th>
          <th>Attributes</th>
          <th>Notes</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each filteredNPCs as npc (npc.id)}
          <tr>
            <td>{npc.name}</td>
            <td>{npc.clan}</td>
            <td>{npc.generation}</td>
            <td>
              {#each Object.entries(npc.attributes || {}) as [key, value]}
                {key}: {value}<br />
              {/each}
            </td>
            <td>{npc.notes}</td>
            <td>
              <button on:click={() => editNPC(npc)}>Edit</button>
              <button on:click={() => deleteNPC(npc.id)}>Delete</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <p>No NPCs found matching your search or no NPCs added yet.</p>
  {/if}
</div>

<style>
  .npc-form {
    display: grid;
    grid-template-columns: 120px 1fr;
    gap: 10px;
    margin-bottom: 20px;
    max-width: 600px;
  }

  .npc-form label {
    text-align: right;
    padding-right: 10px;
  }

  .npc-form input, .npc-form textarea {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
  }

  .npc-form textarea {
    grid-column: 2;
    margin-top: 5px;
  }

  .npc-form button {
    grid-column: 2;
    padding: 10px 15px;
    border: none;
    cursor: pointer;
    margin-top: 10px;
  }

  .npc-form button[type="button"] {
    background-color: #555555;
  }

  .npc-form button[type="button"]:hover {
    background-color: #777777;
  }

  .search-input {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
    margin-bottom: 20px;
  }

  .attributes-grid {
    grid-column: 1 / -1;
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    margin-bottom: 10px;
    padding: 10px;
    border: 1px solid #444;
    border-radius: 5px;
  }

  .attribute-item {
    display: flex;
    align-items: center;
    gap: 5px;
  }

  .attributes-grid label {
    text-align: left;
    padding-right: 0;
  }

  .attributes-grid input[type="number"] {
    width: 50px;
  }
</style>