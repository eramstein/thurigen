package ui

import (
	"eramstein/thurigen/pkg/ng"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Renderer handles all UI rendering
type Renderer struct {
	screenWidth  int
	screenHeight int
	camera       *Camera
}

// NewRenderer creates a new renderer instance
func NewRenderer(width, height int, sim *ng.Simulation) *Renderer {
	r := &Renderer{
		screenWidth:  width,
		screenHeight: height,
	}
	r.camera = NewCamera(width, height)
	return r
}

// Render renders the current game state
func (r *Renderer) Render(sim *ng.Simulation) {
	// Update camera
	r.camera.Update()

	// Begin camera drawing
	rl.BeginMode2D(r.camera.GetCamera())

	r.DisplayRegion(sim.World[0])

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
