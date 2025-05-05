package ng

func IsAdjacent(x1, y1, x2, y2 int) bool {
	return absDiff(x1, x2) <= 1 && absDiff(y1, y2) <= 1
}

func (sim *Simulation) IsOccupied(position Position) bool {
	tile := sim.World[position.Region].Tiles[position.X][position.Y]

	if tile.Occupation != nil || tile.Volume != NoVolume {
		return true
	}

	return false
}

func (sim *Simulation) GetTile(position Position) *Tile {
	return &sim.World[position.Region].Tiles[position.X][position.Y]
}
