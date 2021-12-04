package day4

type Score struct {
	boardIndex	int
	score		int
}

func hasWon(scoreBoard []Score, boardIndex int, score int) []Score {
	for _, currentScore := range scoreBoard {
		if currentScore.boardIndex == boardIndex {
			return scoreBoard
		}
	}
	return append(scoreBoard, Score{score: score, boardIndex: boardIndex})
}
