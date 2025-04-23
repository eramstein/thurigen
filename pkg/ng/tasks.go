package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

// When an objective is added, plan tasks for it
// When a task is completed, plan tasks for the top priority objective
func (sim *Simulation) PlanTasks(character *Character, objective *Objective) {
	// Add tasks for the objective
	switch objective.Type {
	case EatObjective:
		sim.PlanEatingTasks(character, objective)
	case DrinkObjective:
		sim.PlanDrinkingTasks(character, objective)
	case SleepObjective:
		sim.PlanSleepingTasks(character, objective)
	}
	// Set current task to the highest priority task
	character.CurrentTask = sim.GetPriorityTask(character)
}

func (sim *Simulation) GetPriorityTask(character *Character) *Task {
	if len(character.Tasks) == 0 {
		return nil
	}
	// if the character already has a current task, and it's not finished, don't change it
	// TODO: handle vital situations when character has to drop its current task
	if character.CurrentTask != nil && character.CurrentTask.Progress < 100 && character.CurrentTask.Progress > 0 {
		return character.CurrentTask
	}
	// else, return first task that matches the top priority objective
	topObjective := sim.GetTopPriorityObjective(character)
	if topObjective == nil {
		return nil
	}
	for _, task := range character.Tasks {
		if task.Objective.Type == topObjective.Type {
			return task
		}
	}
	return nil
}

func (sim *Simulation) WorkOnPriorityTask(character *Character) {
	if character.CurrentTask == nil {
		return
	}
	task := character.CurrentTask
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
	fmt.Printf("Completing task: %v\n %v\n", &task, task.Objective)
	if character.CurrentTask == task {
		character.CurrentTask = nil
	}
	for i, t := range character.Tasks {
		if t.ID == task.ID {
			character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
			break
		}
	}
	sim.CheckIfObjectiveIsAchieved(character, task.Objective)
	topObjective := sim.GetTopPriorityObjective(character)
	if topObjective != nil {
		sim.PlanTasks(character, topObjective)
	}
}

// Set next task required to achieve an eat objective
func (sim *Simulation) PlanEatingTasks(character *Character, objective *Objective) {
	// Remove previous tasks related to eating objectives
	for i, task := range character.Tasks {
		if task.Objective.Type == EatObjective {
			character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
		}
	}
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
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2-1, Food)
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
	// Remove previous tasks related to drinking objectives
	for i, task := range character.Tasks {
		if task.Objective.Type == DrinkObjective {
			character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
		}
	}
	// Go to the closest water tile if needed, then drink
	closestWater := sim.ScanForTile(character.Position, config.RegionSize/2-1, Water)
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
	// Remove previous tasks related to sleeping objectives
	for i, task := range character.Tasks {
		if task.Objective.Type == SleepObjective {
			character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
		}
	}
	character.AddTask(Task{
		Objective: objective,
		Type:      Sleep,
	})
}
