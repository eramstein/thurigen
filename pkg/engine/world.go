package engine

import "eramstein/thurigen/pkg/config"

func GetInitialWorld() []Region {
	return []Region{makeMainRegion()}
}

func makeMainRegion() Region {
	var tiles [][]Tile
	for x := 0; x < config.RegionSize; x++ {
		tiles = append(tiles, make([]Tile, config.RegionSize))
		for y := 0; y < config.RegionSize; y++ {
			tiles[x][y] = Tile{
				Terrain: func() TerrainType {
					if x == 0 || y == 0 {
						return Rock
					}
					if x > 10 && x < 20 && y > 5 && y < 35 {
						return Water
					}
					return Dirt
				}(),
				Surface: Grass,
				Volume:  RockVolume,
			}
		}
	}
	return Region{
		Tiles: tiles,
	}
}
