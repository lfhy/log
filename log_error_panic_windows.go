//go:build windows
// +build windows

package log

import (
	"os"
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func setStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

// 重定向错误至日志中
func RedirectError(LogFile *os.File) {
	setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(LogFile.Fd()))
	os.Stderr = LogFile
}
