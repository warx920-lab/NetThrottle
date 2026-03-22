package limiter

import (
	"testing"
)

// BenchmarkTokenBucket_Consume 測試在高併發下消耗令牌的性能
func BenchmarkTokenBucket_Consume(b *testing.B) {
	// 設置一個極大的速率，確保測試過程中不會因為等待令牌而阻塞
	// 這樣我們測量的是算法本身的運算開銷
	rate := 1000 * 1024 * 1024 // 1GB/s
	tb := NewTokenBucket(rate, rate)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 模擬每次請求 1KB 的數據量
			tb.Consume(1024)
		}
	})
}

Gemini 說了
這是一個非常專業的舉動。在 GitHub 項目中加入 Benchmark（基準測試） 不僅能證明你的代碼性能優越，還能向潛在的貢獻者展示你對「高性能」的嚴肅態度。

以下是為你準備的 Benchmark 代碼和 Reddit 貼文文稿。

1. Go Benchmark 代碼
請在你的項目中建立文件 pkg/limiter/token_bucket_benchmark_test.go。

這段代碼會測試在高併發（多個 Goroutine 同時請求令牌）的情況下，你的令牌桶算法處理速度有多快。

Go
package limiter

import (
	"testing"
)

// BenchmarkTokenBucket_Consume 測試在高併發下消耗令牌的性能
func BenchmarkTokenBucket_Consume(b *testing.B) {
	// 設置一個極大的速率，確保測試過程中不會因為等待令牌而阻塞
	// 這樣我們測量的是算法本身的運算開銷
	rate := 1000 * 1024 * 1024 // 1GB/s
	tb := NewTokenBucket(rate, rate)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 模擬每次請求 1KB 的數據量
			tb.Consume(1024)
		}
	})
}
