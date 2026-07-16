// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package version contains versioning.
package version

import (
	"fmt"
	"runtime/debug"
)

// tag is set at build time via goreleaser ldflags:
//
//	-X "github.com/akave-ai/akavesdk/private/version.tag=v0.4.4"
var tag = "dev"

// Version encapsulates minimal version information.
type Version struct {
	Tag       string
	Commit    string
	Timestamp string
}

// Format formats version for CLI output.
func (v Version) Format() string {
	return fmt.Sprintf("version: %s\ncommit: %s\ntimestamp: %s\n", v.Tag, v.Commit, v.Timestamp)
}

// Info retrieves process version.
func Info() Version {
	var (
		commit, timestamp = "unknown", "unknown"
	)

	info, ok := debug.ReadBuildInfo()
	if ok {
		commit = info.Main.Version
		for _, setting := range info.Settings {
			if setting.Key == "vcs.time" {
				timestamp = setting.Value
			}
		}
	}

	return Version{
		Tag:       tag,
		Commit:    commit,
		Timestamp: timestamp,
	}
}
