package config

import (
	"fmt"
	"os"
	"path"
	"tinytrail/internal/environment"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	DatabaseURL string `yaml:"database"`
}

func LoadConfig() (*AppConfig, error) {
	configDir := environment.GetConfigDir()
	if configDir == "" {
		return nil, fmt.Errorf("CONFIG_DIR env variable is not set")
	}

	configPath := path.Join(configDir, "config.yaml")
	var config AppConfig

	configFileContent, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(configFileContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}
