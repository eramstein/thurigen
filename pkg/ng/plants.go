package ng

import "fmt"

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

func (sim *Simulation) SpawnPlant(region int, x, y int, variant int) {
	structureConfig := GetStructureConfig(Plant, variant)
	plant, ok := structureConfig.Structure.(*PlantStructure)
	if !ok {
		fmt.Printf("Error: Structure is not a PlantStructure: %T\n", structureConfig.Structure)
		return
	}

	newPlant := &PlantStructure{
		BaseStructure: BaseStructure{
			Type:     Plant,
			Variant:  variant,
			Size:     [2]int{1, 1},
			MoveCost: plant.MoveCost,
		},
		GrowthRate:     plant.GrowthRate,
		ProductionRate: plant.ProductionRate,
		Produces:       plant.Produces,
	}

	newPlant.Position = [2]int{x, y}
	newPlant.Region = region

	sim.World[region].Plants = append(sim.World[region].Plants, newPlant)
	sim.AddStructure(newPlant)
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
