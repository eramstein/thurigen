package ng

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

func (character *Character) PlanEatingTasks(objective *Objective) {
	// Check if the character has the item in their inventory
	item := character.FindInInventory(Food)

	// If the character has the item in their inventory, add a task to eat it
	if item != nil {
		if foodItem, ok := (*item).(*FoodItem); ok {
			character.AddTask(&EatTask{
				BaseTask: BaseTask{
					Objective: objective,
					Type:      Eat,
				},
				Target: foodItem,
			})
		}
	} else {
		character.AddTask(&FindTask{
			BaseTask: BaseTask{
				Objective: objective,
				Type:      Find,
			},
			Target: Food,
		})
	}
}

func (character *Character) PlanDrinkingTasks(objective *Objective) {
}

func (character *Character) PlanSleepingTasks(objective *Objective) {
}
