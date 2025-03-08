package log

import (
	"errors"
	"fmt"
)

// 打印错误日志
func (l *Logger) Error(v ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(l.mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) Errorln(v ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(l.mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) Errorf(format string, args ...interface{}) error {
	if l.logLevel >= LogLevelError {
		l.errorf(l.mark, l.caller, format, args...)
	}
	return fmt.Errorf(format, args...)
}

func (l *Logger) ErrorfCall(call int, format string, args ...interface{}) error {
	if l.logLevel >= LogLevelError {
		l.errorf(l.mark, call, format, args...)
	}
	return fmt.Errorf(format, args...)
}

func (l *Logger) ErrorlnCall(call int, args ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(l.mark, call, args...)
	}
	return errors.New(fmt.Sprint(args...))
}

func (l *Logger) ErrorCall(call int, args ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(l.mark, call, args...)
	}
	return errors.New(fmt.Sprint(args...))
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
	if l.logLevel >= LogLevelError {
		l.error(mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) MErrorln(mark string, v ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(mark, l.caller, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) MErrorf(mark string, format string, args ...interface{}) error {
	if l.logLevel >= LogLevelError {
		l.errorf(mark, l.caller, format, args...)
	}
	return fmt.Errorf(format, args...)
}

// 上层代码行日志
func (l *Logger) PError(v ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(l.mark, l.caller+1, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) PErrorln(v ...any) error {
	if l.logLevel >= LogLevelError {
		l.error(l.mark, l.caller+1, v...)
	}
	return fmt.Errorf("%v", v...)
}

func (l *Logger) PErrorf(format string, args ...interface{}) error {
	if l.logLevel >= LogLevelError {
		l.errorf(l.mark, l.caller+1, format, args...)
	}
	return fmt.Errorf(format, args...)
}

func (l *Logger) error(mark string, call int, args ...any) {
	l.printcallLog(LogLevelError, mark, call+1, args...)
}

func (l *Logger) errorf(mark string, call int, format string, args ...any) {
	l.printfcallLog(LogLevelError, mark, call+1, format, args...)
}
