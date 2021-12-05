package day5

import (
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x		int
	y		int
}

type Vector struct {
	start	Point
	end		Point
}

func parseInput(input []string) []Vector  {
	var output []Vector
	for _, line := range input {
		output = append(output, parseLine(line))
	}
	return output
}

func parseLine(line string) Vector {
	split := strings.Split(line, " -> ")
	return Vector{start: parsePoint(split[0]), end: parsePoint(split[1])}
}

func parsePoint(input string) Point {
	split :=  strings.Split(input, ",")
	xAsInt, errOnX := strconv.Atoi(split[0])
	yAsInt, errOnY := strconv.Atoi(split[1])
	if errOnX != nil || errOnY != nil {
		log.Fatalf("Fail to parse point")
	}
	return Point{x: xAsInt, y: yAsInt}
}


