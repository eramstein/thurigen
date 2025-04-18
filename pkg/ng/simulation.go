package ng

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	loadData()
	sim := &Simulation{
		Speed: 60,
	}
	sim.InitWorld()
	sim.InitCharacters()
	return sim
}

func (sim *Simulation) Update() {
	sim.Time++
	sim.UpdatePlants()
	sim.UpdateCharacters()
}

func loadData() {
	LoadStructuresConfigs()
	LoadItemsConfigs()
}
