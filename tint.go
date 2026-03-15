package tint

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Color defines a custom color object which is defined by ANSI SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

// New creates a new Color with the given attributes.
func New(attrs ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0, len(attrs))}
	c.Add(attrs...)
	return c
}

// Add adds the given attributes to the color.
func (c *Color) Add(attrs ...Attribute) *Color {
	c.params = append(c.params, attrs...)
	return c
}

// Equals reports whether c and other have the same attributes.
func (c *Color) Equals(other *Color) bool {
	if len(c.params) != len(other.params) {
		return false
	}
	for i, a := range c.params {
		if a != other.params[i] {
			return false
		}
	}
	return true
}

// Sprint formats using the default formats for its operands and returns the
// resulting string wrapped in ANSI color sequences.
func (c *Color) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting
// string wrapped in ANSI color sequences.
func (c *Color) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// Sprintln formats using the default formats for its operands and returns the
// resulting string with a newline, wrapped in ANSI color sequences.
func (c *Color) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// SprintFunc returns a function that wraps Sprint with the current color settings.
func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.Sprint(a...)
	}
}

// SprintfFunc returns a function that wraps Sprintf with the current color settings.
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.Sprintf(format, a...)
	}
}

// SprintlnFunc returns a function that wraps Sprintln with the current color settings.
func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.Sprintln(a...)
	}
}

// PrintlnFunc returns a function that wraps Println with the current color settings.
// Note: the returned function has no return values, matching fatih/color's API.
func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		_, _ = c.Println(a...)
	}
}

// PrintFunc returns a function that wraps Print with the current color settings.
func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		_, _ = c.Print(a...)
	}
}

// PrintfFunc returns a function that wraps Printf with the current color settings.
func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		_, _ = c.Printf(format, a...)
	}
}

// FprintFunc returns a function that wraps Fprint with the current color settings.
func (c *Color) FprintFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		_, _ = c.Fprint(w, a...)
	}
}

// FprintfFunc returns a function that wraps Fprintf with the current color settings.
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...interface{}) {
	return func(w io.Writer, format string, a ...interface{}) {
		_, _ = c.Fprintf(w, format, a...)
	}
}

// FprintlnFunc returns a function that wraps Fprintln with the current color settings.
func (c *Color) FprintlnFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		_, _ = c.Fprintln(w, a...)
	}
}

// Print formats using the default formats and writes to standard output.
func (c *Color) Print(a ...interface{}) (int, error) {
	return c.Fprint(Output, a...)
}

// Printf formats according to a format specifier and writes to standard output.
func (c *Color) Printf(format string, a ...interface{}) (int, error) {
	return c.Fprintf(Output, format, a...)
}

// Println formats using the default formats and writes to standard output,
// followed by a newline.
func (c *Color) Println(a ...interface{}) (int, error) {
	return c.Fprintln(Output, a...)
}

// Fprint formats using the default formats and writes to w.
func (c *Color) Fprint(w io.Writer, a ...interface{}) (int, error) {
	s := c.wrap(fmt.Sprint(a...))
	return fmt.Fprint(w, s)
}

// Fprintf formats according to a format specifier and writes to w.
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	s := c.wrap(fmt.Sprintf(format, a...))
	return fmt.Fprint(w, s)
}

// Fprintln formats using the default formats and writes to w, followed by a newline.
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (int, error) {
	s := c.wrap(fmt.Sprintln(a...))
	return fmt.Fprint(w, s)
}

// Set sets the color attributes on the default output and returns the color
// for chaining. This is useful for stateful color changes (e.g., setting a
// color, printing multiple lines, then unsetting).
func (c *Color) Set() *Color {
	if c.isNoColor() {
		return c
	}
	_, _ = fmt.Fprint(Output, c.sequence())
	return c
}

// Unset resets the color on the default output.
func (c *Color) Unset() {
	if c.isNoColor() {
		return
	}
	_, _ = fmt.Fprint(Output, "\033[0m")
}

// DisableColor disables color output for this Color instance.
func (c *Color) DisableColor() {
	v := true
	c.noColor = &v
}

// EnableColor enables color output for this Color instance.
func (c *Color) EnableColor() {
	v := false
	c.noColor = &v
}

// isNoColor returns whether color output is disabled for this Color instance.
func (c *Color) isNoColor() bool {
	// Per-instance override takes precedence.
	if c.noColor != nil {
		return *c.noColor
	}
	// Global setting.
	return NoColor
}

// sequence returns the ANSI escape sequence prefix for the color's attributes.
func (c *Color) sequence() string {
	params := make([]string, len(c.params))
	for i, p := range c.params {
		params[i] = strconv.Itoa(int(p))
	}
	return "\033[" + strings.Join(params, ";") + "m"
}

// wrap wraps s with ANSI color sequences if color is enabled.
func (c *Color) wrap(s string) string {
	if c.isNoColor() {
		return s
	}
	return c.sequence() + s + "\033[0m"
}

// colorPrint is the internal print function used by package-level convenience
// functions. It is safe for concurrent use.
func colorPrint(format string, p Attribute, a ...interface{}) {
	c := getCachedColor(p)
	if len(a) == 0 {
		_, _ = c.Print(format)
	} else {
		_, _ = c.Printf(format, a...)
	}
}

// colorString is the internal string function used by package-level convenience
// functions. It is safe for concurrent use.
func colorString(format string, p Attribute, a ...interface{}) string {
	c := getCachedColor(p)
	if len(a) == 0 {
		return c.Sprint(format)
	}
	return c.Sprintf(format, a...)
}

// colorCache stores cached Color instances for each attribute, avoiding
// allocations on repeated use of convenience functions.
var (
	colorCache   = make(map[Attribute]*Color)
	colorCacheMu sync.Mutex
)

// getCachedColor returns a cached Color for the given attribute.
func getCachedColor(p Attribute) *Color {
	colorCacheMu.Lock()
	defer colorCacheMu.Unlock()
	c, ok := colorCache[p]
	if !ok {
		c = New(p)
		colorCache[p] = c
	}
	return c
}

// Global controls.
var (
	// NoColor disables all color output when set to true.
	// It respects the NO_COLOR environment variable (https://no-color.org/)
	// and TERM=dumb.
	NoColor = noColorFromEnv()

	// Output is the default writer for color output (typically os.Stdout).
	Output io.Writer = os.Stdout

	// Error is the default writer for error color output (typically os.Stderr).
	Error io.Writer = os.Stderr
)

func noColorFromEnv() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb"
}

// Unset resets the terminal colors on the default output.
func Unset() {
	_, _ = fmt.Fprint(Output, "\033[0m")
}
