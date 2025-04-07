package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DisplayTileSidePanel shows information about the selected tile
func (r *Renderer) DisplayTileSidePanel() {
	if r.uiState.SelectedTile == nil {
		return
	}
	// Panel dimensions
	panelWidth := config.TileSidePanelWidth
	panelX := r.screenWidth - panelWidth
	panelY := 0
	panelHeight := r.screenHeight

	// Draw panel background
	rl.DrawRectangle(int32(panelX), int32(panelY), int32(panelWidth), int32(panelHeight), rl.NewColor(240, 240, 240, 255))
	rl.DrawRectangleLines(int32(panelX), int32(panelY), int32(panelWidth), int32(panelHeight), rl.Gray)

	// Get selected tile data
	tileX, tileY := (*r.uiState.SelectedTile)[0], (*r.uiState.SelectedTile)[1]
	tile := r.uiState.DisplayedRegion.Tiles[tileX][tileY]

	// Draw tile information
	yOffset := 20
	lineHeight := 25

	// Tile coordinates
	coordText := fmt.Sprintf("Tile: (%d, %d)", tileX, tileY)
	rl.DrawText(coordText, int32(panelX+10), int32(yOffset), 20, rl.Black)
	yOffset += lineHeight

	// Terrain type
	terrainText := fmt.Sprintf("Terrain: %v", tile.Terrain)
	rl.DrawText(terrainText, int32(panelX+10), int32(yOffset), 20, rl.Black)
	yOffset += lineHeight

	// Surface type
	if tile.Surface != 0 {
		surfaceText := fmt.Sprintf("Surface: %v", tile.Surface)
		rl.DrawText(surfaceText, int32(panelX+10), int32(yOffset), 20, rl.Black)
		yOffset += lineHeight
	}

	// Volume type
	if tile.Volume != 0 {
		volumeText := fmt.Sprintf("Volume: %v", tile.Volume)
		rl.DrawText(volumeText, int32(panelX+10), int32(yOffset), 20, rl.Black)
		yOffset += lineHeight
	}

	// Structure information
	if tile.Occupation != nil && tile.Occupation.Structure != nil {
		structure := tile.Occupation.Structure
		base := structure.GetStructure()
		structureText := fmt.Sprintf("Structure: %v (Variant: %d)", base.Type, base.Variant)
		rl.DrawText(structureText, int32(panelX+10), int32(yOffset), 20, rl.Black)
		yOffset += lineHeight

		// If it's a plant, show growth stage
		if plant, ok := structure.(*ng.PlantStructure); ok {
			growthText := fmt.Sprintf("Growth: %d%%", plant.GrowthStage)
			rl.DrawText(growthText, int32(panelX+10), int32(yOffset), 20, rl.Black)
			yOffset += lineHeight
		}
	}

	// Items information
	if len(tile.Items) > 0 {
		itemsText := fmt.Sprintf("Items: %d", len(tile.Items))
		rl.DrawText(itemsText, int32(panelX+10), int32(yOffset), 20, rl.Black)
		yOffset += lineHeight

		// List each item type and count
		itemCounts := make(map[string]int)
		for _, item := range tile.Items {
			baseItem := (*item).GetItem()
			config := ng.GetItemConfig(baseItem.Type, baseItem.Variant)
			itemCounts[config.Name]++
		}
		for itemName, count := range itemCounts {
			itemText := fmt.Sprintf("  - %s: %d", itemName, count)
			rl.DrawText(itemText, int32(panelX+10), int32(yOffset), 20, rl.Black)
			yOffset += lineHeight
		}
	}
}
