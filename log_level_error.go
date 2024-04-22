package log

import "fmt"

// 打印错误日志
func (l *Logger) Error(v ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, l.mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) Errorln(v ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, l.mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) Errorf(format string, args ...interface{}) error {
	if l.logLevel >= ERRORLevel {
		l.printfcallLog(ERRORPrefix, l.mark, l.caller, format, args...)
	}
	return fmt.Errorf(format, args...)
}

func (l *Logger) ErrorfCall(call int, format string, args ...interface{}) error {
	if l.logLevel >= ERRORLevel {
		l.printfcallLog(ERRORPrefix, l.mark, call, format, args...)
	}
	return fmt.Errorf(format, args...)
}

func (l *Logger) ErrorlnCall(call int, args ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, l.mark, call, args...)
	}
	return fmt.Errorf(fmt.Sprint(args...))
}

func (l *Logger) ErrorCall(call int, args ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, l.mark, call, args...)
	}
	return fmt.Errorf(fmt.Sprint(args...))
}

func (l *Logger) IsError(err error) bool {
	if err != nil {
		l.ErrorlnCall(l.caller+1, "出现错误:", err)
	}
	return err != nil
}

func (l *Logger) NoError(err error) bool {
	if err != nil {
		l.ErrorlnCall(l.caller+1, "出现错误:", err)
	}
	return err == nil
}

// 标记日志
func (l *Logger) MError(mark string, v ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) MErrorln(mark string, v ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) MErrorf(mark string, format string, args ...interface{}) error {
	if l.logLevel >= ERRORLevel {
		l.printfcallLog(ERRORPrefix, mark, l.caller, format, args...)
	}
	return fmt.Errorf(format, args...)
}

// 上层代码行日志
func (l *Logger) PError(v ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, l.mark, l.caller+1, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) PErrorln(v ...any) error {
	if l.logLevel >= ERRORLevel {
		l.printcallLog(ERRORPrefix, l.mark, l.caller+1, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) PErrorf(format string, args ...interface{}) error {
	if l.logLevel >= ERRORLevel {
		l.printfcallLog(ERRORPrefix, l.mark, l.caller+1, format, args...)
	}
	return fmt.Errorf(format, args...)
}
