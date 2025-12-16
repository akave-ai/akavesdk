// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/pdptest"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestExternalCreateBucketCommand(t *testing.T) {
	bucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	testCases := []testCase{
		{
			name:           "Create bucket successfully",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, bucketName},
			expectedOutput: []string{"Bucket created", fmt.Sprintf("Name=%s", bucketName)},
			expectError:    false,
		},
		{
			name:           "Create bucket successfully with metadata encryption",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, bucketName},
			expectedOutput: []string{"Bucket created", fmt.Sprintf("Name=%s", bucketName)},
			expectError:    false,
		},
		{
			name:           "Create bucket already exists",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, bucketName},
			expectedOutput: []string{"BucketAlreadyExists"},
			expectError:    true,
		},
		{
			name:           "Empty bucket name provided",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, ""},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey},
			expectedOutput: []string{"create bucket command expects exactly 1 argument [bucket name]; got 0"},
			expectError:    true,
		},
		{
			name:           "Invalid private key provided",
			args:           []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", "51ccv2", bucketName},
			expectedOutput: []string{"Error: invalid hex character"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalDeleteBucketCommand(t *testing.T) {
	firstBucketName := testrand.String(10)
	secondBucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, firstBucketName})
	require.NoError(t, err)
	_, _, err = captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, secondBucketName})
	require.NoError(t, err)

	testCases := []testCase{
		{
			name:           "Delete non encrypted bucket with encryption",
			args:           []string{"bucket", "delete", "--private-key", privateKey, firstBucketName, "--node-address", nodeAddress, "--metadata-encryption", "--encryption-key", encKey},
			expectedOutput: []string{"BucketNotFound"},
			expectError:    true,
		},
		{
			name:           "Delete non encrypted bucket without encryption",
			args:           []string{"bucket", "delete", "--private-key", privateKey, firstBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Bucket deleted: Name=%s", firstBucketName)},
			expectError:    false,
		},
		{
			name:           "Delete encrypted bucket without encryption",
			args:           []string{"bucket", "delete", "--private-key", privateKey, secondBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{"BucketNotFound"},
			expectError:    true,
		},
		{
			name:           "Delete encrypted bucket with encryption",
			args:           []string{"bucket", "delete", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, secondBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Bucket deleted: Name=%s", secondBucketName)},
			expectError:    false,
		},
		{
			name:           "Delete not existing bucket",
			args:           []string{"bucket", "delete", "--private-key", privateKey, firstBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{"BucketNotFound"},
			expectError:    true,
		},
		{
			name:           "Empty bucket name provided",
			args:           []string{"bucket", "delete", "--private-key", privateKey, "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"bucket", "delete", "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"Error: create bucket command expects exactly 1 argument [bucket name]; got 0"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalViewBucketCommand(t *testing.T) {
	firstBucketName := testrand.String(10)
	secondBucketName := testrand.String(10)
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, firstBucketName})
	require.NoError(t, err)

	_, _, err = captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, secondBucketName})
	require.NoError(t, err)

	testCases := []testCase{
		{
			name:           "View non encrypted bucket without encryption",
			args:           []string{"bucket", "view", "--private-key", privateKey, firstBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Name=%s", firstBucketName)},
			expectError:    false,
		},
		{
			name:           "View non encrypted bucket with encryption",
			args:           []string{"bucket", "view", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, firstBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{"BucketNotFound"},
			expectError:    true,
		},
		{
			name:           "View encrypted bucket without encryption",
			args:           []string{"bucket", "view", "--private-key", privateKey, secondBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{"BucketNotFound"},
			expectError:    true,
		},
		{
			name:           "View encrypted bucket with encryption",
			args:           []string{"bucket", "view", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, secondBucketName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("Name=%s", secondBucketName)},
			expectError:    false,
		},
		{
			name:           "View non-existent bucket",
			args:           []string{"bucket", "view", "--private-key", privateKey, "error", "--node-address", nodeAddress},
			expectedOutput: []string{"BucketNotFound"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"bucket", "view", "--private-key", privateKey, "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalListBucketsCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	// Create 3 buckets for pagination tests
	bucketNames := make([]string, 3)
	for i := range bucketNames {
		bucketNames[i] = fmt.Sprintf("testcli-%s-%d", testrand.String(5), time.Now().UnixNano())
		_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "bucket", "create", "--private-key", privateKey, bucketNames[i]})
		require.NoError(t, err)
	}

	t.Run("List buckets contains all test buckets", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--limit", "200"})
		require.NoError(t, err)
		output := stdout + stderr
		for _, name := range bucketNames {
			require.Contains(t, output, fmt.Sprintf("Bucket: Name=%s", name))
		}
	})

	t.Run("List buckets with limit=2 returns 2 buckets", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--offset", "0", "--limit", "2"})
		require.NoError(t, err)
		output := stdout + stderr
		lines := strings.Split(output, "\n")
		count := 0
		for _, line := range lines {
			if strings.Contains(line, "Bucket: Name=") {
				count++
			}
		}
		require.Equal(t, 2, count)
	})

	t.Run("Pagination returns all test buckets", func(t *testing.T) {
		bucketSet := make(map[string]struct{})
		var offset int
		limit := 2
		for {
			stdout, stderr, err := captureCobraOutput(rootCmd, []string{"bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--offset", fmt.Sprintf("%d", offset), "--limit", fmt.Sprintf("%d", limit)})
			require.NoError(t, err)
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
			require.True(t, exists, "Expected bucket %s to be present in paginated results", name)
		}
	})
}

func TestExternalListBucketsCommandWithMetdadataEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	// Create 3 buckets for pagination tests with metadata encryption
	bucketNames := make([]string, 3)
	for i := range bucketNames {
		bucketNames[i] = fmt.Sprintf("testcli-enc-%s-%d", testrand.String(5), time.Now().UnixNano())
		bucketName := bucketNames[i]
		_, _, err := captureCobraOutput(rootCmd, []string{"--node-address", nodeAddress, "bucket", "create", bucketName, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey})
		require.NoError(t, err)
	}

	t.Run("List buckets contains all test buckets", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(
			rootCmd,
			[]string{"bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, "--limit", "200"},
		)
		require.NoError(t, err)
		output := stdout + stderr

		for _, name := range bucketNames {
			require.Contains(t, output, fmt.Sprintf("Bucket: Name=%s", name))
		}
	})

	t.Run("List buckets with limit=2 returns 2 buckets", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, "--offset", "0", "--limit", "2"})
		require.NoError(t, err)
		output := stdout + stderr
		lines := strings.Split(output, "\n")
		count := 0
		for _, line := range lines {
			if strings.Contains(line, "Bucket: Name=") {
				count++
			}
		}
		require.Equal(t, 2, count)
	})

	t.Run("Pagination returns all test buckets", func(t *testing.T) {
		bucketSet := make(map[string]struct{})
		var offset int
		limit := 2
		for {
			stdout, stderr, err := captureCobraOutput(rootCmd, []string{"bucket", "list", "--node-address", nodeAddress, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, "--offset", fmt.Sprintf("%d", offset), "--limit", fmt.Sprintf("%d", limit)})
			require.NoError(t, err)
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
			require.True(t, exists, "Expected bucket %s to be present in paginated results", name)
		}
	})
}

func TestExternalListFilesCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := "testbucket-" + testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--private-key", privateKey, "--node-address", nodeAddress})
	require.NoError(t, err)

	// Upload 3 files for pagination tests
	fileNames := make([]string, 3)
	for i := range fileNames {
		file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
		require.NoError(t, err)
		fileNames[i] = filepath.Base(file)
		_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", bucketName, file, "--private-key", privateKey, "--node-address", nodeAddress})
		require.NoError(t, err)
	}

	testCases := []testCase{
		{
			name: "List files successfully",
			args: []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s", fileNames[0]),
				fmt.Sprintf("File: Name=%s", fileNames[1]),
				fmt.Sprintf("File: Name=%s", fileNames[2]),
				fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()),
				fmt.Sprintf("EncodedSize=%d", 4000480),
			},
			expectError: false,
		},
		{
			name: "List files with limit=2",
			args: []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "0", "--limit", "2"},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s", fileNames[0]),
				fmt.Sprintf("File: Name=%s", fileNames[1]),
			},
			expectError: false,
		},
		{
			name: "List files with offset=1 and limit=2",
			args: []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "1", "--limit", "2"},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s", fileNames[1]),
				fmt.Sprintf("File: Name=%s", fileNames[2]),
			},
			expectError: false,
		},
		{
			name:           "List files with offset beyond range",
			args:           []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "10", "--limit", "2"},
			expectedOutput: []string{"No files"},
			expectError:    false,
		},
		{
			name:           "List files for non-existent bucket",
			args:           []string{"file", "list", "nonexistent-bucket", "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"No files"},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"file", "list", "--private-key", privateKey, "", "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}

	t.Run("Pagination returns all test files", func(t *testing.T) {
		// List first page (limit=2)
		stdout1, stderr1, err := captureCobraOutput(rootCmd, []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "0", "--limit", "2"})
		require.NoError(t, err)
		filesPage1 := []string{}
		for _, fname := range fileNames {
			if strings.Contains(stdout1+stderr1, filepath.Base(fname)) {
				filesPage1 = append(filesPage1, filepath.Base(fname))
			}
		}

		// List second page (offset=2, limit=2)
		stdout2, stderr2, err := captureCobraOutput(rootCmd, []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress, "--offset", "2", "--limit", "2"})
		require.NoError(t, err)
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
			require.False(t, fileSet[fname], "File %s should not be duplicated", fname)
			fileSet[fname] = true
		}
		// Ensure all expected files are present
		for _, fname := range fileNames {
			require.True(t, fileSet[filepath.Base(fname)], "File %s should be present", filepath.Base(fname))
		}
	})
}

func TestExternalListFilesCommandWithMetadataEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := "testfile-enc-" + testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", bucketName, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, "--node-address", nodeAddress})
	require.NoError(t, err)

	// Upload 3 files to encrypted bucket
	fileNames := make([]string, 3)
	for i := range 3 {
		file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
		require.NoError(t, err)
		fileNames[i] = filepath.Base(file)

		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"file", "upload", bucketName, file, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, "--node-address", nodeAddress})
		require.NoError(t, err, "Failed to upload file %s: %s%s", fileNames[i], stdout, stderr)
	}

	t.Run("List bucket files without metadata encryption flag", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"file", "list", bucketName, "--private-key", privateKey, "--node-address", nodeAddress})
		require.NoError(t, err)
		require.Contains(t, stdout+stderr, "No files")
	})

	t.Run("List bucket files with metadata encryption flag", func(t *testing.T) {
		stdout, stderr, err := captureCobraOutput(rootCmd, []string{"file", "list", bucketName, "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, "--node-address", nodeAddress, "--limit", "200"})
		require.NoError(t, err)
		output := stdout + stderr

		for _, name := range fileNames {
			require.Contains(t, output, fmt.Sprintf("File: Name=%s", name))
		}
		require.Contains(t, output, fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()))
	})
}

func TestExternalFileInfoCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	encryptedBucketName := testrand.String(10)
	_, _, err = captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file1, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file1)
	_, stderr, err := captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file1, "--node-address", nodeAddress})
	require.NoError(t, err)
	rootCID, err := extractRootCID(stderr)
	require.NoError(t, err)

	file2, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	file2Name := filepath.Base(file2)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, file2, "--node-address", nodeAddress})
	require.NoError(t, err)

	testCases := []testCase{
		{
			name: "File info successfully",
			args: []string{"file", "info", "--private-key", privateKey, bucketName, fileName, "--node-address", nodeAddress},
			expectedOutput: []string{
				fmt.Sprintf("File: Name=%s, RootCID=%s", fileName, rootCID),
				fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()),
			},
			expectError: false,
		},
		{
			name: "File info for encrypted bucket without metadata encryption flag",
			args: []string{"file", "info", "--private-key", privateKey, encryptedBucketName, file2Name, "--node-address", nodeAddress},
			expectedOutput: []string{
				"BucketNotFound",
			},
			expectError: true,
		},
		{
			name: "File info for non-encrypted bucket with metadata encryption flag",
			args: []string{"file", "info", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, bucketName, fileName, "--node-address", nodeAddress},
			expectedOutput: []string{
				"BucketNotFound",
			},
			expectError: true,
		},
		{
			name:           "File info for non-existent file",
			args:           []string{"file", "info", "--private-key", privateKey, bucketName, "nonexistent-file", "--node-address", nodeAddress},
			expectedOutput: []string{"file not exists"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"file", "info", "--private-key", privateKey, "", fileName, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"file", "info", "--private-key", privateKey, bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileUploadCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)
	file2, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	file2Name := filepath.Base(file2)

	testCases := []testCase{
		{
			name: "File upload successfully",
			args: []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--disable-erasure-coding"},
			expectedOutput: []string{
				fmt.Sprintf("File uploaded successfully: Name=%s", fileName), fmt.Sprintf("ActualSize=%d", 2*memory.MB.ToInt64()),
				fmt.Sprintf("EncodedSize=%d", 2000028),
			},
			expectError: false,
		},
		{
			name:           "File upload successfully with erasure coding",
			args:           []string{"file", "upload", "--private-key", privateKey, bucketName, file2, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", file2Name)},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"file", "upload", "--private-key", privateKey, "", file, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File path not provided",
			args:           []string{"file", "upload", "--private-key", privateKey, bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file path is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileUploadCommandWithMetadataEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	testCases := []testCase{
		{
			name:           "File upload successfully with metadata encryption",
			args:           []string{"file", "upload", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, bucketName, file, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileUploadCommandWithEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	file2, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	file2Name := filepath.Base(file2)

	testCases := []testCase{
		{
			name:           "File upload successfully",
			args:           []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey, "--disable-erasure-coding"},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
		{
			name:           "File upload successfully with erasure coding",
			args:           []string{"file", "upload", "--private-key", privateKey, bucketName, file2, "--node-address", nodeAddress, "--encryption-key", encKey},
			expectedOutput: []string{fmt.Sprintf("File uploaded successfully: Name=%s", file2Name)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDownloadCommand(t *testing.T) {
	// TODO: add test case for archival mode

	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	require.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		require.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"file", "download", "--private-key", privateKey, bucketName, fileName, tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"file", "download", "--private-key", privateKey, "", fileName, tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"file", "download", "--private-key", privateKey, bucketName, "", tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
		{
			name:           "Destination path not provided",
			args:           []string{"file", "download", "--private-key", privateKey, bucketName, fileName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"destination path is required"},
			expectError:    true,
		},
		// TODO: add case with no private key
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDownloadCommandWithMetadataEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	// Create encrypted bucket and upload file with metadata encryption
	encryptedBucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, file, "--node-address", nodeAddress})
	require.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		require.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully with metadata encryption",
			args:           []string{"file", "download", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, fileName, tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDownloadCommandWithErasureCoding(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	require.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		require.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"file", "download", "--private-key", privateKey, bucketName, fileName, tempDir, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDownloadCommandWithEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey})
	require.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		require.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"file", "download", "--private-key", privateKey, bucketName, fileName, tempDir, "--node-address", nodeAddress, "-e", encryptionKey},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDownloadCommandWithEncryptionAndErasureCoding(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress, "--encryption-key", encKey})
	require.NoError(t, err)

	tempDir, err := os.MkdirTemp("", "test-download")
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		require.NoError(t, err)
	}()

	testCases := []testCase{
		{
			name:           "File download successfully",
			args:           []string{"file", "download", "--private-key", privateKey, bucketName, fileName, tempDir, "--node-address", nodeAddress, "--encryption-key", encKey},
			expectedOutput: []string{fmt.Sprintf("File downloaded successfully: Name=%s", fileName)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDeleteCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	require.NoError(t, err)

	testCases := []testCase{
		{
			name:           "File delete successfully",
			args:           []string{"file", "delete", "--private-key", privateKey, bucketName, fileName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File successfully deleted: Name=%s\n", fileName)},
			expectError:    false,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"file", "delete", "--private-key", privateKey, "", fileName, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"file", "delete", "--private-key", privateKey, bucketName, "", "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalFileDeleteCommandWithMetadataEncryption(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)

	encryptedBucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, file, "--node-address", nodeAddress})
	require.NoError(t, err)

	testCases := []testCase{
		{
			name:           "File delete successfully with metadata encryption",
			args:           []string{"file", "delete", "--private-key", privateKey, "--metadata-encryption", "--encryption-key", encKey, encryptedBucketName, fileName, "--node-address", nodeAddress},
			expectedOutput: []string{fmt.Sprintf("File successfully deleted: Name=%s", fileName)},
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestExternalWalletBalanceCommand(t *testing.T) {
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
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}

func TestPDPExternalArchivalMetadataCommand(t *testing.T) {
	nodeAddress := PickNodeRPCAddress(t)
	privateKey := ipctest.PickPrivateKey(t)
	_ = pdptest.PickPrivateKey(t)

	bucketName := testrand.String(10)
	_, _, err := captureCobraOutput(rootCmd, []string{"bucket", "create", "--private-key", privateKey, bucketName, "--node-address", nodeAddress})
	require.NoError(t, err)

	file, err := createTempFile(t, testrand.String(5), 2*memory.MB.ToInt64())
	require.NoError(t, err)
	fileName := filepath.Base(file)

	_, _, err = captureCobraOutput(rootCmd, []string{"file", "upload", "--private-key", privateKey, bucketName, file, "--node-address", nodeAddress})
	require.NoError(t, err)

	testCases := []testCase{
		{
			name:           "File not fully available",
			args:           []string{"file", "archival-metadata", bucketName, fileName, "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"Not fully available in archival storage"},
			expectError:    false,
		},
		{
			name:           "Verbose mode",
			args:           []string{"file", "archival-metadata", bucketName, fileName, "--private-key", privateKey, "-v", "--node-address", nodeAddress},
			expectedOutput: []string{"Total Chunks:", "Chunk CID:", "Block CID:"},
			expectError:    false,
		},
		{
			name:           "File not found",
			args:           []string{"file", "archival-metadata", bucketName, "nonexistent-file", "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"file not exists"},
			expectError:    true,
		},
		{
			name:           "Bucket name not provided",
			args:           []string{"file", "archival-metadata", "", fileName, "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"bucket name is required"},
			expectError:    true,
		},
		{
			name:           "File name not provided",
			args:           []string{"file", "archival-metadata", bucketName, "", "--private-key", privateKey, "--node-address", nodeAddress},
			expectedOutput: []string{"file name is required"},
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := captureCobraOutput(rootCmd, tc.args)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			for _, expected := range tc.expectedOutput {
				require.Contains(t, stdout+stderr, expected)
			}
		})
	}
}
