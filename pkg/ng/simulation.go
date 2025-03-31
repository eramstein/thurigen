package ng

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	return &Simulation{
		Speed: 60,
		World: GetInitialWorld(),
	}
}

func (sim *Simulation) Update() {
	sim.Time++
}
