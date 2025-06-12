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

func TestShouldPrintFizzForMultiplesOf3ButNotMultiplesOf5(t *testing.T) {
	result := fizzBuzz()

	for i := 0; i < 100; i++ {
		if (i+1)%3 != 0 || (i+1)%5 == 0 {
			continue
		}

		if result[i] != "Fizz" {
			t.Errorf("For count %d expected Fizz but got %s", i, result[i])
		}
	}
}

func TestShouldPrintBuzzForMultiplesOf5(t *testing.T) {
	result := fizzBuzz()

	for i := 0; i < 100; i++ {
		if (i+1)%5 != 0 || (i+1)%3 == 0 {
			continue
		}

		if result[i] != "Buzz" {
			t.Errorf("For count %d expected Buzz but got %s", i, result[i])
		}
	}
}

func TestShouldPrintFizzBuzzForMultiplesOf3And5(t *testing.T) {
	result := fizzBuzz()

	for i := 0; i < 100; i++ {
		if (i+1)%3 != 0 || (i+1)%5 != 0 {
			continue
		}

		if result[i] != "FizzBuzz" {
			t.Errorf("For count %d expected FizzBuzz but got %s", i+1, result[i])
		}
	}
}
