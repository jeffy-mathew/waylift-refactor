package tenniskata

type tennisGame1 struct {
	player1Score int
	player2Score int
	player1Name  string
	player2Name  string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.player1Score += 1
	} else {
		game.player2Score += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.player1Score == game.player2Score {
		switch game.player1Score {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.player1Score >= 4 || game.player2Score >= 4 {
		minusResult := game.player1Score - game.player2Score
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.player1Score
			} else {
				score += "-"
				tempScore = game.player2Score
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
