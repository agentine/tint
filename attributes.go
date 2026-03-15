package tint

// Attribute defines a single ANSI SGR parameter.
type Attribute int

// Base text style attributes.
const (
	Reset      Attribute = 0
	Bold       Attribute = 1
	Faint      Attribute = 2
	Italic     Attribute = 3
	Underline  Attribute = 4
	BlinkSlow  Attribute = 5
	BlinkRapid Attribute = 6
	ReverseVideo Attribute = 7
	Concealed  Attribute = 8
	CrossedOut Attribute = 9
)

// Standard foreground colors (30–37).
const (
	FgBlack   Attribute = 30
	FgRed     Attribute = 31
	FgGreen   Attribute = 32
	FgYellow  Attribute = 33
	FgBlue    Attribute = 34
	FgMagenta Attribute = 35
	FgCyan    Attribute = 36
	FgWhite   Attribute = 37
)

// High-intensity foreground colors (90–97).
const (
	FgHiBlack   Attribute = 90
	FgHiRed     Attribute = 91
	FgHiGreen   Attribute = 92
	FgHiYellow  Attribute = 93
	FgHiBlue    Attribute = 94
	FgHiMagenta Attribute = 95
	FgHiCyan    Attribute = 96
	FgHiWhite   Attribute = 97
)

// Standard background colors (40–47).
const (
	BgBlack   Attribute = 40
	BgRed     Attribute = 41
	BgGreen   Attribute = 42
	BgYellow  Attribute = 43
	BgBlue    Attribute = 44
	BgMagenta Attribute = 45
	BgCyan    Attribute = 46
	BgWhite   Attribute = 47
)

// High-intensity background colors (100–107).
const (
	BgHiBlack   Attribute = 100
	BgHiRed     Attribute = 101
	BgHiGreen   Attribute = 102
	BgHiYellow  Attribute = 103
	BgHiBlue    Attribute = 104
	BgHiMagenta Attribute = 105
	BgHiCyan    Attribute = 106
	BgHiWhite   Attribute = 107
)
