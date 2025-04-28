package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"
)

func (r *Renderer) ToggleTileSelection(x, y int) {
	if r.UiState.SelectedTile != nil && x == r.UiState.SelectedTile[0] && y == r.UiState.SelectedTile[1] {
		r.UiState.SelectedTile = nil
	} else if x >= 0 && x < config.RegionSize && y >= 0 && y < config.RegionSize {
		r.UiState.SelectedTile = &[2]int{x, y}
	}
}

func (r *Renderer) ToggleCharacterSelection(character *ng.Character) {
	if r.UiState.SelectedCharacter == character {
		r.UiState.SelectedCharacter = nil
	} else {
		r.UiState.SelectedCharacter = character
		// Center camera on the selected character
		r.camera.CenterOnPosition(float32(character.Position.X), float32(character.Position.Y))
	}
}

func (r *Renderer) CancelTileSelection() {
	r.UiState.SelectedTile = nil
}

func (r *Renderer) CancelCharacterSelection() {
	r.UiState.SelectedCharacter = nil
}

func (r *Renderer) UpdatePreviousCharacterPositions(sim *ng.Simulation) {
	for _, character := range sim.Characters {
		r.UiState.PreviousCharacterPositions[character.ID] = character.Position
	}
}

func (r *Renderer) InitPreviousCharacterPositions(sim *ng.Simulation) {
	r.UiState.PreviousCharacterPositions = make(map[uint64]ng.Position)
	r.UpdatePreviousCharacterPositions(sim)
}
