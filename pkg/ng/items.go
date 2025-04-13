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
	(*item).GetItem().OnTile = position
	// Add to simulation's items
	sim.Items = append(sim.Items, item)
	// Add to tile's items
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	tile.Items = append(tile.Items, item)
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
