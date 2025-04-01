package ng

import "eramstein/thurigen/pkg/config"

func (sim *Simulation) AddStructure(structure BaseStructure) {
	// Add the structure to the Structures slice
	sim.Structures = append(sim.Structures, &structure)

	// Update the tiles it occupies
	for dx := 0; dx < structure.Size[0]; dx++ {
		for dy := 0; dy < structure.Size[1]; dy++ {
			tileX := structure.Position[0] + dx
			tileY := structure.Position[1] + dy

			if tileX >= 0 && tileX < config.RegionSize && tileY >= 0 && tileY < config.RegionSize {
				tile := &sim.World[structure.Region].Tiles[tileY][tileX]
				tile.Occupation = &TileOccupation{
					Structure:  &structure,         // Assign the structure to the tile
					IsMainTile: dx == 0 && dy == 0, // Mark the main tile
				}
			}
		}
	}
}
