package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Renderer) DisplayCharacters(sim *ng.Simulation) {
	for _, character := range sim.Characters {
		screenX := float32(character.Position.X * config.TilePixelSize)
		screenY := float32(character.Position.Y * config.TilePixelSize)
		prevScreenX := float32(r.UiState.PreviousCharacterPositions[character.ID].X * config.TilePixelSize)
		prevScreenY := float32(r.UiState.PreviousCharacterPositions[character.ID].Y * config.TilePixelSize)

		interpolationFactor := float32(r.UiState.Ticker) / float32(sim.Speed)
		screenX = prevScreenX + (screenX-prevScreenX)*interpolationFactor
		screenY = prevScreenY + (screenY-prevScreenY)*interpolationFactor

		sheet := r.spriteManager.sheets[charactersSpriteSheet.Name]
		if spriteRect, exists := sheet.Sprites[character.ID]; exists {
			// Draw the sprite centered in the tile
			rl.DrawTextureRec(
				sheet.Texture,
				spriteRect,
				rl.NewVector2(screenX, screenY),
				rl.White,
			)
		}

		// Draw task icon if character has a current task that needs an icon
		if character.CurrentTask != nil {
			switch character.CurrentTask.Type {
			case ng.Eat, ng.Drink, ng.Sleep, ng.Build, ng.Chop:
				taskSheet := r.spriteManager.sheets[taskIconsSpriteSheet.Name]
				if taskSpriteRect, exists := taskSheet.Sprites[uint64(character.CurrentTask.Type)]; exists {
					// Draw the task icon in the top-left corner of the character's tile
					rl.DrawTextureRec(
						taskSheet.Texture,
						taskSpriteRect,
						rl.NewVector2(screenX-float32(taskSheet.TileSize)/2, screenY-float32(taskSheet.TileSize)/2), // Small offset from the top-left corner
						rl.White,
					)
				}
			}
		}
	}
}

func (r *Renderer) DisplayCharacterPortraits(sim *ng.Simulation) {
	for i, character := range sim.Characters {
		sheet := r.spriteManager.sheets[charactersSpriteSheet.Name]
		if spriteRect, exists := sheet.Sprites[character.ID]; exists {
			// Calculate position for this portrait
			x := float32(config.PortraitStartX) + float32(i)*(float32(config.PortraitSize)+float32(config.PortraitSpacing))
			y := float32(config.PortraitStartY)

			// Draw a background rectangle for the portrait
			rl.DrawRectangle(
				int32(x-2),
				int32(y-2),
				int32(float32(config.PortraitSize)+4),
				int32(float32(config.PortraitSize)+4),
				rl.NewColor(200, 200, 200, 128),
			)

			// Draw the character portrait
			rl.DrawTexturePro(
				sheet.Texture,
				spriteRect,
				rl.NewRectangle(x, y, float32(config.PortraitSize), float32(config.PortraitSize)),
				rl.Vector2{X: 10, Y: 0},
				0,
				rl.White,
			)
		}
	}
}
