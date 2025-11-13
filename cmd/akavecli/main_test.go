// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestMain(m *testing.M) {
	initFlags()

	// Run tests
	code := m.Run()

	os.Exit(code)
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

func TestExitCode(t *testing.T) {
	pRoot := projectRoot(t)
	modulePath := filepath.Join(pRoot, "cmd", "akavecli")

	t.Run("incorrect user params/flags", func(t *testing.T) {
		cmd := exec.CommandContext(t.Context(), "go", "run", modulePath, "bucket", "create")
		err := cmd.Run()

		require.Error(t, err, "Expected non-zero exit code for invalid parameters")

		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			assert.NotEqual(t, 0, exitError.ExitCode(), "Expected non-zero exit code")
		} else {
			t.Fatalf("Expected exec.ExitError, got %T", err)
		}
	})

	t.Run("non-existent node address", func(t *testing.T) {
		bucketName := testrand.String(10)
		nonExistentAddress := "127.0.0.1:0" // Port 0 is reserved and never assigned

		cmd := exec.CommandContext(t.Context(), "go", "run", modulePath, "bucket", "create", bucketName, "--node-address", nonExistentAddress)
		err := cmd.Run()

		// Expect non-zero exit code due to connection failure
		require.Error(t, err, "Expected non-zero exit code when connecting to non-existent node")

		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			assert.NotEqual(t, 0, exitError.ExitCode(), "Expected non-zero exit code")
		} else {
			t.Fatalf("Expected exec.ExitError, got %T", err)
		}
	})
}

func projectRoot(tb testing.TB) string {
	tb.Helper()

	// Get the path to this file at runtime.
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		tb.Fatal("cannot get current file path")
	}

	dir := filepath.Dir(filename)

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			tb.Fatal("go.mod not found")
		}
		dir = parent
	}
}
