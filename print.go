package log

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/os/glog"
)

func Ctx(ctx context.Context) *Log {
	return &Log{ctx: ctx, logger: glog.DefaultLogger().Clone()}
}

func Debug(args ...any) {
	glog.Debug(ctx, args...)
}

func Debugln(args ...any) {
	glog.Debug(ctx, args...)
}

func Debugf(format string, args ...any) {
	glog.Debugf(ctx, format, args...)
}

func Print(args ...any) {
	glog.Info(ctx, args...)
}

func Println(args ...any) {
	glog.Info(ctx, args...)
}

func Printf(format string, args ...any) {
	glog.Infof(ctx, format, args...)
}

func Info(args ...any) {
	glog.Info(ctx, args...)
}

func Infoln(args ...any) {
	glog.Info(ctx, args...)
}

func Infof(format string, args ...any) {
	glog.Infof(ctx, format, args...)
}

func Error(args ...any) error {
	glog.Error(ctx, args...)
	return errors.New(fmt.Sprint(args...))
}

func Errorln(args ...any) error {
	glog.Error(ctx, args...)
	return errors.New(fmt.Sprint(args...))
}

func Errorf(format string, args ...any) error {
	glog.Errorf(ctx, format, args...)
	return fmt.Errorf(format, args...)
}

func Warnln(args ...any) {
	glog.Warning(ctx, args...)
}

func Warnf(format string, args ...any) {
	glog.Warningf(ctx, format, args...)
}

func Warn(args ...any) {
	glog.Warning(ctx, args...)
}
