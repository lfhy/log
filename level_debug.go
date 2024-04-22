package log

// 打印Debug日志
func Debug(v ...any) any {
	return GetLogout().Debug(v...)
}

func Debugln(v ...any) any {
	return GetLogout().Debugln(v...)
}

func Debugf(format string, args ...interface{}) string {
	return GetLogout().Debugf(format, args...)
}

func DebugCall(call int, args ...any) any {
	return GetLogout().DebugCall(call, args...)
}

func DebuglnCall(call int, args ...any) any {
	return GetLogout().DebuglnCall(call, args...)
}

func DebugfCall(call int, format string, args ...interface{}) string {
	return GetLogout().DebugfCall(call, format, args...)
}

// 标记日志
func MDebug(mark string, v ...any) any {
	return GetLogout().MDebug(mark, v...)
}

func MDebugln(mark string, v ...any) any {
	return GetLogout().MDebugln(mark, v...)
}

func MDebugf(mark string, format string, args ...interface{}) string {
	return GetLogout().MDebugf(mark, format, args...)
}

// 上层代码行日志
func PDebug(v ...any) any {
	return GetLogout().PDebug(v...)
}

func PDebugln(v ...any) any {
	return GetLogout().PDebugln(v...)
}

func PDebugf(format string, args ...interface{}) string {
	return GetLogout().PDebugf(format, args...)
}
