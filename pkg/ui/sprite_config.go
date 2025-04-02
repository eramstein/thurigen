package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// SpriteSheetConfig defines the configuration for a sprite sheet
type SpriteSheetConfig struct {
	Name     string
	Path     string
	TileSize int32
	Sprites  map[int]rl.Rectangle
}

// GetSpriteSheetConfigs returns all sprite sheet configurations
func GetSpriteSheetConfigs() []SpriteSheetConfig {
	return []SpriteSheetConfig{
		{
			Name:     "trees",
			Path:     "assets/images/world/trees.png",
			TileSize: 32,
			Sprites: map[int]rl.Rectangle{
				0: rl.NewRectangle(0, 0, 32, 32),
				1: rl.NewRectangle(32, 0, 32, 32),
				2: rl.NewRectangle(64, 0, 32, 32),
			},
		},
	}
}
