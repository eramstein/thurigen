package ng

import (
	"fmt"
)

func PrintCharacterDetails(character *Character) {
	fmt.Printf("=== Character Details ===\n")
	fmt.Printf("Basic Info:\n")
	fmt.Printf("  Name: %s\n", character.Name)
	fmt.Printf("  ID: %d\n", character.ID)
	fmt.Printf("  Position: Region %d (X: %d, Y: %d)\n",
		character.Position.Region, character.Position.X, character.Position.Y)
	fmt.Printf("  Movement Speed: %.1f\n", character.Stats.Speed)

	fmt.Printf("\nNeeds:\n")
	fmt.Printf("  Food: %d%%\n", character.Needs.Food)
	fmt.Printf("  Water: %d%%\n", character.Needs.Water)
	fmt.Printf("  Sleep: %d%%\n", character.Needs.Sleep)

	if character.Path != nil && len(*character.Path) > 0 {
		fmt.Printf("\nCurrent Path:\n")
		for i, pos := range *character.Path {
			fmt.Printf("  %d. Region %d (X: %d, Y: %d)\n",
				i+1, pos.Region, pos.X, pos.Y)
		}
	}

	if len(character.Tasks) > 0 {
		fmt.Printf("\nCurrent Tasks:\n")
		for i, task := range character.Tasks {
			fmt.Printf("  %d. %v (Progress: %.0f%%)\n",
				i+1, task.Type, task.Progress)

			// Print task target details based on type
			switch task.Type {
			case Move:
				if pos, ok := task.Target.(*Position); ok && pos != nil {
					fmt.Printf("     Target Location: Region %d (X: %d, Y: %d)\n",
						pos.Region, pos.X, pos.Y)
				}
			case Eat:
				if item, ok := task.Target.(*Item); ok && item != nil {
					fmt.Printf("     Target Item: %v\n", item.Type)
					if item.OnTile != nil {
						fmt.Printf("     Item Location: Region %d (X: %d, Y: %d)\n",
							item.OnTile.Region, item.OnTile.X, item.OnTile.Y)
					}
				}
			case Drink:
				if pos, ok := task.Target.(*Position); ok && pos != nil {
					fmt.Printf("     Water Source: Region %d (X: %d, Y: %d)\n",
						pos.Region, pos.X, pos.Y)
				}
			}

			if task.Objective != nil {
				fmt.Printf("     For Objective: %v\n", task.Objective.Type)
			}
		}
	}

	if len(character.Objectives) > 0 {
		fmt.Printf("\nObjectives:\n")
		for i, objective := range character.Objectives {
			fmt.Printf("  %d. %v\n", i+1, objective.Type)
		}
	}

	if len(character.Inventory) > 0 {
		fmt.Printf("\nInventory:\n")
		itemCounts := make(map[string]int)
		for _, item := range character.Inventory {
			itemName := item.Type.String()
			itemCounts[itemName]++
		}
		for itemName, count := range itemCounts {
			fmt.Printf("  - %s: %d\n", itemName, count)
		}
	}
	fmt.Printf("=====================\n")
}
