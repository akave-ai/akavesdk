// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// Config represents the CLI configuration with validation tags
type Config struct {
	// Core settings
	NodeAddress       string `json:"node_address,omitempty" env:"AKAVE_NODE_ADDRESS" default:"127.0.0.1:5000"`
	MaxConcurrency    int    `json:"max_concurrency,omitempty" env:"AKAVE_MAX_CONCURRENCY" default:"10" min:"1" max:"100"`
	BlockPartSize     int64  `json:"block_part_size,omitempty" env:"AKAVE_BLOCK_PART_SIZE" default:"131072" min:"1024"`
	UseConnectionPool bool   `json:"use_connection_pool,omitempty" env:"AKAVE_USE_CONNECTION_POOL" default:"true"`
	DefaultAccount    string `json:"default_account,omitempty" env:"AKAVE_DEFAULT_ACCOUNT"`
	KeystoreDir       string `json:"keystore_dir,omitempty" env:"AKAVE_KEYSTORE_DIR"`

	// Frequently used values cache
	RecentBuckets   []string `json:"recent_buckets,omitempty" max_items:"10"`
	RecentFiles     []string `json:"recent_files,omitempty" max_items:"20"`
	LastNodeAddress string   `json:"last_node_address,omitempty"`
	LastDownloadDir string   `json:"last_download_dir,omitempty"`

	// User preferences
	ProgressStyle string `json:"progress_style,omitempty" env:"AKAVE_PROGRESS_STYLE" default:"bar" enum:"bar,spinner,dots"`
	OutputFormat  string `json:"output_format,omitempty" env:"AKAVE_OUTPUT_FORMAT" default:"table" enum:"table,json,yaml"`
	AutoConfirm   bool   `json:"auto_confirm,omitempty" env:"AKAVE_AUTO_CONFIRM" default:"false"`
	DefaultBucket string `json:"default_bucket,omitempty" env:"AKAVE_DEFAULT_BUCKET"`
}

// ConfigManager handles all configuration operations
type ConfigManager struct {
	configDir  string
	configFile string
	config     *Config
}

var globalConfigManager *ConfigManager

// InitConfigManager initializes the global configuration manager
func InitConfigManager() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	globalConfigManager = &ConfigManager{
		configDir:  filepath.Join(homeDir, ".akave"),
		configFile: filepath.Join(homeDir, ".akave", "config.json"),
	}

	return globalConfigManager.Load()
}

// GetConfig returns the global configuration
func GetConfig() *Config {
	if globalConfigManager == nil || globalConfigManager.config == nil {
		return defaultConfig()
	}
	return globalConfigManager.config
}

// Load loads configuration from file or creates default
func (cm *ConfigManager) Load() error {
	// Create config directory if it doesn't exist (secure permissions)
	if err := os.MkdirAll(cm.configDir, 0o700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Load existing config or create default ONLY if missing
	cfg, err := cm.loadFromFile()
	if err != nil {
		if os.IsNotExist(err) {
			cm.config = defaultConfig()
			if err := cm.Save(); err != nil {
				return fmt.Errorf("failed to save default config: %w", err)
			}
		} else {
			return fmt.Errorf("failed to load config (not overwriting): %w", err)
		}
	} else {
		cm.config = cfg
	}

	// Apply environment variable overrides
	cm.applyEnvironmentOverrides()

	return nil
}

// Save saves configuration to file atomically
func (cm *ConfigManager) Save() error {
	if cm.config == nil {
		return fmt.Errorf("config is nil")
	}

	data, err := json.MarshalIndent(cm.config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write to temporary file first (atomic operation)
	tmp := cm.configFile + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return fmt.Errorf("failed to write temp config file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tmp, cm.configFile); err != nil {
		_ = os.Remove(tmp) // Clean up temp file on failure
		return fmt.Errorf("failed to atomically replace config: %w", err)
	}

	return nil
}

// Get returns a configuration value by key
func (cm *ConfigManager) Get(key string) (interface{}, error) {
	if cm.config == nil {
		return nil, fmt.Errorf("config not loaded")
	}

	value, err := getFieldByJSONTag(cm.config, key)
	if err != nil {
		return nil, fmt.Errorf("unknown configuration key: %s", key)
	}
	return value, nil
}

// Set sets a configuration value by key
func (cm *ConfigManager) Set(key, value string) error {
	if cm.config == nil {
		return fmt.Errorf("config not loaded")
	}

	if err := setFieldByJSONTag(cm.config, key, value); err != nil {
		return err
	}

	return cm.Save()
}

// Reset resets configuration to defaults
func (cm *ConfigManager) Reset(key string) error {
	if key == "" {
		// Reset entire config
		cm.config = defaultConfig()
	} else {
		// Reset specific key
		defaultCfg := defaultConfig()
		defaultValue, err := getFieldByJSONTag(defaultCfg, key)
		if err != nil {
			return fmt.Errorf("unknown configuration key: %s", key)
		}
		if err := setFieldByJSONTagDirect(cm.config, key, defaultValue); err != nil {
			return err
		}
	}

	return cm.Save()
}

// UpdateRecentItem updates a recent items list
func (cm *ConfigManager) UpdateRecentItem(listType, item string) error {
	if cm.config == nil {
		return nil
	}

	var list *[]string
	var maxItems int

	switch listType {
	case "buckets":
		list = &cm.config.RecentBuckets
		maxItems = 10
	case "files":
		list = &cm.config.RecentFiles
		maxItems = 20
	default:
		return fmt.Errorf("unknown list type: %s", listType)
	}

	// Remove if already exists
	for i, existing := range *list {
		if existing == item {
			*list = append((*list)[:i], (*list)[i+1:]...)
			break
		}
	}

	// Add to front
	*list = append([]string{item}, *list...)

	// Keep only max items
	if len(*list) > maxItems {
		*list = (*list)[:maxItems]
	}

	return cm.Save()
}

// UpdateLastUsed updates last used values
func (cm *ConfigManager) UpdateLastUsed(key, value string) error {
	if cm.config == nil {
		return nil
	}

	switch key {
	case "node_address":
		cm.config.LastNodeAddress = value
	case "download_dir":
		cm.config.LastDownloadDir = value
	default:
		return fmt.Errorf("unknown last used key: %s", key)
	}

	return cm.Save()
}

// GetConfigFilePath returns the configuration file path
func (cm *ConfigManager) GetConfigFilePath() string {
	return cm.configFile
}

// Validation and helper methods

func defaultConfig() *Config {
	cfg := &Config{}
	setDefaults(cfg)
	return cfg
}

func (cm *ConfigManager) loadFromFile() (*Config, error) {
	data, err := os.ReadFile(cm.configFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Set defaults for missing fields
	setDefaults(&cfg)

	// Validate configuration
	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, nil
}

func (cm *ConfigManager) applyEnvironmentOverrides() {
	if cm.config == nil {
		return
	}

	val := reflect.ValueOf(cm.config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		envTag := fieldType.Tag.Get("env")
		if envTag == "" {
			continue
		}

		envValue := os.Getenv(envTag)
		if envValue == "" {
			continue
		}

		// Set the environment value
		if err := setFieldValue(field, envValue); err != nil {
			// Log error but continue
			continue
		}
	}
}

// Reflection-based field access

func getFieldByJSONTag(cfg *Config, jsonTag string) (interface{}, error) {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" {
			continue
		}

		// Remove ",omitempty" from tag
		tagName := strings.Split(tag, ",")[0]
		if tagName == jsonTag {
			return field.Interface(), nil
		}
	}

	return nil, fmt.Errorf("field not found")
}

func setFieldByJSONTag(cfg *Config, jsonTag, value string) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" {
			continue
		}

		tagName := strings.Split(tag, ",")[0]
		if tagName == jsonTag {
			// Validate before setting
			if err := validateFieldValue(fieldType, value); err != nil {
				return err
			}
			return setFieldValue(field, value)
		}
	}

	return fmt.Errorf("field not found")
}

func setFieldByJSONTagDirect(cfg *Config, jsonTag string, value interface{}) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" {
			continue
		}

		tagName := strings.Split(tag, ",")[0]
		if tagName == jsonTag {
			field.Set(reflect.ValueOf(value))
			return nil
		}
	}

	return fmt.Errorf("field not found")
}

func setFieldValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer value: %s", value)
		}
		field.SetInt(intVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid boolean value: %s", value)
		}
		field.SetBool(boolVal)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
	return nil
}

// Validation functions

func setDefaults(cfg *Config) {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Skip if field already has a value
		if !isZeroValue(field) {
			continue
		}

		defaultTag := fieldType.Tag.Get("default")
		if defaultTag == "" {
			continue
		}

		// Handle special defaults
		if defaultTag == "home_wallets" {
			if homeDir, err := os.UserHomeDir(); err == nil {
				field.SetString(filepath.Join(homeDir, ".akave_wallets"))
			}
			continue
		}

		_ = setFieldValue(field, defaultTag)
	}
}

func validateConfig(cfg *Config) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if err := validateFieldValue(fieldType, fmt.Sprintf("%v", field.Interface())); err != nil {
			return fmt.Errorf("field %s: %w", fieldType.Name, err)
		}
	}

	return nil
}

func validateFieldValue(fieldType reflect.StructField, value string) error {
	// Validate enum values
	if enumTag := fieldType.Tag.Get("enum"); enumTag != "" {
		validValues := strings.Split(enumTag, ",")
		valid := false
		for _, validValue := range validValues {
			if strings.TrimSpace(validValue) == value {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid value %s (valid: %s)", value, enumTag)
		}
	}

	// Validate min/max for numeric fields
	if fieldType.Type.Kind() == reflect.Int || fieldType.Type.Kind() == reflect.Int64 {
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer: %s", value)
		}

		if minTag := fieldType.Tag.Get("min"); minTag != "" {
			minVal, _ := strconv.ParseInt(minTag, 10, 64)
			if intVal < minVal {
				return fmt.Errorf("value %d below minimum %d", intVal, minVal)
			}
		}

		if maxTag := fieldType.Tag.Get("max"); maxTag != "" {
			maxVal, _ := strconv.ParseInt(maxTag, 10, 64)
			if intVal > maxVal {
				return fmt.Errorf("value %d above maximum %d", intVal, maxVal)
			}
		}
	}

	return nil
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

// Convenience functions for backward compatibility

func UpdateRecentBuckets(bucketName string) {
	if globalConfigManager != nil {
		_ = globalConfigManager.UpdateRecentItem("buckets", bucketName)
	}
}

func UpdateRecentFiles(fileName string) {
	if globalConfigManager != nil {
		_ = globalConfigManager.UpdateRecentItem("files", fileName)
	}
}

func UpdateLastNodeAddress(nodeAddress string) {
	if globalConfigManager != nil {
		_ = globalConfigManager.UpdateLastUsed("node_address", nodeAddress)
	}
}

func UpdateLastDownloadDir(dir string) {
	if globalConfigManager != nil {
		_ = globalConfigManager.UpdateLastUsed("download_dir", dir)
	}
}
