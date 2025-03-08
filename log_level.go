package log

// 设置日志等级
func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

// 返回当前日志等级
func (l *Logger) GetLogLevel() LogLevel {
	return l.logLevel
}

func WithLogLevel(level LogLevel) LoggerOption {
	return func(l *Logger) {
		l.SetLogLevel(level)
	}
}
