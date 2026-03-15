package tint

import "testing"

func TestFg256(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(Fg256(208))
	got := hc.Sprint("orange")
	want := "\033[38;5;208morange\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestBg256(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(Bg256(21))
	got := hc.Sprint("blue bg")
	want := "\033[48;5;21mblue bg\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFgRGB(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(FgRGB(255, 128, 0))
	got := hc.Sprint("rgb")
	want := "\033[38;2;255;128;0mrgb\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestBgRGB(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(BgRGB(0, 0, 128))
	got := hc.Sprint("dark blue")
	want := "\033[48;2;0;0;128mdark blue\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHiColorMixed(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(Bold, Fg256(208), BgRGB(0, 0, 128))
	got := hc.Sprint("mixed")
	want := "\033[1;38;5;208;48;2;0;0;128mmixed\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHiColorSprintf(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(FgRGB(255, 0, 0))
	got := hc.Sprintf("val=%d", 42)
	want := "\033[38;2;255;0;0mval=42\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHiColorSprintln(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(Fg256(10))
	got := hc.Sprintln("line")
	want := "\033[38;5;10mline\n\033[0m"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHiColorNoColor(t *testing.T) {
	orig := NoColor
	NoColor = true
	defer func() { NoColor = orig }()

	hc := HiColor(FgRGB(255, 0, 0))
	got := hc.Sprint("plain")
	if got != "plain" {
		t.Fatalf("expected plain text, got %q", got)
	}
}

func TestHiColorDisableEnable(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(Fg256(100))
	hc.DisableColor()
	got := hc.Sprint("off")
	if got != "off" {
		t.Fatalf("expected plain after DisableColor, got %q", got)
	}

	hc.EnableColor()
	got = hc.Sprint("on")
	want := "\033[38;5;100mon\033[0m"
	if got != want {
		t.Fatalf("expected colored after EnableColor, got %q, want %q", got, want)
	}
}

func TestFg256Boundary(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	// Test boundary values 0 and 255.
	hc0 := HiColor(Fg256(0))
	got := hc0.Sprint("x")
	want := "\033[38;5;0mx\033[0m"
	if got != want {
		t.Fatalf("Fg256(0): got %q, want %q", got, want)
	}

	hc255 := HiColor(Fg256(255))
	got = hc255.Sprint("x")
	want = "\033[38;5;255mx\033[0m"
	if got != want {
		t.Fatalf("Fg256(255): got %q, want %q", got, want)
	}
}
