package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

// Create next task for a given objective
// Add it to the character's tasks if it's not nil
func (sim *Simulation) PlanNextTask(character *Character, objective *Objective) (task *Task) {
	fmt.Printf("Planning tasks for %v %v\n", character.Name, objective.Type)
	// Remove previous tasks related to same objective type
	for i, task := range character.Tasks {
		if task.Objective.Type == objective.Type {
			character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
		}
	}
	// Add tasks for the objective
	switch objective.Type {
	case EatObjective:
		task = sim.PlanEatingTasks(character, objective)
	case DrinkObjective:
		task = sim.PlanDrinkingTasks(character, objective)
	case SleepObjective:
		task = sim.PlanSleepingTasks(character, objective)
	}
	if task != nil {
		character.AddTask(task)
	}
	return task
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
	var topObjectiveFirstTask *Task
	for _, task := range character.Tasks {
		if task.Objective.Type == topObjective.Type {
			topObjectiveFirstTask = task
			break
		}
	}
	if topObjectiveFirstTask != nil {
		return topObjectiveFirstTask
	} else {
		newTask := sim.PlanNextTask(character, topObjective)
		if newTask != nil {
			return newTask
		}
	}
	return nil
}

func (sim *Simulation) SetPriorityTask(character *Character) {
	topObjective := sim.GetTopPriorityObjective(character)
	if topObjective != nil {
		sim.PlanNextTask(character, topObjective)
	}
	character.CurrentTask = sim.GetPriorityTask(character)
}

func (sim *Simulation) WorkOnPriorityTask(character *Character) {
	task := character.CurrentTask
	if task == nil {
		fmt.Printf("No task for %v\n", character.Name)
		return
	}
	switch task.Type {
	case Move:
		sim.MoveForTask(character, task)
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
	fmt.Printf("Completing task:  %v %v %v\n", character.Name, task.Type, task.Objective)
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
}

// Set next task required to achieve an eat objective
func (sim *Simulation) PlanEatingTasks(character *Character, objective *Objective) (task *Task) {
	var newTask *Task
	// Check if the character has the item in their inventory
	itemInInventory := character.FindInInventory(Food)
	// If the character has the item in their inventory, add a task to eat it
	if itemInInventory != nil {
		newTask = &Task{
			Objective: objective,
			Type:      Eat,
			Target:    itemInInventory,
		}
		// If the character is on a tile with a food item, add a task to eat it
	} else if itemOnTile := sim.FindItemInTile(character.Position.Region, character.Position.X, character.Position.Y, Food, false); itemOnTile != nil {
		// claim item
		itemOnTile.OwnedBy = character
		// eat it
		newTask = &Task{
			Objective: objective,
			Type:      Eat,
			Target:    itemOnTile,
		}
	} else {
		// If no food on tile, find the closest food item and add a task to go to it
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2-1, Food, true)
		if closestItem != nil {
			// claim item
			closestItem.OwnedBy = character
			// go to it
			newTask = &Task{
				Objective: objective,
				Type:      Move,
				Target:    closestItem.OnTile,
			}
		} else {
			fmt.Printf("No food found for %v\n", character.Name)
		}
	}
	return newTask
}

func (sim *Simulation) PlanDrinkingTasks(character *Character, objective *Objective) (task *Task) {
	var newTask *Task
	// Go to the closest water tile if needed, then drink
	closestWater := sim.ScanForTile(character.Position, config.RegionSize/2-1, Water)
	fmt.Printf("closestWater: %v\n", closestWater)
	if closestWater == nil {
		return
	}
	if IsAdjacent(character.Position.X, character.Position.Y, closestWater.X, closestWater.Y) {
		newTask = &Task{
			Objective: objective,
			Type:      Drink,
			Target:    closestWater,
		}
	} else {
		// stop one tile before the water tile
		// problem: if closestWater is not accessible, there will be no path found and not task added
		path := sim.World[character.Position.Region].FindPath(character.Position.X, character.Position.Y, closestWater.X, closestWater.Y, 1)
		if len(path) > 0 {
			newTask = &Task{
				Objective: objective,
				Type:      Move,
				Target:    &(path[len(path)-1]),
			}
		} else {
			fmt.Printf("No drinking path found for %v %v\n", character.Name, closestWater)
		}
	}
	return newTask
}

func (sim *Simulation) PlanSleepingTasks(character *Character, objective *Objective) (task *Task) {
	newTask := &Task{
		Objective: objective,
		Type:      Sleep,
	}
	return newTask
}
