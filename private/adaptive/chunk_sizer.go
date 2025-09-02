package adaptive

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/akave-ai/akavesdk/private/memory"
)

const (
	MinChunkSize = 1 * memory.MB
	MaxChunkSize = 64 * memory.MB
	DefaultChunkSize = 32 * memory.MB
	MinUploadSpeed = 1 * memory.MB
	MaxUploadSpeed = 1 * memory.GB
)

type ChunkSizer struct {
	mu sync.RWMutex

	currentSize int64
	uploadSpeed float64
	alpha float64
	lastAdjustment time.Time
	adjustmentInterval time.Duration
	successRate float64
	samples int
}

func New() *ChunkSizer {
	return &ChunkSizer{
		currentSize:        DefaultChunkSize,
		alpha:             0.2, 
		adjustmentInterval: 5 * time.Second,
		successRate:       1.0,
		lastAdjustment:    time.Now(),
	}
}

func (cs *ChunkSizer) GetChunkSize() int64 {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.currentSize
}

func (cs *ChunkSizer) UpdateMetrics(ctx context.Context, bytesUploaded int64, duration time.Duration, success bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	speed := float64(bytesUploaded) / duration.Seconds()

	if cs.samples == 0 {
		cs.uploadSpeed = speed
	} else {
		cs.uploadSpeed = cs.alpha*speed + (1-cs.alpha)*cs.uploadSpeed
	}

	cs.successRate = cs.alpha*float64(boolToInt(success)) + (1-cs.alpha)*cs.successRate

	cs.samples++

	if time.Since(cs.lastAdjustment) < cs.adjustmentInterval {
		return
	}

	cs.adjustChunkSize()
	cs.lastAdjustment = time.Now()
}

func (cs *ChunkSizer) adjustChunkSize() {
	multiplier := 1.0
	if cs.successRate > 0.95 {
		multiplier = 1.25
	} else if cs.successRate < 0.8 {
		multiplier = 0.75
	}

	speedMultiplier := math.Log10(cs.uploadSpeed/MinUploadSpeed) / math.Log10(MaxUploadSpeed/MinUploadSpeed)
	speedMultiplier = math.Max(0.5, math.Min(1.5, speedMultiplier))
	multiplier *= speedMultiplier

	newSize := int64(float64(cs.currentSize) * multiplier)

	if newSize < MinChunkSize {
		newSize = MinChunkSize
	} else if newSize > MaxChunkSize {
		newSize = MaxChunkSize
	}

	cs.currentSize = newSize
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
