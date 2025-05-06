package ng

import "eramstein/thurigen/pkg/config"

func (sim *Simulation) AddStructureOccupation(structure Structure) {
	base := structure.GetStructure()

	// Update the tiles it occupies
	for dx := 0; dx < base.Size[0]; dx++ {
		for dy := 0; dy < base.Size[1]; dy++ {
			tileX := base.Position.X + dx
			tileY := base.Position.Y + dy

			if tileX >= 0 && tileX < config.RegionSize && tileY >= 0 && tileY < config.RegionSize {
				tile := &sim.World[base.Position.Region].Tiles[tileX][tileY]
				tile.Occupation = &TileOccupation{
					Structure:  structure,          // Assign the structure to the tile
					IsMainTile: dx == 0 && dy == 0, // Mark the main tile
				}
				// adjust tile move cost based on structure type
				tile.MoveCost = base.MoveCost
			}
		}
	}
}

func (sim *Simulation) RemoveStructureOccupation(structure Structure) {
	base := structure.GetStructure()

	// Update the tiles it occupies
	for dx := 0; dx < base.Size[0]; dx++ {
		for dy := 0; dy < base.Size[1]; dy++ {
			tileX := base.Position.X + dx
			tileY := base.Position.Y + dy
			if tileX >= 0 && tileX < config.RegionSize && tileY >= 0 && tileY < config.RegionSize {
				sim.World[base.Position.Region].Tiles[tileX][tileY].Occupation = nil
			}
		}
	}
}

// removeFromSlice removes a structure from a slice of structures by ID
func removeFromSlice[T Structure](slice []T, structure T) []T {
	for i, s := range slice {
		if s.GetStructure().ID == structure.GetStructure().ID {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
