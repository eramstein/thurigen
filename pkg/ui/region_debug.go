package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (r *Renderer) HighlightSelectedCharacterPath(character *ng.Character) {
	if r.UiState.SelectedCharacter != nil && r.UiState.SelectedCharacter.Path != nil {
		path := *r.UiState.SelectedCharacter.Path
		for _, pos := range path {
			if pos.Region == r.UiState.DisplayedRegion {
				screenX := float32(pos.X * config.TilePixelSize)
				screenY := float32(pos.Y * config.TilePixelSize)
				// Draw a semi-transparent blue rectangle over the path tile
				rl.DrawRectangle(
					int32(screenX),
					int32(screenY),
					int32(config.TilePixelSize),
					int32(config.TilePixelSize),
					rl.NewColor(0, 0, 255, 128),
				)
			}
		}
	}
}

func (r *Renderer) HighlightSelectedCharacterObjectives(character *ng.Character) {
	if character == nil {
		return
	}

	for _, objective := range character.Objectives {
		for _, task := range objective.Plan {
			if task.Target != nil {
				if pos, ok := task.Target.(*ng.Position); ok && pos != nil {
					if pos.Region == r.UiState.DisplayedRegion {
						screenX := float32(pos.X * config.TilePixelSize)
						screenY := float32(pos.Y * config.TilePixelSize)
						// Draw a semi-transparent yellow rectangle over the objective tile
						rl.DrawRectangle(
							int32(screenX),
							int32(screenY),
							int32(config.TilePixelSize),
							int32(config.TilePixelSize),
							rl.NewColor(255, 255, 0, 128),
						)
					}
				}
			}
		}
	}
}
