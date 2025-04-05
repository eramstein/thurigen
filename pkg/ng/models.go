package ng

import "eramstein/thurigen/pkg/config"

// Simulation represents the main simulation state
type Simulation struct {
	Paused bool
	Speed  int // how many frames until next sim update
	Time   int // in minutes since the start of the simulation
	World  []*Region
	Items  []*Item
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
	Items      []*Item  // Items lying on the tile
	MoveCost   MoveCost // Cached for pathfinding
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
	Position [2]int // Position of the main tile
	Region   int
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
type Item interface {
	GetItem() *BaseItem
}

// All item types embed
type BaseItem struct {
	Type    ItemType
	Variant int
	OnTile  [3]int // If on ground: Region / X / Y
}

// Food items can be consumed by characters to restore nutrition
type FoodItem struct {
	BaseItem
	Nutrition int // 0-100
}
