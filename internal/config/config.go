package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	// Just a placeholder property for now.
	DebugMode bool `toml:"debug_mode"`

	// The keybindings for the app
	Keys *Keys `toml:"keys"`
}

func GetConfig(path string) (*Config, error) {
	config := Config{
		DebugMode: true,
	}

	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &config, nil
		}
		return nil, fmt.Errorf("decoding toml file: %w", err)
	}

	return &config, nil
}
