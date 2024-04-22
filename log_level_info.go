package log

import "fmt"

// 打印信息日志
func (l *Logger) Info(v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Infof(format string, args ...interface{}) string {
	if l.logLevel >= INFOLevel {
		l.printfcallLog(INFOPrefix, l.mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) Infoln(v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Print(v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Println(v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, l.caller, v...)
	}
	return v
}

func (l *Logger) Printf(format string, args ...interface{}) string {
	if l.logLevel >= INFOLevel {
		l.printfcallLog(INFOPrefix, l.mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) PrintlnCall(call int, args ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, call, args...)
	}
	return args
}

func (l *Logger) PrintfCall(call int, format string, args ...interface{}) string {
	if l.logLevel >= INFOLevel {
		l.printfcallLog(INFOPrefix, l.mark, call, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) PrintCall(call int, args ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, call, args...)
	}
	return args
}

func (l *Logger) InfoCall(call int, args ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, call, args...)
	}
	return args
}

func (l *Logger) InfofCall(call int, format string, args ...interface{}) string {
	if l.logLevel >= INFOLevel {
		l.printfcallLog(INFOPrefix, l.mark, call, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *Logger) InfolnCall(call int, args ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, call, args...)
	}
	return args
}

// 标记日志
func (l *Logger) MInfo(mark string, v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MInfoln(mark string, v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, mark, l.caller, v...)
	}
	return v
}

func (l *Logger) MInfof(mark string, format string, args ...interface{}) string {
	if l.logLevel >= INFOLevel {
		l.printfcallLog(INFOPrefix, mark, l.caller, format, args...)
	}
	return fmt.Sprintf(format, args...)
}

// 上层代码行日志
func (l *Logger) PInfo(v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PInfoln(v ...any) any {
	if l.logLevel >= INFOLevel {
		l.printcallLog(INFOPrefix, l.mark, l.caller+1, v...)
	}
	return v
}

func (l *Logger) PInfof(format string, args ...interface{}) string {
	if l.logLevel >= INFOLevel {
		l.printfcallLog(INFOPrefix, l.mark, l.caller+1, format, args...)
	}
	return fmt.Sprintf(format, args...)
}
