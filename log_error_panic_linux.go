//go:build linux
// +build linux

package log

import (
	"os"
	"syscall"
)

// 重定向错误至日志中
func RedirectError(out *os.File) {
	syscall.Dup3(int(out.Fd()), int(os.Stderr.Fd()), 0)
}
