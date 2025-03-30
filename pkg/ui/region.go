package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var terrainColors = map[engine.TerrainType]rl.Color{
	engine.Dirt:  rl.Brown,
	engine.Rock:  rl.Gray,
	engine.Sand:  rl.Yellow,
	engine.Water: rl.Blue,
}

var surfaceColors = map[engine.SurfaceType]rl.Color{
	engine.Grass:       rl.Green,
	engine.WoodSurface: rl.DarkBrown,
}

var volumeColors = map[engine.VolumeType]rl.Color{
	engine.RockVolume: rl.DarkGray,
	engine.WoodVolume: rl.Maroon,
}

func (r *Renderer) DisplayRegion(region engine.Region) {
	for y := 0; y < config.RegionSize; y++ {
		for x := 0; x < config.RegionSize; x++ {
			// Get the tile at current position
			tile := region.Tiles[y][x]

			// Calculate screen position
			screenX := float32(x * config.TilePixelSize)
			screenY := float32(y * config.TilePixelSize)

			// Draw the tile rectangle with the appropriate color
			color := terrainColors[tile.Terrain]
			if tile.Surface != 0 {
				color = surfaceColors[tile.Surface]
			}
			if tile.Volume != 0 {
				color = volumeColors[tile.Volume]
			}

			rl.DrawRectangle(int32(screenX), int32(screenY), int32(config.TilePixelSize), int32(config.TilePixelSize), color)
		}
	}
}
