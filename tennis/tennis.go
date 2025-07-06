/*
Tennis has a rather quirky scoring system, and to newcomers it can be a little difficult to keep track of.
The tennis society has contracted you to build a scoreboard to display the current score during tennis games.

You can read more about Tennis scores on wikipedia which is summarized below:

1. A game is won by the first player to have won at least four points in total and at least two points more than the opponent.
2. The running score of each game is described in a manner peculiar to tennis: scores from zero to three points are described as 
“Love”, “Fifteen”, “Thirty”, and “Forty” respectively.
3. If at least three points have been scored by each player, and the scores are equal, the score is “Deuce”.
4. If at least three points have been scored by each side and a player has one more point than his opponent, 
the score of the game is “Advantage” for the player in the lead.
You need only report the score for the current game. Sets and Matches are out of scope.
*/

package main

type Score int

const (
	WinLose Score = iota
	LoseWin
	Deuce
)

func score(player1 int, player2 int) (Score) {
	if (player1 >= 4 && player1 - player2 >= 2) {
		return WinLose
	} else if(player2 >= 4 && player2 - player1 >= 2) {
		return LoseWin
	}
	return LoseWin
}

func main() {
}
