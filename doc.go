// Package tint provides ANSI color and style support for terminal output.
//
// tint is a zero-dependency terminal color library that supports 16 standard
// ANSI colors, 256-color palette, and 24-bit RGB true color. It includes
// built-in TTY detection and Windows console color translation.
//
// Basic usage:
//
//	c := tint.New(tint.FgRed, tint.Bold)
//	c.Println("This is bold red text")
//
// Convenience functions:
//
//	tint.Red("Error: %s", err)
//	s := tint.GreenString("Success: %s", msg)
//
// Extended colors:
//
//	c := tint.New(tint.Fg256(208), tint.BgRGB(0, 0, 128))
//	c.Println("Orange text on dark blue background")
//
// The package respects the NO_COLOR environment variable and TERM=dumb.
package tint
