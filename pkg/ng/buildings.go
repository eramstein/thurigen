package ng

import "fmt"

func (b *BuildingStructure) GetStructure() *BaseStructure {
	return &b.BaseStructure
}

func (sim *Simulation) AddBuilding(region int, x, y int, buildingVariant BuildingVariant, material MaterialType, completion int) {
	if sim.IsOccupied(Position{Region: region, X: x, Y: y}) {
		fmt.Printf("WARNING: Building %v at %v is occupied\n", buildingVariant, Position{Region: region, X: x, Y: y})
		return
	}
	sim.AddStructure(&BuildingStructure{
		BaseStructure: BaseStructure{
			Type:         Building,
			Variant:      int(buildingVariant),
			Size:         [2]int{1, 1},
			MoveCost:     ImpassableCost,
			Position:     Position{Region: region, X: x, Y: y},
			MaterialType: material,
		},
		Completion: completion,
	})
}
