package limiter

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	// 設定每秒 1000 字節
	tb := NewTokenBucket(1000, 1000)
	
	start := time.Now()
	// 消耗 2000 字節，預計至少需要 1 秒
	tb.Consume(1000)
	tb.Consume(1000)
	elapsed := time.Since(start)

	if elapsed < 900*time.Millisecond {
		t.Errorf("Limiter was too fast, elapsed: %v", elapsed)
	}
}