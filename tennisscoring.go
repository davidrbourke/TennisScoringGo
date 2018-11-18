package tennisscoring

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

// PlayMatch returns a Match
func PlayMatch(points string) Match {

	var match Match
	match.Result = "hi"
	match.Scores = []string{"0", "15", "30", "40", "Deuce", "Ad"}
	match.Player1Score = []int{0, 0, 0, 0, 0}
	match.Player2Score = []int{0, 0, 0, 0, 0}
	match.PlayerGameScoreIndex = []int{0, 0}
	match.PlayerGameScore = []string{"0", "0"}

	for _, point := range points {
		fmt.Printf("%v", string(point))
		winningPlayerIndex := getWinningPlayerIndex(string(point))
		playPoint(winningPlayerIndex, match)
	}

	return match
}

func getWinningPlayerIndex(winningPlayer string) int {
	if winningPlayer == "1" {
		return 0
	}
	return 1
}

func playPoint(winningPlayerIndex int, match Match) {
	var losingPlayerIndex int
	if winningPlayerIndex == 0 {
		losingPlayerIndex = 1
	}

	// Check if deuce
	if match.PlayerGameScoreIndex[winningPlayerIndex] == 2 &&
		match.PlayerGameScoreIndex[losingPlayerIndex] == 3 {
		setDeuce(match)
		return
	}

	if match.PlayerGameScoreIndex[losingPlayerIndex] == 5 {
		setDeuce(match)
		return
	}

	// Check if advantage
	if match.PlayerGameScoreIndex[winningPlayerIndex] == 4 {
		setAdvantage(winningPlayerIndex, losingPlayerIndex, match)
		return
	}

	// Check if player wins game
	if match.PlayerGameScoreIndex[winningPlayerIndex] == 3 || match.PlayerGameScoreIndex[winningPlayerIndex] == 5 {
		playerWinsGame(winningPlayerIndex, match)
		return
	}

	// Normal scoring in game
	match.PlayerGameScoreIndex[winningPlayerIndex]++
	updatePlayerGameScore(match)
}

func updatePlayerGameScore(match Match) {
	match.PlayerGameScore[0] = match.Scores[match.PlayerGameScoreIndex[0]]
	match.PlayerGameScore[1] = match.Scores[match.PlayerGameScoreIndex[1]]
}

func playerWinsGame(winningPlayerIndex int, match Match) {
	if winningPlayerIndex == 0 {
		match.Player1Score[match.CurrentSetIndex]++

		if match.Player1Score[match.CurrentSetIndex] >= 6 &&
			match.Player2Score[match.CurrentSetIndex] <= match.Player1Score[match.CurrentSetIndex]-2 {
			match.CurrentSetIndex++
		}
	} else {
		match.Player2Score[match.CurrentSetIndex]++

		if match.Player2Score[match.CurrentSetIndex] >= 6 &&
			match.Player1Score[match.CurrentSetIndex] <= match.Player2Score[match.CurrentSetIndex]-2 {
			match.CurrentSetIndex++
		}
	}
	setNewGame(match)
}

func setNewGame(match Match) {
	match.PlayerGameScoreIndex[0] = 0
	match.PlayerGameScoreIndex[1] = 0
	updatePlayerGameScore(match)
}

func setDeuce(match Match) {
	match.PlayerGameScoreIndex[0] = 4
	match.PlayerGameScoreIndex[1] = 4
	updatePlayerGameScore(match)
}

func setAdvantage(winningPlayerIndex int, losingPlayerIndex int, match Match) {
	match.PlayerGameScoreIndex[winningPlayerIndex] = 5
	match.PlayerGameScoreIndex[losingPlayerIndex] = 3
	updatePlayerGameScore(match)
}

// Match is a match score struct
type Match struct {
	Result               string
	Scores               []string
	Player1Score         []int
	Player2Score         []int
	PlayerGameScoreIndex []int
	PlayerGameScore      []string
	CurrentSetIndex      int
}
