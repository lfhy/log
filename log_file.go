package log

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 获取时间
func getNowTimeStr() string {
	return time.Now().Format("2006-01-02")
}

// 获取日志名称
func (l *Logger) GetLogFileName() string {
	return l.logfileName
}

// 设置日志名称 不带.log后缀
func (l *Logger) SetLogFileName(name string) {
	l.logfileName = name
}

// 获取文件后缀
func (l *Logger) GetLogFileSuffix() string {
	if l.logfileSuffix == "" {
		l.logfileSuffix = ".log"
	}
	return l.logfileSuffix
}

// 设置日志文件后缀
func (l *Logger) SetLogFileSuffix(suffix string) {
	if strings.HasPrefix(suffix, ".") {
		l.logfileSuffix = suffix
		return
	}
	l.logfileSuffix = "." + suffix
}

// 获取日志文件完整路径
func (l *Logger) GetLogFileFullPath() string {
	dir := l.GetLogFileDir()
	name := l.GetLogFileName()
	if name == "" {
		return filepath.Join(dir, getNowTimeStr()+l.GetLogFileSuffix())
	}
	return filepath.Join(dir, name+l.GetLogFileSuffix())
}

func (l *Logger) GetLogFileDir() string {
	if l.logfileDir == "" {
		return "./"
	}
	return l.logfileDir
}

func (l *Logger) SetLogFileDir(dir string) {
	l.logfileDir = dir
}

// 获取日志文件句柄 不对外暴露
// func (l *Logger) getLogFileHandle() *os.File {
// 	if l.logfile == nil {
// 		return os.Stdout
// 	}
// 	return l.logfile
// }

func (l *Logger) setLogFileHandle(file *os.File) {
	l.logfile = file
}

func (l *Logger) SetNoReload(noReload bool) {
	l.noReload = noReload
}

func (l *Logger) openLogfile() {
	os.MkdirAll(l.GetLogFileDir(), 0755)
	fullPath := l.GetLogFileFullPath()
	LogFile, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开失败使用标准输出
		l.setLogger(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds))
	} else {
		l.setLogFileHandle(LogFile)
		l.setLogger(log.New(LogFile, "", log.LstdFlags|log.Lmicroseconds))
		if l.redirectPanic {
			RedirectError(LogFile)
		}
	}
	// 配置一个定时器在0点重置日志
	go func() {
		var next time.Time
		// 配置定时时间
		now := time.Now() //获取当前时间，放到now里面，要给next用
		if now.Hour() >= 0 && now.Minute() >= 0 {
			next = now.Add(time.Hour * 24) //通过now偏移24小时
		} else {
			next = now
		}
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取执行命令的日期
		// Info("任务将在:", next, "执行")
		t := time.NewTimer(next.Sub(now)) //设置一个定时器
		<-t.C
		//以下为定时执行的操作
		l.reloadLogFile()
	}()
}

// 重载日志
// 每日0点重载
func (l *Logger) reloadLogFile() {
	l.logoutLock.Lock()
	defer l.logoutLock.Unlock()
	if l.noReload {
		return
	}
	// 如果是Std的输出则不处理
	if l.isStdout || l.logfile == nil {
		return
	}
	fi, err := l.logfile.Stat()
	// 如果文件出现问题则不进行移动操作直接进行打开新日志
	if err != nil {
		l.openLogfile()
		return
	}
	l.logfile.Close()
	// 获取旧路径
	LogFileDir := l.GetLogFileDir()
	LogFileName := l.GetLogFileName()
	if LogFileName == "" {
		// 没有名称则为昨天的时间
		LogFileName = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}
	BeforeLogFile := filepath.Join(LogFileDir, LogFileName+l.GetLogFileSuffix())
	nowDate := time.Now().Format("2006-01-02")
	mtime := fi.ModTime().Format("2006-01-02")
	// 判断当前时间是否和文件时间一致 不一致则进行移动
	if mtime != nowDate {
		oldfile := filepath.Join(LogFileDir, LogFileName+"-"+mtime+l.GetLogFileSuffix())
		// 如果旧文件存在则进行追加
		if _, err := os.Stat(oldfile); err == nil {
			addFile(oldfile, BeforeLogFile, oldfile)
		} else {
			// 不存在则进行重命名
			os.Rename(BeforeLogFile, oldfile)
		}
		os.RemoveAll(BeforeLogFile)
	}
	l.openLogfile()
}

// 合并文件
// 参数1:合并文件1
// 参数2:合并文件2
// 参数3:输出路径
// 合并文件会把两个文件进行合并输出成单独的一个文件
func addFile(src1 string, src2 string, dst string) error {
	// 1.打开文件1
	input1, err := os.ReadFile(src1)
	if err != nil {
		input1 = []byte{}
	}

	// 2.打开文件2
	input2, err := os.ReadFile(src2)
	if err != nil {
		input2 = []byte{}
	}

	var buffer bytes.Buffer
	buffer.Write(input1)
	// 添加换行
	buffer.Write([]byte("\n"))
	buffer.Write(input2)

	// 3.写入文件
	return os.WriteFile(dst, buffer.Bytes(), 0644)
}
