package main

import (
	"fmt"
	"testing"
)

func TestWinLose(t *testing.T) {
	tests := []struct {
		player1Points int
		player2Points int
	}{
		{
			player1Points: 4,
			player2Points: 0,
		},
		{
			player1Points: 4,
			player2Points: 1,
		},
		{
			player1Points: 4,
			player2Points: 2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			score1, score2 := score(tt.player1Points, tt.player2Points)

			if score1 != "Win" || score2 != "Lose" {
				t.Errorf("Expected %s - Win but got %s - Lose", score1, score2)
			}
		})
	}
}

func TestLoseWin(t *testing.T) {
	tests := []struct {
		player1Points int
		player2Points int
	}{
		{
			player1Points: 0,
			player2Points: 4,
		},
		{
			player1Points: 1,
			player2Points: 4,
		},
		{
			player1Points: 2,
			player2Points: 4,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			score1, score2 := score(tt.player1Points, tt.player2Points)

			if score1 != "Lose" || score2 != "Win" {
				t.Errorf("Expected %s - Lose but got %s - Win", score1, score2)
			}
		})
	}
}