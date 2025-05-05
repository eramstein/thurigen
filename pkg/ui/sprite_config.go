package ui

import (
	"eramstein/thurigen/pkg/ng"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// SpriteSheetConfig defines the configuration for a sprite sheet
type SpriteSheetConfig struct {
	Name     string
	Path     string
	TileSize int32
	Sprites  map[uint64]rl.Rectangle
}

var structureToSpriteSheet = map[ng.StructureType]SpriteSheetConfig{
	ng.Plant: {
		Name:     "plants",
		Path:     "assets/images/world/plants.png",
		TileSize: 32,
		Sprites: map[uint64]rl.Rectangle{
			1: rl.NewRectangle(0, 0, 32, 32),
			2: rl.NewRectangle(32, 0, 32, 32),
			3: rl.NewRectangle(64, 0, 32, 32),
		},
	},
}

var buildingTypeToSpriteSheet = map[ng.BuildingVariant]SpriteSheetConfig{
	ng.Wall: {
		Name:     "walls",
		Path:     "assets/images/world/walls.png",
		TileSize: 32,
		Sprites: map[uint64]rl.Rectangle{
			uint64(ng.WoodMaterial): rl.NewRectangle(0, 0, 32, 32),
			uint64(ng.RockMaterial): rl.NewRectangle(32, 0, 32, 32),
		},
	},
	ng.Door: {
		Name:     "doors",
		Path:     "assets/images/world/furniture.png",
		TileSize: 32,
		Sprites: map[uint64]rl.Rectangle{
			uint64(ng.WoodMaterial): rl.NewRectangle(0, 0, 32, 32),
			uint64(ng.RockMaterial): rl.NewRectangle(0, 0, 32, 32),
		},
	},
	ng.Window: {
		Name:     "windows",
		Path:     "assets/images/world/furniture.png",
		TileSize: 32,
		Sprites: map[uint64]rl.Rectangle{
			uint64(ng.WoodMaterial): rl.NewRectangle(32, 0, 32, 32),
			uint64(ng.RockMaterial): rl.NewRectangle(32, 0, 32, 32),
		},
	},
}

var charactersSpriteSheet = SpriteSheetConfig{
	Name:     "characters",
	Path:     "assets/images/characters/characters.png",
	TileSize: 32,
	Sprites: map[uint64]rl.Rectangle{
		0: rl.NewRectangle(0, 0, 32, 32),
		1: rl.NewRectangle(32, 0, 32, 32),
		2: rl.NewRectangle(64, 0, 32, 32),
		3: rl.NewRectangle(0, 32, 32, 32),
		4: rl.NewRectangle(32, 32, 32, 32),
		5: rl.NewRectangle(64, 32, 32, 32),
		6: rl.NewRectangle(0, 64, 32, 32),
		7: rl.NewRectangle(32, 64, 32, 32),
		8: rl.NewRectangle(64, 64, 32, 32),
	},
}

var taskIconsSpriteSheet = SpriteSheetConfig{
	Name:     "task_icons",
	Path:     "assets/images/ui/task_icons.png",
	TileSize: 32,
	Sprites: map[uint64]rl.Rectangle{
		uint64(ng.Sleep): rl.NewRectangle(0, 0, 32, 32),
		uint64(ng.Eat):   rl.NewRectangle(32, 0, 32, 32),
		uint64(ng.Drink): rl.NewRectangle(64, 0, 32, 32),
		uint64(ng.Build): rl.NewRectangle(0, 32, 32, 32),
		uint64(ng.Chop):  rl.NewRectangle(64, 32, 32, 32),
	},
}

var terrainSpriteSheet = SpriteSheetConfig{
	Name:     "terrain",
	Path:     "assets/images/world/terrain.png",
	TileSize: 32,
	Sprites: map[uint64]rl.Rectangle{
		0: rl.NewRectangle(0, 0, 32, 32),
		1: rl.NewRectangle(32, 0, 32, 32),
		2: rl.NewRectangle(64, 0, 32, 32),
		3: rl.NewRectangle(0, 32, 32, 32),
		4: rl.NewRectangle(32, 32, 32, 32),
		5: rl.NewRectangle(64, 32, 32, 32),
		6: rl.NewRectangle(0, 64, 32, 32),
		7: rl.NewRectangle(32, 64, 32, 32),
		8: rl.NewRectangle(64, 64, 32, 32),
	},
}

var itemToSpriteSheet = map[ng.ItemType]SpriteSheetConfig{
	ng.Material: {
		Name:     "materials",
		Path:     "assets/images/world/materials.png",
		TileSize: 32,
		Sprites: map[uint64]rl.Rectangle{
			uint64(ng.WoodMaterial): rl.NewRectangle(0, 0, 32, 32),
			uint64(ng.RockMaterial): rl.NewRectangle(32, 0, 32, 32),
		},
	},
}
