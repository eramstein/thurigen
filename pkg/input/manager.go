package input

import (
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Manager handles all input events
type Manager struct {
	mousePosition rl.Vector2
	leftPressed   bool
	rightPressed  bool
}

// NewManager creates a new input manager
func NewManager() *Manager {
	return &Manager{}
}

// Update updates the input state
func (m *Manager) Update(sim *ng.Simulation) {
	if rl.IsKeyPressed(rl.KeySpace) { // Press Space to toggle pause
		sim.Paused = !sim.Paused
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		rl.CloseWindow()
	}

	// Update mouse state
	m.mousePosition = rl.GetMousePosition()
	m.leftPressed = rl.IsMouseButtonPressed(rl.MouseLeftButton)
	m.rightPressed = rl.IsMouseButtonPressed(rl.MouseRightButton)
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
