package ng

import "fmt"

func (sim *Simulation) SpawnItem(item Item, position Position) {
	item.ID = getNextID()
	// Set item's position
	item.OnTile = &position
	// Add to simulation's items
	sim.Items = append(sim.Items, item)
	// Add to tile's items
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	tile.Items = append(tile.Items, &item)
}

func (sim *Simulation) DeleteItem(item *Item) {
	// remove reference of item in inventories
	if item.InInventoryOf != 0 {
		character := sim.GetCharacter(item.InInventoryOf)
		for i, inventoryRef := range character.Inventory {
			if inventoryRef == item {
				character.Inventory = append(character.Inventory[:i], character.Inventory[i+1:]...)
			}
		}
	}
	// remove reference of item on tiles
	if item.OnTile != nil {
		sim.RemoveItemFromTile(item)
	}
	// remove actual item in simulation
	for i, simItem := range sim.Items {
		if simItem.ID == item.ID {
			sim.Items = append(sim.Items[:i], sim.Items[i+1:]...)
		}
	}
}

func (sim *Simulation) ReduceItemDurability(item *Item, amount int) {
	item.Durability -= amount
	if item.Durability <= 0 {
		sim.DeleteItem(item)
	}
}

func (sim *Simulation) RemoveItemFromTile(item *Item) {
	if item.OnTile == nil {
		fmt.Printf("WARNING: Item %v to REMOVE FROM TILE is not on a tile, drop task\n", item)
		return
	}
	for i, tileItem := range sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items {
		if tileItem == item {
			sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items = append(sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items[:i], sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items[i+1:]...)
		}
	}
	item.OnTile = nil
}
