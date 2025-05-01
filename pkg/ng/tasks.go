package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

// Create next task for a given objective
// Add it to the character's tasks if it's not nil
func (sim *Simulation) CreateNextTask(character *Character, objective *Objective) (task *Task) {
	fmt.Printf("Planning tasks for %v %v\n", character.Name, objective.Type)
	// Remove previous tasks related to same objective type
	if UniqueObjectiveTypes[objective.Type] {
		for i, task := range character.Tasks {
			if task.Objective.Type == objective.Type {
				character.Tasks = append(character.Tasks[:i], character.Tasks[i+1:]...)
			}
		}
	}
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
		newTask := sim.CreateNextTask(character, topObjective)
		if newTask != nil {
			return newTask
		}
	}
	return nil
}

func (sim *Simulation) SetPriorityTask(character *Character) {
	topObjective := sim.GetTopPriorityObjective(character)
	if topObjective != nil {
		sim.CreateNextTask(character, topObjective)
	}
	character.CurrentTask = sim.GetPriorityTask(character)
}

func (sim *Simulation) WorkOnPriorityTask(character *Character) {
	task := character.CurrentTask
	if task == nil {
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
	case Build:
		sim.Build(character, task)
	case PickUp:
		sim.PickUp(character, task)
	case Chop:
		sim.Chop(character, task)
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

func (sim *Simulation) GetNextSleepingTask(character *Character, objective *Objective) (task *Task) {
	newTask := &Task{
		Objective: objective,
		Type:      Sleep,
	}
	return newTask
}

// TODO: split this up it became too complex
func (sim *Simulation) GetNextBuildingTask(character *Character, objective *Objective) (task *Task) {
	if len(objective.Plan) == 0 {
		return nil
	}
	var newTask *Task
	// Building objectives have pre-planned set of build tasks
	// Check if conditions to work on first one are met
	nextBuildTask := objective.Plan[0]
	buildTile, ok := nextBuildTask.Target.(*Position)
	if !ok || buildTile == nil {
		fmt.Printf("ERR: Build task target is not a valid Position: %v\n", nextBuildTask.Target)
		return nil
	}

	if nextBuildTask.Type != Build {
		fmt.Printf("ERR: Next build task is not a build task: %v\n", nextBuildTask)
		return nil
	}

	var requiredItemType ItemType
	var requiredItemVariant int

	switch nextBuildTask.ProductType {
	case int(Wall):
		requiredItemType = Material
		requiredItemVariant = nextBuildTask.ProductVariant
	}

	// Check if the character has the item in their inventory
	itemInInventory := character.FindInInventory(requiredItemType, requiredItemVariant)
	// If the character has the item in their inventory
	if itemInInventory != nil {
		// if the character is already adjacent to the build site, build and remove task from plan
		if IsAdjacent(character.Position.X, character.Position.Y, buildTile.X, buildTile.Y) {
			// restore link to objective which was not in the Plan
			nextBuildTask.Objective = objective
			nextBuildTask.MaterialSource = itemInInventory
			newTask = &nextBuildTask
			objective.Plan = objective.Plan[1:]
		} else {
			// if the character is not adjacent to the build site, move to it
			newTask = &Task{
				Objective: objective,
				Type:      Move,
				Target:    buildTile,
			}
		}
	} else if itemOnTile := sim.FindItemInTile(character.Position.Region, character.Position.X, character.Position.Y, requiredItemType, requiredItemVariant, false); itemOnTile != nil {
		// claim item
		itemOnTile.OwnedBy = character.ID
		// pick it up
		newTask = &Task{
			Objective: objective,
			Type:      PickUp,
			Target:    itemOnTile,
		}
	} else {
		// If no material on tile, find the closest material item and add a task to go to it
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2-1, requiredItemType, requiredItemVariant, true)
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
			fmt.Printf("No material found for %v\n", character.Name)
			// go chop a tree (but not an apple tree)
			closestTree := sim.FindClosestPlant(character.Position, Plant, 1, Tree)
			if IsAdjacent(character.Position.X, character.Position.Y, closestTree.Position.X, closestTree.Position.Y) {
				newTask = &Task{
					Objective: objective,
					Type:      Chop,
					Target:    closestTree,
				}
			} else {
				newTask = &Task{
					Objective: objective,
					Type:      Move,
					Target:    &closestTree.Position,
				}
			}
		}
	}

	return newTask
}
