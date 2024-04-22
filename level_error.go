package log

// 打印Error日志
func Error(v ...any) error {
	return GetLogout().Error(v...)
}

func Errorln(v ...any) error {
	return GetLogout().Errorln(v...)
}

func Errorf(format string, args ...interface{}) error {
	return GetLogout().Errorf(format, args...)
}

func ErrorCall(call int, args ...any) error {
	return GetLogout().ErrorCall(call, args...)
}

func ErrorlnCall(call int, args ...any) error {
	return GetLogout().ErrorlnCall(call, args...)
}

func ErrorfCall(call int, format string, args ...interface{}) error {
	return GetLogout().ErrorfCall(call, format, args...)
}

// 标记日志
func MError(mark string, v ...any) error {
	return GetLogout().MError(mark, v...)
}

func MErrorln(mark string, v ...any) error {
	return GetLogout().MErrorln(mark, v...)
}

func MErrorf(mark string, format string, args ...interface{}) error {
	return GetLogout().MErrorf(mark, format, args...)
}

// 上层代码行日志
func PError(v ...any) error {
	return GetLogout().PError(v...)
}

func PErrorln(v ...any) error {
	return GetLogout().PErrorln(v...)
}

func PErrorf(format string, args ...interface{}) error {
	return GetLogout().PErrorf(format, args...)
}

func IsError(err error) bool {
	return GetLogout().IsError(err)
}

func NoError(err error) bool {
	return GetLogout().NoError(err)
}
