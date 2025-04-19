package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter("Henry", Position{Region: 0, X: 30, Y: 30}, CharacterStats{Speed: 1.5})
}

func (sim *Simulation) UpdateCharacters() {
	if sim.Time%config.CharacterNeedsUpdateInterval == 0 {
		for _, character := range sim.Characters {
			character.UpdateNeeds()
		}
	}
	if sim.Time%config.CharacterObjectiveUpdateInterval == 0 {
		for _, character := range sim.Characters {
			sim.UpdateObjectives(character)
		}
	}
	if sim.Time%config.CharacterTaskUpdateInterval == 0 {
		for _, character := range sim.Characters {
			sim.WorkOnPriorityTask(character)
		}
	}
}

func (character *Character) UpdateNeeds() {
	character.Needs.Food += 1
	character.Needs.Water += 1
	character.Needs.Sleep += 1
}

func (sim *Simulation) UpdateObjectives(character *Character) {
	if character.Needs.Food >= 50 && !character.HasObjective(EatObjective) {
		sim.AddObjective(character, EatObjective)
	}

	if character.Needs.Water >= 50 && !character.HasObjective(DrinkObjective) {
		sim.AddObjective(character, DrinkObjective)
	}

	if character.Needs.Sleep >= 50 && !character.HasObjective(SleepObjective) {
		sim.AddObjective(character, SleepObjective)
	}
}

func (sim *Simulation) RemoveCharacter(character *Character) {
	for i, c := range sim.Characters {
		if c == character {
			sim.Characters = append(sim.Characters[:i], sim.Characters[i+1:]...)
		}
	}
}

func (sim *Simulation) MakeCharacter(name string, pos Position, stats CharacterStats) *Character {
	character := &Character{
		ID:        len(sim.Characters),
		Name:      name,
		Position:  pos,
		Stats:     stats,
		Inventory: []*Item{},
		Needs:     Needs{Food: 49, Water: 0, Sleep: 0},
	}
	sim.Characters = append(sim.Characters, character)
	sim.World[pos.Region].Tiles[pos.X][pos.Y].Character = character
	return character
}

func (sim *Simulation) AddObjective(character *Character, objectiveType ObjectiveType) {
	objective := &Objective{
		Type: objectiveType,
	}
	character.Objectives = append(character.Objectives, objective)
	sim.PlanTasks(character, objective)
}

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

func (character *Character) AddTask(task Task) {
	character.Tasks = append(character.Tasks, &task)
}

func (character *Character) HasObjective(objectiveType ObjectiveType) bool {
	for _, objective := range character.Objectives {
		if objective.Type == objectiveType {
			return true
		}
	}
	return false
}

func (character *Character) CompleteObjective(objective *Objective) {
	fmt.Println("Completing objective", objective.Type, objective)
	for i, charObjective := range character.Objectives {
		if charObjective == objective {
			character.Objectives = append(character.Objectives[:i], character.Objectives[i+1:]...)
		}
	}
}

func (sim *Simulation) FollowPath(character *Character, task *Task, extraMove bool) {
	if character.Path == nil {
		return
	}
	path := *character.Path
	if len(path) > 0 {
		moveCost := sim.GetMoveCost(character, path[0])
		if moveCost == -1 {
			return
		}

		// for move tasks, the progress represents how many move points have been spent in the move to next tile in the path
		// an "etra move" is using only left over action points from the same tick
		if !extraMove {
			task.Progress += character.Stats.Speed
		}
		// if enough move points have been spent, set the character to the new position
		if task.Progress >= moveCost {
			sim.SetCharacterPosition(character, path[0])
			if len(path) == 1 {
				character.Path = nil
				sim.CompleteTask(character, task)
				return
			}
			newPath := path[1:]
			character.Path = &newPath
			task.Progress = task.Progress - moveCost
			// get started on next move if excess move points
			if task.Progress >= 1 {
				sim.FollowPath(character, task, true)
			}
		}
	}
}

// How much movement points are needed for a character to move to a tile
// One movement point corresponds to one tick for a character with speed 1 on a default tile
// Returns -1 if the tile is impassable or occupied
func (sim *Simulation) GetMoveCost(character *Character, position Position) float32 {
	if character.Stats.Speed == 0 {
		return -1
	}
	// check if tile is passable or occupied
	targetTile := sim.World[position.Region].Tiles[position.X][position.Y]
	if targetTile.Character != nil || targetTile.MoveCost == ImpassableCost {
		return -1
	}
	// base tile move cost
	moveCost := float32(targetTile.MoveCost)
	// diagonal move cost
	dx := character.Position.X - position.X
	dy := character.Position.Y - position.Y
	if dx != 0 && dy != 0 {
		moveCost *= 1.41421356
	}
	return moveCost
}

func (sim *Simulation) SetCharacterPosition(character *Character, position Position) {
	sim.World[character.Position.Region].Tiles[character.Position.X][character.Position.Y].Character = nil
	sim.World[position.Region].Tiles[position.X][position.Y].Character = character
	character.Position = position
}

func (sim *Simulation) Eat(character *Character, task *Task) {
	// First dereference the *Item to get the Item interface
	itemInterface := *(task.Target.(*Item))
	// Then do the type assertion to *FoodItem
	food, ok := itemInterface.(*FoodItem)
	if !ok {
		fmt.Println("Task target is not a FoodItem", task.Target)
		return
	}
	task.Progress += 10
	if task.Progress >= 100 {
		character.Needs.Food -= food.Nutrition
		if character.Needs.Food < 0 {
			character.Needs.Food = 0
		}
		sim.DeleteItem(task.Target.(*Item))
		sim.CompleteTask(character, task)
	}
}
