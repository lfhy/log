package log

import "fmt"

// 打印调试日志
func (l *Logger) Debug(v ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Debugln(v ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(l.mark, l.caller, v...)
	}
	return v
}

// Debug日志
func (l *Logger) Debugf(format string, args ...interface{}) string {
	if l.logLevel >= LogLevelDebug {
		l.debugf(l.mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) DebugCall(call int, args ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(l.mark, call, args...)
	}
	return args
}

func (l *Logger) DebuglnCall(call int, args ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(l.mark, call, args...)
	}
	return args
}

func (l *Logger) DebugfCall(call int, format string, args ...interface{}) string {
	if l.logLevel >= LogLevelDebug {
		l.debugf(l.mark, call, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

// 标记日志
func (l *Logger) MDebug(mark string, v ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MDebugln(mark string, v ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MDebugf(mark string, format string, args ...interface{}) string {
	if l.logLevel >= LogLevelDebug {
		l.debugf(mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

// 上层代码行日志
func (l *Logger) PDebug(v ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PDebugln(v ...any) any {
	if l.logLevel >= LogLevelDebug {
		l.debug(l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PDebugf(format string, args ...interface{}) string {
	if l.logLevel >= LogLevelDebug {
		l.debugf(l.mark, l.caller+1, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) debug(mark string, call int, args ...any) {
	l.printcallLog(LogLevelDebug, mark, call+1, args...)
}

func (l *Logger) debugf(mark string, call int, format string, args ...any) {
	l.printfcallLog(LogLevelDebug, mark, call+1, format, args...)
}
