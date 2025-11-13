// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package pdptest provides utilities for PDP tests.
package pdptest

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/ipfs/go-cid"
)

const (
	// CalibrationWarmStorageContract is the warm storage contract address on the Filecoin calibration network.
	CalibrationWarmStorageContract = "0x02925630df557F957f70E112bA06e50965417CA0"
	// CalibrationFilecoinRPC is the RPC URL for the Filecoin calibration network.
	CalibrationFilecoinRPC = "https://api.calibration.node.glif.io/rpc/v1"
)

var (
	// PDPPrivateKey is the flag to set the PDP private key for tests.
	PDPPrivateKey = flag.String("pdp-private-key", os.Getenv("PDP_PRIVATE_KEY"), "flag to set deployers hex private key for client tests for pdp")
	// PDPServerURL is the flag to set the PDP server URL for tests.
	PDPServerURL = flag.String("pdp-server-url", os.Getenv("PDP_SERVER_URL"), "PDP server URL for client tests")
)

// PickPrivateKey returns the PDP private key from flag or environment variable.
// Skips the test if not provided.
func PickPrivateKey(t testing.TB) string {
	if *PDPPrivateKey == "" || strings.EqualFold(*PDPPrivateKey, "omit") {
		t.Skip("private key flag missing, example: -PDP_PRIVATE_KEY=<deployers hex private key>")
	}
	return *PDPPrivateKey
}

// PickServerURL returns the PDP server URL from flag or environment variable.
// Skips the test if not provided.
func PickServerURL(t testing.TB) string {
	if *PDPServerURL == "" {
		t.Skip("PDP server URL flag missing, example: -pdp-server-url=<pdp server url>")
	}
	return *PDPServerURL
}

// CalculatePieceCID calculates the PDP piece CID from data.
func CalculatePieceCID(data []byte) (cid.Cid, error) {
	cp := new(commp.Calc)
	if _, err := cp.Write(data); err != nil {
		return cid.Undef, fmt.Errorf("failed to calculate piece hash: %w", err)
	}

	rawCommP, _, err := cp.Digest()
	if err != nil {
		return cid.Undef, fmt.Errorf("failed to get piece commitment: %w", err)
	}

	pieceCid, err := commcid.DataCommitmentToPieceCidv2(rawCommP, uint64(len(data)))
	if err != nil {
		return cid.Undef, fmt.Errorf("failed to convert commitment to CID: %w", err)
	}

	return pieceCid, nil
}
