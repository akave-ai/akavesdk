// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestStreamingCreateBucketCommand(t *testing.T) {
	bucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	testCases := []testCase{
		{
			name:           "Create bucket successfully",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", bucketName},
			expectedOutput: []string{"Bucket created"},
			expectError:    false,
		},
		{
			name:           "Create bucket already exists",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", bucketName},
			expectedOutput: []string{"failed to create bucket: sdk: rpc error: code = AlreadyExists"},
			expectError:    true,
		},
		{
			name:           "Empty bucket name provided",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", ""},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"--node-address", nodeAddress, "bucket", "create"},
			expectedOutput: []string{"Error: create bucket command expects exactly 1 argument [bucket name]; got 0"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingDeleteBucketCommand(t *testing.T) {
	bucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)
	testCases := []testCase{
		{
			name:           "Delete bucket successfully",
			args:           []string{"bucket", "delete", bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Bucket deleted: Name=%s", bucketName)},
			expectError:    false,
		},
		{
			name:           "Delete not existing bucket",
			args:           []string{"bucket", "delete", bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{"failed to delete bucket: sdk: rpc error: code = NotFound"},
			expectError:    true,
		},
		{
			name:           "Empty bucket name provided",
			args:           []string{"bucket", "delete", "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"bucket", "delete", "--node-address", nodeAddress},
			expectedOutput: []string{"Error: delete bucket command expects exactly 1 argument [bucket name]; got 0"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingViewBucketCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name:           "View bucket successfully",
			args:           []string{"bucket", "view", bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Bucket: Name=%s", bucketName)},
			expectError:    false,
		},
		{
			name:           "View non-existent bucket",
			args:           []string{"bucket", "view", "error", "--node-address", nodeAddress},
			expectedOutput: []string{"failed to get bucket: sdk: rpc error: code = NotFound"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"bucket", "view", "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingListBucketsCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName1 := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName1, "--node-address", nodeAddress})
	assert.NoError(t, err)

	bucketName2 := testrand.String(10)
	assert.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName2, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCase := testCase{
		name:           "List buckets successfully",
		args:           []string{"bucket", "list", "--node-address", nodeAddress},
		expectedOutput: []string{fmt.Sprintf("Bucket: Name=%s", bucketName1), fmt.Sprintf("Bucket: Name=%s", bucketName2)},
		expectError:    false,
	}

	t.Run(testCase.name, func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, testCase.args)

		if testCase.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		for _, expected := range testCase.expectedOutput {
			assert.Contains(t, stdout+stderr, expected)
		}
	})
}

func TestStreamingListFilesCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file1, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	file2, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file1, "--node-address", nodeAddress})
	assert.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file2, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCases := []testCase{
		// Streaming API
		{
			name:           "List files successfully",
			args:           []string{"files-streaming", "list", bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File: Name=%s", filepath.Base(file1)), fmt.Sprintf("File: Name=%s", filepath.Base(file2))},
			expectError:    false,
		},
		{
			name:           "List files for non-existent bucket",
			args:           []string{"files-streaming", "list", "nonexistent-bucket", "--node-address", nodeAddress},
			expectedOutput: []string{"failed to list files: sdk: rpc error: code = NotFound"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"files-streaming", "list", "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingFileVersionsCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file1, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file1, "--node-address", nodeAddress})
	assert.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file1, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCases := []testCase{
		// Streaming API
		{
			name:           "List file versions successfully",
			args:           []string{"files-streaming", "versions", bucketName, filepath.Base(file1), "--node-address", nodeAddress},
			expectedOutput: []string{"Version: RootCID="},
			expectError:    false,
		},
		{
			name:           "List file versions for non-existent file",
			args:           []string{"files-streaming", "versions", bucketName, "foobar", "--node-address", nodeAddress},
			expectedOutput: []string{"failed to get file versions: sdk: rpc error: code = NotFound"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				result := stdout + stderr
				if !tc.expectError {
					lines := strings.Split(strings.TrimSpace(result), "\n")
					assert.Len(t, lines, 2)
				}
				assert.Contains(t, result, expected)
			}
		})
	}
}

func TestStreamingFileInfoCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	_, stderr, err := captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)
	rootCID, err := extractRootCID(stderr)
	assert.NoError(t, err)

	testCases := []testCase{
		// Streaming API
		{
			name:           "File info successfully",
			args:           []string{"files-streaming", "info", bucketName, filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File: Name=%s, RootCID=%s", filepath.Base(file), rootCID)},
			expectError:    false,
		},
		{
			name:           "File info for non-existent file",
			args:           []string{"files-streaming", "info", bucketName, "nonexistent-file", "--node-address", nodeAddress},
			expectedOutput: []string{"failed to get file info: sdk: rpc error: code = NotFound"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"files-streaming", "info", "", filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"files-streaming", "info", bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingFileUploadCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	file2, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name: "File uploaded successfully",
			args: []string{"files-streaming", "upload", bucketName, file, "--node-address", nodeAddress, "--disable-erasure-coding"},
			expectedOutput: []string{
				fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file)),
				"RootCID=",
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
		{
			name: "File uploaded successfully with erasure coding",
			args: []string{"files-streaming", "upload", bucketName, file2, "--node-address", nodeAddress},
			expectedOutput: []string{
				fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file2)),
				"RootCID=",
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"files-streaming", "upload", "", file, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File path not provided",
			args:           []string{"files-streaming", "upload", bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file path is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingFileUploadCommandWithEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	file2, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name: "File uploaded successfully",
			args: []string{"files-streaming", "upload", bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey, "--disable-erasure-coding"},
			expectedOutput: []string{
				fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file)),
				"RootCID=",
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
		{
			name: "File upload successfully with erasure coding",
			args: []string{"files-streaming", "upload", bucketName, file2, "--node-address", nodeAddress, "--encryption-key", encKey},
			expectedOutput: []string{
				fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file2)),
				"RootCID=",
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingFileDownloadCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name: "File download successfully",
			args: []string{"files-streaming", "download", bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{
				fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file)),
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
		{
			name: "File download successfully from downloadV2",
			args: []string{"files-streaming", "download", bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress, "--filecoin"},
			expectedOutput: []string{
				fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file)),
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"files-streaming", "download", "", filepath.Base(file), tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"files-streaming", "download", bucketName, "", tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
		{
			name:           "Destination path not provided",
			args:           []string{"files-streaming", "download", bucketName, filepath.Base(file), "", "--node-address", nodeAddress},
			expectedOutput: []string{"destination path is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingFileDownloadCommandWithEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey})
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name: "File download successfully",
			args: []string{"files-streaming", "download", bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress, "--encryption-key", encKey},
			expectedOutput: []string{
				fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file)),
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},

		{
			name: "File download successfully from downloadV2",
			args: []string{"files-streaming", "download", bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress, "--filecoin", "--encryption-key", encKey},
			expectedOutput: []string{
				fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file)),
				fmt.Sprintf("Size=%d", 2*memory.MB.ToInt64()),
				"TransferredSize=",
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestStreamingFileDeleteCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"files-streaming", "upload", bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCases := []testCase{
		// Streaming API
		{
			name:           "File delete successfully",
			args:           []string{"files-streaming", "delete", bucketName, filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File deleted successfully: Name=%s", filepath.Base(file))},
			expectError:    false,
		},
		{
			name:           "File delete for non-existent file",
			args:           []string{"files-streaming", "delete", bucketName, "nonexistent-file", "--node-address", nodeAddress},
			expectedOutput: []string{"failed to delete file: sdk: rpc error: code = NotFound"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"files-streaming", "delete", "", filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"files-streaming", "delete", bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				assert.Contains(t, stdout+stderr, expected)
			}
		})
	}
}
