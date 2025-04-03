package ng

import "eramstein/thurigen/pkg/config"

// Simulation represents the main simulation state
type Simulation struct {
	Paused     bool
	Speed      int // how many frames until next sim update
	Time       int // in minutes since the start of the simulation
	World      []*Region
	Structures []Structure
}

type Region struct {
	Tiles [config.RegionSize][config.RegionSize]Tile
}

type Tile struct {
	Terrain    TerrainType
	Surface    SurfaceType
	Volume     VolumeType
	Occupation *TileOccupation
	MoveCost   MoveCost // Cached for pathfinding
}

type TileOccupation struct {
	Structure  Structure
	IsMainTile bool // True if this is the main tile of the object
}

// Structures occupy one or more tiles. There can be only one structure on a tile.
type Structure interface {
	GetBase() *BaseStructure
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
	*BaseStructure
	GrowthStage     int // 0-100
	ProductionStage int // 0-100
}

func (b *PlantStructure) GetBase() *BaseStructure {
	return b.BaseStructure
}
