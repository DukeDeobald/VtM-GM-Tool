<script>
  import { onMount } from 'svelte';

  export let trackerCombatants;
  export let round;
  export let GetAllCombatants;
  export let RollDice;

  let availableCombatants = [];
  let selectedCombatantUniqueId = null;

  function incrementRound() {
      round++;
  }

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
    }
  }

  function addCombatantToTracker() {
    if (selectedCombatantUniqueId === null) return;

    const [idStr, type] = selectedCombatantUniqueId.split('-');
    const id = parseInt(idStr);

    const combatantToAdd = availableCombatants.find(c => c.id === id && c.type === type);

    if (combatantToAdd && !trackerCombatants.some(tc => tc.id === combatantToAdd.id && tc.type === combatantToAdd.type)) {
      trackerCombatants = [...trackerCombatants, { ...combatantToAdd, roll: 0, totalInitiative: 0, health: 7, status: '' }];
      selectedCombatantUniqueId = null;
    }
  }

  async function rollInitiative() {
      const newTrackerCombatants = [...trackerCombatants];
    for (let i = 0; i < newTrackerCombatants.length; i++) {
      try {
        const rollResult = await RollDice(1);
        newTrackerCombatants[i].roll = rollResult[0];
        const bonus = newTrackerCombatants[i].initiativeBonus || 0;
        newTrackerCombatants[i].totalInitiative = newTrackerCombatants[i].roll + bonus;
      } catch (error) {
        console.error("Error rolling dice for combatant:", newTrackerCombatants[i].name, error);
      }
    }
    trackerCombatants = newTrackerCombatants.sort((a, b) => b.totalInitiative - a.totalInitiative);
  }

  function removeCombatant(id, type) {
    trackerCombatants = trackerCombatants.filter(c => !(c.id === id && c.type === type));
  }

  onMount(() => {
    fetchCombatants();
  });
</script>

<div class="card">
  <h2>Initiative Tracker</h2>

  <div class="controls">
    <select bind:value={selectedCombatantUniqueId}>
        <option value={null}>Select a combatant...</option>
      {#each availableCombatants as combatant (`${combatant.id}-${combatant.type}`)}
        <option value={`${combatant.id}-${combatant.type}`}>{combatant.name} ({combatant.type})</option>
      {/each}
    </select>
    <button on:click={addCombatantToTracker}>Add to Tracker</button>
    <button on:click={rollInitiative} class="secondary">Roll All Initiative</button>
  </div>

  <div class="round-counter">
      <h3>Round: {round}</h3>
      <button on:click={incrementRound}>Next Round</button>
  </div>

  {#if trackerCombatants.length > 0}
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Bonus</th>
          <th>Roll</th>
          <th>Total</th>
          <th>Health</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each trackerCombatants as combatant (`${combatant.id}-${combatant.type}`)}
          <tr>
            <td>{combatant.name}</td>
            <td>{combatant.initiativeBonus || 0}</td>
            <td>{combatant.roll}</td>
            <td>{combatant.totalInitiative}</td>
            <td><input type="number" bind:value={combatant.health} min="0" /></td>
            <td><input type="text" bind:value={combatant.status} placeholder="Normal" /></td>
            <td>
              <button class="secondary" on:click={() => removeCombatant(combatant.id, combatant.type)}>Remove</button>
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
    .controls {
        display: flex;
        gap: 15px;
        margin-bottom: 20px;
        align-items: center;
    }

    .controls select {
        flex-grow: 1;
    }

    .round-counter {
        display: flex;
        align-items: center;
        gap: 20px;
        margin-bottom: 20px;
    }

    td input {
        width: 100px;
    }

    td input[type="text"] {
        width: 150px;
    }
</style>