package ng

import "eramstein/thurigen/pkg/config"

// Simulation represents the main simulation state
type Simulation struct {
	Paused bool
	Speed  int // how many frames until next sim update
	Time   int // in minutes since the start of the simulation
	World  []*Region
}

type Region struct {
	Tiles [config.RegionSize][config.RegionSize]Tile
}

type Tile struct {
	Terrain TerrainType
	Surface SurfaceType
	Volume  VolumeType
}
