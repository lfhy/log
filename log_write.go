package log

func (l *Logger) SetWriterLogLevel(logLevel LogLevel) {
	l.writerLogLevel = logLevel
}

func (l *Logger) Write(b []byte) (n int, err error) {
	str := string(b)
	switch l.writerLogLevel {
	case LogLevelSpeed:
		l.PSpeed(str)
	case LogLevelDebug:
		l.PDebug(str)
	case LogLevelWarn:
		l.PWarn(str)
	case LogLevelError:
		l.PError(str)
		// 默认Info等级
	default:
		l.PInfo(str)
	}
	return len(b), nil
}
