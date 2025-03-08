# log
一个简单的日志库 可以零配置使用

# 快速上手
```go
package main

import "github.com/lfhy/log"

func main() {
	log.SetLogLevel(log.LogLevelDebug)
	log.Debug("Debug日志")
	log.Warn("Warn日志")
	log.Info("Info日志")
	log.Error("Error日志")
}
```
# 使用文档
[GoDoc](https://pkg.go.dev/github.com/lfhy/log)