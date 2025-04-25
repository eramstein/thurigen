package ng

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	sim := &Simulation{
		Speed: 60,
		Calendar: Calendar{
			Minute: 0,
			Hour:   5,
			Day:    0,
		},
	}
	sim.InitWorld()
	sim.InitCharacters()
	return sim
}

func (sim *Simulation) Update() {
	sim.UpdateTime()
	sim.UpdatePlants()
	sim.UpdateCharacters()
}

func (sim *Simulation) UpdateTime() {
	sim.Time++
	sim.Calendar.Minute++
	if sim.Calendar.Minute%60 == 0 {
		sim.Calendar.Hour++
		sim.Calendar.Minute = 0
	}
	if sim.Calendar.Hour >= 24 {
		sim.Calendar.Hour = 0
		sim.Calendar.Day++
	}
}
