package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const _defaultConfigPath = "config/server.yaml"

// Config represents the server configuration.
type Config struct {
	IsSugarLogger bool   `yaml:"isSugarLogger"`
	Difficulty    uint8  `yaml:"difficulty"`
	Addr          string `yaml:"addr"`
	PowTimeout    int    `yaml:"powTimeout"`
}

// NewConfig creates and returns a new Config instance.
func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(loadConfigPath(), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadConfigPath() string {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return _defaultConfigPath
	}

	return path
}
