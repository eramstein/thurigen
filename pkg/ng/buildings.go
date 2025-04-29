package ng

func (b *WallStructure) GetStructure() *BaseStructure {
	return &b.BaseStructure
}

func (sim *Simulation) AddWall(region int, x, y int, material MaterialType, completion int) {
	sim.AddStructure(&WallStructure{
		BaseStructure: BaseStructure{
			Type:     Wall,
			Variant:  int(material),
			Size:     [2]int{1, 1},
			MoveCost: ImpassableCost,
			Position: Position{Region: region, X: x, Y: y},
		},
		Completion: completion,
	})
}
