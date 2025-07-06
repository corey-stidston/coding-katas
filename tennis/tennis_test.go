package main

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		player1Points int
		player2Points int
		player1ExpectedScore Score
		player2ExpectedScore Score
	}{
		{
			player1Points: 4,
			player2Points: 0,
			player1ExpectedScore: Win,
			player2ExpectedScore: Lose,
		},
		{
			player1Points: 4,
			player2Points: 1,
			player1ExpectedScore: Win,
			player2ExpectedScore: Lose,
		},
		{
			player1Points: 4,
			player2Points: 2,
			player1ExpectedScore: Win,
			player2ExpectedScore: Lose,
		},
		{
			player1Points: 6,
			player2Points: 4,
			player1ExpectedScore: Win,
			player2ExpectedScore: Lose,
		},
		{
			player1Points: 0,
			player2Points: 4,
			player1ExpectedScore: Lose,
			player2ExpectedScore: Win,
		},
		{
			player1Points: 1,
			player2Points: 4,
			player1ExpectedScore: Lose,
			player2ExpectedScore: Win,
		},
		{
			player1Points: 2,
			player2Points: 4,
			player1ExpectedScore: Lose,
			player2ExpectedScore: Win,
		},
		{
			player1Points: 4,
			player2Points: 6,
			player1ExpectedScore: Lose,
			player2ExpectedScore: Win,
		},
		{
			player1Points: 3,
			player2Points: 3,
			player1ExpectedScore: Deuce,
			player2ExpectedScore: Deuce,
		},
		{
			player1Points: 4,
			player2Points: 4,
			player1ExpectedScore: Deuce,
			player2ExpectedScore: Deuce,
		},
		{
			player1Points: 0,
			player2Points: 1,
			player1ExpectedScore: Love,
			player2ExpectedScore: Fifteen,
		},
		{
			player1Points: 1,
			player2Points: 0,
			player1ExpectedScore: Fifteen,
			player2ExpectedScore: Love,
		},
		{
			player1Points: 0,
			player2Points: 3,
			player1ExpectedScore: Love,
			player2ExpectedScore: Forty,
		},
		{
			player1Points: 1,
			player2Points: 3,
			player1ExpectedScore: Fifteen,
			player2ExpectedScore: Forty,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			score1, score2 := score(tt.player1Points, tt.player2Points)

			if score1 != tt.player1ExpectedScore && score2 != tt.player2ExpectedScore {
				t.Errorf("Expected %s - %s but got %s - %s", tt.player1ExpectedScore, score1, tt.player2ExpectedScore, score2)
			}
		})
	}
}
