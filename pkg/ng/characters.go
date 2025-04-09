package ng

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter("Henry", 0, [2]int{30, 30})
}

func (sim *Simulation) UpdateCharacters() {
	for _, character := range sim.Characters {
		character.Update()
	}
}

func (character *Character) Update() {
	character.UpdateNeeds()
}

func (character *Character) UpdateNeeds() {
	character.Needs.Food += 1
	character.Needs.Water += 1
	character.Needs.Sleep += 1
}

func (sim *Simulation) RemoveCharacter(character *Character) {
	for i, c := range sim.Characters {
		if c == character {
			sim.Characters = append(sim.Characters[:i], sim.Characters[i+1:]...)
		}
	}
}

func (sim *Simulation) MakeCharacter(name string, region int, position [2]int) *Character {
	character := &Character{
		ID:        len(sim.Characters),
		Name:      name,
		Region:    region,
		Position:  position,
		Inventory: []*Item{},
	}
	sim.Characters = append(sim.Characters, character)
	sim.World[region].Tiles[position[0]][position[1]].Character = character
	return character
}
