package day5

import (
	"../utils"
	"fmt"
)

func Run() {
	fmt.Println("Hello world")
	input := utils.ReadFileInput("day5/input")
	fmt.Println(parseInput(input))
}