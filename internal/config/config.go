package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to get home dir: %w", err)
	}

	configFilePath := homeDir + "/.gatorconfig.json"

	return configFilePath, nil
}

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("Couldn't read config from file: %w", err)
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to unmarshal config json: %w", err)
	}

	return config, nil
}

func (cfg *Config) SetUser(user string) (error) {
	cfg.CurrentUserName = user

	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Failed to marshal config to json: %w", err)
	}

	err = os.WriteFile(configFile, data, 0755)
	if err != nil {
		return fmt.Errorf("Couldn't write config to file: %w", err)
	}

	return nil
}
