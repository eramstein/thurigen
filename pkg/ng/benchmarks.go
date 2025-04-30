package ng

import (
	"fmt"
	"time"
)

// A temporary function containing whatever function we currently want to benchmark
func (sim *Simulation) Benchmark() {
	fmt.Println(len(sim.World[0].Plants))
	// sim.BenchmarkItemSearch(0, [2]int{32, 32}, Food, 1000)
}

// BenchmarkItemSearch compares the performance of different item search methods
func (sim *Simulation) BenchmarkItemSearch(position Position, itemType ItemType, iterations int) {
	// Warm up
	for i := 0; i < 10; i++ {
		sim.ScanForItem(position, 10, itemType, -1, false)
	}

	// Benchmark ScanForItem
	startTime := time.Now()
	for i := 0; i < iterations; i++ {
		sim.ScanForItem(position, 30, itemType, -1, false)
	}
	scanTime := time.Since(startTime)

	// Print results
	fmt.Printf("Benchmark results for %d iterations:\n", iterations)
	fmt.Printf("ScanForItem: %v\n", scanTime)
}
