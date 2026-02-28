package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/vizurth/rate-limiter/internal/postgres"
	"github.com/vizurth/rate-limiter/internal/redis"
)

type Config struct {
	// Server
	Environment string `yaml:"environment" default:"production"`

	Postgres postgres.Config `yaml:"postgres"`

	Redis redis.Config `yaml:"redis"`
}

func New() (*Config, error) {
	var cfg Config

	configPaths := []string{
		"./configs/config.yaml",
		"configs/config.yaml",
		"../configs/config.yaml",
		"../../configs/config.yaml",
	}

	var configPath string
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			configPath = path
			break
		}
	}

	if configPath == "" {
		return &Config{}, fmt.Errorf("config file not found")
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &Config{}, fmt.Errorf("failed to read config: %w", err)
	}

	return &cfg, nil
}
