package tennisscoring

import "testing"

func TestTennisScoring(t *testing.T) {
	res := PlayMatch("111")
	if res.Result != "hi" {
		t.Errorf("error")
	}
}

func TestTennisScoringPlayMatchPlayerOneWinsThreePointFourty(t *testing.T) {
	res := PlayMatch("111")

	if res.PlayerGameScore[0] != "40" {
		t.Errorf("Player score was %v should be 40", res.PlayerGameScore[0])
	}
}

func TestTennisScoringPlayMatchPlayerTwoWinsThreePointFourty(t *testing.T) {
	res := PlayMatch("11122")

	if res.PlayerGameScore[0] != "40" {
		t.Errorf("Player score was %v should be 40", res.PlayerGameScore[0])
	}

	if res.PlayerGameScore[1] != "30" {
		t.Errorf("Player score was %v should be 30", res.PlayerGameScore[1])
	}
}

func TestTennisScoringPlayMatchPlayerOneWinsFirstGame(t *testing.T) {
	res := PlayMatch("111221")

	if res.PlayerGameScore[0] != "0" {
		t.Errorf("Player score was %v should be 0", res.PlayerGameScore[0])
	}

	if res.Player1Score[0] != 1 {
		t.Errorf("Player one games score was %v but should be 1", res.Player1Score[0])
	}

	if res.PlayerGameScore[1] != "0" {
		t.Errorf("Player score was %v should be 0", res.PlayerGameScore[1])
	}
}

func TestTennisScoringPlayMatchDeuce(t *testing.T) {
	res := PlayMatch("121212")

	if res.PlayerGameScore[0] != "Deuce" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[0])
	}

	if res.PlayerGameScore[1] != "Deuce" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[1])
	}
}

func TestTennisScoringPlayMatchAdvantage(t *testing.T) {
	res := PlayMatch("1212121")

	if res.PlayerGameScore[0] != "Ad" {
		t.Errorf("Player score was %v should be Ad", res.PlayerGameScore[0])
	}

	if res.PlayerGameScore[1] != "40" {
		t.Errorf("Player score was %v should be 40", res.PlayerGameScore[1])
	}
}

func TestTennisScoringPlayMatchBackToDeuceAfterAdvantage(t *testing.T) {
	res := PlayMatch("12121212")

	if res.PlayerGameScore[0] != "Deuce" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[0])
	}

	if res.PlayerGameScore[1] != "Deuce" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[1])
	}
}

func TestTennisScoringPlayMatchPlayerWinsSet(t *testing.T) {
	game1 := "1212121211" // p1
	game2 := "2222"       // p2
	game3 := "121211"     // p1
	game4 := "121211"     // p1
	game5 := "121211"     // p1
	game6 := "121211"     // p1
	game7 := "121211"     // p1

	res := PlayMatch(game1 + game2 + game3 + game4 + game5 + game6 + game7)

	if res.PlayerGameScore[0] != "0" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[0])
	}

	if res.PlayerGameScore[1] != "0" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[1])
	}

	if res.Player1Score[0] != 6 {
		t.Errorf("Player 1 set score was %v should have been 6", res.Player1Score[0])
	}

	if res.Player2Score[0] != 1 {
		t.Errorf("Player 2 set score was %v should have been 1", res.Player1Score[0])
	}
}

func TestTennisScoringPlayMatchPlayerWinsSet75(t *testing.T) {
	game1 := "1212121211" // p1
	game2 := "2222"       // p2
	game3 := "121211"     // p1
	game4 := "121211"     // p1
	game5 := "121211"     // p1
	game6 := "121211"     // p1
	game7 := "2222"       // p2
	game8 := "2222"       // p2
	game9 := "2222"       // p2
	game10 := "2222"      // p2
	game11 := "121211"    // p1
	game12 := "121211"    // p1
	game13 := "121211"    // p1

	res := PlayMatch(game1 + game2 + game3 + game4 + game5 + game6 + game7 + game8 + game9 + game10 + game11 + game12 + game13)

	if res.PlayerGameScore[0] != "0" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[0])
	}

	if res.PlayerGameScore[1] != "0" {
		t.Errorf("Player score was %v should be Deuce", res.PlayerGameScore[1])
	}

	if res.Player1Score[0] != 7 {
		t.Errorf("Player 1 set score was %v should have been 7", res.Player1Score[0])
	}

	if res.Player2Score[0] != 5 {
		t.Errorf("Player 2 set score was %v should have been 5", res.Player1Score[0])
	}

	if res.Player1Score[1] != 1 {
		t.Errorf("Player 1 set 2 score was %v should have been 1", res.Player1Score[1])
	}
}
