package engine

// Simulation represents the main simulation state
type Simulation struct {
	Paused bool
	Speed  int // how many frames until next sim update
	Time   int // in minutes since the start of the simulation
}
