package main

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/input"
	"eramstein/thurigen/pkg/ng"
	"eramstein/thurigen/pkg/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = config.WindowWidth
	screenHeight = config.WindowHeight
	fps          = config.FPS
)

func main() {
	// Initialize Raylib
	rl.InitWindow(screenWidth, screenHeight, "Thuringen Simulation")
	rl.SetTargetFPS(fps)
	defer rl.CloseWindow()

	// Load static data (configs, etc)
	loadData()

	// Initialize sim engine
	var sim *ng.Simulation
	var err error

	// Try to load the latest saved state
	sim, err = ng.LoadLatestState()
	if err != nil {
		fmt.Println("No saved state found, creating new simulation:", err)
		sim = ng.NewSimulation()
	} else {
		fmt.Println("Loaded saved simulation state")
	}

	// Initialize UI
	renderer := ui.NewRenderer(screenWidth, screenHeight, sim)

	// Load textures
	if err := renderer.LoadTextures(); err != nil {
		rl.CloseWindow()
		panic(err)
	}

	// Initialize input manager
	inputManager := input.NewManager()

	// Main sim loop
	for !rl.WindowShouldClose() {
		inputManager.Update(sim, renderer)
		if !sim.Paused {
			if renderer.UiState.Ticker == sim.Speed {
				// TODO: is this efficient? we create one routine per second, check how they are handled
				go func() {
					renderer.UpdatePreviousCharacterPositions(sim)
					sim.Update()
				}()
				renderer.UiState.Ticker = 0
			}
			renderer.UiState.Ticker++
		}
		render(renderer, sim)
	}
}

func render(renderer *ui.Renderer, sim *ng.Simulation) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	renderer.Render(sim)
	rl.EndDrawing()
}

func loadData() {
	ng.LoadStructuresConfigs()
	ng.LoadItemsConfigs()
}
