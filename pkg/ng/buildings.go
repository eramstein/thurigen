package ng

func (b *BuildingStructure) GetStructure() *BaseStructure {
	return &b.BaseStructure
}

func (sim *Simulation) AddWall(region int, x, y int, material MaterialType, completion int) {
	sim.AddStructure(&BuildingStructure{
		BaseStructure: BaseStructure{
			Type:         Building,
			Variant:      int(Wall),
			Size:         [2]int{1, 1},
			MoveCost:     ImpassableCost,
			Position:     Position{Region: region, X: x, Y: y},
			MaterialType: material,
		},
		Completion: completion,
	})
}
