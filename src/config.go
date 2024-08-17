package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

type Config struct {
	SocketPath string `json:"socketPath"`
}

const configFileSubPath = "seto/config.json"

func ParseConfig(configPath string) (*Config, error) {
	if configPath == "" {
		configHome := os.Getenv("XDG_CONFIG_HOME")
		if configHome == "" {
			home := os.Getenv("HOME")
			if home == "" {
				return nil, errors.New("environment variable XDG_CONFIG_HOME or HOME must be set")
			}
			configHome = path.Join(home, ".config")
		}

		configPath = path.Join(configHome, configFileSubPath)
	}

	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file (%s): %w", configPath, err)
	}

	var config *Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}

	return config, nil
}
