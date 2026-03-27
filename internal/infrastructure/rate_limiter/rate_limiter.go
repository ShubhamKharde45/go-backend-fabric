package ratelimiter

import (
	"sync"
	"time"

	"github.com/ShubhamKharde45/rate_limiter/internal/domain"
)

type RateLimiter struct {
	store domain.Cache[*domain.Bucket]
	mu    *sync.Mutex
}

func NewRateLimiter(Store domain.Cache[*domain.Bucket], mu *sync.Mutex) *RateLimiter {
	return &RateLimiter{
		store: Store,
		mu:    mu,
	}
}

func (rl *RateLimiter) refill(bucket *domain.Bucket) {

	now := time.Now()
	elapsedTime := now.Sub(bucket.LastRefill).Seconds()
	bucket.Tokens += elapsedTime * bucket.Rate

	if bucket.Tokens > bucket.Capacity {
		bucket.Tokens = bucket.Capacity
	}

	bucket.LastRefill = now

}

func (rl *RateLimiter) IsAllowed(IP string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	allowed := false

	val, _ := rl.store.Get(IP)

	var bucket *domain.Bucket
	if val == nil {
		bucket = &domain.Bucket{
			Tokens:     100,
			Capacity:   100,
			Rate:       1,
			LastRefill: time.Now(),
		}

	} else {
		bucket = val
	}

	_ = rl.store.Set(IP, bucket)

	rl.refill(bucket)

	if bucket.Tokens >= 1 {
		bucket.Tokens--
		rl.store.Update(IP, bucket)
		allowed = true
	}

	return allowed
}
