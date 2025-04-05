package ng

func (b *Plant) GetStructure() *BaseStructure {
	return &b.BaseStructure
}

func (sim *Simulation) UpdatePlants() {
	for _, region := range sim.World {
		for _, plant := range region.Plants {
			plant.Update()
			if plant.Produces.Type != NoItem && plant.ProductionStage >= 100 {
				item := MakeItem(plant.Produces.Type, plant.Produces.Variant)
				sim.SpawnItem(item, plant.Region, plant.Position[0], plant.Position[1])
			}
		}
	}
}

func (plant *Plant) Update() {
	if plant.GrowthStage < 100 {
		plant.GrowthStage += plant.GrowthRate
	}
	if plant.ProductionStage <= 100 {
		plant.ProductionStage += plant.ProductionRate
	}
	if plant.ProductionStage >= 100+plant.ProductionRate {
		plant.ProductionStage = 0
	}
}
