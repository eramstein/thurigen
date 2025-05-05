package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

func (sim *Simulation) UpdateObjectives(character *Character) {
	// periodically un-stuck all objectives to try again
	if sim.Time%config.CharacterObjectiveResetInterval == 0 {
		for _, objective := range character.Objectives {
			objective.Stuck = false
		}
	}

	if character.Needs.Food >= 50 && !character.HasObjective(EatObjective) {
		sim.AddObjective(character, EatObjective, 0)
	}

	if character.Needs.Water >= 50 && !character.HasObjective(DrinkObjective) {
		sim.AddObjective(character, DrinkObjective, 0)
	}

	if character.Needs.Sleep >= 50 && !character.HasObjective(SleepObjective) {
		sim.AddObjective(character, SleepObjective, 0)
	}

	// TODO: better trigger for building a house
	if character.Wants.Confort.SleepConditions <= 0 && !character.HasObjective(BuildObjective) {
		createdObjective := sim.AddObjective(character, BuildObjective, int(BuildHouse))
		sim.PlanObjectiveTasks(character, createdObjective)
	}
}

func (sim *Simulation) AddObjective(character *Character, objectiveType ObjectiveType, variant int) (createdObjective *Objective) {
	fmt.Printf("Adding objective %v %v\n", character.Name, objectiveType)
	objective := &Objective{
		Type:    objectiveType,
		Variant: variant,
		Plan:    []Task{},
	}
	character.Objectives = append(character.Objectives, objective)
	return objective
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
	fmt.Printf("Completing objective %v %v\n", character.Name, objective.Type)
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
	case BuildObjective:
		if len(objective.Plan) == 0 {
			character.CompleteObjective(objective)
		}
	}
}

// Get the top non-stuck priority objective (lowest ObjectiveType is highest priority)
func (sim *Simulation) GetTopPriorityObjective(character *Character) *Objective {
	if len(character.Objectives) == 0 {
		return nil
	}
	lowestObjective := character.Objectives[0]
	for _, objective := range character.Objectives {
		if !objective.Stuck && objective.Type < lowestObjective.Type {
			lowestObjective = objective
		}
	}
	return lowestObjective
}

func (sim *Simulation) PlanObjectiveTasks(character *Character, objective *Objective) {
	switch objective.Type {
	case BuildObjective:
		sim.PlanBuildingTasks(character, objective)
	}
}

func (sim *Simulation) PlanBuildingTasks(character *Character, objective *Objective) {
	switch BuildObjectiveVariant(objective.Variant) {
	case BuildHouse:
		sim.PlanHouseBuildingTasks(character, objective)
	}
}
