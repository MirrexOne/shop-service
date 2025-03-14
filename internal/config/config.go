package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"log"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}
	HTTP struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"level"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error occurred while reading config file")
	}

	if err := yaml.Unmarshal(file, cfg); err != nil {
		log.Fatal("Error occurred while unmarshal config")
	}

	return cfg, nil
}
