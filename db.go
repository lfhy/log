package log

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	"gorm.io/gorm/logger"
)

type ORMLog struct {
	ctx    context.Context
	Level  logger.LogLevel
	logger *glog.Logger
}

type ORMLogOption struct {
	LogOption
	Level logger.LogLevel
}

func NewORMLog(opt ...*ORMLogOption) *ORMLog {
	if len(opt) == 0 {
		opt = append(opt, &ORMLogOption{})
	}
	l := glog.DefaultLogger().Clone()
	if opt[0].Debug {
		l.SetLevel(glog.LEVEL_ALL)
		opt[0].Level = logger.Info
	}
	if opt[0].LogMark != "" {
		l.SetFile(opt[0].LogMark + "-{Y-m-d}.log")
	} else {
		l.SetFile("db-{Y-m-d}.log")
	}

	if opt[0].LogDir != "" {
		l.SetPath(opt[0].LogDir)
	}
	l.SetStdoutPrint(!opt[0].NoStdout)
	return &ORMLog{
		Level:  opt[0].Level,
		logger: l,
	}
}

func (m *ORMLog) LogMode(l logger.LogLevel) logger.Interface {
	m.Level = l
	return m
}

func (m *ORMLog) Ctx() context.Context {
	if m.ctx == nil {
		m.ctx = context.WithValue(context.Background(), CtxMark, "ORM")
	}
	return m.ctx
}

func (m *ORMLog) Info(ctx context.Context, f string, a ...interface{}) {
	if m.Level < logger.Info {
		return
	}
	m.logger.Infof(m.Ctx(), f, a...)
}

func (m *ORMLog) Warn(ctx context.Context, f string, a ...interface{}) {
	if m.Level < logger.Warn {
		return
	}
	m.logger.Warningf(m.Ctx(), f, a...)
}

func (m *ORMLog) Error(ctx context.Context, f string, a ...interface{}) {
	if m.Level < logger.Error {
		return
	}
	m.logger.Errorf(m.Ctx(), f, a...)
}

func (m *ORMLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// glog.Debug(m.Ctx(), "开始时间:", begin)
	sql, rows := fc()
	m.logger.Debug(m.Ctx(), "SQL:", sql, "行数:", rows)
	if err != nil {
		m.logger.Error(m.Ctx(), "错误:", err)
	}
}
