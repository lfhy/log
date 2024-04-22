package log

// 打印Warn日志
func Warn(v ...any) any {
	return GetLogout().Warn(v...)
}

func Warnln(v ...any) any {
	return GetLogout().Warnln(v...)
}

func Warnf(format string, args ...interface{}) string {
	return GetLogout().Warnf(format, args...)
}

func WarnCall(call int, args ...any) any {
	return GetLogout().WarnCall(call, args...)
}

func WarnlnCall(call int, args ...any) any {
	return GetLogout().WarnlnCall(call, args...)
}

func WarnfCall(call int, format string, args ...interface{}) string {
	return GetLogout().WarnfCall(call, format, args...)
}

// 标记日志
func MWarn(mark string, v ...any) any {
	return GetLogout().MWarn(mark, v...)
}

func MWarnln(mark string, v ...any) any {
	return GetLogout().MWarnln(mark, v...)
}

func MWarnf(mark string, format string, args ...interface{}) string {
	return GetLogout().MWarnf(mark, format, args...)
}

// 上层代码行日志
func PWarn(v ...any) any {
	return GetLogout().PWarn(v...)
}

func PWarnln(v ...any) any {
	return GetLogout().PWarnln(v...)
}

func PWarnf(format string, args ...interface{}) string {
	return GetLogout().PWarnf(format, args...)
}
