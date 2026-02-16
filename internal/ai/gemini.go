package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/internal/config"
)

// GenerateRequest represents the request payload for the AI server
type GenerateRequest struct {
	Diff      string `json:"diff"`
	Type      string `json:"type,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Style     string `json:"style"`
	MaxLength int    `json:"max_length,omitempty"`
}

// GenerateResponse represents the response from the AI server
type GenerateResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Message     string `json:"message"`
		Type        string `json:"type"`
		Scope       string `json:"scope"`
		Description string `json:"description"`
		Breaking    bool   `json:"breaking"`
		Body        string `json:"body"`
	} `json:"data,omitempty"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details,omitempty"`
	} `json:"error,omitempty"`
}

// GenerateCommitMessage sends a request to the AI server to generate a commit message
func GenerateCommitMessage(diff string, cfg *config.Config) (string, error) {
	serverURL := config.GetServerURL(cfg)
	apiKey := config.GetAPIKey(cfg)

	// Prepare request payload
	requestBody := GenerateRequest{
		Diff:      diff,
		Style:     cfg.Style,
		MaxLength: 100,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/api/v1/generate", serverURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}

	// Create HTTP client with timeout
	timeout := time.Duration(cfg.Timeout) * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to connect to AI server: %v\nMake sure the server is running at %s", err, serverURL)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	// Parse response
	var response GenerateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	// Check for errors
	if !response.Success {
		if response.Error.Code == "MISSING_API_KEY" || response.Error.Code == "INVALID_API_KEY" {
			return "", fmt.Errorf("authentication failed: %s\nSet API_KEY in .smartcommitrc.json or SMARTCOMMIT_API_KEY environment variable", response.Error.Message)
		}
		return "", fmt.Errorf("server error: %s - %s", response.Error.Code, response.Error.Message)
	}

	// Return the generated commit message
	message := response.Data.Message
	if message == "" || len(message) < 10 {
		return "chore: update files\n\n- Update project files\n- Apply changes from diff", nil
	}

	return message, nil
}
