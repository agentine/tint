package tint

import (
	"bytes"
	"testing"
)

func TestConvenienceStringFunctions(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	tests := []struct {
		name string
		fn   func(string, ...interface{}) string
		code int
	}{
		{"BlackString", BlackString, 30},
		{"RedString", RedString, 31},
		{"GreenString", GreenString, 32},
		{"YellowString", YellowString, 33},
		{"BlueString", BlueString, 34},
		{"MagentaString", MagentaString, 35},
		{"CyanString", CyanString, 36},
		{"WhiteString", WhiteString, 37},
		{"HiBlackString", HiBlackString, 90},
		{"HiRedString", HiRedString, 91},
		{"HiGreenString", HiGreenString, 92},
		{"HiYellowString", HiYellowString, 93},
		{"HiBlueString", HiBlueString, 94},
		{"HiMagentaString", HiMagentaString, 95},
		{"HiCyanString", HiCyanString, 96},
		{"HiWhiteString", HiWhiteString, 97},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn("hello")
			want := "\033[" + itoa(tt.code) + "mhello\033[0m"
			if got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestConvenienceStringFormat(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	got := RedString("n=%d", 5)
	want := "\033[31mn=5\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestConveniencePrintFunctions(t *testing.T) {
	orig := NoColor
	origOut := Output
	NoColor = false
	defer func() { NoColor = orig; Output = origOut }()

	tests := []struct {
		name string
		fn   func(string, ...interface{})
		code int
	}{
		{"Black", Black, 30},
		{"Red", Red, 31},
		{"Green", Green, 32},
		{"Yellow", Yellow, 33},
		{"Blue", Blue, 34},
		{"Magenta", Magenta, 35},
		{"Cyan", Cyan, 36},
		{"White", White, 37},
		{"HiBlack", HiBlack, 90},
		{"HiRed", HiRed, 91},
		{"HiGreen", HiGreen, 92},
		{"HiYellow", HiYellow, 93},
		{"HiBlue", HiBlue, 94},
		{"HiMagenta", HiMagenta, 95},
		{"HiCyan", HiCyan, 96},
		{"HiWhite", HiWhite, 97},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			Output = &buf
			tt.fn("hello")
			got := buf.String()
			want := "\033[" + itoa(tt.code) + "mhello\033[0m"
			if got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestConvenienceNoColor(t *testing.T) {
	orig := NoColor
	NoColor = true
	defer func() { NoColor = orig }()

	got := RedString("hello")
	if got != "hello" {
		t.Fatalf("expected plain text, got %q", got)
	}
}

func itoa(i int) string {
	s := ""
	if i == 0 {
		return "0"
	}
	for i > 0 {
		s = string(rune('0'+i%10)) + s
		i /= 10
	}
	return s
}
