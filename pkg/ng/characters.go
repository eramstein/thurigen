package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
	"sync/atomic"
)

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter(1, "Henry", Position{Region: 0, X: 30, Y: 30}, CharacterStats{Speed: 1.5})
	sim.MakeCharacter(2, "Ella", Position{Region: 0, X: 31, Y: 30}, CharacterStats{Speed: 1})
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
	character.Needs.Food += config.NeedFoodTick
	character.Needs.Water += config.NeedWaterTick
	character.Needs.Sleep += config.NeedSleepTick
}

func (sim *Simulation) RemoveCharacter(character *Character) {
	for i, c := range sim.Characters {
		if c == character {
			sim.Characters = append(sim.Characters[:i], sim.Characters[i+1:]...)
		}
	}
}

func (sim *Simulation) MakeCharacter(id uint64, name string, pos Position, stats CharacterStats) *Character {
	character := &Character{
		ID:        id,
		Name:      name,
		Position:  pos,
		Stats:     stats,
		Inventory: []*Item{},
		Needs:     Needs{Food: 49, Water: 49, Sleep: 49},
	}
	sim.Characters = append(sim.Characters, character)
	sim.World[pos.Region].Tiles[pos.X][pos.Y].Character = character
	return character
}

func (character *Character) AddTask(task Task) {
	task.ID = atomic.AddUint64(&nextTaskID, 1)
	character.Tasks = append(character.Tasks, &task)
}

// How much movement points are needed for a character to move to a tile
// One movement point corresponds to one tick for a character with speed 1 on a default tile
// Returns -1 if the tile is impassable
func (sim *Simulation) GetMoveCost(character *Character, position Position) float32 {
	if character.Stats.Speed == 0 {
		return -1
	}
	// check if tile is passable
	targetTile := sim.World[position.Region].Tiles[position.X][position.Y]
	if targetTile.MoveCost == ImpassableCost {
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
	item := task.Target.(*Item)
	task.Progress += 10
	fmt.Println("Eating", character.Name, item.Type)
	if task.Progress >= 100 {
		character.Needs.Food -= item.Efficiency
		if character.Needs.Food < 0 {
			character.Needs.Food = 0
		}
		sim.DeleteItem(item)
	}
}

func (sim *Simulation) Drink(character *Character, task *Task) {
	position := task.Target.(*Position)
	tile := sim.World[position.Region].Tiles[position.X][position.Y]
	if tile.Terrain != Water {
		return
	}
	character.Needs.Water = 0
	task.Progress = 100
}

func (sim *Simulation) Sleep(character *Character, task *Task) {
	fmt.Println("Sleeping", character.Name)
	character.Needs.Sleep -= config.NeedSleepTick * 2
	if character.Needs.Sleep == 0 {
		task.Progress = 100
	}
}
