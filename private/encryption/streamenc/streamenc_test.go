// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package streamenc_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/encryption/streamenc"
	"github.com/akave-ai/akavesdk/private/testrand"
)

// encryptHelper allocates a buffer with sufficient capacity for the
// ciphertext, copies plaintext into it, and encrypts in-place.
func encryptHelper(t *testing.T, key []byte, plaintext []byte, info string) []byte {
	t.Helper()

	plaintextSize := len(plaintext)
	totalSize := plaintextSize + streamenc.Overhead(plaintextSize)

	buf := make([]byte, plaintextSize, totalSize)
	copy(buf, plaintext)

	n, err := streamenc.Encrypt(key, buf, info)
	require.NoError(t, err)
	return buf[:n]
}

func TestEncryptDecryptRoundtrip(t *testing.T) {
	key := []byte("test-master-key")
	info := "test-info"

	sizes := []struct {
		name          string
		plaintextSize int
	}{
		{"minimum (1 byte plaintext)", 1},
		{"small (< 1 block)", 1000},
		{"exactly block0", streamenc.Block0DataSize},
		{"block0 + 1 byte", streamenc.Block0DataSize + 1},
		{"block0 + full block", streamenc.Block0DataSize + streamenc.BlockNDataSize},
		{"block0 + full block + partial", streamenc.Block0DataSize + streamenc.BlockNDataSize + 512},
		{"five blocks", streamenc.Block0DataSize + 4*streamenc.BlockNDataSize},
		{"1 MB", 1 << 20},
		{"8 MB", 8 << 20},
		{"16 MB", 16 << 20},
	}

	for _, tc := range sizes {
		t.Run(tc.name, func(t *testing.T) {
			original := testrand.Bytes(t, int64(tc.plaintextSize))

			ciphertext := encryptHelper(t, key, original, info)
			require.NotEqual(t, ciphertext, original)

			require.Equal(t, 0, len(ciphertext)%16, "ciphertext length must be divisible by 16")

			n, err := streamenc.DecryptAllBlocks(key, ciphertext, info, streamenc.Version)
			require.NoError(t, err)
			require.Equal(t, original, ciphertext[:n])

			overhead := len(ciphertext) - n
			percent := float64(overhead) / float64(len(ciphertext)) * 100

			t.Logf("Original data size %d, overhead %d (%.2f%%)", n, overhead, percent)
		})
	}
}

func TestEncryptNonDeterministic(t *testing.T) {
	key := []byte("test-key")
	info := "info"
	plaintextSize := 20 << 10 // 20 KB
	original := testrand.Bytes(t, int64(plaintextSize))

	ct1 := encryptHelper(t, key, original, info)
	ct2 := encryptHelper(t, key, original, info)

	require.NotEqual(t, ct1, ct2, "two encryptions of same plaintext must differ")

	r1, err := streamenc.DecryptAllBlocks(key, ct1, info, streamenc.Version)
	require.NoError(t, err)
	r2, err := streamenc.DecryptAllBlocks(key, ct2, info, streamenc.Version)
	require.NoError(t, err)
	require.Equal(t, original, ct1[:r1])
	require.Equal(t, original, ct2[:r2])
}

func TestParsedHeaderValues(t *testing.T) {
	key := []byte("test-key")
	info := "hdr-test"
	plaintextSize := 8000
	original := testrand.Bytes(t, int64(plaintextSize))

	ciphertext := encryptHelper(t, key, original, info)

	ver, _, plaintextSizeFromHeader, err := streamenc.ParseHeader(ciphertext)
	require.NoError(t, err)
	require.Equal(t, uint8(streamenc.Version), ver)
	require.Equal(t, uint32(plaintextSize), plaintextSizeFromHeader)
}

func TestParseHeaderTooShort(t *testing.T) {
	requireParseHeaderErr := func(data []byte) {
		_, _, _, err := streamenc.ParseHeader(data)
		require.Error(t, err)
	}

	requireParseHeaderErr(make([]byte, streamenc.HeaderSize-1))
	requireParseHeaderErr([]byte{})
}

func TestNumBlocks(t *testing.T) {
	cases := []struct {
		plaintextSize int
		expected      int
	}{
		{0, 0},
		{1, 1},
		{streamenc.Block0DataSize, 1},
		{streamenc.Block0DataSize + 1, 2},
		{streamenc.Block0DataSize + streamenc.BlockNDataSize, 2},
		{streamenc.Block0DataSize + streamenc.BlockNDataSize + 1, 3},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected, streamenc.NumBlocks(tc.plaintextSize), "plaintextSize=%d", tc.plaintextSize)
	}
}

func TestOverheadFromPlaintext(t *testing.T) {
	cases := []struct {
		plaintextSize int
	}{
		{0},
		{1},
		{streamenc.Block0DataSize},
		{streamenc.Block0DataSize + 1},
		{streamenc.Block0DataSize + streamenc.BlockNDataSize},
		{streamenc.Block0DataSize + streamenc.BlockNDataSize + 1},
		{1 << 20},
		{16 << 20},
	}
	for _, tc := range cases {
		overhead := streamenc.Overhead(tc.plaintextSize)
		totalSize := tc.plaintextSize + overhead
		assert.Equal(t, 0, totalSize%16, "total ciphertext size must be divisible by 16, plaintextSize=%d", tc.plaintextSize)
		t.Logf("Plain text size %d, overhead %d, total %d", tc.plaintextSize, overhead, totalSize)
	}
}

func TestBlockDataSize(t *testing.T) {
	t.Run("single partial block", func(t *testing.T) {
		size := streamenc.BlockDataSize(100, 0)
		require.Equal(t, 100, size)
	})

	t.Run("block0 of two-block file is full", func(t *testing.T) {
		total := streamenc.Block0DataSize + 500
		size := streamenc.BlockDataSize(total, 0)
		require.Equal(t, streamenc.Block0DataSize, size)
	})

	t.Run("last block of two-block file is remainder", func(t *testing.T) {
		total := streamenc.Block0DataSize + 500
		size := streamenc.BlockDataSize(total, 1)
		require.Equal(t, 500, size)
	})

	t.Run("exactly two full blocks", func(t *testing.T) {
		total := streamenc.Block0DataSize + streamenc.BlockNDataSize
		assert.Equal(t, streamenc.Block0DataSize, streamenc.BlockDataSize(total, 0))
		assert.Equal(t, streamenc.BlockNDataSize, streamenc.BlockDataSize(total, 1))
	})
}

func TestEncryptedBlockOutOfBounds(t *testing.T) {
	require.Panics(t, func() { _ = streamenc.EncryptedBlock([]byte{}, -1) })

	key := []byte("test-key")
	plaintextSize := 100
	original := testrand.Bytes(t, int64(plaintextSize))

	ciphertext := encryptHelper(t, key, original, "info")

	// only block 0 is valid; block 1 is out of bounds and must panic.
	require.NotNil(t, streamenc.EncryptedBlock(ciphertext, 0))
	require.Panics(t, func() { _ = streamenc.EncryptedBlock(ciphertext, 1) })
}

func TestDecryptBlockWrongKey(t *testing.T) {
	key := []byte("correct-key")
	wrongKey := []byte("wrong-key-value!")
	info := "info"
	plaintextSize := 3 << 10 // 3 KB
	original := testrand.Bytes(t, int64(plaintextSize))

	ciphertext := encryptHelper(t, key, original, info)

	_, initialNonce, _, err := streamenc.ParseHeader(ciphertext)
	require.NoError(t, err)

	block0 := streamenc.EncryptedBlock(ciphertext, 0)

	gcmWrong, err := encryption.GCMCipher(wrongKey, info)
	require.NoError(t, err)

	_, err = streamenc.DecryptBlock(gcmWrong, block0, initialNonce, streamenc.Version, 0, plaintextSize)
	require.Error(t, err)
}

func TestDecryptBlockTamperedCiphertext(t *testing.T) {
	key := []byte("test-key")
	info := "info"
	plaintextSize := 16 << 10 // 16 KB
	original := testrand.Bytes(t, int64(plaintextSize))

	ciphertext := encryptHelper(t, key, original, info)

	// Flip a byte in block 0's ciphertext (after the header).
	ciphertext[streamenc.HeaderSize+2] ^= 0xFF

	_, initialNonce, _, err := streamenc.ParseHeader(ciphertext)
	require.NoError(t, err)

	block0 := make([]byte, len(ciphertext))
	copy(block0, ciphertext)

	gcm, err := encryption.GCMCipher(key, info)
	require.NoError(t, err)

	_, err = streamenc.DecryptBlock(gcm, block0, initialNonce, streamenc.Version, 0, plaintextSize)
	require.Error(t, err)
}

func TestDecryptBlockWrongBlockIndex(t *testing.T) {
	key := []byte("test-key")
	info := "info"
	desiredPlaintext := streamenc.Block0DataSize + 100
	original := testrand.Bytes(t, int64(desiredPlaintext))

	ciphertext := encryptHelper(t, key, original, info)

	_, initialNonce, _, err := streamenc.ParseHeader(ciphertext)
	require.NoError(t, err)

	// Take block 0 but decrypt with index 1 (wrong nonce -> auth failure).
	block0 := make([]byte, streamenc.MaxBlockSize)
	copy(block0, ciphertext[:streamenc.MaxBlockSize])

	gcm, err := encryption.GCMCipher(key, info)
	require.NoError(t, err)

	_, err = streamenc.DecryptBlock(gcm, block0, initialNonce, streamenc.Version, 1, desiredPlaintext)
	require.Error(t, err)
}

func TestEncryptInsufficientCapacity(t *testing.T) {
	key := []byte("test-key")
	data := make([]byte, 100) // no extra capacity
	_, err := streamenc.Encrypt(key, data, "info")
	require.ErrorIs(t, err, encryption.ErrBufferTooSmall)
}

func TestEncryptEmptyPlaintext(t *testing.T) {
	key := []byte("test-key")
	_, err := streamenc.Encrypt(key, []byte{}, "info")
	require.EqualError(t, err, "plaintext cannot be empty")
}

func TestPlaintextSizeForTarget(t *testing.T) {
	cases := []struct {
		name       string
		targetSize int
		expected   int
		wantErr    bool
	}{
		{"zero", 0, 0, true},
		{"one block (MaxBlockSize)", streamenc.MaxBlockSize, streamenc.Block0DataSize, false},
		{"two blocks (2*MaxBlockSize)", 2 * streamenc.MaxBlockSize, streamenc.Block0DataSize + streamenc.BlockNDataSize, false},
		{"three blocks (3*MaxBlockSize)", 3 * streamenc.MaxBlockSize, streamenc.Block0DataSize + 2*streamenc.BlockNDataSize, false},
		{"sub-block aligned (48 bytes)", 48, 48 - streamenc.HeaderSize - 16, false},
		{"not a multiple of ecDataBlocks", 1, 0, true},
		{"not a multiple of ecDataBlocks (MaxBlockSize+1)", streamenc.MaxBlockSize + 1, 0, true},
		{"too small (32 bytes)", 32, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := streamenc.MaxPlaintextSizeForTarget(tc.targetSize)
			if tc.wantErr {
				if tc.targetSize >= streamenc.MinCipherTextSize && tc.targetSize%16 != 0 {
					require.ErrorIs(t, err, streamenc.ErrTargetSizeNotAligned)
				} else {
					require.ErrorIs(t, err, streamenc.ErrCiphertextSizeTooSmall)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}

	// Forward+inverse roundtrip: for the maximum plaintext of each block count,
	// Overhead(maxPlain)+maxPlain gives the exact ciphertext size; PlaintextSizeForTarget
	// must recover maxPlain from it.
	for n := 1; n <= 5; n++ {
		var maxPlain int
		if n == 1 {
			maxPlain = streamenc.Block0DataSize
		} else {
			maxPlain = streamenc.Block0DataSize + (n-1)*streamenc.BlockNDataSize
		}
		targetSize := maxPlain + streamenc.Overhead(maxPlain)
		got, err := streamenc.MaxPlaintextSizeForTarget(targetSize)
		require.NoError(t, err)
		assert.Equal(t, maxPlain, got, "roundtrip failed for n=%d blocks", n)
	}
}

func TestEncryptDecryptWithTargetSize(t *testing.T) {
	key := []byte("test-master-key")
	info := "test-info"

	targetSizes := []struct {
		name       string
		targetSize int
	}{
		{"one block", streamenc.MaxBlockSize},
		{"two blocks", 2 * streamenc.MaxBlockSize},
		{"three blocks", 3 * streamenc.MaxBlockSize},
		{"1 MB target size", 1 * 1024 * 1024},
		{"8 MB target size", 8 * 1024 * 1024},
		{"16 MB target size", 16 * 1024 * 1024},
		{"32 MB target size", 32 * 1024 * 1024},
	}

	for _, tc := range targetSizes {
		t.Run(tc.name, func(t *testing.T) {
			plaintextSize, err := streamenc.MaxPlaintextSizeForTarget(tc.targetSize)
			require.NoError(t, err)

			original := testrand.Bytes(t, int64(plaintextSize))

			buf := make([]byte, plaintextSize, tc.targetSize)
			copy(buf, original)

			n, err := streamenc.Encrypt(key, buf, info)
			require.NoError(t, err)
			require.Equal(t, 0, n%16, "encrypted size must be divisible by 16")
			require.LessOrEqual(t, n, tc.targetSize, "encrypted size must not exceed target")

			ciphertext := buf[:n]
			recovered, err := streamenc.DecryptAllBlocks(key, ciphertext, info, streamenc.Version)
			require.NoError(t, err)
			require.Equal(t, original, ciphertext[:recovered])

			overhead := n - plaintextSize
			t.Logf("target size %d, plaintext size %d, actual encrypted size %d, overhead %d bytes (%.2f%%)",
				tc.targetSize, plaintextSize, n, overhead, float64(overhead)/float64(n)*100)
		})
	}
}

func TestCiphertextDivisibleBy16(t *testing.T) {
	key := []byte("test-key")
	info := "info"

	sizes := []int{
		1, 2, 15, 16, 17, 31, 32, 33,
		100, 999, 1000, 1001,
		streamenc.Block0DataSize - 1,
		streamenc.Block0DataSize,
		streamenc.Block0DataSize + 1,
		streamenc.Block0DataSize + streamenc.BlockNDataSize - 1,
		streamenc.Block0DataSize + streamenc.BlockNDataSize,
		streamenc.Block0DataSize + streamenc.BlockNDataSize + 1,
		1 << 20,
	}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			original := testrand.Bytes(t, int64(size))
			ct := encryptHelper(t, key, original, info)
			require.Equal(t, 0, len(ct)%16, "ciphertext length %d is not divisible by 16", len(ct))
		})
	}
}

func TestLastBlockZeroPadding(t *testing.T) {
	key := []byte("test-key")
	info := "info"

	cases := []struct {
		name          string
		plaintextSize int
	}{
		// single block: total = HeaderSize + lastData + pad + tagSize; pad = (16 - (HeaderSize+lastData)%16) % 16
		// 1 byte: (17+1)%16=2, pad=14 > 0
		{"single block, 1 byte", 1},
		// 14 bytes: (17+14)%16=15, pad=1 > 0
		{"single block, 14 bytes", 14},
		// multi-block: pad = (16 - lastData%16) % 16
		// Block0DataSize+1: lastData=1, pad=15 > 0
		{"multi-block, last data 1 byte", streamenc.Block0DataSize + 1},
		// Block0DataSize+15: lastData=15, pad=1 > 0
		{"multi-block, last data 15 bytes", streamenc.Block0DataSize + 15},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := testrand.Bytes(t, int64(tc.plaintextSize))
			ct := encryptHelper(t, key, original, info)

			_, initialNonce, plaintextSize32, err := streamenc.ParseHeader(ct)
			require.NoError(t, err)

			plaintextSize := int(plaintextSize32)
			numBlocks := streamenc.NumBlocks(plaintextSize)
			lastBlockIndex := numBlocks - 1
			lastData := streamenc.BlockDataSize(plaintextSize, lastBlockIndex)

			// Decrypt the last block's ciphertext directly to get the padded plaintext.
			gcm, err := encryption.GCMCipher(key, info)
			require.NoError(t, err)

			lastBlockData := streamenc.EncryptedBlock(ct, lastBlockIndex)
			require.NotNil(t, lastBlockData)

			lastBlockCipher := lastBlockData
			if lastBlockIndex == 0 {
				lastBlockCipher = lastBlockData[streamenc.HeaderSize:]
			}

			blockNonce := streamenc.BlockNonce(initialNonce, lastBlockIndex)
			plainPadded, err := gcm.Open(nil, blockNonce[:], lastBlockCipher, nil)
			require.NoError(t, err)

			// Bytes [lastData:] should be zero padding.
			require.Greater(t, len(plainPadded), lastData, "expected padding bytes present")
			for i := lastData; i < len(plainPadded); i++ {
				require.Equal(t, byte(0), plainPadded[i], "padding byte at index %d should be zero", i)
			}
		})
	}
}

func TestDecryptBlockVersionMismatch(t *testing.T) {
	key := []byte("test-key")
	info := "info"
	original := testrand.Bytes(t, int64(1000))

	ciphertext := encryptHelper(t, key, original, info)

	_, initialNonce, _, err := streamenc.ParseHeader(ciphertext)
	require.NoError(t, err)

	gcm, err := encryption.GCMCipher(key, info)
	require.NoError(t, err)

	block0 := streamenc.EncryptedBlock(ciphertext, 0)
	wrongVersion := uint8(streamenc.Version + 1)

	_, err = streamenc.DecryptBlock(gcm, block0, initialNonce, wrongVersion, 0, len(original))
	require.ErrorIs(t, err, streamenc.ErrVersionMismatch)
}

func TestDecryptAllBlocksVersionMismatch(t *testing.T) {
	key := []byte("test-key")
	info := "info"
	original := testrand.Bytes(t, int64(1000))

	ciphertext := encryptHelper(t, key, original, info)
	wrongVersion := uint8(streamenc.Version + 1)

	_, err := streamenc.DecryptAllBlocks(key, ciphertext, info, wrongVersion)
	require.ErrorIs(t, err, streamenc.ErrVersionMismatch)
}

func BenchmarkEncryptComparison(b *testing.B) {
	key := []byte("benchmark-master-key")
	info := "benchmark-info"

	// targetSize is the desired ciphertext size (must be a multiple of MaxBlockSize).
	sizes := []struct {
		name       string
		targetSize int
	}{
		{"32KB", 1 * streamenc.MaxBlockSize},
		{"64KB", 2 * streamenc.MaxBlockSize},
		{"1MB", 32 * streamenc.MaxBlockSize},
		{"8MB", 256 * streamenc.MaxBlockSize},
		{"16MB", 512 * streamenc.MaxBlockSize},
		{"32MB", 1024 * streamenc.MaxBlockSize},
	}

	for _, tc := range sizes {
		plaintextSize, err := streamenc.MaxPlaintextSizeForTarget(tc.targetSize)
		require.NoError(b, err)

		original := testrand.Bytes(b, int64(plaintextSize))

		b.Run(tc.name+"/normal", func(b *testing.B) {
			b.ReportAllocs()
			buf := make([]byte, plaintextSize, plaintextSize+encryption.Overhead)
			var result []byte
			for b.Loop() {
				b.StopTimer()
				copy(buf, original)
				buf = buf[:plaintextSize]
				b.StartTimer()
				result, err = encryption.Encrypt(key, buf, info)
			}
			_ = result
			require.NoError(b, err)
		})

		b.Run(tc.name+"/stream", func(b *testing.B) {
			b.ReportAllocs()
			buf := make([]byte, plaintextSize, tc.targetSize)
			var result int
			for b.Loop() {
				b.StopTimer()
				copy(buf, original)
				buf = buf[:plaintextSize]
				b.StartTimer()
				result, err = streamenc.Encrypt(key, buf, info)
			}
			_ = result
			require.NoError(b, err)
		})
	}
}

func BenchmarkDecryptComparison(b *testing.B) {
	key := []byte("benchmark-master-key")
	info := "benchmark-info"

	// targetSize is the desired ciphertext size (must be a multiple of MaxBlockSize).
	sizes := []struct {
		name       string
		targetSize int
	}{
		{"32KB", 1 * streamenc.MaxBlockSize},
		{"64KB", 2 * streamenc.MaxBlockSize},
		{"1MB", 32 * streamenc.MaxBlockSize},
		{"8MB", 256 * streamenc.MaxBlockSize},
		{"16MB", 512 * streamenc.MaxBlockSize},
		{"32MB", 1024 * streamenc.MaxBlockSize},
	}

	for _, tc := range sizes {
		plaintextSize, err := streamenc.MaxPlaintextSizeForTarget(tc.targetSize)
		require.NoError(b, err)

		original := testrand.Bytes(b, int64(plaintextSize))

		normalBuf := make([]byte, plaintextSize, plaintextSize+encryption.Overhead)
		copy(normalBuf, original)
		normalCiphertext, err := encryption.Encrypt(key, normalBuf, info)
		require.NoError(b, err)

		b.Run(tc.name+"/normal", func(b *testing.B) {
			b.ReportAllocs()
			cipherBuf := make([]byte, len(normalCiphertext))
			var result []byte
			for b.Loop() {
				b.StopTimer()
				copy(cipherBuf, normalCiphertext)
				b.StartTimer()
				result, err = encryption.Decrypt(key, cipherBuf, info)
			}
			_ = result
			require.NoError(b, err)
		})

		streamBuf := make([]byte, plaintextSize, tc.targetSize)
		copy(streamBuf, original)
		streamTotal, err := streamenc.Encrypt(key, streamBuf, info)
		require.NoError(b, err)
		streamCiphertext := streamBuf[:streamTotal]

		b.Run(tc.name+"/stream", func(b *testing.B) {
			b.ReportAllocs()
			cipherBuf := make([]byte, len(streamCiphertext))
			var result int
			for b.Loop() {
				b.StopTimer()
				copy(cipherBuf, streamCiphertext)
				b.StartTimer()
				result, err = streamenc.DecryptAllBlocks(key, cipherBuf, info, streamenc.Version)
			}
			_ = result
			require.NoError(b, err)
		})
	}
}
