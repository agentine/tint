package tint

import "strconv"

// hiColorAttr represents an extended color attribute (256-color or RGB)
// that generates multi-parameter ANSI sequences.
type hiColorAttr struct {
	seq string
}

// Fg256 returns an Attribute for a 256-color foreground.
// Code values 0–255 map to the standard 256-color palette.
func Fg256(code uint8) *hiColorAttr {
	return &hiColorAttr{seq: "38;5;" + strconv.Itoa(int(code))}
}

// Bg256 returns an Attribute for a 256-color background.
func Bg256(code uint8) *hiColorAttr {
	return &hiColorAttr{seq: "48;5;" + strconv.Itoa(int(code))}
}

// FgRGB returns an Attribute for a 24-bit true-color foreground.
func FgRGB(r, g, b uint8) *hiColorAttr {
	return &hiColorAttr{
		seq: "38;2;" + strconv.Itoa(int(r)) + ";" + strconv.Itoa(int(g)) + ";" + strconv.Itoa(int(b)),
	}
}

// BgRGB returns an Attribute for a 24-bit true-color background.
func BgRGB(r, g, b uint8) *hiColorAttr {
	return &hiColorAttr{
		seq: "48;2;" + strconv.Itoa(int(r)) + ";" + strconv.Itoa(int(g)) + ";" + strconv.Itoa(int(b)),
	}
}

// HiColor creates a new Color with extended (256/RGB) color attributes.
// It accepts a mix of standard Attribute values and *hiColorAttr values.
//
// Example:
//
//	c := tint.HiColor(tint.Bold, tint.Fg256(208))
//	c := tint.HiColor(tint.FgRGB(255, 128, 0), tint.BgRGB(0, 0, 128))
func HiColor(attrs ...interface{}) *HiColorValue {
	hc := &HiColorValue{}
	for _, a := range attrs {
		switch v := a.(type) {
		case Attribute:
			hc.stdAttrs = append(hc.stdAttrs, v)
		case *hiColorAttr:
			hc.hiAttrs = append(hc.hiAttrs, v)
		}
	}
	return hc
}

// HiColorValue holds a combination of standard and extended color attributes.
type HiColorValue struct {
	stdAttrs []Attribute
	hiAttrs  []*hiColorAttr
	noColor  *bool
}

// Sprint formats using default formats and returns the colored string.
func (hc *HiColorValue) Sprint(a ...interface{}) string {
	return hc.wrap(sprint(a...))
}

// Sprintf formats according to a format specifier and returns the colored string.
func (hc *HiColorValue) Sprintf(format string, a ...interface{}) string {
	return hc.wrap(sprintf(format, a...))
}

// Sprintln formats and returns the colored string with a trailing newline.
func (hc *HiColorValue) Sprintln(a ...interface{}) string {
	return hc.wrap(sprintln(a...))
}

// DisableColor disables color for this HiColorValue instance.
func (hc *HiColorValue) DisableColor() {
	v := true
	hc.noColor = &v
}

// EnableColor enables color for this HiColorValue instance.
func (hc *HiColorValue) EnableColor() {
	v := false
	hc.noColor = &v
}

func (hc *HiColorValue) isNoColor() bool {
	if hc.noColor != nil {
		return *hc.noColor
	}
	return NoColor
}

func (hc *HiColorValue) sequence() string {
	parts := make([]string, 0, len(hc.stdAttrs)+len(hc.hiAttrs))
	for _, a := range hc.stdAttrs {
		parts = append(parts, strconv.Itoa(int(a)))
	}
	for _, h := range hc.hiAttrs {
		parts = append(parts, h.seq)
	}
	s := ""
	for i, p := range parts {
		if i > 0 {
			s += ";"
		}
		s += p
	}
	return "\033[" + s + "m"
}

func (hc *HiColorValue) wrap(s string) string {
	if hc.isNoColor() {
		return s
	}
	return hc.sequence() + s + "\033[0m"
}
