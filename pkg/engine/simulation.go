package engine

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	return &Simulation{
		Speed: 60,
	}
}

func (sim *Simulation) Update() {
	sim.Time++
}
