package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/maran1947/localite/internal/utils"
)

const configDirName = ".localite"

type Config struct {
	GeminiApiKey string `json:"GEMINI_API_KEY"`
}

func LoadConfig() (*Config, error) {
	configFilePath, err:= utils.GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}

		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	return &config, err
}

func SaveConfig(key, value string) error {
	config, err := LoadConfig()
	if err != nil {
		config = &Config{}
	}

	switch key {
	case "GEMINI_API_KEY":
		config.GeminiApiKey = value
	default:
		return fmt.Errorf("unknown configuration key: %s", key)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	configFilePath, err:= utils.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("unknown configuration key: %s", err)
	}
	return os.WriteFile(configFilePath, data, 0644)
}
