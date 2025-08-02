package main

import (
	"fmt"
	"testing"
)

func TestYatzy(t *testing.T) {
	tests := []struct {
		dice          dice
		category      category
		expectedScore int
	}{
		{
			dice:          dice{1, 1, 1, 1, 1},
			category:      yatzy,
			expectedScore: 50,
		},
		{
			dice:          dice{2, 1, 1, 1, 1},
			category:      yatzy,
			expectedScore: 0,
		},
		{
			dice:          dice{1, 2, 3, 4, 5},
			category:      chance,
			expectedScore: 15,
		},
		{
			dice:          dice{1, 2, 1, 4, 1},
			category:      ones,
			expectedScore: 3,
		},
		{
			dice:          dice{1, 2, 2, 4, 5},
			category:      twos,
			expectedScore: 4,
		},
		{
			dice:          dice{1, 3, 3, 4, 5},
			category:      threes,
			expectedScore: 6,
		},
		{
			dice:          dice{1, 4, 2, 4, 5},
			category:      fours,
			expectedScore: 8,
		},
		{
			dice:          dice{5, 2, 2, 4, 5},
			category:      fives,
			expectedScore: 10,
		},
		{
			dice:          dice{6, 2, 2, 4, 6},
			category:      sixes,
			expectedScore: 12,
		},
		{ // zero pairs
			dice:          dice{1, 2, 3, 4, 5},
			category:      pair,
			expectedScore: 0,
		},
		{ // one pair
			dice:          dice{1, 1, 3, 4, 5},
			category:      pair,
			expectedScore: 2,
		},
		{ // two pairs - take highest pair
			dice:          dice{1, 1, 3, 3, 5},
			category:      pair,
			expectedScore: 6,
		},
		{ // zero three of a kind
			dice:          dice{1, 2, 3, 4, 5},
			category:      three_of_a_kind,
			expectedScore: 0,
		},
		{ // one three of a kind
			dice:          dice{2, 2, 2, 4, 5},
			category:      three_of_a_kind,
			expectedScore: 6,
		},
		{ // zero four of a kind
			dice:          dice{2, 2, 2, 4, 4},
			category:      four_of_a_kind,
			expectedScore: 0,
		},
		{ // one four of a kind
			dice:          dice{5, 2, 5, 5, 5},
			category:      four_of_a_kind,
			expectedScore: 20,
		},
		{
			dice:          dice{1, 2, 3, 4, 5},
			category:      small_straight,
			expectedScore: 15,
		},
		{ // invalid small straight
			dice:          dice{5, 2, 3, 4, 5},
			category:      small_straight,
			expectedScore: 0,
		},
		{
			dice:          dice{2, 3, 4, 5, 6},
			category:      large_straight,
			expectedScore: 20,
		},
		{ // invalid large straight
			dice:          dice{5, 3, 3, 4, 5},
			category:      large_straight,
			expectedScore: 0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			game := YatzyGame()
			game.Roll(tt.dice)
			result := game.GetScore(tt.category)

			if result != tt.expectedScore {
				t.Errorf("Expected %d but got %d", tt.expectedScore, result)
			}
		})
	}
}
