//go:build !windows

package tint

import (
	"io"
	"os"
)

// NewColorable returns w as-is on Unix/Plan9 since terminals natively
// support ANSI escape sequences.
func NewColorable(w io.Writer) io.Writer {
	return w
}

// NewColorableStdout returns a writer for os.Stdout that supports colored output.
func NewColorableStdout() io.Writer {
	return os.Stdout
}

// NewColorableStderr returns a writer for os.Stderr that supports colored output.
func NewColorableStderr() io.Writer {
	return os.Stderr
}
