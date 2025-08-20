// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"github.com/spf13/cobra"
)

var (
	cacheCmd = &cobra.Command{
		Use:   "cache",
		Short: "Manage CLI cache",
	}

	cacheShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show cache contents",
		Args:  cobra.NoArgs,
		RunE:  cmdCacheShow,
	}

	cacheFavCmd = &cobra.Command{
		Use:   "fav",
		Short: "Manage favorite buckets",
	}

	cacheFavAddCmd = &cobra.Command{
		Use:   "add <bucket>",
		Short: "Add favorite bucket",
		Args:  cobra.ExactArgs(1),
		RunE:  cmdCacheFavAdd,
	}

	cacheFavListCmd = &cobra.Command{
		Use:   "list",
		Short: "List favorite buckets",
		Args:  cobra.NoArgs,
		RunE:  cmdCacheFavList,
	}
)

func initCacheCommands() {
	cacheFavCmd.AddCommand(cacheFavAddCmd)
	cacheFavCmd.AddCommand(cacheFavListCmd)

	cacheCmd.AddCommand(cacheShowCmd)
	cacheCmd.AddCommand(cacheFavCmd)
}

func cmdCacheShow(cmd *cobra.Command, args []string) error {
	// Get data from your existing config system
	config := GetConfig()

	cmd.Printf("Cache Information:\n\n")

	if config != nil {
		if config.LastNodeAddress != "" {
			cmd.Printf("Last Node: %s\n", config.LastNodeAddress)
		}

		if len(config.RecentBuckets) > 0 {
			cmd.Printf("\nRecent Buckets:\n")
			for i, bucket := range config.RecentBuckets {
				cmd.Printf("  %d. %s\n", i+1, bucket)
			}
		}

		if len(config.RecentFiles) > 0 {
			cmd.Printf("\nRecent Files:\n")
			for i, file := range config.RecentFiles {
				cmd.Printf("  %d. %s\n", i+1, file)
			}
		}
	}

	// Show simple cache data
	if globalCache != nil && len(globalCache.FavoriteBuckets) > 0 {
		cmd.Printf("\nFavorite Buckets:\n")
		for i, bucket := range globalCache.FavoriteBuckets {
			cmd.Printf("  %d. %s\n", i+1, bucket)
		}
	}

	return nil
}

func cmdCacheFavAdd(cmd *cobra.Command, args []string) error {
	bucket := args[0]
	if err := AddFavoriteBucket(bucket); err != nil {
		return err
	}
	cmd.Printf("Added '%s' to favorites\n", bucket)
	return nil
}

func cmdCacheFavList(cmd *cobra.Command, args []string) error {
	if globalCache == nil || len(globalCache.FavoriteBuckets) == 0 {
		cmd.Printf("No favorite buckets\n")
		return nil
	}

	cmd.Printf("Favorite Buckets:\n")
	for i, bucket := range globalCache.FavoriteBuckets {
		cmd.Printf("  %d. %s\n", i+1, bucket)
	}

	return nil
}
