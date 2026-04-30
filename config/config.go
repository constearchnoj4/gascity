// Package config provides configuration management for gascity.
// It handles loading, parsing, and validating application configuration
// from environment variables and configuration files.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds the complete application configuration.
type Config struct {
	// Server configuration
	Server ServerConfig

	// Database configuration
	Database DatabaseConfig

	// Application-level settings
	App AppConfig
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig holds database connection settings.
type DatabaseConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// AppConfig holds general application settings.
type AppConfig struct {
	Environment string
	LogLevel    string
	Debug       bool
}

// Load reads configuration from environment variables and returns a Config.
// Missing required variables will result in an error.
func Load() (*Config, error) {
	cfg := &Config{}

	// Server config
	cfg.Server.Host = getEnv("GASCITY_HOST", "0.0.0.0")
	port, err := getEnvInt("GASCITY_PORT", 8080)
	if err != nil {
		return nil, fmt.Errorf("invalid GASCITY_PORT: %w", err)
	}
	cfg.Server.Port = port
	cfg.Server.ReadTimeout = getEnvDuration("GASCITY_READ_TIMEOUT", 15*time.Second)
	cfg.Server.WriteTimeout = getEnvDuration("GASCITY_WRITE_TIMEOUT", 15*time.Second)
	cfg.Server.IdleTimeout = getEnvDuration("GASCITY_IDLE_TIMEOUT", 60*time.Second)

	// Database config
	dsn := getEnv("GASCITY_DATABASE_DSN", "")
	if dsn == "" {
		return nil, fmt.Errorf("GASCITY_DATABASE_DSN is required")
	}
	cfg.Database.DSN = dsn
	maxOpen, err := getEnvInt("GASCITY_DB_MAX_OPEN_CONNS", 25)
	if err != nil {
		return nil, fmt.Errorf("invalid GASCITY_DB_MAX_OPEN_CONNS: %w", err)
	}
	cfg.Database.MaxOpenConns = maxOpen
	maxIdle, err := getEnvInt("GASCITY_DB_MAX_IDLE_CONNS", 5)
	if err != nil {
		return nil, fmt.Errorf("invalid GASCITY_DB_MAX_IDLE_CONNS: %w", err)
	}
	cfg.Database.MaxIdleConns = maxIdle
	cfg.Database.ConnMaxLifetime = getEnvDuration("GASCITY_DB_CONN_MAX_LIFETIME", 5*time.Minute)

	// App config
	cfg.App.Environment = getEnv("GASCITY_ENV", "development")
	cfg.App.LogLevel = getEnv("GASCITY_LOG_LEVEL", "info")
	cfg.App.Debug = getEnvBool("GASCITY_DEBUG", false)

	return cfg, nil
}

// Addr returns the formatted server address string.
func (s *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// getEnv returns the value of the environment variable named by key,
// or fallback if the variable is not set.
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

// getEnvInt returns the integer value of the environment variable named by key,
// or fallback if the variable is not set.
func getEnvInt(key string, fallback int) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback, nil
	}
	return strconv.Atoi(val)
}

// getEnvBool returns the boolean value of the environment variable named by key,
// or fallback if the variable is not set.
func getEnvBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return b
}

// getEnvDuration returns the duration value of the environment variable named by key,
// or fallback if the variable is not set or cannot be parsed.
func getEnvDuration(key string, fallback time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return d
}
