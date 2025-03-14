package main

import (
	"fmt"
	"testing"
)

const errorFormat = "Expected %t but got %t"

func TestReturnsTrueForTypicalLeapYear(t *testing.T) {
	expected := true
	result := isLeapYear(1996)

	if result != expected {
		t.Errorf(errorFormat, expected, result)
	}
}

func TestReturnsFalseForTypicalCommonYears(t *testing.T) {
	expected := false
	tests := []int{1903, 2001}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			if got := isLeapYear(tt); got != expected {
				t.Errorf(errorFormat+" for year %d", expected, got, tt)
			}
		})
	}
}

func TestReturnsFalseForAtypicalCommonYears(t *testing.T) {
	expected := false
	tests := []int{1900, 2300}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			if got := isLeapYear(tt); got != expected {
				t.Errorf(errorFormat+" for year %d", expected, got, tt)
			}
		})
	}
}

func TestReturnsTrueForAtypicalLeapYears(t *testing.T) {
	expected := true
	tests := []int{1600, 2000}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			if got := isLeapYear(tt); got != expected {
				t.Errorf(errorFormat+" for year %d", expected, got, tt)
			}
		})
	}
}
