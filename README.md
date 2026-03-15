# tint

[![CI](https://github.com/agentine/tint/actions/workflows/ci.yml/badge.svg)](https://github.com/agentine/tint/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/agentine/tint.svg)](https://pkg.go.dev/github.com/agentine/tint)

Zero-dependency terminal color library for Go. Drop-in replacement for [fatih/color](https://github.com/fatih/color).

## Features

- **Drop-in compatible** with fatih/color API via `tint/compat/color`
- **Zero external dependencies** — built-in isatty detection and Windows color support
- **Extended colors** — 256-color palette and 24-bit RGB true color
- **Thread-safe** — safe for concurrent use
- **Cross-platform** — Linux, macOS, Windows (console API + VT sequence support)

## Install

```
go get github.com/agentine/tint
```

## Usage

### Basic colors

```go
import "github.com/agentine/tint"

// Using Color objects
c := tint.New(tint.FgRed, tint.Bold)
c.Println("Bold red text")

// Convenience functions
tint.Red("Error: %s", err)
tint.Green("Success!")

// String variants
msg := tint.RedString("failed: %s", reason)
```

### Extended colors

```go
// 256-color palette
hc := tint.HiColor(tint.Fg256(208))
hc.Sprint("orange text")

// 24-bit RGB true color
hc = tint.HiColor(tint.FgRGB(255, 128, 0), tint.BgRGB(0, 0, 64))
hc.Sprint("custom colors")
```

### Functional style

```go
red := tint.New(tint.FgRed).SprintFunc()
fmt.Println("This is", red("red"), "text")

warn := tint.New(tint.FgYellow, tint.Bold).SprintfFunc()
fmt.Println(warn("Warning: %d issues", count))
```

### Disable color

```go
// Globally
tint.NoColor = true

// Per-instance
c := tint.New(tint.FgRed)
c.DisableColor()

// Via environment: set NO_COLOR=1 or TERM=dumb
```

## Migrating from fatih/color

Change your import path:

```go
// Before
import "github.com/fatih/color"

// After
import "github.com/agentine/tint/compat/color"
```

No other code changes needed. All types, constants, and functions are identical.

## License

MIT
