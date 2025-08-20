// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

const (
	blockPartSize  = 128 * memory.KiB
	maxConcurrency = 5
	secretKey      = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

var (
	// nodeRPCAddress is flag to set RPC address of akave node.
	nodeRPCAddress = flag.String("node-rpc-address", os.Getenv("NODE_RPC_ADDRESS"), "flag to set node rpc address")
	// PrivateKey is flag set to private key.
	PrivateKey = flag.String("sdk-private-key", os.Getenv("PRIVATE_KEY"), "flag to set private key")
	// DialURI is flag set to ipc dial uri.
	DialURI = flag.String("sdk-dial-uri", os.Getenv("DIAL_URI"), "flag to set dial URI")
)

// PickNodeRPCAddress picks node PRC address from flag.
func PickNodeRPCAddress(t testing.TB) string {
	if *nodeRPCAddress == "" || strings.EqualFold(*nodeRPCAddress, "omit") {
		t.Skip("node rpc address flag missing, example: -NODE_RPC_ADDRESS=<node rpc address>")
	}
	return *nodeRPCAddress
}

// PickPrivateKey picks hex private key of deployer.
func PickPrivateKey(t testing.TB) string {
	if *PrivateKey == "" || strings.EqualFold(*PrivateKey, "omit") {
		t.Skip("private key flag missing, example: -PRIVATE_KEY=<deployers hex private key>")
	}
	return *PrivateKey
}

// PickDialURI picks ipc dial uri.
func PickDialURI(t testing.TB) string {
	if *DialURI == "" || strings.EqualFold(*DialURI, "omit") {
		t.Skip("private key flag missing, example: -DIAL_URI=<deployers hex private key>")
	}
	return *DialURI
}

func TestCreateSDKClient(t *testing.T) {
	t.Run("invalid max blocks in chunk", func(t *testing.T) {
		_, err := sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithStreamingMaxBlocksInChunk(1))
		require.Error(t, err)
		require.Equal(t, "sdk: streaming max blocks in chunk 1 should be >= 2", err.Error())
	})

	t.Run("invalid erasure coding config", func(t *testing.T) {
		_, err := sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithErasureCoding(17))
		require.Error(t, err)
		require.Equal(t, "sdk: parity blocks count 17 should be <= 16", err.Error())

		_, err = sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true,
			sdk.WithErasureCoding(40),
			sdk.WithStreamingMaxBlocksInChunk(64),
		)
		require.Error(t, err)
		require.Equal(t, "sdk: parity blocks count 40 should be <= 32", err.Error())
	})

	t.Run("invalid encryption key size", func(t *testing.T) {
		_, err := sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte("short")))
		require.Error(t, err)
		require.Equal(t, "sdk: encyption key length should be 32 bytes long", err.Error())
	})
}

func TestRandomFile(t *testing.T) {
	f1 := generate10MiBFile(t, 2025)
	f2 := generate10MiBFile(t, 2025)

	t.Log(f1.Bytes()[:20])
	t.Log(f2.Bytes()[:20])
	require.EqualValues(t, f1.Bytes(), f2.Bytes())
}

func generateAny2MBFile(t *testing.T) *bytes.Buffer {
	data := testrand.Bytes(t, 2*memory.MB.ToInt64())
	return bytes.NewBuffer(data)
}

func generate10MiBFile(t *testing.T, seed int64) *bytes.Buffer {
	data := testrand.BytesD(t, seed, 10*memory.MiB.ToInt64())
	return bytes.NewBuffer(data)
}

func randomBucketName(t require.TestingT, size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	require.NoError(t, err)
	return hex.EncodeToString(b)
}

// checks lengths, first and last N bytes.
func checkFileContents(t *testing.T, n int, expected, actual []byte) {
	t.Helper()
	require.Equal(t, len(expected), len(actual))
	// check first 10 bytes
	require.EqualValues(t, expected[:n], actual[:n])
	// check last 10 bytes
	require.EqualValues(t, expected[len(expected)-n:], actual[len(actual)-n:])
}
