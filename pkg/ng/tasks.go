package ng

import (
	"eramstein/thurigen/pkg/config"
)

func (sim *Simulation) UpdatePriorityTask(character *Character) {
	if len(character.Tasks) == 0 {
		return
	}
	// tasks are sorted by priority, work on the highest priority task
	task := character.Tasks[0]
	switch task.Type {
	case Move:
		character.Move()
	}
}

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
	} else {
		// Find the closest food item and add a task to go to it
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
