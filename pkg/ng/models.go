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
	Items      []Item // TODO: check if map by item type would be better (if we loop often by item type), or map by item id if we do a lot of random access
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
	Objectives  []*Objective
	Ambitions   []*Ambition
	Needs       Needs
	Wants       Wants
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

type Wants struct {
	Confort Confort
	Safety  Safety
}

type Confort struct {
	Total           int
	SleepConditions int // -10 to 10, how confortable last sleep was
}

type Safety struct {
	HasHouse bool
}

type Task struct {
	ID             uint64
	Type           TaskType
	ProductType    int // optional, precises the task is producing based on the Task Type, for example for bulding tasks it's the StructureType to build (e.g. Wall)
	ProductVariant int // optional, further precises the task's product by providing a variant (e.g. Wooden Wall, Stone Wall)
	Objective      *Objective
	Progress       float32 // by default, 0 to 1, as percent of task already done, but can be used otherwise like for movement
	Target         any     // optional target, for example for bulding tasks it's tile on which to build, for eating tasks it's the food item to eat
	MaterialSource *Item   // optional, for bulding tasks it's the material item to use
}

type Objective struct {
	Type    ObjectiveType
	Stuck   bool
	Variant int    // optional, further precises the objective by providing a variant (e.g. "build a house")
	Plan    []Task // A Plan is a list of tasks needed to complete an objective. NOTE, TODO?: risk of circular reference, plan tasks must not point to the objective
	Target  any    // optional target, for example for bulding objectives it's an edifice to complete
}

type Ambition struct {
	Description string
}

type Region struct {
	Tiles    [config.RegionSize][config.RegionSize]Tile // Tiles[X][Y]
	Plants   []*PlantStructure
	Edifices []*Edifice
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
	ID           uint64
	Type         StructureType
	Variant      int
	Size         [2]int // Width and height in tiles
	Position     Position
	Rotation     int      // 0, 90, 180, 270 degrees
	MoveCost     MoveCost // Cost to move through this structure
	MaterialType MaterialType
}

type BuildingStructure struct {
	BaseStructure
	Completion      int // 0-100
	PartOfEdificeID uint64
}

// Plants grow and can produce edible or craft materials (fruits, wood, etc.)
type PlantStructure struct {
	BaseStructure
	PlantType       PlantVariant
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
	ID            uint64
	Type          ItemType
	Variant       int
	OnTile        *Position
	InInventoryOf uint64 // character id
	OwnedBy       uint64 // character id
	Efficiency    int    // for food it's nutrition value
	Durability    int    // for materials it's how many builds they can support
}

type Edifice struct {
	ID         uint64
	Type       EdificeType
	OwnedBy    uint64 // character id
	Buildings  []*BuildingStructure
	IsComplete bool
}
