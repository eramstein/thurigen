package ng

import "eramstein/thurigen/pkg/config"

func (sim *Simulation) AddStructure(structure Structure) {
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
			}
		}
	}
}
