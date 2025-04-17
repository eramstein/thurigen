package ng

import (
	"eramstein/thurigen/pkg/config"
	"math/rand"
)

func (sim *Simulation) InitWorld() {
	mainRegion := makeMainRegion()
	sim.World = []*Region{mainRegion}
	sim.addRandomTrees()
}

func makeMainRegion() *Region {
	regionTiles := fillBlankRegion()
	addLake(regionTiles)
	addMountain(regionTiles)
	fillSurroundedTiles(regionTiles)
	computeMoveCosts(regionTiles)
	return &Region{
		Tiles: *regionTiles,
	}
}

func fillBlankRegion() *[config.RegionSize][config.RegionSize]Tile {
	var tiles [config.RegionSize][config.RegionSize]Tile
	for x := 0; x < config.RegionSize; x++ {
		tiles[x] = [config.RegionSize]Tile{}
		for y := 0; y < config.RegionSize; y++ {
			if (x == 0 || x == config.RegionSize-1) || (y == 0 || y == config.RegionSize-1) {
				tiles[x][y] = Tile{
					Terrain: Dirt,
					Surface: Grass,
					Volume:  RockVolume,
				}
			} else {
				tiles[x][y] = Tile{
					Terrain: Dirt,
					Surface: Grass,
				}
			}
		}
	}
	return &tiles
}

func computeMoveCosts(tiles *[config.RegionSize][config.RegionSize]Tile) {
	for x := 0; x < config.RegionSize; x++ {
		for y := 0; y < config.RegionSize; y++ {
			tiles[x][y].MoveCost = getTileMoseCost(tiles[x][y])
		}
	}
}

func getTileMoseCost(tile Tile) MoveCost {
	if tile.Terrain == Water {
		return ImpassableCost
	}
	if tile.Volume == RockVolume {
		return ImpassableCost
	}
	return DefaultMoveCost
}

func addLake(tiles *[config.RegionSize][config.RegionSize]Tile) {
	// Random lake parameters
	centerX := rand.Intn(config.RegionSize-20) + 10 // Avoid edges
	centerY := rand.Intn(config.RegionSize-20) + 10
	limitSize := rand.Intn(8) + 10
	directions := [4]float64{0.2, 0.2, 0.2, 0.2}

	// Create lake using flood fill with random variations
	floodFillWater(tiles, centerX, centerY, limitSize, 2, 0, directions)
}

func addMountain(tiles *[config.RegionSize][config.RegionSize]Tile) {
	// Random lake parameters
	centerX := rand.Intn(config.RegionSize-20) + 10 // Avoid edges
	centerY := rand.Intn(config.RegionSize-20) + 10
	limitSize := rand.Intn(8) + 20
	directions := [4]float64{0.2, 0.4, 0.2, 0.4}

	// Create mountain using flood fill with random variations
	floodFillMountain(tiles, centerX, centerY, limitSize, 2, 0, directions)
}

func fillSurroundedTiles(tiles *[config.RegionSize][config.RegionSize]Tile) {
	for x := 1; x < config.RegionSize-1; x++ {
		for y := 1; y < config.RegionSize-1; y++ {
			if tiles[x][y-1].Terrain == Water && tiles[x+1][y].Terrain == Water && tiles[x][y+1].Terrain == Water && tiles[x-1][y].Terrain == Water {
				tiles[x][y].Terrain = Water
				tiles[x][y].Surface = NoSurface
			}
			if tiles[x][y-1].Volume == RockVolume && tiles[x+1][y].Volume == RockVolume && tiles[x][y+1].Volume == RockVolume && tiles[x-1][y].Volume == RockVolume {
				tiles[x][y].Volume = RockVolume
			}
		}
	}
}

func floodFillWater(tiles *[config.RegionSize][config.RegionSize]Tile, x, y, limitSize, minSize, currentSize int, directions [4]float64) {
	if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize || tiles[x][y].Terrain == Water {
		return
	}

	if currentSize < limitSize {
		tiles[x][y].Terrain = Water
		tiles[x][y].Surface = NoSurface
		currentSize++
		if currentSize < minSize || rand.Float64() > directions[0] {
			floodFillWater(tiles, x+1, y, limitSize, minSize, currentSize, directions)
		}
		if currentSize < minSize || rand.Float64() > directions[1] {
			floodFillWater(tiles, x-1, y, limitSize, minSize, currentSize, directions)
		}
		if currentSize < minSize || rand.Float64() > directions[2] {
			floodFillWater(tiles, x, y+1, limitSize, minSize, currentSize, directions)
		}
		if currentSize < minSize || rand.Float64() > directions[3] {
			floodFillWater(tiles, x, y-1, limitSize, minSize, currentSize, directions)
		}
	}
}

func floodFillMountain(tiles *[config.RegionSize][config.RegionSize]Tile, x, y, limitSize, minSize, currentSize int, directions [4]float64) {
	if x < 0 || x >= config.RegionSize || y < 0 || y >= config.RegionSize || tiles[x][y].Volume == RockVolume {
		return
	}

	if currentSize < limitSize {
		tiles[x][y].Terrain = Dirt
		tiles[x][y].Surface = NoSurface
		tiles[x][y].Volume = RockVolume
		currentSize++
		if rand.Float64() > directions[0] {
			floodFillMountain(tiles, x+1, y, limitSize, minSize, currentSize, directions)
		}
		if rand.Float64() > directions[1] {
			floodFillMountain(tiles, x-1, y, limitSize, minSize, currentSize, directions)
		}
		if rand.Float64() > directions[2] {
			floodFillMountain(tiles, x, y+1, limitSize, minSize, currentSize, directions)
		}
		if rand.Float64() > directions[3] {
			floodFillMountain(tiles, x, y-1, limitSize, minSize, currentSize, directions)
		}
	}
}

// addRandomTrees adds a random number of trees to the world
func (sim *Simulation) addRandomTrees() {
	// Number of trees to add (adjust these values as needed)
	minTrees := 20
	maxTrees := 40
	numTrees := rand.Intn(maxTrees-minTrees+1) + minTrees

	for i := 0; i < numTrees; i++ {
		x := rand.Intn(config.RegionSize)
		y := rand.Intn(config.RegionSize)
		tile := &sim.World[0].Tiles[x][y]

		if isValidTreeTile(tile) {
			sim.SpawnPlant(0, x, y, rand.Intn(2)+1)
		}
	}
}

// isValidTreeTile checks if a tile is suitable for a tree
func isValidTreeTile(tile *Tile) bool {
	// Trees can't be placed on water or rock
	if tile.Terrain == Water || tile.Volume == RockVolume {
		return false
	}

	// Trees can't be placed on existing structures
	if tile.Occupation != nil {
		return false
	}

	return tile.Surface == Grass || tile.Surface == NoSurface
}
