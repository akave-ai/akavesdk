// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/akave-ai/akavesdk/private/testrand"
)

const encKey = "1234567890123456789012345678901212345678901234567890123456789012"

var (
	// nodeAddress is flag to set RPC address of akave node.
	nodeAddress = flag.String("node-rpc-address", os.Getenv("NODE_RPC_ADDRESS"), "flag to set node rpc address")
)

type testCase struct {
	name           string
	args           []string
	expectedOutput []string
	expectError    bool
}

// PickNodeRPCAddress picks node PRC address from flag.
func PickNodeRPCAddress(t testing.TB) string {
	if *nodeAddress == "" || strings.EqualFold(*nodeAddress, "omit") {
		t.Skip("node rpc address flag missing, example: -NODE_RPC_ADDRESS=<node rpc address>")
	}
	return *nodeAddress
}

func createTempFile(t *testing.T, prefix string, size int64) (string, error) {
	t.Helper()
	tempFile, err := os.CreateTemp(t.TempDir(), fmt.Sprintf("%s-*", prefix))
	if err != nil {
		return "", err
	}
	txt := testrand.String(int(size))
	_, err = tempFile.WriteString(txt)
	if err != nil {
		return "", err
	}
	err = tempFile.Close()
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

func extractRootCID(output string) (string, error) {
	re := regexp.MustCompile(`RootCID=([a-zA-Z0-9]+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return "", fmt.Errorf("failed to extract root CID")
	}
	return matches[1], nil
}

func captureCobraOutput(cmd *cobra.Command, args []string) (string, string, error) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd.SetOut(&stdoutBuf)
	cmd.SetErr(&stderrBuf)

	cmd.SetArgs(args)

	err := cmd.Execute()

	stdout := stdoutBuf.String()
	stderr := stderrBuf.String()

	// Reset flag values to defaults after each command execution to prevent state leakage between tests
	encryptionKey = ""
	useMetadataEncryption = false
	privateKey = ""
	bucketListLimit = 20
	fileListLimit = 20
	bucketListOffset = 0
	fileListOffset = 0

	return stdout, stderr, err
}
