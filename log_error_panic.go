package log

func WithRediectErrorToLog() LoggerOption {
	return func(l *Logger) {
		l.redirectPanic = true
	}
}

func SetRediectErrorToLog() {
	l := GetLogout()
	l.redirectPanic = true
}
