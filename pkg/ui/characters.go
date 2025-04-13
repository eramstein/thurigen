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
	}
}
