package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

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
	requiredItemType, requiredItemVariant := getRequiredItem(nextBuildTask)

	itemInInventory := character.FindInInventory(requiredItemType, requiredItemVariant)
	if itemInInventory != nil {
		newTask = buildFromInventory(character, objective, nextBuildTask, buildTile, itemInInventory)
	} else if itemOnTile := sim.FindItemInTile(character.Position.Region, character.Position.X, character.Position.Y, requiredItemType, requiredItemVariant, false); itemOnTile != nil {
		newTask = pickupMaterial(character, objective, itemOnTile)
	} else {
		newTask = sim.goGetMaterial(character, objective, requiredItemType, requiredItemVariant)
	}

	return newTask
}

func getRequiredItem(nextBuildTask Task) (ItemType, int) {
	var requiredItemType ItemType
	var requiredItemVariant int
	switch nextBuildTask.ProductType {
	case int(Wall):
		requiredItemType = Material
		requiredItemVariant = nextBuildTask.ProductVariant
	}
	return requiredItemType, requiredItemVariant
}

func buildFromInventory(character *Character, objective *Objective, nextBuildTask Task, buildTile *Position, itemInInventory *Item) (task *Task) {
	// if the character is already adjacent to the build site, build and remove task from plan
	if IsAdjacent(character.Position.X, character.Position.Y, buildTile.X, buildTile.Y) {
		// complete plan task: link to objective and used material source
		nextBuildTask.Objective = objective
		nextBuildTask.MaterialSource = itemInInventory
		objective.Plan = objective.Plan[1:]
		return &nextBuildTask
	} else {
		// if the character is not adjacent to the build site, move to it
		return &Task{
			Objective: objective,
			Type:      Move,
			Target:    buildTile,
		}
	}
}

func pickupMaterial(character *Character, objective *Objective, itemOnTile *Item) (task *Task) {
	// claim item
	itemOnTile.OwnedBy = character.ID
	// pick it up
	return &Task{
		Objective: objective,
		Type:      PickUp,
		Target:    itemOnTile,
	}
}

func (sim *Simulation) goGetMaterial(character *Character, objective *Objective, requiredItemType ItemType, requiredItemVariant int) (task *Task) {
	// If no material on tile, find the closest material item and add a task to go to it
	closestItem := sim.ScanForItem(character.Position, config.RegionSize/2-1, requiredItemType, requiredItemVariant, true)
	if closestItem != nil {
		// claim item
		closestItem.OwnedBy = character.ID
		// go to it
		return &Task{
			Objective: objective,
			Type:      Move,
			Target:    closestItem.OnTile,
		}
	} else {
		// go chop a tree (but not an apple tree)
		closestTree := sim.FindClosestPlant(character.Position, Plant, 1, Tree)
		if closestTree == nil {
			fmt.Printf("No tree found for %v\n", character.Name)
			return nil
		}
		if IsAdjacent(character.Position.X, character.Position.Y, closestTree.Position.X, closestTree.Position.Y) {
			return &Task{
				Objective: objective,
				Type:      Chop,
				Target:    closestTree,
			}
		} else {
			return &Task{
				Objective: objective,
				Type:      Move,
				Target:    &closestTree.Position,
			}
		}
	}
}
