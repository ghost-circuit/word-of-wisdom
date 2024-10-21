package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const _defaultConfigPath = "config/server.yaml"

// Config represents the server configuration.
type Config struct {
	IsSugarLogger bool   `yaml:"isSugarLogger"`
	Difficulty    uint8  `yaml:"difficulty"`
	Addr          string `yaml:"addr"`

	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func (c *Config) DatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.DBName,
		c.Postgres.SSLMode,
	)
}

// LoadConfig creates and returns a new Config instance.
func LoadConfig() (*Config, error) {
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
