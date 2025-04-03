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
	sim.UpdatePlants()
}
