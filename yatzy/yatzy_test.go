package main

import (
	"fmt"
	"testing"
)

func TestYatzy(t *testing.T) {
	tests := []struct {
		dice [6]int
		category category
		expectedScore int
	}{
		{
			dice: [6]int{1,1,1,1,1,1},
			category: yatzy,
			expectedScore: 6,
		},
		{
			dice: [6]int{1,2,3,4,5,6},
			category: chance,
			expectedScore: 21,
		},
		{
			dice: [6]int{1,2,1,4,5,1},
			category: ones,
			expectedScore: 3,
		},
		{
			dice: [6]int{1,2,2,4,5,1},
			category: twos,
			expectedScore: 4,
		},
		{
			dice: [6]int{1,3,3,4,5,1},
			category: threes,
			expectedScore: 6,
		},
		{
			dice: [6]int{1,4,2,4,5,1},
			category: fours,
			expectedScore: 8,
		},
		{
			dice: [6]int{5,2,2,4,5,1},
			category: fives,
			expectedScore: 10,
		},
		{
			dice: [6]int{6,2,2,4,6,1},
			category: sixes,
			expectedScore: 12,
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

