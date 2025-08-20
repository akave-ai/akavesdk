// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Simple cache structure that doesn't duplicate config functionality
type SimpleCache struct {
	FavoriteBuckets []string          `json:"favorite_buckets,omitempty"`
	SavedPaths      map[string]string `json:"saved_paths,omitempty"`
}

var globalCache *SimpleCache

// InitCache loads or creates cache
func InitCache() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil // Fail silently
	}

	cacheFile := filepath.Join(homeDir, ".akave", "simple_cache.json")

	// Create directory
	os.MkdirAll(filepath.Dir(cacheFile), 0o700)

	// Try to load existing cache
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		// Create new cache
		globalCache = &SimpleCache{
			SavedPaths: make(map[string]string),
		}
		return saveSimpleCache()
	}

	// Parse existing cache
	globalCache = &SimpleCache{}
	json.Unmarshal(data, globalCache)
	if globalCache.SavedPaths == nil {
		globalCache.SavedPaths = make(map[string]string)
	}

	return nil
}

// Save cache to file
func saveSimpleCache() error {
	if globalCache == nil {
		return nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	cacheFile := filepath.Join(homeDir, ".akave", "simple_cache.json")
	data, _ := json.MarshalIndent(globalCache, "", "  ")
	return os.WriteFile(cacheFile, data, 0o600)
}

// Add favorite bucket
func AddFavoriteBucket(bucket string) error {
	if globalCache == nil {
		return nil
	}

	// Check if already exists
	for _, b := range globalCache.FavoriteBuckets {
		if b == bucket {
			return nil
		}
	}

	globalCache.FavoriteBuckets = append(globalCache.FavoriteBuckets, bucket)
	return saveSimpleCache()
}

// Remove favorite bucket
func RemoveFavoriteBucket(bucket string) error {
	if globalCache == nil {
		return nil
	}

	for i, b := range globalCache.FavoriteBuckets {
		if b == bucket {
			globalCache.FavoriteBuckets = append(globalCache.FavoriteBuckets[:i], globalCache.FavoriteBuckets[i+1:]...)
			return saveSimpleCache()
		}
	}

	return nil
}

// Save path alias
func SavePathAlias(alias, path string) error {
	if globalCache == nil {
		return nil
	}

	globalCache.SavedPaths[alias] = path
	return saveSimpleCache()
}

// Get path by alias
func GetPathAlias(alias string) (string, bool) {
	if globalCache == nil {
		return "", false
	}

	path, exists := globalCache.SavedPaths[alias]
	return path, exists
}
