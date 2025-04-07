package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"
	"fmt"

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

func (r *Renderer) DisplayRegion() {
	region := r.uiState.DisplayedRegion

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

	// First pass: Render all surface rectangles
	for y := startTileY; y < endTileY; y++ {
		for x := startTileX; x < endTileX; x++ {
			r.RenderTileSurface(region.Tiles[x][y], x, y)
		}
	}

	// Second pass: Render all structures
	for y := startTileY; y < endTileY; y++ {
		for x := startTileX; x < endTileX; x++ {
			r.RenderTileStructures(region.Tiles[x][y], x, y)
		}
	}

	// Third pass: Render all items
	for y := startTileY; y < endTileY; y++ {
		for x := startTileX; x < endTileX; x++ {
			r.RenderTileItems(region.Tiles[x][y], x, y)
		}
	}

	//r.DrawGridLines(startTileX, startTileY, endTileX, endTileY)
	r.BorderSelectedTile()
}

func (r *Renderer) BorderSelectedTile() {
	if r.uiState.SelectedTile != nil {
		s := *r.uiState.SelectedTile
		screenX := float32(s[0] * config.TilePixelSize)
		screenY := float32(s[1] * config.TilePixelSize)
		rl.DrawRectangleLines(int32(screenX), int32(screenY), int32(config.TilePixelSize), int32(config.TilePixelSize), rl.White)
	}
}

func (r *Renderer) RenderTileSurface(tile ng.Tile, x, y int) {
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

}

func (r *Renderer) RenderTileStructures(tile ng.Tile, x, y int) {
	// Calculate screen position
	screenX := float32(x * config.TilePixelSize)
	screenY := float32(y * config.TilePixelSize)

	// Draw structure information if the tile is occupied
	if tile.Occupation != nil && tile.Occupation.Structure != nil {
		structure := tile.Occupation.Structure
		base := structure.GetStructure()
		// Get the sprite sheet name for this structure type
		sheetConfig := structureToSpriteSheet[base.Type]
		sheet := r.spriteManager.sheets[sheetConfig.Name]
		if spriteRect, exists := sheet.Sprites[base.Variant]; exists {
			// Draw the sprite centered in the tile
			if plant, ok := structure.(*ng.PlantStructure); ok {
				r.RenderPlant(spriteRect, sheet.Texture, screenX, screenY, plant.GrowthStage)
			} else {
				r.RenderStructure(spriteRect, sheet.Texture, screenX, screenY)
			}
		}
	}
}

func (r *Renderer) RenderTileItems(tile ng.Tile, x, y int) {
	// Calculate screen position
	screenX := float32(x * config.TilePixelSize)
	screenY := float32(y * config.TilePixelSize)

	// Only render if there are items
	if len(tile.Items) > 0 {
		itemCount := len(tile.Items)
		itemText := fmt.Sprintf("%d", itemCount)

		// Position the text in the bottom-right corner of the tile
		textX := screenX + 3*float32(config.TilePixelSize)/4
		textY := screenY + 3*float32(config.TilePixelSize)/4

		// Draw a semi-transparent background for better visibility
		rl.DrawRectangle(int32(textX-2), int32(textY-2), 12, 12, rl.NewColor(0, 0, 0, 128))

		// Draw the text
		rl.DrawText(itemText, int32(textX), int32(textY), 10, rl.White)
	}
}

func (r *Renderer) RenderPlant(spriteRect rl.Rectangle, texture rl.Texture2D, screenX, screenY float32, growthStage int) {
	scale := 0.3 + (float32(growthStage) / 100.0 * 0.7)
	scaledsize := float32(config.TilePixelSize) * scale
	offset := (config.TilePixelSize - scaledsize) / 2
	rl.DrawTexturePro(
		texture,
		spriteRect,
		rl.NewRectangle(screenX+offset, screenY+offset, scaledsize, scaledsize),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

func (r *Renderer) RenderStructure(spriteRect rl.Rectangle, texture rl.Texture2D, screenX, screenY float32) {
	rl.DrawTextureRec(
		texture,
		spriteRect,
		rl.NewVector2(screenX, screenY),
		rl.White,
	)
}

// DrawGridLines draws white lines to separate tiles
func (r *Renderer) DrawGridLines(startTileX, startTileY, endTileX, endTileY int) {
	// Draw vertical lines
	for x := startTileX; x <= endTileX; x++ {
		lineX := float32(x * config.TilePixelSize)
		startY := float32(startTileY * config.TilePixelSize)
		endY := float32(endTileY * config.TilePixelSize)
		rl.DrawLineEx(
			rl.Vector2{X: lineX, Y: startY},
			rl.Vector2{X: lineX, Y: endY},
			1.0, // Line thickness
			rl.White,
		)
	}

	// Draw horizontal lines
	for y := startTileY; y <= endTileY; y++ {
		lineY := float32(y * config.TilePixelSize)
		startX := float32(startTileX * config.TilePixelSize)
		endX := float32(endTileX * config.TilePixelSize)
		rl.DrawLineEx(
			rl.Vector2{X: startX, Y: lineY},
			rl.Vector2{X: endX, Y: lineY},
			1.0, // Line thickness
			rl.White,
		)
	}
}
