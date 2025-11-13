// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package retry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/retry"
)

func TestWithRetry(t *testing.T) {
	ctx := t.Context()

	t.Run("success on first attempt", func(t *testing.T) {
		retry := retry.WithRetry{MaxAttempts: 3, BaseDelay: 10 * time.Millisecond}
		callCount := 0
		err := retry.Do(ctx, func() (bool, error) {
			callCount++
			return false, nil
		})
		require.NoError(t, err)
		require.Equal(t, 1, callCount)
	})

	t.Run("failure without retry", func(t *testing.T) {
		retry := retry.WithRetry{MaxAttempts: 3, BaseDelay: 10 * time.Millisecond}
		callCount := 0
		testErr := errors.New("test error")
		err := retry.Do(ctx, func() (bool, error) {
			callCount++
			return false, testErr
		})
		require.Error(t, err)
		require.Equal(t, testErr, err)
		require.Equal(t, 1, callCount)
	})

	t.Run("retry and success", func(t *testing.T) {
		retry := retry.WithRetry{MaxAttempts: 3, BaseDelay: 10 * time.Millisecond}
		callCount := 0
		testErr := errors.New("test error")
		err := retry.Do(ctx, func() (bool, error) {
			callCount++
			if callCount < 3 {
				return true, testErr
			}
			return false, nil
		})
		require.NoError(t, err)
		require.Equal(t, 3, callCount)
	})

	t.Run("retry exceeds max attempts", func(t *testing.T) {
		retry := retry.WithRetry{MaxAttempts: 2, BaseDelay: 10 * time.Millisecond}
		callCount := 0
		testErr := errors.New("test error")
		err := retry.Do(ctx, func() (bool, error) {
			callCount++
			return true, testErr
		})
		require.Error(t, err)
		require.Equal(t, testErr, err)
		require.Equal(t, 3, callCount) // initial + 2 retries
	})

	t.Run("context cancellation", func(t *testing.T) {
		retry := retry.WithRetry{MaxAttempts: 5, BaseDelay: 100 * time.Millisecond}
		ctx, cancel := context.WithCancel(ctx)
		callCount := 0
		testErr := errors.New("test error")
		go func() {
			time.Sleep(300 * time.Millisecond)
			cancel()
		}()
		err := retry.Do(ctx, func() (bool, error) {
			callCount++
			return true, testErr
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "retry aborted")
		require.Greater(t, callCount, 1)
	})

	t.Run("maxAttempts zero", func(t *testing.T) {
		retry := retry.WithRetry{MaxAttempts: 0, BaseDelay: 10 * time.Millisecond}
		callCount := 0
		testErr := errors.New("test error")
		err := retry.Do(ctx, func() (bool, error) {
			callCount++
			return true, testErr
		})
		require.Error(t, err)
		require.Equal(t, testErr, err)
		require.Equal(t, 1, callCount)
	})
}
