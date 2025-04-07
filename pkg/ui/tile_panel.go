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
	r.RenderText(coordText, panelX+10, yOffset)
	yOffset += lineHeight

	// Terrain type
	terrainText := fmt.Sprintf("%v", tile.Terrain)
	r.RenderText(terrainText, panelX+10, yOffset)
	yOffset += lineHeight

	// Surface type
	if tile.Surface != 0 {
		surfaceText := fmt.Sprintf("%v", tile.Surface)
		r.RenderText(surfaceText, panelX+10, yOffset)
		yOffset += lineHeight
	}

	// Volume type
	if tile.Volume != 0 {
		volumeText := fmt.Sprintf("%v", tile.Volume)
		r.RenderText(volumeText, panelX+10, yOffset)
		yOffset += lineHeight
	}

	// Structure information
	if tile.Occupation != nil && tile.Occupation.Structure != nil {
		structure := tile.Occupation.Structure
		base := structure.GetStructure()
		config := ng.GetStructureConfig(base.Type, base.Variant)
		structureText := fmt.Sprintf(config.Name)
		r.RenderText(structureText, panelX+10, yOffset)
		yOffset += lineHeight

		// If it's a plant, show growth stage
		if plant, ok := structure.(*ng.PlantStructure); ok {
			growthText := fmt.Sprintf("Growth: %d%%", plant.GrowthStage)
			r.RenderText(growthText, panelX+10, yOffset)
			yOffset += lineHeight

			prodText := fmt.Sprintf("Production: %d%%", plant.ProductionStage)
			r.RenderText(prodText, panelX+10, yOffset)
			yOffset += lineHeight
		}
	}

	// Items information
	if len(tile.Items) > 0 {
		itemsText := fmt.Sprintf("Items: %d", len(tile.Items))
		r.RenderText(itemsText, panelX+10, yOffset)
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
			r.RenderText(itemText, panelX+10, yOffset)
			yOffset += lineHeight
		}
	}
}
