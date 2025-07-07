package main

import "testing"

func TestYatzy(t *testing.T) {
	expected := 6
	result := YatzyScore([6]int{1,1,1,1,1,1}, yatzy)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

