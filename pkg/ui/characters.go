package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Renderer) DisplayCharacters(Characters []*ng.Character) {
	for _, character := range Characters {
		screenX := float32(character.Position.X * config.TilePixelSize)
		screenY := float32(character.Position.Y * config.TilePixelSize)

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
			case ng.Eat, ng.Drink, ng.Sleep:
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
