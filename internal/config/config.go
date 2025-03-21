package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    `yaml:"app"`
		Server `yaml:"server"`
		Log    `yaml:"log"`
		DB     `yaml:"database"`
		Hasher `yaml:"hasher"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	Server struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	DB struct {
		URL string `env-required:"true" yaml:"url" env:"PG_URL"`
	}

	Hasher struct {
		Salt string `env-required:"true" env:"HASHER_SALT"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error occurred while unmarshal config: %w", err)
	}

	return cfg, nil
}
