package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate     float64    // 每秒產生的令牌數 (Bytes/s)
	capacity float64    // 桶容量
	tokens   float64    // 當前令牌數
	lastTick time.Time  // 上次更新時間
	mu       sync.Mutex
}

func NewTokenBucket(rate int, burst int) *TokenBucket {
	return &TokenBucket{
		rate:     float64(rate),
		capacity: float64(burst),
		tokens:   float64(burst),
		lastTick: time.Now(),
	}
}

func (tb *TokenBucket) Consume(n int) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	for {
		now := time.Now()
		// 計算時間差並補充令牌
		tb.tokens += now.Sub(tb.lastTick).Seconds() * tb.rate
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastTick = now

		if tb.tokens >= float64(n) {
			tb.tokens -= float64(n)
			return
		}

		// 令牌不足，計算需要休眠的時間
		needed := float64(n) - tb.tokens
		waitDuration := time.Duration(needed / tb.rate * float64(time.Second))
		
		tb.mu.Unlock()
		time.Sleep(waitDuration)
		tb.mu.Lock()
	}
}