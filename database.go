package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	npcsFile        = "npcs.json"
	sessionLogsFile = "session_logs.json"
	charactersFile  = "characters.json"
)

type NPC struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Clan       string         `json:"clan"`
	Generation string         `json:"generation"`
	Notes      string         `json:"notes"`
	Attributes map[string]int `json:"attributes"`
}

type SessionLog struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Character struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Clan       string         `json:"clan"`
	Generation string         `json:"generation"`
	Concept    string         `json:"concept"`
	Sire       string         `json:"sire"`
	Attributes map[string]int `json:"attributes"`
	Abilities  struct {
		Talents    map[string]int `json:"talents"`
		Skills     map[string]int `json:"skills"`
		Knowledges map[string]int `json:"knowledges"`
	} `json:"abilities"`
	Disciplines map[string]int `json:"disciplines"`
	Notes       string         `json:"notes"`
}

func InitDB(dataDir string) {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.MkdirAll(dataDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create data directory: %v", err)
		}
	}

	npcsPath := filepath.Join(dataDir, npcsFile)
	if _, err := os.Stat(npcsPath); os.IsNotExist(err) {
		err := ioutil.WriteFile(npcsPath, []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Failed to create %s: %v", npcsFile, err)
		}
	}

	sessionLogsPath := filepath.Join(dataDir, sessionLogsFile)
	if _, err := os.Stat(sessionLogsPath); os.IsNotExist(err) {
		err := ioutil.WriteFile(sessionLogsPath, []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Failed to create %s: %v", sessionLogsFile, err)
		}
	}

	charactersPath := filepath.Join(dataDir, charactersFile)
	if _, err := os.Stat(charactersPath); os.IsNotExist(err) {
		err := ioutil.WriteFile(charactersPath, []byte("[]"), 0644)
		if err != nil {
			log.Fatalf("Failed to create %s: %v", charactersFile, err)
		}
	}
	log.Println("JSON data files initialized.")
}

func CloseDB() {
	log.Println("JSON data files closed (no action needed).")
}

func readNPCs() ([]NPC, error) {
	data, err := ioutil.ReadFile(filepath.Join("data", npcsFile))
	if err != nil {
		if os.IsNotExist(err) || len(data) == 0 {
			return []NPC{}, nil
		}
		return nil, err
	}
	var npcs []NPC
	err = json.Unmarshal(data, &npcs)
	if err != nil {
		return []NPC{}, err
	}
	return npcs, nil
}

func writeNPCs(npcs []NPC) error {
	data, err := json.MarshalIndent(npcs, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join("data", npcsFile), data, 0644)
}

func readSessionLogs() ([]SessionLog, error) {
	data, err := ioutil.ReadFile(filepath.Join("data", sessionLogsFile))
	if err != nil {
		if os.IsNotExist(err) || len(data) == 0 {
			return []SessionLog{}, nil
		}
		return nil, err
	}
	var logs []SessionLog
	err = json.Unmarshal(data, &logs)
	if err != nil {
		return []SessionLog{}, err
	}
	return logs, nil
}

func writeSessionLogs(logs []SessionLog) error {
	data, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join("data", sessionLogsFile), data, 0644)
}

func readCharacters() ([]Character, error) {
	data, err := ioutil.ReadFile(filepath.Join("data", charactersFile))
	if err != nil {
		if os.IsNotExist(err) || len(data) == 0 {
			return []Character{}, nil
		}
		return nil, err
	}
	var characters []Character
	err = json.Unmarshal(data, &characters)
	if err != nil {
		return []Character{}, err
	}
	return characters, nil
}

func writeCharacters(characters []Character) error {
	data, err := json.MarshalIndent(characters, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join("data", charactersFile), data, 0644)
}
