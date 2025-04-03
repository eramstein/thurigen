package ng

func (sim *Simulation) UpdatePlants() {
	for _, region := range sim.World {
		for _, plant := range region.Plants {
			plant.Update()
		}
	}
}

func (plant *Plant) Update() {
	plant.GrowthStage += 10
}
