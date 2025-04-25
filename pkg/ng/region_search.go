package ng

import "eramstein/thurigen/pkg/config"

func (sim *Simulation) ScanForTile(position Position, maxDistance int, terrain TerrainType) *Position {
	// Check current tile
	tile := &sim.World[position.Region].Tiles[position.X][position.Y]
	if tile.Terrain == terrain {
		return &position
	}

	// Check tiles further and further away until maxDistance
	for distance := 1; distance <= maxDistance; distance++ {
		// top row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y - distance
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := &sim.World[position.Region].Tiles[x][y]
			if tile.Terrain == terrain {
				return &Position{Region: position.Region, X: x, Y: y}
			}
		}
		// bottom row
		for dx := -distance; dx <= distance; dx++ {
			x := position.X + dx
			y := position.Y + distance
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := &sim.World[position.Region].Tiles[x][y]
			if tile.Terrain == terrain {
				return &Position{Region: position.Region, X: x, Y: y}
			}
		}
		// left row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X - distance
			y := position.Y + dy
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := &sim.World[position.Region].Tiles[x][y]
			if tile.Terrain == terrain {
				return &Position{Region: position.Region, X: x, Y: y}
			}
		}
		// right row
		for dy := -distance; dy <= distance; dy++ {
			x := position.X + distance
			y := position.Y + dy
			if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize {
				continue
			}
			tile := &sim.World[position.Region].Tiles[x][y]
			if tile.Terrain == terrain {
				return &Position{Region: position.Region, X: x, Y: y}
			}
		}
	}
	return nil
}
