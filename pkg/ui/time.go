package ui

import (
	"eramstein/thurigen/pkg/ng"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DisplayTime shows the current time
func (r *Renderer) DisplayTime(sim *ng.Simulation) {
	timeText := fmt.Sprintf("Day %d, Hour %d, Minute %d", sim.Calendar.Day, sim.Calendar.Hour, sim.Calendar.Minute)

	// Draw white background
	rl.DrawRectangle(8, 8, 185, 24, rl.White)

	// Draw text
	r.RenderTextWithColor(timeText, 20, 13, rl.Black)
}
