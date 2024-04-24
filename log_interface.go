package log

import (
	"log"
	"os"
	"sync"
)

type LoggerOption func(*Logger)

type Logger struct {
	mark          string          // 日志标签
	logLevel      int             // 日志等级 分为 error warn info debug speed
	short         bool            // 短打印
	caller        int             // 打印行
	Logout        *log.Logger     // 日志输出
	logoutLock    sync.RWMutex    // 日志配置锁
	isStdout      bool            // 是否是STD输出
	StdLogout     *log.Logger     // 终端输出
	useStdout     bool            // 禁用Stdloger进行输出
	push          bool            // 推送开关 开启后会将日志内容推送到管道中
	pushChannel   *chan (LogInfo) // 推送管道
	logfileName   string          // 日志名称
	logfileDir    string          // 日志文件目录
	logfile       *os.File        // 日志文件句柄
	logfileSuffix string          // 日志文件后缀
	redirectPanic bool            // 重定向错误
	disableMark   bool            // 关闭标签
	noCodeLine    bool            // 不打印代码行
}

// 新建Log
func NewLogger(opts ...LoggerOption) *Logger {
	l := &Logger{mark: LogDefaultPrefix, logLevel: INFOLevel, caller: 2}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// 新建一个带标识的日志 标识是唯一字符串
func NewLoggerWithMark(mark string, level int) *Logger {
	return NewLogger(WithMark(mark), WithLogLevel(level))
}
