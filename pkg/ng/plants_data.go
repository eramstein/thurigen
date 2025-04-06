package ng

import (
	"encoding/json"
	"fmt"
	"os"
)

type PlantConfig struct {
	Name string
	PlantStructure
}

// plantConfigs maps plant types and variants to their configurations
var plantConfigs = make(map[PlantType]map[int]PlantConfig)

// LoadPlantConfigs reads plant configurations from the JSON file
func LoadPlantConfigs() error {
	// Read the JSON file
	data, err := os.ReadFile("data/plants.json")
	if err != nil {
		fmt.Println("Error reading plants.json:", err)
		return err
	}

	// Parse the JSON data
	var rawConfig struct {
		Plants map[int]struct {
			Variants map[int]struct {
				Name           string `json:"name"`
				GrowthRate     int    `json:"growthRate"`
				ProductionRate int    `json:"productionRate"`
				Produces       struct {
					Type    int `json:"type"`
					Variant int `json:"variant"`
				} `json:"produces"`
				MoveCost float64 `json:"moveCost"`
			} `json:"variants"`
		} `json:"plants"`
	}

	if err := json.Unmarshal(data, &rawConfig); err != nil {
		fmt.Println("Error unmarshalling rawConfig:", err)
		return err
	}

	// Convert the raw config to our internal format
	for plantTypeInt, plantData := range rawConfig.Plants {
		plantConfigs[PlantType(plantTypeInt)] = make(map[int]PlantConfig)
		for variantInt, variantData := range plantData.Variants {
			plantConfigs[PlantType(plantTypeInt)][variantInt] = PlantConfig{
				Name: variantData.Name,
				PlantStructure: PlantStructure{
					BaseStructure: BaseStructure{
						Type:     Plant,
						Variant:  variantInt,
						Size:     [2]int{1, 1},
						Rotation: 0,
						MoveCost: MoveCost(variantData.MoveCost),
					},
					GrowthRate:     variantData.GrowthRate,
					ProductionRate: variantData.ProductionRate,
					Produces: PlantProduction{
						Type:    ItemType(variantData.Produces.Type),
						Variant: variantData.Produces.Variant,
					},
				},
			}
		}
	}

	return nil
}

// GetPlantConfig returns the configuration for a specific plant type and variant
func GetPlantConfig(plantType PlantType, variant int) PlantConfig {
	if typeConfigs, exists := plantConfigs[plantType]; exists {
		if config, exists := typeConfigs[variant]; exists {
			return config
		}
	}
	return PlantConfig{
		Name: "Unknown Plant",
	}
}
