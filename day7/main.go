package day7

import (
	"../utils"
	"fmt"
	"strings"
)

type Score struct {
	position	int
	fuel		int
}

func Run() {
	inputData := utils.ReadFileInput("day7/input")
	inputNumbers := utils.ConvertLinesToInt(strings.Split(inputData[0], ","))
	max := getMaximum(inputNumbers)
	best := Score{fuel: 0, position: 0}
	for i := 0; i < max; i++ {
		if i == 0 || testDestination(inputNumbers, i) < best.fuel {
			best = Score{fuel: testDestination(inputNumbers, i), position: i}
		}
	}
	fmt.Println(best)
}

func getMaximum(inputData []int) int {
	max := 0
	for _, value := range inputData {
		if value > max {
			max = value
		}
	}
	return max
}

func testDestination(input []int, destination int) int {
	score := 0
	for _, value := range input {
		score += getFuelToPosition(value, destination)
	}
	return score
}

func getFuelToPosition(currentPosition int, destination int) int  {
	if currentPosition == destination {
		return 0
	}
	max := Abs(currentPosition - destination)
	total := 0
	for i := 0; i < max; i++ {
		total = total + i + 1
	}
	return total
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
