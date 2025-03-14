
package main

import "testing"

func TestGetHelloMessage(t *testing.T) {
	expected := "Hello, World!"
	result := getHelloMessage()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
