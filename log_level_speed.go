package log

import "time"

// 计时开始时间
func (l *Logger) StartTime() time.Time {
	startTimeUnix := time.Now()
	l.Speed("执行时间:", startTimeUnix)
	return startTimeUnix
}

// 结束时间时间
func (l *Logger) EndTime(startTime time.Time) time.Duration {
	endTimeUnix := time.Since(startTime)
	if endTimeUnix < 100*time.Millisecond {
		l.Speed("事件耗时:\033[01;32m", endTimeUnix, "\033[0m")
	} else if endTimeUnix > 1*time.Second {
		l.Speed("事件耗时:\033[01;31m", endTimeUnix, "\033[0m")
	} else {
		l.Speed("事件耗时:\033[01;37m", endTimeUnix, "\033[0m")
	}

	return endTimeUnix
}

func (l *Logger) EndTimeWithText(subText string, startTime time.Time) {
	endTimeUnix := time.Since(startTime)
	if endTimeUnix < 100*time.Millisecond {
		l.Speed(subText, "事件耗时:\033[01;32m", endTimeUnix, "\033[0m")
	} else if endTimeUnix > 1*time.Second {
		l.Speed(subText, "事件耗时:\033[01;31m", endTimeUnix, "\033[0m")
	} else {
		l.Speed(subText, "事件耗时:\033[01;37m", endTimeUnix, "\033[0m")
	}
}

// Debug日志
func (l *Logger) Speed(v ...any) {
	if l.logLevel >= SPEEDLevel {
		l.printcallLog(SPEEDPrefix, l.mark, l.caller+1, v...)
	}
}
