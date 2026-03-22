\# ⚡ NetThrottle



NetThrottle is a lightweight, high-performance network traffic shaper for developers. It allows you to simulate weak network conditions (latency and bandwidth limits) for any TCP service.



\## Key Features

\- \*\*Precise Throttling:\*\* Implements the Token Bucket algorithm for byte-accurate rate limiting.

\- \*\*Zero Config:\*\* A single binary tool, no root privileges or complex `tc` commands required.

\- \*\*Bi-directional:\*\* Shapes traffic in both directions.



\## Installation

```bash

go build -o netthrottle ./cmd/netthrottle


Usage
Forward local port 8081 to a database on 5432, limiting speed to 128KB/s with 100ms lag:

Bash
./netthrottle -l :8081 -r localhost:5432 -kb 128 -delay 100
Testing
Bash
go test ./pkg/limiter/...

### 7. `.gitignore`
```text
/netthrottle
\*.exe
\*.bin
vendor/
🛠️ 如何編譯與測試
安裝 Go 語言（版本 1.21+）。

在項目根目錄運行測試：

Bash
go test ./pkg/limiter/...
(如果看到 PASS，說明你的算法邏輯完全正確。)

編譯成可執行文件：

Bash
go build -o netthrottle ./cmd/netthrottle

