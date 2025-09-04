package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	rand.Seed(time.Now().UnixNano())
}

func (a *App) AddNPC(npc NPC) error {
	npcs, err := readNPCs()
	if err != nil {
		log.Println("Error reading NPCs:", err)
		return err
	}

	if npc.Attributes == nil {
		npc.Attributes = make(map[string]int)
	}

	newID := 1
	if len(npcs) > 0 {
		newID = npcs[len(npcs)-1].ID + 1
	}
	npc.ID = newID

	npcs = append(npcs, npc)
	return writeNPCs(npcs)
}

func (a *App) GetNPCs() ([]NPC, error) {
	npcs, err := readNPCs()
	if err != nil {
		log.Println("Error reading NPCs:", err)
		return nil, err
	}

	for i := range npcs {
		if npcs[i].Attributes == nil {
			npcs[i].Attributes = make(map[string]int)
		}
	}

	return npcs, nil
}

func (a *App) UpdateNPC(updatedNPC NPC) error {
	npcs, err := readNPCs()
	if err != nil {
		log.Println("Error reading NPCs for update:", err)
		return err
	}

	found := false
	for i, npc := range npcs {
		if npc.ID == updatedNPC.ID {
			npcs[i] = updatedNPC
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("NPC with ID %d not found", updatedNPC.ID)
	}

	return writeNPCs(npcs)
}

func (a *App) DeleteNPC(id int) error {
	npcs, err := readNPCs()
	if err != nil {
		log.Println("Error reading NPCs for delete:", err)
		return err
	}

	updatedNPCs := []NPC{}
	found := false
	for _, npc := range npcs {
		if npc.ID == id {
			found = true
		} else {
			updatedNPCs = append(updatedNPCs, npc)
		}
	}

	if !found {
		return fmt.Errorf("NPC with ID %d not found", id)
	}

	return writeNPCs(updatedNPCs)
}

func (a *App) AddSessionLog(content string) error {
	logs, err := readSessionLogs()
	if err != nil {
		log.Println("Error reading session logs:", err)
		return err
	}

	newID := 1
	if len(logs) > 0 {
		newID = logs[len(logs)-1].ID + 1
	}
	logEntry := SessionLog{ID: newID, Content: content}

	logs = append(logs, logEntry)
	return writeSessionLogs(logs)
}

func (a *App) GetSessionLogs() ([]SessionLog, error) {
	logs, err := readSessionLogs()
	if err != nil {
		log.Println("Error reading session logs:", err)
		return nil, err
	}
	return logs, nil
}

func (a *App) UpdateSessionLog(updatedLog SessionLog) error {
	logs, err := readSessionLogs()
	if err != nil {
		log.Println("Error reading session logs for update:", err)
		return err
	}

	found := false
	for i, logEntry := range logs {
		if logEntry.ID == updatedLog.ID {
			logs[i] = updatedLog
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Session log with ID %d not found", updatedLog.ID)
	}

	return writeSessionLogs(logs)
}

func (a *App) DeleteSessionLog(id int) error {
	logs, err := readSessionLogs()
	if err != nil {
		log.Println("Error reading session logs for delete:", err)
		return err
	}

	updatedLogs := []SessionLog{}
	found := false
	for _, logEntry := range logs {
		if logEntry.ID == id {
			found = true
		} else {
			updatedLogs = append(updatedLogs, logEntry)
		}
	}

	if !found {
		return fmt.Errorf("Session log with ID %d not found", id)
	}

	return writeSessionLogs(updatedLogs)
}

func (a *App) RollDice(numDice int) []int {
	results := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		results[i] = rand.Intn(10) + 1
	}
	return results
}

func (a *App) AddCharacter(char Character) error {
	chars, err := readCharacters()
	if err != nil {
		log.Println("Error reading characters:", err)
		return err
	}

	if char.Abilities.Talents == nil {
		char.Abilities.Talents = make(map[string]int)
	}
	if char.Abilities.Skills == nil {
		char.Abilities.Skills = make(map[string]int)
	}
	if char.Abilities.Knowledges == nil {
		char.Abilities.Knowledges = make(map[string]int)
	}

	newID := 1
	if len(chars) > 0 {
		newID = chars[len(chars)-1].ID + 1
	}
	char.ID = newID

	chars = append(chars, char)
	return writeCharacters(chars)
}

func (a *App) GetCharacters() ([]Character, error) {
	chars, err := readCharacters()
	if err != nil {
		log.Println("Error reading characters:", err)
		return nil, err
	}

	for i := range chars {
		if chars[i].Abilities.Talents == nil {
			chars[i].Abilities.Talents = make(map[string]int)
		}
		if chars[i].Abilities.Skills == nil {
			chars[i].Abilities.Skills = make(map[string]int)
		}
		if chars[i].Abilities.Knowledges == nil {
			chars[i].Abilities.Knowledges = make(map[string]int)
		}
	}

	return chars, nil
}

func (a *App) UpdateCharacter(updatedChar Character) error {
	chars, err := readCharacters()
	if err != nil {
		log.Println("Error reading characters for update:", err)
		return err
	}

	found := false
	for i, char := range chars {
		if char.ID == updatedChar.ID {
			chars[i] = updatedChar
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Character with ID %d not found", updatedChar.ID)
	}

	return writeCharacters(chars)
}

func (a *App) DeleteCharacter(id int) error {
	chars, err := readCharacters()
	if err != nil {
		log.Println("Error reading characters for delete:", err)
		return err
	}

	updatedChars := []Character{}
	found := false
	for _, char := range chars {
		if char.ID == id {
			found = true
		} else {
			updatedChars = append(updatedChars, char)
		}
	}

	if !found {
		return fmt.Errorf("Character with ID %d not found", id)
	}

	return writeCharacters(updatedChars)
}

type Combatant struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	InitiativeBonus int    `json:"initiativeBonus"`
	Type            string `json:"type"`
}

func (a *App) GetAllCombatants() ([]Combatant, error) {
	var combatants []Combatant

	npcs, err := readNPCs()
	if err != nil {
		log.Println("Error reading NPCs for combatants:", err)
		return nil, err
	}
	for _, npc := range npcs {
		dex := 0
		wits := 0
		if npc.Attributes != nil {
			dex = npc.Attributes["Dexterity"]
			wits = npc.Attributes["Wits"]
		}
		initiativeBonus := dex + wits
		combatants = append(combatants, Combatant{ID: npc.ID, Name: npc.Name, InitiativeBonus: initiativeBonus, Type: "npc"})
	}

	chars, err := readCharacters()
	if err != nil {
		log.Println("Error reading Characters for combatants:", err)
		return nil, err
	}
	for _, char := range chars {
		dex := 0
		wits := 0
		if char.Attributes != nil {
			dex = char.Attributes["Dexterity"]
			wits = char.Attributes["Wits"]
		}
		initiativeBonus := dex + wits
		combatants = append(combatants, Combatant{ID: char.ID, Name: char.Name, InitiativeBonus: initiativeBonus, Type: "character"})
	}

	return combatants, nil
}
