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

type MaterialType int

const (
	NoMaterial MaterialType = iota
	RockMaterial
	WoodMaterial
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
	Building
	Furniture
)

type BuildingVariant int

const (
	NoBuildingVariant BuildingVariant = iota
	Wall
	Door
	Window
)

type PlantVariant int

const (
	NoPlantVariant PlantVariant = iota
	Tree
)

type ItemType int

const (
	NoItem ItemType = iota
	Food
	Material
	Tool
)

type TaskType int

const (
	NoTaskType TaskType = iota
	Move
	Eat
	Drink
	Sleep
	PickUp
	Build
	Chop
)

type ObjectiveType int

// Sorted by priority (lowest number is highest priority)
const (
	NoObjective ObjectiveType = iota
	DrinkObjective
	EatObjective
	SleepObjective
	BuildObjective
)

type EdificeType int

const (
	NoEdifice EdificeType = iota
	House
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
	case Building:
		return "Building"
	case Furniture:
		return "Furniture"
	default:
		return "Unknown Structure"
	}
}

// String returns a human-readable description of the BuildingVariant
func (b BuildingVariant) String() string {
	switch b {
	case NoBuildingVariant:
		return "No Building"
	case Wall:
		return "Wall"
	case Door:
		return "Door"
	case Window:
		return "Window"
	default:
		return "Unknown Building"
	}
}

// String returns a human-readable description of the PlantType
func (p PlantVariant) String() string {
	switch p {
	case NoPlantVariant:
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

func (t TaskType) String() string {
	switch t {
	case NoTaskType:
		return "No Task"
	case Move:
		return "Move Task"
	case Eat:
		return "Eat Task"
	case Drink:
		return "Drink Task"
	case Sleep:
		return "Sleep Task"
	case PickUp:
		return "PickUp Task"
	case Build:
		return "Build Task"
	case Chop:
		return "Chop Task"
	default:
		return "Unknown Task"
	}
}

func (o ObjectiveType) String() string {
	switch o {
	case NoObjective:
		return "No Objective"
	case EatObjective:
		return "Eat Objective"
	case DrinkObjective:
		return "Drink Objective"
	case SleepObjective:
		return "Sleep Objective"
	case BuildObjective:
		return "Build Objective"
	default:
		return "Unknown Objective"
	}
}

func (m MaterialType) String() string {
	switch m {
	case NoMaterial:
		return "No Material"
	case RockMaterial:
		return "Rock Material"
	case WoodMaterial:
		return "Wood Material"
	default:
		return "Unknown Material"
	}
}

func (e EdificeType) String() string {
	switch e {
	case NoEdifice:
		return "No Edifice"
	case House:
		return "House"
	default:
		return "Unknown Edifice"
	}
}
