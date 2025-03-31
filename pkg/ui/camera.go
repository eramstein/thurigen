package ui

import (
	"eramstein/thurigen/pkg/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Camera handles camera operations and state
type Camera struct {
	camera      rl.Camera2D
	cameraSpeed float32
	zoom        float32
}

// NewCamera creates a new camera instance
func NewCamera(width, height int) *Camera {
	c := &Camera{
		cameraSpeed: 500.0,
		zoom:        1.0,
	}

	// Initialize camera
	regionSize := float32(config.RegionSize * config.TilePixelSize)
	c.camera = rl.Camera2D{
		Offset:   rl.Vector2{X: float32(width) / 2, Y: float32(height) / 2},
		Target:   rl.Vector2{X: regionSize / 2, Y: regionSize / 2},
		Rotation: 0,
		Zoom:     1.0,
	}

	return c
}

// Update updates the camera position based on input
func (c *Camera) Update() {
	// Calculate camera bounds based on region size and zoom
	regionSize := float32(config.RegionSize * config.TilePixelSize * c.camera.Zoom)

	maxX := regionSize
	maxY := regionSize
	minX := float32(0.0)
	minY := float32(0.0)

	// Camera movement with WASD
	if rl.IsKeyDown(rl.KeyW) {
		c.camera.Target.Y -= c.cameraSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyS) {
		c.camera.Target.Y += c.cameraSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyA) {
		c.camera.Target.X -= c.cameraSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyD) {
		c.camera.Target.X += c.cameraSpeed * rl.GetFrameTime()
	}

	// Camera panning with right mouse button
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		delta := rl.GetMouseDelta()
		c.camera.Target.X -= delta.X / c.camera.Zoom
		c.camera.Target.Y -= delta.Y / c.camera.Zoom
	}

	// Camera zoom
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		c.zoom += wheel * 0.1
		if c.zoom < 0.1 {
			c.zoom = 0.1
		}
		if c.zoom > 3.0 {
			c.zoom = 3.0
		}
		c.camera.Zoom = c.zoom
	}

	// Clamp camera target to bounds
	if c.camera.Target.X > maxX {
		c.camera.Target.X = maxX
	}
	if c.camera.Target.X < minX {
		c.camera.Target.X = minX
	}
	if c.camera.Target.Y > maxY {
		c.camera.Target.Y = maxY
	}
	if c.camera.Target.Y < minY {
		c.camera.Target.Y = minY
	}
}

// GetCamera returns the current camera state
func (c *Camera) GetCamera() rl.Camera2D {
	return c.camera
}
