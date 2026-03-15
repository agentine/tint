package tint

import "io"

// ColorableWriter wraps an io.Writer with ANSI color support.
// On Unix, this is a simple pass-through. On Windows, ANSI escape
// sequences are translated to console API calls if the terminal
// does not support VT sequences natively.
type ColorableWriter interface {
	io.Writer
}
