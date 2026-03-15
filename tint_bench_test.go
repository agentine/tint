package tint

import "testing"

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(FgRed, Bold)
	}
}

func BenchmarkSprint(b *testing.B) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgRed, Bold)
	for i := 0; i < b.N; i++ {
		c.Sprint("hello world")
	}
}

func BenchmarkSprintf(b *testing.B) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgGreen)
	for i := 0; i < b.N; i++ {
		c.Sprintf("count: %d", 42)
	}
}

func BenchmarkSprintNoColor(b *testing.B) {
	orig := NoColor
	NoColor = true
	defer func() { NoColor = orig }()

	c := New(FgRed, Bold)
	for i := 0; i < b.N; i++ {
		c.Sprint("hello world")
	}
}

func BenchmarkRedString(b *testing.B) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	for i := 0; i < b.N; i++ {
		RedString("error: %s", "something failed")
	}
}

func BenchmarkHiColorSprint(b *testing.B) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(Fg256(208))
	for i := 0; i < b.N; i++ {
		hc.Sprint("hello")
	}
}

func BenchmarkHiColorRGB(b *testing.B) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	hc := HiColor(FgRGB(255, 128, 0))
	for i := 0; i < b.N; i++ {
		hc.Sprint("hello")
	}
}

func BenchmarkColorCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCachedColor(FgRed)
	}
}
