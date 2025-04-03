package ng

type TerrainType int

const (
	Dirt TerrainType = iota
	Rock
	Sand
	Water
)

type SurfaceType int

const (
	NoSurface SurfaceType = iota
	Grass
	WoodSurface
)

type VolumeType int

const (
	NoVolume VolumeType = iota
	RockVolume
	WoodVolume
)

// MoveCost represents the cost of moving through a tile
type MoveCost float64

const (
	DefaultMoveCost   MoveCost = 1.0  // Normal movement cost
	DifficultMoveCost MoveCost = 2.0  // Increased cost for difficult terrain
	ImpassableCost    MoveCost = -1.0 // Represents an impassable tile
)

type StructureType int

const (
	NoStructure StructureType = iota
	Tree
	Wall
	Furniture
)
