package progress

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type ProgressCallback func(Progress)

type Progress struct {
	ID string
	TotalBytes int64
	ProcessedBytes int64
	TotalChunks int32
	ProcessedChunks int32
	TransferRate float64
	EstimatedTimeRemaining float64
	StartTime time.Time
	LastUpdate time.Time
	Error error
}

type Tracker struct {
	mu sync.RWMutex

	operations map[string]*operationState
	callback ProgressCallback
	updateInterval time.Duration
}

type operationState struct {
	progress Progress
	done chan struct{}
	rates []float64
	maxRates int
}

func New(callback ProgressCallback, updateInterval time.Duration) *Tracker {
	return &Tracker{
		operations:     make(map[string]*operationState),
		callback:      callback,
		updateInterval: updateInterval,
	}
}

func (t *Tracker) StartOperation(id string, totalBytes int64, totalChunks int32) {
	t.mu.Lock()
	defer t.mu.Unlock()

	state := &operationState{
		progress: Progress{
			ID:          id,
			TotalBytes:  totalBytes,
			TotalChunks: totalChunks,
			StartTime:   time.Now(),
			LastUpdate:  time.Now(),
		},
		done:     make(chan struct{}),
		maxRates: 10,
	}

	t.operations[id] = state

	go t.monitorProgress(id)
}

func (t *Tracker) UpdateProgress(id string, bytesProcessed int64, chunksProcessed int32) {
	t.mu.Lock()
	defer t.mu.Unlock()

	state, exists := t.operations[id]
	if !exists {
		return
	}

	now := time.Now()
	elapsed := now.Sub(state.progress.LastUpdate).Seconds()
	if elapsed > 0 {
		rate := float64(bytesProcessed) / elapsed
		state.rates = append(state.rates, rate)
		if len(state.rates) > state.maxRates {
			state.rates = state.rates[1:]
		}
	}

	state.progress.ProcessedBytes += bytesProcessed
	state.progress.ProcessedChunks += chunksProcessed
	state.progress.LastUpdate = now

	var totalRate float64
	for _, rate := range state.rates {
		totalRate += rate
	}
	if len(state.rates) > 0 {
		state.progress.TransferRate = totalRate / float64(len(state.rates))
	}

	if state.progress.TransferRate > 0 {
		remainingBytes := state.progress.TotalBytes - state.progress.ProcessedBytes
		state.progress.EstimatedTimeRemaining = float64(remainingBytes) / state.progress.TransferRate
	}

	if t.callback != nil {
		t.callback(state.progress)
	}
}

func (t *Tracker) CompleteOperation(id string, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	state, exists := t.operations[id]
	if !exists {
		return
	}

	state.progress.Error = err
	close(state.done)
	delete(t.operations, id)

	if t.callback != nil {
		t.callback(state.progress)
	}
}

func (t *Tracker) GetProgress(id string) (Progress, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if state, exists := t.operations[id]; exists {
		return state.progress, true
	}
	return Progress{}, false
}

func (t *Tracker) monitorProgress(id string) {
	ticker := time.NewTicker(t.updateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.mu.RLock()
			state, exists := t.operations[id]
			if !exists {
				t.mu.RUnlock()
				return
			}
			progress := state.progress
			done := state.done
			t.mu.RUnlock()

			if t.callback != nil {
				t.callback(progress)
			}

		case <-t.operations[id].done:
			return
		}
	}
}
