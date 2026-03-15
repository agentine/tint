// Package color provides a drop-in replacement for github.com/fatih/color.
//
// To migrate existing code, change only the import path:
//
//	// Before
//	import "github.com/fatih/color"
//
//	// After
//	import "github.com/agentine/tint/compat/color"
//
// All types, constants, functions, and methods are identical to fatih/color.
package color

import (
	"io"

	"github.com/agentine/tint"
)

// Attribute defines a single ANSI SGR parameter.
// This is a type alias for fatih/color compatibility.
type Attribute = tint.Attribute

// Base text style attributes.
const (
	Reset        = tint.Reset
	Bold         = tint.Bold
	Faint        = tint.Faint
	Italic       = tint.Italic
	Underline    = tint.Underline
	BlinkSlow    = tint.BlinkSlow
	BlinkRapid   = tint.BlinkRapid
	ReverseVideo = tint.ReverseVideo
	Concealed    = tint.Concealed
	CrossedOut   = tint.CrossedOut
)

// Standard foreground colors.
const (
	FgBlack   = tint.FgBlack
	FgRed     = tint.FgRed
	FgGreen   = tint.FgGreen
	FgYellow  = tint.FgYellow
	FgBlue    = tint.FgBlue
	FgMagenta = tint.FgMagenta
	FgCyan    = tint.FgCyan
	FgWhite   = tint.FgWhite
)

// High-intensity foreground colors.
const (
	FgHiBlack   = tint.FgHiBlack
	FgHiRed     = tint.FgHiRed
	FgHiGreen   = tint.FgHiGreen
	FgHiYellow  = tint.FgHiYellow
	FgHiBlue    = tint.FgHiBlue
	FgHiMagenta = tint.FgHiMagenta
	FgHiCyan    = tint.FgHiCyan
	FgHiWhite   = tint.FgHiWhite
)

// Standard background colors.
const (
	BgBlack   = tint.BgBlack
	BgRed     = tint.BgRed
	BgGreen   = tint.BgGreen
	BgYellow  = tint.BgYellow
	BgBlue    = tint.BgBlue
	BgMagenta = tint.BgMagenta
	BgCyan    = tint.BgCyan
	BgWhite   = tint.BgWhite
)

// High-intensity background colors.
const (
	BgHiBlack   = tint.BgHiBlack
	BgHiRed     = tint.BgHiRed
	BgHiGreen   = tint.BgHiGreen
	BgHiYellow  = tint.BgHiYellow
	BgHiBlue    = tint.BgHiBlue
	BgHiMagenta = tint.BgHiMagenta
	BgHiCyan    = tint.BgHiCyan
	BgHiWhite   = tint.BgHiWhite
)

// Color is a type alias for tint.Color, providing the same API as fatih/color.Color.
type Color = tint.Color

// New creates a new Color with the given attributes.
func New(attrs ...Attribute) *Color {
	return tint.New(attrs...)
}

// NoColor is a global toggle for disabling color output.
// It respects the NO_COLOR env var and TERM=dumb.
//
// Note: this is a function that reads from tint.NoColor. For direct assignment,
// set tint.NoColor instead, or use the Set* functions below.
func GetNoColor() bool {
	return tint.NoColor
}

// SetNoColor sets the global NoColor flag.
func SetNoColor(v bool) {
	tint.NoColor = v
}

// SetOutput sets the default output writer.
func SetOutput(w io.Writer) {
	tint.Output = w
}

// Convenience print functions — standard colors.

func Black(format string, a ...interface{})   { tint.Black(format, a...) }
func Red(format string, a ...interface{})     { tint.Red(format, a...) }
func Green(format string, a ...interface{})   { tint.Green(format, a...) }
func Yellow(format string, a ...interface{})  { tint.Yellow(format, a...) }
func Blue(format string, a ...interface{})    { tint.Blue(format, a...) }
func Magenta(format string, a ...interface{}) { tint.Magenta(format, a...) }
func Cyan(format string, a ...interface{})    { tint.Cyan(format, a...) }
func White(format string, a ...interface{})   { tint.White(format, a...) }

// Convenience print functions — high-intensity colors.

func HiBlack(format string, a ...interface{})   { tint.HiBlack(format, a...) }
func HiRed(format string, a ...interface{})     { tint.HiRed(format, a...) }
func HiGreen(format string, a ...interface{})   { tint.HiGreen(format, a...) }
func HiYellow(format string, a ...interface{})  { tint.HiYellow(format, a...) }
func HiBlue(format string, a ...interface{})    { tint.HiBlue(format, a...) }
func HiMagenta(format string, a ...interface{}) { tint.HiMagenta(format, a...) }
func HiCyan(format string, a ...interface{})    { tint.HiCyan(format, a...) }
func HiWhite(format string, a ...interface{})   { tint.HiWhite(format, a...) }

// Convenience string functions — standard colors.

func BlackString(format string, a ...interface{}) string   { return tint.BlackString(format, a...) }
func RedString(format string, a ...interface{}) string     { return tint.RedString(format, a...) }
func GreenString(format string, a ...interface{}) string   { return tint.GreenString(format, a...) }
func YellowString(format string, a ...interface{}) string  { return tint.YellowString(format, a...) }
func BlueString(format string, a ...interface{}) string    { return tint.BlueString(format, a...) }
func MagentaString(format string, a ...interface{}) string { return tint.MagentaString(format, a...) }
func CyanString(format string, a ...interface{}) string    { return tint.CyanString(format, a...) }
func WhiteString(format string, a ...interface{}) string   { return tint.WhiteString(format, a...) }

// Convenience string functions — high-intensity colors.

func HiBlackString(format string, a ...interface{}) string   { return tint.HiBlackString(format, a...) }
func HiRedString(format string, a ...interface{}) string     { return tint.HiRedString(format, a...) }
func HiGreenString(format string, a ...interface{}) string   { return tint.HiGreenString(format, a...) }
func HiYellowString(format string, a ...interface{}) string  { return tint.HiYellowString(format, a...) }
func HiBlueString(format string, a ...interface{}) string    { return tint.HiBlueString(format, a...) }
func HiMagentaString(format string, a ...interface{}) string { return tint.HiMagentaString(format, a...) }
func HiCyanString(format string, a ...interface{}) string    { return tint.HiCyanString(format, a...) }
func HiWhiteString(format string, a ...interface{}) string   { return tint.HiWhiteString(format, a...) }

// Unset resets terminal colors.
func Unset() { tint.Unset() }
