//go:build windows

package tint

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32                = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode      = kernel32.NewProc("GetConsoleMode")
)

// IsTerminal reports whether the given file descriptor is a terminal (console).
func IsTerminal(fd uintptr) bool {
	var mode uint32
	r, _, _ := procGetConsoleMode.Call(fd, uintptr(unsafe.Pointer(&mode)))
	return r != 0
}

// isStdoutTerminal reports whether os.Stdout is a terminal.
func isStdoutTerminal() bool {
	return IsTerminal(os.Stdout.Fd())
}

// isStderrTerminal reports whether os.Stderr is a terminal.
func isStderrTerminal() bool {
	return IsTerminal(os.Stderr.Fd())
}
