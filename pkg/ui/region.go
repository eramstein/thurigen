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
	// Get the camera's view bounds
	camera := r.camera.GetCamera()

	// Calculate visible area in world coordinates
	screenLeftTop := rl.GetScreenToWorld2D(rl.Vector2{X: 0, Y: 0}, camera)
	screenRightBottom := rl.GetScreenToWorld2D(rl.Vector2{
		X: float32(r.screenWidth),
		Y: float32(r.screenHeight),
	}, camera)

	// Convert world coordinates to tile indices
	startTileX := int(screenLeftTop.X) / config.TilePixelSize
	startTileY := int(screenLeftTop.Y) / config.TilePixelSize
	endTileX := int(screenRightBottom.X)/config.TilePixelSize + 1 // +1 to handle partial tiles
	endTileY := int(screenRightBottom.Y)/config.TilePixelSize + 1

	// Clamp values to region bounds
	startTileX = clamp(startTileX, 0, config.RegionSize-1)
	startTileY = clamp(startTileY, 0, config.RegionSize-1)
	endTileX = clamp(endTileX, 0, config.RegionSize)
	endTileY = clamp(endTileY, 0, config.RegionSize)

	// Only render tiles within the visible area
	for y := startTileY; y < endTileY; y++ {
		for x := startTileX; x < endTileX; x++ {
			r.RenderTile(region.Tiles[y][x], x, y)
		}
	}
}

func (r *Renderer) RenderTile(tile ng.Tile, x, y int) {
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
		base := structure.GetStructure()
		// Get the sprite sheet name for this structure type
		sheetConfig := structureToSpriteSheet[base.Type]
		sheet := r.spriteManager.sheets[sheetConfig.Name]
		if spriteRect, exists := sheet.Sprites[base.Variant]; exists {
			// Draw the sprite centered in the tile
			if plant, ok := structure.(*ng.Plant); ok {
				scale := 0.3 + (float32(plant.GrowthStage) / 100.0 * 0.7)
				scaledsize := float32(config.TilePixelSize) * scale
				offset := scaledsize / 2
				rl.DrawTexturePro(
					sheet.Texture,
					spriteRect,
					rl.NewRectangle(screenX-offset, screenY-offset, scaledsize, scaledsize),
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
			} else {
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
