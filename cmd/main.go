package main

import (
	"eramstein/thurigen/pkg/engine"
	"eramstein/thurigen/pkg/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 768
	fps          = 60
)

func main() {
	// Initialize Raylib
	rl.InitWindow(screenWidth, screenHeight, "Thuringen Simulation")
	rl.SetTargetFPS(fps)
	defer rl.CloseWindow()

	ticker := 0
	isPaused := false

	// Initialize sim engine
	sim := engine.NewSimulation()

	// Initialize UI
	renderer := ui.NewRenderer(screenWidth, screenHeight, sim)

	// Main sim loop
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) { // Press Space to toggle pause
			isPaused = !isPaused
		}
		if isPaused {
			rl.DrawText("Paused", screenWidth/2-50, screenHeight/2-10, 20, rl.Red)
		}

		if !isPaused {
			if ticker == sim.Speed {
				go func() {
					sim.Update()
				}()
				ticker = 0
			}
			ticker++
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		renderer.Render(sim)
		rl.EndDrawing()
	}
}
