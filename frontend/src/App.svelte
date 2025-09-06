<script>
  import NPCManager from './components/NPCManager.svelte';
  import SessionLogger from './components/SessionLogger.svelte';
  import DiceRoller from './components/DiceRoller.svelte';
  import CharacterSheet from './components/CharacterSheet.svelte';
  import InitiativeTracker from './components/InitiativeTracker.svelte';
  import { GetAllCombatants, RollDice } from '../wailsjs/go/main/App.js';

  let activeTab = 'characters';

  function setActiveTab(tab) {
    activeTab = tab;
  }

  // Initiative Tracker State
  let trackerCombatants = [];
  let round = 1;

  function updateTracker(newTracker) {
      trackerCombatants = newTracker;
  }

  function updateRound(newRound) {
      round = newRound;
  }

</script>

<main>
  <div class="tabs">
    <div class="tab" on:click={() => setActiveTab('characters')} class:active={activeTab === 'characters'}>Character Sheets</div>
    <div class="tab" on:click={() => setActiveTab('npcs')} class:active={activeTab === 'npcs'}>NPCs</div>
    <div class="tab" on:click={() => setActiveTab('initiative')} class:active={activeTab === 'initiative'}>Initiative</div>
    <div class="tab" on:click={() => setActiveTab('dice')} class:active={activeTab === 'dice'}>Dice Roller</div>
    <div class="tab" on:click={() => setActiveTab('log')} class:active={activeTab === 'log'}>Session Log</div>
  </div>

  <div class="content">
    {#if activeTab === 'characters'}
      <CharacterSheet />
    {:else if activeTab === 'npcs'}
      <NPCManager />
    {:else if activeTab === 'initiative'}
      <InitiativeTracker 
        bind:trackerCombatants={trackerCombatants}
        bind:round={round}
        {GetAllCombatants}
        {RollDice}
      />
    {:else if activeTab === 'dice'}
      <DiceRoller />
    {:else if activeTab === 'log'}
      <SessionLogger />
    {/if}
  </div>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }

  .content {
    flex-grow: 1;
    padding: 20px;
    overflow-y: auto;
  }

  .tabs {
    display: flex;
    flex-wrap: wrap;
    border-bottom: 1px solid var(--border-color);
  }

  .tab {
    padding: 15px 25px;
    cursor: pointer;
    background-color: var(--background-secondary);
    border-right: 1px solid var(--border-color);
    transition: background-color 0.3s ease;
    font-size: 16px;
    color: var(--text-secondary);
  }

  .tab:hover {
    background-color: var(--background-tertiary);
    color: var(--text-primary);
  }

  .tab.active {
    background-color: var(--background-primary);
    color: var(--accent-primary);
    border-bottom: 1px solid var(--accent-primary);
    position: relative;
    top: 1px;
    font-weight: bold;
  }
</style>