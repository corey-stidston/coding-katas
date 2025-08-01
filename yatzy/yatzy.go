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
		result = game.calculateYatzy()
	case chance:
		result = game.calculateChance()
	case ones:
		result = game.sumOfMatchingDie(1)
	case twos:
		result = game.sumOfMatchingDie(2)
	case threes:
		result = game.sumOfMatchingDie(3)
	case fours:
		result = game.sumOfMatchingDie(4)
	case fives:
		result = game.sumOfMatchingDie(5)
	case sixes:
		result = game.sumOfMatchingDie(6)
	case pair:
		result = game.sumOfAKind(2)
	case three_of_a_kind:
		result = game.sumOfAKind(3)
	case four_of_a_kind:
		result = game.sumOfAKind(4)
	case small_straight:
		result = game.exactMatch(smallStraight)
	case large_straight:
		result = game.exactMatch(largeStraight)
	}

	return result
}

func (game *yatzyGame) exactMatch(straightType dice) int {
	if game.diceRoll == straightType {
		return game.sumOfDie()
	}
	return 0
}

func (game *yatzyGame) calculateYatzy() int {
	var firstDie = game.diceRoll[0]
	for i := 1; i < len(game.diceRoll); i++ {
		if game.diceRoll[i] != firstDie {
			return 0
		}
	}
	return firstDie * number_of_die
}

func (game *yatzyGame) calculateChance() int {
	var sum int
	for _, die := range game.diceRoll {
		sum += die
	}
	return sum
}

func (game *yatzyGame) sumOfAKind(kind int) int {
	var countOfDie dice
	for i := 0; i < len(game.diceRoll); i++ {
		countOfDie[game.diceRoll[i]-1] += 1
	}

	for i := len(game.diceRoll) - 1; i >= 0; i-- {
		if countOfDie[i] == kind {
			return (i + 1) * kind
		}
	}
	return 0
}

func (game *yatzyGame) sumOfDie() int {
	return reduce(game.diceRoll, func(acc int, current int) int {
		return acc + current
	})
}

func (game *yatzyGame) sumOfMatchingDie(matchingDie int) int {
	return reduce(game.diceRoll, func(acc int, current int) int {
		var value int
		if current == matchingDie {
			value = acc + current
		} else {
			value = acc
		}
		return value
	})
}

func reduce(dice dice, f func(int, int) int) int {
	accumulator := 0
	for _, value := range dice {
		accumulator = f(accumulator, value)
	}
	return accumulator
}

func main() {
	//
}
