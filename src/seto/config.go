package seto

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	util "github.com/nonylene/seto/src/common"
)

type Config struct {
	SocketPath         string   `json:"socketPath"`
	BrowserCommand     []string `json:"browserCommand"`
	CodeRemoteArgument string   `json:"codeRemoteArgument"`
}

func (c *Config) validate() error {
	if c.SocketPath == "" {
		return errors.New("socketPath must be defined")
	}

	if len(c.BrowserCommand) <= 0 {
		return errors.New("browserCommand must be defined")
	}

	if c.CodeRemoteArgument == "" {
		return errors.New("codeRemoteArgument must be defined")
	}

	return nil
}

const configFileSubPath = "seto/config.json"

func ParseConfig(configPath string) (*Config, error) {
	var err error
	if configPath == "" {
		configPath, err = util.GetDefaultConfigPath(configFileSubPath)
		if err != nil {
			return nil, err
		}
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

	err = config.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return config, nil
}
