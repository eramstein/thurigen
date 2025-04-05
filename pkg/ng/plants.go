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
	plant := &PlantStructure{
		BaseStructure: BaseStructure{
			Type:     StructureType(plantType),
			Variant:  variant,
			Size:     [2]int{1, 1},
			Position: [2]int{x, y},
			Region:   region,
			Rotation: 0,
			MoveCost: DifficultMoveCost,
		},
		GrowthStage:     0,
		ProductionStage: 0,
		GrowthRate:      10,
		ProductionRate:  20,
	}
	if plantType == Tree && variant == 1 {
		plant.Produces = PlantProduction{
			Type:    Food,
			Variant: int(Apple),
		}
	}
	return plant
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
