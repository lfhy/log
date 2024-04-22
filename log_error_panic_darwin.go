//go:build darwin
// +build darwin

package log

import (
	"os"
	"syscall"
)

// 重定向错误至日志中
func RedirectError(out *os.File) {
	syscall.Dup2(int(out.Fd()), int(os.Stderr.Fd()))
}
