package ratelimit

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter *rate.Limiter
	mu      sync.Mutex
}

func New(bytesPerSecond int64) *RateLimiter {
	if bytesPerSecond <= 0 {
		return nil
	}

	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(bytesPerSecond), int(bytesPerSecond)), 
	}
}

func (r *RateLimiter) WaitN(ctx context.Context, n int) error {
	if r == nil || r.limiter == nil {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	return r.limiter.WaitN(ctx, n)
}

func (r *RateLimiter) Available() int {
	if r == nil || r.limiter == nil {
		return int(^uint(0) >> 1)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	return r.limiter.Burst() - r.limiter.Tokens()
}

func (r *RateLimiter) SetRate(bytesPerSecond int64) {
	if r == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.limiter.SetLimit(rate.Limit(bytesPerSecond))
	r.limiter.SetBurst(int(bytesPerSecond))
}

type WithBackpressure struct {
	reader io.Reader
	rl     *RateLimiter
	ctx    context.Context
}

func NewReaderWithBackpressure(ctx context.Context, r io.Reader, rl *RateLimiter) *WithBackpressure {
	return &WithBackpressure{
		reader: r,
		rl:     rl,
		ctx:    ctx,
	}
}

func (r *WithBackpressure) Read(p []byte) (n int, err error) {
	if err := r.rl.WaitN(r.ctx, len(p)); err != nil {
		return 0, err
	}
	return r.reader.Read(p)
}
