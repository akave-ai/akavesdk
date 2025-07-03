// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestSkipToPosition(t *testing.T) {
	t.Run("non-resumable upload should return nil without doing anything", func(t *testing.T) {
		// Create a non-resumable file upload (chunkCount = 0)
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		reader := strings.NewReader("test data")

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err)

		buf := make([]byte, 4)
		n, _ := reader.Read(buf)
		require.Equal(t, 4, n)
		require.Equal(t, "test", string(buf))
	})

	t.Run("seekable reader should use Seek method", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 5 // Skip first 5 bytes

		testData := []byte("0123456789abcdef")
		reader := bytes.NewReader(testData)

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err)

		// Verify reading continues from the correct position (after seeking)
		buf := make([]byte, 3)
		n, _ := reader.Read(buf)
		require.Equal(t, 3, n)
		require.Equal(t, "567", string(buf))
	})

	t.Run("seekable reader beyond end should result in EOF on read", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 100 // Skip more bytes than available

		reader := strings.NewReader("hello") // Only 5 bytes available

		// This should succeed (seeking beyond end is allowed)
		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err)

		// But subsequent reads should return EOF immediately
		buf := make([]byte, 1)
		n, readErr := reader.Read(buf)
		require.Equal(t, 0, n)
		require.Equal(t, io.EOF, readErr)
	})

	t.Run("non-seekable reader should use CopyN to discard", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 5 // Skip first 5 bytes

		// Use io.LimitReader which doesn't implement io.Seeker
		baseReader := strings.NewReader("0123456789abcdef")
		reader := io.LimitReader(baseReader, 10) // Limit to 16 bytes

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err)

		// Verify reading continues from the correct position
		buf := make([]byte, 3)
		n, readErr := reader.Read(buf)
		require.NoError(t, readErr)
		require.Equal(t, 3, n)
		require.Equal(t, "567", string(buf))
	})

	t.Run("non-seekable reader with EOF during skip should not return error", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 10 // Try to skip 10 bytes

		// Use io.LimitReader which limits to 5 bytes (doesn't implement io.Seeker)
		baseReader := strings.NewReader("0123456789abcdef")
		reader := io.LimitReader(baseReader, 5) // Only 5 bytes available

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err) // EOF should not cause an error

		// Further reads should return EOF immediately
		buf := make([]byte, 1)
		n, readErr := reader.Read(buf)
		require.Equal(t, 0, n)
		require.Equal(t, io.EOF, readErr)
	})

	t.Run("edge case with exact skip amount", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 5 // Skip exactly 5 bytes

		// Use io.LimitReader to create a reader with exactly 5 bytes
		baseReader := strings.NewReader("0123456789")
		reader := io.LimitReader(baseReader, 5) // Exactly 5 bytes available

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err) // Should succeed

		// Further reads should return EOF immediately since all bytes were consumed
		buf := make([]byte, 1)
		n, readErr := reader.Read(buf)
		require.Equal(t, 0, n)
		require.Equal(t, io.EOF, readErr)
	})

	t.Run("large skip amount with seekable reader", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 1 << 20 // Skip 1MB

		testData := testrand.BytesD(t, 0, 2<<20) // 2MB
		bytesAfterSkip := make([]byte, 4)        // To read after skipping
		copy(bytesAfterSkip, testData[1<<20:1<<20+4])

		reader := bytes.NewReader(testData)

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err)

		// Verify reading continues from the correct position (after seeking)
		buf := make([]byte, 4)
		n, _ := reader.Read(buf)
		require.Equal(t, 4, n)
		require.Equal(t, bytesAfterSkip, buf)
	})

	t.Run("large skip amount with non-seekable reader", func(t *testing.T) {
		fileUpload, err := NewIPCFileUpload("test-bucket", "test-file")
		require.NoError(t, err)

		// Make it resumable
		fileUpload.state.chunkCount = 1
		fileUpload.state.actualFileSize = 64 * 1024 // Skip 64KB

		// Create test data using testrand
		testData := testrand.BytesD(t, 0, 128*1024) // 128KB

		// Use io.LimitReader which doesn't implement io.Seeker
		baseReader := bytes.NewReader(testData)
		reader := io.LimitReader(baseReader, 128*1024)

		err = skipToPosition(reader, fileUpload.state.actualFileSize)
		require.NoError(t, err)

		// Verify reading continues from the correct position
		buf := make([]byte, 4)
		n, readErr := reader.Read(buf)
		require.NoError(t, readErr)
		require.Equal(t, 4, n)
		// Verify we read the expected bytes after skipping
		expected := testData[64*1024 : 64*1024+4]
		require.Equal(t, expected, buf)
	})
}
