package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

	// Try multiple locations in order of priority:
	configPaths := []string{
		// 1. Current directory (highest priority)
		".smartcommitrc.json",
		// 2. Home directory
		filepath.Join(mustGetHomeDir(), ".smartcommitrc.json"),
		// 3. SmartCommit project directory (where the source code is)
		filepath.Join(mustGetGoPath(), "src", "github.com", "JOSIAHTHEPROGRAMMER", "Smart-Commit", ".smartcommitrc.json"),
	}

	var data []byte
	var err error
	foundPath := ""

	for _, path := range configPaths {
		data, err = os.ReadFile(path)
		if err == nil {
			foundPath = path
			break
		}
	}

	if foundPath == "" {
		// No config file found anywhere, use defaults
		return config, nil
	}

	fmt.Printf("DEBUG: Using config from: %s\n", foundPath)

	// Parse JSON
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

// mustGetHomeDir gets home directory or returns empty string
func mustGetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home
}

// mustGetGoPath gets GOPATH or returns empty string
func mustGetGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		// Default GOPATH is ~/go
		return filepath.Join(mustGetHomeDir(), "go")
	}
	return gopath
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
