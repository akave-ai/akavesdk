// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package eip712 provides access to eip712 signer.
package eip712

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// TypedData contains data type and name.
type TypedData struct {
	Name string
	Type string
}

// Domain represents the domain separator.
type Domain struct {
	Name              string
	Version           string
	ChainID           *big.Int
	VerifyingContract common.Address
}

// Sign signs storage data with private key.
func Sign(privateKey *ecdsa.PrivateKey, domain Domain, dataMessage map[string]interface{}, dataTypes map[string][]TypedData) ([]byte, error) {
	hash, err := hashTypedData(domain, dataMessage, dataTypes)
	if err != nil {
		return nil, fmt.Errorf("error hashing data: %w", err)
	}

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("error signing hash: %w", err)
	}

	signature[64] += 27

	return signature, nil
}

// RecoverSignerAddress recovers the signer address from an eip712. For testing purpose.
func RecoverSignerAddress(signature []byte, domain Domain, dataMessage map[string]interface{}, dataTypes map[string][]TypedData) (common.Address, error) {
	hash, err := hashTypedData(domain, dataMessage, dataTypes)
	if err != nil {
		return common.Address{}, fmt.Errorf("error hashing typed data: %w", err)
	}

	sigCopy := make([]byte, len(signature))
	copy(sigCopy, signature)

	sigCopy[64] -= 27

	pubKey, err := crypto.Ecrecover(hash.Bytes(), sigCopy)
	if err != nil {
		return common.Address{}, fmt.Errorf("error recovering public key: %w", err)
	}

	pubKeyECDSA, err := crypto.UnmarshalPubkey(pubKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("error unmarshaling public key: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKeyECDSA), nil
}

func encodeType(primaryType string, types map[string][]TypedData) string {
	var buffer bytes.Buffer
	buffer.WriteString(primaryType)
	buffer.WriteString("(")

	first := true
	for _, field := range types[primaryType] {
		if !first {
			buffer.WriteString(",")
		}
		buffer.WriteString(field.Type)
		buffer.WriteString(" ")
		buffer.WriteString(field.Name)
		first = false
	}

	buffer.WriteString(")")
	return buffer.String()
}

func typeHash(primaryType string, types map[string][]TypedData) common.Hash {
	return crypto.Keccak256Hash([]byte(encodeType(primaryType, types)))
}

// hashTypedData creates the hash that's used for signing.
func hashTypedData(domain Domain, dataMessage map[string]interface{}, dataTypes map[string][]TypedData) (common.Hash, error) {
	domainTypes := map[string][]TypedData{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}

	domainMessage := map[string]interface{}{
		"name":              domain.Name,
		"version":           domain.Version,
		"chainId":           domain.ChainID,
		"verifyingContract": domain.VerifyingContract,
	}

	domainHash, err := encodeData("EIP712Domain", domainMessage, domainTypes)
	if err != nil {
		return common.Hash{}, fmt.Errorf("error hashing domain: %w", err)
	}

	dataHash, err := encodeData("StorageData", dataMessage, dataTypes)
	if err != nil {
		return common.Hash{}, fmt.Errorf("error encoding data: %w", err)
	}

	rawData := []byte{0x19, 0x01}
	rawData = append(rawData, domainHash...)
	rawData = append(rawData, dataHash...)

	return crypto.Keccak256Hash(rawData), nil
}

func encodeData(primaryType string, data map[string]interface{}, types map[string][]TypedData) ([]byte, error) {
	typeHash := typeHash(primaryType, types)

	encodedData := [][]byte{typeHash.Bytes()}

	for _, field := range types[primaryType] {
		value := data[field.Name]
		encodedValue, err := encodeValue(value, field.Type)
		if err != nil {
			return nil, err
		}
		encodedData = append(encodedData, encodedValue)
	}

	return crypto.Keccak256(bytes.Join(encodedData, nil)), nil
}

func encodeValue(value interface{}, typeName string) ([]byte, error) {
	switch typeName {
	case "string":
		strVal, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string, got %T", value)
		}
		return crypto.Keccak256([]byte(strVal)), nil

	case "bytes":
		bytesVal, ok := value.([]byte)
		if !ok {
			return nil, fmt.Errorf("expected []byte, got %T", value)
		}
		return crypto.Keccak256(bytesVal), nil

	case "bytes32":
		bytes32Val, ok := value.([32]byte)
		if !ok {
			return nil, fmt.Errorf("expected [32]byte, got %T", value)
		}
		return bytes32Val[:], nil

	case "uint8":
		uint8Val, ok := value.(uint8)
		if !ok {
			return nil, fmt.Errorf("expected uint8, got %T", value)
		}
		buf := make([]byte, 32)
		buf[31] = uint8Val
		return buf, nil

	case "uint64":
		uint64Val, ok := value.(uint64)
		if !ok {
			return nil, fmt.Errorf("expected uint64, got %T", value)
		}
		buf := make([]byte, 32)
		binary.BigEndian.PutUint64(buf[24:], uint64Val)
		return buf, nil

	case "uint256":
		var bigIntVal *big.Int
		switch v := value.(type) {
		case *big.Int:
			bigIntVal = v
		case big.Int:
			bigIntVal = new(big.Int).Set(&v)
		default:
			return nil, fmt.Errorf("expected *big.Int, got %T", value)
		}

		buf := make([]byte, 32)
		bigIntVal.FillBytes(buf)
		return buf, nil

	case "address":
		var addrVal common.Address
		switch v := value.(type) {
		case common.Address:
			addrVal = v
		case string:
			addrVal = common.HexToAddress(v)
		default:
			return nil, fmt.Errorf("expected common.Address, got %T", value)
		}

		buf := make([]byte, 32)
		copy(buf[12:], addrVal.Bytes())
		return buf, nil

	default:
		return nil, fmt.Errorf("unsupported type: %s", typeName)
	}
}
