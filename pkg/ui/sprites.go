package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// SpriteSheet represents a loaded texture and its sprite regions
type SpriteSheet struct {
	Texture  rl.Texture2D
	Sprites  map[uint64]rl.Rectangle
	TileSize int32
}

// SpriteManager handles all sprite sheets and their loading
type SpriteManager struct {
	sheets map[string]*SpriteSheet
}

// NewSpriteManager creates a new sprite manager
func NewSpriteManager() *SpriteManager {
	return &SpriteManager{
		sheets: make(map[string]*SpriteSheet),
	}
}

// LoadSpriteSheet loads a sprite sheet from a file and defines its sprite regions
func (sm *SpriteManager) LoadSpriteSheet(name, path string, tileSize int32, spriteRegions map[uint64]rl.Rectangle) error {
	texture := rl.LoadTexture(path)
	if texture.ID == 0 {
		return fmt.Errorf("failed to load sprite sheet: %s", path)
	}

	sm.sheets[name] = &SpriteSheet{
		Texture:  texture,
		Sprites:  spriteRegions,
		TileSize: tileSize,
	}

	return nil
}

// UnloadAll unloads all loaded sprite sheets
func (sm *SpriteManager) UnloadAll() {
	for _, sheet := range sm.sheets {
		rl.UnloadTexture(sheet.Texture)
	}
	sm.sheets = make(map[string]*SpriteSheet)
}
