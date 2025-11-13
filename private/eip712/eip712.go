// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package eip712 provides access to eip712 signer.
package eip712

import (
	"crypto/ecdsa"
	"fmt"
	"maps"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var eip712DomainType = []apitypes.Type{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
}

// Sign signs storage data with private key.
func Sign(privateKey *ecdsa.PrivateKey, domain apitypes.TypedDataDomain, primaryType string, dataTypes apitypes.Types, dataMessage map[string]any) ([]byte, error) {
	dataTypesCopy := maps.Clone(dataTypes)
	dataTypesCopy["EIP712Domain"] = eip712DomainType

	typedData := apitypes.TypedData{
		Types:       dataTypesCopy,
		PrimaryType: primaryType,
		Domain:      domain,
		Message:     dataMessage,
	}

	digestHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return nil, err
	}

	signature, err := crypto.Sign(digestHash, privateKey)
	if err != nil {
		return nil, fmt.Errorf("error signing hash: %w", err)
	}

	signature[64] += 27

	return signature, nil
}

// RecoverSignerAddress recovers the signer address from an eip712. For testing purpose.
func RecoverSignerAddress(signature []byte, domain apitypes.TypedDataDomain, primaryType string, dataTypes apitypes.Types, dataMessage map[string]any) (common.Address, error) {
	dataTypesCopy := maps.Clone(dataTypes)
	dataTypesCopy["EIP712Domain"] = eip712DomainType

	typedData := apitypes.TypedData{
		Types:       dataTypesCopy,
		PrimaryType: primaryType,
		Domain:      domain,
		Message:     dataMessage,
	}

	digestHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return common.Address{}, err
	}

	sigCopy := make([]byte, len(signature))
	copy(sigCopy, signature)

	sigCopy[64] -= 27

	pubKey, err := crypto.Ecrecover(digestHash, sigCopy)
	if err != nil {
		return common.Address{}, fmt.Errorf("error recovering public key: %w", err)
	}

	pubKeyECDSA, err := crypto.UnmarshalPubkey(pubKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("error unmarshaling public key: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKeyECDSA), nil
}
