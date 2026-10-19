package config

import (
	"encoding/json"
	"os"
	"fmt"
)

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	
	}

	cfgFile := home + configFileName

	return cfgFile, nil
}

func Read() (Config, error) {
	cfg := Config{}

	cfgFilePath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	cfgFile, err := os.Open(cfgFilePath)
	if err != nil {
		return cfg, err
	}

	defer cfgFile.Close()

	err = json.NewDecoder(cfgFile).Decode(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg Config) error {
	jsonData, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}
	
	cfgFilePath, err := getConfigFilePath()
		if err != nil {
		return err
	}


	err = os.WriteFile(cfgFilePath, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("JSON Config saved successfully!\n")

	return nil
}
