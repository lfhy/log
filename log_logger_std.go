package log

import "log"

// 获取终端输出器
// 终端输出的日志不会带有代码行
// 可以设置终端不进行输出
func (l *Logger) GetStdLogger() *log.Logger {
	if l.StdLogout == nil {
		return l.GetLogger()
	}
	return l.StdLogout
}

func (l *Logger) SetStdLogger(logger *log.Logger) {
	l.StdLogout = logger
}

func WithStdLogger(logger *log.Logger) LoggerOption {
	return func(l *Logger) {
		l.SetStdLogger(logger)
	}
}

// 禁用StdLogger输出
func (l *Logger) DisableStdlogger(disable bool) {
	l.useStdout = !disable
}

func WithDisableStdlogger() LoggerOption {
	return func(l *Logger) {
		l.DisableStdlogger(true)
	}
}
