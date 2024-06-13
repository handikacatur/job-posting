package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func InitConfig() Config {
	appConfig := Config{}
	configFilePath := os.Getenv("APP_ENV")
	if configFilePath == "" {
		configFilePath = ".config.yaml"
	}

	if err := readConfig(&appConfig, configFilePath); err != nil {
		panic("Config not found")
	}

	return appConfig
}

func readConfig(ac *Config, configPath string) error {
	configPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(yamlFile, ac); err != nil {
		return err
	}

	return nil
}
