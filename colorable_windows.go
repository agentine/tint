//go:build windows

package tint

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

var (
	procSetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleMode          = kernel32.NewProc("SetConsoleMode")
)

const (
	enableVirtualTerminalProcessing = 0x0004
)

// consoleScreenBufferInfo contains information about a console screen buffer.
type consoleScreenBufferInfo struct {
	size       [2]int16
	cursorPos  [2]int16
	attributes uint16
	window     [4]int16
	maxSize    [2]int16
}

// colorableWriter translates ANSI escape sequences into Windows console API
// calls. If the console supports virtual terminal sequences (Windows 10+),
// ANSI sequences are passed through directly.
type colorableWriter struct {
	w         io.Writer
	handle    syscall.Handle
	origAttrs uint16
	vtEnabled bool
}

// NewColorable returns a writer that supports colored output on Windows.
// If the underlying writer is not a console handle or the console supports
// ANSI VT sequences natively, the writer is returned as-is.
func NewColorable(w io.Writer) io.Writer {
	if f, ok := w.(*os.File); ok {
		handle := syscall.Handle(f.Fd())
		var info consoleScreenBufferInfo
		r, _, _ := procGetConsoleScreenBufferInfo.Call(
			uintptr(handle),
			uintptr(unsafe.Pointer(&info)),
		)
		if r == 0 {
			// Not a console — return as-is (e.g., pipe, file).
			return w
		}

		// Try to enable VT processing (Windows 10 1607+).
		var mode uint32
		procGetConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
		r2, _, _ := procSetConsoleMode.Call(uintptr(handle), uintptr(mode|enableVirtualTerminalProcessing))
		if r2 != 0 {
			// VT sequences supported natively — pass through.
			return w
		}

		return &colorableWriter{
			w:         w,
			handle:    handle,
			origAttrs: info.attributes,
		}
	}
	return w
}

// NewColorableStdout returns a writer for os.Stdout that supports colored output.
func NewColorableStdout() io.Writer {
	return NewColorable(os.Stdout)
}

// NewColorableStderr returns a writer for os.Stderr that supports colored output.
func NewColorableStderr() io.Writer {
	return NewColorable(os.Stderr)
}

// Write translates ANSI SGR sequences to Windows console attribute calls.
func (cw *colorableWriter) Write(data []byte) (int, error) {
	total := len(data)
	for len(data) > 0 {
		// Find next escape sequence.
		idx := bytes.IndexByte(data, '\033')
		if idx < 0 {
			// No more escape sequences — write remaining text.
			_, err := cw.w.Write(data)
			if err != nil {
				return total - len(data), err
			}
			break
		}

		// Write text before escape.
		if idx > 0 {
			_, err := cw.w.Write(data[:idx])
			if err != nil {
				return total - len(data), err
			}
		}
		data = data[idx:]

		// Parse CSI sequence: ESC [ params m
		if len(data) < 2 || data[1] != '[' {
			_, err := cw.w.Write(data[:1])
			if err != nil {
				return total - len(data), err
			}
			data = data[1:]
			continue
		}

		// Find the end of the CSI sequence (letter terminator).
		end := 2
		for end < len(data) && data[end] != 'm' {
			if data[end] >= 'A' && data[end] <= 'Z' && data[end] != 'M' {
				break
			}
			if data[end] >= 'a' && data[end] <= 'z' {
				break
			}
			end++
		}

		if end >= len(data) || data[end] != 'm' {
			// Not an SGR sequence — write the escape and move on.
			_, err := cw.w.Write(data[:end])
			if err != nil {
				return total - len(data), err
			}
			data = data[end:]
			continue
		}

		// Parse SGR parameters.
		params := string(data[2:end])
		cw.applySGR(params)
		data = data[end+1:]
	}
	return total, nil
}

// applySGR applies SGR (Select Graphic Rendition) parameters to the console.
func (cw *colorableWriter) applySGR(params string) {
	if params == "" || params == "0" {
		// Reset to original attributes.
		procSetConsoleTextAttribute.Call(uintptr(cw.handle), uintptr(cw.origAttrs))
		return
	}

	attrs := cw.origAttrs
	parts := strings.Split(params, ";")
	for i := 0; i < len(parts); i++ {
		code, err := strconv.Atoi(parts[i])
		if err != nil {
			continue
		}

		switch {
		case code == 0:
			attrs = cw.origAttrs
		case code == 1: // Bold — use intensity bit.
			attrs |= 0x0008 // FOREGROUND_INTENSITY
		case code >= 30 && code <= 37:
			attrs = (attrs & 0xFFF0) | ansiToWinFg(code-30)
		case code >= 40 && code <= 47:
			attrs = (attrs & 0xFF0F) | ansiToWinBg(code-40)
		case code >= 90 && code <= 97:
			attrs = (attrs & 0xFFF0) | ansiToWinFg(code-90) | 0x0008
		case code >= 100 && code <= 107:
			attrs = (attrs & 0xFF0F) | ansiToWinBg(code-100) | 0x0080
		}
	}
	procSetConsoleTextAttribute.Call(uintptr(cw.handle), uintptr(attrs))
}

// ansiToWinFg maps ANSI color index (0-7) to Windows foreground attributes.
func ansiToWinFg(idx int) uint16 {
	// ANSI: 0=black 1=red 2=green 3=yellow 4=blue 5=magenta 6=cyan 7=white
	// Windows bits: B=1 G=2 R=4
	table := [8]uint16{0, 4, 2, 6, 1, 5, 3, 7}
	if idx >= 0 && idx < 8 {
		return table[idx]
	}
	return 7
}

// ansiToWinBg maps ANSI color index (0-7) to Windows background attributes.
func ansiToWinBg(idx int) uint16 {
	return ansiToWinFg(idx) << 4
}
