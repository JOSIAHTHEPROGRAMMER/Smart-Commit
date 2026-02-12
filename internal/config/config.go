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
	Model       string `json:"model"`
	Style       string `json:"style"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		DefaultType: "chore",
		Model:       "gemini-3-flash-preview",
		Style:       "detailed",
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

// GetGeminiAPIKey retrieves the Gemini API key from environment
func GetGeminiAPIKey() (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY not found in environment variables")
	}
	return apiKey, nil
}
