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
	RootCID   string
	Name      string
	Size      int64
	CreatedAt time.Time
}

// FileUpload contains single file meta information.
type FileUpload struct {
	BucketName string
	Name       string
	StreamID   string
	CreatedAt  time.Time
}

// FileChunkUpload contains single file chunk meta information.
type FileChunkUpload struct {
	StreamID      string
	Index         int64
	ChunkCID      cid.Cid
	ActualSize    int64
	RawDataSize   uint64
	ProtoNodeSize uint64
	Blocks        []FileBlockUpload
}

// FileDownload contains single file meta information.
type FileDownload struct {
	StreamID   string
	BucketName string
	Name       string
	Chunks     []Chunk
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
	BucketName string
	Name       string
	Chunks     []Chunk
}

// IPCFileListItem contains IPC bucket file list file meta information.
type IPCFileListItem struct {
	RootCID     string
	Name        string
	EncodedSize int64
	CreatedAt   time.Time
}

// IPCFileMeta contains single IPC file meta information.
type IPCFileMeta struct {
	RootCID     string
	Name        string
	BucketName  string
	EncodedSize int64
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
