package log

import "fmt"

// 打印警告日志
func (l *Logger) Warn(v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Warnf(format string, args ...interface{}) string {
	if l.logLevel >= LogLevelWarn {
		l.warnf(l.mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) Warnln(v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) WarnfCall(call int, format string, args ...interface{}) string {
	if l.logLevel >= LogLevelWarn {
		l.warnf(l.mark, call, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) WarnlnCall(call int, v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(l.mark, call, v...)
	}
	return v
}

func (l *Logger) WarnCall(call int, v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(l.mark, call, v...)
	}
	return v
}

// 标记日志
func (l *Logger) MWarn(mark string, v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MWarnln(mark string, v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MWarnf(mark string, format string, args ...interface{}) string {
	if l.logLevel >= LogLevelWarn {
		l.warnf(mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

// 上层代码行日志
func (l *Logger) PWarn(v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PWarnln(v ...any) any {
	if l.logLevel >= LogLevelWarn {
		l.warn(l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PWarnf(format string, args ...interface{}) string {
	if l.logLevel >= LogLevelWarn {
		l.warnf(l.mark, l.caller+1, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) warn(mark string, call int, args ...any) {
	l.printcallLog(LogLevelWarn, mark, call+1, args...)
}

func (l *Logger) warnf(mark string, call int, format string, args ...any) {
	l.printfcallLog(LogLevelWarn, mark, call+1, format, args...)
}
