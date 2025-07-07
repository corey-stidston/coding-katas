/*
Tennis has a rather quirky scoring system, and to newcomers it can be a little difficult to keep track of.
The tennis society has contracted you to build a scoreboard to display the current score during tennis games.

You can read more about Tennis scores on wikipedia which is summarized below:

1. A game is won by the first p to have won at least four points in total and at least two points more than the opponent.
2. The running score of each game is described in a manner peculiar to tennis: scores from zero to three points are described as
“Love”, “Fifteen”, “Thirty”, and “Forty” respectively.
3. If at least three points have been scored by each p, and the scores are equal, the score is “Deuce”.
4. If at least three points have been scored by each side and a p has one more point than his opponent,
the score of the game is “Advantage” for the p in the lead.
You need only report the score for the current game. Sets and Matches are out of scope.
*/

package main

type tennisGame struct {
	p1Points int
	p2Points int
}

type Score string

const (
	Player1Wins      Score = "Player 1 Wins"
	Player2Wins      Score = "Player 2 Wins"
	Player1Advantage Score = "Player 1 Advantage"
	Player2Advantage Score = "Player 2 Advantage"

	LoveFifteen   Score = "Love - Fifteen"
	LoveThirty    Score = "Love - Thirty"
	LoveForty     Score = "Love - Forty"
	FifteenLove   Score = "Fifteen - Love"
	ThirtyLove    Score = "Thirty - Love"
	FortyLove     Score = "Forty - Love"
	FifteenThirty Score = "Fifteen - Thirty"
	ThirtyFifteen Score = "Thirty - Fifteen"
	ThirtyForty   Score = "Thirty - Forty"
	FortyThirty   Score = "Forty - Thirty"
	FifteenForty  Score = "Fifteen - Forty"
	FortyFifteen  Score = "Forty - Fifteen"

	LoveAll    Score = "Love All"
	FifteenAll Score = "Fifteen All"
	ThirtyAll  Score = "Thirty All"
	Deuce      Score = "Deuce"
	Fifteen    Score = "Fifteen"
	Thirty     Score = "Thirty"
	Forty      Score = "Forty"
)

var scoreMap = [4][4]Score{
	{LoveAll, LoveFifteen, LoveThirty, LoveForty},
	{FifteenLove, FifteenAll, FifteenThirty, FifteenForty},
	{ThirtyLove, ThirtyFifteen, ThirtyAll, ThirtyForty},
	{FortyLove, FortyFifteen, FortyThirty, Deuce},
}

func TennisGame() *tennisGame {
	return &tennisGame{
		p1Points: 0,
		p2Points: 0,
	}
}

func (game *tennisGame) WonPoint(playerNumber int) {
	if playerNumber == 1 {
		game.p1Points += 1
	} else {
		game.p2Points += 1
	}
}

func (game *tennisGame) GetScore() Score {
	var result Score

	if game.p1Points >= 4 && game.p1Points-game.p2Points >= 2 {
		result = Player1Wins
	} else if game.p2Points >= 4 && game.p2Points-game.p1Points >= 2 {
		result = Player2Wins
	} else if game.p1Points == game.p2Points && game.p1Points >= 3 && game.p2Points >= 3 {
		result = Deuce
	} else if game.p1Points >= 3 && game.p2Points >= 3 {
		if game.p1Points > game.p2Points {
			result = Player1Advantage
		} else {
			result = Player2Advantage
		}
	} else {
		result = scoreMap[game.p1Points][game.p2Points]
	}
	return result
}

func main() {
	//
}
