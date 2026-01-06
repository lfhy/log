package log

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/os/glog"
)

// Ctx 创建并返回一个带有上下文的 Log 实例。
func Ctx(ctx context.Context) *Log {
	return &Log{ctx: ctx, logger: glog.DefaultLogger().Clone()}
}

// Debug 使用默认上下文记录调试信息。
func Debug(args ...any) {
	glog.Debug(ctx, args...)
}

// Debugln 使用默认上下文记录调试信息。
func Debugln(args ...any) {
	glog.Debug(ctx, args...)
}

// Debugf 使用默认上下文记录调试信息。
func Debugf(format string, args ...any) {
	glog.Debugf(ctx, format, args...)
}

// Print 使用默认上下文记录信息。
func Print(args ...any) {
	glog.Info(ctx, args...)
}

// Println 使用默认上下文记录一行信息。
func Println(args ...any) {
	glog.Info(ctx, args...)
}

// Printf 使用默认上下文格式化并记录信息。
func Printf(format string, args ...any) {
	glog.Infof(ctx, format, args...)
}

// Info 使用默认上下文记录信息。
func Info(args ...any) {
	glog.Info(ctx, args...)
}

// Infoln 使用默认上下文记录一行信息。
func Infoln(args ...any) {
	glog.Info(ctx, args...)
}

// Infof 使用默认上下文格式化并记录信息。
func Infof(format string, args ...any) {
	glog.Infof(ctx, format, args...)
}

// Notice 使用默认上下文记录通知信息。
func Notice(args ...any) {
	glog.Notice(ctx, args...)
}

// Noticeln 使用默认上下文记录一行通知信息。
func Noticeln(args ...any) {
	glog.Notice(ctx, args...)
}

// Noticef 使用默认上下文格式化并记录通知信息。
func Noticef(format string, args ...any) {
	glog.Noticef(ctx, format, args...)
}

// Fatal 使用默认上下文记录致命错误信息，并终止程序。
func Fatal(args ...any) {
	glog.Fatal(ctx, args...)
}

// Fatalln 使用默认上下文记录一行致命错误信息，并终止程序。
func Fatalln(args ...any) {
	glog.Fatal(ctx, args...)
}

// Fatalf 使用默认上下文格式化并记录致命错误信息，并终止程序。
func Fatalf(format string, args ...any) {
	glog.Fatalf(ctx, format, args...)
}

// Panic 使用默认上下文记录恐慌信息，并触发panic。
func Panic(args ...any) {
	glog.Panic(ctx, args...)
}

// Panicln 使用默认上下文记录一行恐慌信息，并触发panic。
func Panicln(args ...any) {
	glog.Panic(ctx, args...)
}

// Panicf 使用默认上下文格式化并记录恐慌信息，并触发panic。
func Panicf(format string, args ...any) {
	glog.Panicf(ctx, format, args...)
}

// Critical 使用默认上下文记录严重错误信息。
func Critical(args ...any) {
	glog.Critical(ctx, args...)
}

// Criticalln 使用默认上下文记录一行严重错误信息。
func Criticalln(args ...any) {
	glog.Critical(ctx, args...)
}

// Criticalf 使用默认上下文格式化并记录严重错误信息。
func Criticalf(format string, args ...any) {
	glog.Criticalf(ctx, format, args...)
}

// Error 使用默认上下文记录错误信息，并返回错误对象。
func Error(args ...any) error {
	glog.Error(ctx, args...)
	return errors.New(fmt.Sprint(args...))
}

// Errorln 使用默认上下文记录一行错误信息，并返回错误对象。
func Errorln(args ...any) error {
	glog.Error(ctx, args...)
	return errors.New(fmt.Sprint(args...))
}

// Errorf 使用默认上下文格式化并记录错误信息，并返回错误对象。
func Errorf(format string, args ...any) error {
	glog.Errorf(ctx, format, args...)
	return fmt.Errorf(format, args...)
}

// Warnln 使用默认上下文记录警告信息。
func Warnln(args ...any) {
	glog.Warning(ctx, args...)
}

// Warnf 使用默认上下文格式化并记录警告信息。
func Warnf(format string, args ...any) {
	glog.Warningf(ctx, format, args...)
}

// Warn 使用默认上下文记录警告信息。
func Warn(args ...any) {
	glog.Warning(ctx, args...)
}
