package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func (l *Logger) printcallLog(logLevel LogLevel, mark string, call int, v ...any) {
	l.printfcallLog(logLevel, mark, call+1, "%s", fmt.Sprintln(v...))
}

func (l *Logger) printfcallLog(logLevel LogLevel, mark string, call int, format string, args ...any) {
	if logLevel == 0 {
		return
	}
	callpc, callfile, callline, _ := runtime.Caller(call)
	callfileName := strings.Split(callfile, "/")
	var infoString string
	var logPrefix string
	if l.noColor {
		logPrefix = LogLevelNoColorPrefixMap[logLevel]
	} else {
		logPrefix = LogLevelColorPrefixMap[logLevel]
	}

	infoString = logPrefix

	if !l.disableMark {
		infoString += "[" + mark + "] "
	}
	if !l.noCodeLine {
		if l.short {
			infoString += "[" + callfileName[len(callfileName)-1] + ":" + strconv.Itoa(callline) + "]: "
		} else {
			infoString += "[" + callfileName[len(callfileName)-1] + ":" + strconv.Itoa(callline) + "][" + runtime.FuncForPC(callpc).Name() + "]" + ":"
		}
	}

	// 推流数据库
	if l.push {
		go func() {
			*l.GetLogChannl() <- LogInfo{CreateTime: time.Now().Unix(), Level: logLevel, CodeFile: callfile, CodeLine: callline, Content: fmt.Sprintf(format, args...)}
		}()
	}
	if !l.isStdout && l.useStdout {
		if l.short {
			l.GetStdLogger().Printf("%s", logPrefix+fmt.Sprintf(format, args...))
		} else {
			l.GetStdLogger().Printf(infoString+format, args...)
		}
	}
	l.GetLogger().Printf(infoString+format, args...)
}

// 设置打印代码层
func (l *Logger) SetCodeCaller(caller int) {
	l.caller = caller
}

// 设置打印代码层
func SetCodeCaller(caller int) {
	GetLogout().SetCodeCaller(caller)
}

func WithCodeCaller(caller int) LoggerOption {
	return func(l *Logger) {
		l.SetCodeCaller(caller)
	}
}

// 设置不打印代码行
func (l *Logger) SetNoPrintCodeLine(noprint bool) {
	l.noCodeLine = noprint
}

// 设置不打印代码行
func SetNoPrintCodeLine(noprint bool) {
	GetLogout().SetNoPrintCodeLine(noprint)
}

// 不打印代码行
func WithNoPrintCodeLine() LoggerOption {
	return func(l *Logger) {
		l.SetNoPrintCodeLine(true)
	}
}
