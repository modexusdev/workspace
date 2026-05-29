package config

import (
	"encoding/json"
	"os"

	"github.com/modexusdev/workspace/internal/models"
)

func LoadConfig() (*models.Config, error) {
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	var configData models.Config

	if err := json.Unmarshal(data, &configData); err != nil {
		return nil, err
	}

	return &configData, nil
}

func SaveConfig(configData *models.Config) error {
	updatedData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigPath, updatedData, 0644)
}
