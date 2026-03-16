package tint

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	c := New(FgRed)
	if len(c.params) != 1 || c.params[0] != FgRed {
		t.Fatalf("expected [FgRed], got %v", c.params)
	}
}

func TestAdd(t *testing.T) {
	c := New(FgRed)
	c.Add(Bold)
	if len(c.params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(c.params))
	}
	if c.params[1] != Bold {
		t.Fatalf("expected Bold, got %v", c.params[1])
	}
}

func TestAddChaining(t *testing.T) {
	c := New().Add(FgRed).Add(Bold)
	if len(c.params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(c.params))
	}
}

func TestEquals(t *testing.T) {
	a := New(FgRed, Bold)
	b := New(FgRed, Bold)
	c := New(FgBlue, Bold)
	if !a.Equals(b) {
		t.Fatal("expected equal")
	}
	if a.Equals(c) {
		t.Fatal("expected not equal")
	}
}

func TestSprint(t *testing.T) {
	// Force color on for testing.
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgRed)
	got := c.Sprint("hello")
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSprintMultipleAttrs(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgRed, Bold)
	got := c.Sprint("hello")
	want := "\033[31;1mhello\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSprintf(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgGreen)
	got := c.Sprintf("count: %d", 42)
	want := "\033[32mcount: 42\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSprintln(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgBlue)
	got := c.Sprintln("hello")
	want := "\033[34mhello\n\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFprint(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	var buf bytes.Buffer
	c := New(FgYellow)
	n, err := c.Fprint(&buf, "hello")
	if err != nil {
		t.Fatal(err)
	}
	want := "\033[33mhello\033[0m"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
	if n != len(want) {
		t.Fatalf("got n=%d, want %d", n, len(want))
	}
}

func TestFprintf(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	var buf bytes.Buffer
	c := New(FgCyan)
	_, err := c.Fprintf(&buf, "val=%d", 7)
	if err != nil {
		t.Fatal(err)
	}
	want := "\033[36mval=7\033[0m"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestFprintln(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	var buf bytes.Buffer
	c := New(FgMagenta)
	_, err := c.Fprintln(&buf, "hello")
	if err != nil {
		t.Fatal(err)
	}
	want := "\033[35mhello\n\033[0m"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestPrintToCustomOutput(t *testing.T) {
	orig := NoColor
	origOut := Output
	NoColor = false
	var buf bytes.Buffer
	Output = &buf
	defer func() { NoColor = orig; Output = origOut }()

	c := New(FgRed)
	_, err := c.Print("test")
	if err != nil {
		t.Fatal(err)
	}
	want := "\033[31mtest\033[0m"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestNoColor(t *testing.T) {
	orig := NoColor
	NoColor = true
	defer func() { NoColor = orig }()

	c := New(FgRed, Bold)
	got := c.Sprint("hello")
	if got != "hello" {
		t.Fatalf("expected plain text, got %q", got)
	}
}

func TestDisableEnableColor(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgRed)
	c.DisableColor()
	got := c.Sprint("hello")
	if got != "hello" {
		t.Fatalf("expected plain text after DisableColor, got %q", got)
	}

	c.EnableColor()
	got = c.Sprint("hello")
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Fatalf("after EnableColor got %q, want %q", got, want)
	}
}

func TestPerInstanceNoColorOverridesGlobal(t *testing.T) {
	orig := NoColor
	NoColor = true
	defer func() { NoColor = orig }()

	c := New(FgRed)
	c.EnableColor()
	got := c.Sprint("hello")
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Fatalf("per-instance EnableColor should override global NoColor, got %q, want %q", got, want)
	}
}

func TestSetUnset(t *testing.T) {
	orig := NoColor
	origOut := Output
	NoColor = false
	var buf bytes.Buffer
	Output = &buf
	defer func() { NoColor = orig; Output = origOut }()

	c := New(FgRed)
	c.Set()
	fmt.Fprint(&buf, "colored")
	c.Unset()

	got := buf.String()
	want := "\033[31mcolored\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSetUnsetNoColor(t *testing.T) {
	orig := NoColor
	origOut := Output
	NoColor = true
	var buf bytes.Buffer
	Output = &buf
	defer func() { NoColor = orig; Output = origOut }()

	c := New(FgRed)
	c.Set()
	fmt.Fprint(&buf, "plain")
	c.Unset()

	got := buf.String()
	if got != "plain" {
		t.Fatalf("expected plain text, got %q", got)
	}
}

func TestSprintFunc(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	fn := New(FgRed).SprintFunc()
	got := fn("hello")
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSprintfFunc(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	fn := New(FgGreen).SprintfFunc()
	got := fn("n=%d", 5)
	want := "\033[32mn=5\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSprintlnFunc(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	fn := New(FgBlue).SprintlnFunc()
	got := fn("hello")
	want := "\033[34mhello\n\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestPrintlnFunc(t *testing.T) {
	orig := NoColor
	origOut := Output
	NoColor = false
	var buf bytes.Buffer
	Output = &buf
	defer func() { NoColor = orig; Output = origOut }()

	fn := New(FgRed).PrintlnFunc()
	fn("test")
	want := "\033[31mtest\n\033[0m"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestAllAttributes(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	attrs := []struct {
		name string
		attr Attribute
		code int
	}{
		{"Reset", Reset, 0},
		{"Bold", Bold, 1},
		{"Faint", Faint, 2},
		{"Italic", Italic, 3},
		{"Underline", Underline, 4},
		{"BlinkSlow", BlinkSlow, 5},
		{"BlinkRapid", BlinkRapid, 6},
		{"ReverseVideo", ReverseVideo, 7},
		{"Concealed", Concealed, 8},
		{"CrossedOut", CrossedOut, 9},
		{"FgBlack", FgBlack, 30},
		{"FgRed", FgRed, 31},
		{"FgGreen", FgGreen, 32},
		{"FgYellow", FgYellow, 33},
		{"FgBlue", FgBlue, 34},
		{"FgMagenta", FgMagenta, 35},
		{"FgCyan", FgCyan, 36},
		{"FgWhite", FgWhite, 37},
		{"FgHiBlack", FgHiBlack, 90},
		{"FgHiRed", FgHiRed, 91},
		{"FgHiGreen", FgHiGreen, 92},
		{"FgHiYellow", FgHiYellow, 93},
		{"FgHiBlue", FgHiBlue, 94},
		{"FgHiMagenta", FgHiMagenta, 95},
		{"FgHiCyan", FgHiCyan, 96},
		{"FgHiWhite", FgHiWhite, 97},
		{"BgBlack", BgBlack, 40},
		{"BgRed", BgRed, 41},
		{"BgGreen", BgGreen, 42},
		{"BgYellow", BgYellow, 43},
		{"BgBlue", BgBlue, 44},
		{"BgMagenta", BgMagenta, 45},
		{"BgCyan", BgCyan, 46},
		{"BgWhite", BgWhite, 47},
		{"BgHiBlack", BgHiBlack, 100},
		{"BgHiRed", BgHiRed, 101},
		{"BgHiGreen", BgHiGreen, 102},
		{"BgHiYellow", BgHiYellow, 103},
		{"BgHiBlue", BgHiBlue, 104},
		{"BgHiMagenta", BgHiMagenta, 105},
		{"BgHiCyan", BgHiCyan, 106},
		{"BgHiWhite", BgHiWhite, 107},
	}

	for _, tt := range attrs {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.attr) != tt.code {
				t.Fatalf("%s: expected %d, got %d", tt.name, tt.code, int(tt.attr))
			}
			c := New(tt.attr)
			got := c.Sprint("x")
			want := fmt.Sprintf("\033[%dmx\033[0m", tt.code)
			if got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestNoColorEnvVar(t *testing.T) {
	// Test that noColorFromEnv respects NO_COLOR.
	orig := os.Getenv("NO_COLOR")
	defer func() { _ = os.Setenv("NO_COLOR", orig) }()

	t.Setenv("NO_COLOR", "1")
	if !noColorFromEnv() {
		t.Fatal("expected noColorFromEnv to return true when NO_COLOR=1")
	}

	t.Setenv("NO_COLOR", "")
	origTerm := os.Getenv("TERM")
	defer func() { _ = os.Setenv("TERM", origTerm) }()
	t.Setenv("TERM", "dumb")
	if !noColorFromEnv() {
		t.Fatal("expected noColorFromEnv to return true when TERM=dumb")
	}

	t.Setenv("TERM", "xterm")
	if noColorFromEnv() {
		t.Fatal("expected noColorFromEnv to return false")
	}
}

func TestGlobalUnset(t *testing.T) {
	orig := NoColor
	origOut := Output
	NoColor = false
	var buf bytes.Buffer
	Output = &buf
	defer func() { NoColor = orig; Output = origOut }()

	Unset()
	if buf.String() != "\033[0m" {
		t.Fatalf("got %q, want %q", buf.String(), "\033[0m")
	}
}

func TestColorCache(t *testing.T) {
	c1 := getCachedColor(FgRed)
	c2 := getCachedColor(FgRed)
	if c1 != c2 {
		t.Fatal("expected same cached instance")
	}
}
