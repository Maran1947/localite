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

func PushToFile(filePath string, keys []string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, key := range keys {
		if _, err := file.WriteString(key + "\n"); err != nil {
			return err
		}
	}
	return nil
}