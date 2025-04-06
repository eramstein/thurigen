package ui

import (
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// SpriteSheetConfig defines the configuration for a sprite sheet
type SpriteSheetConfig struct {
	Name     string
	Path     string
	TileSize int32
	Sprites  map[int]rl.Rectangle
}

var structureToSpriteSheet = map[ng.StructureType]SpriteSheetConfig{
	ng.Plant: {
		Name:     "plants",
		Path:     "assets/images/world/plants.png",
		TileSize: 32,
		Sprites: map[int]rl.Rectangle{
			1: rl.NewRectangle(0, 0, 32, 32),
			2: rl.NewRectangle(32, 0, 32, 32),
			3: rl.NewRectangle(64, 0, 32, 32),
		},
	},
}
