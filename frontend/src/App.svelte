<script>
  import NPCManager from './components/NPCManager.svelte';
  import SessionLogger from './components/SessionLogger.svelte';
  import DiceRoller from './components/DiceRoller.svelte';
  import CharacterSheet from './components/CharacterSheet.svelte';
  import InitiativeTracker from './components/InitiativeTracker.svelte';

  let activeTab = 'npc';

  function setActiveTab(tab) {
    activeTab = tab;
  }
</script>

<main>
  <div class="sidebar">
    <div class="logo-text">VtM</div>
    <button on:click={() => setActiveTab('npc')} class:active={activeTab === 'npc'}>NPC Management</button>
    <button on:click={() => setActiveTab('log')} class:active={activeTab === 'log'}>Session Log</button>
    <button on:click={() => setActiveTab('dice')} class:active={activeTab === 'dice'}>Dice Roller</button>
    <button on:click={() => setActiveTab('character')} class:active={activeTab === 'character'}>Character Sheets</button>
    <button on:click={() => setActiveTab('initiative')} class:active={activeTab === 'initiative'}>Initiative</button>
  </div>

  <div class="content">
    {#if activeTab === 'npc'}
      <NPCManager />
    {:else if activeTab === 'log'}
      <SessionLogger />
    {:else if activeTab === 'dice'}
      <DiceRoller />
    {:else if activeTab === 'character'}
      <CharacterSheet />
    {:else if activeTab === 'initiative'}
      <InitiativeTracker />
    {/if}
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    font-family: 'Arial', sans-serif;
    background-color: #1a1a1a;
    color: #e0e0e0;
  }

  main {
    display: flex;
    height: 100vh;
    background-color: #1a1a1a;
  }

  .sidebar {
    width: 200px;
    background-color: #222222;
    padding: 20px;
    display: flex;
    flex-direction: column;
    border-right: 1px solid #444444;
  }

  .logo-text {
    font-size: 3em;
    font-weight: bold;
    color: #cc0000;
    text-align: center;
    margin-bottom: 30px;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
  }

  .sidebar button {
    padding: 15px 10px;
    margin-bottom: 10px;
    border: none;
    background-color: transparent;
    color: #e0e0e0;
    cursor: pointer;
    font-size: 16px;
    text-align: left;
    transition: background-color 0.3s, color 0.3s;
  }

  .sidebar button:hover {
    background-color: #333333;
    color: #ff3333;
  }

  .sidebar button.active {
    background-color: #444444;
    color: #ff0000;
    border-left: 3px solid #ff0000;
    padding-left: 7px;
  }

  .content {
    flex-grow: 1;
    padding: 20px;
    overflow-y: auto;
    background-color: #1a1a1a;
  }

  :global(h2, h3) {
    color: #ff6666;
  }

  :global(label) {
    color: #cccccc;
  }

  :global(input, textarea, button) {
    background-color: #333333;
    color: #e0e0e0;
    border: 1px solid #555555;
  }

  :global(input:focus, textarea:focus) {
    border-color: #ff0000;
    outline: none;
  }

  :global(button) {
    background-color: #cc0000;
    color: white;
  }

  :global(button:hover) {
    background-color: #ff3333;
  }

  :global(table) {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
  }

  :global(tr) {
    background-color: #1a1a1a;
  }

  :global(th, td) {
    border: 1px solid #444444;
    padding: 8px;
    text-align: left;
    color: #e0e0e0 !important;
    background-color: #222222 !important;
  }

  :global(th) {
    background-color: #333333;
  }
</style>