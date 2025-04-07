package ng

import (
	"encoding/json"
	"fmt"
	"os"
)

type ItemConfig struct {
	Name string
	Item
}

// itemsConfigs maps item types and variants to their configurations
var itemsConfigs = make(map[ItemType]map[int]ItemConfig)

func LoadItemsConfigs() error {
	itemsConfigs[Food] = make(map[int]ItemConfig)
	itemsConfigs[Material] = make(map[int]ItemConfig)
	if err := LoadFoodConfigs(); err != nil {
		return err
	}
	if err := LoadMaterialConfigs(); err != nil {
		return err
	}
	return nil
}

// LoadFoodConfigs reads food configurations from the JSON file
func LoadFoodConfigs() error {
	// Read the JSON file
	data, err := os.ReadFile("data/items_food.json")
	if err != nil {
		fmt.Println("Error reading items_food.json:", err)
		return err
	}

	// Parse the JSON data
	var rawConfig struct {
		Food map[int]struct {
			Name      string `json:"name"`
			Nutrition int    `json:"nutrition"`
		} `json:"food"`
	}

	if err := json.Unmarshal(data, &rawConfig); err != nil {
		fmt.Println("Error unmarshalling rawConfig:", err)
		return err
	}

	// Convert the raw config to our internal format
	for itemTypeInt, itemData := range rawConfig.Food {
		itemsConfigs[Food][itemTypeInt] = ItemConfig{
			Name: itemData.Name,
			Item: &FoodItem{
				BaseItem: BaseItem{
					Type:    ItemType(itemTypeInt),
					Variant: itemTypeInt,
				},
				Nutrition: itemData.Nutrition,
			},
		}
	}

	return nil
}

// LoadMaterialConfigs reads material configurations from the JSON file
func LoadMaterialConfigs() error {
	// Read the JSON file
	data, err := os.ReadFile("data/items_material.json")
	if err != nil {
		fmt.Println("Error reading items_material.json:", err)
		return err
	}

	// Parse the JSON data
	var rawConfig struct {
		Material map[int]struct {
			Name         string `json:"name"`
			MaterialType int    `json:"materialType"`
		} `json:"material"`
	}

	if err := json.Unmarshal(data, &rawConfig); err != nil {
		fmt.Println("Error unmarshalling rawConfig:", err)
		return err
	}

	// Convert the raw config to our internal format
	for itemTypeInt, itemData := range rawConfig.Material {
		itemsConfigs[Material][itemTypeInt] = ItemConfig{
			Name: itemData.Name,
			Item: &MaterialItem{
				BaseItem: BaseItem{
					Type:    ItemType(itemTypeInt),
					Variant: itemTypeInt,
				},
				MaterialType: MaterialType(itemData.MaterialType),
			},
		}
	}

	return nil
}

// GetItemConfig returns the configuration for a specific item type and variant
func GetItemConfig(itemType ItemType, variant int) ItemConfig {
	if typeConfigs, exists := itemsConfigs[itemType]; exists {
		if config, exists := typeConfigs[variant]; exists {
			return config
		}
	}
	return ItemConfig{
		Name: "Unknown Item",
	}
}
