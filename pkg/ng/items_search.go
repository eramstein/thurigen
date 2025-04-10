package ng

func (sim *Simulation) FindCloseItem(region int, position [2]int, itemType ItemType) *Item {
	// Check current tile
	tile := &sim.World[region].Tiles[position[0]][position[1]]
	for _, item := range tile.Items {
		if (*item).GetItem().Type == itemType {
			return item
		}
	}
	// Check adjacent tiles (including diagonal)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// Skip the current tile
			if dx == 0 && dy == 0 {
				continue
			}
			x := position[0] + dx
			y := position[1] + dy

			// Skip if out of bounds
			if x < 0 || x >= len(sim.World[region].Tiles) || y < 0 || y >= len(sim.World[region].Tiles[0]) {
				continue
			}

			// Check items on this tile
			tile := &sim.World[region].Tiles[x][y]
			for _, item := range tile.Items {
				if (*item).GetItem().Type == itemType {
					return item
				}
			}
		}
	}
	return nil
}

func (character *Character) FindInInventory(itemType ItemType) *Item {
	for _, item := range character.Inventory {
		if (*item).GetItem().Type == itemType {
			return item
		}
	}
	return nil
}
