package tint

// Convenience print functions for standard colors.
// Each function prints to Output using Printf-style formatting.

func Black(format string, a ...interface{})   { colorPrint(format, FgBlack, a...) }
func Red(format string, a ...interface{})     { colorPrint(format, FgRed, a...) }
func Green(format string, a ...interface{})   { colorPrint(format, FgGreen, a...) }
func Yellow(format string, a ...interface{})  { colorPrint(format, FgYellow, a...) }
func Blue(format string, a ...interface{})    { colorPrint(format, FgBlue, a...) }
func Magenta(format string, a ...interface{}) { colorPrint(format, FgMagenta, a...) }
func Cyan(format string, a ...interface{})    { colorPrint(format, FgCyan, a...) }
func White(format string, a ...interface{})   { colorPrint(format, FgWhite, a...) }

// High-intensity convenience print functions.

func HiBlack(format string, a ...interface{})   { colorPrint(format, FgHiBlack, a...) }
func HiRed(format string, a ...interface{})     { colorPrint(format, FgHiRed, a...) }
func HiGreen(format string, a ...interface{})   { colorPrint(format, FgHiGreen, a...) }
func HiYellow(format string, a ...interface{})  { colorPrint(format, FgHiYellow, a...) }
func HiBlue(format string, a ...interface{})    { colorPrint(format, FgHiBlue, a...) }
func HiMagenta(format string, a ...interface{}) { colorPrint(format, FgHiMagenta, a...) }
func HiCyan(format string, a ...interface{})    { colorPrint(format, FgHiCyan, a...) }
func HiWhite(format string, a ...interface{})   { colorPrint(format, FgHiWhite, a...) }

// Convenience string functions for standard colors.

func BlackString(format string, a ...interface{}) string   { return colorString(format, FgBlack, a...) }
func RedString(format string, a ...interface{}) string     { return colorString(format, FgRed, a...) }
func GreenString(format string, a ...interface{}) string   { return colorString(format, FgGreen, a...) }
func YellowString(format string, a ...interface{}) string  { return colorString(format, FgYellow, a...) }
func BlueString(format string, a ...interface{}) string    { return colorString(format, FgBlue, a...) }
func MagentaString(format string, a ...interface{}) string { return colorString(format, FgMagenta, a...) }
func CyanString(format string, a ...interface{}) string    { return colorString(format, FgCyan, a...) }
func WhiteString(format string, a ...interface{}) string   { return colorString(format, FgWhite, a...) }

// High-intensity convenience string functions.

func HiBlackString(format string, a ...interface{}) string   { return colorString(format, FgHiBlack, a...) }
func HiRedString(format string, a ...interface{}) string     { return colorString(format, FgHiRed, a...) }
func HiGreenString(format string, a ...interface{}) string   { return colorString(format, FgHiGreen, a...) }
func HiYellowString(format string, a ...interface{}) string  { return colorString(format, FgHiYellow, a...) }
func HiBlueString(format string, a ...interface{}) string    { return colorString(format, FgHiBlue, a...) }
func HiMagentaString(format string, a ...interface{}) string { return colorString(format, FgHiMagenta, a...) }
func HiCyanString(format string, a ...interface{}) string    { return colorString(format, FgHiCyan, a...) }
func HiWhiteString(format string, a ...interface{}) string   { return colorString(format, FgHiWhite, a...) }
