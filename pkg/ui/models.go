package ui

import "eramstein/thurigen/pkg/ng"

type Model struct {
	DisplayedRegion   int
	SelectedTile      *[2]int
	SelectedCharacter *ng.Character
}
