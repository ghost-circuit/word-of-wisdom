package client

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const _defaultConfigPath = "config/client.yaml"

// Config represents the client configuration.
type Config struct {
	IsSugarLogger bool   `yaml:"isSugarLogger"`
	ServerAddr    string `yaml:"serverAddr"`
	RPS           int    `yaml:"rps"`
	TotalRequests int    `yaml:"totalRequests"`
}

// NewConfig creates a new client configuration.
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
