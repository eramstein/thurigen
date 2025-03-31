package ng

import (
	"eramstein/thurigen/pkg/config"
	"math/rand"
)

func GetInitialWorld() []*Region {
	return []*Region{makeMainRegion()}
}

func makeMainRegion() *Region {
	regionTiles := fillBlankRegion()
	addLake(regionTiles)
	fillSurroundedWaterTiles(regionTiles)
	return &Region{
		Tiles: *regionTiles,
	}
}

func fillBlankRegion() *[config.RegionSize][config.RegionSize]Tile {
	var tiles [config.RegionSize][config.RegionSize]Tile
	for x := 0; x < config.RegionSize; x++ {
		tiles[x] = [config.RegionSize]Tile{}
		for y := 0; y < config.RegionSize; y++ {
			tiles[x][y] = Tile{
				Terrain: Dirt,
				Surface: Grass,
			}
		}
	}
	return &tiles
}

func addLake(tiles *[config.RegionSize][config.RegionSize]Tile) {
	// Random lake parameters
	centerX := rand.Intn(config.RegionSize-20) + 10 // Avoid edges
	centerY := rand.Intn(config.RegionSize-20) + 10
	limitSize := rand.Intn(8) + 10
	directions := [4]float64{0.2, 0.2, 0.2, 0.2}

	// Create lake using flood fill with random variations
	floodFill(tiles, centerX, centerY, limitSize, 2, 0, directions)
}

func fillSurroundedWaterTiles(tiles *[config.RegionSize][config.RegionSize]Tile) {
	for x := 1; x < config.RegionSize-1; x++ {
		for y := 1; y < config.RegionSize-1; y++ {
			if tiles[x][y-1].Terrain == Water && tiles[x+1][y].Terrain == Water && tiles[x][y+1].Terrain == Water && tiles[x-1][y].Terrain == Water {
				tiles[x][y].Terrain = Water
				tiles[x][y].Surface = NoSurface
			}
		}
	}
}

func floodFill(tiles *[config.RegionSize][config.RegionSize]Tile, x, y, limitSize, minSize, currentSize int, directions [4]float64) {
	if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize || tiles[x][y].Terrain == Water {
		return
	}

	if currentSize < limitSize {
		tiles[x][y].Terrain = Water
		tiles[x][y].Surface = NoSurface
		currentSize++
		if currentSize < minSize || rand.Float64() > directions[0] {
			floodFill(tiles, x+1, y, limitSize, minSize, currentSize, directions)
		}
		if currentSize < minSize || rand.Float64() > directions[1] {
			floodFill(tiles, x-1, y, limitSize, minSize, currentSize, directions)
		}
		if currentSize < minSize || rand.Float64() > directions[2] {
			floodFill(tiles, x, y+1, limitSize, minSize, currentSize, directions)
		}
		if currentSize < minSize || rand.Float64() > directions[3] {
			floodFill(tiles, x, y-1, limitSize, minSize, currentSize, directions)
		}
	}
}
