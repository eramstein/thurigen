package ui

import (
	"eramstein/thurigen/pkg/config"
	"eramstein/thurigen/pkg/ng"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DisplayTileSidePanel shows information about the selected tile
func (r *Renderer) DisplayTileSidePanel(sim *ng.Simulation) {
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
	tile := sim.World[r.uiState.DisplayedRegion].Tiles[tileX][tileY]

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
			config := ng.GetItemConfig(item.Type, item.Variant)
			itemCounts[config.Name]++
		}
		for itemName, count := range itemCounts {
			itemText := fmt.Sprintf("  - %s: %d", itemName, count)
			r.RenderText(itemText, panelX+10, yOffset)
			yOffset += lineHeight
		}
	}

	// Check for characters at this position
	if tile.Character != nil {
		character := tile.Character
		// separator
		separatorText := "--------------------------------"
		r.RenderText(separatorText, panelX+10, yOffset)
		yOffset += lineHeight

		// Character name and ID
		charText := fmt.Sprintf("Character: %s", character.Name)
		r.RenderText(charText, panelX+10, yOffset)
		yOffset += lineHeight

		// Character needs
		needsText := "Needs:"
		r.RenderText(needsText, panelX+10, yOffset)
		yOffset += lineHeight
		r.RenderText(fmt.Sprintf("  Food: %d%%", character.Needs.Food), panelX+10, yOffset)
		yOffset += lineHeight
		r.RenderText(fmt.Sprintf("  Water: %d%%", character.Needs.Water), panelX+10, yOffset)
		yOffset += lineHeight
		r.RenderText(fmt.Sprintf("  Sleep: %d%%", character.Needs.Sleep), panelX+10, yOffset)
		yOffset += lineHeight

		// Current tasks
		if len(character.Tasks) > 0 {
			tasksText := "Current Tasks:"
			r.RenderText(tasksText, panelX+10, yOffset)
			yOffset += lineHeight
			for _, task := range character.Tasks {
				taskText := fmt.Sprintf("  - %v", task.Type)
				r.RenderText(taskText, panelX+10, yOffset)
				yOffset += lineHeight
			}
		}

		// Objectives
		if len(character.Objectives) > 0 {
			objectivesText := "Objectives:"
			r.RenderText(objectivesText, panelX+10, yOffset)
			yOffset += lineHeight
			for _, objective := range character.Objectives {
				objectiveText := fmt.Sprintf("  - %v", objective.Type)
				r.RenderText(objectiveText, panelX+10, yOffset)
				yOffset += lineHeight
			}
		}

		// Inventory
		if len(character.Inventory) > 0 {
			inventoryText := "Inventory:"
			r.RenderText(inventoryText, panelX+10, yOffset)
			yOffset += lineHeight
			itemCounts := make(map[string]int)
			for _, item := range character.Inventory {
				config := ng.GetItemConfig(item.Type, item.Variant)
				itemCounts[config.Name]++
			}
			for itemName, count := range itemCounts {
				itemText := fmt.Sprintf("  - %s: %d", itemName, count)
				r.RenderText(itemText, panelX+10, yOffset)
				yOffset += lineHeight
			}
		}
	}
}
