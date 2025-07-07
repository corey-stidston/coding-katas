package main

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		player1Points int
		player2Points int
		expectedScore Score
	}{
		{
			player1Points: 4,
			player2Points: 0,
			expectedScore: Player1Wins,
		},
		{
			player1Points: 4,
			player2Points: 1,
			expectedScore: Player1Wins,
		},
		{
			player1Points: 4,
			player2Points: 2,
			expectedScore: Player1Wins,
		},
		{
			player1Points: 6,
			player2Points: 4,
			expectedScore: Player1Wins,
		},
		{
			player1Points: 0,
			player2Points: 4,
			expectedScore: Player2Wins,
		},
		{
			player1Points: 1,
			player2Points: 4,
			expectedScore: Player2Wins,
		},
		{
			player1Points: 2,
			player2Points: 4,
			expectedScore: Player2Wins,
		},
		{
			player1Points: 4,
			player2Points: 6,
			expectedScore: Player2Wins,
		},
		{
			player1Points: 3,
			player2Points: 3,
			expectedScore: Deuce,
		},
		{
			player1Points: 4,
			player2Points: 4,
			expectedScore: Deuce,
		},
		{
			player1Points: 0,
			player2Points: 1,
			expectedScore: LoveFifteen,
		},
		{
			player1Points: 1,
			player2Points: 0,
			expectedScore: FifteenLove,
		},
		{
			player1Points: 0,
			player2Points: 3,
			expectedScore: LoveForty,
		},
		{
			player1Points: 1,
			player2Points: 2,
			expectedScore: FifteenThirty,
		},
		{
			player1Points: 1,
			player2Points: 3,
			expectedScore: FifteenForty,
		},
		{
			player1Points: 3,
			player2Points: 4,
			expectedScore: Player2Advantage,
		},
		{
			player1Points: 10,
			player2Points: 11,
			expectedScore: Player2Advantage,
		},
		{
			player1Points: 117,
			player2Points: 116,
			expectedScore: Player1Advantage,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			game := TennisGame()
			highestScore := max(tt.player1Points, tt.player2Points)

			for i := 0; i < highestScore; i++ {
				if i < tt.player1Points {
					game.WonPoint(1)
				}
				if i < tt.player2Points {
					game.WonPoint(2)
				}
			}

			score := game.GetScore()

			if score != tt.expectedScore {
				t.Errorf("Expected %s but got %s", tt.expectedScore, score)
			}
		})
	}
}
