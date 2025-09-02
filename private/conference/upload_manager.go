package concurrency

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type UploadManager struct {
	maxUploads int64
	activeUploads atomic.Int64
	sem *semaphore.Weighted
	mu sync.RWMutex
	uploads map[string]*UploadStatus
}

type UploadStatus struct {
	ID            string
	TotalBytes    int64
	UploadedBytes atomic.Int64
	ChunksTotal   int32
	ChunksUploaded atomic.Int32
	Error         error
}

func New(maxConcurrentUploads int64) *UploadManager {
	return &UploadManager{
		maxUploads: maxConcurrentUploads,
		sem:        semaphore.NewWeighted(maxConcurrentUploads),
		uploads:    make(map[string]*UploadStatus),
	}
}

func (m *UploadManager) StartUpload(ctx context.Context, id string, totalBytes int64, totalChunks int32) (*UploadStatus, error) {
	if err := m.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}

	status := &UploadStatus{
		ID:          id,
		TotalBytes:  totalBytes,
		ChunksTotal: totalChunks,
	}

	m.mu.Lock()
	m.uploads[id] = status
	m.mu.Unlock()

	m.activeUploads.Add(1)
	return status, nil
}

func (m *UploadManager) UpdateProgress(id string, bytesUploaded int64, chunksUploaded int32) {
	m.mu.RLock()
	status, exists := m.uploads[id]
	m.mu.RUnlock()

	if exists {
		status.UploadedBytes.Add(bytesUploaded)
		if chunksUploaded > 0 {
			status.ChunksUploaded.Add(chunksUploaded)
		}
	}
}

func (m *UploadManager) CompleteUpload(id string, err error) {
	m.mu.Lock()
	if status, exists := m.uploads[id]; exists {
		status.Error = err
		delete(m.uploads, id)
	}
	m.mu.Unlock()

	m.activeUploads.Add(-1)
	m.sem.Release(1)
}

func (m *UploadManager) GetStatus(id string) *UploadStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.uploads[id]
}

func (m *UploadManager) GetActiveUploads() int64 {
	return m.activeUploads.Load()
}

func (m *UploadManager) GetAllStatuses() map[string]*UploadStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()

	statuses := make(map[string]*UploadStatus, len(m.uploads))
	for id, status := range m.uploads {
		statuses[id] = status
	}
	return statuses
}
