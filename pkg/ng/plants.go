package ng

func (b *PlantStructure) GetStructure() *BaseStructure {
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

func (sim *Simulation) SpawnPlant(plant *PlantStructure) {
	sim.World[plant.Region].Plants = append(sim.World[plant.Region].Plants, plant)
	sim.AddStructure(plant)
}

func MakePlant(region int, x, y int, plantType PlantType, variant int) *PlantStructure {
	plant := GetPlantConfig(plantType, variant)
	plant.BaseStructure.Region = region
	plant.BaseStructure.Position = [2]int{x, y}
	return &plant.PlantStructure
}

func (plant *PlantStructure) Update() {
	if plant.GrowthStage < 100 {
		plant.GrowthStage += plant.GrowthRate
	}
	if plant.GrowthStage >= 100 && plant.ProductionStage <= 100 {
		plant.ProductionStage += plant.ProductionRate
	}
	if plant.ProductionStage >= 100+plant.ProductionRate {
		plant.ProductionStage = 0
	}
}
