// Package config - included struct and init function fow work with client configuration
package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Config determines the basic parameters of the agent's operation
type Config struct {
	HostGRPC string `env:"GRPC_ADDRESS"`
	LogLevel string `env:"LOG_LEVEL"`
	FileSize int    `env:"FILE_SIZE"`
}

// GetConfig - method creates a new configuration and sets values from environment variables and command line flags
func GetConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, fmt.Errorf("env parse: %w", err)
	}

	hostGRPC := flag.String("g", "localhost:3200", "адрес эндпоинта grpc-сервера")
	logLevel := flag.String("l", "info", "log level")
	fileSize := flag.Int("s", 4000000, "размер файлов")
	flag.Parse()

	if config.HostGRPC == "" {
		config.HostGRPC = *hostGRPC
	}
	if config.LogLevel == "" {
		config.LogLevel = *logLevel
	}
	if config.FileSize == 0 {
		config.FileSize = *fileSize
	}
	return config, nil
}
