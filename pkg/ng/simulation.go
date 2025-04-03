package ng

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	sim := &Simulation{
		Speed: 60,
	}
	sim.InitWorld()
	return sim
}

func (sim *Simulation) Update() {
	sim.Time++
	for _, structure := range sim.Structures {
		if plant, ok := structure.(*PlantStructure); ok {
			plant.GrowthStage += 10
			if plant.GrowthStage >= 100 {
				plant.GrowthStage = 0
			}
		}
	}
}
