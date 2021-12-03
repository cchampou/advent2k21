package utils

import (
	"log"
	"strconv"
	"strings"
)

func ConvertLinesToInt(raw []string) []int {
	var output []int
	for _, currentValue := range raw {
		newMeasure, err := strconv.Atoi(currentValue)
		if err != nil {
			log.Fatalf("Line could not be converted to number: %v", currentValue)
		}
		output = append(output, newMeasure)
	}
	return output
}

func SplitLines(input []string) [][]string {
	var output [][]string
	for _, currentValue := range input {
		output = append(output, strings.Split(currentValue, " "))
	}
	return output
}