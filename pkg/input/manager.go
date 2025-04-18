package input

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"
	"eramstein/thurigen/pkg/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Manager handles all input events
type Manager struct {
	mousePosition rl.Vector2
	leftPressed   bool
	rightPressed  bool
	camera        *rl.Camera2D
}

// NewManager creates a new input manager
func NewManager() *Manager {
	return &Manager{}
}

// SetCamera sets the camera for coordinate conversion
func (m *Manager) SetCamera(camera rl.Camera2D) {
	m.camera = &camera
}

// Update updates the input state
func (m *Manager) Update(sim *ng.Simulation, renderer *ui.Renderer) {
	m.SetCamera(renderer.GetCamera())

	if rl.IsKeyPressed(rl.KeySpace) { // Press Space to toggle pause
		sim.Paused = !sim.Paused
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		rl.CloseWindow()
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		sim.Benchmark()
	}

	if rl.IsKeyPressed(rl.KeyF5) { // Press F5 to save
		if err := sim.SaveState(); err != nil {
			fmt.Printf("Failed to save simulation state: %v\n", err)
		} else {
			fmt.Println("Simulation state saved successfully")
			fmt.Println(len(sim.World[0].Plants))
		}
	}

	if rl.IsKeyPressed(rl.KeyF4) { // Press F4 to load latest save
		if loadedSim, err := ng.LoadLatestState(); err != nil {
			fmt.Printf("Failed to load latest save: %v\n", err)
		} else {
			// Replace current simulation with loaded one
			*sim = *loadedSim
			fmt.Println("Latest save loaded successfully")
			fmt.Println(len(sim.World[0].Plants))
		}
	}

	// Update mouse state
	m.mousePosition = rl.GetMousePosition()
	m.leftPressed = rl.IsMouseButtonPressed(rl.MouseLeftButton)
	m.rightPressed = rl.IsMouseButtonPressed(rl.MouseRightButton)

	// Handle clicks
	if m.leftPressed && m.camera != nil {
		tileX, tileY := m.ScreenToTileCoordinates(m.mousePosition)
		if tileX >= 0 && tileX < config.RegionSize && tileY >= 0 && tileY < config.RegionSize {
			renderer.ToggleTileSelection(tileX, tileY)
		}
	}
}

// ScreenToTileCoordinates converts screen coordinates to tile coordinates
func (m *Manager) ScreenToTileCoordinates(screenPos rl.Vector2) (int, int) {
	if m.camera == nil {
		return -1, -1
	}
	worldPos := rl.GetScreenToWorld2D(screenPos, *m.camera)
	tileX := int(worldPos.X) / config.TilePixelSize
	tileY := int(worldPos.Y) / config.TilePixelSize
	return tileX, tileY
}

// IsKeyJustPressed returns true if the key was just pressed
func (m *Manager) IsKeyJustPressed(key int32) bool {
	return rl.IsKeyPressed(key)
}

// GetMousePosition returns the current mouse position
func (m *Manager) GetMousePosition() rl.Vector2 {
	return m.mousePosition
}

// IsLeftMousePressed returns true if the left mouse button is pressed
func (m *Manager) IsLeftMousePressed() bool {
	return m.leftPressed
}

// IsRightMousePressed returns true if the right mouse button is pressed
func (m *Manager) IsRightMousePressed() bool {
	return m.rightPressed
}
