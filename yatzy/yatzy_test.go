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
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			result := YatzyScore(tt.dice, tt.category)

			if result != tt.expectedScore {
				t.Errorf("Expected %d but got %d", tt.expectedScore, result)
			}
		})
	}
}

