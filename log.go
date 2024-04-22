package log

var defaultLog *Logger

// 获取日志输出
func GetLogout() *Logger {
	if defaultLog == nil {
		defaultLog = NewLogger(WithCodeCaller(3), WithDisableMark(), WithShortout())
	}
	return defaultLog
}

func SetDefaultLog(log *Logger) {
	defaultLog = log
}

// 设置日志等级
func SetLogLevel(level int) {
	GetLogout().SetLogLevel(level)
}
