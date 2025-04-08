package ng

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter("Henry", [2]int{30, 30})
}

func (sim *Simulation) RemoveCharacter(character *Character) {
	for i, c := range sim.Characters {
		if c == character {
			sim.Characters = append(sim.Characters[:i], sim.Characters[i+1:]...)
		}
	}
}

func (sim *Simulation) MakeCharacter(name string, position [2]int) *Character {
	character := &Character{
		ID:        len(sim.Characters),
		Name:      name,
		Position:  position,
		Inventory: []*Item{},
	}
	sim.Characters = append(sim.Characters, character)
	return character
}
