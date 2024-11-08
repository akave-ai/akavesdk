// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package version contains versioning.
package version

import (
	"fmt"
	"runtime/debug"
)

// Version encapsulates minimal version information.
type Version struct {
	Commit    string
	Timestamp string
}

// Format formats version for CLI output.
func (v Version) Format() string {
	return fmt.Sprintf("commit: %s\ntimestamp: %s\n", v.Commit, v.Timestamp)
}

// Info retrieves process version.
func Info() Version {
	var (
		commit, timestamp = "unknown", "unknown"
	)

	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				commit = setting.Value
			}
			if setting.Key == "vcs.time" {
				timestamp = setting.Value
			}
		}
	}

	return Version{
		Commit:    commit,
		Timestamp: timestamp,
	}
}
