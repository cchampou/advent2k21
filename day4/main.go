package day4

import (
	"../utils"
	"fmt"
)

func Run() {
	inputData := utils.ReadFileInput("day4/input")
	totalCalled := GenerateCalledNumbers(inputData[0])
	boards := generateBoards(inputData[2:])
	var scoreBoard []Score

	for index, _ := range totalCalled {
		called := totalCalled[:index]
		for boardIndex, currentBoard := range boards {
			if isBoardWinning(called, currentBoard) {
				scoreBoard = hasWon(scoreBoard, boardIndex, getBoardScore(called, currentBoard))
			}
		}
	}
	fmt.Println(scoreBoard[0].score)
	fmt.Println(scoreBoard[len(scoreBoard) - 1].score)
}

func isBoardWinning(called []int, board [][]int) bool {
	// check by line
	for _, line := range board {
		isWinning := true
		for _, number := range line {
			if !isNumberCalled(called, number) {
				isWinning = false
			}
		}
		if isWinning {
			return true
		}
	}
	// check by column
	for currentColumnIndex := range board[0] {
		isWinning := true
		for currentLineIndex := range board {
			if !isNumberCalled(called, board[currentLineIndex][currentColumnIndex]) {
				isWinning = false
			}
		}
		if isWinning {
			return true
		}
	}

	return false
}

func isNumberCalled(called []int, number int) bool {
	for _, calledNumber := range called {
		if number == calledNumber {
			return true
		}
	}
	return false
}

func getBoardScore(called []int, board [][]int) int {
	total := 0
	for _, line := range board {
		for _, number := range line {
			if !isNumberCalled(called, number) {
				total += number
			}
		}
	}
	return total * called[len(called) - 1]
}

func generateBoards(input []string) [][][]int {
	var currentBoard []string
	var separatedBoards [][]string
	for _, currentLine := range input {
		if len(currentLine) < 1 {
			separatedBoards = append(separatedBoards, currentBoard)
			currentBoard = []string{}
		} else {
			currentBoard = append(currentBoard, currentLine)
		}
	}
	separatedBoards = append(separatedBoards, currentBoard)
	return convertBoards(separatedBoards)
}

func convertBoards(input [][]string) [][][]int {
	var result [][][]int
	for _, currentBoard := range input {
		var newBoard [][]int
		for _, currentLine := range currentBoard {
			newBoard = append(newBoard, convertLineToInts(currentLine))
		}
		result = append(result, newBoard)
	}
	return result
}