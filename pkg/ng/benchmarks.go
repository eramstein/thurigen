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
		sim.ScanForItem(position, 10, itemType)
		sim.GetClosestItem(position, itemType)
		sim.ScanForItemFullRegion(position, itemType)
	}

	// Benchmark ScanForItem
	startTime := time.Now()
	for i := 0; i < iterations; i++ {
		sim.ScanForItem(position, 30, itemType)
	}
	scanTime := time.Since(startTime)

	// Benchmark GetClosestItem
	startTime = time.Now()
	for i := 0; i < iterations; i++ {
		sim.GetClosestItem(position, itemType)
	}
	closestTime := time.Since(startTime)

	// Benchmark ScanForItemFullRegion
	startTime = time.Now()
	for i := 0; i < iterations; i++ {
		sim.ScanForItemFullRegion(position, itemType)
	}
	fullRegionTime := time.Since(startTime)

	// Print results
	fmt.Printf("Benchmark results for %d iterations:\n", iterations)
	fmt.Printf("ScanForItem: %v\n", scanTime)
	fmt.Printf("GetClosestItem: %v\n", closestTime)
	fmt.Printf("ScanForItemFullRegion: %v\n", fullRegionTime)
}
