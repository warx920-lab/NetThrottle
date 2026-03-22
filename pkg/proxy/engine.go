package proxy

import (
	"io"
	"net"
	"netthrottle/pkg/limiter"
	"time"
)

func Transfer(src, dst net.Conn, lb *limiter.TokenBucket, delay int) {
	buf := make([]byte, 32*1024) 
	for {
		n, err := src.Read(buf)
		if n > 0 {
			// 注入延遲
			if delay > 0 {
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
			// 限流
			lb.Consume(n)
			// 轉發
			_, err = dst.Write(buf[:n])
			if err != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}
}