package color

import (
	"bytes"
	"testing"

	"github.com/agentine/tint"
)

func TestCompatNew(t *testing.T) {
	c := New(FgRed, Bold)
	if c == nil {
		t.Fatal("expected non-nil Color")
	}
}

func TestCompatSprint(t *testing.T) {
	orig := tint.NoColor
	tint.NoColor = false
	defer func() { tint.NoColor = orig }()

	c := New(FgRed)
	got := c.Sprint("hello")
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestCompatConstants(t *testing.T) {
	// Verify constants match tint values.
	if FgRed != tint.FgRed {
		t.Fatal("FgRed mismatch")
	}
	if Bold != tint.Bold {
		t.Fatal("Bold mismatch")
	}
	if BgHiWhite != tint.BgHiWhite {
		t.Fatal("BgHiWhite mismatch")
	}
}

func TestCompatStringFunctions(t *testing.T) {
	orig := tint.NoColor
	tint.NoColor = false
	defer func() { tint.NoColor = orig }()

	got := RedString("err: %s", "fail")
	want := "\033[31merr: fail\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestCompatPrintFunctions(t *testing.T) {
	orig := tint.NoColor
	origOut := tint.Output
	tint.NoColor = false
	var buf bytes.Buffer
	tint.Output = &buf
	defer func() { tint.NoColor = orig; tint.Output = origOut }()

	Red("hello")
	got := buf.String()
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestCompatNoColor(t *testing.T) {
	orig := tint.NoColor
	tint.NoColor = true
	defer func() { tint.NoColor = orig }()

	got := GreenString("ok")
	if got != "ok" {
		t.Fatalf("expected plain text, got %q", got)
	}
}

func TestCompatSetNoColor(t *testing.T) {
	orig := tint.NoColor
	defer func() { tint.NoColor = orig }()

	SetNoColor(true)
	if !GetNoColor() {
		t.Fatal("expected NoColor=true")
	}
	SetNoColor(false)
	if GetNoColor() {
		t.Fatal("expected NoColor=false")
	}
}

func TestCompatSetOutput(t *testing.T) {
	origOut := tint.Output
	defer func() { tint.Output = origOut }()

	var buf bytes.Buffer
	SetOutput(&buf)
	// Verify the output is now the buffer.
	orig := tint.NoColor
	tint.NoColor = false
	defer func() { tint.NoColor = orig }()

	c := New(FgBlue)
	_, err := c.Print("test")
	if err != nil {
		t.Fatal(err)
	}
	want := "\033[34mtest\033[0m"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestCompatUnset(t *testing.T) {
	origOut := tint.Output
	orig := tint.NoColor
	tint.NoColor = false
	var buf bytes.Buffer
	tint.Output = &buf
	defer func() { tint.NoColor = orig; tint.Output = origOut }()

	Unset()
	if buf.String() != "\033[0m" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestCompatTypeAlias(t *testing.T) {
	// Verify Color type alias works — a *Color returned by New()
	// should be assignable to *tint.Color.
	c := New(FgRed)
	var _ *tint.Color = c //nolint:staticcheck // intentional type compatibility check
}

func TestCompatAttributeAlias(t *testing.T) {
	// Verify Attribute type alias — both directions.
	a := FgRed
	var _ tint.Attribute = a //nolint:staticcheck // intentional type compatibility check
	b := tint.FgRed
	var _ Attribute = b //nolint:staticcheck // intentional type compatibility check
}

func TestAllStdPrintFunctions(t *testing.T) {
	orig := tint.NoColor
	origOut := tint.Output
	tint.NoColor = false
	defer func() { tint.NoColor = orig; tint.Output = origOut }()

	fns := []struct {
		name string
		fn   func(string, ...interface{})
	}{
		{"Black", Black},
		{"Green", Green},
		{"Yellow", Yellow},
		{"Blue", Blue},
		{"Magenta", Magenta},
		{"Cyan", Cyan},
		{"White", White},
		{"HiBlack", HiBlack},
		{"HiRed", HiRed},
		{"HiGreen", HiGreen},
		{"HiYellow", HiYellow},
		{"HiBlue", HiBlue},
		{"HiMagenta", HiMagenta},
		{"HiCyan", HiCyan},
		{"HiWhite", HiWhite},
	}
	for _, tt := range fns {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tint.Output = &buf
			tt.fn("x")
			if buf.Len() == 0 {
				t.Fatalf("%s produced no output", tt.name)
			}
		})
	}
}

func TestAllStdStringFunctions(t *testing.T) {
	orig := tint.NoColor
	tint.NoColor = false
	defer func() { tint.NoColor = orig }()

	fns := []struct {
		name string
		fn   func(string, ...interface{}) string
	}{
		{"BlackString", BlackString},
		{"YellowString", YellowString},
		{"BlueString", BlueString},
		{"MagentaString", MagentaString},
		{"CyanString", CyanString},
		{"WhiteString", WhiteString},
	}
	for _, tt := range fns {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn("x")
			if got == "x" {
				t.Fatalf("%s returned uncolored text", tt.name)
			}
		})
	}
}

func TestAllHiStringFunctions(t *testing.T) {
	orig := tint.NoColor
	tint.NoColor = false
	defer func() { tint.NoColor = orig }()

	fns := []struct {
		name string
		fn   func(string, ...interface{}) string
	}{
		{"HiBlackString", HiBlackString},
		{"HiRedString", HiRedString},
		{"HiGreenString", HiGreenString},
		{"HiYellowString", HiYellowString},
		{"HiBlueString", HiBlueString},
		{"HiMagentaString", HiMagentaString},
		{"HiCyanString", HiCyanString},
		{"HiWhiteString", HiWhiteString},
	}
	for _, tt := range fns {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn("x")
			if got == "x" {
				t.Fatalf("%s returned uncolored text", tt.name)
			}
			if len(got) < 5 {
				t.Fatalf("%s output too short: %q", tt.name, got)
			}
		})
	}
}
