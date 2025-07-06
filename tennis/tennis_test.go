package main

import (
	"fmt"
	"testing"
)

func TestWinLose(t *testing.T) {
	tests := []struct {
		player1Points int
		player2Points int
		expectedScore Score
	}{
		{
			player1Points: 4,
			player2Points: 0,
			expectedScore: WinLose,
		},
		{
			player1Points: 4,
			player2Points: 1,
			expectedScore: WinLose,
		},
		{
			player1Points: 4,
			player2Points: 2,
			expectedScore: WinLose,
		},
		{
			player1Points: 0,
			player2Points: 4,
			expectedScore: LoseWin,
		},
		{
			player1Points: 1,
			player2Points: 4,
			expectedScore: LoseWin,
		},
		{
			player1Points: 2,
			player2Points: 4,
			expectedScore: LoseWin,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			score := score(tt.player1Points, tt.player2Points)

			if score != tt.expectedScore {
				t.Errorf("Expected %d but got %d", tt.expectedScore, score)
			}
		})
	}
}
