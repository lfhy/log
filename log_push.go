package log

// 日志推送
func (l *Logger) SetLogPush(push bool) {
	l.push = push
	// 如果是推送 则初始化管道
	if l.push {
		l.GetLogChannl()
	}
}

func (l *Logger) GetLogChannl() *chan (LogInfo) {
	if l.pushChannel == nil {
		c := make(chan (LogInfo))
		l.pushChannel = &c
	}
	return l.pushChannel
}

func WithLogpush() LoggerOption {
	return func(l *Logger) {
		l.SetLogPush(true)
	}
}
