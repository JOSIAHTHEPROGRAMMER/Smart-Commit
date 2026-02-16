package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the smartcommit configuration
type Config struct {
	DefaultType string `json:"default_type"`
	Style       string `json:"style"`
	ServerURL   string `json:"server_url"`
	APIKey      string `json:"api_key"`
	Timeout     int    `json:"timeout"` // in seconds
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		DefaultType: "chore",
		Style:       "detailed",
		ServerURL:   "http://localhost:8080",
		APIKey:      "",
		Timeout:     30,
	}
}

// LoadConfig loads configuration from .smartcommitrc.json
func LoadConfig() (*Config, error) {
	config := DefaultConfig()

	// Try to read config file
	data, err := os.ReadFile(".smartcommitrc.json")
	if err != nil {
		return config, nil
	}

	// Parse JSON
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	return nil
}

// GetServerURL retrieves the server URL from config or environment
func GetServerURL(cfg *Config) string {
	serverURL := os.Getenv("SMARTCOMMIT_SERVER_URL")
	if serverURL != "" {
		return serverURL
	}

	if cfg.ServerURL != "" {
		return cfg.ServerURL
	}

	return "http://localhost:8080"
}

// GetAPIKey retrieves the API key from config or environment
func GetAPIKey(cfg *Config) string {
	apiKey := os.Getenv("SMARTCOMMIT_API_KEY")
	if apiKey != "" {
		return apiKey
	}

	return cfg.APIKey
}
