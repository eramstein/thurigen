package engine

type TerrainType int

const (
	Dirt TerrainType = iota
	Rock
	Sand
	Water
)

type SurfaceType int

const (
	Grass SurfaceType = iota
	WoodSurface
)

type VolumeType int

const (
	RockVolume VolumeType = iota
	WoodVolume
)
