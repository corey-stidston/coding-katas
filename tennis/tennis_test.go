package main

import "testing"

func TestGame(t *testing.T) {
	expectedScore1 := "Lose"
	expectedScore2 := "Win"
	score1, score2 := score(0, 4)

	if score1 != expectedScore1 || score2 != expectedScore2 {
		t.Errorf("Expected %s - %s but got %s - %s", expectedScore1, expectedScore2, score1, score2)
	}
}
