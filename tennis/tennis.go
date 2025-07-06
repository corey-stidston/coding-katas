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

type TennisGame struct {
	p1Points int
	p2Points int
}

type Score string

const (
	Win       Score = "Win"
	Lose      Score = "Lose"
	Deuce     Score = "Deuce"
	Love      Score = "Love"
	Fifteen   Score = "Fifteen"
	Thirty    Score = "Thirty"
	Forty     Score = "Forty"
	Advantage Score = "Advantage"
)

var scoreMap = [4]Score{Love, Fifteen, Thirty, Forty}

func (game *TennisGame) GetScore() (Score, Score) {
	var p1Score Score
	var p2Score Score

	if game.p1Points >= 4 && game.p1Points-game.p2Points >= 2 {
		p1Score = Win
		p2Score = Lose
	} else if game.p2Points >= 4 && game.p2Points-game.p1Points >= 2 {
		p1Score = Lose
		p2Score = Win
	} else if game.p1Points == game.p2Points && game.p1Points >= 3 && game.p2Points >= 3 {
		p1Score = Deuce
		p2Score = Deuce
	} else if game.p1Points >= 3 && game.p2Points >= 3 {
		if game.p1Points > game.p2Points {
			p1Score = Advantage
			p2Score = scoreMap[game.p2Points]
		} else {
			p1Score = scoreMap[game.p1Points]
			p2Score = Advantage
		}
	} else {
		p1Score = scoreMap[game.p1Points]
		p2Score = scoreMap[game.p2Points]
	}
	return p1Score, p2Score
}

func main() {
	//
}
