package ng

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// SaveState saves the current simulation state to a file
func (sim *Simulation) SaveState() error {
	// Create saves directory if it doesn't exist
	savesDir := "saves"
	if err := os.MkdirAll(savesDir, 0755); err != nil {
		return fmt.Errorf("failed to create saves directory: %v", err)
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := filepath.Join(savesDir, fmt.Sprintf("save_%s.json", timestamp))

	// Create save data structure
	saveData := struct {
		Paused     bool         `json:"paused"`
		Speed      int          `json:"speed"`
		Time       int          `json:"time"`
		World      []*Region    `json:"world"`
		Items      []*Item      `json:"items"`
		Characters []*Character `json:"characters"`
	}{
		Paused:     sim.Paused,
		Speed:      sim.Speed,
		Time:       sim.Time,
		World:      sim.World,
		Items:      sim.Items,
		Characters: sim.Characters,
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(saveData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal simulation state: %v", err)
	}

	// Write to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write save file: %v", err)
	}

	return nil
}

// LoadState loads a simulation state from a file
func LoadState(filename string) (*Simulation, error) {
	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read save file: %v", err)
	}

	// Unmarshal JSON
	var saveData struct {
		Paused     bool         `json:"paused"`
		Speed      int          `json:"speed"`
		Time       int          `json:"time"`
		World      []*Region    `json:"world"`
		Items      []*Item      `json:"items"`
		Characters []*Character `json:"characters"`
	}

	if err := json.Unmarshal(data, &saveData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal simulation state: %v", err)
	}

	// Create new simulation
	sim := &Simulation{
		Paused:     saveData.Paused,
		Speed:      saveData.Speed,
		Time:       saveData.Time,
		World:      saveData.World,
		Items:      saveData.Items,
		Characters: saveData.Characters,
	}

	return sim, nil
}

// LoadLatestState finds and loads the most recent save file
func LoadLatestState() (*Simulation, error) {
	// Get all save files
	savesDir := "saves"
	files, err := os.ReadDir(savesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read saves directory: %v", err)
	}

	// Filter for save files and sort by name (which includes timestamp)
	var saveFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "save_") && strings.HasSuffix(file.Name(), ".json") {
			saveFiles = append(saveFiles, file.Name())
		}
	}

	if len(saveFiles) == 0 {
		return nil, fmt.Errorf("no save files found")
	}

	// Sort files by name (which includes timestamp) in descending order
	sort.Sort(sort.Reverse(sort.StringSlice(saveFiles)))

	// Load the most recent save file
	latestFile := filepath.Join(savesDir, saveFiles[0])
	return LoadState(latestFile)
}
