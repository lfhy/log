# log

一个简单易用的Go语言日志库，可以零配置使用，基于gf框架的日志模块。

[![GoDoc](https://pkg.go.dev/badge/github.com/lfhy/log.svg)](https://pkg.go.dev/github.com/lfhy/log)
[![License](https://img.shields.io/github/license/lfhy/log)](https://github.com/lfhy/log/blob/master/LICENSE)

## 特性

- 🚀 **零配置**: 开箱即用，无需复杂配置
- 🌈 **多级别日志**: 支持 Debug、Info、Warn、Error 等多种日志级别
- 🎨 **彩色输出**: 终端日志彩色高亮显示
- 🔧 **上下文支持**: 支持带context的日志记录
- 🗃️ **文件存储**: 支持日志写入文件并按日期分割
- 🔄 **ORM集成**: 内置GORM日志支持
- 📦 **易于集成**: 基于gf框架，稳定可靠

## 安装

```bash
go get github.com/lfhy/log
```

## 快速上手

### 基础使用

```go
package main

import "github.com/lfhy/log"

func main() {
	log.SetLevel(log.LEVEL_ALL)
	log.Debug("Debug日志")
	log.Warn("Warn日志")
	log.Info("Info日志")
	log.Error("Error日志")
}
```

输出示例：
```bash
go run example/print/print.go
2025-12-18T15:34:25.895+08:00 [DEBU] Debug日志
2025-12-18T15:34:25.895+08:00 [WARN] Warn日志
2025-12-18T15:34:25.895+08:00 [INFO] Info日志
2025-12-18T15:34:25.895+08:00 [ERRO] Error日志 
Stack:
1.  github.com/lfhy/log.Error
    /Users/3000y/project/log/print.go:52
2.  main.main
    /Users/3000y/project/log/example/print/print.go:12
```

### 高级使用

#### 初始化配置

```go
// 初始化日志配置
log.Init(&log.LogOption{
    Debug:    true,           // 是否开启调试模式
    LogDir:   "./logs",       // 日志文件目录
    LogMark:  "app",          // 日志文件前缀
    NoStdout: false,          // 是否禁用标准输出
})
```

#### 使用上下文

```go
// 创建带标记的日志实例
logger := log.New("API")

// 或者使用上下文
ctx := context.WithValue(context.Background(), log.CtxMark, "API")
logger := log.Ctx(ctx)

// 记录日志
logger.Info("这是一条带标记的日志")
```

#### 格式化日志

```go
log.Infof("用户 %s 登录成功", username)
log.Errorf("数据库连接失败: %v", err)
```

#### ORM日志集成

```go
// 创建ORM日志实例
ormLog := log.NewORMLog(&log.ORMLogOption{
    LogOption: log.LogOption{
        Debug:    true,
        LogDir:   "./logs",
        LogMark:  "db",
    },
    Level: logger.Info, // GORM日志级别
})

// 在GORM中使用
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: ormLog,
})
```

## API 文档

完整的API文档请参考 [GoDoc](https://pkg.go.dev/github.com/lfhy/log)

### 日志级别

- `LEVEL_ALL` - 所有级别日志
- `LEVEL_DEBU` - 调试级别
- `LEVEL_INFO` - 信息级别
- `LEVEL_WARN` - 警告级别
- `LEVEL_ERRO` - 错误级别
- `LEVEL_CRIT` - 严重错误级别

### 主要方法

- `log.Debug()` / `log.Debugf()` / `log.Debugln()` - 调试日志
- `log.Info()` / `log.Infof()` / `log.Infoln()` - 信息日志
- `log.Warn()` / `log.Warnf()` / `log.Warnln()` - 警告日志
- `log.Error()` / `log.Errorf()` / `log.Errorln()` - 错误日志
- `log.SetLevel()` - 设置日志级别
- `log.Init()` - 初始化日志配置
- `log.New()` - 创建新的日志实例

## 许可证

MIT License，详情请查看 [LICENSE](https://github.com/lfhy/log/blob/master/LICENSE) 文件。