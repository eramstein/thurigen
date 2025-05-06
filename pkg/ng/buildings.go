package ng

import "fmt"

func (b *BuildingStructure) GetStructure() *BaseStructure {
	return &b.BaseStructure
}

func (sim *Simulation) AddBuilding(position Position, buildingVariant BuildingVariant, material MaterialType, completion int, edifice *Edifice) {
	if sim.IsOccupied(position) {
		fmt.Printf("WARNING: Building %v at %v is occupied\n", buildingVariant, position)
		return
	}
	var moveCost MoveCost
	if buildingVariant == Door {
		moveCost = DefaultMoveCost
	} else {
		moveCost = ImpassableCost
	}
	building := BuildingStructure{
		BaseStructure: BaseStructure{
			ID:           getNextID(),
			Type:         Building,
			Variant:      int(buildingVariant),
			Size:         [2]int{1, 1},
			MoveCost:     moveCost,
			Position:     position,
			MaterialType: material,
		},
		Completion:      completion,
		PartOfEdificeID: edifice.ID,
	}
	sim.AddStructureOccupation(&building)

	edifice.Buildings = append(edifice.Buildings, &building)
}

func (sim *Simulation) AddEdifice(region int, edifice Edifice) {
	sim.World[region].Edifices = append(sim.World[region].Edifices, &edifice)
}
