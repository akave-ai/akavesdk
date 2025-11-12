// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package httpext provides utility functions that are built on top of http client.
package httpext

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// RangeDownload downloads a specific byte range from the given URL using the provided HTTP client.
func RangeDownload(ctx context.Context, client *http.Client, url string, offset, length int64) (_ []byte, err error) {
	if length <= 0 || offset < 0 {
		return nil, fmt.Errorf("length must be positive and offset must be non-negative")
	}

	end := offset + length - 1

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", offset, end))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		err = errors.Join(err, resp.Body.Close())
	}()

	// some CDNs may return 200 OK for range requests
	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		return nil, errors.Join(err, fmt.Errorf("download failed with status %d: %s", resp.StatusCode, string(body)))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}
