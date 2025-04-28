package ui

import (
	"eramstein/thurigen/pkg/ng"
)

type Model struct {
	Ticker                     int
	DisplayedRegion            int
	SelectedTile               *[2]int
	SelectedCharacter          *ng.Character
	PreviousCharacterPositions map[uint64]ng.Position
}
