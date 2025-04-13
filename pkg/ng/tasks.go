package ng

import (
	"eramstein/thurigen/pkg/config"
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
		// TODO: find the closest food item and add a task to go to it
		closestItem := sim.ScanForItem(character.Position, config.RegionSize/2, Food)
		if closestItem != nil {
			character.AddTask(&MoveTask{
				BaseTask: BaseTask{
					Objective: objective,
					Type:      Move,
				},
				Target: (*closestItem).GetItem().OnTile,
			})
		}
	}
}

func (sim *Simulation) PlanDrinkingTasks(character *Character, objective *Objective) {
}

func (sim *Simulation) PlanSleepingTasks(character *Character, objective *Objective) {
}
