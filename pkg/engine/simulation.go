package engine

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	return &Simulation{
		Time:  0,
		Speed: 60,
	}
}

func (sim *Simulation) Update() {
	sim.Time++
}
