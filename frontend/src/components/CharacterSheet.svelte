<script>
  import { AddCharacter, GetCharacters, UpdateCharacter, DeleteCharacter } from '../../wailsjs/go/main/App.js';
  import { onMount } from 'svelte';

  let newCharacter = {
    name: '',
    clan: '',
    generation: '',
    concept: '',
    sire: '',
    attributes: { Strength: 1, Dexterity: 1, Stamina: 1, Charisma: 1, Manipulation: 1, Appearance: 1, Perception: 1, Intelligence: 1, Wits: 1 },
    abilities: {
      talents: { Alertness: 0, Athletics: 0, Brawl: 0, Dodge: 0, Empathy: 0, Expression: 0, Intimidation: 0, Leadership: 0, Streetwise: 0, Subterfuge: 0 },
      skills: { AnimalKen: 0, Crafts: 0, Drive: 0, Etiquette: 0, Firearms: 0, Larceny: 0, Melee: 0, Performance: 0, Stealth: 0, Survival: 0 },
      knowledges: { Academics: 0, Computer: 0, Finance: 0, Investigation: 0, Law: 0, Linguistics: 0, Medicine: 0, Occult: 0, Politics: 0, Science: 0 },
    },
    disciplines: {},
    notes: '',
  };

  let characters = [];
  let editingCharacter = null;

  const attributeNames = ["Strength", "Dexterity", "Stamina", "Charisma", "Manipulation", "Appearance", "Perception", "Intelligence", "Wits"];
  const talentNames = ["Alertness", "Athletics", "Brawl", "Dodge", "Empathy", "Expression", "Intimidation", "Leadership", "Streetwise", "Subterfuge"];
  const skillNames = ["AnimalKen", "Crafts", "Drive", "Etiquette", "Firearms", "Larceny", "Melee", "Performance", "Stealth", "Survival"];
  const knowledgeNames = ["Academics", "Computer", "Finance", "Investigation", "Law", "Linguistics", "Medicine", "Occult", "Politics", "Science"];

  let newDisciplineName = '';
  let newDisciplineRating = 0;

  function addDiscipline(character) {
    if (newDisciplineName.trim() === '') return;
    character.disciplines = { ...character.disciplines, [newDisciplineName]: newDisciplineRating };
    newDisciplineName = '';
    newDisciplineRating = 0;
  }

  function removeDiscipline(character, name) {
    const updatedDisciplines = { ...character.disciplines };
    delete updatedDisciplines[name];
    character.disciplines = updatedDisciplines;
  }

  async function addCharacter() {
    try {
      await AddCharacter(newCharacter);
      newCharacter = {
        name: '',
        clan: '',
        generation: '',
        concept: '',
        sire: '',
        attributes: { Strength: 1, Dexterity: 1, Stamina: 1, Charisma: 1, Manipulation: 1, Appearance: 1, Perception: 1, Intelligence: 1, Wits: 1 },
        abilities: {
          talents: { Alertness: 0, Athletics: 0, Brawl: 0, Dodge: 0, Empathy: 0, Expression: 0, Intimidation: 0, Leadership: 0, Streetwise: 0, Subterfuge: 0 },
          skills: { AnimalKen: 0, Crafts: 0, Drive: 0, Etiquette: 0, Firearms: 0, Larceny: 0, Melee: 0, Performance: 0, Stealth: 0, Survival: 0 },
          knowledges: { Academics: 0, Computer: 0, Finance: 0, Investigation: 0, Law: 0, Linguistics: 0, Medicine: 0, Occult: 0, Politics: 0, Science: 0 },
        },
        disciplines: {},
        notes: '',
      };
      await getCharacters();
    } catch (error) {
      console.error("Error adding character:", error);
      alert("Failed to add character: " + error.message);
    }
  }

  async function getCharacters() {
    try {
      characters = await GetCharacters();
    } catch (error) {
      console.error("Error getting characters:", error);
      alert("Failed to load characters: " + error.message);
    }
  }

  function editCharacter(char) {
    editingCharacter = { ...char };
  }

  async function updateCharacter() {
    try {
      await UpdateCharacter(editingCharacter);
      editingCharacter = null;
      await getCharacters();
    } catch (error) {
      console.error("Error updating character:", error);
      alert("Failed to update character: " + error.message);
    }
  }

  async function deleteCharacter(id) {
    if (!confirm("Are you sure you want to delete this character?")) {
      return;
    }
    try {
      await DeleteCharacter(id);
      await getCharacters();
    } catch (error) {
      console.error("Error deleting character:", error);
      alert("Failed to delete character: " + error.message);
    }
  }

  onMount(() => {
    getCharacters();
  });
</script>

<div>
  <h2>Character Sheet</h2>

  <h3>Add New Character</h3>
  <form on:submit|preventDefault={addCharacter} class="character-form">
    <label for="name">Name:</label>
    <input type="text" id="name" bind:value={newCharacter.name} required />

    <label for="clan">Clan:</label>
    <input type="text" id="clan" bind:value={newCharacter.clan} />

    <label for="generation">Generation:</label>
    <input type="text" id="generation" bind:value={newCharacter.generation} />

    <label for="concept">Concept:</label>
    <input type="text" id="concept" bind:value={newCharacter.concept} />

    <label for="sire">Sire:</label>
    <input type="text" id="sire" bind:value={newCharacter.sire} />

    <h4>Attributes</h4>
    {#each attributeNames as attrName}
      <label for="{attrName}">{attrName}:</label>
      <input type="number" id="{attrName}" bind:value={newCharacter.attributes[attrName]} min="1" max="5" />
    {/each}

    <h4>Abilities - Talents</h4>
    {#each talentNames as abilityName}
      <label for="talent-{abilityName}">{abilityName}:</label>
      <input type="number" id="talent-{abilityName}" bind:value={newCharacter.abilities.talents[abilityName]} min="0" max="5" />
    {/each}

    <h4>Abilities - Skills</h4>
    {#each skillNames as abilityName}
      <label for="skill-{abilityName}">{abilityName}:</label>
      <input type="number" id="skill-{abilityName}" bind:value={newCharacter.abilities.skills[abilityName]} min="0" max="5" />
    {/each}

    <h4>Abilities - Knowledges</h4>
    {#each knowledgeNames as abilityName}
      <label for="knowledge-{abilityName}">{abilityName}:</label>
      <input type="number" id="knowledge-{abilityName}" bind:value={newCharacter.abilities.knowledges[abilityName]} min="0" max="5" />
    {/each}

    <h4>Disciplines</h4>
    <div class="discipline-input">
      <input type="text" placeholder="Discipline Name" bind:value={newDisciplineName} />
      <input type="number" bind:value={newDisciplineRating} min="0" max="5" />
      <button type="button" on:click={() => addDiscipline(newCharacter)}>Add Discipline</button>
    </div>
    <div class="discipline-list">
      {#each Object.entries(newCharacter.disciplines) as [name, rating]}
        <span>{name}: {rating} <button type="button" on:click={() => removeDiscipline(newCharacter, name)}>X</button></span>
      {/each}
    </div>

    <h4>Notes</h4>
    <textarea bind:value={newCharacter.notes}></textarea>

    <button type="submit">Add Character</button>
  </form>

  {#if editingCharacter}
    <h3>Edit Character (ID: {editingCharacter.id})</h3>
    <form on:submit|preventDefault={updateCharacter} class="character-form">
      <label for="edit-name">Name:</label>
      <input type="text" id="edit-name" bind:value={editingCharacter.name} required />

      <label for="edit-clan">Clan:</label>
      <input type="text" id="edit-clan" bind:value={editingCharacter.clan} />

      <label for="edit-generation">Generation:</label>
      <input type="text" id="edit-generation" bind:value={editingCharacter.generation} />

      <label for="edit-concept">Concept:</label>
      <input type="text" id="edit-concept" bind:value={editingCharacter.concept} />

      <label for="edit-sire">Sire:</label>
      <input type="text" id="edit-sire" bind:value={editingCharacter.sire} />

      <h4>Attributes</h4>
      {#each attributeNames as attrName}
        <label for="edit-{attrName}">{attrName}:</label>
        <input type="number" id="edit-{attrName}" bind:value={editingCharacter.attributes[attrName]} min="1" max="5" />
      {/each}

      <h4>Abilities - Talents</h4>
      {#each talentNames as abilityName}
        <label for="edit-talent-{abilityName}">{abilityName}:</label>
        <input type="number" id="edit-talent-{abilityName}" bind:value={editingCharacter.abilities.talents[abilityName]} min="0" max="5" />
      {/each}

      <h4>Abilities - Skills</h4>
      {#each skillNames as abilityName}
        <label for="edit-skill-{abilityName}">{abilityName}:</label>
        <input type="number" id="edit-skill-{abilityName}" bind:value={editingCharacter.abilities.skills[abilityName]} min="0" max="5" />
      {/each}

      <h4>Abilities - Knowledges</h4>
      {#each knowledgeNames as abilityName}
        <label for="edit-knowledge-{abilityName}">{abilityName}:</label>
        <input type="number" id="edit-knowledge-{abilityName}" bind:value={editingCharacter.abilities.knowledges[abilityName]} min="0" max="5" />
      {/each}

      <h4>Disciplines</h4>
      <div class="discipline-input">
        <input type="text" placeholder="Discipline Name" bind:value={newDisciplineName} />
        <input type="number" bind:value={newDisciplineRating} min="0" max="5" />
        <button type="button" on:click={() => addDiscipline(editingCharacter)}>Add Discipline</button>
      </div>
      <div class="discipline-list">
        {#each Object.entries(editingCharacter.disciplines) as [name, rating]}
          <span>{name}: {rating} <button type="button" on:click={() => removeDiscipline(editingCharacter, name)}>X</button></span>
        {/each}
      </div>

      <h4>Notes</h4>
      <textarea bind:value={editingCharacter.notes}></textarea>

      <button type="submit">Update Character</button>
      <button type="button" on:click={() => (editingCharacter = null)}>Cancel</button>
    </form>
  {/if}

  <h3>Existing Characters</h3>
  {#if characters.length > 0}
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Clan</th>
          <th>Generation</th>
          <th>Concept</th>
          <th>Sire</th>
          <th>Attributes</th>
          <th>Abilities</th>
          <th>Disciplines</th>
          <th>Notes</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each characters as char (char.id)}
          <tr>
            <td>{char.name}</td>
            <td>{char.clan}</td>
            <td>{char.generation}</td>
            <td>{char.concept}</td>
            <td>{char.sire}</td>
            <td>
              {#each Object.entries(char.attributes) as [key, value]}
                {key}: {value}<br />
              {/each}
            </td>
            <td>
              <strong>Talents:</strong><br />
              {#each Object.entries(char.abilities.talents) as [key, value]}
                {#if value > 0}{key}: {value}<br />{/if}
              {/each}
              <strong>Skills:</strong><br />
              {#each Object.entries(char.abilities.skills) as [key, value]}
                {#if value > 0}{key}: {value}<br />{/if}
              {/each}
              <strong>Knowledges:</strong><br />
              {#each Object.entries(char.abilities.knowledges) as [key, value]}
                {#if value > 0}{key}: {value}<br />{/if}
              {/each}
            </td>
            <td>
              {#each Object.entries(char.disciplines || {}) as [key, value]}
                {key}: {value}<br />
              {/each}
            </td>
            <td>{char.notes}</td>
            <td>
              <button on:click={() => editCharacter(char)}>Edit</button>
              <button on:click={() => deleteCharacter(char.id)}>Delete</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <p>No characters added yet.</p>
  {/if}
</div>

<style>
  .character-form {
    display: grid;
    grid-template-columns: 120px 1fr; /* Adjust column width for labels */
    gap: 10px;
    margin-bottom: 20px;
    max-width: 600px;
  }

  .character-form label {
    text-align: right;
    padding-right: 10px;
  }

  .character-form input[type="text"], .character-form textarea {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
  }

  .character-form input[type="number"] {
    width: 60px;
    text-align: center;
  }

  .character-form h4 {
    grid-column: 1 / -1; /* Span all columns */
    margin-top: 20px;
    margin-bottom: 10px;
    color: #ff6666; /* Lighter red for headings */
  }

  .character-form textarea {
    grid-column: 2; /* Make textarea span the second column */
    margin-top: 5px; /* Add some space above it */
  }

  .character-form button {
    grid-column: 2;
    padding: 10px 15px;
    border: none;
    cursor: pointer;
    margin-top: 10px;
  }

  .character-form button[type="button"] {
    background-color: #555555;
  }

  .character-form button[type="button"]:hover {
    background-color: #777777;
  }

  .discipline-input {
    grid-column: 1 / -1;
    display: flex;
    gap: 10px;
    margin-bottom: 10px;
  }

  .discipline-input input {
    flex-grow: 1;
  }

  .discipline-list {
    grid-column: 1 / -1;
    margin-bottom: 10px;
  }

  .discipline-list span {
    display: inline-block;
    background-color: #444;
    padding: 5px 10px;
    border-radius: 5px;
    margin-right: 5px;
    margin-bottom: 5px;
  }

  .discipline-list button {
    background-color: #cc0000;
    color: white;
    border: none;
    border-radius: 3px;
    padding: 2px 5px;
    cursor: pointer;
    font-size: 0.8em;
    margin-left: 5px;
  }
</style>