package ui

import (
	"eramstein/thurigen/pkg/engine"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Renderer handles all UI rendering
type Renderer struct {
	screenWidth       int
	screenHeight      int
	backgroundTexture rl.Texture2D
}

// NewRenderer creates a new renderer instance
func NewRenderer(width, height int, sim *engine.Simulation) *Renderer {
	r := &Renderer{
		screenWidth:  width,
		screenHeight: height,
	}
	// Load the background texture
	r.backgroundTexture = rl.LoadTexture("assets/images/world/domain_castle.png")
	return r
}

// Render renders the current game state
func (r *Renderer) Render(sim *engine.Simulation) {
	// Draw background
	rl.DrawTexture(r.backgroundTexture, 0, 0, rl.White)

	// Draw time on top
	r.DisplayTime(sim)

	if sim.Paused {
		rl.DrawText("Paused", int32(r.screenWidth/2-50), int32(r.screenHeight/2-10), 20, rl.Red)
	}
}

// DisplayTime shows the current time
func (r *Renderer) DisplayTime(sim *engine.Simulation) {
	turnText := fmt.Sprintf("Minutes: %d", sim.Time)
	rl.DrawText(turnText, 10, 10, 20, rl.Black)
}
