// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// StatsProvider interface for objects that provide upload/download statistics.
type StatsProvider interface {
	Stats() (chunks, blocks, bytes int64)
}

// trackProgress creates a progress bar and starts a goroutine that periodically updates it
// Returns the progress bar and a done channel that should be closed when the operation completes.
func trackProgress(cmd *cobra.Command, totalSize int64, statsProvider StatsProvider, operation string) (*progressbar.ProgressBar, chan struct{}) {
	bar := progressbar.NewOptions64(totalSize,
		progressbar.OptionSetDescription(operation+"..."),
		progressbar.OptionSetWriter(os.Stderr),
	)

	const tickerInterval = 400 * time.Millisecond

	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(tickerInterval)
		defer ticker.Stop()

		var lastBytes int64
		startTime := time.Now()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				chunks, blocks, bytes := statsProvider.Stats()
				handleProgressUpdates(cmd, bar, chunks, blocks, bytes, &lastBytes, startTime, operation)
			}
		}
	}()
	return bar, done
}

func handleProgressUpdates(
	cmd *cobra.Command,
	bar *progressbar.ProgressBar,
	chunks, blocks, bytes int64,
	lastBytes *int64,
	startTime time.Time,
	operation string,
) {

	bytesIncrement := bytes - *lastBytes
	if bytesIncrement > 0 {
		if err := bar.Add64(bytesIncrement); err != nil {
			if !strings.Contains(err.Error(), "current number exceeds max") {
				cmd.PrintErrf("failed to add bytes to progress: %v", err)
			}
		}
		*lastBytes = bytes
	}

	// Average speed in MB/s over the entire duration
	var speedMBps float64
	totalDuration := time.Since(startTime).Seconds()
	if totalDuration > 0 && bytes > 0 {
		speedMBps = float64(bytes) / (1024 * 1024) / totalDuration
	}

	var speedStr string
	if speedMBps > 0 {
		speedStr = fmt.Sprintf(", Speed: %6.2f MB/s", speedMBps)
	}

	desc := fmt.Sprintf("%s... Chunks: %3d, Blocks: %4d%s", operation, chunks, blocks, speedStr)
	bar.Describe(desc)
}
