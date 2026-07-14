// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package encryption_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestEncryption(t *testing.T) {
	type TestData struct {
		name string
		key  string
		data string
		info string
	}

	testData := []TestData{
		{
			name: "without info",
			key:  "foo",
			data: "big brown fox jumps over the lazy dog",
			info: "",
		},
		{
			name: "with info",
			key:  "foo",
			data: "big brown fox jumps over the lazy dog",
			info: "info",
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			t.Logf("%s len(data) %d", td.name, len(td.data))
			data := make([]byte, len(td.data), len(td.data)+encryption.Overhead)
			copy(data, []byte(td.data))

			encrypted, err := encryption.Encrypt([]byte(td.key), data, td.info)
			require.NoError(t, err)

			t.Logf("%s encrypted data: %s", td.name, base64.StdEncoding.EncodeToString(encrypted))
			t.Logf("%s encrypted len(data): %d", td.name, len(encrypted))

			decrypted, err := encryption.Decrypt([]byte(td.key), encrypted, td.info)
			require.NoError(t, err)

			t.Logf("%s descrypted data: %s", td.name, string(decrypted))

			require.Equal(t, td.data, string(decrypted))
		})
	}
}

func TestEncryptionDeterminismAndNonDeterminism(t *testing.T) {
	key := []byte("key")
	dataStr := "quick brown fox jumps over the lazy dog"

	t.Run("non deterministic encryption", func(t *testing.T) {
		data1 := make([]byte, len(dataStr), len(dataStr)+encryption.Overhead)
		copy(data1, []byte(dataStr))

		data2 := make([]byte, len(dataStr), len(dataStr)+encryption.Overhead)
		copy(data2, []byte(dataStr))

		encryptedData1, err := encryption.Encrypt(key, data1, "")
		require.NoError(t, err)
		encryptedData2, err := encryption.Encrypt(key, data2, "")
		require.NoError(t, err)
		require.NotEqual(t, encryptedData1, encryptedData2)

		decryptedData1, err := encryption.Decrypt(key, encryptedData1, "")
		require.NoError(t, err)
		decryptedData2, err := encryption.Decrypt(key, encryptedData2, "")
		require.NoError(t, err)
		require.Equal(t, dataStr, string(decryptedData1))
		require.Equal(t, dataStr, string(decryptedData2))
	})

	t.Run("deterministic encryption", func(t *testing.T) {
		encryptedData1, err := encryption.EncryptD(key, []byte(dataStr), "")
		require.NoError(t, err)
		encryptedData2, err := encryption.EncryptD(key, []byte(dataStr), "")
		require.NoError(t, err)
		require.Equal(t, encryptedData1, encryptedData2)

		decryptedData, err := encryption.Decrypt(key, encryptedData1, "")
		require.NoError(t, err)

		require.Equal(t, dataStr, string(decryptedData))
	})
}

func TestDataOverhead(t *testing.T) {
	dataSizes := []int64{1, 16}
	key, _ := encryption.DeriveKey([]byte("key"), "some_info")
	for i, size := range dataSizes {
		data := make([]byte, size*memory.MB.ToInt64(), size*memory.MB.ToInt64()+encryption.Overhead)
		copy(data, testrand.Bytes(t, size*memory.MB.ToInt64()))

		originalData := make([]byte, size*memory.MB.ToInt64())
		copy(originalData, data)

		encrypted, err := encryption.Encrypt(key, data, fmt.Sprintf("%d", i))
		require.NoError(t, err)
		require.NotEqual(t, originalData[:10], encrypted[:10])
		encryptedSize := len(encrypted)
		dataSize := len(originalData)
		t.Logf("Data size: %d, Encrypted size: %d, overhead: %d", dataSize, encryptedSize, encryptedSize-dataSize)
	}
}

func TestCeilDiv(t *testing.T) {
	tests := []struct {
		name     string
		run      func() any
		expected any
	}{
		{
			name:     "exact division with int",
			run:      func() any { return encryption.CeilDiv(8, 4) },
			expected: 2,
		},
		{
			name:     "rounds up with int",
			run:      func() any { return encryption.CeilDiv(9, 4) },
			expected: 3,
		},
		{
			name:     "rounds up with int64",
			run:      func() any { return encryption.CeilDiv(int64(10), int64(3)) },
			expected: int64(4),
		},
		{
			name:     "single chunk with uint",
			run:      func() any { return encryption.CeilDiv(uint(1), uint(8)) },
			expected: uint(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.run())
		})
	}
}

func BenchmarkDecrypt(b *testing.B) {
	key := []byte("key")
	info := "benchmark"
	dataSize := 16 * 1024 * 1024
	data := make([]byte, dataSize, dataSize+encryption.Overhead)
	copy(data, testrand.Bytes(b, int64(dataSize)))
	encrypted, err := encryption.Encrypt(key, data, info)
	require.NoError(b, err)

	original := make([]byte, len(encrypted))
	copy(original, encrypted)
	buf := make([]byte, len(encrypted))

	b.ReportAllocs()

	var result []byte
	for b.Loop() {
		b.StopTimer()
		copy(buf, original)
		b.StartTimer()
		result, err = encryption.Decrypt(key, buf, info)
	}
	_ = result
	require.NoError(b, err)
}

func BenchmarkEncrypt(b *testing.B) {
	key := []byte("key")
	info := "benchmark"
	dataSize := 1024 * 1024
	data := make([]byte, dataSize, dataSize+encryption.Overhead)
	copy(data, testrand.Bytes(b, int64(dataSize)))

	var result []byte
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		currentInfo := fmt.Sprintf("%s-%d", info, i)
		b.StartTimer()
		result, err = encryption.Encrypt(key, data, currentInfo)
	}
	_ = result
	require.NoError(b, err)
}
