package main

import "testing"

func TestShouldPrintNumbersFrom1To100(t *testing.T) {
	expected := `1
2
3
4
5
6
7
8
9
10`
	result := fizzBuzz()

	if expected != result {
		t.Errorf("Expected %s but got %s", expected, result)
	}

}
