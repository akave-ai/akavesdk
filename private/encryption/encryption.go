// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package encryption provides functions for encrypting and decrypting data using AES-GCM.
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/hkdf"
)

const keyLength = 32

// Encrypt encrypts the given data using the master key and the given info.
func Encrypt(key, data, info []byte) ([]byte, error) {
	gcm, err := makeGCMCipher(key, info)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// EncryptD encrypts the given data using the master key and the given info. Deterministic encryption.
// WARNING: This function produces the same output for the same input. Only use it when deterministic
// encryption is specifically required. For general encryption needs, use Encrypt() instead.
func EncryptD(key, data, info []byte) ([]byte, error) {
	gcm, err := makeGCMCipher(key, info)
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

// Decrypt decrypts the given data using the master key and the given info.
func Decrypt(key, encryptedData, info []byte) ([]byte, error) {
	gcm, err := makeGCMCipher(key, info)
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < gcm.NonceSize() {
		return nil, io.ErrUnexpectedEOF
	}
	nonce, encryptedData := encryptedData[:gcm.NonceSize()], encryptedData[gcm.NonceSize():]
	return gcm.Open(nil, nonce, encryptedData, nil)
}

// DeriveKey derives a key from the master key and the given info.
func DeriveKey(key, info []byte) ([]byte, error) {
	r := hkdf.New(sha256.New, key, nil, info)
	derivedKey := make([]byte, keyLength)
	if _, err := io.ReadFull(r, derivedKey); err != nil {
		return nil, err
	}
	return derivedKey, nil
}

func makeGCMCipher(originKey, info []byte) (cipher.AEAD, error) {
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
