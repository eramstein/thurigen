package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

func (sim *Simulation) SetCurrentTask(character *Character) {
	topObjective := sim.GetTopPriorityObjective(character)
	if topObjective != nil {
		nextTask := sim.CreateNextTask(character, topObjective)
		if nextTask != nil {
			character.CurrentTask = nextTask
		} else {
			topObjective.Stuck = true
			fmt.Printf("Objective stuck because no task: %v\n", topObjective)
		}
	}
}

func (sim *Simulation) WorkOnCurrentTask(character *Character) {
	task := character.CurrentTask
	if task == nil {
		return
	}
	switch task.Type {
	case Move:
		sim.MoveForTask(character)
	case Eat:
		sim.Eat(character)
	case Drink:
		sim.Drink(character)
	case Sleep:
		sim.Sleep(character)
	case Build:
		sim.Build(character)
	case PickUp:
		sim.PickUp(character)
	case Chop:
		sim.Chop(character)
	}
	if task.Progress >= 100 {
		sim.CompleteTask(character)
	}
}

// Create next task for a given objective
// Add it to the character's tasks if it's not nil
func (sim *Simulation) CreateNextTask(character *Character, objective *Objective) (task *Task) {
	switch objective.Type {
	case EatObjective:
		task = sim.GetNextEatingTask(character, objective)
	case DrinkObjective:
		task = sim.GetNextDrinkingTask(character, objective)
	case SleepObjective:
		task = sim.GetNextSleepingTask(character, objective)
	case BuildObjective:
		task = sim.GetNextBuildingTask(character, objective)
	}
	return task
}

func (sim *Simulation) CompleteTask(character *Character) {
	if character.CurrentTask == nil {
		return
	}
	fmt.Printf("Completing task:  %v %v %v\n", character.Name, character.CurrentTask.Type, character.CurrentTask.Objective.Type)
	sim.CheckIfObjectiveIsAchieved(character, character.CurrentTask.Objective)
	character.CurrentTask = nil
}

func (sim *Simulation) CancelTask(character *Character) {
	if character.CurrentTask == nil {
		return
	}
	fmt.Printf("Cancelling task:  %v %v %v\n", character.Name, character.CurrentTask.Type, character.CurrentTask.Objective)
	character.CurrentTask = nil
}

// Set next task required to achieve an eat objective
func (sim *Simulation) GetNextEatingTask(character *Character, objective *Objective) (task *Task) {
	var newTask *Task
	// Check if the character has the item in their inventory
	itemInInventory := character.FindInInventory(Food, -1)
	// If the character has the item in their inventory, add a task to eat it
	if itemInInventory != nil {
		newTask = &Task{
			Objective: objective,
			Type:      Eat,
			Target:    itemInInventory,
		}
		// If the character is on a tile with a food item, add a task to eat it
	} else if itemOnTile := sim.FindItemInTile(character.Position.Region, character.Position.X, character.Position.Y, Food, -1, false); itemOnTile != nil {
		// claim item
		itemOnTile.OwnedBy = character.ID
		// eat it
		newTask = &Task{
			Objective: objective,
			Type:      Eat,
			Target:    itemOnTile,
		}
	} else {
		// If no food on tile, find the closest food item and add a task to go to it
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2-1, Food, -1, true)
		if closestItem != nil {
			// claim item
			closestItem.OwnedBy = character.ID
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

func (sim *Simulation) GetNextDrinkingTask(character *Character, objective *Objective) (task *Task) {
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
		// problem: if closestWater is not accessible, there will be no path found and no task added
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

func (sim *Simulation) GetNextSleepingTask(character *Character, objective *Objective) (task *Task) {
	newTask := &Task{
		Objective: objective,
		Type:      Sleep,
	}
	return newTask
}
