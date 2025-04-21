package ng

func (sim *Simulation) SpawnItem(item *Item, position Position) {
	// Set item's position
	item.OnTile = &position
	// Add to simulation's items
	sim.Items = append(sim.Items, item)
	// Add to tile's items
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	tile.Items = append(tile.Items, item)
}

func (sim *Simulation) DeleteItem(item *Item) {
	// remove reference of item in inventories
	if item.InInventoryOf != nil {
		for i, inventoryRef := range item.InInventoryOf.Inventory {
			if inventoryRef == item {
				item.InInventoryOf.Inventory = append(item.InInventoryOf.Inventory[:i], item.InInventoryOf.Inventory[i+1:]...)
			}
		}
	}
	// remove reference of item on tiles
	if item.OnTile != nil {
		for i, tileItemRef := range sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items {
			if tileItemRef == item {
				sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items = append(sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items[:i], sim.World[item.OnTile.Region].Tiles[item.OnTile.X][item.OnTile.Y].Items[i+1:]...)
			}
		}
	}
	// remove actual item in simulation
	for i, simItem := range sim.Items {
		if simItem == item {
			sim.Items = append(sim.Items[:i], sim.Items[i+1:]...)
		}
	}
}

func MakeItem(itemType ItemType, variant int) Item {
	item := Item{
		Type:    itemType,
		Variant: variant,
	}
	switch itemType {
	case Food:
		item.Efficiency = 50
	}
	return item
}
