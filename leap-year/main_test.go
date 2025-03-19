package main

import (
	"fmt"
	"testing"
)

func TestReturnsCorrectValueWhenInputIsValid(t *testing.T) {
	tests := []struct {
		testName string
		year     int
		expected bool
	}{
		{
			"Typical leap year", 1996, true,
		},
		{
			"Typical common year", 1903, false,
		},
		{
			"Typical common year", 2001, false,
		},
		{
			"Atypical common year", 1900, false,
		},
		{
			"Atypical common year", 2300, false,
		},
		{
			"Atypical leap year", 1600, true,
		},
		{
			"Atypical leap year", 2000, true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			if got, _ := isLeapYear(tt.year); got != tt.expected {
				t.Errorf(tt.testName+": Expected %t but got %t for year %d", tt.expected, got, tt.year)
			}
		})
	}
}

func TestReturnsErrorWhenInputIsInvalid(t *testing.T) {
	_, error := isLeapYear(-1000)

	if error == nil {
		t.Errorf("Expected error to be returned.")
	}
}
