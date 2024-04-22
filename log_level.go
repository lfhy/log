package log

// 设置日志等级
func (l *Logger) SetLogLevel(level int) {
	l.logLevel = level
}

// 返回当前日志等级
func (l *Logger) GetLogLevel() int {
	return l.logLevel
}

func WithLogLevel(level int) LoggerOption {
	return func(l *Logger) {
		l.SetLogLevel(level)
	}
}
