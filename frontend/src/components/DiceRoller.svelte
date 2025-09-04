<script>
  import { RollDice } from '../../wailsjs/go/main/App.js';

  let numDice = 1;
  let rollResults = [];

  async function rollDice() {
    try {
      if (numDice < 1) {
        alert("Number of dice must be at least 1.");
        return;
      }
      rollResults = await RollDice(numDice);
    } catch (error) {
      console.error("Error rolling dice:", error);
      alert("Failed to roll dice: " + error.message);
    }
  }
</script>

<div>
  <h2>Dice Roller</h2>

  <div>
    <label for="numDice">Number of d10s:</label>
    <input type="number" id="numDice" bind:value={numDice} min="1" />
    <button on:click={rollDice}>Roll Dice</button>
  </div>

  {#if rollResults.length > 0}
    <h3>Results:</h3>
    <p>{rollResults.join(', ')}</p>
  {/if}
</div>

<style>
  div {
    margin-bottom: 15px;
  }

  label {
    margin-right: 10px;
  }

  input[type="number"] {
    width: 60px;
    text-align: center;
    margin-right: 10px;
  }

  button {
    padding: 8px 15px;
  }
</style>
