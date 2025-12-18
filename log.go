package log

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/os/glog"
)

type LogLevel int

const (
	LEVEL_ALL  LogLevel = LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_DEV  LogLevel = LEVEL_ALL
	LEVEL_PROD LogLevel = LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_NONE LogLevel = 0
	LEVEL_DEBU LogLevel = 1 << iota // 16
	LEVEL_INFO                      // 32
	LEVEL_NOTI                      // 64
	LEVEL_WARN                      // 128
	LEVEL_ERRO                      // 256
	LEVEL_CRIT                      // 512
	LEVEL_PANI                      // 1024
	LEVEL_FATA                      // 2048
)

var DebugMode bool

type LogCtx string

const (
	CtxMark LogCtx = "mark"
)

var ctx = context.Background()

type Log struct {
	ctx    context.Context
	logger *glog.Logger
}

func (l *Log) Ctx(ctx context.Context) *Log {
	var mark string
	if l.ctx != nil {
		mark, _ = ctx.Value(CtxMark).(string)
	}
	if mark != "" {
		l.ctx = context.WithValue(ctx, CtxMark, mark)
	} else {
		l.ctx = ctx
	}
	return l
}

func (l *Log) Debug(args ...any) {
	l.logger.Debug(l.ctx, args...)
}

func (l *Log) Debugln(args ...any) {
	l.logger.Debug(l.ctx, args...)
}

func (l *Log) Debugf(format string, args ...any) {
	l.logger.Debugf(l.ctx, format, args...)
}

func (l *Log) Print(args ...any) {
	l.logger.Info(l.ctx, args...)
}

func (l *Log) Println(args ...any) {
	l.logger.Info(l.ctx, args...)
}

func (l *Log) Printf(format string, args ...any) {
	l.logger.Infof(l.ctx, format, args...)
}

func (l *Log) Info(args ...any) {
	l.logger.Info(l.ctx, args...)
}

func (l *Log) Infoln(args ...any) {
	l.logger.Info(l.ctx, args...)
}

func (l *Log) Infof(format string, args ...any) {
	l.logger.Infof(l.ctx, format, args...)
}

func (l *Log) Error(args ...any) error {
	l.logger.Error(l.ctx, args...)
	return errors.New(fmt.Sprint(args...))
}

func (l *Log) Errorln(args ...any) error {
	l.logger.Error(l.ctx, args...)
	return errors.New(fmt.Sprint(args...))
}

func (l *Log) Errorf(format string, args ...any) error {
	l.logger.Errorf(l.ctx, format, args...)
	return fmt.Errorf(format, args...)
}

func (l *Log) Warn(args ...any) {
	l.logger.Warning(l.ctx, args...)
}

func (l *Log) Warnln(args ...any) {
	l.logger.Warning(l.ctx, args...)
}

func (l *Log) Warnf(format string, args ...any) {
	l.logger.Warningf(l.ctx, format, args...)
}

func (l *Log) IsError(err error) bool {
	if err != nil {
		glog.Error(l.ctx, err)
	}
	return err != nil
}

func (l *Log) Logger(logger *glog.Logger) *Log {
	l.logger = logger
	return l
}

func (l *Log) File(file string) *Log {
	l.logger.SetFile(file + "-{Y-m-d}.log")
	return l
}

func (l *Log) NoStd() *Log {
	l.logger.SetStdoutPrint(false)
	return l
}

func (l *Log) SetLevel(level LogLevel) *Log {
	l.logger.SetLevel(int(level))
	return l
}

type LogOption struct {
	Debug    bool
	LogDir   string
	LogMark  string
	NoStdout bool
}

func Init(opt *LogOption) {
	if opt.Debug {
		DebugMode = true
		glog.DefaultLogger().SetLevel(glog.LEVEL_ALL)
	}
	if opt.LogMark != "" {
		glog.SetFile(opt.LogMark + "-{Y-m-d}.log")
	}
	glog.SetCtxKeys(CtxMark)
	glog.SetFlags(glog.F_FILE_SHORT | glog.F_CALLER_FN)
	glog.DefaultLogger().SetStackSkip(1)
	if opt.LogDir == "" {
		return
	}
	glog.DefaultLogger().SetStdoutPrint(!opt.NoStdout)
	glog.DefaultLogger().SetPath(opt.LogDir)
}

func SetDebug(debug bool) {
	if debug {
		DebugMode = true
		glog.DefaultLogger().SetLevel(glog.LEVEL_ALL)
	} else {
		DebugMode = false
		glog.DefaultLogger().SetLevel(glog.LEVEL_INFO)
	}
}

func SetLevel(level LogLevel) {
	glog.DefaultLogger().SetLevel(int(level))
}

func IsError(err error) bool {
	if err != nil {
		glog.Error(ctx, err)
	}
	return err != nil
}

func New(mark ...string) *Log {
	if len(mark) == 0 {
		return Ctx(ctx)
	}
	ctx := context.WithValue(ctx, CtxMark, mark[0])
	return Ctx(ctx)
}
