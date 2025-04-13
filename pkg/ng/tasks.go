package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

func (b *BaseTask) GetTask() *BaseTask {
	return b
}

func (b *EatTask) GetTask() *BaseTask {
	return &b.BaseTask
}

func (b *DrinkTask) GetTask() *BaseTask {
	return &b.BaseTask
}

func (b *SleepTask) GetTask() *BaseTask {
	return &b.BaseTask
}

func (b *FindTask) GetTask() *BaseTask {
	return &b.BaseTask
}

func (b *MoveTask) GetTask() *BaseTask {
	return &b.BaseTask
}

func (sim *Simulation) PlanEatingTasks(character *Character, objective *Objective) {
	// Check if the character has the item in their inventory
	itemInInventory := character.FindInInventory(Food)

	// If the character has the item in their inventory, add a task to eat it
	if itemInInventory != nil {
		character.AddTask(&EatTask{
			BaseTask: BaseTask{
				Objective: objective,
				Type:      Eat,
			},
			Target: (*itemInInventory).(*FoodItem),
		})
	} else {
		// Find the closest food item and add a task to go to it
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2, Food)
		if closestItem != nil {
			base := (*closestItem).GetItem()
			character.AddTask(&MoveTask{
				BaseTask: BaseTask{
					Objective: objective,
					Type:      Move,
				},
				Target: base.OnTile,
			})
			path := sim.World[character.Position.Region].FindPath(character.Position.X, character.Position.Y, base.OnTile.X, base.OnTile.Y)
			fmt.Println("closestItem", base.OnTile)
			if path != nil {
				for _, pos := range path {
					fmt.Println("pos", pos)
				}
			} else {
				fmt.Println("no path")
			}
		}
	}
}

func (sim *Simulation) PlanDrinkingTasks(character *Character, objective *Objective) {
}

func (sim *Simulation) PlanSleepingTasks(character *Character, objective *Objective) {
}
