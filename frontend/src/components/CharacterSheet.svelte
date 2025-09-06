
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
    }
  }

  async function getCharacters() {
    try {
      characters = await GetCharacters();
    } catch (error) {
      console.error("Error getting characters:", error);
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
    }
  }

  onMount(() => {
    getCharacters();
  });
</script>

<div class="card">
  <h2>Add New Character</h2>
  <form on:submit|preventDefault={addCharacter} class="character-form">
    <div class="form-grid">
        <input type="text" placeholder="Name" bind:value={newCharacter.name} required />
        <input type="text" placeholder="Clan" bind:value={newCharacter.clan} />
        <input type="text" placeholder="Generation" bind:value={newCharacter.generation} />
        <input type="text" placeholder="Concept" bind:value={newCharacter.concept} />
        <input type="text" placeholder="Sire" bind:value={newCharacter.sire} />
    </div>

    <h4>Attributes</h4>
    <div class="attributes-grid">
        {#each attributeNames as attrName}
            <div class="attribute-item">
                <label for="{attrName}">{attrName}</label>
                <input type="number" id="{attrName}" bind:value={newCharacter.attributes[attrName]} min="1" max="5" />
            </div>
        {/each}
    </div>

    <h4>Abilities</h4>
    <div class="abilities-grid">
        <div class="ability-group">
            <h5>Talents</h5>
            {#each talentNames as abilityName}
                <div class="ability-item">
                    <label for="talent-{abilityName}">{abilityName}</label>
                    <input type="number" id="talent-{abilityName}" bind:value={newCharacter.abilities.talents[abilityName]} min="0" max="5" />
                </div>
            {/each}
        </div>
        <div class="ability-group">
            <h5>Skills</h5>
            {#each skillNames as abilityName}
                <div class="ability-item">
                    <label for="skill-{abilityName}">{abilityName}</label>
                    <input type="number" id="skill-{abilityName}" bind:value={newCharacter.abilities.skills[abilityName]} min="0" max="5" />
                </div>
            {/each}
        </div>
        <div class="ability-group">
            <h5>Knowledges</h5>
            {#each knowledgeNames as abilityName}
                <div class="ability-item">
                    <label for="knowledge-{abilityName}">{abilityName}</label>
                    <input type="number" id="knowledge-{abilityName}" bind:value={newCharacter.abilities.knowledges[abilityName]} min="0" max="5" />
                </div>
            {/each}
        </div>
    </div>

    <h4>Disciplines</h4>
    <div class="discipline-input">
      <input type="text" placeholder="Discipline Name" bind:value={newDisciplineName} />
      <input type="number" bind:value={newDisciplineRating} min="0" max="5" />
      <button type="button" on:click={() => addDiscipline(newCharacter)}>Add</button>
    </div>
    <div class="discipline-list">
      {#each Object.entries(newCharacter.disciplines) as [name, rating]}
        <span>{name}: {rating} <button type="button" class="small" on:click={() => removeDiscipline(newCharacter, name)}>X</button></span>
      {/each}
    </div>

    <h4>Notes</h4>
    <textarea bind:value={newCharacter.notes} placeholder="Character notes..."></textarea>

    <button type="submit">Add Character</button>
  </form>
</div>

{#if editingCharacter}
  <div class="card">
    <h2>Edit Character (ID: {editingCharacter.id})</h2>
    <form on:submit|preventDefault={updateCharacter} class="character-form">
        <div class="form-grid">
            <input type="text" placeholder="Name" bind:value={editingCharacter.name} required />
            <input type="text" placeholder="Clan" bind:value={editingCharacter.clan} />
            <input type="text" placeholder="Generation" bind:value={editingCharacter.generation} />
            <input type="text" placeholder="Concept" bind:value={editingCharacter.concept} />
            <input type="text" placeholder="Sire" bind:value={editingCharacter.sire} />
        </div>

        <h4>Attributes</h4>
        <div class="attributes-grid">
            {#each attributeNames as attrName}
                <div class="attribute-item">
                    <label for="edit-{attrName}">{attrName}</label>
                    <input type="number" id="edit-{attrName}" bind:value={editingCharacter.attributes[attrName]} min="1" max="5" />
                </div>
            {/each}
        </div>

        <h4>Abilities</h4>
        <div class="abilities-grid">
            <div class="ability-group">
                <h5>Talents</h5>
                {#each talentNames as abilityName}
                    <div class="ability-item">
                        <label for="edit-talent-{abilityName}">{abilityName}</label>
                        <input type="number" id="edit-talent-{abilityName}" bind:value={editingCharacter.abilities.talents[abilityName]} min="0" max="5" />
                    </div>
                {/each}
            </div>
            <div class="ability-group">
                <h5>Skills</h5>
                {#each skillNames as abilityName}
                    <div class="ability-item">
                        <label for="edit-skill-{abilityName}">{abilityName}</label>
                        <input type="number" id="edit-skill-{abilityName}" bind:value={editingCharacter.abilities.skills[abilityName]} min="0" max="5" />
                    </div>
                {/each}
            </div>
            <div class="ability-group">
                <h5>Knowledges</h5>
                {#each knowledgeNames as abilityName}
                    <div class="ability-item">
                        <label for="edit-knowledge-{abilityName}">{abilityName}</label>
                        <input type="number" id="edit-knowledge-{abilityName}" bind:value={editingCharacter.abilities.knowledges[abilityName]} min="0" max="5" />
                    </div>
                {/each}
            </div>
        </div>

        <h4>Disciplines</h4>
        <div class="discipline-input">
          <input type="text" placeholder="Discipline Name" bind:value={newDisciplineName} />
          <input type="number" bind:value={newDisciplineRating} min="0" max="5" />
          <button type="button" on:click={() => addDiscipline(editingCharacter)}>Add</button>
        </div>
        <div class="discipline-list">
          {#each Object.entries(editingCharacter.disciplines) as [name, rating]}
            <span>{name}: {rating} <button type="button" class="small" on:click={() => removeDiscipline(editingCharacter, name)}>X</button></span>
          {/each}
        </div>

        <h4>Notes</h4>
        <textarea bind:value={editingCharacter.notes} placeholder="Character notes..."></textarea>

      <button type="submit">Update Character</button>
      <button type="button" class="secondary" on:click={() => (editingCharacter = null)}>Cancel</button>
    </form>
  </div>
{/if}

<div class="card">
  <h2>Existing Characters</h2>
  {#if characters.length > 0}
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Clan</th>
          <th>Generation</th>
          <th>Attributes</th>
          <th>Abilities</th>
          <th>Disciplines</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each characters as char (char.id)}
          <tr>
            <td>{char.name}</td>
            <td>{char.clan}</td>
            <td>{char.generation}</td>
            <td>
              {#each Object.entries(char.attributes) as [key, value]}
                <span class="stat">{key}: {value}</span>
              {/each}
            </td>
            <td>
              {#each Object.entries(char.abilities.talents) as [key, value]}
                {#if value > 0}<span class="stat">{key}: {value}</span>{/if}
              {/each}
              {#each Object.entries(char.abilities.skills) as [key, value]}
                {#if value > 0}<span class="stat">{key}: {value}</span>{/if}
              {/each}
              {#each Object.entries(char.abilities.knowledges) as [key, value]}
                {#if value > 0}<span class="stat">{key}: {value}</span>{/if}
              {/each}
            </td>
            <td>
              {#each Object.entries(char.disciplines || {}) as [key, value]}
                <span class="stat">{key}: {value}</span>
              {/each}
            </td>
            <td>
              <button on:click={() => editCharacter(char)}>Edit</button>
              <button class="secondary" on:click={() => deleteCharacter(char.id)}>Delete</button>
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

    .abilities-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 20px;
    }

    .ability-group h5 {
        color: var(--text-primary);
        margin-bottom: 10px;
    }

    .ability-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 5px;
    }

    .ability-item label {
        font-size: 0.9em;
    }

    .discipline-input {
        display: flex;
        gap: 10px;
        align-items: center;
    }

    .discipline-input input[type="number"] {
        width: 80px;
    }

    .discipline-list span {
        display: inline-block;
        background-color: var(--background-tertiary);
        padding: 5px 10px;
        border-radius: 3px;
        margin-right: 5px;
        margin-bottom: 5px;
    }

    .discipline-list button.small {
        padding: 2px 5px;
        font-size: 0.8em;
        margin-left: 5px;
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
