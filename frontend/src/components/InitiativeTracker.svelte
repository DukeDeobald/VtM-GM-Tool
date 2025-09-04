<script>
  import { GetAllCombatants, RollDice } from '../../wailsjs/go/main/App.js';
  import { onMount } from 'svelte';

  let availableCombatants = [];
  let trackerCombatants = [];
  let selectedCombatantUniqueId = null;

  async function fetchCombatants() {
    try {
      const fetchedCombatants = await GetAllCombatants();
      if (fetchedCombatants) {
        availableCombatants = fetchedCombatants;
      } else {
        availableCombatants = [];
      }

      if (availableCombatants.length > 0) {
        selectedCombatantUniqueId = `${availableCombatants[0].id}-${availableCombatants[0].type}`;
      }
    } catch (error) {
      console.error("Error fetching combatants:", error);
      alert("Failed to load combatants: " + error.message);
    }
  }

  function addCombatantToTracker() {
    if (selectedCombatantUniqueId === null) return;

    const [idStr, type] = selectedCombatantUniqueId.split('-');
    const id = parseInt(idStr);

    const combatantToAdd = availableCombatants.find(c => c.id === id && c.type === type);

    if (combatantToAdd && !trackerCombatants.some(tc => tc.id === combatantToAdd.id && tc.type === combatantToAdd.type)) {
      trackerCombatants = [...trackerCombatants, { ...combatantToAdd, roll: 0, totalInitiative: 0 }];
      selectedCombatantUniqueId = null;
    }
  }

  async function rollInitiative() {
    for (let i = 0; i < trackerCombatants.length; i++) {
      try {
        const rollResult = await RollDice(1);
        trackerCombatants[i].roll = rollResult[0];
        trackerCombatants[i].totalInitiative = trackerCombatants[i].roll + trackerCombatants[i].initiativeBonus;
      } catch (error) {
        console.error("Error rolling dice for combatant:", trackerCombatants[i].name, error);
        alert("Failed to roll dice for " + trackerCombatants[i].name + ": " + error.message);
      }
    }
    trackerCombatants = trackerCombatants.sort((a, b) => b.totalInitiative - a.totalInitiative);
  }

  function removeCombatant(id, type) {
    trackerCombatants = trackerCombatants.filter(c => !(c.id === id && c.type === type));
  }

  onMount(() => {
    fetchCombatants();
  });
</script>

<div>
  <h2>Initiative Tracker</h2>

  <h3>Add Combatant</h3>
  <select bind:value={selectedCombatantUniqueId}>
    {#each availableCombatants as combatant (`${combatant.id}-${combatant.type}`)}
      <option value={`${combatant.id}-${combatant.type}`}>{combatant.name} ({combatant.type})</option>
    {/each}
  </select>
  <button on:click={addCombatantToTracker}>Add to Tracker</button>

  <h3>Current Initiative</h3>
  <button on:click={rollInitiative}>Roll All Initiative</button>

  {#if trackerCombatants.length > 0}
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Bonus</th>
          <th>Roll</th>
          <th>Total</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each trackerCombatants as combatant (`${combatant.id}-${combatant.type}`)}
          <tr>
            <td>{combatant.name}</td>
            <td>{combatant.initiativeBonus}</td>
            <td>{combatant.roll}</td>
            <td>{combatant.totalInitiative}</td>
            <td>
              <button on:click={() => removeCombatant(combatant.id, combatant.type)}>Remove</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <p>No combatants in tracker. Add some from the dropdown.</p>
  {/if}
</div>

<style>
  select, button {
    margin-right: 10px;
    margin-bottom: 10px;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
  }

  th, td {
    border: 1px solid #444444;
    padding: 8px;
    text-align: left;
    color: #e0e0e0 !important;
    background-color: #222222 !important;
  }

  th {
    background-color: #333333;
  }
</style>
