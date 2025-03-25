package engine

import (
	"fmt"
	"time"
)

// NewSimulation creates a new Simulation instance
func NewSimulation() *Simulation {
	return &Simulation{
		Time:  0,
		Speed: 60,
	}
}

// Runs the simulation: events, character actions, etc.
func (sim *Simulation) Run() {
	speed := time.Millisecond * 1000 // Example speed variable
	ticker := time.NewTicker(speed)
	defer ticker.Stop()

	go func() {
		fmt.Println("Routiner starts", sim.Time)
		for range ticker.C {
			fmt.Println("Routiner tick", sim.Time)
			sim.Update()
		}
	}()
}

func (sim *Simulation) Update() {
	sim.Time++
}
