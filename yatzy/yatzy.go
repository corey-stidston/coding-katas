/*
The game of Yatzy is a simple dice game. Each player rolls five six-sided dice.
They can re-roll some or all of the dice up to three times (including the original roll).

The player then places the roll in a category, such as ones, twos, fives, pair, two pairs etc (see the rules below).
If the roll is compatible with the category, the player gets a score for the roll according to the rules.
If the roll is not compatible with the category, the player scores zero for the roll.

For example, suppose a player scores 5,6,5,5,2 in the fives category they would score 15 (three fives).
The score for that go is then added to their total and the category cannot be used again in the remaining goes for that game.
A full game consists of one go for each category. Thus, for their last go in a game, a player must choose their only remaining category.

Your task is to score a GIVEN roll in a GIVEN category. You do NOT have to program the random dice rolling.
The game is NOT played by letting the computer choose the highest scoring category for a given roll.

Yatzy categories -

Chance: The player scores the sum of all dice, no matter what they read.

Yatzy: If all dice have the same number, the player scores 50 points.

Ones, Twos, Threes, Fours, Fives, Sixes: The player scores the sum of the dice that reads one, two, three, four, five or six, respectively.

Pair: The player scores the sum of the two highest matching dice.

Two pairs: If there are two pairs of dice with the same number, the player scores the sum of these dice.

Three of a kind: If there are three dice with the same number, the player scores the sum of these dice.

Four of a kind: If there are four dice with the same number, the player scores the sum of these dice.

Small straight: dice (1,2,3,4,5) the player scores the sum of all dice (15)

Large straight: dice (2,3,4,5,6) the player scores the sum of all dice (20)

Full house: If the dice are two of a kind and three of a kind, the player scores the sum of all the dice.
*/

package main

type category string

const (
	number_of_die   int      = 5
	unknown         category = "unknown"
	chance          category = "chance"
	yatzy           category = "yatzy"
	ones            category = "ones"
	twos            category = "twos"
	threes          category = "threes"
	fours           category = "fours"
	fives           category = "fives"
	sixes           category = "sixes"
	pair            category = "pair"
	three_of_a_kind category = "three_of_a_kind"
	four_of_a_kind  category = "four_of_a_kind"
	small_straight  category = "small_straight"
	large_straight  category = "large_straight"
	two_pairs       category = "two_pairs"
	full_house      category = "full_house"
)

type dice [number_of_die]int

type yatzyGame struct {
	diceRoll dice
}

var smallStraight = dice{1,2,3,4,5}
var largeStraight = dice{2,3,4,5,6}

func YatzyGame() *yatzyGame {
	return &yatzyGame{}
}

func (game *yatzyGame) Roll(dice dice) {
	game.diceRoll = dice
}

func (game *yatzyGame) GetScore(category category) int {
	var result int

	switch category {
	case yatzy:
		result = game.yatzy()
	case chance:
		result = game.chance()
	case ones, twos, threes, fours, fives, sixes:
		result = game.sumOfMatchingDie(category)
	case pair:
		result = game.sumOfAKind(2)
	case three_of_a_kind:
		result = game.sumOfAKind(3)
	case four_of_a_kind:
		result = game.sumOfAKind(4)
	case small_straight:
		result = game.straight(smallStraight)
	case large_straight:
		result = game.straight(largeStraight)
	case two_pairs:
		result = game.twoPairs()
	}

	return result
}

func (game *yatzyGame) twoPairs() int {
	dieCounts := game.getDieCounts()

	pairs := 0
	sumOfPairs := 0
	for index, dieCount := range dieCounts {
		if dieCount == 2 {
			pairs++
			sumOfPairs += 2 * (index + 1)
		}
	}

	if pairs == 2 {
		return sumOfPairs
	}
	return 0
}

func (game *yatzyGame) straight(straightType dice) int {
	if game.diceRoll == straightType {
		return game.sumOfDie()
	}
	return 0
}

func (game *yatzyGame) yatzy() int {
	var dieCounts = game.getDieCounts()
	for _, die := range dieCounts {
		if die == 5 {
			return 50
		}
	}
	return 0
}

func (game *yatzyGame) chance() int {
	var sum int
	for _, die := range game.diceRoll {
		sum += die
	}
	return sum
}

func (game *yatzyGame) getDieCounts() [6]int {
	var dieCounts [6]int
	for _, die := range game.diceRoll {
		dieCounts[die-1]++
	}
	return dieCounts
}

func (game *yatzyGame) sumOfAKind(kind int) int {
	var dieCounts = game.getDieCounts()

	for i := len(game.diceRoll) - 1; i >= 0; i-- {
		if dieCounts[i] == kind {
			return (i + 1) * kind
		}
	}
	return 0
}

func (game *yatzyGame) sumOfDie() int {
	sum := 0
    for _, die := range game.diceRoll {
        sum += die
    }
    return sum
}

func (game *yatzyGame) sumOfMatchingDie(category category) int {
	categoryMap := map[string]int {
		string(ones): 1,
		string(twos): 2,
		string(threes): 3,
		string(fours): 4,
		string(fives): 5,
		string(sixes): 6,
	}
	sum := 0
    for _, die := range game.diceRoll {
        if die == categoryMap[string(category)] {
            sum += die
        }
    }
    return sum
}

func main() {
	//
}
