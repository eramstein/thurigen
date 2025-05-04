package ng

import (
	"fmt"
)

func printTaskDetails(task *Task, indent string) {
	fmt.Printf("%sType: %v (Progress: %.0f%%)\n",
		indent, task.Type, task.Progress)

	// Print task target details based on type
	switch task.Type {
	case Move:
		if pos, ok := task.Target.(*Position); ok && pos != nil {
			fmt.Printf("%sTarget Location: Region %d (X: %d, Y: %d)\n",
				indent, pos.Region, pos.X, pos.Y)
		}
	case Eat:
		if item, ok := task.Target.(*Item); ok && item != nil {
			fmt.Printf("%sTarget Item: %v\n", indent, item.Type)
			if item.OnTile != nil {
				fmt.Printf("%sItem Location: Region %d (X: %d, Y: %d)\n",
					indent, item.OnTile.Region, item.OnTile.X, item.OnTile.Y)
			}
		}
	case Drink:
		if pos, ok := task.Target.(*Position); ok && pos != nil {
			fmt.Printf("%sWater Source: Region %d (X: %d, Y: %d)\n",
				indent, pos.Region, pos.X, pos.Y)
		}
	}

	if task.Objective != nil {
		fmt.Printf("%sFor Objective: %v\n", indent, task.Objective.Type)
	}
}

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

	if character.CurrentTask != nil {
		fmt.Printf("\nCurrent Task:\n")
		printTaskDetails(character.CurrentTask, "  ")
	}

	if character.Path != nil && len(*character.Path) > 0 {
		fmt.Printf("\nCurrent Path:\n")
		for i, pos := range *character.Path {
			fmt.Printf("  %d. Region %d (X: %d, Y: %d)\n",
				i+1, pos.Region, pos.X, pos.Y)
		}
	}

	if character.CurrentTask != nil {
		fmt.Printf("\nCurrent Task:\n")
		printTaskDetails(character.CurrentTask, "  ")
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
