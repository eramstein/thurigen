package ui

import (
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
	uiState       *Model
}

// NewRenderer creates a new renderer instance
func NewRenderer(width, height int, sim *ng.Simulation) *Renderer {
	r := &Renderer{
		screenWidth:   width,
		screenHeight:  height,
		spriteManager: NewSpriteManager(),
		uiState: &Model{
			DisplayedRegion: sim.World[0],
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

	return nil
}

// Render renders the current game state
func (r *Renderer) Render(sim *ng.Simulation) {
	// Update camera
	r.camera.Update()

	// Begin camera drawing
	rl.BeginMode2D(r.camera.GetCamera())

	r.DisplayRegion()

	// End camera drawing
	rl.EndMode2D()

	if sim.Paused {
		rl.DrawText("Paused", int32(r.screenWidth/2-50), int32(r.screenHeight/2-10), 20, rl.Red)
	}

	// Draw time on top
	r.DisplayTime(sim)
}

// DisplayTime shows the current time
func (r *Renderer) DisplayTime(sim *ng.Simulation) {
	turnText := fmt.Sprintf("Minutes: %d", sim.Time)
	rl.DrawText(turnText, 10, 10, 20, rl.Black)
}
