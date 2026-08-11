// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func Test_ConfigSetAndGet(t *testing.T) {
	// Create temporary config directory
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	// Override getConfigPath for testing
	originalGetConfigPath := getConfigPath
	getConfigPath = func() (string, error) {
		return configPath, nil
	}
	defer func() { getConfigPath = originalGetConfigPath }()

	// Test setting node-address
	config := &Config{}
	config.NodeAddress = "test-node:5000"
	if err := saveConfig(config); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Test loading config
	loaded, err := loadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loaded.NodeAddress != "test-node:5000" {
		t.Errorf("expected node-address to be 'test-node:5000', got '%s'", loaded.NodeAddress)
	}
}

func Test_ConfigMaskSensitiveValue(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		expected string
	}{
		{
			name:     "mask private key",
			key:      "private-key",
			value:    "0x1234567890abcdef",
			expected: "0x12...cdef",
		},
		{
			name:     "mask encryption key",
			key:      "encryption-key",
			value:    "mySecretKey123456",
			expected: "mySe...3456",
		},
		{
			name:     "no mask for non-sensitive",
			key:      "node-address",
			value:    "connect.akave.ai:5000",
			expected: "connect.akave.ai:5000",
		},
		{
			name:     "short value not masked",
			key:      "private-key",
			value:    "short",
			expected: "short",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maskSensitiveValue(tt.key, tt.value)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func Test_ConfigJSON(t *testing.T) {
	maxConcurrency := 20
	blockPartSize := int64(262144)
	useConnectionPool := false
	metadataEncryption := true

	config := &Config{
		NodeAddress:        "test-node:5000",
		PrivateKey:         "0x123456",
		EncryptionKey:      "testkey",
		Account:            "myaccount",
		MaxConcurrency:     &maxConcurrency,
		BlockPartSize:      &blockPartSize,
		UseConnectionPool:  &useConnectionPool,
		MetadataEncryption: &metadataEncryption,
	}

	// Marshal to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("failed to marshal config: %v", err)
	}

	// Unmarshal back
	var decoded Config
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal config: %v", err)
	}

	// Verify all fields
	if decoded.NodeAddress != config.NodeAddress {
		t.Errorf("node-address mismatch")
	}
	if decoded.PrivateKey != config.PrivateKey {
		t.Errorf("private-key mismatch")
	}
	if decoded.EncryptionKey != config.EncryptionKey {
		t.Errorf("encryption-key mismatch")
	}
	if decoded.Account != config.Account {
		t.Errorf("account mismatch")
	}
	if *decoded.MaxConcurrency != *config.MaxConcurrency {
		t.Errorf("maxConcurrency mismatch")
	}
	if *decoded.BlockPartSize != *config.BlockPartSize {
		t.Errorf("blockPartSize mismatch")
	}
	if *decoded.UseConnectionPool != *config.UseConnectionPool {
		t.Errorf("useConnectionPool mismatch")
	}
	if *decoded.MetadataEncryption != *config.MetadataEncryption {
		t.Errorf("metadata-encryption mismatch")
	}
}

func Test_ConfigNonExistentFile(t *testing.T) {
	// Create temporary directory
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "nonexistent", "config.json")

	// Override getConfigPath for testing
	originalGetConfigPath := getConfigPath
	getConfigPath = func() (string, error) {
		return configPath, nil
	}
	defer func() { getConfigPath = originalGetConfigPath }()

	// Loading non-existent config should return empty config
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("failed to load non-existent config: %v", err)
	}

	if config.NodeAddress != "" {
		t.Errorf("expected empty config, got non-empty node-address")
	}
}

func Test_ConfigFilePermissions(t *testing.T) {
	// Create temporary config directory
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	// Override getConfigPath for testing
	originalGetConfigPath := getConfigPath
	getConfigPath = func() (string, error) {
		return configPath, nil
	}
	defer func() { getConfigPath = originalGetConfigPath }()

	// Save config
	config := &Config{PrivateKey: "sensitive-data"}
	if err := saveConfig(config); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Check file permissions (should be 0600)
	info, err := os.Stat(configPath)
	if err != nil {
		t.Fatalf("failed to stat config file: %v", err)
	}

	mode := info.Mode().Perm()
	if mode != 0600 {
		t.Errorf("expected file permissions 0600, got %o", mode)
	}
}
