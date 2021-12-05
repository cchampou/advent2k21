package day5

import (
	"../utils"
	"fmt"
)

func Run() {
	input := utils.ReadFileInput("day5/input")
	parsedData := parseInput(input)
	vectors := parsedData
	grid := initGrid(vectors)
	grid = draw(vectors, grid)
	printGrid(grid)
	fmt.Println(countScore(grid))
}