package ng

func (b *BaseItem) GetItem() *BaseItem {
	return b
}

func (b *FoodItem) GetItem() *BaseItem {
	return &b.BaseItem
}

func (b *MaterialItem) GetItem() *BaseItem {
	return &b.BaseItem
}

func (sim *Simulation) SpawnItem(item *Item, position Position) {
	// Set item's position
	(*item).GetItem().OnTile = &position
	// Add to simulation's items
	sim.Items = append(sim.Items, item)
	// Add to tile's items
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	tile.Items = append(tile.Items, item)
}

func (sim *Simulation) DeleteItem(item *Item) {
	base := (*item).GetItem()
	// remove reference of item in inventories
	if base.InInventoryOf != nil {
		for i, inventoryRef := range base.InInventoryOf.Inventory {
			if inventoryRef == item {
				base.InInventoryOf.Inventory = append(base.InInventoryOf.Inventory[:i], base.InInventoryOf.Inventory[i+1:]...)
			}
		}
	}
	// remove reference of item on tiles
	if base.OnTile != nil {
		for i, tileItemRef := range sim.World[base.OnTile.Region].Tiles[base.OnTile.X][base.OnTile.Y].Items {
			if tileItemRef == item {
				sim.World[base.OnTile.Region].Tiles[base.OnTile.X][base.OnTile.Y].Items = append(sim.World[base.OnTile.Region].Tiles[base.OnTile.X][base.OnTile.Y].Items[:i], sim.World[base.OnTile.Region].Tiles[base.OnTile.X][base.OnTile.Y].Items[i+1:]...)
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
	switch itemType {
	case Food:
		return &FoodItem{
			BaseItem: BaseItem{
				Type:    itemType,
				Variant: variant,
			},
			Nutrition: 50, // Default nutrition value
		}
	default:
		return &BaseItem{
			Type:    itemType,
			Variant: variant,
		}
	}
}
