// Package main prints a greeting to stdout.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// run contains the application logic and returns an error for any failures.
func run() error {
	fmt.Fprintln(outputWriter, "Hello, Agent")
	return nil
}

// outputWriter is the canonical target for user-facing output.
// Replaceable in tests via os.Std/out reassignment.
var outputWriter io.Writer = os.Stdout
