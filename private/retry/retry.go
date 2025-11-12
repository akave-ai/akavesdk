// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package retry provides functionality to retry operations with exponential backoff.
package retry

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

// WithRetry defines retry behavior.
type WithRetry struct {
	MaxAttempts int
	BaseDelay   time.Duration
}

// Do calls function until success, max retries reached or operation is cancelled with exponential backoff.
func (retry WithRetry) Do(ctx context.Context, f func() (bool, error)) error {
	for attempt := 0; ; attempt++ {
		needsRetry, err := f()
		if err == nil {
			return nil
		}
		if !needsRetry || attempt >= retry.MaxAttempts {
			return err
		}

		backoff := retry.BaseDelay * (1 << attempt)
		jitter := time.Duration(rand.Int64N(int64(retry.BaseDelay)))
		delay := backoff + jitter

		select {
		case <-ctx.Done():
			return errors.Join(fmt.Errorf("retry aborted: %w", ctx.Err()), err)
		case <-time.After(delay):
		}
	}
}
