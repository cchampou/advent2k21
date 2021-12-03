package day2

import (
	"../utils"
	"fmt"
	"strconv"
)

type Position struct {
	depth      int
	horizontal int
	aim        int
}

func Run() {
	var currentPosition = Position{depth: 0, horizontal: 0}

	instructions := utils.SplitLines(utils.ReadFileInput("day2/input"))

	for _, currentInstruction := range instructions {
		direction := currentInstruction[0]
		value, _ := strconv.Atoi(currentInstruction[1])
		switch direction {
		case "forward":
			currentPosition.horizontal += value
			currentPosition.depth += currentPosition.aim * value
		case "down":
			currentPosition.aim += value
		case "up":
			currentPosition.aim -= value
		}
	}
	fmt.Println(currentPosition.depth * currentPosition.horizontal)
}
