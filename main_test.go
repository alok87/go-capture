package capture

import (
	"fmt"
	"os"
	"testing"
)

func TestStdout(t *testing.T) {
	out := Stdout(func() {
		fmt.Fprint(os.Stdout, "stdout")
	})

	if out != "stdout" {
		t.Errorf("expected: stdout, got: %v", out)
	}
}

func TestStderr(t *testing.T) {
	out := Stderr(func() {
		fmt.Fprint(os.Stderr, "stderr")
	})

	if out != "stderr" {
		t.Errorf("expected: stderr, got: %v", out)
	}
}

func TestOutput(t *testing.T) {
	out := Output(func() {
		fmt.Fprint(os.Stdout, "stdout")
		fmt.Fprint(os.Stderr, "stderr")
	})

	if out != "stdoutstderr" {
		t.Errorf("expected: stdoutstderr, got: %v", out)
	}
}
