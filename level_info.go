package log

// 打印Info日志
func Info(v ...any) any {
	return GetLogout().Info(v...)
}

func Infoln(v ...any) any {
	return GetLogout().Infoln(v...)
}

func Printf(format string, args ...interface{}) string {
	return GetLogout().Printf(format, args...)
}
func Print(v ...any) any {
	return GetLogout().Print(v...)
}

func Println(v ...any) any {
	return GetLogout().Println(v...)
}

func Infof(format string, args ...interface{}) string {
	return GetLogout().Infof(format, args...)
}

func InfoCall(call int, args ...any) any {
	return GetLogout().InfoCall(call, args...)
}

func InfolnCall(call int, args ...any) any {
	return GetLogout().InfolnCall(call, args...)
}

func InfofCall(call int, format string, args ...interface{}) string {
	return GetLogout().InfofCall(call, format, args...)
}

// 标记日志
func MInfo(mark string, v ...any) any {
	return GetLogout().MInfo(mark, v...)
}

func MInfoln(mark string, v ...any) any {
	return GetLogout().MInfoln(mark, v...)
}

func MInfof(mark string, format string, args ...interface{}) string {
	return GetLogout().MInfof(mark, format, args...)
}

// 上层代码行日志
func PInfo(v ...any) any {
	return GetLogout().PInfo(v...)
}

func PInfoln(v ...any) any {
	return GetLogout().PInfoln(v...)
}

func PInfof(format string, args ...interface{}) string {
	return GetLogout().PInfof(format, args...)
}
