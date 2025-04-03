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
		info []byte
	}

	testData := []TestData{
		{
			name: "without info",
			key:  "foo",
			data: "big brown fox jumps over the lazy dog",
			info: nil,
		},
		{
			name: "with info",
			key:  "foo",
			data: "big brown fox jumps over the lazy dog",
			info: []byte("info"),
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			t.Logf("%s len(data) %d", td.name, len(td.data))

			encrypted, err := encryption.Encrypt([]byte(td.key), []byte(td.data), td.info)
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
	data := "quick brown fox jumps over the lazy dog"

	t.Run("non deterministic encryption", func(t *testing.T) {
		encryptedData1, err := encryption.Encrypt(key, []byte(data), nil)
		require.NoError(t, err)
		encryptedData2, err := encryption.Encrypt(key, []byte(data), nil)
		require.NoError(t, err)
		require.NotEqual(t, encryptedData1, encryptedData2)

		decryptedData1, err := encryption.Decrypt(key, encryptedData1, nil)
		require.NoError(t, err)
		decryptedData2, err := encryption.Decrypt(key, encryptedData2, nil)
		require.NoError(t, err)
		require.Equal(t, data, string(decryptedData1))
		require.Equal(t, data, string(decryptedData2))
	})

	t.Run("deterministic encryption", func(t *testing.T) {
		encryptedData1, err := encryption.EncryptD(key, []byte(data), nil)
		require.NoError(t, err)
		encryptedData2, err := encryption.EncryptD(key, []byte(data), nil)
		require.NoError(t, err)
		require.Equal(t, encryptedData1, encryptedData2)

		decryptedData, err := encryption.Decrypt(key, encryptedData1, nil)
		require.NoError(t, err)

		require.Equal(t, data, string(decryptedData))
	})
}

func TestDataOverhead(t *testing.T) {
	dataSizes := []int64{1, 16}
	key, _ := encryption.DeriveKey([]byte("key"), []byte("some_info"))
	for i, size := range dataSizes {
		data := testrand.Bytes(t, size*memory.MB.ToInt64())
		encrypted, err := encryption.Encrypt(key, data, []byte(fmt.Sprintf("%d", i)))
		require.NoError(t, err)
		require.NotEqual(t, data[:10], encrypted[:10])
		encryptedSize := len(encrypted)
		dataSize := len(data)
		t.Logf("Data size: %d, Encrypted size: %d, overhead: %d", dataSize, encryptedSize, encryptedSize-dataSize)
	}
}
