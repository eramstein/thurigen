package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	terrainColors = map[ng.TerrainType]rl.Color{
		ng.Dirt:  rl.Brown,
		ng.Rock:  rl.Gray,
		ng.Sand:  rl.Yellow,
		ng.Water: rl.Blue,
	}

	surfaceColors = map[ng.SurfaceType]rl.Color{
		ng.Grass:       rl.DarkGreen,
		ng.WoodSurface: rl.DarkBrown,
	}

	volumeColors = map[ng.VolumeType]rl.Color{
		ng.RockVolume: rl.DarkGray,
		ng.WoodVolume: rl.Maroon,
	}
)

func (r *Renderer) DisplayRegion(region *ng.Region) {
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

			// Draw the main tile
			rl.DrawRectangle(int32(screenX), int32(screenY), int32(config.TilePixelSize), int32(config.TilePixelSize), color)

			// Draw structure information if the tile is occupied
			if tile.Occupation != nil && tile.Occupation.Structure != nil {
				structure := tile.Occupation.Structure
				variantIndex := structure.Variant
				// Get the appropriate sprite sheet based on structure type
				if structure.Type == ng.Tree {
					if sheet, exists := r.spriteManager.GetSpriteSheet("trees"); exists {
						if spriteRect, exists := sheet.Sprites[variantIndex]; exists {
							// Draw the sprite centered in the tile
							rl.DrawTextureRec(
								sheet.Texture,
								spriteRect,
								rl.NewVector2(screenX, screenY),
								rl.White,
							)
						}
					}
				}
			}
		}
	}
}
