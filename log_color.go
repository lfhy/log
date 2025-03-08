package log

// 设置输出不打印颜色
func (l *Logger) SetNoColor(value bool) {
	l.noColor = value
}

// 设置保存的日志不打印颜色
// func (l *Logger) SetLogFileNoColor(value bool) {
// 	l.logFileNoColor = value
// }

// 设置输出不打印颜色
func SetNoColor(value bool) {
	GetLogout().SetNoColor(value)
}

// 设置保存的日志不打印颜色
// func SetLogFileNoColor(value bool) {
// 	GetLogout().SetLogFileNoColor(value)
// }

// 设置输出不打印颜色
func WithNoColor() LoggerOption {
	return func(l *Logger) {
		l.SetNoColor(true)
	}
}

// 设置保存的日志不打印颜色
// func WithLogFileNoColor() LoggerOption {
// 	return func(l *Logger) {
// 		l.SetLogFileNoColor(true)
// 	}
// }
