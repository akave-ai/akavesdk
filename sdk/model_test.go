// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package sdk

import (
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIPCFileUploadStateConcurrency(t *testing.T) {
	upload, err := NewIPCFileUpload("test-bucket", "test-file")
	require.NoError(t, err)
	require.NotNil(t, upload)

	mockTx := types.NewTransaction(0, [20]byte{}, nil, 0, nil, nil)

	var wg sync.WaitGroup
	var deleteWg sync.WaitGroup

	deleteWg.Add(1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 10 {
			if i == 6 {
				deleteWg.Done()
			}

			chunk := createMockChunk(int64(i))
			err := upload.state.preCreateChunk(chunk, mockTx)
			require.NoError(t, err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		deleteWg.Wait()
		for i := range 5 {
			chunk := createMockChunk(int64(i))
			upload.state.chunkUploaded(chunk)
		}
	}()

	wg.Wait()

	assert.Equal(t, int64(10), upload.state.chunkCount, "chunk count should be 10")
	assert.Equal(t, int64(10*1024), upload.state.actualFileSize, "actual file size should be 10 * 1024")
	assert.Equal(t, int64(10*1100), upload.state.encodedFileSize, "encoded file size should be 10 * 1100")

	preCreatedChunks := upload.state.listPreCreatedChunks()
	assert.Equal(t, 5, len(preCreatedChunks), "should have 5 remaining pre-created chunks")
}

func createMockChunk(index int64) IPCFileChunkUploadV2 {
	cidStr := "bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi"
	mockCID, _ := cid.Decode(cidStr)

	return IPCFileChunkUploadV2{
		Index:       index,
		ChunkCID:    mockCID,
		ActualSize:  1024,
		RawDataSize: 1024,
		EncodedSize: 1100,
		BucketID:    [32]byte{1, 2, 3},
		FileName:    "test-file",
	}
}
