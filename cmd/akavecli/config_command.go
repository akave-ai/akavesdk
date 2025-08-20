// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage CLI configuration",
		Long:  "Configure default values, shortcuts, and preferences for the Akave CLI",
	}

	configGetCmd = &cobra.Command{
		Use:   "get [key]",
		Short: "Get configuration value",
		Args:  cobra.MaximumNArgs(1),
		RunE:  cmdConfigGet,
	}

	configSetCmd = &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set configuration value",
		Args:  cobra.ExactArgs(2),
		RunE:  cmdConfigSet,
	}

	configListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all configuration values",
		Args:  cobra.NoArgs,
		RunE:  cmdConfigList,
	}

	configResetCmd = &cobra.Command{
		Use:   "reset [key]",
		Short: "Reset configuration to defaults",
		Args:  cobra.MaximumNArgs(1),
		RunE:  cmdConfigReset,
	}

	configEditCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit configuration file in default editor",
		Args:  cobra.NoArgs,
		RunE:  cmdConfigEdit,
	}
)

func initConfigCommands() {
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configResetCmd)
	configCmd.AddCommand(configEditCmd)
}

func cmdConfigGet(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmdConfigList(cmd, args)
	}

	key := args[0]
	value, err := globalConfigManager.Get(key)
	if err != nil {
		return err
	}

	cmd.Printf("%s = %v\n", key, value)
	return nil
}

func cmdConfigSet(cmd *cobra.Command, args []string) error {
	key := args[0]
	value := args[1]

	if err := globalConfigManager.Set(key, value); err != nil {
		return err
	}

	cmd.Printf("Set %s = %s\n", key, value)
	return nil
}

func cmdConfigList(cmd *cobra.Command, args []string) error {
	config := GetConfig()
	if config == nil {
		return fmt.Errorf("configuration not loaded")
	}

	formatter := NewOutputFormatter()

	// Support JSON/YAML output for scripting
	if formatter.format == "json" || formatter.format == "yaml" {
		payload := buildConfigPayload(config)
		return formatter.FormatSingle(cmd, payload, "config")
	}

	// Pretty table output
	printConfigTable(cmd, config)
	return nil
}

func cmdConfigReset(cmd *cobra.Command, args []string) error {
	key := ""
	if len(args) > 0 {
		key = args[0]
	}

	if key == "" {
		if !ConfirmAction("Reset all configuration to defaults?", false) {
			cmd.Printf("Configuration reset cancelled\n")
			return nil
		}
	}

	if err := globalConfigManager.Reset(key); err != nil {
		return err
	}

	if key == "" {
		cmd.Printf("Configuration reset to defaults\n")
	} else {
		cmd.Printf("Reset %s to default value\n", key)
	}
	return nil
}

func cmdConfigEdit(cmd *cobra.Command, args []string) error {
	editor := getEditor()
	if editor == "" {
		return fmt.Errorf("no editor found. Set EDITOR or VISUAL environment variable")
	}

	configFile := globalConfigManager.GetConfigFilePath()
	cmd.Printf("Opening %s with %s...\n", configFile, editor)

	c := exec.Command(editor, configFile)
	c.Stdin, c.Stdout, c.Stderr = cmd.InOrStdin(), cmd.OutOrStdout(), cmd.ErrOrStderr()
	return c.Run()
}

// Helper functions

func buildConfigPayload(config *Config) map[string]interface{} {
	return map[string]interface{}{
		"config_file": globalConfigManager.GetConfigFilePath(),
		"core": map[string]interface{}{
			"node_address":        config.NodeAddress,
			"max_concurrency":     config.MaxConcurrency,
			"block_part_size":     config.BlockPartSize,
			"use_connection_pool": config.UseConnectionPool,
			"default_account":     config.DefaultAccount,
			"keystore_dir":        config.KeystoreDir,
		},
		"preferences": map[string]interface{}{
			"progress_style": config.ProgressStyle,
			"output_format":  config.OutputFormat,
			"auto_confirm":   config.AutoConfirm,
			"default_bucket": config.DefaultBucket,
		},
		"recent": map[string]interface{}{
			"buckets":           config.RecentBuckets,
			"files":             config.RecentFiles,
			"last_node_address": config.LastNodeAddress,
			"last_download_dir": config.LastDownloadDir,
		},
	}
}

func printConfigTable(cmd *cobra.Command, config *Config) {
	cmd.Printf("Configuration file: %s\n\n", globalConfigManager.GetConfigFilePath())

	// Core settings
	cmd.Printf("Core Settings:\n")
	cmd.Printf("  node_address       = %s\n", config.NodeAddress)
	cmd.Printf("  max_concurrency    = %d\n", config.MaxConcurrency)
	cmd.Printf("  block_part_size    = %d\n", config.BlockPartSize)
	cmd.Printf("  use_connection_pool = %t\n", config.UseConnectionPool)
	cmd.Printf("  default_account    = %s\n", config.DefaultAccount)
	cmd.Printf("  keystore_dir       = %s\n", config.KeystoreDir)

	// User preferences
	cmd.Printf("\nUser Preferences:\n")
	cmd.Printf("  progress_style     = %s\n", config.ProgressStyle)
	cmd.Printf("  output_format      = %s\n", config.OutputFormat)
	cmd.Printf("  auto_confirm       = %t\n", config.AutoConfirm)
	cmd.Printf("  default_bucket     = %s\n", config.DefaultBucket)

	// Cache information
	cmd.Printf("\nCached Values:\n")
	cmd.Printf("  last_node_address  = %s\n", config.LastNodeAddress)
	cmd.Printf("  last_download_dir  = %s\n", config.LastDownloadDir)

	// Recent items
	if len(config.RecentBuckets) > 0 {
		cmd.Printf("\nRecent Buckets:\n")
		for i, bucket := range config.RecentBuckets {
			cmd.Printf("  %d. %s\n", i+1, bucket)
		}
	}

	if len(config.RecentFiles) > 0 {
		cmd.Printf("\nRecent Files:\n")
		for i, file := range config.RecentFiles {
			if i >= 10 { // Show only top 10
				break
			}
			cmd.Printf("  %d. %s\n", i+1, file)
		}
	}
}

func getEditor() string {
	editors := []string{"EDITOR", "VISUAL"}
	for _, env := range editors {
		if editor := os.Getenv(env); editor != "" {
			return editor
		}
	}
	return ""
}
