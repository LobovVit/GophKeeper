// Package config - included struct and init function fow work with client configuration
package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

// Config determines the basic parameters of the server's operation
type Config struct {
	HostGRPC      string        `env:"GRPC_ADDRESS"`
	MigrationDir  string        `env:"MGR_DIR"`
	LogLevel      string        `env:"LOG_LEVEL"`
	DSN           string        `env:"DATABASE_DSN" `
	Files         string        `env:"FILES"`
	TokenLifetime time.Duration `env:"TOKEN_LIFETIME"`
}

// GetConfig - method creates a new configuration and sets values from environment variables and command line flags
func GetConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, fmt.Errorf("env parse: %w", err)
	}

	hostGRPC := flag.String("g", "localhost:3200", "адрес эндпоинта grpc-сервера")
	migrationDir := flag.String("m", "./migrations", "файлы миграции")
	files := flag.String("f", "./files", "хранилище файлов")
	logLevel := flag.String("l", "info", "log level")
	dsn := flag.String("d", "postgresql://postgres:password@10.66.66.3:5432/postgres?sslmode=disable", "строка подключения к БД")
	flag.Parse()

	config.TokenLifetime = 30000 * time.Second
	if config.HostGRPC == "" {
		config.HostGRPC = *hostGRPC
	}
	if config.MigrationDir == "" {
		config.MigrationDir = *migrationDir
	}
	if config.Files == "" {
		config.Files = *files
	}
	if config.LogLevel == "" {
		config.LogLevel = *logLevel
	}
	if config.DSN == "" {
		config.DSN = *dsn
	}
	return config, nil
}
