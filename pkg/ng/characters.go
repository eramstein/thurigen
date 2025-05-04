package ng

import (
	"eramstein/thurigen/pkg/config"
	"fmt"
)

func (sim *Simulation) GetCharacter(id uint64) *Character {
	for _, character := range sim.Characters {
		if character.ID == id {
			return character
		}
	}
	return nil
}

func (sim *Simulation) InitCharacters() {
	sim.MakeCharacter(1, "Henry", Position{Region: 0, X: 30, Y: 30}, CharacterStats{Speed: 1.8})
	//sim.MakeCharacter(2, "Ella", Position{Region: 0, X: 35, Y: 35}, CharacterStats{Speed: 1.45})
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
			if character.CurrentTask == nil {
				sim.SetCurrentTask(character)
			}
			sim.WorkOnCurrentTask(character)
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
		Needs:     Needs{Food: 0, Water: 0, Sleep: 49},
	}
	sim.Characters = append(sim.Characters, character)
	sim.World[pos.Region].Tiles[pos.X][pos.Y].Character = character
	return character
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

func (sim *Simulation) Eat(character *Character) {
	task := character.CurrentTask
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

func (sim *Simulation) Drink(character *Character) {
	task := character.CurrentTask
	position := task.Target.(*Position)
	tile := sim.World[position.Region].Tiles[position.X][position.Y]
	if tile.Terrain != Water {
		return
	}
	task.Progress += 50
	fmt.Println("Drinking", character.Name)
	if task.Progress >= 100 {
		character.Needs.Water = 0
		task.Progress = 100
	}
}

func (sim *Simulation) Sleep(character *Character) {
	task := character.CurrentTask
	fmt.Println("Sleeping", character.Name)
	character.Needs.Sleep -= config.NeedSleepTick * 5
	if character.Needs.Sleep <= 0 {
		character.Needs.Sleep = 0
		task.Progress = 100
		character.UpdateSleepConditionsWants(WantSleepOnFloor)
	}
}

func (sim *Simulation) Build(character *Character) {
	task := character.CurrentTask
	materialSource := task.MaterialSource
	if materialSource == nil || materialSource.Durability <= 0 {
		fmt.Printf("WARNING: Material source %v to BUILD is inexistant or has no durability\n", materialSource)
		sim.CancelTask(character)
		return
	}
	task.Progress += 10
	sim.ReduceItemDurability(materialSource, 1)
	if task.Progress >= 100 {
		switch task.ProductType {
		case int(Wall):
			targetTile := task.Target.(*Position)
			sim.AddWall(targetTile.Region, targetTile.X, targetTile.Y, MaterialType(task.ProductVariant), 0)
		}
	}
}

func (sim *Simulation) PickUp(character *Character) {
	task := character.CurrentTask
	item := task.Target.(*Item)
	if item.OnTile == nil || !IsAdjacent(character.Position.X, character.Position.Y, item.OnTile.X, item.OnTile.Y) {
		fmt.Printf("WARNING: Item %v to PICKUP is not on a tile or not adjacent\n", item)
		sim.CancelTask(character)
		return
	}
	fmt.Printf("Picking up %v\n", item)
	character.Inventory = append(character.Inventory, item)
	item.InInventoryOf = character.ID
	sim.RemoveItemFromTile(item)
	task.Progress = 100
}

func (sim *Simulation) Chop(character *Character) {
	task := character.CurrentTask
	tree := task.Target.(*PlantStructure)
	if tree == nil {
		fmt.Printf("WARNING: Tree %v to CHOP is inexistant\n", tree)
		sim.CancelTask(character)
		return
	}
	task.Progress += 20
	if task.Progress >= 100 {
		sim.ChopTree(tree)
	}
}
