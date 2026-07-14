// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package encryption provides functions for encrypting and decrypting data using AES-GCM.
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"golang.org/x/exp/constraints"
)

const (
	keyLength = 32

	// Overhead is of 16 bytes for AES-GCM tag, 12 bytes for nonce.
	Overhead = 28
)

// ErrBufferTooSmall is returned when the provided buffer does not have sufficient capacity for encryption.
var ErrBufferTooSmall = errors.New("buffer too small for encryption")

// Encrypt encrypts the given data using the master key and the given info.
// data MUST have sufficient capacity for len(data) + EncryptionOverhead.
// Encryption is performed in-place, and the resliced data is returned.
func Encrypt(key, data []byte, info string) ([]byte, error) {
	gcm, err := GCMCipher(key, info)
	if err != nil {
		return nil, err
	}

	if cap(data) < len(data)+Overhead {
		return nil, ErrBufferTooSmall
	}

	// Shift originalData to make room for nonce at front
	originalData := data[gcm.NonceSize() : gcm.NonceSize()+len(data)]
	copy(originalData, data)

	// Generate nonce
	nonce := data[:gcm.NonceSize()]
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(originalData[:0], nonce, originalData, nil)

	// Return resliced data containing nonce + ciphertext
	return data[:gcm.NonceSize()+len(ciphertext)], nil
}

// EncryptD encrypts the given data using the master key and the given info. Deterministic encryption.
// WARNING: This function produces the same output for the same input. Only use it when deterministic
// encryption is specifically required. For general encryption needs, use Encrypt() instead.
func EncryptD(key, data []byte, info string) ([]byte, error) {
	gcm, err := GCMCipher(key, info)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, key)
	_, err = h.Write(data)
	if err != nil {
		return nil, err
	}

	nonce := h.Sum(nil)[:gcm.NonceSize()]

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Decrypt decrypts the given data in-place using the master key and the given info.
// Decryption is performed in-place, and the resliced encryptedData is returned.
func Decrypt(key, encryptedData []byte, info string) ([]byte, error) {
	gcm, err := GCMCipher(key, info)
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < gcm.NonceSize()+gcm.Overhead() {
		return nil, io.ErrUnexpectedEOF
	}

	nonce := encryptedData[:gcm.NonceSize()]
	ciphertext := encryptedData[gcm.NonceSize():]
	return gcm.Open(ciphertext[:0], nonce, ciphertext, nil)
}

// DeriveKey derives a key from the master key and the given info.
func DeriveKey(key []byte, info string) ([]byte, error) {
	return hkdf.Key(sha256.New, key, nil, info, keyLength)
}

// GCMCipher creates a new AES-GCM cipher using a key derived from the master key and info.
func GCMCipher(originKey []byte, info string) (cipher.AEAD, error) {
	key, err := DeriveKey(originKey, info)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

// CeilDiv returns the ceiling of a/b for positive integers.
// User must provide positive integers for a and b, otherwise the behavior is undefined.
func CeilDiv[T constraints.Integer](a, b T) T {
	return (a + b - 1) / b
}
