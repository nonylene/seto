package common

import (
	"errors"
	"os"
	"path"
)

func GetDefaultConfigPath(subPath string) (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home := os.Getenv("HOME")
		if home == "" {
			return "", errors.New("environment variable XDG_CONFIG_HOME or HOME must be set")
		}
		configHome = path.Join(home, ".config")
	}

	return path.Join(configHome, subPath), nil
}
