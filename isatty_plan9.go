//go:build plan9

package tint

import (
	"os"
)

// IsTerminal reports whether the given file descriptor is a terminal.
// On Plan9, we check if /dev/cons can be opened.
func IsTerminal(fd uintptr) bool {
	f, err := os.Open("/dev/cons")
	if err != nil {
		return false
	}
	f.Close()
	return true
}

// isStdoutTerminal reports whether os.Stdout is a terminal.
func isStdoutTerminal() bool {
	return IsTerminal(os.Stdout.Fd())
}

// isStderrTerminal reports whether os.Stderr is a terminal.
func isStderrTerminal() bool {
	return IsTerminal(os.Stderr.Fd())
}
