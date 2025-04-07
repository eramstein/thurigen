package ng

import (
	"encoding/json"
	"fmt"
	"os"
)

type StructureConfig struct {
	Name string
	Structure
}

// structureConfigs maps structure types and variants to their configurations
var structureConfigs = make(map[StructureType]map[int]StructureConfig)

func LoadStructuresConfigs() error {
	structureConfigs[Plant] = make(map[int]StructureConfig)
	if err := LoadPlantConfigs(); err != nil {
		fmt.Println("Error LoadPlantConfigs:", err)
		return err
	}
	return nil
}

// LoadPlantConfigs reads plant configurations from the JSON file
func LoadPlantConfigs() error {
	// Read the JSON file
	data, err := os.ReadFile("data/structures_plants.json")
	if err != nil {
		fmt.Println("Error reading plants.json:", err)
		return err
	}

	// Parse the JSON data
	var rawConfig struct {
		Plant map[int]struct {
			Name           string `json:"name"`
			GrowthRate     int    `json:"growthRate"`
			ProductionRate int    `json:"productionRate"`
			Produces       struct {
				Type    int `json:"type"`
				Variant int `json:"variant"`
			} `json:"produces"`
			MoveCost float64 `json:"moveCost"`
		} `json:"plant"`
	}

	if err := json.Unmarshal(data, &rawConfig); err != nil {
		fmt.Println("Error unmarshalling rawConfig:", err)
		return err
	}

	// Convert the raw config to our internal format
	for plantTypeInt, plantData := range rawConfig.Plant {
		structureConfigs[Plant][plantTypeInt] = StructureConfig{
			Name: plantData.Name,
			Structure: &PlantStructure{
				BaseStructure: BaseStructure{
					Type:     Plant,
					Variant:  plantTypeInt,
					Size:     [2]int{1, 1},
					Rotation: 0,
					MoveCost: MoveCost(plantData.MoveCost),
				},
				GrowthRate:     plantData.GrowthRate,
				ProductionRate: plantData.ProductionRate,
				Produces: PlantProduction{
					Type:    ItemType(plantData.Produces.Type),
					Variant: plantData.Produces.Variant,
				},
			},
		}
	}

	return nil
}

// GetStructureConfig returns the configuration for a specific structure type and variant
func GetStructureConfig(structureType StructureType, variant int) StructureConfig {
	if typeConfigs, exists := structureConfigs[structureType]; exists {
		if config, exists := typeConfigs[variant]; exists {
			return config
		}
	}
	return StructureConfig{
		Name: "Unknown StructureConfig",
	}
}
