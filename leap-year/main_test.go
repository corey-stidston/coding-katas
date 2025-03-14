package main

import "testing"

func TestReturnsTrueForTypicalLeapYear(t *testing.T) {
	expected := true
	result := isLeapYear(1996)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func TestReturnsFalseForTypicalNonLeapYear(t *testing.T) {
	expected := false
	result := isLeapYear(2001)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}
