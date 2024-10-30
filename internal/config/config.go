package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/maran1947/localite/internal/utils"
)

type LocaliteConfig map[string]string

func (c *LocaliteConfig) SetConfigValue(key, value string) {
	(*c)[key] = value
}

func (c *LocaliteConfig) GetConfigValue(key string) (string, bool) {
	value, exists := (*c)[key]
	return value, exists
}

func (c *LocaliteConfig) DeleteConfigValue(key string) {
	delete(*c, key)
}

func LoadConfig() (*LocaliteConfig, error) {
	configFilePath, err := utils.GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &LocaliteConfig{}, nil
		}

		return nil, err
	}

	var config LocaliteConfig
	err = json.Unmarshal(data, &config)
	return &config, err
}

func SaveConfig(key, value string) error {
	localiteConfig, err := LoadConfig()
	if err != nil {
		localiteConfig = &LocaliteConfig{}
	}

	localiteConfig.SetConfigValue(key, value)

	data, err := json.MarshalIndent(localiteConfig, "", "  ")
	if err != nil {
		return err
	}

	configFilePath, err := utils.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("unknown configuration key: %s", err)
	}
	return os.WriteFile(configFilePath, data, 0644)
}

func UpdateConfig(localiteConfig LocaliteConfig) error {
	data, err := json.MarshalIndent(localiteConfig, "", "  ")
	if err != nil {
		return err
	}

	configFilePath, err := utils.GetConfigFilePath()
	return os.WriteFile(configFilePath, data, 0644)
}
