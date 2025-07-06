package main

import (
	"fmt"
	"testing"
)

func TestWinLose(t *testing.T) {
	tests := []struct {
		player1Points int
		player2Points int
		expectedScore string
	}{
		{
			player1Points: 4,
			player2Points: 0,
			expectedScore: "Win - Lose",
		},
		{
			player1Points: 4,
			player2Points: 1,
			expectedScore: "Win - Lose",
		},
		{
			player1Points: 4,
			player2Points: 2,
			expectedScore: "Win - Lose",
		},
		{
			player1Points: 0,
			player2Points: 4,
			expectedScore: "Lose - Win",
		},
		{
			player1Points: 1,
			player2Points: 4,
			expectedScore: "Lose - Win",
		},
		{
			player1Points: 2,
			player2Points: 4,
			expectedScore: "Lose - Win",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			score := score(tt.player1Points, tt.player2Points)

			if score != tt.expectedScore {
				t.Errorf("Expected %s but got %s", tt.expectedScore, score)
			}
		})
	}
}
