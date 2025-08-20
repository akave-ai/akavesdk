// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestMain(m *testing.M) {
	initFlags()

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func TestCmdCarCreate(t *testing.T) {
	t.Skip("Fix me, this test does not work")

	nodeAddress := PickNodeRPCAddress(t)
	fileName := "test.car"
	destination := t.TempDir()
	path := filepath.Join(destination, fileName)
	_, _, err := captureCobraOutput(rootCmd, []string{"car", "create", path, "--node-address", nodeAddress})
	assert.NoError(t, err)

	testCases := []testCase{
		{
			name:           "CAR file created successfully",
			args:           []string{"car", "create", path, "--node-address", nodeAddress},
			expectedOutput: []string{"CAR file: Name=test.car"},
			expectError:    false,
		},
		{
			name:           "Destination path is required",
			args:           []string{"car", "create", "", "--node-address", nodeAddress},
			expectedOutput: []string{"destination file path is required"},
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

func TestWalletFlow(t *testing.T) {
	// Create temporary directory for test wallets
	tmpKeystoreDir := t.TempDir()

	walletName := testrand.String(10)
	// Create a new wallet
	stdout, stderr, err := captureCobraOutput(rootCmd, []string{"wallet", "create", walletName, "--keystore", tmpKeystoreDir})
	assert.NoError(t, err)
	assert.Contains(t, stdout+stderr, fmt.Sprintf("Wallet (%s) created successfully", walletName))

	// List wallets and verify the created wallet exists
	stdout, stderr, err = captureCobraOutput(rootCmd, []string{"wallet", "list", "--keystore", tmpKeystoreDir})
	assert.NoError(t, err)
	assert.Contains(t, stdout+stderr, walletName)

	// Export private key to verify wallet creation
	stdout, stderr, err = captureCobraOutput(rootCmd, []string{"wallet", "export-key", walletName, "--keystore", tmpKeystoreDir})
	assert.NoError(t, err)
	assert.Contains(t, stdout+stderr, "Private key:")
}
