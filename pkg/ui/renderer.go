package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Renderer handles all UI rendering
type Renderer struct {
	screenWidth   int
	screenHeight  int
	camera        *Camera
	spriteManager *SpriteManager
	fontManager   *FontManager
	defaultFont   rl.Font // Cache the default font
	uiState       *Model
}

// NewRenderer creates a new renderer instance
func NewRenderer(width, height int, sim *ng.Simulation) *Renderer {
	r := &Renderer{
		screenWidth:   width,
		screenHeight:  height,
		spriteManager: NewSpriteManager(),
		fontManager:   NewFontManager(),
		uiState: &Model{
			DisplayedRegion: 0,
		},
	}
	r.camera = NewCamera(width, height)
	return r
}

// GetCamera returns the current camera state
func (r *Renderer) GetCamera() rl.Camera2D {
	return r.camera.GetCamera()
}

// LoadTextures loads all required textures for the renderer
func (r *Renderer) LoadTextures() error {
	// Load each sprite sheet
	for _, config := range structureToSpriteSheet {
		if err := r.spriteManager.LoadSpriteSheet(config.Name, config.Path, config.TileSize, config.Sprites); err != nil {
			return fmt.Errorf("failed to load sprite sheet %s: %v", config.Name, err)
		}
	}

	// Load character sprite sheet
	if err := r.spriteManager.LoadSpriteSheet(charactersSpriteSheet.Name, charactersSpriteSheet.Path, charactersSpriteSheet.TileSize, charactersSpriteSheet.Sprites); err != nil {
		return fmt.Errorf("failed to load sprite sheet %s: %v", charactersSpriteSheet.Name, err)
	}

	// Load fonts
	if err := r.fontManager.LoadFont("default", "assets/fonts/Roboto-Regular.ttf", config.BaseFontSize); err != nil {
		return fmt.Errorf("failed to load default font: %v", err)
	}

	// Cache the default font
	r.defaultFont = r.fontManager.GetFont("default")

	return nil
}

// Render renders the current game state
func (r *Renderer) Render(sim *ng.Simulation) {
	// Update camera
	r.camera.Update()

	// Begin camera drawing
	rl.BeginMode2D(r.camera.GetCamera())

	r.DisplayRegion(sim)
	r.DisplayCharacters(sim.Characters)

	// End camera drawing
	rl.EndMode2D()

	if sim.Paused {
		r.RenderTextWithColor("Paused", r.screenWidth/2-50, r.screenHeight/2-10, rl.Red)
	}

	// Draw time on top
	r.DisplayTime(sim)

	// Draw side panel
	r.DisplayTileSidePanel(sim)
}

// DisplayTime shows the current time
func (r *Renderer) DisplayTime(sim *ng.Simulation) {
	turnText := fmt.Sprintf("Minutes: %d", sim.Time)
	r.RenderText(turnText, 10, 10)
}

// RenderText renders text at a specific position
func (r *Renderer) RenderText(text string, x, y int) {
	rl.DrawTextEx(r.defaultFont, text, rl.Vector2{X: float32(x), Y: float32(y)}, float32(r.defaultFont.BaseSize), 1.0, rl.Black)
}

// RenderTextWithColor renders text at a specific position with a specific color
func (r *Renderer) RenderTextWithColor(text string, x, y int, color rl.Color) {
	rl.DrawTextEx(r.defaultFont, text, rl.Vector2{X: float32(x), Y: float32(y)}, float32(r.defaultFont.BaseSize), 1.0, color)
}
