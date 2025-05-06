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
				sim.SpawnItem(&item, plant.Position)
			}
		}
	}
}

func (sim *Simulation) SpawnPlant(region int, x, y int, variant int, plantType PlantVariant) {
	structureConfig := GetStructureConfig(Plant, variant)
	plant, ok := structureConfig.Structure.(*PlantStructure)
	if !ok {
		fmt.Printf("Error: Structure is not a PlantStructure: %T\n", structureConfig.Structure)
		return
	}

	newPlant := &PlantStructure{
		BaseStructure: BaseStructure{
			ID:       getNextID(),
			Type:     Plant,
			Variant:  variant,
			Size:     [2]int{1, 1},
			MoveCost: plant.MoveCost,
			Position: Position{Region: region, X: x, Y: y},
		},
		PlantType:       plantType,
		GrowthRate:      plant.GrowthRate,
		ProductionRate:  plant.ProductionRate,
		Produces:        plant.Produces,
		ProductionStage: 99,
	}
	sim.World[region].Plants = append(sim.World[region].Plants, newPlant)
	sim.AddStructureOccupation(newPlant)
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

func (sim *Simulation) ChopTree(tree *PlantStructure) {
	sim.SpawnItem(&Item{Type: Material, Variant: int(WoodMaterial), Durability: tree.GrowthStage}, tree.Position)
	sim.RemovePlant(tree)
}

func (sim *Simulation) RemovePlant(plant *PlantStructure) {
	sim.RemoveStructureOccupation(plant)
	sim.World[plant.Position.Region].Plants = removeFromSlice(sim.World[plant.Position.Region].Plants, plant)
}
