// cmd/akavecli/utils.go - CLEAN VERSION (No command redeclarations)
// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// OutputFormatter handles different output formats
type OutputFormatter struct {
	format string
}

func NewOutputFormatter() *OutputFormatter {
	config := GetConfig()
	format := "table"
	if config != nil && config.OutputFormat != "" {
		format = config.OutputFormat
	}
	return &OutputFormatter{format: format}
}

func (of *OutputFormatter) FormatBuckets(cmd *cobra.Command, buckets []interface{}) error {
	switch of.format {
	case "json":
		return of.outputJSON(cmd, map[string]interface{}{"buckets": buckets})
	case "yaml":
		return of.outputYAML(cmd, map[string]interface{}{"buckets": buckets})
	default:
		return of.outputBucketsTable(cmd, buckets)
	}
}

func (of *OutputFormatter) FormatFiles(cmd *cobra.Command, files []interface{}) error {
	switch of.format {
	case "json":
		return of.outputJSON(cmd, map[string]interface{}{"files": files})
	case "yaml":
		return of.outputYAML(cmd, map[string]interface{}{"files": files})
	default:
		return of.outputFilesTable(cmd, files)
	}
}

func (of *OutputFormatter) FormatSingle(cmd *cobra.Command, item interface{}, itemType string) error {
	switch of.format {
	case "json":
		return of.outputJSON(cmd, map[string]interface{}{itemType: item})
	case "yaml":
		return of.outputYAML(cmd, map[string]interface{}{itemType: item})
	default:
		return of.outputSingleItem(cmd, item)
	}
}

func (of *OutputFormatter) outputJSON(cmd *cobra.Command, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	cmd.Printf("%s\n", jsonData)
	return nil
}

func (of *OutputFormatter) outputYAML(cmd *cobra.Command, data interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	cmd.Printf("%s", yamlData)
	return nil
}

func (of *OutputFormatter) outputBucketsTable(cmd *cobra.Command, items []interface{}) error {
	if len(items) == 0 {
		cmd.Printf("No buckets found\n")
		return nil
	}

	cmd.Printf("%-30s %-20s\n", "NAME", "CREATED")
	cmd.Printf("%-30s %-20s\n", strings.Repeat("-", 30), strings.Repeat("-", 20))

	for _, item := range items {
		if bucketMap, ok := item.(map[string]interface{}); ok {
			name := getStringFromInterface(bucketMap["Name"])
			createdAt := getTimeFromInterface(bucketMap["CreatedAt"])
			cmd.Printf("%-30s %-20s\n", name, formatTime(createdAt))
		}
	}

	return nil
}

func (of *OutputFormatter) outputFilesTable(cmd *cobra.Command, items []interface{}) error {
	if len(items) == 0 {
		cmd.Printf("No files found\n")
		return nil
	}

	cmd.Printf("%-30s %-15s %-20s\n", "NAME", "SIZE", "CREATED")
	cmd.Printf("%-30s %-15s %-20s\n", strings.Repeat("-", 30), strings.Repeat("-", 15), strings.Repeat("-", 20))

	for _, item := range items {
		if fileMap, ok := item.(map[string]interface{}); ok {
			name := getStringFromInterface(fileMap["Name"])
			size := getInt64FromInterface(fileMap["Size"])
			createdAt := getTimeFromInterface(fileMap["CreatedAt"])
			cmd.Printf("%-30s %-15s %-20s\n", name, formatFileSize(size), formatTime(createdAt))
		}
	}

	return nil
}

func (of *OutputFormatter) outputSingleItem(cmd *cobra.Command, item interface{}) error {
	if itemMap, ok := item.(map[string]interface{}); ok {
		for key, value := range itemMap {
			cmd.Printf("%s: %v\n", key, value)
		}
	} else {
		cmd.Printf("%v\n", item)
	}
	return nil
}

// Confirmation and user interaction
func ConfirmAction(message string, defaultYes bool) bool {
	config := GetConfig()
	if config != nil && config.AutoConfirm {
		return true
	}

	prompt := "[y/N]"
	if defaultYes {
		prompt = "[Y/n]"
	}

	fmt.Printf("%s %s: ", message, prompt)

	var response string
	fmt.Scanln(&response)

	response = strings.TrimSpace(strings.ToLower(response))

	if response == "" {
		return defaultYes
	}

	return response == "y" || response == "yes"
}

// Enhanced success/error messaging
func PrintSuccess(cmd *cobra.Command, operation string, details ...string) {
	config := GetConfig()
	if config != nil && config.OutputFormat == "json" {
		result := map[string]interface{}{
			"status":    "success",
			"operation": operation,
		}
		if len(details) > 0 {
			result["details"] = details
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		cmd.Printf("%s\n", data)
	} else {
		cmd.Printf("✓ %s", operation)
		if len(details) > 0 {
			cmd.Printf(": %s", strings.Join(details, ", "))
		}
		cmd.Printf("\n")
	}
}

func PrintError(cmd *cobra.Command, operation string, err error) {
	config := GetConfig()
	if config != nil && config.OutputFormat == "json" {
		result := map[string]interface{}{
			"status":    "error",
			"operation": operation,
			"error":     err.Error(),
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		cmd.Printf("%s\n", data)
	} else {
		cmd.PrintErrf("✗ %s failed: %v\n", operation, err)
	}
}

// Utility functions for type conversion and formatting
func getStringFromInterface(v interface{}) string {
	if str, ok := v.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", v)
}

func getInt64FromInterface(v interface{}) int64 {
	switch val := v.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case float64:
		return int64(val)
	default:
		return 0
	}
}

func getTimeFromInterface(v interface{}) time.Time {
	switch val := v.(type) {
	case time.Time:
		return val
	case string:
		if t, err := time.Parse(time.RFC3339, val); err == nil {
			return t
		}
	}
	return time.Time{}
}

func formatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return "N/A"
	}
	return t.Format("2006-01-02 15:04")
}

func formatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%.0fms", float64(d.Nanoseconds())/1e6)
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%.1fm", d.Minutes())
	}
	return fmt.Sprintf("%.1fh", d.Hours())
}
