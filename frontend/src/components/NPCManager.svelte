
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

  const presets = [
    {
      name: 'Select a Preset...',
      npc: { attributes: { Strength: 1, Dexterity: 1, Stamina: 1, Charisma: 1, Manipulation: 1, Appearance: 1, Perception: 1, Intelligence: 1, Wits: 1 } }
    },
    {
      name: 'Thug',
      npc: {
        name: 'Thug',
        clan: 'Brujah (Enforcer)',
        generation: '12th',
        notes: 'A hired muscle, ready for a fight.',
        attributes: { Strength: 4, Dexterity: 3, Stamina: 3, Charisma: 2, Manipulation: 1, Appearance: 2, Perception: 2, Intelligence: 1, Wits: 2 },
      }
    },
    {
      name: 'Police Officer',
      npc: {
        name: 'Police Officer',
        clan: 'Mortal',
        generation: '',
        notes: 'A mortal police officer, possibly a ghoul.',
        attributes: { Strength: 3, Dexterity: 3, Stamina: 3, Charisma: 2, Manipulation: 2, Appearance: 2, Perception: 3, Intelligence: 2, Wits: 3 },
      }
    },
    {
      name: 'Fledgling Vampire',
      npc: {
        name: 'Fledgling',
        clan: 'Caitiff',
        generation: '13th',
        notes: 'A newly embraced vampire, struggling to survive.',
        attributes: { Strength: 2, Dexterity: 2, Stamina: 2, Charisma: 2, Manipulation: 2, Appearance: 2, Perception: 2, Intelligence: 2, Wits: 2 },
      }
    },
    {
        name: 'Grizzled Detective',
        npc: {
            name: 'Detective',
            clan: 'Mortal',
            generation: '',
            notes: 'A seasoned investigator with a keen eye for detail.',
            attributes: { Strength: 2, Dexterity: 3, Stamina: 2, Charisma: 2, Manipulation: 3, Appearance: 2, Perception: 4, Intelligence: 3, Wits: 4 },
        }
    },
    {
        name: 'Gangrel Enforcer',
        npc: {
            name: 'Enforcer',
            clan: 'Gangrel',
            generation: '11th',
            notes: 'A tough and feral vampire, skilled in combat and survival.',
            attributes: { Strength: 4, Dexterity: 4, Stamina: 5, Charisma: 1, Manipulation: 2, Appearance: 2, Perception: 3, Intelligence: 2, Wits: 3 },
        }
    },
    {
        name: 'Toreador Artisan',
        npc: {
            name: 'Artisan',
            clan: 'Toreador',
            generation: '10th',
            notes: 'A charismatic and influential artist, a master of social gatherings.',
            attributes: { Strength: 2, Dexterity: 3, Stamina: 2, Charisma: 5, Manipulation: 4, Appearance: 4, Perception: 3, Intelligence: 3, Wits: 3 },
        }
    },
    {
        name: 'Ventrue Executive',
        npc: {
            name: 'Executive',
            clan: 'Ventrue',
            generation: '9th',
            notes: 'A powerful and wealthy executive with considerable influence.',
            attributes: { Strength: 2, Dexterity: 2, Stamina: 3, Charisma: 4, Manipulation: 5, Appearance: 3, Perception: 4, Intelligence: 4, Wits: 4 },
        }
    }
  ];

  let selectedPreset = presets[0];

  function applyPreset() {
      if(selectedPreset.name === 'Select a Preset...') {
          newNPC = { name: '', clan: '', generation: '', notes: '', attributes: { Strength: 1, Dexterity: 1, Stamina: 1, Charisma: 1, Manipulation: 1, Appearance: 1, Perception: 1, Intelligence: 1, Wits: 1 } };
      } else {
        newNPC = { ...selectedPreset.npc };
      }
  }

  let npcs = [];
  let editingNPC = null;
  let searchTerm = '';

  $: filteredNPCs = npcs.filter(npc =>
    npc.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    (npc.clan && npc.clan.toLowerCase().includes(searchTerm.toLowerCase())) ||
    (npc.generation && npc.generation.toLowerCase().includes(searchTerm.toLowerCase())) ||
    (npc.notes && npc.notes.toLowerCase().includes(searchTerm.toLowerCase()))
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
      selectedPreset = presets[0];
      await getNPCs();
    } catch (error) {
      console.error("Error adding NPC:", error);
    }
  }

  async function getNPCs() {
    try {
      npcs = await GetNPCs();
    } catch (error) {
      console.error("Error getting NPCs:", error);
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
    }
  }

  onMount(() => {
    getNPCs();
  });
</script>

<div class="card">
  <h2>Add New NPC</h2>
  <div class="preset-selector">
      <label for="preset">Load from Preset:</label>
      <select id="preset" bind:value={selectedPreset} on:change={applyPreset}>
          {#each presets as preset}
              <option value={preset}>{preset.name}</option>
          {/each}
      </select>
  </div>
  <form on:submit|preventDefault={addNPC} class="npc-form">
    <div class="form-grid">
        <input type="text" placeholder="Name" bind:value={newNPC.name} required />
        <input type="text" placeholder="Clan" bind:value={newNPC.clan} />
        <input type="text" placeholder="Generation" bind:value={newNPC.generation} />
    </div>

    <h4>Attributes</h4>
    <div class="attributes-grid">
      {#each npcAttributeNames as attrName}
        <div class="attribute-item">
          <label for="{attrName}">{attrName}</label>
          <input type="number" id="{attrName}" bind:value={newNPC.attributes[attrName]} min="1" max="10" />
        </div>
      {/each}
    </div>

    <textarea placeholder="Notes" bind:value={newNPC.notes}></textarea>

    <button type="submit">Add NPC</button>
  </form>
</div>

{#if editingNPC}
  <div class="card">
    <h2>Edit NPC (ID: {editingNPC.id})</h2>
    <form on:submit|preventDefault={updateNPC} class="npc-form">
        <div class="form-grid">
            <input type="text" placeholder="Name" bind:value={editingNPC.name} required />
            <input type="text" placeholder="Clan" bind:value={editingNPC.clan} />
            <input type="text" placeholder="Generation" bind:value={editingNPC.generation} />
        </div>

        <h4>Attributes</h4>
        <div class="attributes-grid">
            {#each npcAttributeNames as attrName}
                <div class="attribute-item">
                    <label for="edit-{attrName}">{attrName}</label>
                    <input type="number" id="edit-{attrName}" bind:value={editingNPC.attributes[attrName]} min="1" max="10" />
                </div>
            {/each}
        </div>

        <textarea placeholder="Notes" bind:value={editingNPC.notes}></textarea>

      <button type="submit">Update NPC</button>
      <button type="button" class="secondary" on:click={() => (editingNPC = null)}>Cancel</button>
    </form>
  </div>
{/if}

<div class="card">
  <h2>Existing NPCs</h2>
  <input type="text" placeholder="Search NPCs..." bind:value={searchTerm} />

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
                <span class="stat">{key}: {value}</span>
              {/each}
            </td>
            <td>{npc.notes}</td>
            <td>
              <button on:click={() => editNPC(npc)}>Edit</button>
              <button class="secondary" on:click={() => deleteNPC(npc.id)}>Delete</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <p>No NPCs found.</p>
  {/if}
</div>

<style>
    .preset-selector {
        display: flex;
        align-items: center;
        gap: 15px;
        margin-bottom: 20px;
    }

    .npc-form {
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

    .form-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
        gap: 15px;
    }

    .attributes-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
        gap: 15px;
    }

    .attribute-item {
        display: flex;
        flex-direction: column;
    }

    .attribute-item label {
        margin-bottom: 5px;
        font-size: 0.9em;
        color: var(--text-secondary);
    }

    .stat {
        display: inline-block;
        background-color: var(--background-tertiary);
        padding: 2px 6px;
        border-radius: 3px;
        margin: 2px;
        font-size: 0.9em;
    }

    td {
        vertical-align: top;
    }
</style>
