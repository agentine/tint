package tint

import (
	"bytes"
	"sync"
	"testing"
)

func TestConcurrentSprint(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgRed, Bold)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			got := c.Sprint("hello")
			want := "\033[31;1mhello\033[0m"
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}()
	}
	wg.Wait()
}

func TestConcurrentFprint(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	c := New(FgGreen)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var buf bytes.Buffer
			_, err := c.Fprint(&buf, "test")
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestConcurrentConvenienceFunctions(t *testing.T) {
	orig := NoColor
	NoColor = false
	defer func() { NoColor = orig }()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = RedString("err")
			_ = GreenString("ok")
			_ = BlueString("info")
		}()
	}
	wg.Wait()
}

func TestConcurrentColorCache(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = getCachedColor(FgRed)
			_ = getCachedColor(FgBlue)
			_ = getCachedColor(FgGreen)
		}()
	}
	wg.Wait()
}
