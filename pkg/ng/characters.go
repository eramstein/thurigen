package ng

import "eramstein/thurigen/pkg/config"

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter("Henry", 0, [2]int{30, 30})
}

func (sim *Simulation) UpdateCharacters() {
	if sim.Time%config.CharacterNeedsUpdateInterval == 0 {
		for _, character := range sim.Characters {
			character.UpdateNeeds()
		}
	}
	if sim.Time%config.CharacterObjectiveUpdateInterval == 0 {
		for _, character := range sim.Characters {
			character.UpdateObjectives()
		}
	}
}

func (character *Character) UpdateNeeds() {
	character.Needs.Food += 1
	character.Needs.Water += 1
	character.Needs.Sleep += 1
}

func (character *Character) UpdateObjectives() {
	if character.Needs.Food >= 50 && !character.HasObjective(EatObjective) {
		character.AddObjective(EatObjective)
	}

	if character.Needs.Water >= 50 && !character.HasObjective(DrinkObjective) {
		character.AddObjective(DrinkObjective)
	}

	if character.Needs.Sleep >= 50 && !character.HasObjective(SleepObjective) {
		character.AddObjective(SleepObjective)
	}
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
		Needs:     Needs{Food: 49, Water: 0, Sleep: 0},
	}
	sim.Characters = append(sim.Characters, character)
	sim.World[region].Tiles[position[0]][position[1]].Character = character
	return character
}

func (character *Character) AddObjective(objectiveType ObjectiveType) {
	character.Objectives = append(character.Objectives, &Objective{
		Type: objectiveType,
	})
}

func (character *Character) HasObjective(objectiveType ObjectiveType) bool {
	for _, objective := range character.Objectives {
		if objective.Type == objectiveType {
			return true
		}
	}
	return false
}
