package ng

// ScanForItem searches the closest item of a given type by looping tiles around a position
func (sim *Simulation) ScanForItem(position Position, maxDistance int, itemType ItemType) *Item {
	// Check current tile
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	for _, item := range tile.Items {
		if (*item).GetItem().Type == itemType {
			return item
		}
	}

	// Check tiles further and further away until maxDistance
	for distance := 1; distance <= maxDistance; distance++ {
		// top row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y - distance
			item := sim.FindItemInTile(position.Region, x, y, itemType)
			if item != nil {
				return item
			}
		}
		// bottom row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y + distance
			item := sim.FindItemInTile(position.Region, x, y, itemType)
			if item != nil {
				return item
			}
		}
		// left row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X - distance
			y := position.Y + dy
			item := sim.FindItemInTile(position.Region, x, y, itemType)
			if item != nil {
				return item
			}
		}
		// right row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X + distance
			y := position.Y + dy
			item := sim.FindItemInTile(position.Region, x, y, itemType)
			if item != nil {
				return item
			}
		}
	}
	return nil
}

// GetClosestItem searches the closest item of a given type by looping items in a region
func (sim *Simulation) GetClosestItem(position Position, itemType ItemType) *Item {
	var closestItem *Item
	minDistance := -1

	// Loop through all items in the simulation
	for _, item := range sim.Items {
		baseItem := (*item).GetItem()

		// Check if this item is of the right type and in the right region
		if baseItem.Type == itemType && baseItem.OnTile.Region == position.Region {
			// Calculate distance to the target position
			dx := baseItem.OnTile.X - position.X
			dy := baseItem.OnTile.Y - position.Y
			distance := dx*dx + dy*dy // Using squared distance to avoid square root calculation

			// Update closest item if this is closer
			if minDistance == -1 || distance < minDistance {
				minDistance = distance
				closestItem = item
			}
		}
	}
	return closestItem
}

// TODO: check if this is faster than GetClosestItem or ScanForItem
func (sim *Simulation) ScanForItemFullRegion(position Position, itemType ItemType) *Item {
	var closestItem *Item
	minDistance := -1

	// Iterate through all tiles in the region
	for x := 0; x < len(sim.World[position.Region].Tiles); x++ {
		for y := 0; y < len(sim.World[position.Region].Tiles[0]); y++ {
			tile := &sim.World[position.Region].Tiles[x][y]

			// Check items on this tile
			for _, item := range tile.Items {
				if (*item).GetItem().Type == itemType {
					// Calculate distance to the target position
					dx := x - position.X
					dy := y - position.Y
					distance := dx*dx + dy*dy // Using squared distance to avoid square root calculation

					if distance <= 1 {
						return item
					}
					// Update closest item if this is closer
					if minDistance == -1 || distance < minDistance {
						minDistance = distance
						closestItem = item
					}
				}
			}
		}
	}

	return closestItem
}

func (sim *Simulation) FindItemInTile(region int, x int, y int, itemType ItemType) *Item {
	// Skip if out of bounds
	if x < 0 || x >= len(sim.World[region].Tiles) || y < 0 || y >= len(sim.World[region].Tiles[0]) {
		return nil
	}

	// Check items on this tile
	tile := &sim.World[region].Tiles[x][y]
	for _, item := range tile.Items {
		if (*item).GetItem().Type == itemType {
			return item
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
