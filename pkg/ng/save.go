package ng

import (
	"encoding/gob"
	"eramstein/thurigen/pkg/config"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func init() {
	// Register all types that need to be encoded/decoded
	gob.Register([]*Region{})
	gob.Register([]*Item{})
	gob.Register([]*Character{})
	gob.Register([]*PlantStructure{})
	gob.Register(&PlantStructure{})
	gob.Register(&BaseStructure{})
	gob.Register(&PlantProduction{})
	gob.Register(&Position{})
	gob.Register(&TileOccupation{})
	gob.Register(&Item{})
	gob.Register(&Calendar{})
	gob.Register(&Wants{})
	gob.Register(&Confort{})
	gob.Register(&WallStructure{})

	// Register tile-related types
	gob.Register([config.RegionSize][config.RegionSize]Tile{})
	gob.Register(&Tile{})
	gob.Register(TerrainType(0))
	gob.Register(SurfaceType(0))
	gob.Register(VolumeType(0))
	gob.Register(MoveCost(0))
}

// SaveState saves the current simulation state to a file
func (sim *Simulation) SaveState() error {
	// Create saves directory if it doesn't exist
	savesDir := "saves"
	if err := os.MkdirAll(savesDir, 0755); err != nil {
		return fmt.Errorf("failed to create saves directory: %v", err)
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := filepath.Join(savesDir, fmt.Sprintf("save_%s.bin", timestamp))

	// Create save data structure
	saveData := struct {
		Paused     bool         `json:"paused"`
		Speed      int          `json:"speed"`
		Time       int          `json:"time"`
		Calendar   Calendar     `json:"calendar"`
		World      []*Region    `json:"world"`
		Items      []*Item      `json:"items"`
		Characters []*Character `json:"characters"`
	}{
		Paused:     sim.Paused,
		Speed:      sim.Speed,
		Time:       sim.Time,
		Calendar:   sim.Calendar,
		World:      sim.World,
		Items:      sim.Items,
		Characters: sim.Characters,
	}

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create save file: %v", err)
	}
	defer file.Close()

	// Create gob encoder
	encoder := gob.NewEncoder(file)

	// Encode the data
	if err := encoder.Encode(saveData); err != nil {
		return fmt.Errorf("failed to encode simulation state: %v", err)
	}

	return nil
}

// LoadState loads a simulation state from a file
func LoadState(filename string) (*Simulation, error) {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open save file: %v", err)
	}
	defer file.Close()

	// Create gob decoder
	decoder := gob.NewDecoder(file)

	// Unmarshal binary data
	var saveData struct {
		Paused     bool         `json:"paused"`
		Speed      int          `json:"speed"`
		Time       int          `json:"time"`
		Calendar   Calendar     `json:"calendar"`
		World      []*Region    `json:"world"`
		Items      []*Item      `json:"items"`
		Characters []*Character `json:"characters"`
	}

	if err := decoder.Decode(&saveData); err != nil {
		return nil, fmt.Errorf("failed to decode simulation state: %v", err)
	}

	// Create new simulation
	sim := &Simulation{
		Paused:     saveData.Paused,
		Speed:      saveData.Speed,
		Time:       saveData.Time,
		Calendar:   saveData.Calendar,
		World:      saveData.World,
		Items:      saveData.Items,
		Characters: saveData.Characters,
	}

	sim.ReconnectReferences()

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
		if !file.IsDir() && strings.HasPrefix(file.Name(), "save_") && strings.HasSuffix(file.Name(), ".bin") {
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

// ReconnectReferences reconnects all the references after loading a save
// (some data is duplicated for performances reasons, for example which tile a character is on is stored both on the character and on the tile)
func (sim *Simulation) ReconnectReferences() {
	// Reconnect character-tile references
	for _, character := range sim.Characters {
		// Set correct character reference in tile
		sim.World[character.Position.Region].Tiles[character.Position.X][character.Position.Y].Character = character
	}
}
