package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
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

func TestNoBankStatements(t *testing.T) {
	bankAccount := BankAccount()

	output := captureOutput(func() {
		bankAccount.printStatement()
	})
	if output != `Date       || Amount || Balance
` {
		t.Errorf("Expected statement header in output, got: %q", output)
	}
}

// Single deposit
func TestSingleDepositStatement(t *testing.T) {
	bankAccount := BankAccount()
	bankAccount.deposit(100)

	output := captureOutput(func() {
		bankAccount.printStatement()
	})

	today := time.Now().Format(time.DateOnly)

	expectedOutput := fmt.Sprintf(`Date       || Amount || Balance
%s || 100    || 100
`, today)

	if output != expectedOutput {
		t.Errorf("Expected %q, got: %q", expectedOutput, output)
	}
}

func TestMultipleDepositStatement(t *testing.T) {
	bankAccount := BankAccount()
	bankAccount.deposit(100)
	bankAccount.deposit(340)

	output := captureOutput(func() {
		bankAccount.printStatement()
	})

	today := time.Now().Format(time.DateOnly)

	expectedOutput := fmt.Sprintf(`Date       || Amount || Balance
%s || 100    || 100
%s || 340    || 440
`, today, today)

	if output != expectedOutput {
		t.Errorf("Expected %q, got: %q", expectedOutput, output)
	}
}

// Single deposit, single withdrawal
// Negative balance not allowed
