# Changelog

## v0.1.0 — 2026-03-16

Initial release. A zero-dependency replacement for [fatih/color](https://github.com/fatih/color).

### Features

- **Core Color type** — `New()`, `Add()`, full Sprint/Sprintf/Sprintln, Print/Printf/Println, Fprint/Fprintf/Fprintln API with func-returning variants (SprintFunc, etc.)
- **42 ANSI attributes** — Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid, ReverseVideo, Concealed, CrossedOut, all 16 foreground colors (standard + high-intensity), all 16 background colors
- **256-color support** — `Fg256(code)` and `Bg256(code)` for extended palette
- **RGB true-color support** — `FgRGB(r, g, b)` and `BgRGB(r, g, b)` for 24-bit color
- **HiColor mixer** — combine extended colors with standard attributes
- **16 convenience print functions** — `Black()`, `Red()`, `Green()`, ..., `HiWhite()`
- **16 string functions** — `BlackString()`, `RedString()`, ..., `HiWhiteString()`
- **Built-in isatty detection** — Unix (TIOCGWINSZ syscall), Windows (GetConsoleMode), Plan9 (/dev/cons). Zero external dependencies.
- **Built-in Windows colorable writer** — ANSI-to-ConsoleAPI translation with VT passthrough detection for modern terminals
- **Global controls** — `NoColor`, `Output`, `Error`, `Unset()`, `NO_COLOR` env var support
- **Thread-safe** — safe for concurrent use
- **Drop-in compatibility** — `compat/color` package provides 100% fatih/color API compatibility (type aliases, all constants, all functions)
- **CI** — tested on Ubuntu, macOS, Windows across Go 1.21, 1.22, 1.23
