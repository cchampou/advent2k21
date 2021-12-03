package main

import (
	"../utils"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	depth int
	horizontal int
	aim int
}

func main() {
	var currentPosition = Position{ depth: 0, horizontal: 0 }

	instructions := splitLines(utils.ReadFileInput())

	for _, currentInstruction := range instructions{
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

func splitLines(input []string) [][]string {
	var output [][]string
	for _, currentValue := range input {
		output = append(output, strings.Split(currentValue, " "))
	}
	return output
}