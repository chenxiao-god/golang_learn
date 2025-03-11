package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DB   DBConfig
	HTTP HTTPConfig
	JWT  JWTConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type HTTPConfig struct {
	Port int
}

type JWTConfig struct {
	Secret string
}

func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
