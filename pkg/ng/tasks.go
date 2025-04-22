package ng

import (
	"eramstein/thurigen/pkg/config"
)

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
	case Drink:
		sim.Drink(character, task)
	case Sleep:
		sim.Sleep(character, task)
	}
	if task.Progress >= 100 {
		sim.CompleteTask(character, task)
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
			path := sim.World[character.Position.Region].FindPath(character.Position.X, character.Position.Y, closestItem.OnTile.X, closestItem.OnTile.Y, 0)
			character.Path = &path
			character.AddTask(Task{
				Objective: objective,
				Type:      Move,
				Target:    closestItem.OnTile,
			})
		}
	}
}

func (sim *Simulation) PlanDrinkingTasks(character *Character, objective *Objective) {
	closestWater := sim.ScanForTile(character.Position, config.RegionSize/2, Water)
	if closestWater == nil {
		return
	}
	if IsAdjacent(character.Position.X, character.Position.Y, closestWater.X, closestWater.Y) {
		character.AddTask(Task{
			Objective: objective,
			Type:      Drink,
			Target:    closestWater,
		})
	} else {
		path := sim.World[character.Position.Region].FindPath(character.Position.X, character.Position.Y, closestWater.X, closestWater.Y, 1)
		if len(path) > 0 {
			character.Path = &path
			character.AddTask(Task{
				Objective: objective,
				Type:      Move,
				Target:    path[len(path)-1],
			})
		}
	}
}

func (sim *Simulation) PlanSleepingTasks(character *Character, objective *Objective) {
	character.AddTask(Task{
		Objective: objective,
		Type:      Sleep,
	})
}
