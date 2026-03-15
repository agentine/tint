# tint — Terminal Color Library for Go

## Overview

**Replaces:** [fatih/color](https://github.com/fatih/color) (26,630 importers, 7k+ stars, archived Oct 2024, author on indefinite sabbatical)

**Package:** `github.com/agentine/tint`

**Language:** Go

**Why:** fatih/color is the most widely used terminal color package in Go (26,630 importers). The author archived all his open source projects in mid-2024 and released a final v1.18.0 in October 2024 before stepping away indefinitely. The gofrs community discussed adopting it but declined. No maintained fork exists. Issues continue to be filed (as recently as Oct 2025) with no response. The package depends on mattn/go-colorable and mattn/go-isatty (still maintained but single-maintainer packages). A zero-dependency replacement with API compatibility eliminates the dependency chain risk entirely.

## Design Goals

1. **Drop-in compatible** with fatih/color API
2. **Zero dependencies** — built-in isatty detection and Windows color support
3. **Extended color support** — 256-color and RGB/true-color (24-bit) in addition to basic 16 colors
4. **Thread-safe** — safe for concurrent use
5. **Modern Go** — requires Go 1.21+, uses generics where beneficial

## Architecture

### Package Structure

```
tint/
├── tint.go              # Core Color type, attributes, Sprint/Fprint/Println API
├── attributes.go        # Color attribute constants (FgRed, Bold, etc.)
├── writer.go            # ColorableWriter for Windows ANSI translation
├── isatty.go            # TTY detection (unix)
├── isatty_windows.go    # TTY detection (windows)
├── colorable_windows.go # Windows console API color translation
├── colorable_other.go   # No-op on Unix (ANSI native)
├── hicolor.go           # 256-color and RGB/true-color support
├── global.go            # Package-level convenience functions (Red, Blue, etc.)
├── compat/
│   └── color/           # Drop-in fatih/color replacement package
│       └── color.go     # Re-exports with identical API surface
└── doc.go               # Package documentation
```

### Core Components

#### 1. Color Type

```go
type Color struct {
    params  []Attribute
    noColor *bool  // per-instance override
}

func New(attrs ...Attribute) *Color
func (c *Color) Add(attrs ...Attribute) *Color
func (c *Color) Sprint(a ...interface{}) string
func (c *Color) Sprintf(format string, a ...interface{}) string
func (c *Color) Sprintln(a ...interface{}) string
func (c *Color) Print(a ...interface{}) (int, error)
func (c *Color) Printf(format string, a ...interface{}) (int, error)
func (c *Color) Println(a ...interface{}) (int, error)
func (c *Color) Fprint(w io.Writer, a ...interface{}) (int, error)
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (int, error)
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (int, error)
func (c *Color) SprintFunc() func(a ...interface{}) string
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string
func (c *Color) PrintlnFunc() func(a ...interface{})
func (c *Color) Set() *Color          // Set terminal color
func (c *Color) Unset()               // Reset terminal color
func (c *Color) DisableColor()
func (c *Color) EnableColor()
```

#### 2. Attributes

Standard 16-color ANSI attributes:
- Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid, ReverseVideo, Concealed, CrossedOut
- FgBlack, FgRed, FgGreen, FgYellow, FgBlue, FgMagenta, FgCyan, FgWhite
- FgHiBlack, FgHiRed, FgHiGreen, FgHiYellow, FgHiBlue, FgHiMagenta, FgHiCyan, FgHiWhite
- BgBlack, BgRed, ..., BgHiWhite

#### 3. Extended Color Support (enhancement)

```go
func Fg256(code uint8) Attribute      // 256-color foreground
func Bg256(code uint8) Attribute      // 256-color background
func FgRGB(r, g, b uint8) Attribute   // True-color foreground
func BgRGB(r, g, b uint8) Attribute   // True-color background
```

#### 4. Convenience Functions

```go
func Black(format string, a ...interface{})   { ... }
func Red(format string, a ...interface{})     { ... }
func Green(format string, a ...interface{})   { ... }
// ... etc for all 16 colors
func BlackString(format string, a ...interface{}) string { ... }
func RedString(format string, a ...interface{}) string   { ... }
// ... etc
```

#### 5. Global Controls

```go
var NoColor = os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb"
var Output io.Writer = colorable.NewColorableStdout()
var Error  io.Writer = colorable.NewColorableStderr()
func Unset()  // Reset terminal colors
```

#### 6. Built-in isatty

Platform-specific TTY detection without external dependencies:
- Unix: `syscall.SYS_IOCTL` with `TIOCGWINSZ`
- Windows: `GetConsoleMode` via syscall
- Plan9: check `/dev/cons`

#### 7. Built-in Windows Colorable Writer

Translates ANSI escape sequences to Windows Console API calls:
- Parse ANSI SGR sequences from output stream
- Map to `SetConsoleTextAttribute` calls
- Support for 16-color, 256-color, and true-color (Windows 10+)
- Pass-through on terminals that support ANSI natively (Windows Terminal, WSL)

#### 8. Compatibility Package

`tint/compat/color` provides 100% fatih/color API compatibility:
- Same type names, function signatures, constants
- Import path swap: `github.com/fatih/color` → `github.com/agentine/tint/compat/color`
- No behavior changes for existing code

## Deliverables

1. Core `tint` package with full Color API
2. Built-in isatty detection (zero deps)
3. Built-in Windows colorable writer (zero deps)
4. 256-color and RGB/true-color extensions
5. `tint/compat/color` drop-in compatibility package
6. Comprehensive test suite with platform-specific tests
7. Benchmarks comparing to fatih/color
8. README with migration guide

## Non-Goals

- GUI/rich-text rendering
- Image/pixel color manipulation
- Complex layout or box drawing (that's a TUI framework concern)
