package ng

func (sim *Simulation) PlanHouseBuildingTasks(character *Character, objective *Objective) {
	// Define house dimensions (5x5 tiles)
	houseWidth := 5
	houseHeight := 5

	// Calculate starting position (top-left corner)
	startX := character.Position.X
	startY := character.Position.Y

	doorX := startX + houseWidth/2
	doorY := startY + houseHeight - 1

	// Add walls for the rectangle
	// Top wall
	for x := startX; x < startX+houseWidth; x++ {
		objective.Plan = append(objective.Plan, Task{
			Type:           Build,
			ProductType:    int(Wall),
			ProductVariant: int(WoodMaterial),
			Target: &Position{
				Region: character.Position.Region,
				X:      x,
				Y:      startY,
			},
		})
	}

	// Bottom wall
	for x := startX; x < startX+houseWidth; x++ {
		y := startY + houseHeight - 1
		if x == doorX && y == doorY {
			continue
		}
		objective.Plan = append(objective.Plan, Task{
			Type:           Build,
			ProductType:    int(Wall),
			ProductVariant: int(WoodMaterial),
			Target: &Position{
				Region: character.Position.Region,
				X:      x,
				Y:      y,
			},
		})
	}

	// Left wall (excluding corners)
	for y := startY + 1; y < startY+houseHeight-1; y++ {
		objective.Plan = append(objective.Plan, Task{
			Type:           Build,
			ProductType:    int(Wall),
			ProductVariant: int(WoodMaterial),
			Target: &Position{
				Region: character.Position.Region,
				X:      startX,
				Y:      y,
			},
		})
	}

	// Right wall (excluding corners)
	for y := startY + 1; y < startY+houseHeight-1; y++ {
		objective.Plan = append(objective.Plan, Task{
			Type:           Build,
			ProductType:    int(Wall),
			ProductVariant: int(WoodMaterial),
			Target: &Position{
				Region: character.Position.Region,
				X:      startX + houseWidth - 1,
				Y:      y,
			},
		})
	}

	// Add door in the middle of the bottom wall
	objective.Plan = append(objective.Plan, Task{
		Type:           Build,
		ProductType:    int(Door),
		ProductVariant: int(WoodMaterial),
		Target: &Position{
			Region: character.Position.Region,
			X:      doorX,
			Y:      doorY,
		},
	})
}
