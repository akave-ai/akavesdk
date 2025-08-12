// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package main

import (
	"flag"
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

var (
	// privateKeyTest is flag to set deployers hex private key.
	privateKeyTest = flag.String("cmd-private-key", os.Getenv("PRIVATE_KEY"), "flag to set deployers hex private key for client tests")
)

// PickPrivateKey picks hex private key of deployer.
func PickPrivateKey(t testing.TB) string {
	if *privateKeyTest == "" || strings.EqualFold(*privateKeyTest, "omit") {
		t.Skip("private key flag missing, example: -PRIVATE_KEY=<deployers hex private key>")
	}
	return *privateKeyTest
}

func TestIPCCreateBucketCommand(t *testing.T) {
	bucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	testCases := []testCase{
		{
			name:           "Create bucket successfully",
			args:           []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey, bucketName},
			expectedOutput: []string{"Bucket created"},
			expectError:    false,
		},
		{
			name:           "Create bucket already exists",
			args:           []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey, bucketName},
			expectedOutput: []string{"BucketAlreadyExists"},
			expectError:    true,
		},
		{
			name:           "Empty bucket name provided",
			args:           []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey, ""},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey},
			expectedOutput: []string{"create bucket command expects exactly 1 argument [bucket name]; got 0"},
			expectError:    true,
		},
		{
			name:           "Invalid private key provided",
			args:           []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", "51ccv2", bucketName},
			expectedOutput: []string{"Error: invalid hex character"},
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

func TestIPCDeleteBucketCommand(t *testing.T) {
	bucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey, bucketName})
	require.NoError(t, err)
	testCases := []testCase{
		{
			name:           "Delete bucket successfully",
			args:           []string{"ipc", "bucket", "delete", "--private-key", privateKey, bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Bucket deleted: Name=%s", bucketName)},
			expectError:    false,
		},
		{
			name:           "Delete not existing bucket",
			args:           []string{"ipc", "bucket", "delete", "--private-key", privateKey, bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{"failed to get bucket: not found"},
			expectError:    true,
		},
		{
			name:           "Empty bucket name provided",
			args:           []string{"ipc", "bucket", "delete", "--private-key", privateKey, "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "bucket", "delete", "--private-key", privateKey, "--node-address", nodeAddress},
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

func TestIPCViewBucketCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey, bucketName})
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name:           "View bucket successfully",
			args:           []string{"ipc", "bucket", "view", "--private-key", privateKey, bucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Name=%s", bucketName)},
			expectError:    false,
		},
		{
			name:           "View non-existent bucket",
			args:           []string{"ipc", "bucket", "view", "--private-key", privateKey, "error", "--node-address", nodeAddress},
			expectedOutput: []string{"failed to get bucket: not found"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "bucket", "view", "--private-key", privateKey, "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
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

func TestIPCListBucketsCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	// Create 3 buckets for pagination tests
	bucketNames := make([]string, 3)
	for i := range bucketNames {
		bucketNames[i] = "testcli-" + testrand.String(10)
		_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "ipc", "bucket", "create", "--private-key", privateKey, bucketNames[i]})
		assert.NoError(t, err)
	}

	t.Run("List buckets contains all test buckets", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey})
		assert.NoError(t, err)
		output := stdout + stderr
		for _, name := range bucketNames {
			assert.Contains(t, output, fmt.Sprintf("Bucket: Name=%s", name))
		}
	})

	t.Run("List buckets with limit=2 returns 2 buckets", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--offset", "0", "--limit", "2"})
		assert.NoError(t, err)
		output := stdout + stderr
		lines := strings.Split(output, "\n")
		count := 0
		for _, line := range lines {
			if strings.Contains(line, "Bucket: Name=") {
				count++
			}
		}
		assert.Equal(t, 2, count)
	})

	t.Run("Pagination returns all test buckets", func(t *testing.T) {
		bucketSet := make(map[string]struct{})
		var offset int
		limit := 2
		for {
			stdout, stderr, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--offset", fmt.Sprintf("%d", offset), "--limit", fmt.Sprintf("%d", limit)})
			assert.NoError(t, err)
			output := stdout + stderr
			lines := strings.Split(output, "\n")
			found := false
			for _, line := range lines {
				if after, ok := strings.CutPrefix(line, "Bucket: Name="); ok {
					name := strings.SplitN(after, ",", 2)[0]
					bucketSet[name] = struct{}{}
					found = true
				}
			}
			if !found {
				break
			}
			offset += limit
		}
		// Check that all test buckets are present in paginated results
		for _, name := range bucketNames {
			_, exists := bucketSet[name]
			assert.True(t, exists, "Expected bucket %s to be present in paginated results", name)
		}
	})
}

func TestIPCListFilesCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", bucketName, "--private-key", privateKey, "--node-address", nodeAddress})
	assert.NoError(t, err)

	// Upload 6 files for pagination tests
	fileNames := make([]string, 3)
	for i := range fileNames {
		file, err := createTempFile(t, 2*memory.MB.ToInt64())
		assert.NoError(t, err)
		fileNames[i] = file
		_, _, err = captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", bucketName, file, "--private-key", privateKey, "--node-address", nodeAddress})
		assert.NoError(t, err)
	}

	testCases := []testCase{
		{
			name: "List files successfully",
			args: []string{"ipc", "file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[0])),
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[1])),
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[2])),
				fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()),
				fmt.Sprintf("EncodedSize=%d", 4000480),
			},
			expectError: false,
		},
		{
			name: "List files with limit=2",
			args: []string{"ipc", "file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "0", "--limit", "2"},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[0])),
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[1])),
			},
			expectError: false,
		},
		{
			name: "List files with offset=1 and limit=2",
			args: []string{"ipc", "file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "1", "--limit", "2"},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[1])),
				fmt.Sprintf("File: Name=%s", filepath.Base(fileNames[2])),
			},
			expectError: false,
		},
		{
			name:           "List files with offset beyond range",
			args:           []string{"ipc", "file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "10", "--limit", "2"},
			expectedOutput: []string{"No files"},
			expectError:    false,
		},
		{
			name:           "List files for non-existent bucket",
			args:           []string{"ipc", "file", "list", "nonexistent-bucket", "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"No files"},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "file", "list", "--private-key", privateKey, "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
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

	t.Run("Pagination returns all test files", func(t *testing.T) {
		// List first page (limit=2)
		stdout1, stderr1, err := captureCobraOutput(rootCmd, []string{"ipc", "file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "0", "--limit", "2"})
		assert.NoError(t, err)
		filesPage1 := []string{}
		for _, fname := range fileNames {
			if strings.Contains(stdout1+stderr1, filepath.Base(fname)) {
				filesPage1 = append(filesPage1, filepath.Base(fname))
			}
		}

		// List second page (offset=2, limit=2)
		stdout2, stderr2, err := captureCobraOutput(rootCmd, []string{"ipc", "file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "2", "--limit", "2"})
		assert.NoError(t, err)
		filesPage2 := []string{}
		for _, fname := range fileNames {
			if strings.Contains(stdout2+stderr2, filepath.Base(fname)) {
				filesPage2 = append(filesPage2, filepath.Base(fname))
			}
		}

		// Combine and check for duplicates
		//nolint:gocritic
		allFiles := append(filesPage1, filesPage2...)
		fileSet := make(map[string]bool)
		for _, fname := range allFiles {
			assert.False(t, fileSet[fname], "File %s should not be duplicated", fname)
			fileSet[fname] = true
		}
		// Ensure all expected files are present
		for _, fname := range fileNames {
			assert.True(t, fileSet[filepath.Base(fname)], "File %s should be present", filepath.Base(fname))
		}
	})
}

func TestIPCFileInfoCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	_, stderr, err := captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)
	rootCID, err := extractRootCID(stderr)
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name: "File info successfully",
			args: []string{"ipc", "file", "info", "--private-key", privateKey, bucketName, filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s, RootCID=%s", filepath.Base(file), rootCID),
				fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()),
			},
			expectError: false,
		},
		{
			name:           "File info for non-existent file",
			args:           []string{"ipc", "file", "info", "--private-key", privateKey, bucketName, "nonexistent-file", "--node-address", nodeAddress},
			expectedOutput: []string{"file not exists"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "file", "info", "--private-key", privateKey, "", filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"ipc", "file", "info", "--private-key", privateKey, bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
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

func TestIPCFileUploadCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)
	file2, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name: "File upload successfully",
			args: []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--disable-erasure-coding"},
			expectedOutput: []string{
				fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file)), fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()),
				fmt.Sprintf("EncodedSize=%d", 2000028),
			},
			expectError: false,
		},
		{
			name:           "File upload successfully with erasure coding",
			args:           []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file2, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file2))},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "file", "upload", "--private-key", privateKey, "", file, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File path not provided",
			args:           []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, "", "--node-address", nodeAddress},
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

func TestIPCFileUploadCommandWithEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	file2, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name:           "File upload successfully",
			args:           []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey, "--disable-erasure-coding"},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file))},
			expectError:    false,
		},
		{
			name:           "File upload successfully with erasure coding",
			args:           []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file2, "--node-address", nodeAddress, "--encryption-key", encKey},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", filepath.Base(file2))},
			expectError:    false,
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

func TestIPCFileDownloadCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file))},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, "", filepath.Base(file), tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, bucketName, "", tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
		{
			name:           "Destination path not provided",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, bucketName, filepath.Base(file), "", "--node-address", nodeAddress},
			expectedOutput: []string{"destination path is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
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

func TestIPCFileDownloadCommandWithErasureCoding(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file))},
			expectError:    false,
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

func TestIPCFileDownloadCommandWithEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey})
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress, "-e", encryptionKey},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file))},
			expectError:    false,
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

func TestIPCFileDownloadCommandWithEncryptionAndErasureCoding(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey})
	assert.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	assert.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"ipc", "file", "download", "--private-key", privateKey, bucketName, filepath.Base(file), tempDir, "--node-address", nodeAddress, "--encryption-key", encKey},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", filepath.Base(file))},
			expectError:    false,
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

func TestIPCFileDeleteCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"ipc", "bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	assert.NoError(t, err)

	file, err := createTempFile(t, 2*memory.MB.ToInt64())
	assert.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"ipc", "file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name:           "File delete successfully",
			args:           []string{"ipc", "file", "delete", "--private-key", privateKey, bucketName, filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File successfully deleted: Name=%s\n", filepath.Base(file))},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"ipc", "file", "delete", "--private-key", privateKey, "", filepath.Base(file), "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"ipc", "file", "delete", "--private-key", privateKey, bucketName, "", "--node-address", nodeAddress},
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

func TestIPCWalletBalanceCommand(t *testing.T) {
	// Create temporary directory for test wallets
	tmpKeystoreDir := t.TempDir()
	nodeAddress := PickNodeRPCAddress(t)

	// Use the first anvil default account private key for testing
	// This is the well-known first default account from anvil that has initial balance
	testPrivateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	walletName := testrand.String(10)

	// Import the test wallet with known private key
	stdout, stderr, err := captureCobraOutput(rootCmd, []string{
		"wallet", "import", walletName, testPrivateKey, "--keystore", tmpKeystoreDir,
	})
	require.NoError(t, err)
	require.Contains(t, stdout+stderr, "Wallet imported successfully")

	testCases := []testCase{
		{
			name: "Get wallet balance successfully",
			args: []string{
				"--node-address", nodeAddress,
				"wallet", "balance", walletName,
				"--keystore", tmpKeystoreDir,
			},
			expectedOutput: []string{"Wallet:", walletName, "Address:", "Balance:", "AKVT"},
			expectError:    false,
		},
		{
			name: "Get balance for default wallet (no wallet name specified)",
			args: []string{
				"--node-address", nodeAddress,
				"wallet", "balance",
				"--keystore", tmpKeystoreDir,
			},
			expectedOutput: []string{"Wallet:", walletName, "Address:", "Balance:", "AKVT"},
			expectError:    false,
		},
		{
			name: "Get balance for non-existent wallet",
			args: []string{
				"--node-address", nodeAddress,
				"wallet", "balance", "non-existent-wallet",
				"--keystore", tmpKeystoreDir,
			},
			expectedOutput: []string{"failed to read wallet file"},
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
