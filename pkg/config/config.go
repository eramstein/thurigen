package config

const (
	WindowWidth        = 1344
	WindowHeight       = 960
	FPS                = 60
	RegionSize         = 64
	TilePixelSize      = 32
	TileSidePanelWidth = 350
	BaseFontSize       = 16

	CharacterNeedsUpdateInterval     = 1
	CharacterObjectiveUpdateInterval = 5
	CharacterTaskUpdateInterval      = 1

	// TODO: restore to 1 after testing buildign
	NeedFoodTick  = 0
	NeedWaterTick = 0
	NeedSleepTick = 0

	// Character portrait UI settings
	PortraitSize    = TilePixelSize * 2 // Twice the size of regular character icons
	PortraitSpacing = 10                // Space between portraits
	PortraitStartX  = 300               // Start X position for portraits
	PortraitStartY  = 10                // Start Y position for portraits
)
