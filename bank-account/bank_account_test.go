package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Helper to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old
	return buf.String()
}

func TestPrintStatementEmpty(t *testing.T) {
	output := captureOutput(func() {
		PrintStatement()
	})
	if output != "Date       || Amount || Balance" {
		t.Errorf("Expected statement header in output, got: %q", output)
	}
}
