package ng

import (
	"eramstein/thurigen/pkg/config"
)

func (sim *Simulation) WorkOnPriorityTask(character *Character) {
	if len(character.Tasks) == 0 {
		return
	}
	// tasks are sorted by priority, work on the highest priority task
	task := character.Tasks[0]
	switch task.Type {
	case Move:
		sim.FollowPath(character, task, false)
	case Eat:
		sim.Eat(character, task)
	}
}

func (sim *Simulation) CompleteTask(character *Character, task *Task) {
	for i, t := range character.Tasks {
		if t == task {
			character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
			break
		}
	}
	sim.CheckIfObjectiveIsAchieved(character, task.Objective)
	if len(character.Objectives) > 0 {
		sim.PlanTasks(character, character.Objectives[0])
	}
}

func (sim *Simulation) CheckIfObjectiveIsAchieved(character *Character, objective *Objective) {
	if objective.Type == EatObjective {
		if character.Needs.Food < 40 {
			character.CompleteObjective(objective)
		}
	}
}

// Set next task required to achieve an eat objective
func (sim *Simulation) PlanEatingTasks(character *Character, objective *Objective) {
	// Check if the character has the item in their inventory
	itemInInventory := character.FindInInventory(Food)

	// If the character has the item in their inventory, add a task to eat it
	if itemInInventory != nil {
		character.AddTask(Task{
			Objective: objective,
			Type:      Eat,
			Target:    itemInInventory,
		})
		// If the character is on a tile with a food item, add a task to eat it
	} else if itemOnTile := sim.FindItemInTile(character.Position.Region, character.Position.X, character.Position.Y, Food); itemOnTile != nil {
		character.AddTask(Task{
			Objective: objective,
			Type:      Eat,
			Target:    itemOnTile,
		})
	} else {
		// If no food on tile, find the closest food item and add a task to go to it
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2, Food)
		if closestItem != nil {
			base := (*closestItem).GetItem()
			path := sim.World[character.Position.Region].FindPath(character.Position.X, character.Position.Y, base.OnTile.X, base.OnTile.Y)
			character.Path = &path
			character.AddTask(Task{
				Objective: objective,
				Type:      Move,
				Target:    base.OnTile,
			})
		}
	}
}

func (sim *Simulation) PlanDrinkingTasks(character *Character, objective *Objective) {
}

func (sim *Simulation) PlanSleepingTasks(character *Character, objective *Objective) {
}
