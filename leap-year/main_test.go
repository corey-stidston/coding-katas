package main

import "testing"

func TestReturnsTrueForLeapYear(t *testing.T) {
	expected := true
	result := isLeapYear(1996)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}
