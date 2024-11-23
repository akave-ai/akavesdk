// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package sdk is the Akave SDK.
package sdk

import (
	"time"

	"github.com/ipfs/go-cid"
)

// BucketCreateResult is the result of bucket creation.
type BucketCreateResult struct {
	ID        string
	CreatedAt time.Time
	TxHash    string
}

// Bucket is a bucket.
type Bucket struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

// Chunk is a piece of metadata of some file.
type Chunk struct {
	CID   string
	Size  int64
	Index int64
}

// FileBlock is a piece of metadata of some file.
type FileBlock struct {
	CID  string
	Data []byte

	Permit      string
	NodeAddress string
	NodeID      string
}

// FileBlockSP is a piece of metadata of some file.
type FileBlockSP struct {
	CID  string
	Data []byte

	Permit      string
	NodeAddress string
	NodeID      string

	SPBaseURL string
}

// FileUpload represents a file and some metadata.
type FileUpload struct {
	RootCID    string
	BucketName string
	FileName   string
	FileSize   int64
	Blocks     []FileBlock
	TxHash     string
}

// FileDownload represents a file download and some metadata.
type FileDownload struct {
	BucketName string
	FileName   string
	Blocks     []FileBlock
}

// FileDownloadSP represents a file download and some metadata.
type FileDownloadSP struct {
	BucketName string
	FileName   string
	Blocks     []FileBlockSP
}

// FileListItem contains bucket file list file meta information.
type FileListItem struct {
	RootCID   string
	Name      string
	Size      int64
	CreatedAt time.Time
}

// FileMeta contains single file meta information.
type FileMeta struct {
	RootCID   string
	Name      string
	Size      int64
	CreatedAt time.Time
}

// FileUploadV2 contains single file meta information.
type FileUploadV2 struct {
	BucketID  string
	Name      string
	StreamID  string
	CreatedAt time.Time
}

// FileChunkUploadV2 contains single file chunk meta information.
type FileChunkUploadV2 struct {
	StreamID      string
	Index         int64
	ChunkCID      cid.Cid
	RawDataSize   uint64
	ProtoNodeSize uint64
	Blocks        []FileBlock
}

// FileDownloadV2 contains single file meta information.
type FileDownloadV2 struct {
	StreamID string
	BucketID string
	Name     string
	Chunks   []Chunk
}

// FileChunkDownloadV2 contains single file chunk meta information.
type FileChunkDownloadV2 struct {
	CID    string
	Index  int64
	Size   int64
	Blocks []FileBlock
}

// FileMetaV2 contains single file meta information.
type FileMetaV2 struct {
	StreamID   string
	RootCID    string
	BucketID   string
	Name       string
	Size       int64
	CreatedAt  time.Time
	CommitedAt time.Time
}
