package log

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 获取日志输出器
func (l *Logger) GetLogger() *log.Logger {
	l.logoutLock.RLock()
	defer l.logoutLock.RUnlock()
	if l.Logout == nil {
		l.logoutLock.RUnlock()
		l.isStdout = true
		l.SetLogger(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
		l.logoutLock.RLock()
	}
	return l.Logout
}

func (l *Logger) SetLogger(logger *log.Logger) {
	l.logoutLock.Lock()
	defer l.logoutLock.Unlock()
	l.setLogger(logger)
}

func (l *Logger) setLogger(logger *log.Logger) {
	if file, ok := logger.Writer().(*os.File); ok {
		l.setLogFileHandle(file)
	}
	l.Logout = logger
}

func SetLogger(logger *log.Logger) {
	GetLogout().setLogger(logger)
}

// 指定日志输出器
func WithLogger(logger *log.Logger) LoggerOption {
	return func(l *Logger) {
		l.SetLogger(logger)
	}
}

// 指定日志文件
func WithLogFile(logfile string) LoggerOption {
	return func(l *Logger) {
		dir := filepath.Dir(logfile)
		l.SetLogFileDir(dir)
		l.SetLogFileName(strings.TrimSuffix(filepath.Base(logfile), ".log"))
		l.openLogfile()
	}
}

func SetLogFile(logfile string) {
	l := GetLogout()
	dir := filepath.Dir(logfile)
	l.SetLogFileDir(dir)
	filename := filepath.Base(logfile)
	ext := filepath.Ext(filename)
	l.SetLogFileName(strings.TrimSuffix(filename, ext))
	l.SetLogFileSuffix(ext)
	l.openLogfile()
}

// 指定日志文件夹
func WithLogDir(logdir string) LoggerOption {
	return func(l *Logger) {
		// 创建父目录
		l.SetLogFileDir(logdir)
		l.openLogfile()
	}
}

func WithSetNoReload(noReload bool) LoggerOption {
	return func(l *Logger) {
		// 创建父目录
		l.SetNoReload(noReload)
	}
}

func SetLogDir(logdir string) {
	l := GetLogout()
	// 创建父目录
	l.SetLogFileDir(logdir)
	l.openLogfile()
}

func SetNoReload(noReload bool) {
	l := GetLogout()
	l.SetNoReload(noReload)
}
