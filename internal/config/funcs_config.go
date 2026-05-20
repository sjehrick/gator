package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	cfg := Config{}

	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}

	dat, err := os.ReadFile(home + configFileName)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
