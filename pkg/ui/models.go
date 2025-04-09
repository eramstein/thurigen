package ui

import "eramstein/thurigen/pkg/ng"

type Model struct {
	DisplayedRegion *ng.Region
	SelectedTile    *[2]int
	Simulation      *ng.Simulation
}
