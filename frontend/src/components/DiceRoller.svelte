
<script>
  import { RollDice } from '../../wailsjs/go/main/App.js';

  let numDice = 1;
  let difficulty = 6;
  let rollResults = [];
  let successes = 0;
  let resultMessage = "";

  async function rollDice() {
    try {
      if (numDice < 1) {
        numDice = 1;
      }
      if (difficulty < 2) {
        difficulty = 2;
      }
      if (difficulty > 10) {
        difficulty = 10;
      }

      rollResults = await RollDice(numDice);
      
      let successCount = 0;

      rollResults.forEach(die => {
        if (die >= difficulty) {
          successCount++;
        }
        if (die === 1) {
            successCount--;
        }
      });

      successes = successCount;

      resultMessage = `${successes} Successes`;

    } catch (error) {
      console.error("Error rolling dice:", error);
    }
  }
</script>

<div class="card">
    <h2>Dice Roller</h2>
    <div class="dice-roller-controls">
        <label for="numDice">Number of d10s:</label>
        <input type="number" id="numDice" bind:value={numDice} min="1" />
        <label for="difficulty">Difficulty:</label>
        <input type="number" id="difficulty" bind:value={difficulty} min="2" max="10" />
        <button on:click={rollDice}>Roll Dice</button>
    </div>

    {#if rollResults.length > 0}
        <div class="results">
            <h3>Results:</h3>
            <p class="roll">{rollResults.join(', ')}</p>
            <p class="message">{resultMessage}</p>
        </div>
    {/if}
</div>

<style>
    .dice-roller-controls {
        display: flex;
        align-items: center;
        gap: 15px;
        margin-bottom: 20px;
    }

    .dice-roller-controls input {
        width: 80px;
    }

    .results {
        margin-top: 20px;
        padding: 15px;
        background-color: var(--background-primary);
        border-radius: 5px;
    }

    .results .roll {
        font-size: 1.2em;
        font-weight: bold;
        word-wrap: break-word;
    }

    .results .message {
        font-size: 1.1em;
        margin-top: 10px;
        color: var(--accent-primary);
    }
</style>
