package main

import (
	"testing"
)

func TestShouldPrint100Items(t *testing.T) {
	expectedCount := 100
	result := fizzBuzz()

	if len(result) != expectedCount {
		t.Errorf("Expected length to be %d but got %d", expectedCount, len(result))
	}
}

func TestShouldPrintFizzForMultiplesOf3(t *testing.T) {
	result := fizzBuzz()

	for i := 0; i < 100; i++ {
		if (i+1)%3 != 0 {
			continue
		}

		if result[i] != "Fizz" {
			t.Errorf("Expected Fizz but got %s", result[i])
		}
	}
}
