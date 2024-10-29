package utils

import (
	"os"
	"path/filepath"
)

const configDirName = ".localite"
const configFileName = "config.json"

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, configDirName)
	if err := os.MkdirAll(configDir, 0700); err != nil { 
		return "", err
	}

	configFilePath := filepath.Join(configDir, configFileName)
	return configFilePath, nil
}