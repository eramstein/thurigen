package ng

import "eramstein/thurigen/pkg/config"

// Simulation represents the main simulation state
type Simulation struct {
	Paused     bool
	Speed      int // how many frames until next sim update
	Time       int // in minutes since the start of the simulation
	World      []*Region
	Structures []*BaseStructure
}

type Region struct {
	Tiles [config.RegionSize][config.RegionSize]Tile
}

type Tile struct {
	Terrain    TerrainType
	Surface    SurfaceType
	Volume     VolumeType
	Occupation *TileOccupation
	MoveCost   MoveCost // Cached for path
}

type TileOccupation struct {
	Structure  *BaseStructure // Pointer to a structure occupying this tile
	IsMainTile bool           // True if this is the main tile of the object
}

type BaseStructure struct {
	Type     StructureType
	Variant  int
	Size     [2]int // Width and height in tiles
	Position [2]int // Position of the main tile
	Region   int
	Rotation int      // 0, 90, 180, 270 degrees
	MoveCost MoveCost // Cost to move through this structure
}

type FurnitureStructure struct {
	BaseStructure
	IsDecorative    bool // Indicates if the furniture is purely decorative
	CanContainItems bool // Indicates if the furniture can hold items
	StorageCapacity int  // Number of items it can hold (if applicable)
}

type TreeStructure struct {
	BaseStructure
	IsFruitBearing bool   // Indicates if the tree produces fruit
	FruitType      string // Type of fruit produced (if applicable)
}
