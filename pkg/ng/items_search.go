package ng

import (
	"eramstein/thurigen/pkg/config"
)

// ScanForItem searches the closest item of a given type by looping tiles around a position
// if variant is irrelevant pass -1
func (sim *Simulation) ScanForItem(position Position, maxDistance int, itemType ItemType, variant int, unclaimedOnly bool) *Item {
	// Check current tile
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	for _, item := range tile.Items {
		if item.Type == itemType && (item.Variant == variant || variant == -1) && (!unclaimedOnly || item.OwnedBy == 0) {
			return item
		}
	}

	// Check tiles further and further away until maxDistance
	for distance := 1; distance <= maxDistance; distance++ {
		// top row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y - distance
			item := sim.FindItemInTile(position.Region, x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
		// bottom row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y + distance
			item := sim.FindItemInTile(position.Region, x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
		// left row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X - distance
			y := position.Y + dy
			item := sim.FindItemInTile(position.Region, x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
		// right row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X + distance
			y := position.Y + dy
			item := sim.FindItemInTile(position.Region, x, y, itemType, variant, unclaimedOnly)
			if item != nil {
				return item
			}
		}
	}
	return nil
}

func (sim *Simulation) FindItemInTile(region int, x int, y int, itemType ItemType, variant int, unclaimedOnly bool) *Item {
	// Skip if out of bounds
	if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
		return nil
	}

	// Check items on this tile
	tile := &sim.World[region].Tiles[x][y]
	for _, item := range tile.Items {
		if item.Type == itemType && (!unclaimedOnly || item.OwnedBy == 0) {
			return item
		}
	}
	return nil
}

func (character *Character) FindInInventory(itemType ItemType, variant int) *Item {
	for _, item := range character.Inventory {
		if item.Type == itemType && (item.Variant == variant || variant == -1) {
			return item
		}
	}
	return nil
}
