//go:build !plan9

package tint

import (
	"os"
	"testing"
)

func TestIsTerminal(t *testing.T) {
	// In test environments, stdout may or may not be a terminal.
	// Just verify the function runs without panic.
	_ = IsTerminal(os.Stdout.Fd())
	_ = IsTerminal(os.Stderr.Fd())
}

func TestIsTerminalInvalidFd(t *testing.T) {
	// Invalid fd should return false, not panic.
	if IsTerminal(999999) {
		t.Fatal("expected false for invalid fd")
	}
}

func TestIsTerminalPipe(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	if IsTerminal(r.Fd()) {
		t.Fatal("pipe read end should not be a terminal")
	}
	if IsTerminal(w.Fd()) {
		t.Fatal("pipe write end should not be a terminal")
	}
}

func TestIsStdoutTerminal(t *testing.T) {
	// Just verify it runs without panic.
	_ = isStdoutTerminal()
}

func TestIsStderrTerminal(t *testing.T) {
	_ = isStderrTerminal()
}
