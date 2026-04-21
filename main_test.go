package main

import (
	"bytes"
	"testing"
)

// TestRun verifies that run() produces the expected output.
func TestRun(t *testing.T) {
	old := outputWriter
	defer func() { outputWriter = old }()

	var buf bytes.Buffer
	outputWriter = &buf

	err := run()
	if err != nil {
		t.Fatalf("run() returned unexpected error: %v", err)
	}

	got := buf.String()
	expected := "Hello, Agent\n"
	if got != expected {
		t.Errorf("output = %q, want %q", got, expected)
	}
}
