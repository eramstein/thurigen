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
	Plant
	Wall
	Furniture
)

type PlantType int

const (
	NoPlant PlantType = iota
	Tree
)

type ItemType int

const (
	NoItem ItemType = iota
	Food
	Material
	Tool
)

type MaterialType int

const (
	NoMaterialType MaterialType = iota
	WoodMaterial
	StoneMaterial
)

type TaskType int

const (
	NoTask TaskType = iota
	GotoTask
	EatTask
)

type ObjectiveType int

const (
	NoObjective ObjectiveType = iota
	EatObjective
	BuildObjective
)

// String returns a human-readable description of the TerrainType
func (t TerrainType) String() string {
	switch t {
	case Dirt:
		return "Dirt"
	case Rock:
		return "Rock"
	case Sand:
		return "Sand"
	case Water:
		return "Water"
	default:
		return "Unknown Terrain"
	}
}

// String returns a human-readable description of the SurfaceType
func (s SurfaceType) String() string {
	switch s {
	case NoSurface:
		return "No Surface"
	case Grass:
		return "Grass"
	case WoodSurface:
		return "Wood Surface"
	default:
		return "Unknown Surface"
	}
}

// String returns a human-readable description of the VolumeType
func (v VolumeType) String() string {
	switch v {
	case NoVolume:
		return "No Volume"
	case RockVolume:
		return "Rock Volume"
	case WoodVolume:
		return "Wood Volume"
	default:
		return "Unknown Volume"
	}
}

// String returns a human-readable description of the StructureType
func (s StructureType) String() string {
	switch s {
	case NoStructure:
		return "No Structure"
	case Plant:
		return "Plant"
	case Wall:
		return "Wall"
	case Furniture:
		return "Furniture"
	default:
		return "Unknown Structure"
	}
}

// String returns a human-readable description of the PlantType
func (p PlantType) String() string {
	switch p {
	case NoPlant:
		return "No Plant"
	case Tree:
		return "Tree"
	default:
		return "Unknown Plant"
	}
}

// String returns a human-readable description of the ItemType
func (i ItemType) String() string {
	switch i {
	case NoItem:
		return "No Item"
	case Food:
		return "Food"
	case Material:
		return "Material"
	case Tool:
		return "Tool"
	default:
		return "Unknown Item"
	}
}

// String returns a human-readable description of the MaterialType
func (m MaterialType) String() string {
	switch m {
	case NoMaterialType:
		return "No Material Type"
	case WoodMaterial:
		return "Wood Material"
	case StoneMaterial:
		return "Stone Material"
	default:
		return "Unknown Material Type"
	}
}
