package ng

import (
	"eramstein/thurigen/pkg/config"
	"time"
)

// Simulation represents the main simulation state
type Simulation struct {
	Paused     bool
	Speed      int // how many frames until next sim update
	Time       int // in minutes since the start of the simulation
	Calendar   Calendar
	World      []*Region
	Items      []*Item
	Characters []*Character
}

type Calendar struct {
	Minute int
	Hour   int
	Day    int
}

type Day struct {
	Date time.Time
}

type Position struct {
	Region int
	X      int
	Y      int
}

type Character struct {
	ID          uint64
	Name        string
	Position    Position
	Stats       CharacterStats
	Inventory   []*Item
	CurrentTask *Task
	Tasks       []*Task
	Objectives  []*Objective
	Ambitions   []*Ambition
	Needs       Needs
	Path        *[]Position
}

type CharacterStats struct {
	Speed float32 // movement speed, modifier applied to move cost
}

type Needs struct {
	Food  int // 0-100
	Water int // 0-100
	Sleep int // 1 = 1 minute of sleep needed
}

type Task struct {
	ID        uint64
	Type      TaskType
	Objective *Objective
	Progress  float32     // by default, 0 to 1, as percent of task already done, but can be used otherwise like for movement
	Target    interface{} // Can be any type of struct
}

type Objective struct {
	Type ObjectiveType
}

type Ambition struct {
	Description string
}

type Region struct {
	Tiles  [config.RegionSize][config.RegionSize]Tile // Tiles[X][Y]
	Plants []*PlantStructure
}

type Tile struct {
	Terrain    TerrainType
	Surface    SurfaceType
	Volume     VolumeType
	Occupation *TileOccupation
	Character  *Character // Character standing on the tile
	Items      []*Item    // Items lying on the tile
	MoveCost   MoveCost   // Cached for pathfinding
}

type TileOccupation struct {
	Structure  Structure
	IsMainTile bool // True if this is the main tile of the object
}

// Structures occupy one or more tiles. There can be only one structure on a tile.
type Structure interface {
	GetStructure() *BaseStructure
}

// All structure types embed
type BaseStructure struct {
	Type     StructureType
	Variant  int
	Size     [2]int // Width and height in tiles
	Position Position
	Rotation int      // 0, 90, 180, 270 degrees
	MoveCost MoveCost // Cost to move through this structure
}

// Plants grow and can produce edible or craft materials (fruits, wood, etc.)
type PlantStructure struct {
	BaseStructure
	GrowthStage     int // 0-100
	ProductionStage int // 0-100
	GrowthRate      int // How many growth stages per update
	ProductionRate  int // How many production stages per update
	Produces        PlantProduction
}

type PlantProduction struct {
	Type    ItemType
	Variant int
}

// Items are small objects that can be on a tile, on a character, or in a container
// They can be owned, traded, discarded, consumed, used for crafting, etc.
type Item struct {
	Type          ItemType
	Variant       int
	OnTile        *Position
	InInventoryOf *Character
	Efficiency    int // for food it's nutrition value
}
