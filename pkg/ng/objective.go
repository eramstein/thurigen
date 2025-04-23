package ng

import "fmt"

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

func (sim *Simulation) AddObjective(character *Character, objectiveType ObjectiveType) {
	objective := &Objective{
		Type: objectiveType,
	}
	character.Objectives = append(character.Objectives, objective)
	sim.PlanTasks(character, objective)
}

func (character *Character) HasObjective(objectiveType ObjectiveType) bool {
	for _, objective := range character.Objectives {
		if objective.Type == objectiveType {
			return true
		}
	}
	return false
}

func (character *Character) CompleteObjective(objective *Objective) {
	fmt.Println("Completing objective", objective.Type, objective)
	for i, charObjective := range character.Objectives {
		if charObjective == objective {
			character.Objectives = append(character.Objectives[:i], character.Objectives[i+1:]...)
		}
	}
}

func (sim *Simulation) CheckIfObjectiveIsAchieved(character *Character, objective *Objective) {
	switch objective.Type {
	case EatObjective:
		if character.Needs.Food < 40 {
			character.CompleteObjective(objective)
		}
	case DrinkObjective:
		if character.Needs.Water < 40 {
			character.CompleteObjective(objective)
		}
	case SleepObjective:
		if character.Needs.Sleep < 10 {
			character.CompleteObjective(objective)
		}
	}
}

// Get the top priority objective (lowest ObjectiveType is highest priority)
func (sim *Simulation) GetTopPriorityObjective(character *Character) *Objective {
	if len(character.Objectives) == 0 {
		return nil
	}
	lowestObjective := character.Objectives[0]
	for _, objective := range character.Objectives {
		if objective.Type < lowestObjective.Type {
			lowestObjective = objective
		}
	}
	return lowestObjective
}
