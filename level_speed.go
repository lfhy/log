package log

import "time"

// 计时开始时间
func StartTime() time.Time {
	return GetLogout().StartTime()
}

// 结束时间时间
func EndTime(startTime time.Time) time.Duration {
	return GetLogout().EndTime(startTime)
}

func EndTimeWithText(subText string, startTime time.Time) {
	GetLogout().EndTimeWithText(subText, startTime)
}

// Debug日志
func Speed(v ...any) {
	GetLogout().Speed(v...)
}
