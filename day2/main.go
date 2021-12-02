package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth int
	horizontal int
	aim int
}

func main() {
	fmt.Println("Program started")

	var instructions [][]string
	var currentPosition = Position{ depth: 0, horizontal: 0 }

	f, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Fail to open the input file")
		return
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		text := sc.Text()
		instructions = append(instructions, strings.Split(text, " "))
	}

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
	fmt.Println(currentPosition)
	fmt.Println(currentPosition.depth * currentPosition.horizontal)
}