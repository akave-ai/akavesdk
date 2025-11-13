// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package httpext_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/httpext"
)

func TestRangeDownload(t *testing.T) {
	data := []byte("Hello, World! This is some test data for the client.")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rangeHeader := r.Header.Get("Range")
		if rangeHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var start, end int
		_, err := fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if start < 0 || end >= len(data) || start > end {
			w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
			return
		}
		w.WriteHeader(http.StatusPartialContent)
		_, _ = w.Write(data[start : end+1])
	}))
	defer server.Close()

	ctx := t.Context()

	// Test full download
	result, err := httpext.RangeDownload(ctx, http.DefaultClient, server.URL, 0, int64(len(data)))
	require.NoError(t, err)
	assert.Equal(t, data, result)

	// Test partial download
	partial, err := httpext.RangeDownload(ctx, http.DefaultClient, server.URL, 7, 5) // "World"
	require.NoError(t, err)
	assert.Equal(t, []byte("World"), partial)

	// Test invalid offset
	_, err = httpext.RangeDownload(ctx, http.DefaultClient, server.URL, -1, 5)
	assert.Error(t, err)

	// Test invalid length
	_, err = httpext.RangeDownload(ctx, http.DefaultClient, server.URL, 0, 0)
	assert.Error(t, err)

	// Test length <= 0
	_, err = httpext.RangeDownload(ctx, http.DefaultClient, server.URL, 0, -1)
	assert.Error(t, err)
}
