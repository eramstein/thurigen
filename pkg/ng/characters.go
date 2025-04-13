package ng

import "eramstein/thurigen/pkg/config"

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter("Henry", Position{Region: 0, X: 30, Y: 30})
}

func (sim *Simulation) UpdateCharacters() {
	if sim.Time%config.CharacterNeedsUpdateInterval == 0 {
		for _, character := range sim.Characters {
			character.UpdateNeeds()
		}
	}
	if sim.Time%config.CharacterObjectiveUpdateInterval == 0 {
		for _, character := range sim.Characters {
			sim.UpdateObjectives(character)
		}
	}
}

func (character *Character) UpdateNeeds() {
	character.Needs.Food += 1
	character.Needs.Water += 1
	character.Needs.Sleep += 1
}

func (sim *Simulation) UpdateObjectives(character *Character) {
	if character.Needs.Food >= 50 && !character.HasObjective(EatObjective) {
		sim.AddObjective(character, EatObjective)
	}

	if character.Needs.Water >= 50 && !character.HasObjective(DrinkObjective) {
		sim.AddObjective(character, DrinkObjective)
	}

	if character.Needs.Sleep >= 50 && !character.HasObjective(SleepObjective) {
		sim.AddObjective(character, SleepObjective)
	}
}

func (sim *Simulation) RemoveCharacter(character *Character) {
	for i, c := range sim.Characters {
		if c == character {
			sim.Characters = append(sim.Characters[:i], sim.Characters[i+1:]...)
		}
	}
}

func (sim *Simulation) MakeCharacter(name string, pos Position) *Character {
	character := &Character{
		ID:        len(sim.Characters),
		Name:      name,
		Position:  pos,
		Inventory: []*Item{},
		Needs:     Needs{Food: 49, Water: 0, Sleep: 0},
	}
	sim.Characters = append(sim.Characters, character)
	sim.World[pos.Region].Tiles[pos.X][pos.Y].Character = character
	return character
}

func (sim *Simulation) AddObjective(character *Character, objectiveType ObjectiveType) {
	objective := &Objective{
		Type: objectiveType,
	}
	character.Objectives = append(character.Objectives, objective)
	sim.PlanTasks(character, objective)
}

func (sim *Simulation) PlanTasks(character *Character, objective *Objective) {
	switch objective.Type {
	case EatObjective:
		sim.PlanEatingTasks(character, objective)
	case DrinkObjective:
		sim.PlanDrinkingTasks(character, objective)
	case SleepObjective:
		sim.PlanSleepingTasks(character, objective)
	}
}

func (character *Character) AddTask(task Task) {
	character.Tasks = append(character.Tasks, &task)
}

func (character *Character) HasObjective(objectiveType ObjectiveType) bool {
	for _, objective := range character.Objectives {
		if objective.Type == objectiveType {
			return true
		}
	}
	return false
}
