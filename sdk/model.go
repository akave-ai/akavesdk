// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package sdk is the Akave SDK.
package sdk

import (
	"slices"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/exp/maps"
)

// BucketCreateResult is the result of bucket creation.
type BucketCreateResult struct {
	Name      string
	CreatedAt time.Time
}

// Bucket is a bucket.
type Bucket struct {
	Name      string
	CreatedAt time.Time
}

// Chunk is a piece of metadata of some file.
type Chunk struct {
	CID         string
	EncodedSize int64
	Size        int64
	Index       int64
}

// AkaveBlockData is a akavenode block metadata.
type AkaveBlockData struct {
	Permit      string
	NodeAddress string
	NodeID      string
}

// FilecoinBlockData is a filecoin block metadata.
type FilecoinBlockData struct {
	BaseURL string
}

// FileBlockUpload is a piece of metadata of some file used for upload.
type FileBlockUpload struct {
	CID  string
	Data []byte

	Permit      string
	NodeAddress string
	NodeID      string
}

// FileBlockDownload is a piece of metadata of some file used for download.
type FileBlockDownload struct {
	CID  string
	Data []byte

	Filecoin *FilecoinBlockData
	Akave    *AkaveBlockData
}

// FileListItem contains bucket file list file meta information.
type FileListItem struct {
	RootCID     string
	Name        string
	Size        int64
	DataBlocks  int64
	TotalBlocks int64
	CreatedAt   time.Time
}

// FileUpload contains single file meta information.
type FileUpload struct {
	dataCounters

	BucketName  string
	Name        string
	StreamID    string
	DataBlocks  int64
	TotalBlocks int64
	CreatedAt   time.Time
}

// FileChunkUpload contains single file chunk meta information.
type FileChunkUpload struct {
	StreamID      string
	Index         int64
	ChunkCID      cid.Cid
	RawDataSize   uint64
	ProtoNodeSize uint64
	Blocks        []FileBlockUpload
}

// FileDownload contains single file meta information.
type FileDownload struct {
	StreamID    string
	BucketName  string
	Name        string
	DataBlocks  int64
	TotalBlocks int64
	Chunks      []Chunk
}

// FileChunkDownload contains single file chunk meta information.
type FileChunkDownload struct {
	CID         string
	Index       int64
	EncodedSize int64
	Size        int64
	Blocks      []FileBlockDownload
}

// FileMeta contains single file meta information.
type FileMeta struct {
	StreamID    string
	RootCID     string
	BucketName  string
	Name        string
	EncodedSize int64
	Size        int64
	DataBlocks  int64
	TotalBlocks int64
	CreatedAt   time.Time
	CommitedAt  time.Time
}

// IPCBucketCreateResult is the result of ipc bucket creation.
type IPCBucketCreateResult struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

// IPCBucket is an IPC bucket.
type IPCBucket struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

// IPCFileDownload represents an IPC file download and some metadata.
type IPCFileDownload struct {
	dataCounters

	BucketName string
	Name       string
	Chunks     []Chunk
}

// IPCFileListItem contains IPC bucket file list file meta information.
type IPCFileListItem struct {
	RootCID     string
	Name        string
	EncodedSize int64
	ActualSize  int64
	CreatedAt   time.Time
}

// IPCFileMeta contains single IPC file meta information.
type IPCFileMeta struct {
	RootCID     string
	Name        string
	BucketName  string
	EncodedSize int64
	ActualSize  int64
	IsPublic    bool
	CreatedAt   time.Time
}

// IPCFileMetaV2 contains single file meta information.
type IPCFileMetaV2 struct {
	RootCID     string
	BucketName  string
	Name        string
	EncodedSize int64
	Size        int64
	CreatedAt   time.Time
	CommittedAt time.Time
}

// IPCFileChunkUploadV2 contains single file chunk meta information.
type IPCFileChunkUploadV2 struct {
	Index         int64
	ChunkCID      cid.Cid
	ActualSize    int64
	RawDataSize   uint64
	ProtoNodeSize uint64
	Blocks        []FileBlockUpload
	BucketID      [32]byte
	FileName      string
}

// IPCFileUpload contains ipc single file meta information.
type IPCFileUpload struct {
	dataCounters

	BucketName string
	Name       string

	state uploadState
}

// NewIPCFileUpload creates a new IPCFileUpload.
func NewIPCFileUpload(bucketName, name string) (*IPCFileUpload, error) {
	dagRoot, err := NewDAGRoot()
	if err != nil {
		return nil, err
	}

	return &IPCFileUpload{
		dataCounters: newDataCounters(),
		state:        uploadState{dagRoot: dagRoot, preCreatedChunks: make(map[int]chunkWithTx)},
		BucketName:   bucketName,
		Name:         name,
	}, nil
}

type uploadState struct {
	dagRoot *DAGRoot

	mu               sync.RWMutex
	preCreatedChunks map[int]chunkWithTx // chunks created in upload pipline but not yet uploaded

	isCommitted     bool
	chunkCount      int64
	actualFileSize  int64
	encodedFileSize int64
}

// adds a chunk to the pre-created chunks map, updates chunk count and file sizes.
func (us *uploadState) preCreateChunk(chunk IPCFileChunkUploadV2, tx *types.Transaction) error {
	us.mu.Lock()
	us.preCreatedChunks[int(chunk.Index)] = chunkWithTx{chunk: chunk, tx: tx}
	us.mu.Unlock()

	us.chunkCount++
	us.actualFileSize += chunk.ActualSize
	us.encodedFileSize += int64(chunk.ProtoNodeSize)

	return us.dagRoot.AddLink(chunk.ChunkCID, chunk.RawDataSize, chunk.ProtoNodeSize)
}

// removes a chunk from the pre-created chunks map.
func (us *uploadState) chunkUploaded(chunk IPCFileChunkUploadV2) {
	us.mu.Lock()
	delete(us.preCreatedChunks, int(chunk.Index))
	us.mu.Unlock()
}

func (us *uploadState) listPreCreatedChunks() []chunkWithTx {
	us.mu.RLock()
	defer us.mu.RUnlock()

	keys := maps.Keys(us.preCreatedChunks)
	slices.Sort(keys) // sort to preserve order in which they were added

	res := make([]chunkWithTx, 0, len(us.preCreatedChunks))
	for _, key := range keys {
		if chunk, exists := us.preCreatedChunks[key]; exists {
			res = append(res, chunk)
		}
	}

	return res
}

type chunkWithTx struct {
	chunk IPCFileChunkUploadV2
	tx    *types.Transaction
}

// dataCounters tracks the number of chunks, blocks, and bytes processed.
// TODO: use as values; in places where they are used, use structs as pointers.
type dataCounters struct {
	chunksCounter *atomic.Int64
	blocksCounter *atomic.Int64
	bytesCounter  *atomic.Int64
}

// newDataCounters creates a new instance of dataCounters.
func newDataCounters() dataCounters {
	return dataCounters{
		chunksCounter: new(atomic.Int64),
		blocksCounter: new(atomic.Int64),
		bytesCounter:  new(atomic.Int64),
	}
}

// Stats returns current data counters.
func (dc dataCounters) Stats() (int64, int64, int64) {
	return dc.chunksCounter.Load(), dc.blocksCounter.Load(), dc.bytesCounter.Load()
}
