// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage CLI configuration settings",
		Long: `Configure persistent settings for akavecli to avoid repetitive parameter entry.

Settings are stored in ~/.akave/config.json and automatically loaded for all commands.
You can still override any setting using command-line flags.`,
	}

	configSetCmd = &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration value",
		Long: `Set a persistent configuration value. Available keys:
  node-address        - The address of the node RPC (default: 127.0.0.1:5000)
  private-key         - Private key for signing transactions
  encryption-key      - Encryption key for encrypting file data
  account             - Wallet name to use
  maxConcurrency      - Maximum concurrency level (default: 10)
  blockPartSize       - Size of each block part in bytes (default: 131072)
  useConnectionPool   - Use connection pool (true/false, default: true)
  metadata-encryption - Enable metadata encryption (true/false, default: false)

Example:
  akavecli config set node-address connect.akave.ai:5000
  akavecli config set private-key 0x1234567890abcdef...
  akavecli config set encryption-key mySecretKey123`,
		Args: cobra.ExactArgs(2),
		Run:  runConfigSet,
	}

	configGetCmd = &cobra.Command{
		Use:   "get <key>",
		Short: "Get a configuration value",
		Long: `Get the current value of a configuration setting.

Example:
  akavecli config get node-address`,
		Args: cobra.ExactArgs(1),
		Run:  runConfigGet,
	}

	configListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all configuration values",
		Long: `Display all current configuration settings.

Sensitive values (private-key, encryption-key) will be partially masked.`,
		Args: cobra.NoArgs,
		Run:  runConfigList,
	}

	configUnsetCmd = &cobra.Command{
		Use:   "unset <key>",
		Short: "Remove a configuration value",
		Long: `Remove a configuration setting.

Example:
  akavecli config unset private-key`,
		Args: cobra.ExactArgs(1),
		Run:  runConfigUnset,
	}

	configPathCmd = &cobra.Command{
		Use:   "path",
		Short: "Show the configuration file path",
		Long:  `Display the path to the configuration file.`,
		Args:  cobra.NoArgs,
		Run:   runConfigPath,
	}
)

// Config represents the CLI configuration.
type Config struct {
	NodeAddress        string `json:"node-address,omitempty"`
	PrivateKey         string `json:"private-key,omitempty"`
	EncryptionKey      string `json:"encryption-key,omitempty"`
	Account            string `json:"account,omitempty"`
	MaxConcurrency     *int   `json:"maxConcurrency,omitempty"`
	BlockPartSize      *int64 `json:"blockPartSize,omitempty"`
	UseConnectionPool  *bool  `json:"useConnectionPool,omitempty"`
	MetadataEncryption *bool  `json:"metadata-encryption,omitempty"`
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configUnsetCmd)
	configCmd.AddCommand(configPathCmd)
}

// getConfigPath returns the path to the config file.
// This variable can be overridden in tests.
var getConfigPath = func() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(homeDir, ".akave", "config.json"), nil
}

// loadConfig loads the configuration from file.
func loadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Config file doesn't exist yet, return empty config
			return &Config{}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// saveConfig saves the configuration to file.
func saveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// maskSensitiveValue masks sensitive values for display.
func maskSensitiveValue(key, value string) string {
	sensitiveKeys := map[string]bool{
		"private-key":    true,
		"encryption-key": true,
	}

	if sensitiveKeys[key] && len(value) > 8 {
		return value[:4] + "..." + value[len(value)-4:]
	}
	return value
}

func runConfigSet(cmd *cobra.Command, args []string) {
	key := args[0]
	value := args[1]

	config, err := loadConfig()
	if err != nil {
		cmd.PrintErrf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Set the value based on key
	switch key {
	case "node-address":
		config.NodeAddress = value
	case "private-key":
		config.PrivateKey = value
	case "encryption-key":
		config.EncryptionKey = value
	case "account":
		config.Account = value
	case "maxConcurrency":
		var intVal int
		if _, err := fmt.Sscanf(value, "%d", &intVal); err != nil {
			cmd.PrintErrf("Error: maxConcurrency must be an integer\n")
			os.Exit(1)
		}
		config.MaxConcurrency = &intVal
	case "blockPartSize":
		var int64Val int64
		if _, err := fmt.Sscanf(value, "%d", &int64Val); err != nil {
			cmd.PrintErrf("Error: blockPartSize must be an integer\n")
			os.Exit(1)
		}
		config.BlockPartSize = &int64Val
	case "useConnectionPool":
		var boolVal bool
		if value == "true" || value == "1" || value == "yes" {
			boolVal = true
		} else if value == "false" || value == "0" || value == "no" {
			boolVal = false
		} else {
			cmd.PrintErrf("Error: useConnectionPool must be true or false\n")
			os.Exit(1)
		}
		config.UseConnectionPool = &boolVal
	case "metadata-encryption":
		var boolVal bool
		if value == "true" || value == "1" || value == "yes" {
			boolVal = true
		} else if value == "false" || value == "0" || value == "no" {
			boolVal = false
		} else {
			cmd.PrintErrf("Error: metadata-encryption must be true or false\n")
			os.Exit(1)
		}
		config.MetadataEncryption = &boolVal
	default:
		cmd.PrintErrf("Error: unknown configuration key: %s\n", key)
		cmd.PrintErrln("Run 'akavecli config set --help' for available keys")
		os.Exit(1)
	}

	if err := saveConfig(config); err != nil {
		cmd.PrintErrf("Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configuration updated: %s = %s\n", key, maskSensitiveValue(key, value))
}

func runConfigGet(cmd *cobra.Command, args []string) {
	key := args[0]

	config, err := loadConfig()
	if err != nil {
		cmd.PrintErrf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	var value string
	switch key {
	case "node-address":
		value = config.NodeAddress
	case "private-key":
		value = config.PrivateKey
	case "encryption-key":
		value = config.EncryptionKey
	case "account":
		value = config.Account
	case "maxConcurrency":
		if config.MaxConcurrency != nil {
			value = fmt.Sprintf("%d", *config.MaxConcurrency)
		}
	case "blockPartSize":
		if config.BlockPartSize != nil {
			value = fmt.Sprintf("%d", *config.BlockPartSize)
		}
	case "useConnectionPool":
		if config.UseConnectionPool != nil {
			value = fmt.Sprintf("%t", *config.UseConnectionPool)
		}
	case "metadata-encryption":
		if config.MetadataEncryption != nil {
			value = fmt.Sprintf("%t", *config.MetadataEncryption)
		}
	default:
		cmd.PrintErrf("Error: unknown configuration key: %s\n", key)
		os.Exit(1)
	}

	if value == "" {
		fmt.Printf("%s is not set\n", key)
	} else {
		fmt.Println(value)
	}
}

func runConfigList(cmd *cobra.Command, args []string) {
	config, err := loadConfig()
	if err != nil {
		cmd.PrintErrf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Current configuration:")
	fmt.Println()

	if config.NodeAddress != "" {
		fmt.Printf("  node-address: %s\n", config.NodeAddress)
	}
	if config.PrivateKey != "" {
		fmt.Printf("  private-key: %s\n", maskSensitiveValue("private-key", config.PrivateKey))
	}
	if config.EncryptionKey != "" {
		fmt.Printf("  encryption-key: %s\n", maskSensitiveValue("encryption-key", config.EncryptionKey))
	}
	if config.Account != "" {
		fmt.Printf("  account: %s\n", config.Account)
	}
	if config.MaxConcurrency != nil {
		fmt.Printf("  maxConcurrency: %d\n", *config.MaxConcurrency)
	}
	if config.BlockPartSize != nil {
		fmt.Printf("  blockPartSize: %d\n", *config.BlockPartSize)
	}
	if config.UseConnectionPool != nil {
		fmt.Printf("  useConnectionPool: %t\n", *config.UseConnectionPool)
	}
	if config.MetadataEncryption != nil {
		fmt.Printf("  metadata-encryption: %t\n", *config.MetadataEncryption)
	}
}

func runConfigUnset(cmd *cobra.Command, args []string) {
	key := args[0]

	config, err := loadConfig()
	if err != nil {
		cmd.PrintErrf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	switch key {
	case "node-address":
		config.NodeAddress = ""
	case "private-key":
		config.PrivateKey = ""
	case "encryption-key":
		config.EncryptionKey = ""
	case "account":
		config.Account = ""
	case "maxConcurrency":
		config.MaxConcurrency = nil
	case "blockPartSize":
		config.BlockPartSize = nil
	case "useConnectionPool":
		config.UseConnectionPool = nil
	case "metadata-encryption":
		config.MetadataEncryption = nil
	default:
		cmd.PrintErrf("Error: unknown configuration key: %s\n", key)
		os.Exit(1)
	}

	if err := saveConfig(config); err != nil {
		cmd.PrintErrf("Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configuration key '%s' has been unset\n", key)
}

func runConfigPath(cmd *cobra.Command, args []string) {
	configPath, err := getConfigPath()
	if err != nil {
		cmd.PrintErrf("Error getting config path: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(configPath)
}

// applyConfigDefaults loads config from file and applies it to global flags if not set via command line.
func applyConfigDefaults(cmd *cobra.Command) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	// Only apply config values if the flag wasn't explicitly set on command line
	if !cmd.Flags().Changed("node-address") && config.NodeAddress != "" {
		nodeRPCAddress = config.NodeAddress
	}
	if !cmd.Flags().Changed("private-key") && config.PrivateKey != "" {
		privateKey = config.PrivateKey
	}
	if !cmd.Flags().Changed("encryption-key") && config.EncryptionKey != "" {
		encryptionKey = config.EncryptionKey
	}
	if !cmd.Flags().Changed("account") && config.Account != "" {
		accountName = config.Account
	}
	if !cmd.Flags().Changed("maxConcurrency") && config.MaxConcurrency != nil {
		maxConcurrency = *config.MaxConcurrency
	}
	if !cmd.Flags().Changed("blockPartSize") && config.BlockPartSize != nil {
		blockPartSize = *config.BlockPartSize
	}
	if !cmd.Flags().Changed("useConnectionPool") && config.UseConnectionPool != nil {
		useConnectionPool = *config.UseConnectionPool
	}
	if !cmd.Flags().Changed("metadata-encryption") && config.MetadataEncryption != nil {
		useMetadataEncryption = *config.MetadataEncryption
	}

	return nil
}
