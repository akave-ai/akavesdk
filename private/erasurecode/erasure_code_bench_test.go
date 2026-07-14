// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package erasurecode_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func BenchmarkEncode(b *testing.B) {
	ec, err := erasurecode.New(16, 16)
	require.NoError(b, err)

	data := make([]byte, 16*1024*1024) // 16 MiB
	for i := range data {
		data[i] = byte(i)
	}

	encryptionKey := testrand.Bytes(b, 32)
	testData := make([]byte, len(data), len(data)+encryption.Overhead)

	b.ReportAllocs()

	var encoded []byte
	for b.Loop() {
		b.StopTimer()
		copy(testData, data)
		b.StartTimer()
		encryptedData, err := encryption.Encrypt(encryptionKey, testData, "benchmark")
		if err != nil {
			b.Fatal(err)
		}
		encoded, err = ec.Encode(encryptedData)
		if err != nil {
			b.Fatal(err)
		}
		if len(encoded) <= len(data) {
			b.Fatalf("expected encoded data to be larger than original: encoded=%d original=%d", len(encoded), len(data))
		}
	}
}

func BenchmarkEncodeInPlace(b *testing.B) {
	ec, err := erasurecode.New(16, 16)
	require.NoError(b, err)

	data := make([]byte, 16*1024*1024) // 16 MiB
	for i := range data {
		data[i] = byte(i)
	}

	encryptionKey := testrand.Bytes(b, 32)
	encryptedDataLen := len(data) + encryption.Overhead
	requiredCapacity := encryptedDataLen + erasurecode.WrapOverhead
	// Allocate with extra capacity (enables in-place wrapping)
	testData := make([]byte, len(data), requiredCapacity)

	b.ReportAllocs()

	var encoded []byte
	for b.Loop() {
		b.StopTimer()
		copy(testData, data)
		b.StartTimer()
		encryptedData, err := encryption.Encrypt(encryptionKey, testData, "benchmark")
		if err != nil {
			b.Fatal(err)
		}
		encoded, err = ec.Encode(encryptedData)
		if err != nil {
			b.Fatal(err)
		}
		if len(encoded) <= len(data) {
			b.Fatalf("expected encoded data to be larger than original: encoded=%d original=%d", len(encoded), len(data))
		}
	}
}
