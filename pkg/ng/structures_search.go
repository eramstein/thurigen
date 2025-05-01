package ng

import (
	"math"
)

func (sim *Simulation) FindClosestPlant(position Position, structureType StructureType, variant int, plantType PlantType) *PlantStructure {
	region := sim.World[position.Region]
	var closestPlant *PlantStructure
	minDistance := math.MaxFloat64

	for _, plant := range region.Plants {
		// Skip if plant doesn't match criteria
		if plant.Type != structureType || (variant != -1 && plant.Variant != variant) || plant.PlantType != plantType {
			continue
		}

		// Calculate distance to plant
		dx := float64(plant.Position.X - position.X)
		dy := float64(plant.Position.Y - position.Y)
		distance := math.Sqrt(dx*dx + dy*dy)

		// Update closest plant if this one is closer
		if distance < minDistance {
			minDistance = distance
			closestPlant = plant
		}
	}

	return closestPlant
}
