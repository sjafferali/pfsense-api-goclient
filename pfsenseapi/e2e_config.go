// +build e2e

package pfsenseapi

import (
	"encoding/json"
	"fmt"
	"os"
)

// E2EConfig holds configuration for end-to-end tests
type E2EConfig struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoadE2EConfig loads the e2e test configuration from a file
// It first checks the E2E_CONFIG_FILE environment variable,
// then falls back to "e2e_config.json" in the current directory
func LoadE2EConfig() (*E2EConfig, error) {
	configPath := os.Getenv("E2E_CONFIG_FILE")
	if configPath == "" {
		configPath = "e2e_config.json"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", configPath, err)
	}

	var config E2EConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Validate required fields
	if config.URL == "" {
		return nil, fmt.Errorf("URL is required in config")
	}
	if config.Username == "" {
		return nil, fmt.Errorf("username is required in config")
	}
	if config.Password == "" {
		return nil, fmt.Errorf("password is required in config")
	}

	return &config, nil
}
