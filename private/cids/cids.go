// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package cids provides utilities for working with Content Identifiers (CIDs).
package cids

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// VerifyRaw checks if the provided CID matches the data.
func VerifyRaw(providedCID string, data []byte) error {
	parsedCID, err := cid.Decode(providedCID)
	if err != nil {
		return fmt.Errorf("failed to decode provided CID: %w", err)
	}

	return Verify(parsedCID, data)
}

// Verify checks if the provided CID matches the data.
func Verify(c cid.Cid, data []byte) error {
	calculatedCID, err := calculateStandardCID(c, data)
	if err != nil {
		return err
	}

	if !calculatedCID.Equals(c) {
		return fmt.Errorf("CID mismatch: provided %s, calculated %s", c.String(), calculatedCID.String())
	}

	return nil
}

func calculateStandardCID(c cid.Cid, data []byte) (cid.Cid, error) {
	prefix := c.Prefix()
	hash, err := multihash.Sum(data, prefix.MhType, -1)
	if err != nil {
		return cid.Undef, fmt.Errorf("failed to create multihash: %w", err)
	}

	switch prefix.Version {
	case 0:
		return cid.NewCidV0(hash), nil
	case 1:
		return cid.NewCidV1(c.Type(), hash), nil
	default:
		return cid.Undef, fmt.Errorf("unsupported CID version: %d", prefix.Version)
	}
}
