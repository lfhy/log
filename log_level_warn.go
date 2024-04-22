package log

import "fmt"

// 打印警告日志
func (l *Logger) Warn(v ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Warnf(format string, args ...interface{}) string {
	if l.logLevel >= WARNLevel {
		l.printfcallLog(WARNPrefix, l.mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) Warnln(v ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, l.mark, 2, v...)
	}
	return v
}

func (l *Logger) WarnfCall(call int, format string, args ...interface{}) string {
	if l.logLevel >= WARNLevel {
		l.printfcallLog(WARNPrefix, l.mark, call, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) WarnlnCall(call int, args ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, l.mark, call, args...)
	}
	return args
}

func (l *Logger) WarnCall(call int, args ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, l.mark, call, args...)
	}
	return args
}

// 标记日志
func (l *Logger) MWarn(mark string, v ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MWarnln(mark string, v ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MWarnf(mark string, format string, args ...interface{}) string {
	if l.logLevel >= WARNLevel {
		l.printfcallLog(WARNPrefix, mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

// 上层代码行日志
func (l *Logger) PWarn(v ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PWarnln(v ...any) any {
	if l.logLevel >= WARNLevel {
		l.printcallLog(WARNPrefix, l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PWarnf(format string, args ...interface{}) string {
	if l.logLevel >= WARNLevel {
		l.printfcallLog(WARNPrefix, l.mark, l.caller+1, format, args...)
	}
	return fmt.Sprintf(format, args...)
}
