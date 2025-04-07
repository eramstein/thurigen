package ui

import "eramstein/thurigen/pkg/config"

func (r *Renderer) ToggleTileSelection(x, y int) {
	if r.uiState.SelectedTile != nil && x == r.uiState.SelectedTile[0] && y == r.uiState.SelectedTile[1] {
		r.uiState.SelectedTile = nil
	} else if x >= 0 && x < config.RegionSize && y >= 0 && y < config.RegionSize {
		r.uiState.SelectedTile = &[2]int{x, y}
	}
}
